FROM golang:1.15

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY pkg/ pkg/

# Copy the static source
COPY static/ static/

# Build
RUN go build -v ./pkg/tnijto.go

ENTRYPOINT ["./tnijto"]
