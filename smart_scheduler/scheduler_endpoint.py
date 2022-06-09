from pickle import TRUE
from re import T
from flask import Flask

app = Flask(__name__)

@app.route('/', methods=['POST'])
def hello():
    return "Hello World!"
print("what2")
app.run(host="0.0.0.0", port=8081, debug=True)