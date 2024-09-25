FROM ubuntu:22.04 AS build
WORKDIR /app

RUN apt-get update && \
  apt-get install -y git build-essential \
  && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

RUN git clone https://github.com/ggerganov/whisper.cpp.git /app
RUN make

FROM ubuntu:22.04 AS runtime
WORKDIR /app

RUN apt-get update && \
  apt-get install -y curl ffmpeg libsdl2-dev \
  && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

COPY --from=build /app /app

RUN /app/models/download-ggml-model.sh large-v3-q5_0 /app/models

EXPOSE 8081

CMD [ "/app/server", \
      "--host", "0", \
      "--port", "8081", \
      "--convert", \
      "--best-of", "5", \
      "--beam-size", "5", \
      "-m", "/app/models/ggml-large-v3-q5_0.bin" ]
