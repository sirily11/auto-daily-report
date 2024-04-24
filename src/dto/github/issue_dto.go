package github

type ActiveLockReason string

const (
	OffTopicLockReason  ActiveLockReason = "off-topic"
	ResolvedLockReason  ActiveLockReason = "resolved"
	SpamLockReason      ActiveLockReason = "spam"
	TooHeatedLockReason ActiveLockReason = "too heated"
)

const (
	AssignedAction   string = "assigned"
	ClosedAction     string = "closed"
	DeletedAction    string = "deleted"
	EditedAction     string = "edited"
	LabeledAction    string = "labeled"
	OpenedAction     string = "opened"
	ReopenedAction   string = "reopened"
	UnassignedAction string = "unassigned"
	UnlabeledAction  string = "unlabeled"
)

type Label struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AssigneeDto struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

type IssueDto struct {
	Id               int              `json:"id"`
	ActiveLockReason ActiveLockReason `json:"active_lock_reason,omitempty"`
	Assignee         AssigneeDto      `json:"assignee"`
	Body             string           `json:"body,omitempty"`
	ClosedAt         string           `json:"closed_at,omitempty"`
	CreatedAt        string           `json:"created_at"`
	Labels           []Label          `json:"labels"`
	Title            string           `json:"title"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

type IssueWebhookDTO struct {
	Action     string     `json:"action"`
	Issue      IssueDto   `json:"issue"`
	Repository Repository `json:"repository"`
	// Whether the issue is reported or not
	IsReported bool `json:"is_reported"`
}
