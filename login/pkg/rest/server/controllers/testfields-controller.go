package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ppinnamani/chandu/login/pkg/rest/server/models"
	"github.com/ppinnamani/chandu/login/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TestfieldsController struct {
	testfieldsService *services.TestfieldsService
}

func NewTestfieldsController() (*TestfieldsController, error) {
	testfieldsService, err := services.NewTestfieldsService()
	if err != nil {
		return nil, err
	}
	return &TestfieldsController{
		testfieldsService: testfieldsService,
	}, nil
}

func (testfieldsController *TestfieldsController) CreateTestfields(context *gin.Context) {
	// validate input
	var input models.Testfields
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger testfields creation
	if _, err := testfieldsController.testfieldsService.CreateTestfields(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Testfields created successfully"})
}

func (testfieldsController *TestfieldsController) UpdateTestfields(context *gin.Context) {
	// validate input
	var input models.Testfields
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfields update
	if _, err := testfieldsController.testfieldsService.UpdateTestfields(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Testfields updated successfully"})
}

func (testfieldsController *TestfieldsController) FetchTestfields(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfields fetching
	testfields, err := testfieldsController.testfieldsService.GetTestfields(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testfields)
}

func (testfieldsController *TestfieldsController) DeleteTestfields(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testfields deletion
	if err := testfieldsController.testfieldsService.DeleteTestfields(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Testfields deleted successfully",
	})
}

func (testfieldsController *TestfieldsController) ListTestfields(context *gin.Context) {
	// trigger all testfields fetching
	testfields, err := testfieldsController.testfieldsService.ListTestfields()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testfields)
}

func (*TestfieldsController) PatchTestfields(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*TestfieldsController) OptionsTestfields(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*TestfieldsController) HeadTestfields(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
