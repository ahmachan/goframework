package Model
import "goframework/framework"


type User struct {
	Id     int       `gorm:"column:id" json:"id"`
	Name   string    `gorm:"column:name" json:"name"`
	Passwd string    `gorm:"column:passwd" json:"passwd"`
	Email  string    `gorm:"column:email" json:"email"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfo(id int) (User, error) {

	User := User{}
	err := framework.GetORM().Where("id = ?", id).First(&User).Error
	return User, err
}

func GetUserList() ([]User,error) {
	UserList := []User{}
	err := framework.GetORM().Find(&UserList).Error
	if err != nil {
		return nil,err
	}
	return UserList,nil
}

func (this *User)GetInfoById(id int) (User, error) {

	User := User{}
	err := framework.GetORM().Where("id = ?", id).First(&User).Error
	return User, err
}

func (this *User)GetList() ([]User,error) {
	UserList := []User{}
	err := framework.GetORM().Find(&UserList).Error
	if err != nil {
		return nil,err
	}
	return UserList,nil
}
