package handler

import (
	"net/http"
	"sidu/document"
	. "sidu/entity"
	"sidu/helper"

	"github.com/gin-gonic/gin"
)

type documentHandler struct {
	service document.Service
}

func NewDocumentHandler(documentService document.Service) *documentHandler {
	return &documentHandler{documentService}
}

func (h *documentHandler) GetDocuments(c *gin.Context) {
	documents, err := h.service.GetDocuments()
	if err != nil {
		response := helper.APIResponse("Error to get documents", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := document.FormatDocuments(documents)
	response := helper.APIResponse("List of documents", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *documentHandler) GetDocument(c *gin.Context) {
	var input document.DocumentUriInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to get detail of document", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	detailDocument, err := h.service.GetDocumentById(input.ID)

	if err != nil {
		errorMessage := gin.H{"errors": "Document with that id not found"}
		response := helper.APIResponse("Failed to get detail of document", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := document.FormatDocument(detailDocument)
	response := helper.APIResponse("Document detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *documentHandler) CreateDocument(c *gin.Context) {
	var input document.CreateDocumentInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create document", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(User)

	if currentUser.ID == 0 {
		response := helper.APIResponse("U're not authorized to do this !", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	newDocument, err := h.service.CreateDocument(input)

	if err != nil {
		response := helper.APIResponse("Failed to create document", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := document.FormatDocument(newDocument)
	response := helper.APIResponse("Successfully create document", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *documentHandler) UpdateDocument(c *gin.Context) {
	var uri document.DocumentUriInput
	var input document.UpdateDocumentInput

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update document", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to update document", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(User)

	if currentUser.ID == 0 {
		response := helper.APIResponse("U're not authorized to do this !", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedDocument, err := h.service.UpdateDocument(uri.ID, input)

	if err != nil {
		response := helper.APIResponse("Failed to update document", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := document.FormatDocument(updatedDocument)
	response := helper.APIResponse("Successfully update document", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *documentHandler) DeleteDocument(c *gin.Context) {
	var uri document.DocumentUriInput

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to delete document", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isDeleted, err := h.service.DeleteDocument(uri.ID)

	if err != nil {
		response := helper.APIResponse("Failed to delete document", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_deleted": isDeleted,
	}

	metaMessage := "Document cannot be deleted !"

	if isDeleted {
		metaMessage = "Document has been deleted !"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}
