build:
	go build -o build/gospring-demo main.go

stop:
	docker-compose down

pepare: stop
	docker-compose up --build

clean:
	rm -rf build/