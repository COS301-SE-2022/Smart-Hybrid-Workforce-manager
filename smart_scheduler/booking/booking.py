from typing import Dict
from datetime import datetime

# change in future to be a global constant
TIME_FMT: str = '%Y-%m-%dT%H:%M:%S.%f%z'  # %z formats Z as being UTC+0


class Booking:
    start_key: str = 'start'
    end_key: str = 'end'
    date_created_key = 'date_created'

    def __init__(self, booking_dict: Dict[str]):
        self.booking = booking_dict


def parse_dict(booking: Dict[str, str | bool | None]) -> Dict[str, str | bool | None | datetime]:
    """
    Parses a dict representing a booking, specifically, it parses the dates, and converts them to
    datetime objects instead of strings
    :param booking: The booking dict
    :return: A new booking dict that have parsed dates
    """
    parsed_booking: Dict[str, str | bool | None | datetime] = dict(booking)  # copy booking
    if Booking.start_key in parsed_booking and parsed_booking[Booking.start_key] is not None:
        parsed_booking[Booking.start_key] = datetime.strptime(parsed_booking[Booking.start_key], TIME_FMT)

    if Booking.end_key in parsed_booking and parsed_booking[Booking.end_key] is not None:
        parsed_booking[Booking.end_key] = datetime.strptime(parsed_booking[Booking.end_key], TIME_FMT)

    if Booking.date_created_key in parsed_booking \
            and parsed_booking[Booking.date_created_key] is not None:
        parsed_booking[Booking.date_created_key] = datetime.strptime(
            parsed_booking[Booking.date_created_key], TIME_FMT
        )
    return parsed_booking
