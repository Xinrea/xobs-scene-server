FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
WORKDIR /work/
COPY xobs-scene-server .
COPY static static
COPY config.json .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
RUN chmod +x xobs-scene-server
ENTRYPOINT ["./entrypoint.sh"]
CMD ["localhost:8080"]
