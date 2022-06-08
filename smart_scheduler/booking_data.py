import json
import requests
from typing import List, Dict, Tuple


ENDPOINT: str = 'http://arche-api:8080/api/booking/information'


def get_bookings() -> List[Dict]:
    # get bookings for which booked = true
    resp: requests.Response = requests.post(ENDPOINT, data='{"booked": true}')

    # return booking_list
