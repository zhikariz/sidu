package document

import . "sidu/entity"

type DocumentFormatter struct {
	ID          int    `json:"id"`
	DocumentNo  string `json:"document_no"`
	Description string `json:"description"`
	Disposition string `json:"disposition"`
	IsApproved  string `json:"is_approved"`
}

func FormatDocument(document Document) (documentFormatter DocumentFormatter) {
	documentFormatter.ID = int(document.ID)
	documentFormatter.DocumentNo = document.DocumentNo
	documentFormatter.Description = document.Description
	documentFormatter.Disposition = document.Disposition
	if document.IsApproved == 1 {
		documentFormatter.IsApproved = "Approved"
	} else {
		documentFormatter.IsApproved = "Pending"
	}
	return
}

func FormatDocuments(documents []Document) []DocumentFormatter {
	documentsFormatter := []DocumentFormatter{}
	for _, document := range documents {
		documentFormatter := FormatDocument(document)
		documentsFormatter = append(documentsFormatter, documentFormatter)
	}
	return documentsFormatter
}
