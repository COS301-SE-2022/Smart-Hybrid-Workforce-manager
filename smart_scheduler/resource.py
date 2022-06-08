from typing import Dict
import json
from datetime import datetime

"""
This module is responsible resource related operations
It currently assumes that rooms are flat and that no nesting
of rooms occur
"""

ENDPOINT: str = 'http://arche-api:8080/api/booking/information'

TIME_FMT: str = '%Y-%m-%dT%H:%M:%S.%fZ'


class Resource:
    """
    Class representing a resource
    """
    def __init__(self, resource_dict: Dict = None):
        self.resource = {
            "id": None,
            "room_id": None,
            "name": None,
            "location": None,
            "resource_type": "DESK",
            "decorations": {},
            "date_created": None
        }
        if resource_dict is not None:
            for key, val in resource_dict.items():
                self.resource[key] = val

    def __str__(self):
        return str(self.resource)


class ResourceEncoder(json.JSONEncoder):
    """
    Class used when encoding a resource
    """
    def default(self, o):
        # if passed in object is a datetime object
        if isinstance(o, datetime):
            return datetime.strftime(o, TIME_FMT)

        # if passed in object is a resource
        if isinstance(o, Resource):
            return o.resource

        # otherwise use the default encoder
        return json.JSONEncoder.default(self, o)


def parse_dict(resource: Dict[str, any]) -> Resource:
    """
    Parses a dict representing a resource
    :param resource:
    :return:
    """
    pass