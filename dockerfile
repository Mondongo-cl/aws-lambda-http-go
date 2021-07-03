# Build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /go/src/app
COPY ./src/ .
# RUN go mod edit -module http-rest-echo-go
RUN go build .
RUN go get -d -v ./...
RUN go install -v ./...

# final stage
FROM alpine:latest
LABEL Name=appName Version=0.0.1
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/http-rest-echo-go /app
ENTRYPOINT ./app
EXPOSE 5001

## to start container 
## docker run -it -p  5001:5001 http-rest-echo-go