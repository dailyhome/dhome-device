# diot_device
bootstraps your dailyiot device 

#### Getting Started
Define the device unique identifier by editing `deviceid` file 
```
DEVICEID=MyRasp
```
Deploy the device stack
```
./deploy.sh
```

#### Overview of diot_device
1. Register device to the dailyiot platform
2. Provide gateway to route request to skills and provide health data
3. Deploy and manages skill on the go

> DailyIOT Device Stack 

    Device platform use swarm for its core
    FaasSwarm: Manage skill dynamically
    DeviceGateway: Lightweight and fast http router
    Skills : Handle Skill Specific Call (run in a priiviledged mode)

<p align="center">
   <img src="https://farm2.staticflickr.com/1756/40740438330_b4efa720db_o.jpg">
</p>


##### TODO
- [X] Implement switch skill
- [X] Implement dummy skill
- [X] Automatic device registration on Startup
- [X] Support local deployment of multiple devices for testing
- [ ] Optional Automatic ngrok tunnel creation on Startup
- [ ] Health Checkup for Skills
- [ ] Change password and token to docker secrets
- [ ] Template and SDK for Writing different IOT Skill 
- [ ] Write Skill implementation Documentation
- [ ] Travis integration
- [ ] Device Setup via Wifi Hotspot like Alexa
