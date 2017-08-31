FROM golang:1.9

WORKDIR /go/src/github.com/littlebrownham/supermarket
COPY . /go/src/github.com/littlebrownham/supermarket

EXPOSE 50200

RUN bin/build

CMD ["build/supermarket"]

ENV \ 
    SUPERMARKET_PORT=50200 \ 
    SUPERMARKET_HOST=0.0.0.0



