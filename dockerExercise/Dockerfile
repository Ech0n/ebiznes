FROM ubuntu:24.04

RUN apt-get update

#python
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt-get install -y python3.10
RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.10 1

#sdkman
RUN apt-get install -y curl unzip zip
RUN curl -s https://get.sdkman.io | bash

#java 8, kotlin, gradle za pomocą sdkman
SHELL ["/bin/bash", "-c"]    
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install java 8.0.442-zulu
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install kotlin
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install gradle

WORKDIR /app

COPY /app /app

#uruchamianie
RUN source "/root/.sdkman/bin/sdkman-init.sh" && gradle build
ENV JAVA_HOME="/root/.sdkman/candidates/java/current"
CMD ["/root/.sdkman/candidates/gradle/current/bin/gradle", "run"]