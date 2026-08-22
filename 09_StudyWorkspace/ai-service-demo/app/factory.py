from fastapi import FastAPI

from app.api.v1.router import router as api_router
from app.config import get_settings
from app.error_handlers import register_exception_handlers
from app.logger import setup_logging
from app.middleware.request_id import RequestIDMiddleware


def create_app() -> FastAPI:
    settings = get_settings()

    setup_logging()

    app = FastAPI(
        title=settings.app_name,
        version=settings.app_version,
        description="Production-style AI backend API",
    )

    app.add_middleware(
        RequestIDMiddleware,
    )

    register_exception_handlers(app)

    app.include_router(
        api_router,
        prefix="/api/v1",
    )

    return app
