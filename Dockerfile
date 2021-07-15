# Build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /go/src/app
RUN go env -w GOPROXY=direct
RUN git clone https://github.com/Mondongo-cl/http-rest-echo-go.git
WORKDIR /go/src/app/http-rest-echo-go/src
RUN go build .
RUN go get -d -v ./...
RUN go install -v ./...

# final stage
FROM alpine:latest
LABEL Name=http-rest-echo-go Version=0.2
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/http-rest-echo-go /app
ENTRYPOINT ./app
# EXPOSE 5001

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
