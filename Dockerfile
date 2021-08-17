FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o pulsar_mate .


FROM ttbb/pulsar:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/pulsar/mate

COPY --from=build /opt/sh/compile/pkg/pulsar_mate /opt/sh/pulsar/mate/pulsar_mate

COPY config/broker_original.conf /opt/sh/pulsar/conf/broker_original.conf

WORKDIR /opt/sh/pulsar

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/pulsar/mate/scripts/start.sh"]