package types

type Tag string

const (
	Generic    Tag = "generic"
	HTML       Tag = "html"
	JavaScript Tag = "javascript"
	JsExt      Tag = "jsext"
)

var AllTags = []Tag{
	Generic, HTML, JavaScript, JsExt,
}
