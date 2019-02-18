ARG GOLANG_VERSION
ARG IMAGE_BASE

FROM golang:${GOLANG_VERSION}-alpine AS golang

FROM ${IMAGE_BASE}

# Libraries used in the build process of this image.
RUN apk add --update --no-cache \
        build-base \
        bash \
        git

# change default shell from ash to bash
RUN sed -i -e "s/bin\/ash/bin\/bash/" /etc/passwd

# Installs Golang.
COPY --from=golang /usr/local/go /usr/local/go

RUN export PATH="/usr/local/go/bin:$PATH" && \
    go version

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Define our workding outside of $GOPATH (we're using go modules).
WORKDIR /src

# Copy our module dependencies definitions.
COPY go.mod .
COPY go.sum .

# Install module dependencies.
RUN go mod download
