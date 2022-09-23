import random
import datetime
from typing import Dict, List


class UserGenerator:
    def __init__(self, config: Dict) -> None:
        self.history: List[Dict] = []
        self.config = config["create_user"]
        self.office_days_options = self.config["office_days_options"]
        self.office_days_probabilities = self.config["office_days_probabilities"]
        self.time_bins: List[datetime.datetime] = [self.config["lowest_preferred_start_time"]]
        self.email_domains = self.config["email_domains"]
        step_minutes: datetime.timedelta = datetime.timedelta(minutes=self.config["preferred_time_step_minutes"])
        time_slot: datetime.datetime = self.config["lowest_preferred_start_time"] + step_minutes
        while time_slot <= self.config["highest_preferred_end_time"]:
            self.time_bins.append(time_slot)
            time_slot += step_minutes

        self.team_num_bins = list(range(len(self.config["team_probabilities"])))
        self.team_num_probs = self.config["team_probabilities"]

        self.role_num_bins = list(range(len(self.config["role_probabilities"])))
        self.role_num_probs = self.config["role_probabilities"]

    # returns true if the name was already used
    def search_history_names(self, first_name: str, last_name: str) -> bool:
        for user in self.history:
            if first_name == user["first_name"] and last_name == user["last_name"]:
                return True

    def generate(self, first_names: List[str], last_names: List[str], *, passwords: List[str] = None,
                 teams: List[str] = None, roles: List[str] = None, desks: List[str] = None,
                 profile_pics: List[str] = None, seed: int = None) -> Dict:
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
            "work_from_home": False,
            "building_id": None,
            "preferred_desk": None,
            "picture": None,
            "_teams": [],
            "_roles": [],
        }

        # names
        user["first_name"] = random.choice(first_names)
        user["last_name"] = random.choice(last_names)
        while self.search_history_names(user["first_name"], user["last_name"]):
            user["first_name"] = random.choice(first_names)
            user["last_name"] = random.choice(last_names)

        # office days
        user["office_days"] = random.choices(self.office_days_options, self.office_days_probabilities, k=1)[0]

        # email
        user["email"] = user["first_name"].lower() + "." + user["last_name"].lower() + str(user["office_days"]) + random.choice(self.email_domains)

        # work from home
        wfh_prob = self.config["work_from_home_probability"]  # work from home probability
        user["work_from_home"] = random.choices([False, True], weights=[1 - wfh_prob, wfh_prob], k=1)[0]

        if self.config["password_override"] is not None:
            user["password"] = self.config["password_override"]
        else:
            user["password"] = random.choice(passwords)

        # office time
        start_time_i: int = random.randint(0, len(self.time_bins) - 2)
        end_time_i: int = random.randint(start_time_i + 1, len(self.time_bins) - 1)
        user["preferred_start_time"] = self.time_bins[start_time_i].replace(tzinfo=None).isoformat() + "Z"
        user["preferred_end_time"] = self.time_bins[end_time_i].replace(tzinfo=None).isoformat() + "Z"

        if profile_pics is not None and len(profile_pics) > 0:
            user["picture"] = random.choice(profile_pics)

        if desks is not None and len(desks) > 0:
            user["preferred_desk"] = random.choice(desks)

        if teams is not None:
            num_teams = random.choices(self.team_num_bins, self.team_num_probs, k=1)[0]
            num_teams = min(num_teams, len(teams))
            if num_teams > 0:
                user["_teams"] = random.sample(teams, num_teams)

        if roles is not None:
            num_roles = random.choices(self.role_num_bins, self.role_num_probs, k=1)[0]
            num_roles = min(num_roles, len(roles))
            if num_roles > 0:
                user["_roles"] = random.sample(roles, num_roles)

        return user
