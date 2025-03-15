package types

type Tag string

const (
	Generic    Tag = "generic"
	HTML       Tag = "html"
	JavaScript Tag = "javascript"

	JsExt Tag = "jsext"

	Django    Tag = "django"
	GLPI      Tag = "glpi"
	PHPBB     Tag = "phpbb"
	Subrion   Tag = "subrion"
	Umbraco   Tag = "umbraco"
	WordPress Tag = "wordpress"
)

var AllTags = append([]Tag{
	Generic, HTML, JavaScript, JsExt,
}, AllWorkflows...)

var AllWorkflows = []Tag{
	Django, GLPI,
	PHPBB,
	Subrion, Umbraco,
	WordPress,
}
