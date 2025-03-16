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

var AllTags = []Tag{
	Generic, HTML, JavaScript, JsExt,
}

var AllTagsAndWorkflows = append(append([]Tag{}, AllTags...), AllWorkflows...)

var AllWorkflows = []Tag{
	Django, GLPI,
	PHPBB,
	Subrion, Umbraco,
	WordPress,
}

func (t Tag) IsWorkflow() bool {
	for _, workflow := range AllWorkflows {
		if t == workflow {
			return true
		}
	}
	return false
}
