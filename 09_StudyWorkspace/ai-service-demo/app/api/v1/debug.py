from fastapi import APIRouter

from app.exceptions import AIServiceException

router = APIRouter(
    prefix="/debug",
    tags=["debug"],
)


@router.get("/error-test")
async def error_test_endpoint() -> None:
    raise AIServiceException(
        message="test error",
        code=50001,
        status_code=500,
    )
