from fastapi import APIRouter

from app.logger import logger

router = APIRouter()


@router.get("/healthz")
async def healthz():

    logger.info("health check called")

    return {"status": "ok"}
