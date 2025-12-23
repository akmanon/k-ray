package models

type Findings struct {
	Severity  string `json:"severity"`
	Namespace string `json:"namespace"`
	Resource  string `json:"resource"`
	Reason    string `json:"reason"`
	Message   string `json:"message,omitempty"`
	Restarts  int32  `json:"restarts,omitempty"`
}
