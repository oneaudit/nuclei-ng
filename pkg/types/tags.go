package types

type Tag string

const (
	GENERIC Tag = "generic"
	HTML    Tag = "html"
	JsExt   Tag = "jsext"
)

var AllTags = []Tag{
	GENERIC, HTML, JsExt,
}
