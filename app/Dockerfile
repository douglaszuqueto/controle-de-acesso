FROM golang:alpine as builder

RUN mkdir -p /cda
COPY  ./ /cda
WORKDIR /cda
RUN GOPATH=$(pwd) go build -ldflags "-s -w" -o app

FROM alpine
WORKDIR /cda/
COPY --from=builder /cda/app .
EXPOSE 3000 8080
CMD ["./app"]
