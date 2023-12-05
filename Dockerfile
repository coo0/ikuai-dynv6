FROM alpine:latest
 
WORKDIR /build
COPY ikuai-dynv6 .

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

CMD ["./ikuai-dynv6", "-c", "/etc/ikuai-dynv6/config.yml"]