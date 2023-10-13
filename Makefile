PROJECT_NAME := $(shell basename `pwd`)
REPOSITORY_NAME := $(shell basename `pwd`)
REGISTRY_NAME=ghcr.io/jrmanes
ENV:="dev"
DIR := ${CURDIR}
DIR_MIGRATIONS := "${DIR}/db/migrations"

clear:
	clear

go_fmt:
	gofmt ./...

deploy: clear line build
	echo "---> Deploying Service..." &&\
	cd infra &&\
	terraform workspace select ${ENV} &&\
	terraform plan -out ${ENV}.tfplan &&\
	terraform apply ${ENV}.tfplan
	#terraform apply --auto-approve

destroy: line
	cd infra &&\
	terraform workspace select ${ENV} &&\
	terraform destroy --auto-approve


# Docker
docker_compose_up_build:
	docker-compose up --build

docker_compose_up:
	docker-compose up

build_local:
	GOOS=linux GOARCH=amd64 go build -o ./${PROJECT_NAME} ./cmd/server/main.go || exit 1

dc_up_build_local: build_local
	#docker-compose up -d db
	make db_migrate_up
	#docker-compose down db
	docker-compose up --build

dc_up_infra:
	docker-compose -f docker-compose-infra.yaml up


# DB migrations
db_migrate_init:
	docker run -v ${DIR_MIGRATIONS}:/migrations --network host migrate/migrate \
				-path=/migrations/ \
				create -ext sql -dir migrations/ -seq init_schema

db_migrate_create:
	docker run -v ${DIR_MIGRATIONS}:/migrations --network host migrate/migrate \
 				-path=/migrations/ \
				-verbose create -ext sql -dir migrations/ -seq create_${NAME}_table

db_migrate_up:
	docker run -v ${DIR_MIGRATIONS}:/migrations --network host migrate/migrate \
				-path=/migrations/ \
 				-database postgres://postgres:password@localhost:5432/tyr?sslmode=disable -verbose up

db_migrate_down:
	docker run -v ${DIR_MIGRATIONS}:/migrations --network host migrate/migrate \
				-path=/migrations/ \
 				-database postgres://postgres:password@localhost:5432/tyr?sslmode=disable -verbose down -all