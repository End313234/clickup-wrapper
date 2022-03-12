package clickup

type HistoryItemTaskCreated struct {
	Id       string `json:"id,omitempty"`
	Type     int    `json:"type,omitempty"`
	Date     string `json:"date,omitempty"`
	Field    string `json:"field,omitempty"`
	ParentId string `json:"parent_id,omitempty"`
	Data     struct {
		StatusType string `json:"status_type,omitempty"`
	} `json:"data,omitempty"`
	Source string `json:"source,omitempty"`
	User   User   `json:"user,omitempty"`
	Before Status `json:"before,omitempty"`
	After  Status `json:"after,omitempty"`
}

// Payload for taskCreated event
type TaskCreatedPayload struct {
	TaskId       string                   `json:"task_id,omitempty"`
	WebhookId    string                   `json:"webhook_id,omitempty"`
	Event        string                   `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCreated `json:"history_items,omitempty"`
}

type TypeConfigOption struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Value      string `json:"value,omitempty"`
	Type       string `json:"type,omitempty"`
	Color      string `json:"color,omitempty"`
	OrderIndex int    `json:"orderindex,omitempty"`
}

type TypeConfig struct {
	Default            int                `json:"default,omitempty"`
	Placeholder        string             `json:"placeholder,omitempty"`
	NewDropDown        bool               `json:"new_drop_down,omitempty"`
	Options            []TypeConfigOption `json:"options,omitempty"`
	ValuesSet          []string           `json:"values_set,omitempty"`
	UserId             string             `json:"userid,omitempty"`
	CreatedAt          string             `json:"date_created,omitempty"`
	HideFromGuests     bool               `json:"hide_from_guests,omitempty"`
	TeamId             string             `json:"team_id,omitempty"`
	Deleted            bool               `json:"deleted,omitempty"`
	DeletedBy          User               `json:"deleted_by,omitempty"`
	Pinned             bool               `json:"pinned,omitempty"`
	Required           bool               `json:"required,omitempty"`
	RequiredOnSubtasks bool               `json:"required_on_subtasks,omitempty"`
	LinkedSubcategory  string             `json:"linked_subcategory"`
}

type CustomField struct {
	Id         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Type       string     `json:"type,omitempty"`
	TypeConfig TypeConfig `json:"type_config,omitempty"`
}

type HistoryItemTaskUpdated struct {
	Id       string `json:"id,omitempty"`
	Type     int    `json:"type,omitempty"`
	Date     string `json:"date,omitempty"`
	Field    string `json:"field,omitempty"`
	ParentId string `json:"parent_id,omitempty"`
	Data     struct {
		StatusType string `json:"status_type,omitempty"`
	} `json:"data,omitempty"`
	Source      string      `json:"source,omitempty"`
	User        User        `json:"user,omitempty"`
	Before      string      `json:"before,omitempty"`
	After       string      `json:"after,omitempty"`
	CustomField CustomField `json:"custom_field,omitempty"`
}

// Payload for taskUpdated event
type TaskUpdatedPayload struct {
	TaskId       string                   `json:"task_id,omitempty"`
	WebhookId    string                   `json:"webhook_id,omitempty"`
	Event        string                   `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCreated `json:"history_items,omitempty"`
}

// Payload for taskDeleted event
type TaskDeletedPayload struct {
	TaskId    string `json:"task_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for taskPriorityUpdated event
type TaskPriorityUpdatedPayload struct {
	TaskId       string                   `json:"task_id,omitempty"`
	WebhookId    string                   `json:"webhook_id,omitempty"`
	Event        string                   `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCreated `json:"history_items,omitempty"`
}

// Payload for taskStatusUpdated event
type TaskStatusUpdatedPayload struct {
	TaskId       string                   `json:"task_id,omitempty"`
	WebhookId    string                   `json:"webhook_id,omitempty"`
	Event        string                   `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCreated `json:"history_items,omitempty"`
}

// Payload for taskAssigneeUpdated event
type TaskAssigneeUpdatedPayload struct {
	TaskId       string                   `json:"task_id,omitempty"`
	WebhookId    string                   `json:"webhook_id,omitempty"`
	Event        string                   `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCreated `json:"history_items,omitempty"`
}

type HistoryItemDataTaskDueDateUpdated struct {
	StatusType     string `json:"status_type,omitempty"`
	DueDateTime    string `json:"due_date_time,omitempty"`
	OldDueDateTime string `json:"old_due_date_time,omitempty"`
}

