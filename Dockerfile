FROM ubuntu:latest
WORKDIR /work/
COPY obs-scene-server .
COPY static .
COPY config.json .
RUN chmod +x obs-scene-server
ENTRYPOINT ["/work/obs-scene-server"]
