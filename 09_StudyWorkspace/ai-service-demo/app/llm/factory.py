# Factory 的返回类型不是：
# MockLLM
# 也不是：
# UppercaseLLM
# 而是它们共同的抽象类型：
# BaseLLM
# 这表示：
# Factory 可以返回任何符合 BaseLLM 接口的具体实现。
from app.llm.base import BaseLLM

# Factory 必须知道有哪些对象可以创建。
# 当前有两个具体实现：
# mock       → MockLLM
# uppercase  → UppercaseLLM
# Factory 知道具体类，调用者不需要知道具体类。
from app.llm.mock import MockLLM
from app.llm.uppercase import UppercaseLLM


def create_llm(provider: str) -> BaseLLM:
    """Create an LLM provider from its configured name."""

    # 标准化输入
    normalized_provider = provider.strip().lower()

    # 根据名称创建对象
    if normalized_provider == "mock":
        return MockLLM()

    if normalized_provider == "uppercase":
        return UppercaseLLM()

    raise ValueError(f"Unsupported LLM provider: {provider}")
