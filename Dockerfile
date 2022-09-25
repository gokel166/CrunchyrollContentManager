FROM ubuntu:22.04

LABEL org.opencontainers.image.source="https://github.com/gokel166/CrunchyrollContentManager"

WORKDIR /app

HEALTHCHECK --interval=30s --timeout=5s \
    CMD wget -q -O - http://$HOSTNAME:14235/health-check || exit 1

COPY . .

ENTRYPOINT [ "bin/mono"]

CMD [ "serve" ]