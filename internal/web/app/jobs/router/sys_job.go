package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/internal/web/app/jobs/apis"
	"go-admin/internal/web/app/jobs/models"
	"go-admin/internal/web/app/jobs/service/dto"
	"go-admin/internal/web/middleware"
	"go-admin/internal/web/middleware/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysJobRouter)
}

// 需认证的路由代码
func registerSysJobRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/sysjob").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		sysJob := &models.SysJob{}
		r.GET("", actions.PermissionAction(), actions.IndexAction(sysJob, new(dto.SysJobSearch), func() interface{} {
			list := make([]models.SysJob, 0)
			return &list
		}))
		r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.SysJobById), func() interface{} {
			return &dto.SysJobItem{}
		}))
		r.POST("", actions.CreateAction(new(dto.SysJobControl)))
		r.PUT("", actions.PermissionAction(), actions.UpdateAction(new(dto.SysJobControl)))
		r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.SysJobById)))
	}
	sysJob := apis.SysJob{}

	v1.GET("/job/remove/:id", sysJob.RemoveJobForService)
	v1.GET("/job/start/:id", sysJob.StartJobForService)
}
