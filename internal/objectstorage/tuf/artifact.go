package tuf

// Hashes represents the Hashes structure
type Hashes struct {
	Sha256 string `json:"sha256"`
}

// TargetsInfo represents the TargetsInfo structure
type TargetsInfo struct {
	Length int            `json:"length"`
	Hashes Hashes         `json:"hashes"`
	Custom map[string]any `json:"custom,omitempty"`
}

// Target represents the Target structure
type Target struct {
	Path string      `json:"path"`
	Info TargetsInfo `json:"info"`
}

// ArtifactPayload represents the payload structure
type ArtifactPayload struct {
	Targets           []Target `json:"targets"`
	AddTaskIDToCustom bool     `json:"add_task_id_to_custom"`
	PublishTargets    bool     `json:"publish_targets"`
}
