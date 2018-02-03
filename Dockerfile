FROM golang:1.9.2-alpine3.6 AS build
MAINTAINER Brett Mandler <brettmandler@gmail.com>
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep
WORKDIR /go/src/arithmatic/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . /go/src/arithmatic/
RUN go build -o /bin/arithmatic -tags=jsoniter .

FROM scratch
COPY --from=build /bin/arithmatic /bin/arithmatic
ENTRYPOINT ["/bin/arithmatic"]

EXPOSE 8080
