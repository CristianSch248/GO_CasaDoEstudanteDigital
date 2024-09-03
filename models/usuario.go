package models

type Usuario struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Nome          string `gorm:"type:varchar(255)"`
	Email         string `gorm:"type:varchar(255);unique"`
	Senha         string `gorm:"type:varchar(255)"`
	Tipo          int    `gorm:"type:int"`
	Matricula     string `gorm:"type:varchar(255)"`
	Telefone      string `gorm:"type:varchar(255)"`
	Ativo         bool   `gorm:"default:true"`
	IDApartamento uint   `gorm:"type:int;constraint:OnDelete:CASCADE"`
	Vagas         []Vaga `gorm:"foreignKey:IDAluno;constraint:OnDelete:CASCADE"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
