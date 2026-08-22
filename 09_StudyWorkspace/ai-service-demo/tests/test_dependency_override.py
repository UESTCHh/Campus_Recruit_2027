# 文件作用
# 该文件用于验证：
# app.dependency_overrides[get_chat_service]
# 能否在 API 测试中把生产依赖替换为测试专用依赖。
# 为什么新增这个测试
# 当前普通 API 测试使用真实依赖链：
# Settings
# → Factory
# → MockLLM
# → ChatService
# 这会让测试结果受到配置和具体 Provider 的影响。
# 使用依赖覆盖后，测试可以明确指定：
# 本次测试必须使用 FakeLLM
# 从而做到：
# 不依赖 .env；
# 不依赖当前 Provider 配置；
# 不访问真实模型；
# 返回结果完全可控；
# 测试失败时更容易定位。

from collections.abc import Iterator

import pytest
from fastapi.testclient import TestClient

from app.dependencies import get_chat_service
from app.llm.base import BaseLLM
from app.llm.types import LLMResult
from app.main import app
from app.services.chat_service import ChatService


# 这是测试专用 Provider。
# 它：
# 不等待网络；
# 不读取配置；
# 不调用 Factory；
# 不调用真实模型；
# 返回固定且容易识别的结果。
class FakeLLM(BaseLLM):
    async def generate(self, message: str) -> LLMResult:
        return LLMResult(
            content=f"fake override: {message}",
            model="fake-override-model",
            provider="fake",
        )


def get_fake_chat_service() -> ChatService:
    return ChatService(llm=FakeLLM())


@pytest.fixture
def client_with_fake_chat_service() -> Iterator[TestClient]:
    # 它表示：
    # 当 FastAPI 准备调用 get_chat_service() 时
    # 不要调用原函数，改为调用 get_fake_chat_service()。
    app.dependency_overrides[get_chat_service] = get_fake_chat_service

    with TestClient(app) as client:
        yield client

    app.dependency_overrides.clear()
    # app 是多个测试共同使用的全局 FastAPI 应用对象。
    # 如果不清除：
    # app.dependency_overrides.clear()
    # 后续其他测试可能仍然使用 FakeLLM。


def test_chat_uses_dependency_override(
    client_with_fake_chat_service: TestClient,
) -> None:
    response = client_with_fake_chat_service.post(
        "/api/v1/chat",
        json={
            "message": "hello override",
            "session_id": "override-test",
        },
    )

    assert response.status_code == 200
    assert response.json() == {
        "answer": "fake override: hello override",
        "session_id": "override-test",
        "model": "fake-override-model",
    }
