package model

import "time"

type profile struct {
	id        string
	firstName string
	lastName  string
	email     string
	password  string
	createdAt time.Time
}

//create a type called Profiles that implements the slice of profile type
type Profiles []profile

//New Profile's constructor
func NewProfile() *profile {
	p := profile{
		createdAt: time.Now(),
	}
	return &p
}

// *profile -> pointer to a struct (somewhere in memory)
// profile -> struct

