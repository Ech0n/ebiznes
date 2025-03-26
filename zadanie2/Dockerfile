FROM ubuntu:24.04

RUN apt-get update

#sdkman
RUN apt-get install -y curl unzip zip
RUN curl -s https://get.sdkman.io | bash

RUN curl -O https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip && \
    unzip ngrok-stable-linux-amd64.zip -d /usr/local/bin && \
    rm ngrok-stable-linux-amd64.zip

#java 8, kotlin, gradle za pomocą sdkman
SHELL ["/bin/bash", "-c"]
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install java
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install sbt

VOLUME /build
WORKDIR /app

COPY /play-scala /app

RUN source "/root/.sdkman/bin/sdkman-init.sh" && sbt dist

RUN unzip /app/target/universal/my-play-scala-app-1.0.zip -d /build/
    # chmod +x /build/bin/*

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 9000 4040

# CMD ["/entrypoint.sh"]