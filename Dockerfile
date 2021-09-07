FROM alpine:latest as builder

RUN apk --no-cache --update add ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY trumpet /trumpet

EXPOSE 5353
CMD ["./trumpet", "server"]
