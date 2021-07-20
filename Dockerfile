FROM golang:alpine as builder

RUN apk add --no-cache git gcc libc-dev
WORKDIR /go/src/app
# RUN git clone https://github.com/Mondongo-cl/http-rest-echo-go.git
## RUN apk add go
COPY . .
WORKDIR /go/src/app/src
RUN go build .
RUN go get -v ./... 
RUN go install -v .

# final stage
FROM alpine:latest
## FROM golang:latest
LABEL Name=http-rest-echo-go Version=1.0
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/* /usr/local/bin/
### /usr/local/bin/
## RUN chmod 777 /usr/local/bin/http-rest-echo-go
ENV dbusername=root
ENV dbpassword=123456
ENV dbhostname=localhost
ENV dbport=3306
ENV httplistenerport=9000
ENV databasename=testdb
ENTRYPOINT ["sh","-c","http-rest-echo-go -databasename ${databasename} -dbusername ${dbusername} -dbpassword ${dbpassword} -dbhostname ${dbhostname} -dbport ${dbport} -httplistenerport ${httplistenerport}"]
##ENTRYPOINT [ "http-rest-echo-go" ]
EXPOSE ${httplistenerport}

## to start container 
## docker run -d -p  5001:5001 http-rest-echo-go
