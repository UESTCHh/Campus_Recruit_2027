from app.llm.base import BaseLLM
from app.llm.types import LLMResult


class ChatService:
    def __init__(
        self,
        llm: BaseLLM,
    ) -> None:
        self._llm = llm

    async def chat(
        self,
        message: str,
    ) -> LLMResult:

        return await self._llm.generate(message)
