package scripts

type Meta struct {
	Language   string   `yaml:"language"`
	Action     string   `yaml:"action"`
	Entrypoint string   `yaml:"entrypoint"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args"`
}
