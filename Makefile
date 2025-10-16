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