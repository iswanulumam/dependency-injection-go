package following

type Following struct {
	ID       int
	Username string
	FullName string
}

func NewFollowing(username, fullName string) Following {
	return Following{0, username, fullName}
}
