package entity

type Post struct {
	ID        		    uint 			`gorm:"primary_key"`
	Comment 			string 			`gorm:"column:comment"json:"comment"`
	
	UserId 				uint			`gorm:"column:uId;foreignkey:uId"json:"uId"`
	Users				User			`json:"users"`	
}
