import unittest
from datetime import datetime

import booking


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
        result = booking.parse_dict(booking1)
        self.assertEqual(result, booking_parsed)


if __name__ == '__main__':
    unittest.main()
