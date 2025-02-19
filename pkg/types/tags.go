package types

type Tag string

const (
	GENERIC Tag = "generic"
	HTML    Tag = "html"
)

var AllTags = []Tag{
	GENERIC, HTML,
}
