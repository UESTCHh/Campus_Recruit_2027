from fastapi import APIRouter

from app.schemas import EchoRequest, EchoResponse

router = APIRouter(
    prefix="/utility",
    tags=["utility"],
)


@router.post(
    "/echo",
    response_model=EchoResponse,
)
async def echo_endpoint(
    request: EchoRequest,
) -> EchoResponse:
    return EchoResponse(
        message=request.message,
        length=len(request.message),
    )
