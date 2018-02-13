# Markov

A library to generate strings based off the Markov stochastic model.

## Usage

```go
package main

import (
	"github.com/ceruleis/markov"
	"github.com/ceruleis/markov/storage"
)

func main() {
	var chain markov.Chain

	if store, err := storage.Load("sqlite3", "markov.db"); err != nil {
		panic(err.Error)
	} else {
		chain.Storage = store
	}

	// A higher prefix length means a more accurate representation of 'real'
	// text, at the cost of a much longer 'training' time for the pool to grow.
	// Between 2 and 3 offer the best trade-off.
	chain.PrefixLen = 3

	// Adds the line into the chain store, splits into prefix/suffix pairs.
	chain.AddLine("To be fair, you have to have a very high IQ to understand Rick and Morty. The humour is extremely subtle, and without a solid grasp of theoretical physics most of the jokes will go over a typical viewer’s head. There’s also Rick’s nihilistic outlook, which is deftly woven into his characterisation- his personal philosophy draws heavily from Narodnaya Volya literature, for instance. The fans understand this stuff; they have the intellectual capacity to truly appreciate the depths of these jokes, to realise that they’re not just funny- they say something deep about LIFE. As a consequence people who dislike Rick & Morty truly ARE idiots- of course they wouldn’t appreciate, for instance, the humour in Rick’s existential catchphrase “Wubba Lubba Dub Dub,” which itself is a cryptic reference to Turgenev’s Russian epic Fathers and Sons. I’m smirking right now just imagining one of those addlepated simpletons scratching their heads in confusion as Dan Harmon’s genius wit unfolds itself on their television screens. What fools.. how I pity them.")

	// Generates a line of, at maximum, 12 words long.
	chain.GenerateLine(12)
}

```

## License

ISC licensed. See the [LICENSE](./LICENSE) file for the full license text.