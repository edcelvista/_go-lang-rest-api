FROM alpine:3.22.1
LABEL maintainer "administrator@edcelvista.com"
RUN wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz && export PATH=$PATH:/usr/local/go/bin

WORKDIR /opt

COPY bin/app /usr/local/bin/

RUN chmod ugo+x /usr/local/bin/app

EXPOSE 8443

# command: ["sleep", "infinity"]
CMD app
