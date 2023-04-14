FROM golang:1.19-alpine as builder

ARG SERVICE_VERSION
ENV SERVICE_VERSION=$SERVICE_VERSION

RUN apk add git

WORKDIR /app

COPY . ./

RUN go build -o /app/fluxcorgi cmd/app/main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/fluxcorgi /app/fluxcorgi
COPY --from=builder /app/third_party /app/third_party
COPY --from=builder /app/swagger/docs.swagger.json /app/swagger/docs.swagger.json

EXPOSE 8080

ENTRYPOINT [ "/app/fluxcorgi" ]
