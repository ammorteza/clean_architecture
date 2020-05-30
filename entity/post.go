package entity

type Post struct {
	ID        		    uint 			`sql:"primary_key"`
	Comment 			string 			`sql:"column:comment;unique"json:"comment"`
	
	UserId 				uint			`sql:"column:uId;foreignkey:uId"json:"uId"`
	Users				User			`json:"users"`	
}
