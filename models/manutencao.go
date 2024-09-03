package models

type Manutencao struct {
	ID                   uint   `gorm:"primaryKey;autoIncrement"`
	IDAluno              uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	IDUsuarioConfirmador uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	IDApartamento        uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	Caso                 string `gorm:"type:varchar(255)"`
	MaterialUsado        string `gorm:"type:varchar(255)"`
	DescricaoAtividade   string `gorm:"type:varchar(255)"`
	DtManutencao         string `gorm:"type:date"`
	HoraManutencao       string `gorm:"type:time"`
	Status               int    `gorm:"type:int"`
}

func (Manutencao) TableName() string {
	return "manutencoes"
}
