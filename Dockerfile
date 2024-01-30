FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY . .

RUN go build -o /go/go-monitoring ./cmd/app

FROM scratch

COPY --from=build /go/go-monitoring /go/go-monitoring

CMD ["/go/go-monitoring"]