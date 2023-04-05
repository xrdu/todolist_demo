package model

// SessionStorage 存放用户会话的容器，简单用内存中一个Map实现
type SessionStorage struct {
	data map[string]*User
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		data: make(map[string]*User),
	}
}

func (s *SessionStorage) Put(sessionId string, user *User) {
	s.data[sessionId] = user
}

func (s *SessionStorage) Get(sessionId string) *User {
	return s.data[sessionId]
}

func (s *SessionStorage) Del(sessionId string) {
	delete(s.data, sessionId)
}

func (s *SessionStorage) Exist(sessionId string) bool {
	return s.Get(sessionId) != nil
}
