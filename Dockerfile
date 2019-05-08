FROM alpine
COPY gopath/bin/app /app
RUN apk --update add ca-certificates
EXPOSE 50051/tcp
EXPOSE 8080/tcp
ENTRYPOINT ["/app"]
