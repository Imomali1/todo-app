# Build stage
FROM --platform=linux/x86_64 golang:alpine AS builder

RUN mkdir oseas
WORKDIR /oseas
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Run stage
FROM --platform=linux/x86_64 alpine:latest

COPY --from=builder /oseas/main .

EXPOSE 8050
CMD ["/main"]