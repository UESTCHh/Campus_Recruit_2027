from datetime import datetime

from pydantic import BaseModel, Field


class HealthResponse(BaseModel):
    status: str
    time: datetime


class EchoRequest(BaseModel):
    message: str = Field(min_length=1, max_length=500)


class EchoResponse(BaseModel):
    message: str
    length: int


class ChatRequest(BaseModel):
    message: str = Field(min_length=1, max_length=2000)

    session_id: str = "demo-session"


class ChatResponse(BaseModel):
    answer: str
    session_id: str
    model: str


class ErrorResponse(BaseModel):
    code: int
    message: str
