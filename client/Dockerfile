# docker build -t grpc-server ./server --build-arg app_env=dev

FROM golang:1.8

ARG app_env
ENV APP_ENV $app_env

# Installing ubuntu packages
RUN apt-get update && apt-get install -y \
    unzip 

# Including the resources needed
WORKDIR /go/src/github.com/hannesrabo/go-micro-gRPC

# Moving share resources and scripts
COPY ./shared /go/src/github.com/hannesrabo/go-micro-gRPC/shared
COPY ./build /go/src/github.com/hannesrabo/go-micro-gRPC/build

# Install protobuf and protobuf for go
RUN ./build/scripts/install-protobuff-ubuntu.sh
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN export PATH=$PATH:$GOPATH/bin

# Compiling the protobuffer for gRPC
RUN ./build/scripts/build-protobuff.sh

#########################################################################

COPY ./client/app /go/src/github.com/hannesrabo/go-micro-gRPC/client
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