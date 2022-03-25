                                         # syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18

WORKDIR /app

ADD . /app

RUN go build -o main .

CMD [ "/app/main" ]
