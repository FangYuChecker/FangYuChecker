package Init

type User struct {
	Balance float64
}

var User1 User
var Judge bool

func Init() {
	User1.Balance = 0.0
	Judge = true
}
