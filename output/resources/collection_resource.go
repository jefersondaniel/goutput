package resources

import "github.com/jefersondaniel/goutput/output/interfaces"

type CollectionResource struct {
    BaseResource
}

func CreateCollectionResource(resourceKey string, data interface{}, transformer interfaces.TransformerInterface) interfaces.ResourceInterface {
    return &ItemResource{BaseResource{resourceKey, data, transformer}}
}
