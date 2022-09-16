package confluence

import (
	"github.com/madeiramadeirabr/openapi2confluence/internal/dto/config"
	goconfluence "github.com/virtomize/confluence-go-api"
)

func CfgToQueryFind(cfg *config.Config) *goconfluence.Content {
	return &goconfluence.Content{
		ID:   cfg.ConfluencePageId,
		Type: "page", // TODO no futuro deixar mais generico
		Ancestors: []goconfluence.Ancestor{
			{
				ID: cfg.AncestorId,
			},
		},
		Space: goconfluence.Space{
			Key: cfg.SpaceKey,
		},
		Title: cfg.Title,
	}
}

func CfgToContentUpdate(cfg *config.Config, v int, val string) *goconfluence.Content {
	return &goconfluence.Content{
		ID:   cfg.ConfluencePageId,
		Type: "page", // TODO no futuro deixar mais generico
		Version: &goconfluence.Version{
			Number: v,
		},
		Title: cfg.Title,
		Space: goconfluence.Space{
			Key: cfg.SpaceKey,
		},
		Ancestors: []goconfluence.Ancestor{
			{
				ID: cfg.AncestorId,
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          val,
				Representation: "storage",
			},
		},
	}
}

func CfgToContentCreate(cfg *config.Config, val string) *goconfluence.Content {
	return &goconfluence.Content{
		Title: cfg.Title,
		Type:  "page", // TODO no futuro deixar mais generico
		Version: &goconfluence.Version{
			Number: 1,
		},
		Space: goconfluence.Space{
			Key: cfg.SpaceKey,
		},
		Ancestors: []goconfluence.Ancestor{
			{
				ID: cfg.AncestorId,
			},
		},
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          val,
				Representation: "storage",
			},
		},
	}
}
