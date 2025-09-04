package settings

import (
	"maps"
	"net/http"

	"git.gammaspectra.live/git/go-away/utils"
)

type Settings struct {
	Bind Bind `yaml:"bind"`

	Backends map[string]Backend `yaml:"backends"`

	BindDebug   string `yaml:"bind-debug"`
	BindMetrics string `yaml:"bind-metrics"`

	Strings utils.Strings `yaml:"strings"`

	// Links to add to challenge/error pages like privacy/impressum.
	Links []Link `yaml:"links"`

	ChallengeHttpCode int `yaml:"challenge-http-code"`

	ChallengeTemplate string `yaml:"challenge-template"`

	// ChallengeTemplateOverrides Key/Value overrides for the current chosen template
	ChallengeTemplateOverrides map[string]string `yaml:"challenge-template-overrides"`
}

type Link struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

var DefaultSettings = Settings{
	Strings:           DefaultStrings,
	ChallengeHttpCode: http.StatusTeapot,
	ChallengeTemplate: "anubis",
	ChallengeTemplateOverrides: func() map[string]string {
		m := make(map[string]string)
		maps.Copy(m, map[string]string{
			"Theme": "",
			"Logo":  "",
		})
		return m
	}(),

	Bind: Bind{
		Address:         ":8080",
		Network:         "tcp",
		SocketMode:      "0770",
		Proxy:           false,
		TLSAcmeAutoCert: "",
	},
	Backends: make(map[string]Backend),
}
