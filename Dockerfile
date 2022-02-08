FROM  gcr.io/distroless/static:nonroot

COPY trumpet /trumpet

EXPOSE 5353
CMD ["./trumpet", "server"]
