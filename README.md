# API to validate IBAN number
Supports an API to check if the given IBAN number is valid
# Run Webserver
```go
go run main.go
```
# URL to probe 
```go
127.0.0.1:8080/validate?iban=DE89 3704 0044 0532 0130 00
```
# Expected Responses

```go
{
    "status": 400,
    "error": "Given IBAN is not valid: ABCDEFGHIJK"
}
```
```go
{
    "status": 200,
    "data": "Given IBAN is valid"
}
```
# API definition
One can open and copy paste the contents of `api.yaml` to https://editor.swagger.io/ in browser 
