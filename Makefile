.PHONY: build deploy dev clean

build:
	@go run cmd/build/main.go

deploy:
	@go run cmd/deploy/main.go

dev:
	@air

clean:
	@rm -rf dist/
	@rm -rf tmp/
	@echo "✓ Cleaned build artifacts"

