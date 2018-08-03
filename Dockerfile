FROM alpine:3.6

EXPOSE 80

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
 