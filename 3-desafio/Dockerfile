FROM golang:1.17-alpine3.15 AS builder
ARG SRC_PATH=.
ENV GO111MODULE=on
COPY /starters /starters
COPY ${SRC_PATH}/ /build/container/
WORKDIR /build/container
RUN ls -lha && go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app

###
FROM alpine:3.15 AS dist
RUN mkdir /app && \
    addgroup -S user && adduser -S user -G user \
    && chown -R user:user /app \
    && touch /app/.env

USER user
COPY --from=builder --chown=user /build/container/app /app/

WORKDIR /app
EXPOSE 8080
CMD ["./app"]
