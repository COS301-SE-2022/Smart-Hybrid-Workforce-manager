from audioop import reverse
from os import sched_getscheduler
import sched
from tracemalloc import start
from xml.etree.ElementTree import QName
import docker
import requests
import json 
from datetime import datetime

import booking
import user
import role
import resource
import teams
import room

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
                    # create desk booking
                    self.bookingsList.append({"user_id": userID,
                    "resource_type": "DESK",
                    "resource_preference_id": self.roomList.rooms[rooms[0]].pop(),
                    "start": (datetime.today().strftime('%Y-%m-%d') + startTime),
                    "end": (datetime.today().strftime('%Y-%m-%d') + endTime),
                    "booked": False})
                    # create parking booking
                    self.bookingsList.append({"user_id": userID,
                    "resource_type": "PARKING",
                    "resource_preference_id": None,
                    "start": startTime,
                    "end": endTime,
                    "booked": False})


    def createBookings(self, verbose: bool = False):
        batchBooking = {"user_id" : "00000000-0000-0000-0000-000000000000", "bookings" :  self.bookingsList }
        response = requests.post("http://arche-api:8080/api/batch-booking/create", data=json.dumps(batchBooking))
        if verbose:
            print(response)

    def printBookings(self):
        print("BOOKINGS")
        for booking in self.bookingsList:
            print(booking)



scheduler = SmartScheduler()
# print(scheduler.roomList)
# print(scheduler.resourceList)
# scheduler.teamList.add_member("1", "1")
# scheduler.teamList.add_member("1", "2")
# scheduler.teamList.add_member("1", "3")
# scheduler.teamList.add_member("2", "1")
# scheduler.teamList.add_member("2", "2")

scheduler.schedule()
scheduler.printBookings()


# print(scheduler.teamList.teams)
# scheduler.sortTeams()
# print()
# print(scheduler.teamList)

# for team in scheduler.teamList.teams:
#     print(team, "   ", scheduler.teamList.get_team(team))

# temp = scheduler.getSortedTeams()
# print()
# for team in temp:
#     print(team, "   ", scheduler.teamList.get_team(team))

# scheduler.teamList.remove_member('12121212-dc08-4a06-9983-8b374586e459', '11111111-1111-4a06-9983-8b374586e459')

# temp = scheduler.getSortedTeams()
# print()
# for team in temp:
#     print(team, "   ", scheduler.teamList.get_team(team))

# def get_teams():
#     global teamList
#     teamList = teams.fetch_team_users()

# def get_rooms():
#     global roomList
#     roomList = teams.fetch_team_users()

# print("USERS")
# _users = user.fetch_users()
# for r in _users:
#     # print(r.user['preferred_start_time'])
#     print(r)
# print()

# print("RESOURCES")
# _resources = resource.fetch_resources()
# for r in _resources:
#     print(r)
#     for key, val in r.resource.items():
#         print(key, "   ", val, type(val))
# print()

# print("BOOKINGS")
# _bookings = booking.fetch_bookings(None)
# for b in _bookings:
#     print(b)
# print()

# print("TEAMS")
# _teams = teams.fetch_team_users()
# print(_teams)
# print()

# print("ROOMS")
# _rooms: room.Rooms = room.Rooms()
# _rooms.add_resource(resource.Resource({"id": "123-123", "room_id": "456-456"}))
# _rooms.add_resource(resource.Resource({"id": "321-321", "room_id": "456-456"}))
# _rooms.add_resource(resource.Resource({"id": "123-321", "room_id": "456-456"}))
# _rooms.add_resource(resource.Resource({"id": "321-123", "room_id": "456-456"}))
# _rooms.add_resource(resource.Resource({"id": "321-123", "room_id": "456-456"}))

# _rooms.add_resource(resource.Resource({"id": "789-987", "room_id": "753-357"}))
# _rooms.add_resource(resource.Resource({"id": "951-159", "room_id": "753-357"}))
# # print(_rooms)
# # print(''.join(_rooms.rooms_size(2, 'gt')))
# print(', '.join(_rooms.rooms_size(2, 'ge')))
# _rooms.remove_resource(resource.Resource({"id": "789-987", "room_id": "753-357"}))
# print(', '.join(_rooms.rooms_size(2, 'ge')))
# print('asfd', ', '.join(_rooms.rooms_size(3, 'ge')))
# print()