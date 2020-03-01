# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:latest as builder
LABEL maintainer="Jeehong Lee <hongsgo@gmail.com>"
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY static ./static
COPY template/template/html ./template/html
# Expose port 8080 to the outside world
EXPOSE 1323
# Command to run the executable
CMD ["./main"]