package models

type TitanSoldier struct {
	Id             int64  `json:"soldier_id"`
	Name           string `json:"soldier_name" binding:"required"`
	Division       string `json:"soldier_division" binding:"required"`
	Specialization string `json:"soldier_specialization" binding:"required"`
}
