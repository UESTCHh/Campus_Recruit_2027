# 文件作用
# 这个文件负责组装应用运行时需要的对象：
# 读取配置
#     ↓
# 创建 Provider
#     ↓
# 创建 ChatService
#     ↓
# 交给 FastAPI Router
# 它属于应用的“组装层”。
from app.config import get_settings
from app.llm.factory import create_llm
from app.services.chat_service import ChatService


def get_chat_service() -> ChatService:
    settings = get_settings()
    llm = create_llm(settings.llm_provider)

    return ChatService(llm=llm)


# 完整职责划分是：
# Settings：提供配置名称

# Factory：根据名称创建 Provider

# Dependency：组装 Provider 和 ChatService

# ChatService：使用 Provider

# Router：使用 ChatService
