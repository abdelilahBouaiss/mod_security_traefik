ckage main

import (
	"context"
	"net/http"
	"os/exec"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	RuleFile string `json:"ruleFile,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// ModSecurityPlugin a ModSecurity plugin.
type ModSecurityPlugin struct {
	next     http.Handler
	ruleFile string
	name     string
}

// New creates a new ModSecurity plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &ModSecurityPlugin{
		next:     next,
		ruleFile: config.RuleFile,
		name:     name,
	}, nil
}

func (m *ModSecurityPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Use ModSecurity to inspect the request
	cmd := exec.Command("modsecurity", "--rules", m.ruleFile, "--input", req.URL.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(rw, "Forbidden", http.StatusForbidden)
		return
	}

	if strings.Contains(string(output), "DENY") {
		http.Error(rw, "Forbidden", http.StatusForbidden)
		return
	}

	m.next.ServeHTTP(rw, req)
}

