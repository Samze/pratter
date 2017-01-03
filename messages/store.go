package messages

type MessageStore interface {
	AddMessage(user string, msg Message)
	GetMessages(user string) []Message
}

type MemoryMessageStore struct {
	messages map[string][]Message
}

func NewMemoryMessageStore() MemoryMessageStore {
	messages := make(map[string][]Message)
	return MemoryMessageStore{messages}
}

func (s *MemoryMessageStore) AddMessage(user string, msg Message) {
	s.messages[user] = append(s.messages[user], msg)
}

func (s *MemoryMessageStore) GetMessages(user string) []Message {
	return s.messages[user]
}
