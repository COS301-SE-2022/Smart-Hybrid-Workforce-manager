from typing import Dict, List, Set

import resource
from resource import Resource
import operator

"""
This module is responsible for room related operations
"""


# Current version does not call endpoint, infers rooms from passed in resources

class Rooms:
    """
    Class representing Rooms, an object of this class
    will contain a Dict that uses room_ids as keys, the elements
    of this dict will be sets containing team_ids, a more efficient
    solution can still be investigated
    """

    def __init__(self, resources: List[Resource] = None):
        """
        Initialises the object with a list of resource, infers the rooms from the resources
        """
        self.rooms: Dict[str, Set[str]] = {}
        if resources is not None:
            for r in resources:
                # TODO: @JonathanEnslin implement getters and setters for all the classes
                if r.resource["id"] is None:
                    continue  # only add resources that belong to a room
                if r.resource["room_id"] not in self.rooms:
                    # add room
                    self.rooms[r.resource["room_id"]] = set()
                self.rooms[r.resource["room_id"]].add(r.resource["id"])

    def add_resource(self, resource: Resource):
        """
        Adds the given resource to the relevant room, creates the room if it does not already exist
        :param resource: The resource to add the room
        """
        if resource.resource["room_id"] not in self.rooms:
            self.rooms[resource.resource["room_id"]] = set()
        self.rooms[resource.resource["room_id"]].add(resource.resource["id"])

    def remove_resource(self, resource: Resource):
        """
        Removes the passed resource if it is contained
        """
        if resource.resource["room_id"] in self.rooms:
            self.rooms[resource.resource["room_id"]].discard(resource.resource["id"])

    def rooms_size(self, size: int, compare_with: str = 'eq') -> List[str]:
        """
        Returns a list of room ids that have the specified number of resources contained in them
        :param size: The size room that should be used while searching
        :param compare_with: The type of comparison to use, 'eq' for exact size matches, 'lt' for rooms with
         size less than the specified size or 'gt' for rooms greater than the specified size, or le for <= and ge for >=
         to 'eq' if an invalid operator is passed
        """
        comparator: operator
        comparators: Dict[str, any] = {
            'eq': operator.eq,
            'gt': operator.gt,
            'lt': operator.lt,
            'le': operator.le,
            'ge': operator.ge,
        }
        if compare_with not in comparators:
            comparator = comparators['eq']
        else:
            comparator = comparators[compare_with]

        rooms: List[str] = []
        for _key, _val in self.rooms.items():
            if comparator(len(_val), size):
                rooms.append(_key)
        return rooms

    def __str__(self):
        rooms_str: str = '{\n'
        for _key, _val in self.rooms.items():
            rooms_str += f'"{_key}": {_val},\n'
        return rooms_str + '}'

def room_size(room_list: List[Dict[str, str]], size: int, compare_with: str = 'eq') -> Dict[str, List[str]]:
        """
        Returns a list of room ids that have the specified number of resources contained in them
        :param room_list: The list of rooms to search
        :param size: The size room that should be used while searching
        :param compare_with: The type of comparison to use, 'eq' for exact size matches, 'lt' for rooms with
         size less than the specified size or 'gt' for rooms greater than the specified size, or le for <= and ge for >=
         to 'eq' if an invalid operator is passed
        """
        comparator: operator
        comparators: Dict[str, any] = {
            'eq': operator.eq,
            'gt': operator.gt,
            'lt': operator.lt,
            'le': operator.le,
            'ge': operator.ge,
        }
        if compare_with not in comparators:
            comparator = comparators['eq']
        else:
            comparator = comparators[compare_with]

        rooms: Dict[str, List[str]] = []
        for _room in room_list:
            if comparator(len(_room['resource_ids']), size):
                rooms.append(_room['id'])
        return rooms

if __name__ == '__main__':
    # resource.ENDPOINT = 'http://localhost:8100/api/resource/information'
    # _resources = resource.fetch_resources()
    # for r in _resources:
    #     print(r)
        # for key, val in r.resource.items():
        #     print("   ", val, type(val))
    _rooms: Rooms = Rooms()
    _rooms.add_resource(Resource({"id": "123-123", "room_id": "456-456"}))
    _rooms.add_resource(Resource({"id": "321-321", "room_id": "456-456"}))
    _rooms.add_resource(Resource({"id": "123-321", "room_id": "456-456"}))
    _rooms.add_resource(Resource({"id": "321-123", "room_id": "456-456"}))
    _rooms.add_resource(Resource({"id": "321-123", "room_id": "456-456"}))

    _rooms.add_resource(Resource({"id": "789-987", "room_id": "753-357"}))
    _rooms.add_resource(Resource({"id": "951-159", "room_id": "753-357"}))
    print(_rooms)
    print(''.join(_rooms.rooms_size(2, 'gt')))
    print(', '.join(_rooms.rooms_size(2, 'ge')))
    # TODO: @JonathanEnslin add proper tests
