import asyncio

from app.llm.base import BaseLLM
from app.llm.types import LLMResult


class MockLLM(BaseLLM):
    async def generate(self, message: str) -> LLMResult:
        await asyncio.sleep(0.05)
        return LLMResult(
            content=f"模拟回答: {message}",
            model="mock-model-v1",
            provider="mock",
        )
