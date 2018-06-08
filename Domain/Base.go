package Domain
import (
	"goframework/Model"
)


type Base struct {
	User Domain_User
}

type Domain_User struct {
	Model Model.Base
}
