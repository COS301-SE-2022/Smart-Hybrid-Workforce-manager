import json
from pprint import pprint
from typing import Dict, List

import requests


def register_user(user: Dict, token: str, call_count: List) -> bool:
    call_count[0] += 1
    body: Dict[str] = user
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/user/register", json=body, headers=headers)
    return resp.status_code == 200


def get_user(email: str, token: str, call_count: List) -> Dict:
    call_count[0] += 1
    body: Dict[str] = {
        "email": email
    }
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/user/information", json=body, headers=headers)
    return resp.json()[0]


def update_user_info(user: Dict, token: str, call_count: List) -> bool:
    call_count[0] += 1
    body: Dict[str] = user
    body["id"] = get_user(user["email"], token, call_count)["id"]
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/user/update", json=body, headers=headers)
    return resp.status_code == 200


def create_user_team_assoc(user: Dict, token: str, call_count: List) -> bool:
    all_success = True
    for team_id in user["_teams"]:
        call_count[0] += 1
        body: Dict[str] = {
            "team_id": team_id,
            "user_id": get_user(user["email"], token, call_count)["id"]
        }
        headers: Dict[str] = {
            "Authorization": token
        }
        resp = requests.post("http://localhost:8100/api/team/user/create", json=body, headers=headers)
        all_success = all_success and resp.status_code == 200
    return all_success


def create_user_role_assoc(user: Dict, token: str, call_count: List) -> bool:
    all_success = True
    for role_id in user["_roles"]:
        call_count[0] += 1
        body: Dict[str] = {
            "role_id": role_id,
            "user_id": get_user(user["email"], token, call_count)["id"]
        }
        headers: Dict[str] = {
            "Authorization": token
        }
        resp = requests.post("http://localhost:8100/api/role/user/create", json=body, headers=headers)
        all_success = all_success and resp.status_code == 200
    return all_success


def create_new_user(user: Dict, token: str, call_count: List) -> bool:
    success_all = True
    success_all = success_all and register_user(user, token, call_count)
    success_all = success_all and update_user_info(user, token, call_count)
    success_all = success_all and create_user_team_assoc(user, token, call_count)
    success_all = success_all and create_user_role_assoc(user, token, call_count)
    return success_all
