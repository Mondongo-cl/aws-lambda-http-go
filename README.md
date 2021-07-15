# AWS Lambda http golang Function

## Build

### Command 
``` powershell
 go build
```
### Results
```
no visible results, a new executable file is created
```

## Run

### Command 
``` powershell
 go run .
```
### Results
```

allow access to the executable to network and then, the service will print out to the console  the text:

starting hello world service...
Press any key to close ...
```

## Test

### Command 
``` 
 go test .\dataaccess\ .\business\
```
### Results
```
ok      github.com/Mondongo-cl/http-rest-echo-go/dataaccess     0.079s  coverage: 86.7% of statements
ok      github.com/Mondongo-cl/http-rest-echo-go/business       0.092s  coverage: 10.1% of statements
```

## Code Coverage

### Command 

```
 go test -cover .\dataaccess\ .\business\
 ```
### Results
```
ok      github.com/Mondongo-cl/http-rest-echo-go/dataaccess     (cached)        coverage: 86.7% of statements
ok      github.com/Mondongo-cl/http-rest-echo-go/business       0.101s  coverage: 10.1% of statements
<div>
```
</div>
 
## Code Coverage Reports

### Command 
``` powershell
  go test -coverprofile cover.out  .\dataaccess\ . \business\

  go tool cover -html cover.out -o coverage.html

  start coverage.html
```

### Results
``` 
the command generate a html with detailed coverage report page
```
