FROM golang:1.15

RUN apt-get update
RUN apt-get install -y pkg-config && \
    apt-get install -y libaio1 && \
    apt-get install unzip

ENV CLIENT_FILENAME instantclient_12_1.zip
COPY /oracle/${CLIENT_FILENAME} .
COPY /oracle/oci8.pc /usr/lib/pkgconfig/oci8.pc

ENV LD_LIBRARY_PATH /usr/lib:/usr/local/lib:/usr/instantclient_12_1

# to build the application with mattn/go-oci8, it is necessary to extract all files, including the SDK.
RUN unzip ${CLIENT_FILENAME} -d /usr &&  \
    ln -s /usr/instantclient_12_1/libclntsh.so.12.1 /usr/instantclient_12_1/libclntsh.so && \
    ln -s /usr/instantclient_12_1/libclntshcore.so.12.1 /usr/instantclient_12_1/libclntshcore.so && \
    ln -s /usr/instantclient_12_1/libocci.so.12.1 /usr/instantclient_12_1/libocci.so
    
RUN go get -u github.com/mattn/go-oci8
ENV GO111MODULE=on\
    CGO_ENABLED=1\
    GOOS=linux\
    GOARCH=amd64
WORKDIR /go/src/app
COPY go.mod .
RUN go mod download
COPY . .
RUN go install
