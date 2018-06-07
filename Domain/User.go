package Domain

import (
	"phalgo-sample/Model"
	"errors"
)

func GetHello(id int) (Model.User, error) {

	user, err := Model.GetInfo(id)
	if err != nil {
		return user, errors.New("UserInfo There is no")
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
