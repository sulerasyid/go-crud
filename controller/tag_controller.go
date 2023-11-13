package controller

import (
	"github.com/sulerasyid/go-crud/data/request"
	"github.com/sulerasyid/go-crud/data/response"
	"github.com/sulerasyid/go-crud/helper"

	"net/http"
	"strconv"

	"github.com/sulerasyid/go-crud/service"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService service.TagsService
	Logger     service.Logger
}

func NewTagController(service service.TagsService, logger service.Logger) *TagController {
	return &TagController{tagService: service, Logger: logger}
}

func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)

	controller.Logger.LogAccess("%s %s %s\n", ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.URL)
	controller.Logger.LogError("%s", err)
	helper.ErrorPanic(err)

	controller.tagService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	controller.Logger.LogAccess("%s %s %s\n", ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.URL)
	controller.Logger.LogError("%s", err)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagRequest.Id = id

	controller.tagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	controller.Logger.LogAccess("%s %s %s\n", ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.URL)
	controller.Logger.LogError("%s", err)
	helper.ErrorPanic(err)
	controller.tagService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	controller.Logger.LogAccess("%s %s %s\n", ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.URL)
	controller.Logger.LogError("%s", err)
	helper.ErrorPanic(err)

	tagResponse := controller.tagService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagService.FindAll()
	controller.Logger.LogAccess("%s %s %s\n", ctx.Request.RemoteAddr, ctx.Request.Method, ctx.Request.URL)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
