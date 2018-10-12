FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/ethanduncan/trumail-go-api
RUN go install -v github.com/ethanduncan/trumail-go-api
RUN go build -o trumail-go-api

ENTRYPOINT [ "./trumail-go-api" ]