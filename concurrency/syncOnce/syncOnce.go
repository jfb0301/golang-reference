import(
	"sync"
)

type SlowComplicatedParser interface {
	Parse(string) string 
}

var parser SlowComplicatedParser
var once sync.Once

func parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	// Do all sorts of setup and loading here 
}

