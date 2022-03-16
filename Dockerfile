FROM alpine:latest  
RUN apk --no-cache add ca-certificates
## non privileged user
USER 1000
# EXPOSE 9000
WORKDIR /app/
COPY goreleaser-hello-world /app/goreleaser-hello-world

ENTRYPOINT ["/app/goreleaser-hello-world"]