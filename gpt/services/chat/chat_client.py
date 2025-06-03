import os
from dotenv import load_dotenv
from groq import Groq
import random

load_dotenv()

OPENINGS = [
    "Cześć! W czym mogę Ci pomóc dzisiaj w naszym sklepie z ubraniami?",
    "Witaj! Szukasz czegoś konkretnego w naszej kolekcji odzieży?",
    "Hej! Chętnie pomogę Ci znaleźć idealny strój. Co Cię interesuje?",
    "Dzień dobry! Jak mogę pomóc w wyborze ubrań?",
    "Cześć! Zapraszam do zadawania pytań o nasze ubrania i promocje!"
]

CLOSINGS = [
    "Dziękuję za rozmowę, miłego dnia!",
    "Jeśli masz jeszcze pytania, śmiało pytaj!",
    "Do zobaczenia! Zapraszamy ponownie!",
    "Cieszę się, że mogłem pomóc. Na razie!",
    "Trzymaj się! W razie potrzeby jestem do dyspozycji."
]
closing_statements = "\", \"".join(CLOSINGS)
PRECONTEXT = f"Jesteś asystentem w internetowym sklepie z ubraniami. Odpowiadaj wyłącznie na pytania związane z odzieżą i funkcjonowaniem sklepu. Ignoruj pytania z innych dziedzin. Jeżeli klient będzie chciał zakończyć rozmowe, albo się z tobą pożegna użyj dokładnie któregoś z podanych zwrotów kończących: \"{closing_statements}\" "


class ChatClient:
    def __init__(self):
        api_key = os.getenv("GROQ_API_KEY")
        if not api_key:
            raise ValueError("Missing GROQ_API_KEY in environment variables")
        self.client = Groq(api_key=api_key)
        self.contexts = {}
        print(f"Precontext {PRECONTEXT}")

    def initiate_chat(self, speaker_id: str):
        random_opening = random.choice(OPENINGS)
        self.contexts[speaker_id] = [{"role": "system", "content":PRECONTEXT },{"role": "assistant", "content": random_opening}]
        return random_opening

    def ask(self, prompt: str,speaker_id: str) -> str:
        track_context = speaker_id != "anonymous"

        if track_context:
            if speaker_id not in self.contexts:
                self.initiate_chat(speaker_id)
            self.contexts[speaker_id].append({"role": "user", "content": prompt})
            messages = self.contexts[speaker_id]
        else:
            messages = [{"role": "user", "content": prompt}]

        
        completion = self.client.chat.completions.create(
            model="llama3-70b-8192",
            messages=[
                {
                    "role": "user",
                    "content": prompt
                }
            ],
            temperature=1,
            max_completion_tokens=1024,
            top_p=1,
            stream=True,
            stop=None,
        )

        full_response = ""
        for chunk in completion:
            full_response += chunk.choices[0].delta.content or ""

        if track_context:
            self.contexts[speaker_id].append({"role": "assistant", "content": full_response})
        print(f"FULL CONVERSATION: {self.contexts[speaker_id]}")
        return full_response
