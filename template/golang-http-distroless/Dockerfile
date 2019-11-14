ARG BUILDER_BASE_IMAGE="golang:1.13.4"
ARG GOPROXY=""

FROM openfaas/of-watchdog:0.7.2 as watchdog

FROM ${BUILDER_BASE_IMAGE} as builder
ENV CGO_ENABLED=0
ENV GOCACHE=/root/.cache
RUN mkdir -p /go/src/handler
WORKDIR /go/src/handler/function
COPY function/go.mod function/go.sum ./
RUN go mod download
WORKDIR /go/src/handler
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build --installsuffix cgo --ldflags='-s -w' -o /handler -v .

FROM ${BUILDER_BASE_IMAGE} as gocache
COPY --from=builder /root/.cache /root/.cache

FROM gcr.io/distroless/static:nonroot
ENV fprocess="/handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"
COPY --from=watchdog /fwatchdog /fwatchdog
COPY --from=builder /handler /handler
CMD ["/fwatchdog"]
