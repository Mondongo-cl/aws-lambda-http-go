#FROM  public.ecr.aws/lambda/provided:latest as builder
FROM  golang:latest as builder
WORKDIR /go/src/app
COPY . ./http-rest-echo-go/
RUN go env -w GOPROXY=direct
WORKDIR /go/src/app/http-rest-echo-go/src
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 GO_EXTLINK_ENABLED=0 go build .
# RUN go get -d -v ./...
# RUN go install -v ./...

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/go:1
#FROM public.ecr.aws/lambda/provided:latest
LABEL Name=http-rest-echo-go Version=0.2
ENV GOPROXY=direct
ENV _HANDLER=http-rest-echo-go
ENV AWS_LAMBDA_RUNTIME_API=go.1x
ENV servername=
ENV serverport=
ENV username=
ENV password=
ENV database=
ENV _LAMBDA_SERVER_PORT=9000
COPY --from=builder /go/src/app/http-rest-echo-go/src/http-rest-echo-go ${LAMBDA_TASK_ROOT}
CMD ["http-rest-echo-go"]

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
