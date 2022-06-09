import json
from typing import Dict, List, Set

import requests

"""
This module is responsible for team related operations
"""

ENDPOINT: str = 'http://arche-api:8080/api/team/user/information'


class Teams:
    """
    Class representing teams, an object of this class
    will contain a Dict that uses team_ids as keys, the elements
    of this dict will be lists containing user_ids, a more efficient
    solution can still be investigated, a separate dict will also be
    maintained containing the date created for each team
    """

    def __init__(self, teams: Dict[str, Set] = None, teams_meta: Dict[str, List] = None):
        if teams is None:
            self.teams: Dict[str, Set[str]] = {}  # set is used to ensure users aren't added twice
        if teams_meta is None:  # not yet used TODO: @JonathanEnslin see if needed
            self.teams_meta: Dict = {}

    def add_team(self, team_id: str, user_ids: List[str] = None):
        """
        Adds a team or adds a list of member to a team
        :param team_id: The ID of the team to be added or that members should be added to
        :param user_ids:
        """
        if team_id not in self.teams:
            # add team
            self.teams[team_id] = set()  # create empty set
        # add users that are not already contained
        for _id in user_ids:
            self.teams[team_id].add(_id)

    def add_member(self, team_id: str, user_id: str):
        """
        Adds a member to a team, team is created if it does not
        yet exist, this function uses add_team
        :param team_id: The id of the team to add to (or create)
        :param user_id: The user_id to add
        """
        self.add_team(team_id, [user_id])

    def exists(self, team_id: str) -> bool:
        """
        Used to find if a team has been added
        :param team_id: The if of the team to check
        :return: True if it exists otherwise False
        """
        return team_id in self.teams

    def team_size(self, team_id: str) -> int:
        """
        Gets the size (amount of members in team), returns -1
        if the team is not present
        :param team_id: The id of the team to get the size of
        :return: The size of the team
        """
        return len(self.teams[team_id]) if team_id in self.teams else -1

    def __str__(self):
        teams_str: str = '{\n'
        for _key, _val in self.teams.items():
            teams_str += f'"{_key}": {_val},\n'
        return teams_str + '}'


def fetch_team_users(teams_filter: Dict = None) -> Teams:
    """
    Uses the api to fetch all teams containing users,
    it creates a Teams object and adds the users to their relevant
    teams
    :param teams_filter: A dict that is applied as a filter,
    can be used as per api documentation, if None, all teams are fetched
    :return: A teams object
    """
    req_data = json.dumps(teams_filter)  # no custom encoder so far
    resp: requests.Response = requests.post(ENDPOINT, data=req_data)
    resp_list: List[Dict] = json.loads(resp.content) # no custom hook yet
    teams = Teams()
    for user_team in resp_list:
        teams.add_member(user_team["team_id"], user_team["user_id"])
    return teams


if __name__ == '__main__':
    ENDPOINT: str = 'http://localhost:8100/api/team/user/information'
    # TODO: @JonathanEnslin add proper tests
    # _teams = Teams()
    # print(_teams)
    # print()
    # _teams.add_team("123-123", [
    #     "111-111",
    #     "111-222",
    #     "222-111",
    #     "111-111"
    # ])
    # _teams.add_member("123-123", "111-111")
    # _teams.add_member("123-123", "111-333")
    # _teams.add_member("321-123", "111-111")
    # print(_teams)
    # print(_teams.team_size("123-123"))
    # print(_teams.exists("8787878"))
    _teams = fetch_team_users()
    print(_teams)
