FROM golang:1.19-alpine

WORKDIR /go/src
COPY ./api .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
