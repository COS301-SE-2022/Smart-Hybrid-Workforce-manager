import json
from typing import Dict, List
from datetime import datetime

import requests

"""
This module is responsible for user related operations
It currently assumes that teams are flat
"""

ENDPOINT: str = 'http://arche-api/api/user/information'

# TIME_FMT: str = '%Y-%m-%dT%H:%M:%SZ'
TIME_FMT: str = '0000-01-01T%H:%M:%SZ'


class User:
    """
    Class representing a user
    """
    def __init__(self, user_dict: Dict = None):
        self.user = {
            "id": None,
            "identifier": None,
            "first_name": None,
            "last_name": None,
            "email": None,
            "picture": None,
            "date_created": None,
            "work_from_home": False,
            "parking": 'STANDARD',
            "office_days": 0,
            "preferred_start_time": None,
            "preferred_end_time": None,
        }

        if user_dict is not None:
            for _key, _val in user_dict.items():
                self.user[_key] = _val

    def __str__(self):
        return str(self.user)


class UserEncoder(json.JSONEncoder):
    """
    Used to properly encode a user
    """
    def default(self, o):
        # if passed in object is a datetime instance
        if isinstance(o, datetime):
            return datetime.strftime(o, TIME_FMT)

        # if passed in object is a User
        if isinstance(o, User):
            return o.user

        # otherwise default
        return json.JSONEncoder.default(self, o)

# no need for parse dict since a hook will be used


def user_hook(obj: any):
    # TODO: @JonathanEnslin check error handling
    value = obj.get('preferred_start_time')
    if value and isinstance(value, str):
        obj['preferred_start_time'] = datetime.strptime(obj['preferred_start_time'], TIME_FMT).time()

    value = obj.get('preferred_end_time')
    if value and isinstance(value, str):
        obj['preferred_end_time'] = datetime.strptime(obj['preferred_end_time'], TIME_FMT).time()
    return obj


def fetch_users(user_filter: User = None) -> List[User]:
    """
    Uses the API to fetch users, fetches with the passed filter, or all users
    if no filter is passed, fetches all by default
    :param user_filter: The filter to be applied or None to fetch all
    :return: A list of Users
    """
    request: Dict | User = user_filter if user_filter is not None else {}
    req_data = json.dumps(request, cls=UserEncoder)
    # TODO: @JonathanEnslin do err handling for all fetch functions
    resp: requests.Response = requests.post(ENDPOINT, data=req_data)
    resp_list: List[Dict] = json.loads(resp.content, object_hook=user_hook)
    return [User(u) for u in resp_list]


if __name__ == '__main__':
    ENDPOINT: str = 'http://localhost:8100/api/user/information'
    _resources = fetch_users()
    for r in _resources:
        print(r)
