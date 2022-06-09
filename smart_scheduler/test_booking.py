import unittest
import json
from datetime import datetime

import booking as _booking


class Test(unittest.TestCase):
    def test_parse_dict(self):
        booking1 = {
            "id": "33333333-1111-4a06-9983-8b374586e459",
            "user_id": "11111111-1111-4a06-9983-8b374586e459",
            "resource_type": "DESK",
            "resource_preference_id": "22222222-dc08-4a06-9983-8b374586e459",
            "start": "2022-05-09T09:54:16.123456Z",
            "end": "2022-05-09T13:54:16.123456Z",
            "booked": False,
            "date_created": "2022-06-08T11:07:56.123456Z"
        }
        booking_parsed = {
            "id": "33333333-1111-4a06-9983-8b374586e459",
            "user_id": "11111111-1111-4a06-9983-8b374586e459",
            "resource_type": "DESK",
            "resource_preference_id": "22222222-dc08-4a06-9983-8b374586e459",
            "start": datetime(2022, 5, 9, 9, 54, 16, 123456),
            "end": datetime(2022, 5, 9, 13, 54, 16, 123456),
            "booked": False,
            "date_created": datetime(2022, 6, 8, 11, 7, 56, 123456)
        }
        result = _booking.parse_dict(booking1)
        self.assertEqual(result.booking, booking_parsed)

    def test_booking_encoder(self):
        year, month, day, hour, minutes, seconds, mu_seconds = 2022, 12, 11, 10, 23, 34, 123456
        id_str = "123-123"
        booking = _booking.Booking()
        booking.booking["id"] = id_str  # no need to test with real id
        booking.booking["start"] = datetime(year, month, day, hour, minutes, seconds, mu_seconds)
        json_str = json.dumps(booking, cls=_booking.BookingEncoder)
        expect_date = f'{year}-{month}-{day}T{hour}:{minutes}:{seconds}.{123456}Z'
        self.assertEqual(json_str, f'{{"id": "{id_str}", "user_id": null, "resource_type": null, '
                                   f'"resource_preference_id": null, "start": "{expect_date}", '
                                   f'"end": null, "booked": null, "date_created": null}}')


if __name__ == '__main__':
    unittest.main()
