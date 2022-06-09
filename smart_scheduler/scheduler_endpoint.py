from flask import Flask
import smart_scheduler

app = Flask(__name__)

@app.route('/', methods=['POST'])
def hello():
    print("what on erf") 
    scheduler = smart_scheduler.SmartScheduler()
    scheduler.schedule()
    return scheduler.printBookings()
    # return "Hello World!"
print("what2")  # please
app.run(host="0.0.0.0", port=8081, debug=True)