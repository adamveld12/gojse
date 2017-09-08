from alpine

ENV JSEMINE_USERNAME ""
ENV JSEMINE_PASSWORD ""

WORKDIR /app

RUN apk add --update ca-certificates \
  && rm -rf /var/cache/apk/*

COPY ./jseminer /app/

CMD '/app/jseminer'
