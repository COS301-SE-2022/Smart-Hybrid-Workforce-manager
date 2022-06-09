import json
from typing import Dict, List, Set

import requests

"""
This module is responsible for role related operations
"""

ENDPOINT: str = 'http://arche-api:8080/api/role/user/information'


class Roles:
    """
    Class representing roles, an object of this class
    will contain a Dict that uses role_ids as keys, the elements
    of this dict will be sets containing user_ids, a more efficient
    solution can still be investigated, a separate dict will also be
    maintained containing the date created for each role
    """

    def __init__(self, roles: Dict[str, Set] = None, roles_meta: Dict[str, List] = None):
        if roles is None:
            self.roles: Dict[str, Set[str]] = {}  # set is used to ensure users aren't added twice
        if roles_meta is None:  # not yet used TODO: @JonathanEnslin see if needed
            self.roles_meta: Dict = {}

    def add_role(self, role_id: str, user_ids: List[str] = None):
        """
        Adds a role or adds a list of users to a role
        :param role_id: The ID of the role to be added or that users should be added to
        :param user_ids:
        """
        if role_id not in self.roles:
            # add role
            self.roles[role_id] = set()  # create empty set
        # add users that are not already contained
        for _id in user_ids:
            self.roles[role_id].add(_id)

    def add_user(self, role_id: str, user_id: str):
        """
        Adds a user to a role, role is created if it does not
        yet exist, this function uses add_role
        :param role_id: The id of the role to add to (or create)
        :param user_id: The user_id to add
        """
        self.add_role(role_id, [user_id])

    def remove_user(self, role_id, user_id):
        """
        Removes a user_id from a role if the users exists
        :param role_id: The role to remove from
        :param user_id: The user_id to remove
        """
        if self.exists(role_id, user_id):
            self.roles[role_id].discard(user_id)

    def remove_role(self, role_id):
        if role_id in self.roles:
            del self.roles[role_id]

    def exists(self, role_id: str, user_id: str | None = None) -> bool:
        """
        Used to find if a role has been added, if
        :param role_id: The id of the role to check
        :param user_id: The optional user_id to be checked for
        :return: True if it exists otherwise False
        """
        if user_id is not None:
            if role_id not in self.roles:
                return False
            return user_id in self.roles[role_id]
        return role_id in self.roles

    def get_role(self, role_id: str) -> List[str] | None:
        """
        Returns a list of user_ids or None if role does not exist
        :param role_id: The role to retrieve
        :return: A list of user_ids or None if role doesn't exist
        """
        return list(self.roles[role_id]) if role_id in self.roles else None

    def __str__(self):
        roles_str: str = '{\n'
        for _key, _val in self.roles.items():
            roles_str += f'"{_key}": {_val},\n'
        return roles_str + '}'


def fetch_roles_users(roles_filter: Dict = None) -> Roles:
    """
    Uses the api to fetch all roles containing users,
    it creates a Roles object and adds the users to their relevant
    roles
    :param roles_filter: A dict that is applied as a filter,
    can be used as per api documentation, if None, all roles are fetched
    :return: A roles object
    """
    req_data = json.dumps(roles_filter)  # no custom encoder so far
    resp: requests.Response = requests.post(ENDPOINT, data=req_data)
    resp_list: List[Dict] = json.loads(resp.content)  # no custom hook yet
    roles = Roles()
    for user_roles in resp_list:
        roles.add_user(user_roles["role_id"], user_roles["user_id"])
    return roles


if __name__ == '__main__':
    ENDPOINT: str = 'http://localhost:8100/api/role/user/information'
    # TODO: @JonathanEnslin add proper tests
    # _roles = Roles()
    # print(_roles)
    # print()
    # _roles.add_role("123-123", [
    #     "111-111",
    #     "111-222",
    #     "222-111",
    #     "111-111"
    # ])
    # _roles.add_user("123-123", "111-111")
    # _roles.add_user("123-123", "111-333")
    # _roles.add_user("321-123", "111-111")
    # print(_roles)
    # print(_roles.exists("8787878", "4654"))
    # print(_roles.exists("8787878"))
    # print(_roles.exists("123-123"))
    # print(_roles.exists("123-123", "56454"))
    # print(_roles.exists("123-123", "111-111"))
    # _roles.remove_user("123-123", "111-333")
    # role_123_123 = _roles.get_role("123-123")
    # print(role_123_123)
    # role_123_123.append("789-987")
    # print(role_123_123)
    # print(_roles)
    _roles = fetch_roles_users()
    print(_roles)
