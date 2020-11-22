FROM golang:1.14

WORKDIR ${GOPATH}/src/slowly

COPY . .

RUN make build
ENTRYPOINT ["./slowly"]
