# 文件作用
# 集中声明并读取应用配置，包括：
# 应用名称和版本
# 运行环境
# LLM Provider
# LLM Model
# 数据库连接 URL
from functools import lru_cache
from typing import Literal

from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    """
    Application settings loaded from environment variables.

    Values can come from:
    - system environment variables
    - .env file
    """

    app_name: str = "My First AI Backend"
    app_version: str = "0.2.0"
    environment: str = "dev"

    llm_provider: Literal["mock", "uppercase"] = "mock"
    llm_model: str = "mock-model-v1"

    database_url: str = "sqlite+aiosqlite:///./data/app.db"

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore",
    )


@lru_cache
def get_settings() -> Settings:
    """
    Cached settings instance.

    Avoid creating Settings repeatedly on every request.
    """
    return Settings()
