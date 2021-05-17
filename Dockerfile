FROM golang:1.16.4-alpine3.13

COPY . src/

RUN go get -u github.com/gorilla/mux

RUN cd src/ && go build -o ../bin

EXPOSE 8090

ENTRYPOINT [ "./bin/go-juice-inventory" ]
