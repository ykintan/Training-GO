FROM golang:1.22.1-alpine AS build

#install neccessary packages
RUN apk add --no-cache git

#set destination for copy
WORKDIR /app

#copy the Go modules manifest and download dependencies
COPY go.mod  go.sum ./
RUN go mod download

#copy resource code
COPY . .

#install go lint 
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

#run the linters with increase timeout
RUN golangci-lint run --timeout 5m

#build the Go Application
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-hello-world .

#stage2 - run the stage
FROM alpine:edge

#set the working directory
WORKDIR /app

#copy the binary from the build stage
COPY --from=build /app/go-crud-user .

#install cA certifikat
run apk --no-cache add ca-certificates

#run
CMD [ "./go-hello-world" ]