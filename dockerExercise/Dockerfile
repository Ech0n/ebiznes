FROM ubuntu:24.04

RUN apt-get update

#python
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt-get install -y python3.10
RUN apt-get clean
RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.10 1

# FROM ubuntu:24.04 as java_kotlin_ubuntu

#java
# RUN apt-get update
RUN apt-get install -y openjdk-8-jdk
RUN apt-get clean
    
#kotlin
RUN apt-get install -y curl unzip zip
RUN curl -s https://get.sdkman.io | bash
SHELL ["/bin/bash", "-c"]    
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install kotlin

#gradle
RUN source "/root/.sdkman/bin/sdkman-init.sh" && sdk install gradle 8.13

#jdbc
RUN curl -O https://repo1.maven.org/maven2/org/xerial/sqlite-jdbc/3.45.3.0/sqlite-jdbc-3.45.3.0.jar

WORKDIR /WORKDIR

COPY . .

#uruchamianie
RUN source "/root/.sdkman/bin/sdkman-init.sh" && gradle build
CMD ["/root/.sdkman/candidates/gradle/current/bin/gradle", "run"]