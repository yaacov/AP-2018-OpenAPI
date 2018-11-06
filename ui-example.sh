#!/usr/bin/env bash

# docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' openapi
# docker rm openapi 

# Test api v1
docker run --name openapi -p 8080:8080 -e SWAGGER_JSON=/api/api-v1.yaml -v $(pwd)/bookstore:/api:Z swaggerapi/swagger-ui

# Test api v2
docker run --name openapi -p 8080:8080 -e SWAGGER_JSON=/api/api-v2.yaml -v $(pwd)/bookstore:/api:Z swaggerapi/swagger-ui
