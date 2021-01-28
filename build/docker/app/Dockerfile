ARG IMAGE=golang:1.15.7-alpine3.13

# #####################
# base
FROM ${IMAGE} as base

ARG APP_NAME

ENV WORK_DIR /${APP_NAME}/

RUN apk --update add --no-cache git openssh

WORKDIR ${WORK_DIR}
COPY ./ ./
RUN go mod download

# #####################
# development
FROM base as development

ARG APP_NAME

RUN apk --update add tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && apk del tzdata \
    && rm -rf /var/cache/apk/*

RUN apk update add --no-cache \
           git \
           openssh \
           bash \
           gcc \
           mercurial \
           musl-dev \
           curl \
        && go get -u -v github.com/pilu/fresh

WORKDIR /${APP_NAME}/cmd/sample-api/

CMD fresh