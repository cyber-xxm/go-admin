package actions

import (
	"errors"
	dto3 "go-admin/internal/database/dto"
	dto2 "go-admin/internal/web/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"gorm.io/gorm"
)

// IndexAction 通用查询动作
func IndexAction(m dto3.ActiveRecord, d dto2.Index, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		list := f()
		object := m.Generate()
		req := d.Generate()
		var count int64

		//查询列表
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
			return
		}

		//数据权限检查
		p := GetPermissionFromContext(c)

		err = db.WithContext(c).Model(object).
			Scopes(
				dto2.MakeCondition(req.GetNeedSearch()),
				dto2.Paginate(req.GetPageSize(), req.GetPageIndex()),
				Permission(object.TableName(), p),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(&count).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("MsgID[%s] Index error: %s", msgID, err)
			response.Error(c, 500, err, "查询失败")
			return
		}
		response.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
		c.Next()
	}
}
