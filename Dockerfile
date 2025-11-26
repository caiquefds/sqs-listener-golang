FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app .

FROM scratch

COPY --from=builder /app/app /app/app

WORKDIR /app

ENTRYPOINT ["./app"]
