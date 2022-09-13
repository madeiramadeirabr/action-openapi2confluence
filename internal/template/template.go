package template

import "fmt"

var (
	template string = `<ac:structured-macro ac:name="swagger-integration" ac:schema-version="1" data-layout="default" ac:local-id="%s" ac:macro-id="%s"><ac:plain-text-body><![CDATA[%s]]></ac:plain-text-body></ac:structured-macro>`
)

func Render(specContent, macroId, localId string) string {
	return fmt.Sprintf(template, localId, macroId, specContent)
}
