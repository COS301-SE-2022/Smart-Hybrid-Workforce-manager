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
    def __init__(self):
        self.teamList = teams.fetch_team_users()
        self.resourceList = resource.fetch_resources()
        self.roomList = self.loadResources()
        self.userList = user.fetch_users()
        self.bookingsList = []

    # only loads desks at the moment
    def loadResources(self):
        _rooms: room.Rooms = room.Rooms()
        for resource in self.resourceList:
            if resource.resource['resource_type'] == 'DESK':
                _rooms.add_resource(resource)
        return _rooms

    # returns a list of team IDs sorted in descending order of the number of members in the team
    def getSortedTeams(self):
        return sorted(self.teamList.teams, key=lambda t: self.teamList.team_size(t), reverse=True)

    def schedule(self):
        sortedTeams = self.getSortedTeams() # gets team IDs sorted in descending order of the number of members in the team
        for team in sortedTeams:
            while self.teamList.teams[team]:
                teamSize = self.teamList.team_size(team)
                rooms = rooms = self.roomList.rooms_size(teamSize, 'ge') # gets rooms that can fir the entire team

                while not rooms: # if no room is big enough, find the largest room
                    teamSize -= 1
                    rooms = self.roomList.rooms_size(teamSize, 'eq')

                for i in range(teamSize):
                    userID = self.teamList.teams[team].pop()
                    startTime = list(filter(lambda u: u.user['id'] == userID, self.userList))[0].user['preferred_start_time'].strftime('T%H:%M:%SZ')
                    endTime = list(filter(lambda u: u.user['id'] == userID, self.userList))[0].user['preferred_end_time'].strftime('T%H:%M:%SZ')
                    bookDate = date.today() + timedelta(days=7)

                    # create desk booking
                    self.bookingsList.append({"user_id": userID,
                    "resource_type": "DESK",
                    "resource_preference_id": self.roomList.rooms[rooms[0]].pop(),
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

    def get_fitness(self):
        return fitness.get_fitness(self.bookingsList)

# scheduler = SmartScheduler()
# scheduler.schedule()
# scheduler.printBookings()
# scheduler.printBookings()
