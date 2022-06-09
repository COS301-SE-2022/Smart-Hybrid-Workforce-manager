import json
from typing import Dict, List
from datetime import datetime

"""
This module is responsible for user related operations
It currently assumes that teams are flat
"""

ENDPOINT: str = 'http://arche-api/api/user/information'

TIME_FMT: str = '%Y-%m-%dT%H:%M:%S.%fZ'


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

