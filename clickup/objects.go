package clickup

type User struct {
	Id             int    `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty"`
	Color          string `json:"color,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
	Initials       string `json:"initials,omitempty"`
	Role           int    `json:"role,omitempty"`
	CustomRole     int    `json:"customRole,omitempty"`
	LastActive     string `json:"last_active,omitempty"`
	JoinedAt       string `json:"date_joined,omitempty"`
	InvitedAt      string `json:"date_invited,omitempty"`
}

type Member struct {
	User      User `json:"user,omitempty"`
	InvitedBy User `json:"invited_by,omitempty"`
}

type Team struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Color   string   `json:"color,omitempty"`
	Avatar  string   `json:"avatar,omitempty"`
	Members []Member `json:"members,omitempty"`
}

type Status struct {
	Status     string `json:"status,omitempty"`
	Type       string `json:"type,omitempty"`
	OrderIndex int    `json:"orderindex,omitempty"`
	Color      string `json:"color,omitempty"`
}

type BasicFeature struct {
	Enabled bool `json:"enabled,omitempty"`
}

type DueDatesFeature struct {
	Enabled            bool `json:"enabled,omitempty"`
	StartDate          bool `json:"start_date,omitempty"`
	RemapDueDates      bool `json:"remap_due_dates,omitempty"`
	RemapClosedDueDate bool `json:"remap_closed_due_date,omitempty"`
}

type Prioritie struct {
	Id         string `json:"id,omitempty"`
	Priority   string `json:"priority,omitempty"`
	Color      string `json:"color,omitempty"`
	OrderIndex string `json:"orderindex,omitempty"`
}

type PrioritiesFeature struct {
	Enabled    bool        `json:"enabled,omitempty"`
	Priorities []Prioritie `json:"priorities,omitempty"`
}

type CheckUnresolvedFeature struct {
	Enabled    bool     `json:"enabled,omitempty"`
	Subtasks   bool     `json:"subtasks,omitempty"`
	Checklists []string `json:"checklists,omitempty"`
	Comments   []string `json:"comments,omitempty"`
}

type Features struct {
	DueDates          DueDatesFeature        `json:"due_dates,omitempty"`
	TimeTracking      BasicFeature           `json:"time_tracking,omitempty"`
	Tags              BasicFeature           `json:"tags,omitempty"`
	TimeEstimates     BasicFeature           `json:"time_estimates,omitempty"`
	CheckLists        BasicFeature           `json:"check_lists,omitempty"`
	CustomFields      BasicFeature           `json:"custom_fields,omitempty"`
	RemapDependencies BasicFeature           `json:"remap_dependencies,omitempty"`
	DependencyWarning BasicFeature           `json:"dependency_warning,omitempty"`
	Portfolios        BasicFeature           `json:"portfolios,omitempty"`
	Sprints           BasicFeature           `json:"sprints,omitempty"`
	Points            BasicFeature           `json:"points,omitempty"`
	CustomItems       BasicFeature           `json:"custom_items,omitempty"`
	Priorities        PrioritiesFeature      `json:"priorities,omitempty"`
	CheckUnresolved   CheckUnresolvedFeature `json:"check_unresolved,omitempty"`
	Zoom              BasicFeature           `json:"zoom,omitempty"`
	Milestones        BasicFeature           `json:"milestones,omitempty"`
	Emails            BasicFeature           `json:"emails,omitempty"`
}

type Space struct {
	Id                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Private           bool     `json:"private,omitempty"`
	Statuses          []Status `json:"statuses,omitempty"`
	MultipleAssignees bool     `json:"multiple_assignees,omitempty"`
	Features          Features `json:"features,omitempty"`
	Archived          bool     `json:"archived,omitempty"`
}

type Spaces struct {
	Spaces []Space `json:"spaces,omitempty"`
}
