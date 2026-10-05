package token

import "time"


//interface to declare functions that will be implemnted in jwt_maker file
type Maker interface {
	CreateToken(email string, role string, duration time.Duration)(string,error)
}

