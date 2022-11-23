package car

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"car/external"
)

type PolicyHistory struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	CarId int `json:"carId"`
	PolicyApplicationId int `json:"policyApplicationId"`
	Status string `json:"status"`

}

func (self *PolicyHistory) onPostPersist() (err error){
	policyApplied := NewPolicyApplied()
	model.Copy(policyApplied, self)

	Publish(policyApplied)
	policyDenied := NewPolicyDenied()
	model.Copy(policyDenied, self)

	Publish(policyDenied)

	return nil
}
func (self *PolicyHistory) onPrePersist() (err error){ return nil }
func (self *PolicyHistory) onPreUpdate() (err error){ return nil }
func (self *PolicyHistory) onPostUpdate() (err error){ return nil }
func (self *PolicyHistory) onPreRemove() (err error){ return nil }
func (self *PolicyHistory) onPostRemove() (err error){ return nil }


func AddPolicy(policyCreated *PolicyCreated){
	/** Example 1:  new item
	policyHistory := &PolicyHistory{}
	policyHistoryrepository.save(policyHistory)

	policyApplied := NewPolicyApplied()
	model.Copy(policyApplied, policyHistory)
	Publish(policyApplied)
	policyDenied := NewPolicyDenied()
	model.Copy(policyDenied, policyHistory)
	Publish(policyDenied)
	*/

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(policyCreated.id, 10, 64)
	policyHistory, err := policyHistoryrepository().FindById(int(id))
	if err != nil {

	}
	*/
}
