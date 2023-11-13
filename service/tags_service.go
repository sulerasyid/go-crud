package service

import (
	"github.com/sulerasyid/go-crud/data/request"
	"github.com/sulerasyid/go-crud/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
