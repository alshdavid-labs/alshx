package scripts

type Meta struct {
	Version    int      `yaml:"version"`
	Language   string   `yaml:"language"`
	Action     string   `yaml:"action"`
	Entrypoint string   `yaml:"entrypoint"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args"`
}
