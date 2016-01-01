FROM golang:onbuild
RUN go get gopkg.in/redis.v3
EXPOSE 8080