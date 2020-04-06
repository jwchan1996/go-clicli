FROM golang:1.13-alpine as builder

COPY ./ /api/
WORKDIR /api/
RUN go mod vendor && \
    CGO_ENABLED=0 go build -v -o publisher

EXPOSE 8084
ENTRYPOINT []

CMD [ "./api" ]