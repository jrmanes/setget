PROJECT_NAME := $(shell basename `pwd`)
REGISTRY_NAME=ghcr.io/jrmanes

# Docker
docker_build:
	docker build -t ${REGISTRY_NAME}/${PROJECT_NAME}:latest .
.PYONHY: docker_build

docker_push:
	docker push ${REGISTRY_NAME}/${PROJECT_NAME}:latest
.PYONHY: docker_push

docker_all: docker_build docker_push
.PYONHY: docker_all

# HELM
helm_dep:
	helm repo update
	helm dependency build ./infra/setget
.PYONHY: helm_dep

helm_install:
	helm install -f infra/setget/values.yaml setget ./infra/setget
.PYONHY: helm_install

helm_uninstall:
	helm uninstall setget
.PYONHY: helm_uninstall