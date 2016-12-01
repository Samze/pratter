package messages

type MessageStore struct {
	messages map[string][]message
}

func NewMessageStore() MessageStore {
	messages := make(map[string][]message)
	return MessageStore{messages}
}

func (s *MessageStore) addMessage(user string, msg message) {
	s.messages[user] = append(s.messages[user], msg)
}

func (s *MessageStore) getMessages(user string) []message {
	return s.messages[user]
}