type HistoryItemTaskDueDateUpdated struct {
	Id       string                            `json:"id,omitempty"`
	Type     int                               `json:"type,omitempty"`
	Date     string                            `json:"date,omitempty"`
	Field    string                            `json:"field,omitempty"`
	ParentId string                            `json:"parent_id,omitempty"`
	Data     HistoryItemDataTaskDueDateUpdated `json:"data,omitempty"`
	Source   string                            `json:"source,omitempty"`
	User     User                              `json:"user,omitempty"`
	Before   string                            `json:"before,omitempty"`
	After    string                            `json:"after,omitempty"`
}

// Payload for taskDueDateUpdated event
type TaskDueDateUpdatedPayload struct {
	TaskId       string                          `json:"task_id,omitempty"`
	WebhookId    string                          `json:"webhook_id,omitempty"`
	Event        string                          `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskDueDateUpdated `json:"history_items,omitempty"`
}

type HistoryItemAfterAndBeforeTaskTagUpdated struct {
	Name    string `json:"name,omitempty"`
	TagFg   string `json:"tag_fg,omitempty"`
	TagBg   string `json:"tag_bg,omitempty"`
	Creator int    `json:"creator,omitempty"`
}

type HistoryItemTaskTagUpdated struct {
	Id       string                                    `json:"id,omitempty"`
	Type     int                                       `json:"type,omitempty"`
	Date     string                                    `json:"date,omitempty"`
	Field    string                                    `json:"field,omitempty"`
	ParentId string                                    `json:"parent_id,omitempty"`
	Data     HistoryItemDataTaskDueDateUpdated         `json:"data,omitempty"`
	Source   string                                    `json:"source,omitempty"`
	User     User                                      `json:"user,omitempty"`
	Before   []HistoryItemAfterAndBeforeTaskTagUpdated `json:"before,omitempty"`
	After    []HistoryItemAfterAndBeforeTaskTagUpdated `json:"after,omitempty"`
}

// Payload for taskTagUpdated event
type TaskTagUpdatedPayload struct {
	TaskId       string                                    `json:"task_id,omitempty"`
	WebhookId    string                                    `json:"webhook_id,omitempty"`
	Event        string                                    `json:"event,omitempty"`
	HistoryItems []HistoryItemAfterAndBeforeTaskTagUpdated `json:"history_items,omitempty"`
}

type HistoryItemDataTaskMoved struct {
	MuteNotifications bool `json:"mute_notifications,omitempty"`
}

type HistoryItemAfterAndBeforeTaskMoved struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Category struct {
		Id     string `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		Hidden bool   `json:"hidden,omitempty"`
	}
	Project struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
}

type HistoryItemTaskMoved struct {
	Id       string                               `json:"id,omitempty"`
	Type     int                                  `json:"type,omitempty"`
	Date     string                               `json:"date,omitempty"`
	Field    string                               `json:"field,omitempty"`
	ParentId string                               `json:"parent_id,omitempty"`
	Data     HistoryItemDataTaskDueDateUpdated    `json:"data,omitempty"`
	Source   string                               `json:"source,omitempty"`
	User     User                                 `json:"user,omitempty"`
	Before   []HistoryItemAfterAndBeforeTaskMoved `json:"before,omitempty"`
	After    []HistoryItemAfterAndBeforeTaskMoved `json:"after,omitempty"`
}

