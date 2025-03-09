package types

type Tag string

const (
	Generic    Tag = "generic"
	HTML       Tag = "html"
	JavaScript Tag = "javascript"

	JsExt Tag = "jsext"

	Django    Tag = "django"
	PHPBB     Tag = "phpbb"
	Subrion   Tag = "subrion"
	Umbraco   Tag = "umbraco"
	WordPress Tag = "wordpress"
)

var AllTags = []Tag{
	Generic, HTML, JavaScript,
	JsExt,
	Django, PHPBB, Subrion, Umbraco, WordPress,
}

var AllWorkflows = []Tag{
	Django, PHPBB, Subrion, Umbraco, WordPress,
}
