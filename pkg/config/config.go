package config

import "text/template"

// AppConfig holds the app configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
