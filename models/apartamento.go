package models

type Apartamento struct {
	ID     uint  `gorm:"primaryKey;autoIncrement"`
	Numero int64 `gorm:"type:bigint"`
	Bloco  int64 `gorm:"type:bigint"`
	Vagas  int   `gorm:"type:int;default:8"`
}

func (Apartamento) TableName() string {
	return "apartamentos"
}
