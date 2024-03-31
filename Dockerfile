FROM golang:1.21 AS gobuild

WORKDIR /build

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

COPY ./ ./

RUN chmod +x migrate.sh migrate_test.sh

RUN go mod download && \
    CGO_ENABLED=0 go build server/cmd/app/main.go && \
    CGO_ENABLED=0 go build -o test server/cmd/test/main.go

CMD [ "./main" ]