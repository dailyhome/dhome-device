from flask import Flask
from flask import request
from raspsys import relay_ctl
import sys
import signal
import json


app = Flask(__name__)
# dummy state
state = {
"1" : true,
"2" : false,
"3" : true,
"4" : false,
"5" : false,
"6" : true,
"7" : true,
"8" : false
}

@app.route("/")
def root():
    return json.dumps({"state": state})


def handler(signal, frame):
    relay_ctl.reset()
    sys.exit(0)

if __name__ == '__main__':
    signal.signal(signal.SIGINT, handler)
    app.run(host='0.0.0.0', port=8080, debug=False)
