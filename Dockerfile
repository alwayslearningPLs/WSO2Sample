ARG WSO2_VERSION=4.1.0
FROM ubuntu:22.04
ARG WSO2_VERSION

RUN apt-get update > /dev/null 2>&1 && apt-get install --yes openjdk-11-jdk wget unzip > /dev/null 2>&1 && rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME="/usr/lib/jvm/java-11-openjdk-amd64" \
  WSO2_HOME="/opt/wso2am-${WSO2_VERSION}"

ENV PATH="${JAVA_HOME}/bin:${WSO2_HOME}/bin:$PATH"

RUN echo "export PATH=${PATH}" >> /root/.bashrc && \
  wget --quiet https://github.com/wso2/product-apim/releases/download/v${WSO2_VERSION}/wso2am-${WSO2_VERSION}.zip -O /tmp/wso2am-${WSO2_VERSION}.zip && \
  unzip -q /tmp/wso2am-${WSO2_VERSION}.zip -d /opt

EXPOSE 9443

ENTRYPOINT ["/bin/bash", "-c", "api-manager.sh"]

# CMD ["here", "your", "arguments"]
