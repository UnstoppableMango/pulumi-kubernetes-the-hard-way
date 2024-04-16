FROM --platform=${BUILDPLATFORM} ubuntu:24.04

RUN echo 'root:root' | chpasswd
RUN printf '#!/bin/sh\nexit 0' > /usr/sbin/policy-rc.d \
    && printf "systemctl start systemd-logind" >> /etc/profile

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
    && apt install -y \
        curl \
        dbus \
        dbus-user-session \
        net-tools \
        openssh-server \
        systemd \
        systemd-sysv \
        sudo \
        wget \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 6969 test \
    && echo 'test:test' | chpasswd
# RUN echo 'test ALL=(ALL) NOPASSWD: ALL' | tee -a /etc/sudoers
# Test changes

RUN systemctl enable ssh

EXPOSE 22

ENTRYPOINT ["/sbin/init"]
