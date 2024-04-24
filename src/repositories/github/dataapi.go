package github

type DataAPIRequest struct {
	DataSource string         `json:"dataSource"`
	Database   string         `json:"database"`
	Collection string         `json:"collection"`
	Filter     map[string]any `json:"filter,omitempty"`
	Projection map[string]any `json:"projection,omitempty"`
	Sort       map[string]any `json:"sort,omitempty"`
	Limit      int            `json:"limit,omitempty"`
	Skip       int            `json:"skip,omitempty"`
	Document   any            `json:"document,omitempty"`
}
