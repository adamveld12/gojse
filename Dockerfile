FROM alpine

ENV GOJSE_USERNAME ""
ENV GOJSE_PASSWORD ""

WORKDIR /app

RUN apk add --update ca-certificates \
  && rm -rf /var/cache/apk/*

COPY ./gojse /app/

CMD '/app/gojse'
