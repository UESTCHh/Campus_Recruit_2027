# AI Service Demo

# AI Service Demo

A production-style AI backend service built with **FastAPI**.

This project demonstrates a clean backend architecture with:
- layered application structure
- mock LLM service integration
- unified error handling
- request ID middleware
- Docker deployment
- automated testing


## Features

- FastAPI REST API
- Pydantic v2 configuration management
- Middleware based request tracing
- Unified JSON error responses
- Mock AI service layer
- Docker Compose deployment
- Health check endpoint
- Pytest test suite


## Project Structure

ai-service-demo/
│
├── app/
│   ├── llm/
│   │   ├── base.py
│   │   └── mock.py
│   │
│   ├── middleware/
│   │   └── request_id.py
│   │
│   ├── services/
│   │   └── chat_service.py
│   │
│   ├── config.py
│   ├── error_handlers.py
│   ├── exceptions.py
│   ├── logger.py
│   ├── main.py
│   └── schemas.py
│
├── tests/
│
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── pyproject.toml


## Requirements

- Python 3.12+
- Docker
- Docker Compose


## Run locally

Install dependencies:

```bash
uv sync
Start development server:
uv run fastapi dev app/main.py
API documentation:
http://127.0.0.1:8000/docs
Run with Docker
Build image:
make build
Start service:
make up
Stop service:
make down
API Endpoints
Health Check
GET /healthz
Example:
{
  "status": "ok"
}
Mock Chat
POST /chat/mock
Example request:
{
  "message": "Hello AI"
}
Echo
POST /echo
Testing
Run tests:
make test
Current test status:
7 passed
Development Commands
Show available commands:
make help
Available:
make build
make up
make down
make restart
make test
make lint
make logs
make shell
Docker Health Check
The container includes health monitoring:
GET http://localhost:8000/healthz
Docker automatically marks the service healthy after startup.
License
MIT

---

保存：

nano：A production-style AI backend service built with **FastAPI**.

This project demonstrates a clean backend architecture with:
- layered application structure
- mock LLM service integration
- unified error handling
- request ID middleware
- Docker deployment
- automated testing


## Features

- FastAPI REST API
- Pydantic v2 configuration management
- Middleware based request tracing
- Unified JSON error responses
- Mock AI service layer
- Docker Compose deployment
- Health check endpoint
- Pytest test suite


## Project Structure
