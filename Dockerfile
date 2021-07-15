FROM public.ecr.aws/lambda/go:1 as builder
# install GIT
WORKDIR /go/src/app
RUN yum install -y git
RUN git clone https://github.com/Mondongo-cl/http-rest-echo-go.git
RUN yum install -y golang
RUN go env -w GOPROXY=direct
WORKDIR /go/src/app/http-rest-echo-go/src
RUN go build .
# RUN go get -d -v ./...
# RUN go install -v ./...


# copy artifacts to a clean image
FROM public.ecr.aws/lambda/go:1
LABEL Name=http-rest-echo-go Version=0.2
COPY --from=builder /go/src/app/http-rest-echo-go/src/http-rest-echo-go ${LAMBDA_TASK_ROOT}
CMD ["http-rest-echo-go"]
EXPOSE 5001

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
