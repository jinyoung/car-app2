package car

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverPolicyCreated_AddPolicy(data map[string]interface{}){
	
	event := NewPolicyCreated()
	mapstructure.Decode(data,&event)

	AddPolicy(event);
}

