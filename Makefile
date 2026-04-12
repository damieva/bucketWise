MONGO_CONTAINER_NAME := local-mongodb
MONGO_PORT := 27017
MONGO_VOLUME := mongo_data
MONGO_DATABASE := bucketWise

.PHONY: swagger run start-up-local-env stop-local-env clean-local-env

swagger:
	cd cmd/api && swag init \
		-g main.go \
		-d .,../../pkg/dto \
		-o ./docs

run: swagger
	 go run cmd/api/main.go

start-up-local-env:
	@echo "🚀 Iniciando entorno local con MongoDB y colecciones iniciales..."
	docker volume create $(MONGO_VOLUME)
	docker run -d \
		--name $(MONGO_CONTAINER_NAME) \
		-p $(MONGO_PORT):27017 \
		-v $(MONGO_VOLUME):/data/db \
		-v $(PWD)/scripts/mongo_init.sh:/docker-entrypoint-initdb.d/mongo_init.sh \
		-e MONGO_INITDB_DATABASE=$(MONGO_DATABASE) \
		mongo:8.0
	@echo "✅ MongoDB con colecciones iniciales en el puerto $(MONGO_PORT)"

stop-local-env:
	@echo "🛑 Deteniendo contenedor MongoDB..."
	-docker stop $(MONGO_CONTAINER_NAME)
	-docker rm $(MONGO_CONTAINER_NAME)

clean-local-env: stop-local-env
	@echo "🧹 Eliminando volumen MongoDB..."
	-docker volume rm $(MONGO_VOLUME)

docker-build:
	docker build -t $(IMAGE):$(TAG) .

docker-run:
	docker run --rm \
		-p 8080:8080 \
		-e MONGO_URI=mongodb://host.docker.internal:$(MONGO_PORT) \
		-e MONGO_DATABASE=$(MONGO_DATABASE) \
		$(IMAGE):$(TAG)

docker-push:
	docker push $(IMAGE):$(TAG)

docker-release: docker-build docker-push

help:
	@echo ""
	@echo "BucketWise Makefile"
	@echo ""
	@echo "General:"
	@echo "  make swagger                Generate Swagger docs"
	@echo "  make run                    Run API locally (Go)"
	@echo ""
	@echo "Local environment:"
	@echo "  make start-up-local-env     Start local MongoDB"
	@echo "  make stop-local-env         Stop local MongoDB"
	@echo "  make clean-local-env        Remove MongoDB container and volume"
	@echo ""
	@echo "Docker (TAG is required):"
	@echo "  make docker-build TAG=x     Build Docker image"
	@echo "  make docker-push  TAG=x     Push Docker image to registry"
	@echo "  make docker-release TAG=x  Build and push image"
	@echo ""
	@echo "Examples:"
	@echo "  make docker-build IMAGE=ghcr.io/damieva/bucketwise TAG=v0.1.0"
	@echo "  make docker-push IMAGE=ghcr.io/damieva/bucketwise TAG=v0.1.0"
	@echo "  make docker-release IMAGE=ghcr.io/damieva/bucketwise TAG=v0.1.0"
	@echo ""
