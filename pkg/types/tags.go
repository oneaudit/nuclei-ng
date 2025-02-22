package types

type Tag string

const (
	Generic    Tag = "generic"
	HTML       Tag = "html"
	JavaScript Tag = "javascript"
	JsExt      Tag = "jsext"
	WordPress  Tag = "wordpress"
)

var AllTags = []Tag{
	Generic, HTML, JavaScript, JsExt, WordPress,
}
