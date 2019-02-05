FROM alpine:edge

RUN apk add --update --no-cache \
        ttf-dejavu \
        ttf-droid \
        ttf-freefont \
        ttf-liberation \
        ttf-ubuntu-font-family \
        wkhtmltopdf

RUN mkdir -p /app/public
COPY public/ /app/public
COPY server /app/server

WORKDIR /app

EXPOSE 8515/tcp

CMD ["/app/server"]

ARG BUILD_DATE
ARG VCS_REF

LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.vcs-url="https://github.com/opsway/documents.git" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.schema-version="1.0" \
      org.label-schema.vendor="OpsWay <start@opsway.com>" \
      org.label-schema.name="documents" \
      org.label-schema.description="Documents generation API" \
      org.label-schema.url="https://github.com/opsway/documents"
