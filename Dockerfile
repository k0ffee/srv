FROM golang:1.12.1 as builder
WORKDIR /var/tmp
USER daemon
COPY Makefile .
COPY srv.go .
ENV GOCACHE=/var/tmp/.gocache
RUN make
RUN make test

FROM scratch
WORKDIR /
COPY --from=builder /var/tmp/srv .
USER 1
ENTRYPOINT ["/srv", "8080"]
