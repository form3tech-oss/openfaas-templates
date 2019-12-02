# syntax=docker/dockerfile:1.1.3-experimental

FROM openfaas/of-watchdog:0.7.2 as watchdog

FROM golang:1.13.4 as builder
ARG GITHUB_USERNAME=""
ARG GITHUB_TOKEN=""
ARG GO111MODULE="off"
ARG GOFLAGS=""
ARG GOPRIVATE=""
ARG GOPROXY=""
ARG MODULE_PATH="handler"
RUN [ "${GITHUB_USERNAME}" != "" ] && echo "machine github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" >> /root/.netrc || :
RUN mkdir -p /go/src/$MODULE_PATH
WORKDIR /go/src/$MODULE_PATH
COPY . .
RUN mv ./function ./function.tmp && mv ./function.tmp/function . && rm -rf ./function.tmp
RUN sed -i "s|handler/function|$MODULE_PATH/function|g" go.mod handler.go
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build --ldflags='-d -s -w' -o /handler -tags netgo -v .

FROM gcr.io/distroless/static:nonroot
ENV fprocess="/handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"
COPY --from=watchdog /fwatchdog /fwatchdog
COPY --from=builder /handler /handler
CMD ["/fwatchdog"]
