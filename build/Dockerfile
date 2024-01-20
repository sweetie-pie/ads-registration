# choosing a base image
FROM golang:1.20-alpine

# entring app directory
WORKDIR /app

# copy files
COPY go.mod go.sum ./

# download go modules
RUN go mod download

# copy other files
COPY internal/ internal/
COPY main.go main.go

# build go app
RUN go build -o main

# execute go app
CMD ./main