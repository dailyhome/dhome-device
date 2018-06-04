from flask import Flask
from flask import request
from raspsys import relay_ctl 
import sys
import signal
import json


validSocket = ["1", "2", "3", "4", "5", "6", "7", "8", "all"]
app = Flask(__name__)

@app.route("/")
def root():
    return json.dumps({"enable" : "/enable?socket=<no/all>", 
                       "disable": "/disable?socket=<no/all>", 
                       "state": "/state [?socket=<no/all>]",
                       "reset": "/reset"})

@app.route('/enable')
def enable():
    socket = request.args.get('value')
    if socket == 'all':
        for socket in range(1,9):
            relay_ctl.turnon(str(socket))
        state = relay_ctl.state()
        return json.dumps(state)
    elif socket in validSocket[:-1]:
        relay_ctl.turnon(socket)
        state = relay_ctl.state()
        return json.dumps(state)
    else:
        print "Invalid Request: invalid socket, valid sockets: " + ' '.join(validSocket)
        return json.dumps({ "error": "invalid socket" })

@app.route('/disable')
def disable():
    socket = request.args.get('value')
    if socket == 'all':
        for socket in range(1,9):
            relay_ctl.turnoff(str(socket))
        state = relay_ctl.state()
        return json.dumps(state)
    elif socket in validSocket[:-1]:
        relay_ctl.turnoff(socket)
        state = relay_ctl.state()
        return json.dumps(state)
    else:
        print "Invalid Request: invalid socket, valid sockets: " + ' '.join(validSocket)
        return json.dumps({ "error": "invalid socket" })

@app.route('/state')
def state():
    socket = request.args.get('value')
    state = relay_ctl.state()
    if socket == 'all' or socket == None:
        return json.dumps(state)
    elif socket in validSocket[:-1]:
        return json.dumps({ socket: state[socket] })
    else:
        print "Invalid Request: invalid socket, valid sockets: " + ' '.join(validSocket)
        return json.dumps({ "error": "invalid socket" })

@app.route('/reset')
def reset():
    relay_ctl.reset()
    state = relay_ctl.state()
    return json.dumps(state)

def handler(signal, frame):
    relay_ctl.reset()
    sys.exit(0)

if __name__ == '__main__':
    signal.signal(signal.SIGINT, handler)
    app.run(host='0.0.0.0', port=8080, debug=False)
