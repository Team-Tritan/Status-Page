package config

type Config struct {
	Port     string
	Sentry   string
	Services ServicesStruct
}

type ServicesStruct struct {
	Services []ServiceConfig
}

type ServiceConfig struct {
	Title       string
	Hostname    string
	Port        string
	Description string
}

func LoadConfig() Config {
	return Config{
		Port:   ":3000",
		Sentry: "https://f9d0e8db7b6dd42f126d144ce3815521@sentry.tritan.dev/14",
		Services: ServicesStruct{
			[]ServiceConfig{
				{
					Title:       "Core Router",
					Hostname:    "lightning.as393577.net",
					Port:        "179",
					Description: "The core router for our ASN that has Arelion as an upstream.",
				},
				{
					Title:       "Spacecoast Router",
					Hostname:    "charter.edge.as393577.net",
					Port:        "8006",
					Description: "A remote router in Florida for our ASN that is has Charter as a upstream.",
				},
				{
					Title:       "IXP Router 1",
					Hostname:    "kan.lightning.as393577.net",
					Port:        "179",
					Description: "Our core internet exchange router, located in Kansas USA.",
				},
				{
					Title:       "IXP Router 2: Amsterdam",
					Hostname:    "ams.lightning.as393577.net",
					Port:        "179",
					Description: "Our EU internet exchange router, located in Amsterdam NL.",
				},
			},
		},
	}
}
