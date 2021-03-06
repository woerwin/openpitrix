# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM openpitrix/openpitrix-builder as builder

WORKDIR /go/src/openpitrix.io/openpitrix/
COPY . .

RUN mkdir -p /openpitrix_bin
RUN go generate openpitrix.io/openpitrix/pkg/version && \
	CGO_ENABLED=0 GOOS=linux GOBIN=/openpitrix_bin go install -ldflags '-w -s' -tags netgo openpitrix.io/openpitrix/cmd/... && \
	CGO_ENABLED=0 GOOS=linux GOBIN=/openpitrix_bin go install -ldflags '-w -s' -tags netgo openpitrix.io/openpitrix/metadata/cmd/...

RUN find /openpitrix_bin -type f -exec upx {} \;

FROM alpine:3.7
RUN apk add --update ca-certificates && update-ca-certificates
COPY --from=builder /openpitrix_bin/* /usr/local/bin/

CMD ["sh"]