# Build stage
FROM --platform=linux/x86_64 golang:alpine AS builder

RUN mkdir todo
WORKDIR /todo
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/cmd .

# Run stage
FROM --platform=linux/x86_64 alpine:latest

COPY --from=builder /todo/bin/cmd .

EXPOSE 8080
CMD ["/cmd"]