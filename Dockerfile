FROM ubuntu:latest

RUN mkdir app

RUN apt-get update && apt-get install -y ca-certificates

COPY ./bin/opengovernance-linux app/

