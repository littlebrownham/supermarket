FROM scratch

WORKDIR /go/src/github.com/littlebrownham/supermarket
COPY build/supermarket /go/src/github.com/littlebrownham/supermarket

EXPOSE 50200

CMD ["./supermarket"]
