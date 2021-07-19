FROM public.ecr.aws/lambda/provided:al2 as builder
# install GIT
WORKDIR /go/src/app
# RUN git clone https://github.com/Mondongo-cl/http-rest-echo-go.git
COPY . ./http-rest-echo-go/
WORKDIR /go/src/app/http-rest-echo-go/src
RUN go build .
RUN go get -v ./... 
run go install -v .
# final stage
FROM alpine:latest
LABEL Name=http-rest-echo-go Version=1.2
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/http-rest-echo-go/src/http-rest-echo-go /usr/local/bin/

ENV dbusername=root
ENV dbpassword=123456
ENV dbhostname=localhost
ENV dbport=3306
ENV httplistenerport=9000
ENV databasename=testdb
ENTRYPOINT ["sh","-c","http-rest-echo-go -databasename ${databasename} -dbusername ${dbusername} -dbpassword ${dbpassword} -dbhostname ${dbhostname} -dbport ${dbport} -httplistenerport ${httplistenerport}"]

EXPOSE ${httplistenerport}

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
