package models

type Vaga struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	DtEntrada     string  `gorm:"type:date"`
	DtSaida       string  `gorm:"type:date"`
	Observacao    string  `gorm:"type:varchar(255)"`
	Ativo         bool    `gorm:"type:boolean"`
	IDAluno       uint    `gorm:"type:int;constraint:OnDelete:CASCADE"`
	IDApartamento uint    `gorm:"type:int;constraint:OnDelete:CASCADE"`
	Usuario       Usuario `gorm:"foreignKey:IDAluno"`
}

func (Vaga) TableName() string {
	return "vagas"
}
