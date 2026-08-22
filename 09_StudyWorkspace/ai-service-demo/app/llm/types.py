from dataclasses import dataclass


# dataclass 减少样板代码；
# frozen=True 表示创建后不应修改；
# slots=True 限制随意添加属性并减少部分内存开销。
@dataclass(frozen=True, slots=True)
class LLMResult:
    content: str
    model: str
    provider: str
