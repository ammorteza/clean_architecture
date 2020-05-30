package entity

type User struct {
    ID        		    uint 			`sql:"primary_key"`
	Name 				string			`sql:"column:name"json:"name"`
	CommentCount		uint			`sql:"column:commentCount"json:"comment_count"`
    
    Posts    			[]Post 			`sql:"foreignkey:fId"json:"posts"`
}