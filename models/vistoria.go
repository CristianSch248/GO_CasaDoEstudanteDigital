package models

type Vistoria struct {
	ID                   uint   `gorm:"primaryKey;autoIncrement"`
	IDAluno              uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	IDUsuarioConfirmador uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	IDApartamento        uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	DtVistoria           string `gorm:"type:date"`
	HoraVistoria         string `gorm:"type:time"`
	Status               int    `gorm:"type:int"`
	Observacoes          string `gorm:"type:varchar(255)"`
}

func (Vistoria) TableName() string {
	return "vistorias"
}
