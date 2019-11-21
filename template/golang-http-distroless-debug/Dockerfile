# syntax=docker/dockerfile:1.1.3-experimental

FROM openfaas/of-watchdog:0.7.2 as watchdog

FROM golang:1.13.4 as builder
ARG GITHUB_USERNAME=""
ARG GITHUB_TOKEN=""
ARG GO111MODULE="off"
ARG GOFLAGS=""
ARG GOPRIVATE=""
ARG GOPROXY=""
RUN --mount=type=cache,target=/root/.cache/go-build go get -u github.com/go-delve/delve/cmd/dlv
RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler
RUN [ "${GITHUB_USERNAME}" != "" ] && echo "machine github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" >> /root/.netrc || :
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /handler -tags netgo -v .

FROM gcr.io/distroless/base:debug
ENV fprocess="/dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"
COPY --from=watchdog /fwatchdog /fwatchdog
COPY --from=builder /go/bin/dlv /dlv
COPY --from=builder /handler /handler
ENTRYPOINT [""]
CMD ["/fwatchdog"]
