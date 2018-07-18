# docker build -t grpc-client ./client --build-arg app_env=dev
# docker run -it -p 8080:8080 -v /home/hrabo/go/src/github.com/hannesrabo/go-micro-gRPC/client/app:/go/src/go-micro-gRPC/client grpc-client

FROM golang:1.8

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/hannesrabo/go-micro-gRPC/client
WORKDIR /go/src/github.com/hannesrabo/go-micro-gRPC/client

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
    then \
    app; \
    else \
    go get github.com/pilu/fresh && \
    fresh; \
    fi

EXPOSE 8080