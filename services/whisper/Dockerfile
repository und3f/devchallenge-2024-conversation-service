FROM ubuntu:22.04 AS build
WORKDIR /app

RUN apt-get update && \
  apt-get install -y wget tar build-essential \
  && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

RUN wget https://github.com/ggerganov/whisper.cpp/archive/69339af2d104802f3f201fd419163defba52890e.tar.gz -O /tmp/whisper.cpp.tar.gz && \
    tar -xvf /tmp/whisper.cpp.tar.gz --strip-components 1 && \
    make

FROM ubuntu:22.04 AS runtime
WORKDIR /app

RUN apt-get update && \
  apt-get install -y curl ffmpeg libsdl2-dev \
  && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

COPY --from=build /app /app

RUN /app/models/download-ggml-model.sh large-v2-q5_0 /app/models

EXPOSE 8081

CMD [ "/app/server", \
      "--host", "0", \
      "--port", "8081", \
      "--threads", "8", \
      "--convert", \
      "--best-of", "5", \
      "--beam-size", "5", \
      "-m", "/app/models/ggml-large-v2-q5_0.bin" ]
