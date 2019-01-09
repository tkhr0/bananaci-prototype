package config

type (
	// root
	Config struct {
		// .services
		Services []Service `yaml:"services"`
	}

	// .services.n
	Service struct {
		Name  string `yaml:"name"`
		Build `yaml:"build"`
	}

	// .services.n.build
	Build struct {
		Context    string `yaml:"context"`
		Dockerfile string `yaml:"dockerfile"`
	}
)
