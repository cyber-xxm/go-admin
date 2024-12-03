package dto

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/database/dto"
)

type Index interface {
	Generate() Index
	Bind(ctx *gin.Context) error
	GetPageIndex() int
	GetPageSize() int
	GetNeedSearch() interface{}
}

type Control interface {
	Generate() Control
	Bind(ctx *gin.Context) error
	GenerateM() (dto.ActiveRecord, error)
	GetId() interface{}
}
