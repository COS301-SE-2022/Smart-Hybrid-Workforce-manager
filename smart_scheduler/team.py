from typing import Dict, List, Set

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
    def __init__(self, teams: Dict[str, List] = None, teams_meta: Dict[str, List] = None):
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

    def team_size(self, team_id: str) -> int:
        """
        Gets the size (amount of members in team), returns -1
        if the team is not present
        :param team_id: The id of the team to get the size of
        :return: The size of the team
        """
        return len(self.teams[team_id]) if team_id in self.teams else -1
    

if __name__ == '__main__':
    aset = set()
    aset.add("nope")
    print(aset)
    aset.add("nope")
    aset.add("nope2")
    print(aset)
