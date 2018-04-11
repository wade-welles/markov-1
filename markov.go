package markov

import (
	"github.com/elliotspeck/markov/storage"
	"strings"
)

type Chain struct {
	PrefixLen int
	Storage   storage.Storage
}

func (chain *Chain) AddLine(line string) error {
	if chain.PrefixLen < 2 {
		chain.PrefixLen = 2
	}

	var links []storage.Link
	words := strings.Fields(line)

	for i := 0; i <= len(words)-(chain.PrefixLen + 1); i++ {
		prefix := strings.Join(words[i:i+chain.PrefixLen], " ")
		suffix := words[i+chain.PrefixLen]

		link := storage.Link{Prefix: prefix, Suffix: suffix}
		links = append(links, link)
	}

	for _, link := range links {
		if err := chain.Storage.AddLink(link); err != nil {
			return err
		}
	}

	return nil
}

func (chain *Chain) GenerateLine(length int) (string, error) {
	return chain.Storage.GenerateLine(length)
}
