
# main.py
from fastapi import FastAPI, Request, Form
from fastapi.responses import HTMLResponse, JSONResponse
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from pydantic import BaseModel
from chat_client import ChatClient

app = FastAPI()
templates = Jinja2Templates(directory="templates")

chatClient = ChatClient()

def analyze_and_respond(user_input: str) -> str:
    return f"Echo: {user_input}"

class ChatRequest(BaseModel):
    message: str
    speaker_id: str

@app.post("/api/chat")
async def chat_api(req: ChatRequest):
    response = chatClient.ask(req.message, req.speaker_id)
    return {"response": response}


class InitRequest(BaseModel):
    speaker_id: str

@app.post("/api/init")
async def init_chat(request: InitRequest):
    message = None
    if request.speaker_id != "anonymous":
        message = chatClient.initiate_chat(request.speaker_id)
    return JSONResponse(
        content={
            "status": "ok",
            "new_chat": message is not None,
            "message": message
        }
    )