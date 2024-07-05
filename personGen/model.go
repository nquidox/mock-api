package personGen

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
)

var DBC *gorm.DB

type Person struct {
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Surname        string `json:"surname"`
	Address        string `json:"address"`
	PassportSerie  int    `json:"passportSerie"`
	PassportNumber int    `json:"passportNumber"`
}

type RandomData struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func (r *RandomData) TableName(name string) string {
	return name
}

func (r *RandomData) Count(tn string) int {
	DBC.Table(tn).Last(r)
	return r.Key
}

func (r *RandomData) Get(tn string, key int) string {
	DBC.Table(tn).Where("key = ?", key).First(r)
	return r.Value
}

func (p *Person) Create() {
	gender := true
	if rand.Intn(2) == 1 {
		gender = false
	}

	var r RandomData
	if gender == true {
		p.Name = r.Get("male_names", rand.Intn(r.Count("male_names")))
		p.Surname = r.Get("male_surnames", rand.Intn(r.Count("male_patronymics")))
		p.Patronymic = r.Get("male_patronymics", rand.Intn(r.Count("male_surname")))
	} else {
		p.Name = r.Get("female_names", rand.Intn(r.Count("female_names")))
		p.Surname = r.Get("female_surnames", rand.Intn(r.Count("female_patronymics")))
		p.Patronymic = r.Get("female_patronymics", rand.Intn(r.Count("female_surname")))
	}
	addrKey := rand.Intn(r.Count("addresses"))
	p.Address = r.Get("addresses", addrKey)
}

func (p *Person) Read() error {
	err := DBC.Table("people").
		Where("passport_serie = ? AND passport_number = ?", p.PassportSerie, p.PassportNumber).
		First(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Person) Save() {
	err := DBC.Create(p).Error
	if err != nil {
		log.Error(err)
	}
}
