package dao

import (
	"GoChat/common"
	"GoChat/global"
	"GoChat/models"
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// GetUserList 查询所有用户
func GetUserList() ([]*models.UserBasic, error) {
	var list []*models.UserBasic
	// SELECT * FROM users;
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("用户获取列表失败")
	}
	return list, nil
}

// FindUserByNameAndPwd 根据用户名和密码查询用户
func FindUserByNameAndPwd(name, password string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	// SELECT * FROM users WHERE name = "name" AND pass_word = "password" LIMIT 1;
	if tx := global.DB.Where("name=? and pass_word=?", name, password).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}

	token := strconv.Itoa(int(time.Now().Unix()))

	// 登录识别
	temp := common.Md5encoder(token)

	// 更新登录时间
	// UPDATE users SET login_time = "timeStamp" WHERE id = "user.ID";
	if tx := global.DB.Model(&user).Where("id=?", user.ID).Update("identity", temp); tx.RowsAffected == 0 {
		return nil, errors.New("写入identity失败")
	}
	return &user, nil
}

// FindUserByName 根据用户名查询用户
// 用户登录时使用
func FindUserByName(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	// SELECT * FROM users WHERE name = "name" LIMIT 1;
	if tx := global.DB.Where("name=?", name).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

// 用户注册时使用
func FindUser(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	// SELECT * FROM users WHERE name = "name" LIMIT 1;
	if tx := global.DB.Where("name=?", name).First(&user); tx.RowsAffected == 1 {
		return nil, errors.New("当前用户已存在")
	}
	return &user, nil
}

// FindUserByPhone 根据手机号查询用户
func FindUserByPhone(phone string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	// SELECT * FROM users WHERE phone = "phone" LIMIT 1;
	if tx := global.DB.Where("phone = ?", phone).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

// FindUserByEmail 根据邮箱查询用户
func FindUerByEmail(email string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("email = ?", email).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

// FindUserID 根据ID查询用户
func FindUserID(ID uint) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where(ID).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

// CreateUser 创建用户
func CreateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Create(&user)
	if tx.RowsAffected == 0 {
		zap.S().Info("创建新用户失败")
		return nil, errors.New("新增用户失败")
	}
	return &user, nil
}

// UpdateUser 更新用户
func UpdateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Model(&user).Updates(models.UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Salt:     user.Salt,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("更新用户失败")
		return nil, errors.New("更新用户失败")
	}
	return &user, nil
}

// DeleteUser 删除用户
func DeleteUser(user models.UserBasic) error {
	// DELETE FROM users WHERE id = "user.ID";
	if tx := global.DB.Delete(&user); tx.RowsAffected == 0 {
		zap.S().Info("删除失败")
		return errors.New("删除用户失败")
	}
	return nil
}
