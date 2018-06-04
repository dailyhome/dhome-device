package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	gatewayUrl string
	deviceId   string
	skills     []string
	token      string
)

type RegistationInfo struct {
	Method   string   `json: method`
	DeviceID string   `json: deviceid`
	Skills   []string `json: skills`
}

type HealthStat struct {
	Status      bool            `json: status`
	SkillStatus map[string]bool `json: skills`
}

func validSkill(skill string) bool {
	for _, name := range skills {
		if name == skill {
			return true
		}
	}
	return false
}

// Get device info with health data
func Device(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

// Get device Health info
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	health := HealthStat{}
	health.SkillStatus = make(map[string]bool)

	//  TODO: check skill health
	for _, name := range skills {
		health.SkillStatus[name] = true
	}

	health.Status = true
	data, _ := json.Marshal(health)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Route a request to a skill and return the response
func Skill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	method := ps.ByName("method")
	skill := ps.ByName("name")
	if !validSkill(skill) {
		http.Error(w, fmt.Sprintf("Invalid skill request"), http.StatusInternalServerError)
		return
	}
	value := ps.ByName("value")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+skill+":8080/"+method+"?value="+value, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", token)
	resp, reqerr := client.Do(req)
	if reqerr != nil {
		http.Error(w, fmt.Sprintf("failed to perform skill req, error: %v", reqerr), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, readerr := ioutil.ReadAll(resp.Body)
	if readerr != nil {
		http.Error(w, fmt.Sprintf("failed to perform skill req, error: %v", readerr), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func register() error {
	gatewayUrl = os.Getenv("DAILYIOT")
	if gatewayUrl == "" {
		return fmt.Errorf("DAILYIOT url can't be empty")
	}
	deviceId = os.Getenv("DEVICEID")
	if gatewayUrl == "" {
		return fmt.Errorf("DEVICEID can't be empty")
	}
	token = os.Getenv("DIOTTOKEN")
	if gatewayUrl == "" {
		return fmt.Errorf("DIOTTOKEN can't be empty")
	}
	skillsstr := os.Getenv("SKILLS")
	if skillsstr != "" {
		jsonerr := json.Unmarshal([]byte(skillsstr), &skills)
		if jsonerr != nil {
			return fmt.Errorf("failed to unmarshal device skills, error: %v", jsonerr)
		}
	}
	info := RegistationInfo{DeviceID: deviceId, Skills: skills, Method: "register"}
	js, jserr := json.Marshal(info)
	if jserr != nil {
		return fmt.Errorf("failed to marshal registration, error: %v", jserr)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", gatewayUrl, bytes.NewBuffer([]byte(js)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", token)
	resp, reqerr := client.Do(req)
	if reqerr != nil {
		return fmt.Errorf("failed to perform request, error: %v", reqerr)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to perform request, statuscode: %s", resp.StatusCode)
	}
	return nil
}

func main() {
	router := httprouter.New()
	router.GET("/", Device)
	router.GET("/health", Health)
	router.POST("/skill/:name/:method/:value", Skill)

	err := register()
	if err != nil {
		log.Fatal("Failed to register device with diot platform, error: ", err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
