# 文件作用
# 集中创建：
# 全局 AsyncEngine
# 全局 AsyncSession 工厂
# 每次使用时提供独立 Session 的异步依赖函数
# 为什么独立成文件
# 不能让 Router、Service 或 Repository 分别创建自己的 Engine。
# 数据库连接基础设施应集中管理。

from collections.abc import AsyncIterator

from sqlalchemy.ext.asyncio import (
    AsyncEngine,
    AsyncSession,
    async_sessionmaker,
    create_async_engine,
)

from app.config import get_settings

settings = get_settings()

database_engine: AsyncEngine = create_async_engine(
    settings.database_url,
    echo=False,
)

database_session_factory: async_sessionmaker[AsyncSession] = async_sessionmaker(
    bind=database_engine,
    expire_on_commit=False,
)


async def get_database_session() -> AsyncIterator[AsyncSession]:
    """
    Provide one database session for one unit of work.

    The session is closed automatically after the caller finishes.
    """

    async with database_session_factory() as session:
        yield session
