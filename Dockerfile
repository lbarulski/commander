FROM golang:1.10-alpine3.7 as builder
RUN apk --no-cache add git \
    && go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/commander
COPY . .
RUN dep ensure -vendor-only
RUN mkdir dist \
    && CGO_ENABLED=0 go build -o dist/initiator initiator/initiator.go \
    && CGO_ENABLED=0 go build -o dist/shell shell/shell.go \
    && CGO_ENABLED=0 go build -o dist/executor executor/executor.go

FROM alpine:3.7
RUN apk --no-cache add ca-certificates openssh \
    && echo "PasswordAuthentication no" >> /etc/ssh/sshd_config \
    && echo "/usr/local/bin/shell" > /etc/shells \
    && chmod 700 /bin/sh /bin/ash /bin/busybox \
    && ssh-keygen -A
WORKDIR /root
COPY entrypoint.sh /root/entrypoint.sh
COPY --from=builder /go/src/commander/dist/ /usr/local/bin/
EXPOSE 22
ENTRYPOINT ["/root/entrypoint.sh"]