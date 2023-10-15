FROM golang:1.21.3-bullseye as builder
LABEL maintainer="Jose Ramon Mañes - github.com/jrmanes"
ENV GOPATH /go
ENV GOBIN /go/bin
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/

######## Start a new stage from scratch #######
FROM alpine:3.18.4
ARG UID=10001
ARG USR_NAME=setget
ENV USR_HOME=/home/${USER_NAME}

RUN apk update \
    && apk add --no-cache bash \
    && adduser ${USR_NAME} \
    -D \
    -g ${USR_NAME} \
    -h ${USR_HOME} \
    -s /sbin/nologin \
    -u ${UID}

# Copy the Pre-built binary file from the previous stage
COPY --chown=${USR_NAME}:${USR_NAME} --from=builder /app/main .

USER ${USR_HOME}

# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./main"]