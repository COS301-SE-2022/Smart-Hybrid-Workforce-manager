import random
import datetime
from typing import Dict, List


class UserGenerator:
    def __init__(self, config: Dict) -> None:
        self.history: List[Dict] = []
        self.config = config["create_user"]
        self.office_days_options = self.config["office_days_options"]
        self.office_days_probabilities = self.config["office_days_probabilities"]
        self.time_bins: List[datetime.datetime] = []
        

    # returns true if the name was already used
    def search_history_names(self, first_name: str, last_name: str) -> bool:
        for user in self.history:
            if first_name == user["first_name"] and last_name == user["last_name"]:
                return True

    def generate(self, first_names: List[str], last_names: List[str], passwords: List[str], seed: int = None):
        if seed is not None:
            random.seed(seed)

        # noinspection PyDictCreation
        user = {
                "first_name": None,
                "last_name": None,
                "email": None,
                "password": None,
                "office_days": 0,
                "preferred_start_time": "2022-08-24T09:00:00.000Z",
                "preferred_end_time": "2022-08-24T16:00:00.000Z",
                "work_from_home": False
        }

        # names
        user["first_name"] = random.choice(first_names)
        user["last_name"] = random.choice(last_names)
        while self.search_history_names(user["first_name"], user["last_name"]):
            user["first_name"] = random.choice(first_names)
            user["last_name"] = random.choice(last_names)

        # office days
        user["office_days"] = random.choices(self.office_days_options, self.office_days_probabilities)

        # email
        user["email"] = user["first_name"] + user["last_name"] + str(user["office_days"])

        # work from home
        wfh_prob = self.config["work_from_home_probability"]  # work from home probability
        user["work_from_home"] = random.choices([False, True], weights=[1 - wfh_prob, wfh_prob])

        if self.config["password_override"] is not None:
            user["password"] = self.config["password_override"]
        else:
            user["password"] = random.choice(passwords)
