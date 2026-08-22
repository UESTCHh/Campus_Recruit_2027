# 文件作用
# 负责：
# 定义聊天相关的 HTTP 路径；
# 接收和校验 ChatRequest；
# 通过 Depends 获取 ChatService；
# 调用聊天业务；
# 将 LLMResult 转换成 ChatResponse。
from fastapi import APIRouter, Depends

from app.dependencies import get_chat_service
from app.schemas import ChatRequest, ChatResponse
from app.services.chat_service import ChatService

router = APIRouter(
    prefix="/chat",
    tags=["chat"],
)

# 两个路由装饰器
# 它们把同一个函数注册为两个 HTTP 接口。
# 通用接口
# @router.post("", response_model=ChatResponse)
# 由于 Router 已有：
# prefix="/chat"
# 空字符串最终形成：
# /chat
# 再与应用的 /api/v1 组合，最终得到：
# POST /api/v1/chat
# 旧接口
# @router.post("/mock", response_model=ChatResponse, deprecated=True)
# 最终仍然是：
# POST /api/v1/chat/mock
# deprecated=True 表示：
# 这个接口暂时仍能使用，但不建议新客户端继续使用。


# 它会在 Swagger 文档中显示为已弃用。
@router.post("", response_model=ChatResponse)
@router.post("/mock", response_model=ChatResponse, deprecated=True)
# 为什么两个接口共用一个函数
# 不推荐复制两个函数：
# async def chat_endpoint(...):
#     ...
# async def mock_chat_endpoint(...):
#     ...
# 因为这样会重复：
# 调用 ChatService；
# 读取 result.content；
# 组装 ChatResponse；
# 维护相同业务逻辑。
# 当前设计：
# /api/v1/chat
#           ┐
#           ├──→ chat_endpoint()
#           │         ↓
# /chat/mock┘    ChatService
# 两个地址共用一套实现。
async def chat_endpoint(
    request: ChatRequest,
    chat_service: ChatService = Depends(get_chat_service),
) -> ChatResponse:
    result = await chat_service.chat(
        request.message,
    )

    return ChatResponse(
        answer=result.content,
        session_id=request.session_id,
        model=result.model,
    )
