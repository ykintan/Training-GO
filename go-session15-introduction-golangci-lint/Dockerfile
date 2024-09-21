FROM golang:1.22.1-alpine AS build

#set destination for copy
WORKDIR /app

#copy resource code
COPY . .

#build the Go Application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-hello-world .

#stage2 - run the stage
FROM alpine:edge

#set the working directory
WORKDIR /app

#copy the binary from the build stage
COPY --from=build /app/go-hello-world .

#optiomal: expose port
EXPOSE 8080

#run
CMD [ "./go-hello-world" ]