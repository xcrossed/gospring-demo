# Version: 0.1
FROM golang:1.14 as builder

ARG CI_COMMIT_TAG

ENV CI_COMMIT_TAG=$CI_COMMIT_TAG

RUN mkdir -p /data/build
ADD ./ /data/build
RUN cd /data/build && make clean release


FROM alpine:3.5
MAINTAINER XXX "xxx@cmcm.com"

ARG env_name
ARG gin_mode

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk add tzdata ca-certificates && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    rm -rf /var/cache/apk/* && mkdir -pv /data/apps/minigame/logs
ADD conf /data/apps/gospring-demo/conf/
COPY --from=builder /data/build/bin /data/apps/gospring-demo/bin/
WORKDIR /data/apps/gospring-demo

EXPOSE 8000

ENV SERVICE_NAME=gospring-demo
ENV GIN_MODE=${gin_mode}
ENV PLUGIN_CONF_PATH=/data/apps/gospring-demo/conf/${env_name}
ENV ZONEINFO /opt/zoneinfo.zip

HEALTHCHECK --timeout=4s CMD wget -q -O /dev/null localhost:8000/version || exit 1

CMD /data/apps/gospring-demo/bin/gospring-demo -config /data/apps/gospring-demo/conf/$env_name/config.json
