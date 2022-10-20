package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/madeiramadeirabr/openapi2confluence/internal/dto/config"
)

func TestShouldReturnConfig(t *testing.T) {
	path := "/bar/foo"
	confluencePageId := "10000"
	title := "Teste Pagina"
	spaceKey := "GPT"
	ancestorId := "11000"
	localId := "00000000-0000-0000-0000-0000000000000"
	macroId := "11111111-1111-1111-1111-1111111111111"
	token := "aaaabbbbbccccddddeee"
	email := "email@teste.com"
	host := "madeiramadeira.confluence.com"
	env := "STAGING"

	envAin := "STG"
	titleAin := fmt.Sprintf("[%s] %s", envAin, title)

	os.Setenv("OPENAPI2CONFLUENCE_ENV", env)
	os.Setenv("OPENAPI2CONFLUENCE_CONFLUENCE_HOST", host)
	os.Setenv("OPENAPI2CONFLUENCE_CONFLUENCE_EMAIL", email)
	os.Setenv("OPENAPI2CONFLUENCE_CONFLUENCE_API_KEY", token)
	os.Setenv("", titleAin)

	cfg, err := config.NewConfig(
		path,
		confluencePageId,
		title,
		spaceKey,
		ancestorId,
		localId,
		macroId,
		"",
	)

	if err != nil {
		t.Error(t)
	}

	if cfg.Path != path {
		t.Errorf("the value of '%s' is different that '%s'", cfg.Path, path)
	}

	if cfg.AncestorId != ancestorId {
		t.Errorf("the value of '%s' is different that '%s'", cfg.AncestorId, ancestorId)
	}

	if cfg.SpaceKey != spaceKey {
		t.Errorf("the value of '%s' is different that '%s'", cfg.SpaceKey, spaceKey)
	}

	if cfg.Token != token {
		t.Errorf("the value of '%s' is different that '%s'", cfg.Token, token)
	}

	if cfg.Email != email {
		t.Errorf("the value of '%s' is different that '%s'", cfg.Email, email)
	}

	if cfg.Host != host {
		t.Errorf("the value of '%s' is different that '%s'", cfg.Host, host)
	}

	if cfg.Environment != env {
		t.Errorf("the value of '%s' is different that '%s'", cfg.Environment, env)
	}

	if cfg.Title != titleAin {
		t.Errorf("the value of '%s' is different that '%s'", cfg.Title, titleAin)
	}

	if cfg.MacroId != macroId {
		t.Errorf("the value of '%s' is different that '%s'", cfg.MacroId, macroId)
	}

	if cfg.LocalId != localId {
		t.Errorf("the value of '%s' is different that '%s'", cfg.LocalId, localId)
	}
}
