package datastructures

import "time"

// RepoCreateEvent represents the payload for a repository creation event webhook
type RepoCreateEvent struct {
	Action       string       `json:"action"`
	Repository   Repository   `json:"repository"`
	Organization User         `json:"organization"`
	Sender       User         `json:"sender"`
}

type Repository struct {
	ID                           int              `json:"id"`
	Owner                        User             `json:"owner"`
	Name                         string           `json:"name"`
	FullName                     string           `json:"full_name"`
	Description                  string           `json:"description"`
	Empty                        bool             `json:"empty"`
	Private                      bool             `json:"private"`
	Fork                         bool             `json:"fork"`
	Template                     bool             `json:"template"`
	Mirror                       bool             `json:"mirror"`
	Size                         int              `json:"size"`
	Language                     string           `json:"language"`
	LanguagesURL                 string           `json:"languages_url"`
	HTMLURL                      string           `json:"html_url"`
	URL                          string           `json:"url"`
	Link                         string           `json:"link"`
	SSHURL                       string           `json:"ssh_url"`
	CloneURL                     string           `json:"clone_url"`
	OriginalURL                  string           `json:"original_url"`
	Website                      string           `json:"website"`
	StarsCount                   int              `json:"stars_count"`
	ForksCount                   int              `json:"forks_count"`
	WatchersCount                int              `json:"watchers_count"`
	OpenIssuesCount              int              `json:"open_issues_count"`
	OpenPRCounter                int              `json:"open_pr_counter"`
	ReleaseCounter               int              `json:"release_counter"`
	DefaultBranch                string           `json:"default_branch"`
	Archived                     bool             `json:"archived"`
	CreatedAt                    time.Time        `json:"created_at"`
	UpdatedAt                    time.Time        `json:"updated_at"`
	ArchivedAt                   time.Time        `json:"archived_at"`
	Permissions                  Permissions      `json:"permissions"`
	HasIssues                    bool             `json:"has_issues"`
	InternalTracker              InternalTracker  `json:"internal_tracker"`
	HasWiki                      bool             `json:"has_wiki"`
	HasPullRequests              bool             `json:"has_pull_requests"`
	HasProjects                  bool             `json:"has_projects"`
	ProjectsMode                 string           `json:"projects_mode"`
	HasReleases                  bool             `json:"has_releases"`
	HasPackages                  bool             `json:"has_packages"`
	HasActions                   bool             `json:"has_actions"`
	IgnoreWhitespaceConflicts    bool             `json:"ignore_whitespace_conflicts"`
	AllowMergeCommits            bool             `json:"allow_merge_commits"`
	AllowRebase                  bool             `json:"allow_rebase"`
	AllowRebaseExplicit          bool             `json:"allow_rebase_explicit"`
	AllowSquashMerge             bool             `json:"allow_squash_merge"`
	AllowFastForwardOnlyMerge    bool             `json:"allow_fast_forward_only_merge"`
	AllowRebaseUpdate            bool             `json:"allow_rebase_update"`
	DefaultDeleteBranchAfterMerge bool            `json:"default_delete_branch_after_merge"`
	DefaultMergeStyle            string           `json:"default_merge_style"`
	DefaultAllowMaintainerEdit   bool             `json:"default_allow_maintainer_edit"`
	AvatarURL                    string           `json:"avatar_url"`
	Internal                     bool             `json:"internal"`
	MirrorInterval               string           `json:"mirror_interval"`
	ObjectFormatName             string           `json:"object_format_name"`
	MirrorUpdated                time.Time        `json:"mirror_updated"`
	Topics                       []string         `json:"topics"`
	Licenses                     []string         `json:"licenses"`
}

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

type InternalTracker struct {
	EnableTimeTracker               bool `json:"enable_time_tracker"`
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
	EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
}

type User struct {
	ID               int       `json:"id"`
	Login            string    `json:"login"`
	LoginName        string    `json:"login_name"`
	SourceID         int       `json:"source_id"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	AvatarURL        string    `json:"avatar_url"`
	HTMLURL          string    `json:"html_url"`
	Language         string    `json:"language"`
	IsAdmin          bool      `json:"is_admin"`
	LastLogin        time.Time `json:"last_login"`
	Created          time.Time `json:"created"`
	Restricted       bool      `json:"restricted"`
	Active           bool      `json:"active"`
	ProhibitLogin    bool      `json:"prohibit_login"`
	Location         string    `json:"location"`
	Website          string    `json:"website"`
	Description      string    `json:"description"`
	Visibility       string    `json:"visibility"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	StarredReposCount int      `json:"starred_repos_count"`
	Username         string    `json:"username"`
}
