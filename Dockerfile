ARG SERVICE=conversation-backend

FROM golang:1.23.1-alpine AS base

FROM base AS build

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE

RUN go build -v -o /usr/local/bin/${SERVICE} ./cmd/service/main.go

# Run application
FROM base AS production
ARG SERVICE

ENV REDIS_ADDR="localhost:6379"
EXPOSE 8081

COPY --from=build /usr/local/bin/${SERVICE} /usr/local/bin/${SERVICE}
COPY --from=build /usr/src/app/README.md README.md

CMD ["conversation-backend"]
