package plugin // import "github.com/asjustas/home-cloud-plugin"

type Plugin interface {
	GetName() string
	GetVersion() string
	Run()
}
