package model

import (
	"berpar/common"
	"time"

	"gorm.io/gorm"
)

type Consumer struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Username     string    `gorm:"type:varchar(255)" json:"username"`
	Password     string    `gorm:"type:varchar(100)" json:"password"`
	Sex          int       `gorm:"type:tinyint" json:"sex"`
	PhoneNum     string    `gorm:"type:char(15)" json:"phoneNum"`
	Email        string    `gorm:"type:char(30)" json:"email"`
	Birth        time.Time `gorm:"type:datetime" json:"birth"`
	Introduction string    `gorm:"type:varchar(255)" json:"introduction"`
	Location     string    `gorm:"type:varchar(45)" json:"location"`
	Avator       string    `gorm:"type:varchar(255)" json:"avator"`
	CreateTime   time.Time `gorm:"type:datetime" json:"createTime"`
	UpdateTime   time.Time `gorm:"type:dateTime" json:"updateTime"`
}

func AddUser(name, passWord string) *common.ResponseBody {
	user := Consumer{Username: name, Password: passWord}
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	err := Db.Create(&user).Error
	if err != nil {
		return nil
	}
	return common.SuccessMessage("成功添加用户！", user)
}

func VerifyPasswd(userName, passWord string) *common.ResponseBody {
	var user []Consumer
	err := Db.Model(&Consumer{}).
		Where(&Consumer{Username: userName, Password: passWord}).
		Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return common.SuccessMessage("登录失败", 0)
	}
	return common.SuccessMessage("登录成功", user)
}

func GetAllUser() *common.ResponseBody {
	var userList []Consumer
	err := Db.Model(&Consumer{}).Where(&Consumer{}).Find(&userList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("所有用户", userList)
}

func SelectByUserId(id uint) *common.ResponseBody {
	if id==0 {
		return nil
	}		
	var user []Consumer
	err := Db.Model(&Consumer{}).Where(&Consumer{ID: id}).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("用户id", user)
}

func DeleteUser(id uint) *common.ResponseBody {
	err := Db.Delete(&Consumer{}, id).Error
	if err != nil {
		return common.SuccessMessage("成功删除失败！", id)
	}
	return common.SuccessMessage("成功删除用户！", id)
}

func UpdateUserMsg(id int,email string)*common.ResponseBody{
	err:=Db.Model(&Consumer{ID: uint(id)}).
	Select("email").
	Updates(map[string]interface{}{"email": email}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新用户信息！",email)
}

func UpdateUserPassword(id int,passWord string)*common.ResponseBody{
	err:=Db.Model(&Consumer{ID: uint(id)}).
	Select("password").
	Updates(map[string]interface{}{"password": passWord}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新用户密码！",passWord)
}

func UpdateUserPic(id int,avator string)*common.ResponseBody{
	err:=Db.Model(&Consumer{ID: uint(id)}).
	Select("avator").
	Updates(map[string]interface{}{"avator": avator}).
	Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return common.SuccessMessage("成功更新用户图片！",avator)
}