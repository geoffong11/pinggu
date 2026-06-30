package yamlparser

type ConfigPost struct {
	Endpoint string `yaml:"endpoint"`
	Body     string `yaml:"body"`
}

type ConfigGet struct {
	Endpoint string `yaml:"endpoint"`
}

type HttpRequests struct {
	Get  []ConfigGet  `yaml:"GET"`
	Post []ConfigPost `yaml:"POST"`
}

type Config struct {
	Users        int          `yaml:"users"`
	Duration     int          `yaml:"duration"`
	HttpRequests HttpRequests `yaml:"requests"`
}
