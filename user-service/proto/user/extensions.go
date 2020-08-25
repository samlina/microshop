package microshop_user_service


import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

//User模型拓展方法
//在插入数据之前，手动设置ID的值
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}