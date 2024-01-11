package dto

type UserAvailabilityDTO struct {
	EmailAvailable    bool `json:"emailAvailable"`
	UsernameAvailable bool `json:"usernameAvailable"`
}
