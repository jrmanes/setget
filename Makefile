PROJECT_NAME := $(shell basename `pwd`)
REPOSITORY_NAME := $(shell basename `pwd`)
REGISTRY_NAME=ghcr.io/jrmanes
ENV:="dev"
DIR := ${CURDIR}

go_fmt:
	go fmt ./...

# Docker
docker_compose_up_build:
	docker-compose up --build

docker_compose_up:
	docker-compose up
:password@localhost:5432/tyr?sslmode=disable -verbose down -all