// Payload for taskMoved event
type TaskMovedPayload struct {
	TaskId       string                 `json:"task_id,omitempty"`
	WebhookId    string                 `json:"webhook_id,omitempty"`
	Event        string                 `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskMoved `json:"history_items,omitempty"`
}

type HistoryItemCommentTaskCommentPosted struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Parent  string `json:"parent,omitempty"`
	Type    int    `json:"type,omitempty"`
	Comment struct {
		Text       string `json:"text,omitempty"`
		Attributes struct {
			BlockId string `json:"block-id,omitempty"`
		} `json:"attributes,omitempty"`
	}
	TextContent             string   `json:"text_content,omitempty"`
	X                       string   `json:"x,omitempty"`
	Y                       string   `json:"y,omitempty"`
	ImageX                  string   `json:"image_x,omitempty"`
	ImageY                  string   `json:"image_y,omitempty"`
	Page                    int      `json:"page,omitempty"`
	CommentNumber           int      `json:"comment_number"`
	PageId                  int      `json:"page_id,omitempty"`
	PageName                string   `json:"page_name,omitempty"`
	ViewId                  string   `json:"view_id,omitempty"`
	ViewName                string   `json:"view_name,omitempty"`
	Team                    string   `json:"team,omitempty"`
	User                    User     `json:"user,omitempty"`
	NewThreadCount          int      `json:"new_thread_count,omitempty"`
	NewMentionedThreadCount int      `json:"new_mentioned_thread_count,omitempty"`
	EmailAttachments        []string `json:"email_attachments,omitempty"`
	ThreadedUsers           []User   `json:"threaded_users,omitempty"`
	ThreadedReplies         int      `json:"threaded_replies,omitempty"`
	ThreadedAssignees       int      `json:"threaded_assignees,omitempty"`
	ThreadedUnresolvedCount int      `json:"threaded_unresolved_count,omitempty"`
	ThreadFollowers         []User   `json:"threaded_followers,omitempty"`
	GroupThreadFollowers    []User   `json:"group_thread_followers,omitempty"`
	Reactions               []string `json:"reactions,omitempty"`
	Emails                  []string `json:"emails,omitempty"`
}

type HistoryItemTaskCommentPosted struct {
	Id       string                              `json:"id,omitempty"`
	Type     int                                 `json:"type,omitempty"`
	Date     string                              `json:"date,omitempty"`
	Field    string                              `json:"field,omitempty"`
	ParentId string                              `json:"parent_id,omitempty"`
	Data     HistoryItemDataTaskDueDateUpdated   `json:"data,omitempty"`
	Source   string                              `json:"source,omitempty"`
	User     User                                `json:"user,omitempty"`
	Before   string                              `json:"before,omitempty"`
	After    string                              `json:"after,omitempty"`
	Comment  HistoryItemCommentTaskCommentPosted `json:"comment,omitempty"`
}

// Payload for taskCommentPosted event
type TaskCommentPostedPayload struct {
	TaskId       string                         `json:"task_id,omitempty"`
	WebhookId    string                         `json:"webhook_id,omitempty"`
	Event        string                         `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCommentPosted `json:"history_items,omitempty"`
}

// Payload for taskCommentUpdated event
type TaskCommentUpdatedPayload struct {
	TaskId       string                         `json:"task_id,omitempty"`
	WebhookId    string                         `json:"webhook_id,omitempty"`
	Event        string                         `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskCommentPosted `json:"history_items,omitempty"`
}

type HistoryItemDataTaskTimeEstimateUpdated struct {
	TimeEstimateString    string `json:"time_estimate_string,omitempty"`
	OldTimeEstimateString string `json:"old_time_estimate_string,omitempty"`
	RolledUpTimeEstimate  uint   `json:"rolled_up_time_estimate,omitempty"`
	TimeEstimate          uint   `json:"time_estimate,omitempty"`
	TimeEstimatesByUser   []struct {
		UserId                 int    `json:"userid,omitempty"`
		UserTimeEstimate       string `json:"user_time_estimate,omitempty"`
		UserRollupTimeEstimate string `json:"user_rollup_time_estimate,omitempty"`
	} `json:"time_estimates_by_user,omitempty"`
}

type HistoryItemTaskTimeEstimateUpdated struct {
	Id       string                                 `json:"id,omitempty"`
	Type     int                                    `json:"type,omitempty"`
	Date     string                                 `json:"date,omitempty"`
	Field    string                                 `json:"field,omitempty"`
	ParentId string                                 `json:"parent_id,omitempty"`
	Data     HistoryItemDataTaskTimeEstimateUpdated `json:"data,omitempty"`
	Source   string                                 `json:"source,omitempty"`
	User     User                                   `json:"user,omitempty"`
	Before   string                                 `json:"before,omitempty"`
	After    string                                 `json:"after,omitempty"`
}

// Payload for taskTimeEstimateUpdated event
type TaskTimeEstimateUpdated struct {
	TaskId       string                               `json:"task_id,omitempty"`
	WebhookId    string                               `json:"webhook_id,omitempty"`
	Event        string                               `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskTimeEstimateUpdated `json:"history_items,omitempty"`
}

