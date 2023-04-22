FROM golang:1.15.7-buster

COPY go.mod go.sum /go/src/elkWorkshop/
WORKDIR /go/src/elkWorkshop
RUN go mod download
COPY . /go/src/elkWorkshop
RUN go build -o /usr/bin/elkWorkshop elkWorkshop/cmd/api

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/elkWorkshop"]