package dependencies

import (
	"context"
	"errors"

	"github.com/madeiramadeirabr/openapi2confluence/internal/dto/config"
)

type CtxKey string

const KeyCfg CtxKey = "CfgCtx"

func SetCfgCtx(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, KeyCfg, cfg)
}

func GetCfgFromCtx(ctx context.Context) (*config.Config, error) {
	v := ctx.Value(KeyCfg)

	cfg, ok := v.(*config.Config)

	if !ok {
		return nil, errors.New("valor do contexto da cfg não é um ponteiro para dto.Config")
	}

	return cfg, nil
}
