FROM golang

RUN apt-get update
# install git
RUN apt-get -y install git

# Expose service ports.
EXPOSE 8000