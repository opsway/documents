# Buils binary.
ARG GOLANG_VERSION
ARG IMAGE_BASE

FROM golang:${GOLANG_VERSION}-alpine AS golang

# Libraries used in the build process of this image.
RUN apk add --update --no-cache \
        git
        
ENV GOOS=linux\
    GOARCH=amd64\
    CGO_ENABLED=0

# Define our workding outside of $GOPATH (we're using go modules).
WORKDIR /src

# Copy our source code.
COPY . .

# Build our binary.
RUN go build  \
       -o entrypoint \
       main.go

# Last instructions of this build.
FROM ${IMAGE_BASE}

ARG VERSION
ARG BUILD_DATE
ARG VCS_REF

COPY public/ /app/public
COPY --from=golang /src/entrypoint /app/entrypoint

WORKDIR /app

ENV \
    DOCUMENTS_PORT=8515

EXPOSE ${DOCUMENTS_PORT}

CMD ["/app/entrypoint"]

# Build-time metadata as defined at http://label-schema.org
LABEL org.label-schema.build-date=${BUILD_DATE} \
      org.label-schema.name="documents" \
      org.label-schema.description="Service stateless API for converting HTML documents to PDF" \
      org.label-schema.url="https://opsway.github.io/documents/" \
      org.label-schema.vcs-ref=${VCS_REF} \
      org.label-schema.vcs-url="https://github.com/opsway/documents" \
      org.label-schema.vendor="OpsWay <start@opsway.com>" \
      org.label-schema.version=${VERSION} \
      org.label-schema.schema-version="1.0"
