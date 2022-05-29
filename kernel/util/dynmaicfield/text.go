package dynmaicfield

type Text struct{}

func (c *Text) SearchFieldRender() {

}

func (c *Text) ValueSet() {

}

func (c *Text) ValueGet() {

}

func (c *Text) EditFieldRender() {

}

func (c *Text) DisplayValueRender() {

}

// ticket template config render
// if he his angular
func (c *Text) TemplateRender(id int) map[string]string {

	return map[string]string{
		"name":      "",
		"label":     "",
		"fieldType": "",
		"default":   "",
		"display":   "",
		"regex":     "",
		"regexHint": "",
		"hint":      "",
		"hintType":  "",
	}
}

// terimal user conifg
