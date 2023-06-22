## Build the Go API
##
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY ./src/go.mod ./

COPY ./src/go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./city_api ./src/main.go


## Production image
##
FROM alpine:latest AS prod

WORKDIR /

COPY --from=builder /app/city_api .

EXPOSE 3000

CMD [ "/city_api" ]  