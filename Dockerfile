FROM  golang:1.12.4-alpine as  builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/shrikar007/customer-rest-api
RUN apk add --no-cache protobuf git make bash build-base \
	&& rm -rf /var/cache/apk/*
ADD . ./
RUN  go build -a -o /main .

FROM alpine
WORKDIR /
RUN apk add --no-cache ca-certificates \
	&& update-ca-certificates \
    # cleanup
    && rm -rf /var/cache/apk/*

COPY --from=builder /main .
RUN chmod +x main
EXPOSE 8086

ENTRYPOINT ["/main"]