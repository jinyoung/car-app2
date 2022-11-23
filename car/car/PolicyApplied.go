package car

import (
	"time"
)

type PolicyApplied struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	CarId int `json:"carId" type:"int"` 
	PolicyApplicationId int `json:"policyApplicationId" type:"int"` 
	Status string `json:"status" type:"string"` 
	
}

func NewPolicyApplied() *PolicyApplied{
	event := &PolicyApplied{EventType:"PolicyApplied", TimeStamp:time.Now().String()}

	return event
}
