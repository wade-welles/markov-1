package markov

import (
	"github.com/ceruleis/markov/storage"
)

type Chain struct {
	PrefixLen int
	Storage   storage.Storage
}

func main() {
	var chain Chain
	var err error
	chain.Storage, err = storage.Load("sqlite3", "markov.db")
	if err != nil {
		panic(err.Error())
	}
}