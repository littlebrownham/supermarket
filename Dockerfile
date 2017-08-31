FROM golang:1.9

WORKDIR /opt/go/src/github.com/littlebrownham/supermarket
COPY . /opt/go/src/github.com/littlebrownham/supermarket

EXPOSE 50200

RUN bin/build

CMD ["build/supermarket"]
CMD ["bin/test"]

ENV \ 
    SUPERMARKET_PORT=50200 \ 
    SUPERMARKET_HOST=0.0.0.0



