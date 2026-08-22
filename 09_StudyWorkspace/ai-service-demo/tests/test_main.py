# 文件作用
# 这是现有 API 集成测试，可能已经验证：
# /api/v1/chat/mock
# 本任务暂时保留旧接口，正是为了避免这些测试立即失效。
# 我们还需要确认：
# 当前测试客户端如何创建；
# 请求 JSON 格式；
# 期望的回答文本；
# 空消息的 422 测试是否已经存在。

# 该文件通过 TestClient 测试完整 HTTP 调用链：
# HTTP 请求
#   ↓
# FastAPI
#   ↓
# Router
#   ↓
# Dependency Injection
#   ↓
# ChatService
#   ↓
# Provider
#   ↓
# HTTP 响应
from fastapi.testclient import TestClient

from app.main import app

client = TestClient(app)


def test_healthz() -> None:
    response = client.get("/api/v1/healthz")

    assert response.status_code == 200
    assert response.json() == {"status": "ok"}
    assert response.headers["x-request-id"]


def test_request_id_is_preserved() -> None:
    response = client.get(
        "/api/v1/healthz",
        headers={
            "X-Request-ID": "test-request-id",
        },
    )

    assert response.status_code == 200
    assert response.headers["x-request-id"] == "test-request-id"


def test_echo() -> None:
    response = client.post(
        "/api/v1/utility/echo",
        json={
            "message": "hello",
        },
    )

    assert response.status_code == 200
    assert response.json() == {
        "message": "hello",
        "length": 5,
    }


def test_echo_rejects_empty_message() -> None:
    response = client.post(
        "/api/v1/utility/echo",
        json={
            "message": "",
        },
    )

    assert response.status_code == 422

    body = response.json()

    assert body["code"] == 42200
    assert body["message"] == "request validation failed"
    assert body["request_id"]
    assert body["details"]


def test_chat() -> None:
    response = client.post(
        "/api/v1/chat",
        json={
            "message": "测试LLM",
            "session_id": "test-session",
        },
    )

    assert response.status_code == 200
    assert response.json() == {
        "answer": "模拟回答: 测试LLM",
        "session_id": "test-session",
        "model": "mock-model-v1",
    }


def test_mock_chat() -> None:
    response = client.post(
        "/api/v1/chat/mock",
        json={
            "message": "测试LLM",
            "session_id": "test-session",
        },
    )

    assert response.status_code == 200
    assert response.json() == {
        "answer": "模拟回答: 测试LLM",
        "session_id": "test-session",
        "model": "mock-model-v1",
    }


def test_custom_exception_handler() -> None:
    response = client.get("/api/v1/debug/error-test")

    assert response.status_code == 500

    body = response.json()

    assert body["code"] == 50001
    assert body["message"] == "test error"
    assert body["request_id"]


def test_not_found_uses_unified_error_format() -> None:
    response = client.get("/does-not-exist")

    assert response.status_code == 404

    body = response.json()

    assert body["code"] == 40400
    assert body["message"] == "Not Found"
    assert body["request_id"]
