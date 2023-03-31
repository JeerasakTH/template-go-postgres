package controller

import (
	"context"

	db "github.com/JeerasakTH/template-go-postgres/database"
	"github.com/JeerasakTH/template-go-postgres/model"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func GetAllCouponDetail(resource *db.PostgresDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := "select id, name, start_date, end_date, status, coupon_count, coupon_type, reward from coupon_detail"
		var couponDetail []model.CouponDetail
		err := resource.DB.Select(&couponDetail, query)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Select error",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
			"data":    couponDetail,
		})
	}
}

func GetCouponDetailByID(resource *db.PostgresDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		couponID := c.Query("id")
		query := "select id, name, start_date, end_date, status, coupon_count, coupon_type, reward from coupon_detail where id=$1"
		couponDetail := model.CouponDetail{}
		err := resource.DB.Get(&couponDetail, query, couponID)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Insert Error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    couponDetail,
		})
	}
}

func CreateCoupon(resource *db.PostgresDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		tx, err := resource.DB.Begin()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Begin Error",
			})
			return
		}

		couponDetail := model.CouponDetail{}
		if err := c.ShouldBind(&couponDetail); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "ShouldBind Error",
			})
			return
		}

		q := "insert into coupon_detail (id, name, start_date, end_date, status, coupon_count, coupon_type, reward) values (:id, :name, :start_date, :end_date, :status, :coupon_count, :coupon_type, :reward)"
		result, err := resource.DB.NamedExec(q, &couponDetail)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Insert Error",
			})
			return
		}

		// Check ความเปลี่ยนแปลง
		affected, err := result.RowsAffected()
		if err != nil {
			// ถ้าพังตรงไหนก็โรลแบคกลับ
			tx.Rollback()
			return
		}

		if affected <= 0 {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Cannot insert",
			})
			return
		}

		// ผ่านทั้งหมดให้คอมมิท
		err = tx.Commit()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Commit error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    couponDetail,
		})
	}
}

func UpdateCouponByID(resource *db.PostgresDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Query("name")
		id := c.Query("id")
		tx, err := resource.DB.Begin()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Begin Error",
			})
			return
		}

		query := "update coupon_detail set name=$1 where id=$2"
		result, err := resource.DB.Exec(query, name, id)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Insert Error",
			})
			return
		}

		// Check ความเปลี่ยนแปลง
		affected, err := result.RowsAffected()
		if err != nil {
			// ถ้าพังตรงไหนก็โรลแบคกลับ
			tx.Rollback()
			return
		}

		if affected <= 0 {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Cannot insert",
			})
			return
		}

		// ผ่านทั้งหมดให้คอมมิท
		err = tx.Commit()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Commit error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    map[string]string{"name": name},
		})
	}
}

func DeleteCoupon(resource *db.PostgresDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Query("id")
		query := "delete from coupon_detail where id=$1"
		result, err := resource.DB.Exec(query, id)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Delete Error",
			})
			return
		}

		// Check ความเปลี่ยนแปลง
		affected, err := result.RowsAffected()
		if affected <= 0 {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"data":    "Cannot Delete",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    nil,
		})
	}
}
