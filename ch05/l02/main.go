package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	isValidUser := func(u user) bool {
		return u.name != "" && u.number != 0
	}
	return isValidUser(mToSend.sender) && isValidUser(mToSend.recipient)
}
