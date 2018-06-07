package Domain
import (
	"phalgo-sample/Model"
)


type Base struct {
	User Domain_User
}

type Domain_User struct {
	Model Model.Base
}
