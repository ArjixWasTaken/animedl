package allProviders

import (
	"github.com/ArjixWasTaken/animedl/animedl/providers"
	"github.com/ArjixWasTaken/animedl/animedl/providers/gogoanime"
)

func GetProviders() []*providers.Provider {
	Providers := make([]*providers.Provider, 0, 1)
	Providers = append(Providers, gogoanime.GogoanimeProvider)

	return Providers
}

func GetProviderByName(name string) *providers.Provider {

	for _, provider := range GetProviders() {
		if provider.Name == name {
			return provider
		}
	}

	return nil
}
