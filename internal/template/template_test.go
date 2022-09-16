package template_test

import (
	"testing"

	"github.com/madeiramadeirabr/openapi2confluence/internal/template"
)

func TestShouldRenderTemplate(t *testing.T) {
	aim := `<ac:structured-macro ac:name="swagger-integration" ac:schema-version="1" data-layout="default" ac:local-id="00000000-0000-0000-0000-00000000000" ac:macro-id="11111111-1111-1111-1111-11111111111"><ac:plain-text-body><![CDATA[teste]]></ac:plain-text-body></ac:structured-macro>`

	get := template.Render(
		"teste",
		"11111111-1111-1111-1111-11111111111",
		"00000000-0000-0000-0000-00000000000",
	)

	if aim != get {
		t.Errorf("the result '%s' is not equals to '%s'", get, aim)
	}
}
