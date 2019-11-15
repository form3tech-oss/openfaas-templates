ARG BUILDER_BASE_IMAGE="golang:1.13.4"

FROM openfaas/of-watchdog:0.7.2 as watchdog

FROM ${BUILDER_BASE_IMAGE} as builder
ARG GITHUB_USERNAME=""
ARG GITHUB_TOKEN=""
ARG GOPRIVATE=""
ARG GOPROXY=""
ENV GOCACHE=/root/.cache
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler/function
RUN [ "${GITHUB_USERNAME}" != "" ] && echo "machine github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" >> /root/.netrc || :
COPY function/go.mod function/go.sum ./
RUN go mod download
WORKDIR /go/src/handler
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" --installsuffix cgo -o /handler -v .

FROM ${BUILDER_BASE_IMAGE} as gocache
COPY --from=builder /root/.cache /root/.cache

FROM gcr.io/distroless/base:debug
ENV fprocess="/dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec /handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"
COPY --from=watchdog /fwatchdog /fwatchdog
COPY --from=builder /go/bin/dlv /dlv
COPY --from=builder /handler /handler
ENTRYPOINT [""]
CMD ["/fwatchdog"]
