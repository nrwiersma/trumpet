package main

import (
	"fmt"
	stdlog "log"
	"net"
	"os"
	"path/filepath"

	"github.com/brutella/dnssd"
	"github.com/brutella/dnssd/log"
	"github.com/hamba/cmd/v2"
	"github.com/hamba/cmd/v2/term"
	"github.com/hamba/logger/v2"
	lctx "github.com/hamba/logger/v2/ctx"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func runServer(_ term.Term) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		ctx := c.Context

		l, err := cmd.NewLogger(c)
		if err != nil {
			return err
		}

		cfg, err := parseConfig(c.String(flagService))
		if err != nil {
			return err
		}

		resp, err := createResponder(cfg, l)
		if err != nil {
			return err
		}

		l.Info("Starting server", lctx.Int("port", 5353))
		if err = resp.Respond(ctx); err != nil {
			l.Error("Responder exited with an error", lctx.Error("error", err))
		}

		return nil
	}
}

type config struct {
	Groups []struct {
		Name     string `yaml:"name"`
		Services []struct {
			Type   string            `yaml:"type"`
			Domain string            `yaml:"domain"`
			Host   string            `yaml:"host"`
			IPs    []string          `yaml:"ips"`
			Port   int               `yaml:"port"`
			TXT    map[string]string `yaml:"txt"`
		} `yaml:"services"`
	} `yaml:"groups"`
}

func parseConfig(f string) (config, error) {
	var cfg config

	b, err := os.ReadFile(filepath.Clean(f))
	if err != nil {
		return cfg, fmt.Errorf("could not read service configuration: %w", err)
	}

	if err = yaml.Unmarshal(b, &cfg); err != nil {
		return cfg, fmt.Errorf("badly formatted service configuration: %w", err)
	}
	return cfg, err
}

func createResponder(cfg config, l *logger.Logger) (dnssd.Responder, error) {
	log.Debug = &log.Logger{Logger: stdlog.New(l.Writer(logger.Debug), "", 0)}
	log.Info = &log.Logger{Logger: stdlog.New(l.Writer(logger.Info), "", 0)}

	resp, err := dnssd.NewResponder()
	if err != nil {
		return nil, fmt.Errorf("could not create responder: %w", err)
	}

	for _, grp := range cfg.Groups {
		for _, svc := range grp.Services {
			ips := make([]net.IP, 0, len(svc.IPs))
			for _, ip := range svc.IPs {
				v := net.ParseIP(ip)
				if v == nil {
					return nil, fmt.Errorf("invalid ip address %q", ip)
				}
				ips = append(ips, v)
			}

			srv, err := dnssd.NewService(dnssd.Config{
				Name:   grp.Name,
				Type:   svc.Type,
				Domain: svc.Domain,
				Host:   svc.Host,
				IPs:    ips,
				Port:   svc.Port,
			})
			if err != nil {
				return nil, fmt.Errorf("invalid service %q: %w", svc.Type, err)
			}

			hdl, err := resp.Add(srv)
			if err != nil {
				return nil, fmt.Errorf("could not add service: %w", err)
			}

			if len(svc.TXT) > 0 {
				hdl.UpdateText(svc.TXT, resp)
			}

			l.Info("Added service", lctx.Str("name", grp.Name), lctx.Str("type", svc.Type))
		}
	}

	return resp, nil
}
