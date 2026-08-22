.PHONY: help build up down restart test lint format type ci logs shell


help:
	@echo "Available commands:"
	@echo "  make build     Build docker image"
	@echo "  make up        Start service"
	@echo "  make down      Stop service"
	@echo "  make restart   Restart service"
	@echo "  make test      Run tests"
	@echo "  make lint      Run ruff"
	@echo "  make format    Format code"
	@echo "  make type      Run mypy"
	@echo "  make ci        Run all checks"
	@echo "  make logs      Show logs"
	@echo "  make shell     Enter container"


build:
	docker compose build


up:
	docker compose up


down:
	docker compose down


restart:
	docker compose down
	docker compose up --build


test:
	uv run pytest -q


lint:
	uv run ruff check .


format:
	uv run ruff format .


type:
	uv run mypy app


ci:
	make lint
	make type
	make test
	make build


logs:
	docker compose logs -f


shell:
	docker exec -it ai-service-demo bash
