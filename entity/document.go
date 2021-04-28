package entity

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	DocumentNo  string `json:"document_no"`
	Description string `json:"description"`
	Disposition string `json:"remark"`
	IsApproved  int    `json:"is_approved"`
}
