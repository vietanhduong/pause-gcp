package gke

type Cluster struct {
	Project   string      `json:"project"`
	Name      string      `json:"name,omitempty"`
	Location  string      `json:"location,omitempty"`
	NodePools []*NodePool `json:"nodePools,omitempty"`
}

type NodePool struct {
	Name             string       `json:"name,omitempty"`
	InstanceGroups   []string     `json:"instanceGroups,omitempty"`
	Autoscaling      *Autoscaling `json:"autoscaling,omitempty"`
	Locations        []string     `json:"locations,omitempty"`
	InitialNodeCount int32        `json:"initialNodeCount,omitempty"`
	CurrentSize      int32        `json:"currentSize,omitempty"`
}

type Autoscaling struct {
	Enabled           bool   `json:"enabled,omitempty"`
	MinNodeCount      int32  `json:"minNodeCount,omitempty"`
	MaxNodeCount      int32  `json:"maxNodeCount,omitempty"`
	Autoprovisioned   bool   `json:"autoprovisioned,omitempty"`
	LocationPolicy    string `json:"locationPolicy,omitempty"`
	TotalMinNodeCount int32  `json:"totalMinNodeCount,omitempty"`
	TotalMaxNodeCount int32  `json:"totalMaxNodeCount,omitempty"`
}
