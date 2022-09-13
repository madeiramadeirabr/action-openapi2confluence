package dependencies

import (
	"context"

	"github.com/madeiramadeirabr/openapi2confluence/internal/dto/config"
)

type CtxKey string

const KeyCfg CtxKey = "CfgCtx"

func SetCfgCtx(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, KeyCfg, cfg)
}
