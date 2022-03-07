package entities

type Animal struct {
	ID    int    `gorm:"autoIncrement;primaryKey" json:"id"`
	Name  string `gorm:"size:100;unique;not null" required:"true" json:"name"`
	Class string `required:"true" gorm:"not null;size:20" json:"class"`
	Legs  int    `required:"true" gorm:"not null" json:"legs"`
}
