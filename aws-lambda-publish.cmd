@echo off
cd src
set GOOS=linux
set CGO_ENABLED=0 
set GOARCH=amd64
echo Building...
go mod tidy
go build -o ../app
cd ..
echo upload function
tar -a -c -f app.zip app
del app.
aws lambda update-function-code --function-name echo-zip-fun --zip-file fileb://app.zip 
echo update function configuration
aws lambda update-function-configuration  --function-name echo-zip-fun --handler app --description "echo golang bin file" --environment "Variables={servername=www.cyberpojos.com,serverport=3306,username=root,password=acemq3306,database=testdb}" 
del app.zip