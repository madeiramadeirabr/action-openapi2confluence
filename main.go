package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/madeiramadeirabr/openapi2confluence/internal/confluence"
	"github.com/madeiramadeirabr/openapi2confluence/internal/dependencies"
	"github.com/madeiramadeirabr/openapi2confluence/internal/dto/config"
	"github.com/madeiramadeirabr/openapi2confluence/internal/reader"
	"github.com/madeiramadeirabr/openapi2confluence/internal/template"
	goconfluence "github.com/virtomize/confluence-go-api"
)

var (
	cfg *config.Config
)

func main() {
	ctx := getCtx()

	cli := confluence.NewClient(http.DefaultClient, cfg.Host, cfg.Email, cfg.Token)

	spec, err := reader.LoadOpenApiSpec(cfg.Path)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[🔍] Procurando pagina... %s\n", cfg.Title)
	search, err := cli.FindContent(ctx, confluence.CfgToQueryFind(cfg), 0)
	if err != nil {
		panic(fmt.Errorf("cant search current page: %w", err))
	}

	version := 1
	has := len(search.Results) > 0
	content := &goconfluence.Content{}
	if has {
		// recupera a versão atual
		fmt.Println("[🤔] versão atual: ", search.Results[0].Version.Number)
		version = search.Results[0].Version.Number + 1
		fmt.Printf("[👷] Atualizando pagina: %s\n", cfg.Title)
		fmt.Printf("[✅] Versão: %d\n", version)
		contentPayload := confluence.CfgToContentUpdate(
			cfg,
			version,
			template.Render(spec, cfg.MacroId, cfg.LocalId),
		)
		// atualizar
		content, err = cli.UpdateContent(ctx, contentPayload)
	} else {
		// criar
		// fmt.Printf("[👷] Criando pagina: %s\n", cfg.Title)
		// contentPayload := confluence.CfgToContentCreate(cfg, spec)
		// content, err = cli.CreateContent(ctx, contentPayload)
		// TODO resolver bug de criar
		fmt.Printf("[❌] Primeiro crie a pagina (%s) no confluence com o puglin do swagger\n", cfg.Title)
		err = errors.New("pagina do confluence nao criada")
	}

	if err != nil {
		panic(err)
	}

	if has {
		fmt.Printf("[👌] Atualizado com sucesso\n")
	} else {
		fmt.Printf("[👌] Criado com sucesso\n")
	}

	fmt.Println("[🌐] Link: " + content.Links.Base + content.Links.TinyUI)
}

func init() {
	var (
		path             string
		localId          string
		macroId          string
		confluencePageId string
		spaceKey         string
		ancestorId       string
		title            string
		env              string
		err              error
	)

	// valores das flags
	flag.StringVar(&path, "p", "", "caminho do arquivo openapi")
	flag.StringVar(&confluencePageId, "id", "", "id da pagina no confluence")
	flag.StringVar(&spaceKey, "s", "", "space key da pagina")
	flag.StringVar(&ancestorId, "a", "", "id da pagina pai")
	flag.StringVar(&title, "t", "", "titulo da pagina")
	flag.StringVar(&localId, "lid", "", "locale id")
	flag.StringVar(&macroId, "mid", "", "macro id")
	flag.StringVar(&env, "env", "", "envorriment")
	flag.Parse()

	cfg, err = config.NewConfig(
		path,
		confluencePageId,
		title,
		spaceKey,
		ancestorId,
		localId,
		macroId,
		env,
	)
	if err != nil {
		panic(err)
	}
}

func getCtx() context.Context {
	ctx := context.Background()

	ctx = dependencies.SetCfgCtx(context.Background(), cfg)

	return ctx
}
