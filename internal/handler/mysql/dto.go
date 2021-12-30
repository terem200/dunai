package mysql

// dto - expected request body for /mysql group routes
type dto struct {
	Query string `json:"query" binding:"required"`
}
