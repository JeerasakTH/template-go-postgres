package model

type CouponDetail struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	StartDate   string  `db:"start_date" json:"start_date"`
	EndDate     string  `db:"end_date" json:"end_date"`
	Status      int     `db:"status" json:"status"` // open-close
	CouponCount int     `db:"coupon_count" json:"coupon_count"`
	CouponType  int     `db:"coupon_type" json:"coupon_type"` // 1 one 2 many
	Reward      float64 `db:"reward" json:"reward"`
}

type CouponCode struct {
	ID       int    `db:"id" json:"id"`
	Status   int    `db:"status" json:"status"` // use-unused
	Total    int    `db:"total" json:"total"`   // ทั้งหมด
	Count    int    `db:"count" json:"count"`   // คงเหลือ
	Code     string `db:"code" json:"code"`
	ConfigID int    `db:"config_id" json:"config_id"`
}
