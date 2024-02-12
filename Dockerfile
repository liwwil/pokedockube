
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pokedoc-app .
RUN CGO_ENABLED=0 GOOS=linux go build -o pokedoc-app .

FROM scratch

WORKDIR /app
COPY --from=builder /app/pokedoc-app .
CMD ["./pokedoc-app"]