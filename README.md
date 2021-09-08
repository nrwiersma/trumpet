![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=trumpet&tag=dns%20service%20discovery)

[![Go Report Card](https://goreportcard.com/badge/github.com/nrwiersma/trumpet)](https://goreportcard.com/report/github.com/nrwiersma/trumpet)
[![Build Status](https://github.com/nrwiersma/trumpet/actions/workflows/test.yml/badge.svg)](https://github.com/nrwiersma/trumpet/actions)
[![Coverage Status](https://coveralls.io/repos/github/nrwiersma/trumpet/badge.svg?branch=master)](https://coveralls.io/github/nrwiersma/trumpet?branch=master)
[![GitHub release](https://img.shields.io/github/release/nrwiersma/trumpet.svg)](https://github.com/nrwiersma/trumpet/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/nrwiersma/trumpet/master/LICENSE)

## About

Trumpet is a DNS Service Discovery server

## Usage

Start the binary manually:

```shell
./trumpet server --services=example/config.yaml
```

or with docker

```shell
docker run --rm -v example/config.yaml:/config.yaml -e SERVICES=/config.yaml ghcr.io/nrwiersma/trumpet:v0.1.0
```
