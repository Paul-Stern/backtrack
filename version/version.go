package version

var (
	Version    = "dev"
	Commit     = "none"
	FullCommit = "none"
	BuildTime  = "unknown"
)

func GetVersion() string {
	return Version
}

func GetCommit() string {
	return Commit
}

func GetBuild() string {
	return BuildTime
}

func GetFullCommit() string {
	return FullCommit
}

func GetFullVersion() string {
	return Version + "-" + Commit + "-" + BuildTime
}
