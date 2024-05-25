FROM traefik:v2.5

COPY main.go /plugins-local/src/modsecurity-plugin/modsecurity-plugin.go
COPY modsecurity.conf /etc/modsecurity/modsecurity.conf

ENV GO111MODULE=on

CMD ["--experimental.plugins.modsecurity.modulename=github.com/yourusername/modsecurity-plugin", "--experimental.plugins.modsecurity.version=v0.0.1"]
