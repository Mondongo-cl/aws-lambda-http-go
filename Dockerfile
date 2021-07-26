FROM golang:latest as builder
# install GIT
WORKDIR /go/src/app
COPY . ./http-rest-echo-go/
# RUN yum install -y git
# RUN git clone https://github.com/Mondongo-cl/http-rest-echo-go.git
# RUN yum install -y golang
# RUN go get -d -v ./...

RUN go env -w GOPROXY=direct
WORKDIR /go/src/app/http-rest-echo-go/src
RUN GOOS=linux go build .
# RUN go get -d -v ./...
# RUN go install -v ./...

FROM scratch as test
WORKDIR /go/src/app/
COPY --from=builder /go/src/app/http-rest-echo-go .
RUN src/app/http-rest-echo-go

# copy artifacts to a clean image
# FROM public.ecr.aws/lambda/go:1
FROM public.ecr.aws/lambda/provided:latest
LABEL Name=http-rest-echo-go Version=0.2
ENV GOPROXY=direct
ENV _LAMBDA_SERVER_PORT=5001
ENV _HANDLER=http-rest-echo-go
ENV servername=
ENV serverport=
ENV username=
ENV password=
ENV database=
COPY --from=builder /go/src/app/http-rest-echo-go/src/http-rest-echo-go ${LAMBDA_TASK_ROOT}
# ARG servername=${servername}
# ARG serverport=${serverport}
# ARG username=${username}
# ARG password=${password}
# ARG database=${database}
CMD [ "http-rest-echo-go"]
# EXPOSE 5001

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
