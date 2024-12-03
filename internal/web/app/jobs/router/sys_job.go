package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/internal/web/app/jobs/apis"
	models2 "go-admin/internal/web/app/jobs/models"
	dto2 "go-admin/internal/web/app/jobs/service/dto"
	"go-admin/internal/web/middleware"
	actions2 "go-admin/internal/web/middleware/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysJobRouter)
}

// 需认证的路由代码
func registerSysJobRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/sysjob").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		sysJob := &models2.SysJob{}
		r.GET("", actions2.PermissionAction(), actions2.IndexAction(sysJob, new(dto2.SysJobSearch), func() interface{} {
			list := make([]models2.SysJob, 0)
			return &list
		}))
		r.GET("/:id", actions2.PermissionAction(), actions2.ViewAction(new(dto2.SysJobById), func() interface{} {
			return &dto2.SysJobItem{}
		}))
		r.POST("", actions2.CreateAction(new(dto2.SysJobControl)))
		r.PUT("", actions2.PermissionAction(), actions2.UpdateAction(new(dto2.SysJobControl)))
		r.DELETE("", actions2.PermissionAction(), actions2.DeleteAction(new(dto2.SysJobById)))
	}
	sysJob := apis.SysJob{}

	v1.GET("/job/remove/:id", sysJob.RemoveJobForService)
	v1.GET("/job/start/:id", sysJob.StartJobForService)
}
