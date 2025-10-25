package datastructures

type RepoCreatePushMirrorBody struct {
	Interval string `json:"interval"`
	RemoteAddress string `json:"remote_address"`
	RemoteUsername string `json:"remote_username"`
	RemotePassword string `json:"remote_password"`
	SyncOnCommit bool `json:"sync_on_commit"`
}
