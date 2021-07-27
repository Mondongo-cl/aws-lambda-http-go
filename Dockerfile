FROM golang:latest as builder
WORKDIR /go/src/app
COPY . .
WORKDIR /go/src/app/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
RUN go get -v ./... 
RUN go install -v .
ENTRYPOINT [ "sh" ]
# final stage
 FROM alpine:latest
# 
LABEL Name=http-rest-echo-go Version=1.0
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/src/http-rest-echo-go /usr/local/bin/
ENV dbusername=root
ENV dbpassword=123456
ENV dbhostname=localhost
ENV dbport=3306
ENV httplistenerport=9000
ENV databasename=testdb
ENTRYPOINT ["sh","-c","http-rest-echo-go -databasename ${databasename} -dbusername ${dbusername} -dbpassword ${dbpassword} -dbhostname ${dbhostname} -dbport ${dbport} -httplistenerport ${httplistenerport}"]
EXPOSE ${httplistenerport}
# to start container 
# docker run -d -p  5001:5001 http-rest-echo-go
