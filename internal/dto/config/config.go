package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

const (
	configFilename  = "config"
	configExtension = "yaml"
)

type Config struct {
	Path             string
	ConfluencePageId string
	Title            string
	SpaceKey         string
	AncestorId       string
	Token            string
	Host             string
	Email            string
	Environment      string
	LocalId          string
	MacroId          string
}

func NewConfig(path, confluencePageId, title, spaceKey, ancestorId, localId, macroId, envFlag string) (*Config, error) {
	// carrega config de env ou yaml
	token, host, email, env, err := loadEnvOrConfigFile()
	if err != nil {
		return nil, err
	}

	// valida os dados inputado
	if path == "" {
		return nil, errors.New("informa um path ex: -p /xxx/xxx")
	}

	if confluencePageId == "" {
		return nil, errors.New("informa a pagina do confluence ex: -id xxxxxxxxx")
	}

	if confluencePageId == "" {
		return nil, errors.New("informa o titulo da pagina do confluence ex: -t xxxxxxxxx")
	}

	if spaceKey == "" {
		return nil, errors.New("informa o space da pagina no confluence ex: -s XXX")
	}

	if ancestorId == "" {
		return nil, errors.New("informa a pagina pai da pagina ex: -a xxxxxxxxx")
	}

	if token == "" {
		return nil, errors.New("informa o token (confluence_api_key) no arquivo de config ou env (OPENAPI2CONFUENLCE_CONFLUENCE_API_KEY)")
	}

	if email == "" {
		return nil, errors.New("informa o email (confluence_email) no arquivo de config ou env (OPENAPI2CONFUENLCE_CONFLUENCE_EMAIL)")
	}

	if host == "" {
		return nil, errors.New("informa o host (confluence_host) no arquivo de config ou env (OPENAPI2CONFUENLCE_CONFLUENCE_HOST)")
	}

	if localId == "" {
		return nil, errors.New("informa o local id")
	}

	if macroId == "" {
		return nil, errors.New("informa o macro id")
	}

	if envFlag != "" {
		env = envFlag
	}

	abbreviatedEnv := ""
	switch strings.ToLower(env) {
	case "staging":
		abbreviatedEnv = "STG"
	case "production":
		abbreviatedEnv = "PRD"
	case "sandbox":
		abbreviatedEnv = "SBX"
	default:
		return nil, fmt.Errorf("valor '%s' é inválido para env", env)
	}

	computedTitle := fmt.Sprintf("[%s] %s", abbreviatedEnv, title)

	return &Config{
		Path:             path,
		ConfluencePageId: confluencePageId,
		Title:            computedTitle,
		SpaceKey:         spaceKey,
		AncestorId:       ancestorId,
		Environment:      env,
		Token:            token,
		Email:            email,
		Host:             host,
		MacroId:          macroId,
		LocalId:          localId,
	}, nil
}

func loadEnvOrConfigFile() (string, string, string, string, error) {
	var (
		tokenKey string = "confluence_api_key"
		hostKey  string = "confluence_host"
		emailKey string = "confluence_email"
		envKey   string = "env"
	)

	configPath, err := createIfNotExistsProgramConfig()
	if err != nil {
		return "", "", "", "", err
	}

	viper.SetDefault(tokenKey, "")
	viper.SetDefault(hostKey, "")
	viper.SetDefault(emailKey, "")
	viper.SetDefault(envKey, "")

	viper.SetEnvPrefix("OPENAPI2CONFLUENCE")
	viper.BindEnv(tokenKey)
	viper.BindEnv(emailKey)
	viper.BindEnv(hostKey)
	viper.BindEnv(envKey)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return "", "", "", "", fmt.Errorf(
			"nao pode ler as config configure os ENVs ou crie o arquivo de config %s/config.yaml: %w",
			configPath,
			err,
		)
	}

	return viper.GetString(tokenKey), viper.GetString(hostKey), viper.GetString(emailKey), viper.GetString(envKey), nil
}

func createIfNotExistsProgramConfig() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("nao pode recuperar o home dir: %w", err)
	}

	configPath := path.Join(homeDir, ".partnertools", "confluence")
	configFile := fmt.Sprintf("%s.%s", configFilename, configExtension)

	fullPath := path.Join(configPath, configFile)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("nao pode criar o config dir: %w", err)
		}
	}

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return "", fmt.Errorf("nao pode criar o arquivo de config: %w", err)
	}

	defer f.Close()

	return configPath, nil
}
