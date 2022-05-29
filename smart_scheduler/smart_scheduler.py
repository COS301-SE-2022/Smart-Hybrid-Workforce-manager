import docker
import requests
import json 

def init():
    # response = requests.post("http://arche-api:8080/api/booking/information", data="{}")
    response = requests.post("http://arche-api:8080/api/booking/information", data='{}')
    global bookingsList
    bookingsList = json.loads(response.content)

def printBookings():
    response = requests.post("http://arche-api:8080/api/booking/information", data='{}')
    bookingsList = json.loads(response.content)
    i = 1
    for booking in bookingsList:
        print('Booking ' + str(i) + ': ')
        i+=1
        for field in booking:
            print(field + ' : ' + str(booking[field]))
        print()

# only creates the first booking for now
def finaliseBookings():
    for booking in bookingsList:
        if not booking['booked']:
            response = requests.post("http://arche-api:8080/api/booking/create", data='{"id": "' + str(booking['id']) + '", "user_id": "' + str(booking['user_id']) + '", "resource_type": "' + str(booking['resource_type']) + '", "resource_preference_id": "' + str(booking['resource_preference_id']) + '", "start": "' + str(booking['start']) + '", "end": "' + str(booking['end']) + '", "booked": true}')
            print(response)
            break

init()
# finaliseBookings()
# printBookings()