type HistoryItemAfterAndBeforeTaskTimeTrackedUpdated struct {
	Id        string `json:"id,omitempty"`
	Start     string `json:"start,omitempty"`
	End       string `json:"end,omitempty"`
	Time      string `json:"time,omitempty"`
	Source    string `json:"source,omitempty"`
	DateAdded string `json:"date_added,omitempty"`
}

type HistoryItemTaskTimeTrackedUpdated struct {
	Id       string `json:"id,omitempty"`
	Type     int    `json:"type,omitempty"`
	Date     string `json:"date,omitempty"`
	Field    string `json:"field,omitempty"`
	ParentId string `json:"parent_id,omitempty"`
	Data     struct {
		TotalTime  string `json:"total_time,omitempty"`
		RollupTime string `json:"rollup_time,omitempty"`
	} `json:"data,omitempty"`
	Source string `json:"source,omitempty"`
	User   User   `json:"user,omitempty"`
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

// Payload for taskTimeTrackedUpdated event
type TaskTimeTrackedUpdatedPayload struct {
	TaskId    string `json:"task_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Data      struct {
		Description string `json:"description,omitempty"`
		IntervalId  string `json:"interval_id,omitempty"`
	} `json:"data,omitempty"`
	Event        string                              `json:"event,omitempty"`
	HistoryItems []HistoryItemTaskTimeTrackedUpdated `json:"history_items,omitempty"`
}

// Payload for listCreated event
type ListCreatedPayload struct {
	ListId    string `json:"list_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

type HistoryItemListUpdated struct {
	Id       string            `json:"id,omitempty"`
	Type     int               `json:"type,omitempty"`
	Date     string            `json:"date,omitempty"`
	Field    string            `json:"field,omitempty"`
	ParentId string            `json:"parent_id,omitempty"`
	Source   map[string]string `json:"source,omitempty"`
	User     User              `json:"user,omitempty"`
	Before   string            `json:"before,omitempty"`
	After    string            `json:"after,omitempty"`
}

// Payload for listUpdated event
type ListUpdatedPayload struct {
	ListId      string                   `json:"list_id,omitempty"`
	WebhookId   string                   `json:"webhook_id,omitempty"`
	Event       string                   `json:"event,omitempty"`
	HistoryItem []HistoryItemListUpdated `json:"history_item,omitempty"`
}

// Payload for listDeleteds event
type ListDeletedPayload struct {
	ListId    string `json:"list_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for folderCreated event
type FolderCreatedPayload struct {
	FolderId  string `json:"folder_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for folderUpdated event
type FolderUpdatedPayload struct {
	FolderId  string `json:"folder_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for folderDeleted event
type FolderDeletedPayload struct {
	FolderId  string `json:"folder_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for spaceCreated event
type SpaceCreatedPayload struct {
	SpaceId   string `json:"space_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for spaceCreated event
type SpaceUpdatedPayload struct {
	SpaceId   string `json:"space_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for spaceDeleted event
type SpaceDeletedPayload struct {
	SpaceId   string `json:"space_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for goalCreated event
type GoalCreatedPayload struct {
	GoalId    string `json:"goal_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for goalUpdated event
type GoalUpdatedPayload struct {
	GoalId    string `json:"goal_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for goalDeleted event
type GoalDeletedPayload struct {
	GoalId    string `json:"goal_id,omitempty"`
	WebhookId string `json:"webhook_id,omitempty"`
	Event     string `json:"event,omitempty"`
}

// Payload for keyResultCreated event
type KeyResultCreatedPayload struct {
	GoalId      string `json:"goal_id,omitempty"`
	WebhookId   string `json:"webhook_id,omitempty"`
	KeyResultId string `json:"key_result_id,omitempty"`
	Event       string `json:"event,omitempty"`
}

// Payload for keyResultUpdated event
type KeyResultUpdatedPayload struct {
	GoalId      string `json:"goal_id,omitempty"`
	WebhookId   string `json:"webhook_id,omitempty"`
	KeyResultId string `json:"key_result_id,omitempty"`
	Event       string `json:"event,omitempty"`
}

// Payload for keyResultDeleted event
type KeyResultDeletedPayload struct {
	GoalId      string `json:"goal_id,omitempty"`
	WebhookId   string `json:"webhook_id,omitempty"`
	KeyResultId string `json:"key_result_id,omitempty"`
	Event       string `json:"event,omitempty"`
}
