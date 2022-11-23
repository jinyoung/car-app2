package car

import (
	"gopkg.in/jeevatkm/go-model.v1"

	"strconv"

	"gorm.io/gorm" //JPA in Java
)

type PolicyHistory struct {
	gorm.Model                 // extends Entity;  @Entity
	Id                  int    `gorm:"primaryKey" json:"id" type:"int"`
	CarId               int    `json:"carId"`
	PolicyApplicationId int    `json:"policyApplicationId"`
	Status              string `json:"status"`
}

func (this *PolicyHistory) onPostPersist() (err error) { // TODO: Override in Golang
	// policyApplied := NewPolicyApplied()
	// model.Copy(policyApplied, this) // BeanUtils.copyProperties(src, trg)

	// Publish(policyApplied) // Event publish

	// policyDenied := NewPolicyDenied()
	// model.Copy(policyDenied, this)

	//Publish(policyDenied)

	return nil // null
}

func (self *PolicyHistory) onPrePersist() (err error) { return nil }
func (self *PolicyHistory) onPreUpdate() (err error)  { return nil }
func (self *PolicyHistory) onPostUpdate() (err error) { return nil }
func (self *PolicyHistory) onPreRemove() (err error)  { return nil }
func (self *PolicyHistory) onPostRemove() (err error) { return nil }

func AddPolicy(policyCreated *PolicyCreated) {
	/** Example 1:  new item	*/
	policyHistory := &PolicyHistory{}

	//s2 := strconv.FormatInt(int(policyCreated.CarId), 10)
	policyHistory.CarId, _ = strconv.Atoi(policyCreated.CarId)
	policyHistory.PolicyApplicationId = policyCreated.Id

	policyHistoryrepository.save(policyHistory)

	policyApplied := NewPolicyApplied()
	model.Copy(policyApplied, policyHistory)
	Publish(policyApplied)

	// policyDenied := NewPolicyDenied()
	// model.Copy(policyDenied, policyHistory)
	// Publish(policyDenied)

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(policyCreated.id, 10, 64)
	policyHistory, err := policyHistoryrepository().FindById(int(id))
	if err != nil {

	}
	*/
}
