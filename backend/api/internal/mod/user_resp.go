package mod

type UserResponse struct {
	Success bool
	Message string
}

// FriendList: mod for friend list
type FriendList struct {
	Success bool
	Friends []string
	Count   int
}

// RetrieveUpdates: mod for retrieving updates
type RetrieveUpdates struct {
	Success    bool
	Message    string
	Recipients []string
}
