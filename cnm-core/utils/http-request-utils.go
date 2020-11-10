package utils

import (

	"net/http"

)

func  GetUserId( r *http.Request) (*int, error) {
	header :=   r.Header
	token:=header.Get("Authorization")
	claim,_,err:=TokenClaim(token)
	return  claim.UserId,err
}
