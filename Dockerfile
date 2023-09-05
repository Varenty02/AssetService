FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main
  # Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
ENV CENSYS_API_ID=7f5dba87-fe88-4863-b4b8-650cef0742a2
ENV CENSYS_API_SECRET=TDRgPPXKlcjQ384vHI7NlyqNT1lA72Oq
EXPOSE 3009
CMD [ "/app/main" ]