from xml.etree.ElementTree import QName
import docker
import requests
import json 

def init():
    response = requests.post("http://arche-api:8080/api/booking/information", data='{"booked":false}')
    global bookingsList
    bookingsList = json.loads(response.content)
    bookingsList = [booking for booking in bookingsList if not booking['booked']]

# only used for testing purposes
def printBookings():
    i = 1
    for booking in bookingsList:
        print('Booking ' + str(i) + ': ')
        i+=1
        for field in booking:
            print(field + ' : ' + str(booking[field]))
        print()

def setAllBookingsTrue():
    for booking in bookingsList:
        booking['booked'] = True

# books all unfinalised bookings, no logic is performed yet.
def finaliseBookings():
    setAllBookingsTrue()
    batchBooking = {"user_id"  : "00000000-0000-0000-0000-000000000000", 
                    "bookings" :  bookingsList }
    response = requests.post("http://arche-api:8080/api/batch-booking/create", data=json.dumps(batchBooking))
    print(response)

init()
# printBookings()
# finaliseBookings()
