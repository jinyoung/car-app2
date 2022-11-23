package car-app

import (
	"time"
)

type PolicyDenied struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"`
	CarId int `json:"carId" type:"int"`
	PolicyApplicationId int `json:"policyApplicationId" type:"int"`
	Status string `json:"status" type:"string"`
}

func NewPolicyDenied() *PolicyDenied{
	event := &PolicyDenied{EventType:"PolicyDenied", TimeStamp:time.Now().String()}
	return event
}
