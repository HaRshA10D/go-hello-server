FROM golang:1.14-alpine as build

RUN apk add --update make

WORKDIR /go/src/go-hello-server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make copy-dev-config
RUN make build

# --------

FROM alpine:latest

COPY --from=build /go/src/go-hello-server/build .
COPY --from=build /go/src/go-hello-server/application.yml .

RUN echo -e "#!/bin/bash\n/go-hello-server server" > run.sh

RUN chmod u+x run.sh

EXPOSE 9000

CMD ["sh", "run.sh"]
