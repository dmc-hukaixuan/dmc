package ticket

// 在 tag 中添加 omitempty 忽略空值
// 注意这里 hobby,omitempty 合起来是 json tag 值，中间用英文逗号分隔
type FeildData struct {
	Name                 string            `json:"name"`
	Default              string            `json:"default"`
	FieldType            string            `json:"type"`
	Label                string            `json:"label"`
	Placeholder          string            `json:"placeholder,omitempty"`
	Display              int               `json:"display"`
	Impacts              []string          `json:"impacts,omitempty"`
	DependsOn            []string          `json:"dependsOn,omitempty"`
	PromptCode           int               `json:"promptCode,omitempty"`
	PromptMessage        string            `json:"PromptMessage,omitempty"`
	AutoComplete         bool              `json:"autoComplete,omitempty"`
	Options              map[string]string `json:"options,omitempty"`
	OptionsType          string            `json:"optionsType,omitempty"`
	OptionsValueComments map[string]string `json:"optionsValueComments,omitempty"`
	HintMessage          string            `json:"hint,omitempty"`
	HintType             int               `json:"hintType,omitempty"`
	Width                int               `json:"width,omitempty"`
	RegexError           string            `json:"regexError,omitempty"`
	Regex                string            `json:"regex,omitempty"`
}
