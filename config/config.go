package config

type Config struct {
	Port     string
	Sentry   string
	MongoURI string
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
		Port:     ":3000",
		Sentry:   "https://f9d0e8db7b6dd42f126d144ce3815521@sentry.tritan.dev/14",
		MongoURI: "mongodb://data.myinfra.lol:27017",

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
					Title:       "IXP Router 1: United States",
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
				{
					Title:       "Development Router: AS57308",
					Hostname:    "66.103.134.6",
					Port:        "179",
					Description: "A core router for our downstreamed development ASN.",
				},
				{
					Title:       "Hypervisor 01",
					Hostname:    "colo-fl.myinfra.lol",
					Port:        "8006",
					Description: "A 4u hypervisor used for AI and ML workloads.",
				},
				{
					Title:       "Hypervisor 02",
					Hostname:    "spacecoast.myinfra.lol",
					Port:        "8006",
					Description: "A 1u hypervisor used for network testing.",
				},
				{
					Title:       "Hypervisor 03",
					Hostname:    "colo-fl03.myinfra.lol",
					Port:        "9090",
					Description: "A 1u hypervisor used for web hosting.",
				},
				{
					Title:       "Hypervisor 04",
					Hostname:    "colo-fl04.myinfra.lol",
					Port:        "9090",
					Description: "A 1u hypervisor used for customer hosting.",
				},
				{
					Title:       "Hypervisor 05",
					Hostname:    "colo-fl05.myinfra.lol",
					Port:        "9090",
					Description: "A 1u hypervisor.",
				},
				{
					Title:       "Hypervisor 06",
					Hostname:    "colo-fl06.myinfra.lol",
					Port:        "8006",
					Description: "A 2u hypervisor with an Epyc CPU used for core services.",
				},
				{
					Title:       "Grafana",
					Hostname:    "grafana.myinfra.lol",
					Port:        "3000",
					Description: "A Grafana instance used for monitoring.",
				},
				{
					Title:       "3cx",
					Hostname:    "3cx.myinfra.lol",
					Port:        "443",
					Description: "A 3cx instance used for VOIP.",
				},
				{
					Title:       "Sentry",
					Hostname:    "sentry.myinfra.lol",
					Port:        "9000",
					Description: "A Sentry instance used for error reporting.",
				},
				{
					Title:       "Unifi Controller",
					Hostname:    "unifi.myinfra.lol",
					Port:        "443",
					Description: "A Unifi Controller instance used for managing our network.",
				},
				{
					Title:       "Mail Server: SMTP",
					Hostname:    "mail.webmailapp.net",
					Port:        "26",
					Description: "A SMTP server used for sending emails.",
				},
				{
					Title:       "Mail Server: IMAP",
					Hostname:    "mail.webmailapp.net",
					Port:        "993",
					Description: "A IMAP server used for getting emails.",
				},
				{
					Title:       "Mail Server: Web",
					Hostname:    "mail.webmailapp.net",
					Port:        "443",
					Description: "A mail server frontend used for sending emails.",
				},
				{
					Title:       "Mail Filter",
					Hostname:    "filter.webmailapp.net",
					Port:        "8006",
					Description: "A spam filter used for filtering emails.",
				},
				{
					Title:       "Dawg Squad Bot",
					Hostname:    "verify.myinfra.lol",
					Port:        "80",
					Description: "A Discord bot used for verification.",
				},
				{
					Title:       "Customer VM 1",
					Hostname:    "66.103.134.42",
					Port:        "22",
					Description: "A customer VM used for hosting a website.",
				},
				{
					Title:       "Customer VM 2",
					Hostname:    "66.103.134.47",
					Port:        "22",
					Description: "A customer VM used for hosting a website.",
				},
				{
					Title:       "Customer VM 3",
					Hostname:    "66.103.134.33",
					Port:        "22",
					Description: "A customer VM used for hosting a website.",
				},
				{
					Title:       "Customer VM 4",
					Hostname:    "23.142.248.100",
					Port:        "22",
					Description: "A customer VM used for hosting a website.",
				},
				{
					Title:       "Database",
					Hostname:    "data.myinfra.lol",
					Port:        "27017",
					Description: "A MongoDB instance used for storing data.",
				},
				{
					Title:       "Tritan Bot",
					Hostname:    "data.myinfra.lol",
					Port:        "65535",
					Description: "A Discord multipurpose bot.",
				},
				{
					Title:       "Lewd API",
					Hostname:    "yaoi.myinfra.lol",
					Port:        "4001",
					Description: "A lewd yaoi API.",
				},
				{
					Title:       "Bin",
					Hostname:    "bin.myinfra.lol",
					Port:        "7777",
					Description: "A hastebin instance used for storing data.",
				},
				{
					Title:       "Xfy's Portfolio",
					Hostname:    "portfolio.myinfra.lol",
					Port:        "4801",
					Description: "A portfolio website for Xfy.",
				},
				{
					Title:       "Tritan Dev's Website",
					Hostname:    "main-site.myinfra.lol",
					Port:        "4000",
					Description: "The website for Tritan Dev.",
				},
				{
					Title:       "File Uploader",
					Hostname:    "uploader.myinfra.lol",
					Port:        "8080",
					Description: "A file uploader to share screenshots with custom domains.",
				},
			},
		},
	}
}
