import asyncio

# UppercaseLLM 必须遵守 BaseLLM 规定的接口。
# 关系是：
# BaseLLM：规定 generate() 应该长什么样
# UppercaseLLM：提供 generate() 的具体实现
# 对应上一个实验：
# Animal       → BaseLLM
# Bird         → UppercaseLLM
# speak()      → generate()
from app.llm.base import BaseLLM

# 无论是：
# MockLLM
# UppercaseLLM
# 未来的 OpenAILLM
# 未来的 DeepSeekLLM
# 都统一返回：
# LLMResult
# 调用者便可以统一访问：
# result.content
# result.model
# result.provider
from app.llm.types import LLMResult


# UppercaseLLM 是 BaseLLM 的一个具体实现
# 由于 BaseLLM 中的 generate() 是抽象方法
# 所以 UppercaseLLM 必须实现它，否则不能实例化。
class UppercaseLLM(BaseLLM):
    """Return the user message converted to uppercase."""

    # 这里虽然只是将字符串转大写，但仍然使用 async def。
    # 原因不是转大写需要异步，而是：
    # 所有 Provider 必须遵守相同的异步接口。
    # 真实 Provider 以后会在这里执行：
    # response = await http_client.post(...)
    # 这样 ChatService 不需要区分某个 Provider 是同步还是异步。
    async def generate(self, message: str) -> LLMResult:
        # 这不是业务需要，而是模拟真实模型请求可能存在的网络等待。
        # 它不会阻塞整个事件循环。

        await asyncio.sleep(0.05)

        # 转换大写
        # Provider 元数据
        return LLMResult(
            content=message.upper(),
            model="uppercase-model-v1",
            provider="uppercase",
        )
