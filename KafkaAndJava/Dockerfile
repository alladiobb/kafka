FROM openjdk:11-slim  

WORKDIR /java/src
ENV PATH="/java/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]