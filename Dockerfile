FROM golang:1.19-alpine

RUN apk update && apk add git

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/hotel-booking-api

COPY . .

RUN go mod init hotel-booking-api
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17

WORKDIR cmd
RUN GOOS=linux go build -o app

ENTRYPOINT ["./app"]

EXPOSE 3000