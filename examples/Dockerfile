FROM --platform=${BUILDPLATFORM} ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
    && apt install -y \
        curl \
        openssh-server \
        wget \
    && rm -rf /var/lib/apt/lists/*

