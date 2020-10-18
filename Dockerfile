FROM depscloud/devbase:latest AS builder

WORKDIR /go/src/beacon

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .

ARG VERSION="local"
RUN go install -ldflags="-X 'main.version=${VERSION}'" ./cmd/beacon/

FROM depscloud/base:latest

COPY --from=builder /go/bin/beacon /usr/bin/beacon

WORKDIR /home/depscloud
USER 13490:13490

ENTRYPOINT [ "/usr/bin/beacon" ]
