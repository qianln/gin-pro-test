package models

import (
	"database/sql/driver"
	"fmt"
	"gin-pro/app/core/system"
	"gin-pro/app/global/consts"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	ID        int64     `gorm:"primaryKey" json:"id"`
	CreatedAt BaseTime  `gorm:"created_at" json:"created_at"` //日期时间字段统一设置为字符串即可
	UpdatedAt BaseTime  `gorm:"updated_at" json:"updated_at"`
	DeletedAt *BaseTime `json:"deleted_at"`
}

func UseDbConn() *gorm.DB {
	if system.DbMysql == nil {
		system.ZapLog.Fatal(fmt.Sprintf(consts.ErrorsGormNotInitGlobalPointer, "Mysql", "Mysql"))
	}
	return system.DbMysql
}

type BaseTime struct {
	time.Time
}

func (t BaseTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(system.DateFormat))
	return []byte(output), nil
}

func (t *BaseTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *BaseTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = BaseTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

/**
自定义一些非终止条件
*/

func (b *BaseModel) Where(key string, operator string, value any) *BaseModel {
	return &BaseModel{
		DB: b.DB.Where(fmt.Sprintf("%s %s ?", key, operator), value),
	}
}

func (b *BaseModel) OrderBy(column string, direction ...string) *BaseModel {
	if direction[0] == "" {
		direction[0] = "asc"
	}
	return &BaseModel{
		DB: b.DB.Order(fmt.Sprintf("%s %s", column, direction[0])),
	}
}

func (b *BaseModel) WhereIn(key string, value any) *BaseModel {
	return &BaseModel{
		DB: b.DB.Where(fmt.Sprintf("%s  in  ?", key), value),
	}
}

func (b *BaseModel) WhereNotIn(key string, value any) *BaseModel {
	return &BaseModel{
		DB: b.DB.Not(key, value),
	}
}

func (b *BaseModel) WhereBetween(key string, value1, value2 string) *BaseModel {
	return &BaseModel{
		DB: b.DB.Where(fmt.Sprintf("%s  between  ? AND ?", key), value1, value2),
	}
}

func (b *BaseModel) WhereLike(key string, value string) *BaseModel {
	return &BaseModel{
		DB: b.DB.Where(fmt.Sprintf("%s  LIKE  ? ", key), value),
	}
}

func (b *BaseModel) WhereMap(value map[string]any) *BaseModel {
	return &BaseModel{
		DB: b.DB.Where(value),
	}
}

func (b *BaseModel) OrWhere(key string, operator string, value any) *BaseModel {
	return &BaseModel{
		DB: b.DB.Or(fmt.Sprintf("%s %s ?", key, operator), value),
	}
}

func (b *BaseModel) When(key bool, value map[string]any) *BaseModel {
	if key {
		return &BaseModel{
			DB: b.DB.Where(value),
		}
	}
	return &BaseModel{
		DB: b.DB,
	}
}
