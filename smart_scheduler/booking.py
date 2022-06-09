"""
This module contains the code relating to Bookings
"""
import json
from typing import Dict, List
from datetime import datetime

import requests

ENDPOINT: str = 'http://arche-api:8080/api/booking/information'
ENDPOINT_BATCH: str = 'http://arche-api:8080/api/batch-booking/information'

# change in future to be a global constant
TIME_FMT: str = '%Y-%m-%dT%H:%M:%S.%fZ'
# TIME_FMT: str = '%Y-%m-%dT%H:%M:%S.%f%z'  # %z formats Z as being UTC+0


class Booking:
    """
    Represents a booking, encapsulates a dict
    """
    def __init__(self, booking_dict: Dict = None):
        self.booking = {
            "id": None,
            "user_id": None,
            "resource_type": None,
            "resource_preference_id": None,
            "start": None,
            "end": None,
            "booked": None,
            "date_created": None,
        }
        if booking_dict is not None:
            for key, val in booking_dict.items():
                self.booking[key] = val

    def __str__(self):
        return str(self.booking)


class BookingEncoder(json.JSONEncoder):
    """
    Class to be used as an encoder when encoding Booking class objects
    """
    def default(self, o):
        # if passed in object is datetime object
        if isinstance(o, datetime):
            return datetime.strftime(o, TIME_FMT)

        # if passed in object is a Booking
        if isinstance(o, Booking):
            return o.booking

        # otherwise use the default behavior
        return json.JSONEncoder.default(self, o)


def parse_dict(booking: Dict[str, str | bool | None]) -> Booking:
    """
    Parses a dict representing a booking, specifically, it parses the dates,
    and converts them to datetime objects instead of strings
    :param booking: The booking dict
    :return: A new booking dict that have parsed dates
    """
    parsed_booking: Dict[str, any] = dict(booking)  # copy booking
    parsed_booking["start"] = datetime.strptime(parsed_booking["start"], TIME_FMT)

    parsed_booking["end"] = datetime.strptime(parsed_booking["end"], TIME_FMT)

    if "date_created" in parsed_booking and parsed_booking["date_created"] is not None:
        parsed_booking["date_created"] = datetime.strptime(
            parsed_booking["date_created"], TIME_FMT
        )
    return Booking(parsed_booking)


# gets bookings that match at least one of the passed in filters
def fetch_bookings(filters: List[Booking] = None) -> List[Booking]:
    """
    This method calls the api and fetches bookings, it fetches
    all bookings matching at least one filter, as specified
    per the API documentation
    :param filters: The filters used when fetching bookings`
    :return: The bookings that have been fetched
    """
    req_dict: Dict = {
        "user_id": "00000000-0000-0000-0000-000000000000",
        "bookings": filters if filters is not None else [{}]
    }

    data = json.dumps(req_dict, cls=BookingEncoder)
    resp: requests.Response = requests.post(ENDPOINT_BATCH, data=data)
    resp_list: List[Dict] = json.loads(resp.content)
    bookings = [parse_dict(d) for d in resp_list]  # parse all the returned dictionaries
    return bookings


if __name__ == '__main__':
    ENDPOINT_BATCH: str = 'http://localhost:8100/api/batch-booking/information'
    # booking_filters = [Booking({"booked": False})]
    # _bookings = fetch_bookings(booking_filters)
    _bookings = fetch_bookings(None)
    for b in _bookings:
        print(b)
