package webmodel

type RoleInput struct {
    Name 		string  `json:"name" binding:"required"`
}