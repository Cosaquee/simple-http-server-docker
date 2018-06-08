FROM debian:stretch
LABEL "author" "your name <your_email@email.com>"

RUN apt update && apt install -y git
RUN echo "8.8.8.8" > /etc/resolve.conf
