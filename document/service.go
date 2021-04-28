package document

import (
	"errors"
	. "sidu/entity"
)

type Service interface {
	CreateDocument(input CreateDocumentInput) (Document, error)
	UpdateDocument(id int, input UpdateDocumentInput) (Document, error)
	DeleteDocument(id int) (bool, error)
	GetDocumentById(id int) (Document, error)
	GetDocuments() ([]Document, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) CreateDocument(input CreateDocumentInput) (Document, error) {
	document := Document{
		DocumentNo:  input.DocumentNo,
		Description: input.Description,
		Disposition: input.Disposition,
		IsApproved:  0,
	}

	newDocument, err := s.repository.Save(document)

	if err != nil {
		return newDocument, err
	}

	return newDocument, nil
}

func (s *service) UpdateDocument(id int, input UpdateDocumentInput) (Document, error) {
	document, err := s.repository.FindById(id)
	if err != nil {
		return document, err
	}

	document.DocumentNo = input.DocumentNo
	document.Description = input.Description
	document.Disposition = input.Disposition
	document.IsApproved = input.IsApproved

	updatedDocument, err := s.repository.Update(document)

	if err != nil {
		return updatedDocument, err
	}
	return updatedDocument, nil
}

func (s *service) DeleteDocument(id int) (bool, error) {
	document, err := s.repository.FindById(id)

	if err != nil {
		return false, err
	}

	isDeleted, err := s.repository.Delete(document)

	if err != nil {
		return false, err
	}

	return isDeleted, nil
}

func (s *service) GetDocumentById(id int) (Document, error) {
	document, err := s.repository.FindById(id)
	if err != nil {
		return document, err
	}

	if document.ID == 0 {
		return document, errors.New("No document found with that id")
	}
	return document, nil
}

func (s *service) GetDocuments() ([]Document, error) {
	documents, err := s.repository.FindAll()
	if err != nil {
		return documents, err
	}
	return documents, nil
}
