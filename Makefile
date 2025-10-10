swagger:
	cd cmd/api && swag init -g main.go -d .,../../pkg/dto -o ./docs
