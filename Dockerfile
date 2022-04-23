FROM ubuntu:latest
WORKDIR /work/
COPY xobs-scene-server .
COPY static .
COPY config.json .
RUN chmod +x xobs-scene-server
ENTRYPOINT ["/work/xobs-scene-server"]
