FROM golang:alpine AS buildIma
ARG project="local-life"
WORKDIR /app/project
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go mod tidy && go build -o bin/${project} main.go

FROM alpine

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
LABEL maintainer=chenyi date="2023-01-01"
WORKDIR /app
COPY --from=buildIma /app/project/bin/ /app/bin/
COPY --from=buildIma /app/project/conf/ /app/conf/

EXPOSE 8080
ENTRYPOINT  ["./bin/local-life"]
# docker build --network host -t locallife:v1 -f Dockerfile .
# docker run -it locallife:v1
