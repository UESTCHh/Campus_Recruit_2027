from fastapi import APIRouter

from app.api.v1.chat import router as chat_router
from app.api.v1.debug import router as debug_router
from app.api.v1.health import router as health_router
from app.api.v1.utility import router as utility_router

router = APIRouter()


router.include_router(
    health_router,
)

router.include_router(
    chat_router,
)

router.include_router(
    utility_router,
)

router.include_router(
    debug_router,
)
