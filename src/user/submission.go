package user

import ()

type User struct {
	UserID       string
	Username     string
	FirstName    string
	LastName     string
	Email        string
	UTCTimeStamp int64 // Remember: Only new records are admissible into the DB.
}
