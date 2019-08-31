FROM golang:1.12-alpine

ARG git_tag
ARG git_commit

RUN apk add --no-cache git build-base curl

WORKDIR /go/src/github.com/justpark/auth/keto

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on

RUN go mod download

ADD keto/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -X github.com/justpark/auth/keto/cmd.Version=$git_tag -X github.com/justpark/auth/keto/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/justpark/auth/keto/cmd.GitHash=$git_commit" -a -installsuffix cgo -o keto

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/github.com/justpark/auth/keto/keto /usr/bin/keto

ENTRYPOINT ["keto"]

CMD ["serve"]
