from flask import Flask
import smart_scheduler

app = Flask(__name__)

@app.route('/', methods=['POST'])
def hello():
    scheduler = smart_scheduler.SmartScheduler()
    scheduler.schedule()
    scheduler.createBookings()
    return scheduler.printBookings()
app.run(host="0.0.0.0", port=8081, debug=True)