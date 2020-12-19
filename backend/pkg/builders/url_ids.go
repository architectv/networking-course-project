package builders

import (
	"yak/backend/pkg/models"
)

type UrlIdsBuilder struct {
	UrlIds *models.UrlIds
}

func NewUrlIdsBuilder() *UrlIdsBuilder {
	urlIds := &models.UrlIds{
		ProjectId: 0,
		BoardId:   0,
		ListId:    0,
		TaskId:    0,
	}
	return &UrlIdsBuilder{UrlIds: urlIds}
}

func (u *UrlIdsBuilder) Build() *models.UrlIds {
	return u.UrlIds
}

func (u *UrlIdsBuilder) WithProject(id int) *UrlIdsBuilder {
	u.UrlIds.ProjectId = id
	return u
}

func (u *UrlIdsBuilder) WithBoard(id int) *UrlIdsBuilder {
	u.UrlIds.BoardId = id
	return u
}

func (u *UrlIdsBuilder) WithList(id int) *UrlIdsBuilder {
	u.UrlIds.ListId = id
	return u
}

func (u *UrlIdsBuilder) WithTask(id int) *UrlIdsBuilder {
	u.UrlIds.TaskId = id
	return u
}
