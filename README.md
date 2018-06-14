# diot_device
bootstraps your dailyiot device stack

### Overview of diot_device
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
   
   
### Getting Started

> For getting started with swarm on Raspberry Pi with raspbian follow the [instructions](https://github.com/dailyiot/diot_device/blob/master/doc/setup_raspberry.md)     
> Alternatively,    
> You can deploy the dummy implementation locally by running the [dummy skills](https://github.com/dailyiot/skills/blob/master/README.md#dummy)
   
Define the device unique identifier by editing `configuration` file 
```bash
DEVICEID=MyRasp
```

Build Locally (optional)
```bash
./build.sh
```

Deploy the device stack
```bash
./deploy.sh
```

### Configuration
> Device stack can be deployed along with daily-iot platform in same swarm cluster otherwise independently   
> It uses same network as of openfaas functions `func_functions` for making `device-gateway` accessible from platform    
> and create a private network based on the `DEVICEID` to independently manage skills

#### Run on independent swarm cluster

For a independent swarm cluster daily-iot platform address and device address need to be defined by changing `docker-compose.yml` file:    
Change the dailyiot gateway address  
```yaml
DAILYIOT: "http://your-openfaas/function/diot-gateway"
```
and    
Change the public address of your device  
```yaml
DEVICEADDR: "http://<device_public_address>"
```
    
#### Run multiple device stack in same host

Multiple device stack can be deployed in a same host by changing the `DEVICEID` and exposed `PORT` in the `configuration` file
```bash
DEVICEID=MySecondDevice
PORT=6207
```
    
#### TODO
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
