package datastructures

// RepoCreatePushMirrorBody represents the body of a push mirror configuration for a repository creation event
type RepoCreatePushMirrorBody struct {
	Interval string `json:"interval"`
	RemoteAddress string `json:"remote_address"`
	RemoteUsername string `json:"remote_username"`
	RemotePassword string `json:"remote_password"`
	SyncOnCommit bool `json:"sync_on_commit"`
}
