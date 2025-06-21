
test:
	go test -v ./...

run:
	docker compose up -d --build

mocks:
	mockery