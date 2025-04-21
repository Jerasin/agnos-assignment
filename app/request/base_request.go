package request

type ParamID struct {
	ID string `json:"ID"  binding:"required"`
}

type BasePaginationModel struct {
	Page      string `json:"page"`
	PageSize  string `json:"pageSize"`
	Search    string `json:"search"`
	SortField string `json:"sortField"`
	SortValue string `json:"sortValue"`
}
