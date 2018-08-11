# build stage
FROM golang:1.11beta3-alpine3.8 as build-stage
WORKDIR /app
COPY ./main.go /app/
RUN mkdir build && go build -o build/what-is-my-ip

# production stage
FROM alpine:3.8 as production-stage
COPY --from=build-stage /app/build /usr/bin/
EXPOSE 8090
CMD ["/usr/bin/what-is-my-ip"]