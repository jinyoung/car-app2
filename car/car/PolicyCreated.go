package car

import (
	"time"
)

type PolicyCreated struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	PolicyId string `json:"policyId" type:"string"` 
	CarId string `json:"carId" type:"string"` 
	
}

func NewPolicyCreated() *PolicyCreated{
	event := &PolicyCreated{EventType:"PolicyCreated", TimeStamp:time.Now().String()}

	return event
}
