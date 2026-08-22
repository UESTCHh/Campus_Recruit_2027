from app.llm.base import BaseLLM
from app.llm.types import LLMResult
from app.services.chat_service import ChatService


class FakeLLM(BaseLLM):
    async def generate(
        self,
        message: str,
    ) -> LLMResult:

        return LLMResult(
            content=f"fake:{message}",
            model="fake-model",
            provider="fake",
        )


def test_chat_service_returns_fake_result() -> None:

    import asyncio

    async def run() -> None:

        service = ChatService(
            llm=FakeLLM(),
        )

        result = await service.chat(
            "你好",
        )

        assert result.content == "fake:你好"
        assert result.model == "fake-model"
        assert result.provider == "fake"

    asyncio.run(run())
