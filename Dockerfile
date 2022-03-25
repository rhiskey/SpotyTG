# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get github.com/rhiskey/spotytg/auths
RUN go get github.com/rhiskey/spotytg/spotifydl

COPY *.go ./

RUN go build -o /spotytg

CMD [ "/spotytg" ]