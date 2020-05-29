package entity

type User struct {
    ID        		    uint 			`gorm:"primary_key"`
	Name 				string			`gorm:"column:name"json:"name"`
	CommentCount		uint			`gorm:"column:commentCount"json:"comment_count"`
    
    Posts    			[]Post 			`gorm:"foreignkey:fId"json:"posts"`
}