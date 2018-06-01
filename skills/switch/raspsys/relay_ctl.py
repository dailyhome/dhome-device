#!/usr/bin/python
import RPi.GPIO as GPIO
from threading import Lock
import copy

pinlock = Lock()

# Each socket is represent by each pin
pinList = [2, 3, 4, 17, 27, 22, 10, 9]

socketState = {
    "1" : False,
    "2" : False,
    "3" : False,
    "4" : False,
    "5" : False,
    "6" : False,
    "7" : False,
    "8" : False,
}

def init():
    GPIO.setmode(GPIO.BCM)
    # loop through pins and set mode and state to 'low'
    for i in pinList: 
        GPIO.setup(i, GPIO.OUT) 
        GPIO.output(i, GPIO.HIGH)

def reset():
    pinlock.acquire()
    try:
        GPIO.cleanup()
        for key, value in socketState.iteritems():
            socketState[key] = False
    finally:
        pinlock.release()

def turnon(socket):
    pinlock.acquire()
    try:
        socketState[socket] = True
        GPIO.output(pinList[int(socket)-1], GPIO.LOW)
    finally:
        pinlock.release()

def turnoff(socket):
    pinlock.acquire()
    try:
        socketState[socket] = False
        GPIO.output(pinList[int(socket)-1], GPIO.HIGH)
    finally:
        pinlock.release()

def state():
    return copy.deepcopy(socketState)
