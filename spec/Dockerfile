FROM alpine
WORKDIR /usr/local/bin
ARG _path
ARG APP
ARG CSV
ENV APP=${APP}
COPY ${_path}/${APP} .
COPY ${CSV} /csv
RUN chmod +x ${APP}
RUN apk  add --no-cache --update \
    tzdata \
    ca-certificates

RUN echo -e '#!/bin/sh \n ${APP} $@' > entrypoint.sh
RUN chmod +x entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
