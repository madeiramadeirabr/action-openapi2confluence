FROM alpine:3.16

ADD https://github.com/madeiramadeirabr/action-openapi2confluence/releases/download/v0.6/openapi2confluence_linux_amd64  /usr/bin/openapi2confluence

RUN chmod +x /usr/bin/lifecycledoc

COPY entrypoint.sh entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
