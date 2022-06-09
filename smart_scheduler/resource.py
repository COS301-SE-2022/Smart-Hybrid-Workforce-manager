from typing import Dict, List
import json
from datetime import datetime

import requests

"""
This module is responsible resource related operations
It currently assumes that rooms are flat and that no nesting
of rooms occur
"""

ENDPOINT: str = 'http://arche-api:8080/api/resource/information'

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
            for _key, _val in resource_dict.items():
                self.resource[_key] = _val

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
    :param resource: The resource dict to be parsed
    :return: A new Resource object
    """
    parsed_resource: Dict[str, any] = dict(resource)

    if "date_created" in parsed_resource and parsed_resource["date_created"] is not None:
        parsed_resource["date_created"] = datetime.strptime(parsed_resource["date_created"], TIME_FMT)

    return Resource(parsed_resource)


def fetch_resources(resource_filter: Resource = None) -> List[Resource]:
    """
    Uses the API to fetch all resources matching the passed filter,
    if no filter is passed then all resources will be fetched
    :param resource_filter: The filter to be used for fetching from the api
    :return: A list of Resources matching the filter
    """
    def hook(obj):
        # parse decorations as a dict
        value = obj.get("decorations")
        if value and isinstance(value, str):
            obj["decorations"] = json.loads(value, object_hook=hook)
        # TODO: @JonathanEnslin make user of hook to also parse date instead of parse_dict()
        return obj

    request: Dict | Resource = resource_filter if resource_filter is not None else {}
    req_data = json.dumps(request, cls=ResourceEncoder)
    resp: requests.Response = requests.post(ENDPOINT, data=req_data)
    resp_list: List[Dict] = json.loads(resp.content, object_hook=hook)
    return [parse_dict(_r) for _r in resp_list]


if __name__ == '__main__':
    ENDPOINT: str = 'http://localhost:8100/api/resource/information'
    _resources = fetch_resources()
    for r in _resources:
        print(r)
        for key, val in r.resource.items():
            print("   ", val, type(val))
