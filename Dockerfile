# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

# --- Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# ---

ADD . /app

RUN go build -o main .

CMD [ "/app/main" ]
