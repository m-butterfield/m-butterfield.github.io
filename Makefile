cloudrunbasecommand := gcloud run deploy --project=mattbutterfield --region=us-central1 --platform=managed
deployservercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/mattbutterfield.com mattbutterfield
deployworkercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/mattbutterfield.com-worker mattbutterfield-worker

terraformbasecommand := cd infra && terraform
terraformvarsarg := -var-file=secrets.tfvars

export DB_SOCKET=host=localhost dbname=mattbutterfield

build: build-server build-worker

build-server:
	go build -o bin/server cmd/server/main.go

build-worker:
	go build -o bin/worker cmd/worker/main.go

deploy: docker-build docker-push
	$(deployservercommand)
	$(deployworkercommand)

deploy-server: docker-build-server docker-push-server
	$(deployservercommand)

deploy-worker: docker-build-worker docker-push-worker
	$(deployworkercommand)

docker-build:
	docker-compose build

docker-build-server:
	docker-compose build server

docker-build-worker:
	docker-compose build worker

docker-push:
	docker-compose push

docker-push-server:
	docker-compose push server

docker-push-worker:
	docker-compose push worker

reset-db:
	dropdb --if-exists social
	createdb social
	go run cmd/migrate/main.go

migrate:
	go run cmd/migrate/main.go

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix
	cd infra/ && terraform fmt

run-server: export USE_LOCAL_FS=true
run-server: export SQL_LOGS=true
run-server: export WORKER_BASE_URL=http://localhost:8001/
run-server:
	go run cmd/server/main.go

run-worker: export SQL_LOGS=true
run-worker:
	go run cmd/worker/main.go

test: export DB_SOCKET=host=localhost dbname=mattbutterfield_test
test:
	dropdb --if-exists mattbutterfield_test && createdb mattbutterfield_test && psql -d mattbutterfield_test -f schema.sql
	go test -v ./app/...

tf-plan:
	$(terraformbasecommand) plan $(terraformvarsarg)

tf-apply:
	$(terraformbasecommand) apply $(terraformvarsarg)

tf-refresh:
	$(terraformbasecommand) apply $(terraformvarsarg) -refresh-only

update-deps:
	go get -u ./...
	go mod tidy
	npm upgrade
	cd infra && terraform init -upgrade && cd -
