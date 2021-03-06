package client

const (
	LabelSelectorType                  = "labelSelector"
	LabelSelectorFieldMatchExpressions = "matchExpressions"
	LabelSelectorFieldMatchLabels      = "matchLabels"
)

type LabelSelector struct {
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchLabels      map[string]string          `json:"matchLabels,omitempty"`
}
