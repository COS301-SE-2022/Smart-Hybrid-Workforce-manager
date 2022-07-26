import logging

from flask import Flask
from flask import request
import smart_scheduler

app = Flask(__name__)

app.logger.setLevel(logging.DEBUG)


@app.route('/', methods=['POST'])
def scheduler_call():
    scheduler_data = request.json  # This is holds the data send by the scheduler
    scheduler = smart_scheduler.SmartScheduler()
    scheduler.schedule()
    scheduler.createBookings()
    return scheduler.printBookings()


@app.route('/echo', methods=['POST'])  # This endpoint is used for testing purposes
def echo_call():
    data = request.json
    app.logger.debug(data)  # Print received data
    return data


app.run(host="0.0.0.0", port=8081, debug=True)
