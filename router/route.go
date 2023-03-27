package router

import (
	"github.com/JeerasakTH/template-go-postgres/controller"
	db "github.com/JeerasakTH/template-go-postgres/database"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, resource *db.PostgresDB) {
	r.POST("/coupon", controller.CreateCoupon(resource))
	r.GET("/coupon", controller.GetAllCouponDetail(resource))
	r.GET("/couponID", controller.GetCouponDetailByID(resource))
	r.POST("/update_couponID", controller.UpdateCouponByID(resource))
	r.POST("/delete_couponID", controller.DeleteCoupon(resource))
}
