FROM golang:1.21.4-alpine

WORKDIR /app

COPY . .

RUN go get -u github.com/cosmtrek/air@v1.49.0 \
    && go install github.com/cosmtrek/air \
    && go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]