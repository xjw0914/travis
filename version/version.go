package version

const Maj = "0"
const Min = "1"
const Fix = "0"

var (
	// Version is the current version of Travis
	// Must be a string because build scripts will read this file.
	Version = "0.1.0"

	// GitCommit is set with --ldflags "-X main.gitCommit=$(git rev-parse --short HEAD)"
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
