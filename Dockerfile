FROM golang:alpine AS builder

WORKDIR usr/src/houseService

#dependencies
COPY go.mod go.sum ./
RUN go mod download

#build
COPY . .
RUN go build -o /usr/local/bin/houseService cmd/main.go

FROM alpine AS runner
COPY --from=builder /usr/local/bin/houseService /
COPY .env /.env

CMD ["/houseService"]