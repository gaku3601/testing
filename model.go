package main

type User struct {
	ID   int64
	Name string
}

// set User's table name to be `user`
func (User) TableName() string {
	return "user"
}

type Friend struct {
	From int64
	To   int64
}

// set friend's table name to be `friend`
func (Friend) TableName() string {
	return "friend"
}
