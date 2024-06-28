FROM golang:1.22

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

COPY core .
COPY go.work .
COPY authentication . 

RUN go mod tidy

RUN go build -v -o /usr/local/bin/app .

EXPOSE 3000

CMD ["app"]