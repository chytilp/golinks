package webmodel

type CategoryInput struct {
    Name 		string  `json:"name" binding:"required"`
    ParentID 	uint 	`json:"parentId"`
}