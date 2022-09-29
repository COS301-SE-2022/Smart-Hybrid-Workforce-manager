import csv
import json
import os
from typing import List, Dict
import time
import requests
import config
import generate_user
import user
import webbrowser

def read_in_col(path: str) -> List:
    names: List
    with open(path) as f:
        csvreader = csv.reader(f)
        _names = list(csvreader)

    names: List = []
    for namel in _names:
        names.append(namel[0])
    return names


def login_get_token(call_count: List, email: str = None, password: str = None) -> str:
    call_count[0] += 1
    body: dict[str] = {
        "id": None,
        "secret": "admin",
        "active": None,
        "FailedAttempts": None,
        "LastAccessed": None,
        "Identifier": "admin@admin.com"
    }
    resp = requests.post("http://localhost:8100/api/user/login", json=body)
    return "bearer " + str(resp.json()["token"])


def retrieve_rooms_in_building(building_id: str, token: str, call_count: List) -> List[Dict]:
    call_count[0] += 1
    body: Dict[str] = {
        "building_id": building_id
    }
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/resource/room/information", json=body, headers=headers)
    return resp.json()


def get_resources_in_room(room_id: str, token: str, call_count: List) -> List[Dict]:
    call_count[0] += 1
    body: Dict[str] = {
        "room_id": room_id
    }
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/resource/information", json=body, headers=headers)
    return resp.json()


def get_resources_in_building(building_id: str, token: str, call_count: List) -> List[Dict]:
    call_count[0] += 1
    resources = []
    rooms = retrieve_rooms_in_building(building_id, token, call_count)
    for room in rooms:
        resources.extend(get_resources_in_room(room["id"], token, call_count))
    return resources


def get_resource_id_list(resources: List[Dict], resource_type: str) -> List[str]:
    resource_ids = []
    for resource in resources:
        if resource["resource_type"] == resource_type:
            resource_ids.append(resource["id"])
    return resource_ids


def get_teams(token: str, call_count: List) -> List[Dict]:
    call_count[0] += 1
    body = {}
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/team/information", json=body, headers=headers)
    return resp.json()


def get_roles(token: str, call_count: List) -> List[Dict]:
    call_count[0] += 1
    body = {}
    headers: Dict[str] = {
        "Authorization": token
    }
    resp = requests.post("http://localhost:8100/api/role/information", json=body, headers=headers)
    return resp.json()


def get_role_ids(token: str, call_count: List) -> List[str]:
    role_ids = []
    roles = get_roles(token, call_count)
    for role in roles:
        role_ids.append(role["id"])
    return role_ids


def get_team_ids(token: str, call_count: List) -> List[str]:
    team_ids = []
    teams = get_teams(token, call_count)
    for team in teams:
        team_ids.append(team["id"])
    return team_ids


if __name__ == "__main__":
    call_count = [0]
    conf = config.parse_config("./mock-data-config.json")
    print("Attempting login...")
    auth_token = login_get_token(call_count)
    print("TOKEN:", auth_token)
    print("Login succeeded.")
    print()

    resource_ids: List | None = None
    if conf["building_id"] is not None:
        print("Fetching resources in building...")
        resources = get_resources_in_building(conf["building_id"], auth_token, call_count)
        print("Resources fetched.")
        resource_ids = get_resource_id_list(resources, "DESK")
        print()

    print("Fetching teams...")
    team_ids = get_team_ids(auth_token, call_count)
    print("Teams fetched.")

    print()

    print("Fetching roles...")
    role_ids = get_role_ids(auth_token, call_count)
    print("Roles fetched.")

    print("Loading names and image URLs...")
    fnames = read_in_col("fnames.csv")
    lnames = read_in_col("lnames.csv")
    pictures = read_in_col("image_urls.txt")
    print("Loaded")

    generated_users: List = []
    gen = generate_user.UserGenerator(conf)
    print(f"Generating {conf['create_user']['num_users']} users...")
    for _ in range(conf['create_user']['num_users']):
        generated_users.append(gen.generate(fnames, lnames, profile_pics=pictures,
                                            desks=resource_ids, teams=team_ids, roles=role_ids))
    print(f"Generated.")
    print()
    with open("./preview.json", "w") as pf:
        pf.write(json.dumps(generated_users, indent=4))

    webbrowser.open("preview.json")

    user_input = input("Please close the preview.json file, and enter PROCEED to persist the user, or CANCEL to abort:")
    user_input = user_input.lstrip().rstrip().upper()
    while user_input != "PROCEED" and user_input != "CANCEL":
        user_input = input(
            "Invalid input: Please close the preview.json file, and enter PROCEED to persist the users, or CANCEL to abort")
        user_input = user_input.lstrip().rstrip().upper()

    try:
        os.remove("preview.json")
    except:
        ...

    if user_input == "CANCEL":
        print("Cancelling")
        exit(0)


    print("Persisting users...")
    all_success = True
    start = time.time()
    for _user in generated_users:
        print(f"     Persisting: {_user['email']}", end="")
        success = user.create_new_user(_user, auth_token, call_count)
        all_success = all_success and success
        print(":", ("SUCCESS" if success else "FAILED"))
    end = time.time()
    print("Done:", ("SUCCESS" if all_success else "FAILED"), "\nNum API calls:", call_count[0], "\nTime elapsed:", f'{end - start}s')
