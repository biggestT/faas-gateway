FROM golang

WORKDIR /app

COPY . .

RUN ./cmd/gendata.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go

EXPOSE 8080
ENTRYPOINT ["/app/main"]
