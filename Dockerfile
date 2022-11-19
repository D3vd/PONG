FROM golang:1.19 as build

WORKDIR /src/pong

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/pong

# Final Stage
FROM ubuntu:latest

WORKDIR /app

COPY --from=build /app/pong /app/

ENV PORT 80
ENV ENV prod

EXPOSE 80

CMD /app/pong
