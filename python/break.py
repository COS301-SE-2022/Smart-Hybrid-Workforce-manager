import datetime
import threading
from tracemalloc import start
from turtle import st
from typing import Dict
import requests

def login_get_token() -> str:
    body: dict[str] = {
        "id": None,
        "secret": None,
        "active": None,
        "FailedAttempts": None,
        "LastAccessed": None,
        "Identifier": None
    }
    resp = requests.post("http://localhost:8100/api/user/login", json=body)
    return "bearer " + str(resp.json()["token"])

def book(token: str, i) -> int:
    body: Dict[str] = {
        "user_id": "00000000-0000-0000-0000-000000000000",
        "resource_type": "DESK",
        "resource_preference_id": None,
        "start": f"2022-11-0{i+1}T12:12:00.000Z",
        "end": f"2022-11-0{i+1}T13:12:00.000Z",
        "automated": True,
        "booked": False
    }
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/booking/create", json=body, headers=headers)
    print(i, "STATUS", resp)
    return resp.status_code

token = login_get_token()
print("logged in: ", token)
for i in range(5):
    # book(token, i)
    x = threading.Thread(target=book, args=(token,2))
    x.start()
x.join()