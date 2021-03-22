#----
# Build stage
#----
FROM golang:1.15-buster as buildstage


WORKDIR /$GOPATH/src/mailer
ARG MODULE_NAME=metrics-collector
ENV MODULENAME=$MODULE_NAME
# Install git
RUN apt-get update && apt-get install -y git make gcc tar rsync

# Enable go modules
ENV GO111MODULE on

# Populate the module cache based on the go.{mod,sum} files for maas
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

# Compile with static linking
RUN make build distMode=dir ignoreMissing=yes
RUN mv ./dist/$MODULE_NAME/$MODULE_NAME /service

#----
# Microservice stage
#----
FROM scratch
LABEL maintainer="alessandro.distefano@phd.unict.it"
LABEL author="aleskandro"
LABEL name="mora-metrics-collector"
LABEL description="Metrics collector with Mongo Driver for MORA real testbed"

COPY ./conf/app.prod.conf ./conf/app.conf
# Copy built executable
COPY --from=buildstage /service ./
CMD ["./service"]
