##
## STEP 1
##

# Specify the base image for building the application
FROM golang:1.20-alpine AS build

# Create a working directory
WORKDIR /app

# Copy directory files
COPY . .

# Build application
RUN go build -ldflags="-s -w" -trimpath -o products-api

##
## STEP 2 
##
FROM alpine

WORKDIR /

# Move Deps from build image
COPY --from=build /app/products-api /products-api

RUN echo DB_USER=$DB_USER >> /.env
RUN echo DB_PASS=$DB_PASS >> /.env
RUN echo DB_PORT=$DB_PORT >> /.env
RUN echo DB_NAME=$DB_NAME >> /.env

EXPOSE 8080

ENTRYPOINT ["/products-api"]