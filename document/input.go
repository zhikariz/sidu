package document

type CreateDocumentInput struct {
	DocumentNo  string `json:"document_no" binding:"required"`
	Description string `json:"description" binding:"required"`
	Disposition string `json:"disposition" binding:"required"`
}

type UpdateDocumentInput struct {
	DocumentNo  string `json:"document_no" binding:"required"`
	Description string `json:"description" binding:"required"`
	Disposition string `json:"disposition" binding:"required"`
	IsApproved  int    `json:"is_approved" binding:"required"`
}

type DocumentUriInput struct {
	ID int `uri:"id" binding:"required"`
}
