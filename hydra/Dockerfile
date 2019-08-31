# To compile this image manually run:
#
# $ GO111MODULE=on GOOS=linux GOARCH=amd64 go build && docker build -t oryd/hydra:v1.0.0-rc.7_oryOS.10 . && rm hydra
FROM golang:1.12-alpine

ARG git_tag
ARG git_commit

RUN apk add --no-cache git build-base curl

WORKDIR /go/src/github.com/justpark/auth/hydra

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on

RUN go mod download

ADD hydra/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -X github.com/justpark/auth/hydra/cmd.Version=$git_tag -X github.com/justpark/auth/hydra/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/justpark/auth/hydra/cmd.GitHash=$git_commit" -a -installsuffix cgo -o hydra

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 hydra /usr/bin/hydra

USER 1000

ENTRYPOINT ["hydra"]
CMD ["serve", "all"]
