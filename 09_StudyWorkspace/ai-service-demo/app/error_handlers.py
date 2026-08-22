from typing import Any

from fastapi import FastAPI, Request
from fastapi.encoders import jsonable_encoder
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse
from starlette.exceptions import HTTPException as StarletteHTTPException

from app.exceptions import AIServiceException
from app.logger import logger


def get_request_id(request: Request) -> str:
    """从当前请求中读取 Request ID。"""
    return getattr(
        request.state,
        "request_id",
        "unknown",
    )


def register_exception_handlers(app: FastAPI) -> None:
    """向 FastAPI 应用注册全部异常处理器。"""

    @app.exception_handler(AIServiceException)
    async def handle_ai_service_exception(
        request: Request,
        exc: AIServiceException,
    ) -> JSONResponse:
        request_id = get_request_id(request)

        logger.error(
            "request_id=%s code=%s path=%s message=%s",
            request_id,
            exc.code,
            request.url.path,
            exc.message,
        )

        return JSONResponse(
            status_code=exc.status_code,
            content={
                "code": exc.code,
                "message": exc.message,
                "request_id": request_id,
            },
        )

    @app.exception_handler(RequestValidationError)
    async def handle_validation_exception(
        request: Request,
        exc: RequestValidationError,
    ) -> JSONResponse:
        request_id = get_request_id(request)
        details: list[dict[str, Any]] = jsonable_encoder(exc.errors())

        logger.warning(
            "request_id=%s code=42200 path=%s validation_failed",
            request_id,
            request.url.path,
        )

        return JSONResponse(
            status_code=422,
            content={
                "code": 42200,
                "message": "request validation failed",
                "request_id": request_id,
                "details": details,
            },
        )

    @app.exception_handler(StarletteHTTPException)
    async def handle_http_exception(
        request: Request,
        exc: StarletteHTTPException,
    ) -> JSONResponse:
        request_id = get_request_id(request)

        if isinstance(exc.detail, str):
            message = exc.detail
        else:
            message = "HTTP request failed"

        logger.warning(
            "request_id=%s code=%s path=%s message=%s",
            request_id,
            exc.status_code,
            request.url.path,
            message,
        )

        return JSONResponse(
            status_code=exc.status_code,
            content={
                "code": exc.status_code * 100,
                "message": message,
                "request_id": request_id,
            },
        )

    @app.exception_handler(Exception)
    async def handle_unexpected_exception(
        request: Request,
        exc: Exception,
    ) -> JSONResponse:
        request_id = get_request_id(request)

        logger.error(
            "request_id=%s code=50000 path=%s unexpected_error",
            request_id,
            request.url.path,
            exc_info=(
                type(exc),
                exc,
                exc.__traceback__,
            ),
        )

        return JSONResponse(
            status_code=500,
            content={
                "code": 50000,
                "message": "internal server error",
                "request_id": request_id,
            },
        )
