package resources

import "github.com/jefersondaniel/goutput/output/interfaces"

type ItemResource struct {
    BaseResource
}

func CreateItemResource(resourceKey string, data interface{}, transformer interfaces.TransformerInterface) interfaces.ResourceInterface {
    return &ItemResource{BaseResource{resourceKey, data, transformer}}
}
