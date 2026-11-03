package main

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Membership
	Name string
}

func newUser(name string, membershipType string) User {
	messageCharLimit := 100
	if membershipType == "premium" {
		messageCharLimit = 1_000
	}

	return User{
		Name: name,
		/*
			Membership: Membership{
				Type: membershipType,
				MessageCharLimit: messageCharLimit,
			},
		*/
		Type:             membershipType,
		MessageCharLimit: messageCharLimit,
	}
}
