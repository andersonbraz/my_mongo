FROM golang:alpine as builder
RUN mkdir /build 
ADD main.go /build/
WORKDIR /build
RUN apk add --no-cache curl
RUN apk add --no-cache git
#RUN go get -d -v github.com/gorilla/mux \
#	&& go get -d -v gopkg.in/mgo.v2/bson \
#	&& go get -d -v gopkg.in/mgo.v2
RUN go get -d -v go.mongodb.org/mongo-driver/mongo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cleacgo -ldflags '-extldflags "-static"' -o main .
#FROM scratch
FROM alpine:latest 
EXPOSE 8030 
WORKDIR /app/
COPY --from=builder /build/main .
CMD ["./main"]