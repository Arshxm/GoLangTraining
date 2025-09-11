package main

import "gorm.io/gorm"

type TLDRProvider interface {
	Retrieve(string) string
	List() []string
}

type TLDREntity struct {
	gorm.Model
	Key string `gorm:"primaryKey;size:100"`
	Val string `gorm:"size:1000"`
}

func (t *TLDREntity) Retrieve(key string) string {
	if {
		return ""
	}
	return t.Val
}

func (t *TLDREntity) List() []string {
	return []string{t.Key}
}
