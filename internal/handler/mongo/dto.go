package mongo

type dtoGet struct {
	Collection string                 `json:"collection"`
	Query      map[string]interface{} `json:"query"`
}

type dtoModify struct {
	Collection string                   `json:"collection"`
	Query      []map[string]interface{} `json:"query"`
}
