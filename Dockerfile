FROM golang:alpine3.15 AS builder
ARG main_path

WORKDIR /app

COPY . .

RUN go build -o "app" -mod=vendor ${main_path:-./}

FROM alpine:3.15.5

RUN apk update \
	&& apk -U upgrade \
	&& apk add --no-cache ca-certificates bash \
	&& update-ca-certificates --fresh \
	&& rm -rf /var/cache/apk/*

# adds app user and fix app folder's permission
RUN	addgroup klever \
	&& adduser -S klever -u 1000 -G klever

USER klever

COPY --from=builder --chown=klever:klever app /usr/local/bin/
RUN chmod +x /usr/local/bin/app

ENTRYPOINT [ "/usr/local/bin/app" ]