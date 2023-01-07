package model

import (
	"berpar/common"

	"gorm.io/gorm"
)

type Admin struct {
	ID uint `gorm:"primarykey" json:"id"`
	Name string `gorm:"type:varchar(45)" json:"name"`
	Password string `gorm:"type:varchar(45)" json:"password"`
}


// admin
func Veritypasswd(username,password string)*common.ResponseBody{
	var admin Admin
	err:=Db.Model(&Admin{}).
	Where(&Admin{Name: username,Password: password}).
	Find(&admin).Error
	if err==gorm.ErrRecordNotFound {
		return common.AuthErrorMessage()
	}
	if err != nil{
		return nil
	}
	return common.SuccessMessage("user登录成功!",username)
}