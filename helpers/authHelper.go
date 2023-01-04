package helpers

import (
	"errors"
	"log"
	"net/http"
)

func CheckUserType(r *http.Request, user_type string) (err error) {
	var role string = r.Context().Value("userType").(string)
	if role != user_type {
		log.Println(role, user_type)
		err := errors.New("Unauthorized to access resources not matching user type")
		return err
	}
	return nil
}
func MatchUserTypeToUId(r *http.Request, user_id string) (err error) {
	UserId := r.Context().Value("UserId")
	var role string = r.Context().Value("UserType").(string)
	if UserId == user_id && role == "USER" {
		err := errors.New("Unauthorized to access resources not matching userId and user_id")
		return err
	}
	err = CheckUserType(r, role)
	return err
}
