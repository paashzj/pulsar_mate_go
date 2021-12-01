FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o pulsar_mate .
WORKDIR /opt/sh/compile/cmd/config
RUN go build -o config_gen .


FROM ttbb/pulsar:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/pulsar/mate

COPY --from=build /opt/sh/compile/pkg/pulsar_mate /opt/sh/pulsar/mate/pulsar_mate
COPY --from=build /opt/sh/compile/cmd/config/config_gen /opt/sh/pulsar/mate/config_gen

COPY config/client_original.conf /opt/sh/pulsar/conf/client_original.conf
COPY config/broker_original.conf /opt/sh/pulsar/conf/broker_original.conf
COPY config/standalone_original.conf /opt/sh/pulsar/conf/standalone_original.conf

WORKDIR /opt/sh/pulsar

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/pulsar/mate/scripts/start.sh"]