FROM golang:latest

RUN export GO111MODULE="on"

WORKDIR /app

COPY . /app

RUN go mod download
RUN go build main.go

EXPOSE 1883

CMD [ "./main" ]