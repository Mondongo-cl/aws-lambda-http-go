FROM public.ecr.aws/lambda/provided:al2 as build
# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct
RUN yum install -y git

RUN git clone https://github.com/Mondongo-cl/http-rest-echo-go.git
WORKDIR /go/src/app/http-rest-echo-go/src
RUN go build .
RUN go get -d -v ./...
RUN go install -v ./...


# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
LABEL Name=http-rest-echo-go Version=0.2
COPY --from=builder /go/bin/http-rest-echo-go /app
# ENTRYPOINT ./app
# EXPOSE 5001
ENTRYPOINT [ "/app" ]           

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
