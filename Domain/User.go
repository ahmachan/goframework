package Domain

import (
	"goframework/Model"
	"errors"
)

func GetHello(id int) (Model.User, error) {

	user, err := Model.GetUserInfo(id)
	if err != nil {
		return user, errors.New("empty userinfo")
	}
	return user, nil
}

func GetUserList() ([]Model.User, error) {

	user, err := Model.GetUserList()
	if err != nil {
		return user, errors.New("no userinfo list")
	}
	return user, nil

}

/*
func (this *Domain_User)GetUserInfo(id int) (Model.User, error) {

	user, err := this.Model.GetInfoById(id)
	if err != nil {
		return user, errors.New("UserInfo There is no")
	}
	return user, nil
}

func (this *Domain_User)GetUserList() ([]Model.User, error) {

	user, err := this.Model.GetList()
	if err != nil {
		return user, errors.New("UserInfo There is no")
	}
	return user, nil

}
*/
