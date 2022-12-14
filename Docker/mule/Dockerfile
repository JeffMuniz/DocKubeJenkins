FROM anapsix/alpine-java:8_jdk_nashorn

LABEL maintainer="https://manik.magar.me"

# Define environment variables.
ENV BUILD_DATE=12082019
ENV MULE_HOME=/opt/mule-standalone
ENV MULE_VERSION=3.9.0
ENV MULE_MD5=39b773bf20702f614faf30b2ffca4716
ENV TINI_SUBREAPER=
ENV TZ=America/Sao_Paulo

# SSL Cert for downloading mule zip
RUN apk --no-cache update && \
    apk --no-cache upgrade && \
    apk --no-cache add ca-certificates && \
    update-ca-certificates && \
    apk --no-cache add openssl && \
    apk add --update tzdata && \
    rm -rf /var/cache/apk/*

RUN adduser -D -g "" mule mule

RUN mkdir /opt/mule-standalone-${MULE_VERSION} && \
    ln -s /opt/mule-standalone-${MULE_VERSION} ${MULE_HOME} && \
    chown mule:mule -R /opt/mule*

RUN echo ${TZ} > /etc/timezone

USER mule

# For checksum, alpine linux needs two spaces between checksum and file name
RUN cd ~ && wget https://repository-master.mulesoft.org/nexus/content/repositories/releases/org/mule/distributions/mule-standalone/${MULE_VERSION}/mule-standalone-${MULE_VERSION}.tar.gz && \
    echo "${MULE_MD5}  mule-standalone-${MULE_VERSION}.tar.gz" | md5sum -c && \
    cd /opt && \
    tar xvzf ~/mule-standalone-${MULE_VERSION}.tar.gz && \
    rm ~/mule-standalone-${MULE_VERSION}.tar.gz

# Define mount points.
VOLUME ["${MULE_HOME}/logs", "${MULE_HOME}/conf", "${MULE_HOME}/apps", "${MULE_HOME}/domains"]

# Define working directory.
WORKDIR ${MULE_HOME}

ENV ENVIRONMENT=${ENVIRONMENT}

CMD /opt/mule-standalone/bin/mule console -M-Dmule.env=${ENVIRONMENT}

# Default http port
EXPOSE 8081