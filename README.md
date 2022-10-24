# Openapi2confluence

# Descrição 

Github action que faz a atualização das nossas documentações no confluence baseada no openapi


# Contexto de negócio 

Sistemas que contém a documentação no padrão do openapi

# Squad Owner

partnertools

# Get started

## Github - Action 

Configure o seu workflow como o exemplo abaixo:

```
name: Openapi2confluence
on:
  push:
    branches:
      - staging
      - production

jobs:
  SyncApiDocConfluence:
    runs-on: ubuntu-latest
    timeout-minutes: 3
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      
      - name: action sync 
        uses: madeiramadeirabr/action-openapi2confluence@production
        with:
          path: path/to/openapi.yml
          id:  9999999999
          spaceKey: ABC
          ancestorId: 9999999999
          title: Doc cool of my api
          localId: 00000000-0000-0000-0000-000000000000
          macroId: 00000000-0000-0000-0000-000000000000
          env: staging
          confluenceHost: https://youcompany.atlassian.net
          confluenceAuth:  Basic xxxxxxxxxxxxxxxxx
```

## Cli

Para rodar o script voce precisa configurar suas credenciais do confluence em `~/.partnertools/confluence/config.yaml`. Com a seguinte estrutura:

```
confluence_api_key: <TOKEN>
confluence_host: https://madeiramadeira.atlassian.net
```

ou configurar envs com o prefixo `OPENAPI2CONFLUENCE_`

para roda o script segue o exemplo a baixo:

```
openapi2confluence -p <path do openapi> \
    -id <id page for edit> \
    -t "Title of page" \
    -s <key space: ex: GPT> \ 
    -a <id da ancestor page>
    -mid <macro id>
    -lid <local id>
    -env <env name>
```

# Padrão de branchs

* Feature - feature/xxxxx

* Bugfix - bugfix/xxxxx

# Padrão de commmits

Resuma de forma sucinta o que foi adicionado, removido ou refatorado
