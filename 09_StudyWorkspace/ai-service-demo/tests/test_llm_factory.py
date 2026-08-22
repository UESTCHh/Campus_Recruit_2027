# 文件作用
# 该文件专门测试：
# create_llm()
# 是否能：
# 创建 MockLLM；
# 创建 UppercaseLLM；
# 拒绝未知 Provider。
import pytest

from app.llm.factory import create_llm
from app.llm.mock import MockLLM
from app.llm.uppercase import UppercaseLLM


def test_create_llm_returns_mock_provider() -> None:
    llm = create_llm("mock")

    # isinstance 直接检查对象与类之间的关系。
    # 它比比较类名字符串更可靠，因为类名字符串可能被重命名或手工拼错。
    assert isinstance(llm, MockLLM)


def test_create_llm_returns_uppercase_provider() -> None:
    llm = create_llm("uppercase")

    assert isinstance(llm, UppercaseLLM)


def test_create_llm_rejects_unknown_provider() -> None:
    with pytest.raises(
        ValueError,
        match="Unsupported LLM provider: invalid",
    ):
        create_llm("invalid")
    # 这段代码表示：
    # 执行 create_llm("invalid") 时，预期它抛出 ValueError。

    # 如果没有抛出错误，测试失败。
    # 如果抛出的不是 ValueError，测试也会失败。
