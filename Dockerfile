FROM golang:1.13.8-alpine3.11

RUN echo -e "http://nl.alpinelinux.org/alpine/v3.11/main\nhttp://nl.alpinelinux.org/alpine/v3.11/community" > /etc/apk/repositories

RUN apk update && \
	apk add git \
	build-base \
	tzdata \
	postgresql-dev

RUN echo "America/Sao_Paulo" > /etc/timezone

ADD . /go/src/github.com/fernandochimi/bank-account

WORKDIR /go/src/

RUN go get github.com/gorilla/mux && \
	go get github.com/jinzhu/gorm && \
	go get github.com/dgrijalva/jwt-go && \
	go get github.com/joho/godotenv && \
	go get github.com/lib/pq && \
	go get golang.org/x/crypto/bcrypt