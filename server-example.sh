#!/usr/bin/env bash

# Install gen
go get -u -v github.com/wzshiming/gen/cmd/gen

# API v1
gen route github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service
gen client -o ./bookstore/v1/client/client_gen.go github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service
gen openapi -o ./bookstore/v1/openapi/openapi.json github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service

# Run a testing server.
gen run github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service


# API v2
gen route github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
gen client -o ./bookstore/v2/client/client_gen.go github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
gen openapi -o ./bookstore/v2/openapi/openapi.json github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service

# Run a testing server.
gen run github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
