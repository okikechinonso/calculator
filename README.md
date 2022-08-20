# Simple Calculator



Run ``go mod tidy`` to download all dependencies

Run `go run ./...` to start the project

Run `go test ./...` for tests

Run on terminal using
```azure
    curl --header "Content-Type: application/json" \
--request POST \
--data '{"numerator":7,"denominator":56}' \
http://localhost:5550/api/v1/division
```
Run on postman:  
Link: `localhost:5550/api/v1/division`  
payload:
```azure
{
    "numerator": 7,
    "denominator": 56
}
```


