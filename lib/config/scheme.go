package config

type (
	// root
	Config struct {
		// .services
		Images   []BuildImage `yaml:"images"`
		Services Services     `yaml:"services"`
	}

	// .services.n
	BuildImage struct {
		Name  string `yaml:"name"`
		Build `yaml:"build"`
	}

	// .services.n.build
	Build struct {
		Context    string `yaml:"context"`
		Dockerfile string `yaml:"dockerfile"`
	}

	Services struct {
		Nginx
		Rails
	}

	Nginx struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	}

	Rails struct {
		Name   string `yaml:"name"`
		Port   int    `yaml:"port"`
		DbHost int    `yaml:"dbHost"`
	}
)
