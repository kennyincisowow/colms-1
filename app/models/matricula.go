package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Matricula struct {
	Id string `gorm:"primary_key;"`
	//Fecha  string
	Semestre string
	AlumnoId string `gorm:"size:191"`
	Alumno   Alumno // `gorm:"embedded"`
}

func (tab Matricula) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

func (tab *Matricula) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Alumno   Alumno //para crear el FK `gorm:"foreignkey:AlumnoId"`
