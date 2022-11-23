package car

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PolicyHistoryDB struct {
	db *gorm.DB
}

var policyHistoryrepository *PolicyHistoryDB //TODO: PolicyHistoryRepository

func PolicyHistoryDBInit() {
	var err error
	policyHistoryrepository = &PolicyHistoryDB{}
	policyHistoryrepository.db, err = gorm.Open(sqlite.Open("PolicyHistory_table.db"), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	policyHistoryrepository.db.AutoMigrate(&PolicyHistory{})

}

func PolicyHistoryRepository() *PolicyHistoryDB {
	return policyHistoryrepository
}

func (self *PolicyHistoryDB) save(entity interface{}) error { //TODO: ?? interface{}

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *PolicyHistoryDB) GetList() []PolicyHistory { //TODO: findAll

	entities := []PolicyHistory{}
	self.db.Find(&entities)

	return entities
}

func (self *PolicyHistoryDB) FindById(id int) (*PolicyHistory, error) {
	entity := &PolicyHistory{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		return entity, nil
	}
}

func (self *PolicyHistoryDB) Delete(entity *PolicyHistory) error {
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *PolicyHistoryDB) Update(id int, params map[string]string) (*PolicyHistory, error) { //TODO: attribute dirty flag check 을 안하는? goorm
	entity := &PolicyHistory{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &PolicyHistory{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}
