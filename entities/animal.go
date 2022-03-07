package entities

type Animal struct {
	ID    int    `gorm:"autoIncrement;primaryKey"`
	Name  string `gorm:"size:100;unique;not null" required:"true"`
	Class string `required:"true" gorm:"not null;size:20"`
	Legs  int8   `required:"true" gorm:"not null"`
}

// func (animal Animal)Create()  {

// }
