FROM golang:1.16.12 AS build-env

WORKDIR /home-server

ADD . /home-server

RUN sh /home-server/build.sh

# Final stage
FROM debian:latest

RUN apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends ca-certificates curl

EXPOSE 8000

WORKDIR /home-server
COPY --from=build-env /home-server/output/ /home-server

CMD ["/home-server/home-server"]
