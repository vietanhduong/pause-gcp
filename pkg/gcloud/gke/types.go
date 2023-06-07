package gke

const (
	// Operation Type

	SetNodePoolSize = "SET_NODE_POOL_SIZE"

	// Operation Status

	Done    = "DONE"
	Running = "RUNNING"
)

type Operation struct {
	Name          string `json:"name,omitempty"`
	OperationType string `json:"operation_type,omitempty"`
	StartTime     string `json:"start_time,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	TargetLink    string `json:"target_link,omitempty"`
	Zone          string `json:"zone,omitempty"`
	Status        string `json:"status,omitempty"`
}

type OperationFilter struct {
	Status        string
	OperationType string
}
