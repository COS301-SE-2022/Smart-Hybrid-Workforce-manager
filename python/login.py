import requests

def login_get_token(email: str = None, password: str = None) -> str:
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

print(login_get_token())