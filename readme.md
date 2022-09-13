# Openapi2confluence

# Descrição 

Github action que faz a atualização das nossas documentações no confluence baseada no openapi


# Contexto de negócio 

Sistemas que contém a documentação no padrão do openapi

# Squad Owner

partnertools

# Get started

Para rodar o script voce precisa configurar suas credenciais do confluence em `~/.partnertools/confluence/config.yaml`. Com a seguinte estrutura:

```
confluence_api_key: <TOKEN>
confluence_email: <SEU EMAIL>
confluence_host: https://madeiramadeira.atlassian.net
```

ou configurar envs com o prefixo `OPENAPI2CONFLUENCE_`

para roda o script segue o exemplo a baixo:

```
openapi2confluence -p <path do openapi> \
    -id <id da pagina se houver> \
    -t "Titulo da pagina" \
    -s <key space: ex: GPT> \ 
    -a <id da pagina pai>
    -mid <id da macro> // podem ser conseguidos quando a pagina for criada
    -lid <local id> // podem ser conseguidos quando a pagina for criada
```

# Exit codes

* 0 - Sucesso
* > 0 - Error

# Padrão de branchs

* Feature - feature/xxxxx

* Bugfix - bugfix/xxxxx

# Padrão de commmits

Resuma de forma sucinta o que foi adicionado, removido ou refatorado
