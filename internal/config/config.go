package config

var (
	username string
	password string
	debug    bool
	dryRun   bool
	endpoint string
)

const DefaultEndpoint = "https://api.five9.com/wsadmin/v14/AdminWebService"

func SetUsername(u string) { username = u }
func Username() string     { return username }

func SetPassword(p string) { password = p }
func Password() string     { return password }

func SetDebug(d bool) { debug = d }
func Debug() bool     { return debug }

func SetDryRun(d bool) { dryRun = d }
func DryRun() bool     { return dryRun }

func SetEndpoint(e string) {
	if e != "" {
		endpoint = e
	}
}

func Endpoint() string {
	if endpoint != "" {
		return endpoint
	}
	return DefaultEndpoint
}
