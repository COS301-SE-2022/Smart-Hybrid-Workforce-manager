from audioop import reverse
from xml.etree.ElementTree import QName
import docker
import requests
import json 
from datetime import date, timedelta

import booking
import user
import role
import resource
import teams
import room
import fitness

class SmartScheduler:
    def __init__(self, scheduler_data):
        self.teamList = scheduler_data["teams"]
        self.resourceList = scheduler_data["resources"]
        self.roomList = scheduler_data["rooms"]
        self.userList = scheduler_data["users"]
        self.bookingsList = []
        # self.teamList = teams.fetch_team_users()
        # self.resourceList = resource.fetch_resources()
        # self.roomList = self.loadResources()
        # self.userList = user.fetch_users()
        # self.bookingsList = []

    # only loads desks at the moment
    def loadResources(self):
        _rooms: room.Rooms = room.Rooms()
        for resource in self.resourceList:
            if resource.resource['resource_type'] == 'DESK':
                _rooms.add_resource(resource)
        return _rooms

    # returns a list of team IDs sorted in descending order of team priority
    def getSortedTeamsPriority(self):
        return sorted(self.teamList, key=lambda t: t['priority'], reverse=True)

    # returns a list of team IDs sorted in descending order of the number of members in the team
    def getSortedTeamsSize(self):
        return sorted(self.teamList, key=lambda t: len(t['user_ids']), reverse=True)

    def schedule(self):
        sortedTeams = self.getSortedTeamsSize() # gets team IDs sorted in descending order of the number of members in the team
        for team in sortedTeams:
            while team['user_ids']:
                # teamSize = len([t for t in self.teamList if t['id'] == team['id']]['user_ids'])
                teamSize = len(team['user_ids'])
                room_id = room.room_size(self.roomList, teamSize, 'ge') # gets rooms that can fit the entire team

                while not room_id: # if no room is big enough, find the largest room
                    teamSize -= 1
                    room_id = room.room_size(self.roomList, teamSize, 'eq')

                room_id = room_id[0]
                assigned_room_resources = []
                for r in self.roomList:
                    if r['id'] == room_id:
                        assigned_room_resources = r['resource_ids']

                for i in range(teamSize):
                    userID = team['user_ids'].pop()
                    current_user = list(filter(lambda u: u['id'] == userID, self.userList))[0]
                    # assigned_desk_id = assigned_room_resources[0] if current_user.get('preferred_desk') is None else current_user.get('preferred_desk')
                    assigned_desk_id = assigned_room_resources.pop()
                    # startTime = current_user.get('preferred_start_time').strftime('T%H:%M:%SZ')
                    # endTime = current_user.get('preferred_end_time').strftime('T%H:%M:%SZ')
                    startTime = current_user.get('preferred_start_time')
                    endTime = current_user.get('preferred_end_time')
                    bookDate = date.today() + timedelta(days=7)

                    # create desk booking
                    self.bookingsList.append({"user_id": userID,
                    "resource_type": "DESK",
                    "resource_preference_id": assigned_desk_id,
                    "start": (bookDate.strftime('%Y-%m-%d') + startTime),
                    "end": (bookDate.strftime('%Y-%m-%d') + endTime),
                    "booked": False})

                    # create parking booking
                    self.bookingsList.append({"user_id": userID,
                    "resource_type": "PARKING",
                    "resource_preference_id": None,
                    "start": (bookDate.strftime('%Y-%m-%d') + startTime),
                    "end": (bookDate.strftime('%Y-%m-%d') + endTime),
                    "booked": False})

    def createBookings(self, verbose: bool = False):
        batchBooking = {"user_id" : "00000000-0000-0000-0000-000000000000", "bookings" :  self.bookingsList }
        response = requests.post("http://arche-api:8080/api/batch-booking/create", data=json.dumps(batchBooking))
        if verbose:
            print(response)

    def printBookings(self):
        ret_str = "BOOKINGS\n"
        print("BOOKINGS")
        for booking in self.bookingsList:
            ret_str += str(booking) + "\n"
            print(booking)
        print(self.bookingsList)
        return ret_str

    def get_bookings(self):
        ret_dict = {}
        ret_dict['bookings'] = self.bookingsList
        return ret_dict

    def get_fitness(self):
        return fitness.get_fitness(self.bookingsList)

# scheduler = SmartScheduler()
# scheduler.schedule()
# scheduler.printBookings()
# scheduler.printBookings()
