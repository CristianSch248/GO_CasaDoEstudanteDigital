package models

type Patrimonio struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	IDApartamento uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	Descricao     string `gorm:"type:varchar(255)"`
	Estado        string `gorm:"type:varchar(255)"`
}

func (Patrimonio) TableName() string {
	return "patrimonios"
}
