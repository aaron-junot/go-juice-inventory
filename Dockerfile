FROM golang:1.16.4-alpine3.13

RUN adduser --system --disabled-password app

COPY . src/

RUN go get -u github.com/gorilla/mux
RUN go get github.com/lib/pq

RUN cd src/ && go build -o ../bin

EXPOSE 8090

USER app

ENTRYPOINT [ "./bin/go-juice-inventory" ]
