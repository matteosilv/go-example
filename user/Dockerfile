FROM golang as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM scratch

COPY --from=build /go/src/app/main /

CMD [ "/main" ]