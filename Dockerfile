FROM ubuntu:latest

RUN mkdir app

RUN apt-get update && apt-get install -y ca-certificates

COPY ./dist/linux_linux_arm64/checkctl app/

