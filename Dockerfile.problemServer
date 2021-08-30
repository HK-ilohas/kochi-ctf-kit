ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}

WORKDIR /kochi-ctf-kit
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build
RUN chmod +x ./problemServer
CMD ./problemServer
