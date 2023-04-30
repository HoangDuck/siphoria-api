pro:
	docker rmi -f web-service:1.0
	docker-compose up
dev:
	swag init -g cmd/main.go --output docs
	cd cmd && go run main.go -env dev

production:
	cd cmd && go run main.go -env pro

dbdev:
	swag init -g cmd/main.go --output docs
	cd cmd && go run main.go -env dbdev

migrate:
	cd cmd && go run main.go -env dbdev
swag:
	swag init -g cmd/main.go --output docs

fast:
	cd cmd && go run main.go -env dev

pushtest:
	swag init -g cmd/main.go --output docs
	git add docs
	git commit -m "Chore: update swagger"
	git push origin test