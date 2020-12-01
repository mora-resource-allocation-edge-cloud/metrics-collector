FROM golang:1.15-buster
ENV GO111MODULE on
ARG MODULE_PATH=load-generator
ARG MODULE_NAME=load-generator
ENV MODULE_NAME=$MODULE_NAME
ENV MODULE_PATH=$MODULE_PATH
WORKDIR /$GOPATH/src/$MODULE_PATH

# Install dependencies
RUN apt-get update && apt-get install -y git make gcc tar rsync

# Live-Reload Go project
RUN go get github.com/cespare/reflex
RUN go get github.com/beego/bee
# Add utils scripts for the container
COPY docker/root/ /

# Populate the module cache based on the go.{mod,sum} files for maas
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Entrypoint.sh just set umask to avoid loss of permissions on volume,
# then exec any other command as speciefied below (i.e., reflex to auto-reload on code change)
# TODO test on other OSes
ENTRYPOINT ["entrypoint.sh"]
CMD ["reflex", "-c", "/etc/reflex/reflex.conf"]
