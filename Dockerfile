FROM        golang
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
WORKDIR     /app/cmd/word-card-api
RUN         go build -o app
ENTRYPOINT  ["./app"]
EXPOSE      8080