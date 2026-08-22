from uuid import uuid4

from fastapi import Request, Response
from starlette.middleware.base import (
    BaseHTTPMiddleware,
    RequestResponseEndpoint,
)


class RequestIDMiddleware(BaseHTTPMiddleware):
    """为每个 HTTP 请求添加唯一的 Request ID。"""

    header_name = "X-Request-ID"

    async def dispatch(
        self,
        request: Request,
        call_next: RequestResponseEndpoint,
    ) -> Response:
        incoming_request_id = request.headers.get(self.header_name)

        if incoming_request_id and 0 < len(incoming_request_id.strip()) <= 128:
            request_id = incoming_request_id.strip()
        else:
            request_id = uuid4().hex

        request.state.request_id = request_id

        response = await call_next(request)
        response.headers[self.header_name] = request_id

        return response
