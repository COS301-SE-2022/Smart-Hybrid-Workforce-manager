import datetime
from tracemalloc import start
from turtle import st
from typing import Dict
import requests

def retrieve_rooms_in_building(start: datetime.datetime) -> int:
    body: Dict[str] = {
        "start_date": start.replace(tzinfo=None).isoformat() + "Z"
    }
    resp = requests.post("http://localhost:8100/api/scheduler/execute", json=body)
    return resp.status_code

start_date: datetime.datetime = datetime.datetime.now()
start_date = start_date - datetime.timedelta(days=100)
i = 0
while start_date <= (datetime.datetime.now() - datetime.timedelta(days=40)):
    print(i, ": ", start_date)
    start_date += datetime.timedelta(days=7)
    resp = retrieve_rooms_in_building(start_date)
    print("RESPONSE:", resp)
    i += 1