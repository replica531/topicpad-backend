FROM golang:1.19

WORKDIR /go/src
COPY ./api .

RUN go get -u gorm.io/gorm && \
    go get -u gorm.io/driver/mysql

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
