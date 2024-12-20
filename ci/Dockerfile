FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive
ENV PATH=$PATH:/usr/local/go/bin

# Install dependencies
RUN apt-get update && \
    apt-get install -y build-essential make git wget tar gcc-12 kmod iputils-ping net-tools curl jq && \
    apt-get install -y linux-headers-generic linux-modules-$(uname -r) linux-modules-extra-$(uname -r) && \
    rm -rf /usr/local/go && \
    wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz -O - | tar -C /usr/local -zxvf - && \
    echo "export PATH=$PATH:/usr/local/go/bin" >> /root/.profile

# Install packet rusher
RUN git clone https://github.com/HewlettPackard/PacketRusher /root/PacketRusher && \
    cd /root/PacketRusher/lib/gtp5g && \
    apt-get install -y linux-headers-$(uname -r) && \
    export CC=gcc-12 && \
    make clean && make && \
    mkdir -p /lib/modules/$(uname -r)/kernel/drivers/net && \
    mkdir -p /etc/modules-load.d && \
    depmod -a && \
    make install && \
    cd /root/PacketRusher && \
    . /root/.profile && \
    go mod download && go build cmd/packetrusher.go

WORKDIR /root

VOLUME [ "/root" ]
