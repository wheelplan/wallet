FROM golang:1.20rc1-alpine3.17

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./... && ls -a | egrep -v "^\.$|^\.\.$" | xargs rm -rf

CMD ["app"]
