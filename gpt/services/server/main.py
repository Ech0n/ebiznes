from fastapi import FastAPI, Request, Form, Header
from fastapi.responses import HTMLResponse
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
import uvicorn
import httpx
import asyncio

app = FastAPI()
app.mount("/static", StaticFiles(directory="static"), name="static")
templates = Jinja2Templates(directory="templates")


@app.get("/", response_class=HTMLResponse)
async def get_chat(request: Request):
    return templates.TemplateResponse("chat.html", {"request": request})

@app.post("/send", response_class=HTMLResponse)
async def send_message(request: Request, message: str = Form(...), x_client_id: str = Header(None)):
    speaker_id = x_client_id or "anonymous"
    print(f"Sending message from client_id: {speaker_id}, message: {message}")
    response = await send_message_to_server(message, speaker_id)
    print(f"RESPONSE: \"{response}\"")
    
    return templates.TemplateResponse("message.html", {"request": request, "message": message, "response": response})

async def send_message_to_server(message: str, speaker_id: str) -> str:
    url = "http://chat:5000/api/chat"
    data = {"message": message, "speaker_id": speaker_id}

    async with httpx.AsyncClient() as client:
        response = await client.post(url, json=data)
        response.raise_for_status()
        json_resp = response.json()
        return json_resp.get("response", "")
    

@app.post("/start-chat", response_class=HTMLResponse)
async def start_chat(request: Request, x_client_id: str = Header(None)):
    url = "http://chat:5000/api/init"
    data = {"speaker_id": x_client_id}

    async with httpx.AsyncClient() as client:
        response = await client.post(url, json=data)
        response.raise_for_status()
        response_data = response.json()

    initial_message = response_data.get("message")
    print(f"Initial message {initial_message}")
    return templates.TemplateResponse(
        "start_chat_partial.html",
        {
            "request": request,
            "response": initial_message
        }
    )