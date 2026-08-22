from abc import ABC, abstractmethod

from app.llm.types import LLMResult


class BaseLLM(ABC):
    @abstractmethod
    async def generate(self, message: str) -> LLMResult:
        """Generate one response for a user message."""
