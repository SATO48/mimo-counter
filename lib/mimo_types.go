package lib

type APIResponse[T any] struct {
	Links    map[string]string `json:"links"`
	Data     T                 `json:"data"`
	Included []T               `json:"included"`
}

type Layer struct {
	ID         string          `json:"id,omitempty"`
	Attributes LayerAttributes `json:"attributes"`
}

type LayerAttributes struct {
	InputValues struct {
		Text string `json:"tvGroup_Content__Text_TypeMultiline"`
	} `json:"input-values"`
}
