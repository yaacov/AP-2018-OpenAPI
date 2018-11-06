# AP-2018-OpenAPI
Demo for August Penguin 2018 OpenAPI talk

## Links

  - https://www.openapis.org/
  - https://github.com/OAI/OpenAPI-Specification
  - https://github.com/OAI/OpenAPI-Specification/blob/master/IMPLEMENTATIONS.md#implementations
  - https://apis.guru/awesome-openapi3/
  - https://github.com/wzshiming/gen

## Book store

#### Interactive UI

``` bash
# Test api v2
docker run --name openapi -p 8080:8080 -e SWAGGER_JSON=/api/api-v2.yaml -v $(pwd)/bookstore:/api:Z swaggerapi/swagger-ui
```

#### Server

``` bash
# Test server v2
gen route github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
gen client -o ./bookstore/v2/client/client_gen.go github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
gen openapi -o ./bookstore/v2/openapi/openapi.json github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service

# Run a testing server.
gen run github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
```
