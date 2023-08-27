package models

import (
	"github.com/lib/pq"
)

type Fact struct {
	//gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

type SegmentQuery struct { //for post/delete
	Segment string `json:"segment" gorm:"text;not null;default:null; primarykey"`
}

type UserQuery struct { // for patch
	ID     uint64   `json:"id"`
	AddSeg []string `json:"addseg"`
	DelSeg []string `json:"delseg"`
}

type User struct {
	ID       uint64         `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Segments pq.StringArray `json:"segments" gorm:"type:text[]"`
}

// type UserOne struct {
// 	ID uint64 `json:"id"`
// }
