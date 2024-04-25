FROM --platform=${BUILDPLATFORM} ubuntu:noble-20240225

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
    && apt install -y \
        curl \
        net-tools \
        openssh-server \
        sudo \
        wget \
    && rm -rf /var/lib/apt/lists/*

RUN echo 'root:root' | chpasswd \
    && sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config \
    && printf '#!/bin/sh\nexit 0' > /usr/sbin/policy-rc.d \
    && useradd -rm -d /home/test -s /bin/bash -g root -G sudo -u 6969 test \
    && echo 'test:test' | chpasswd

RUN service ssh start

# Fix for CI until we get tests refactored
RUN mkdir /config && chown 'test:root' /config

EXPOSE 22

ENTRYPOINT ["/sbin/init"]
