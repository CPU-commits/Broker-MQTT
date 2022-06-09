package auth

type Auth struct {
	Users         map[string]string   // Usernames (key) with passwords (value)
	AllowedTopics map[string][]string // Usernames and topics
}

func (a *Auth) Authenticate(user, password []byte) bool {
	if pass, ok := a.Users[string(user)]; ok && pass == string(password) {
		return true
	}
	return false
}

func (a *Auth) ACL(user []byte, topic string, write bool) bool {
	// Topics restricctions
	if topics, ok := a.AllowedTopics[string(user)]; ok {
		for _, t := range topics {
			if t == topic {
				return true
			}
		}
		return false
	}
	// Allow all topics
	return true
}
