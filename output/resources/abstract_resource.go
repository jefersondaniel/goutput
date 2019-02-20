package resources

import "github.com/jefersondaniel/goutput/output/interfaces"

type BaseResource struct {
    resourceKey string
    data interface{}
    transformer interfaces.TransformerInterface
}

func (self BaseResource) GetResourceKey() string {
    return  self.resourceKey
}

func (self BaseResource) GetData() interface{} {
    return  self.data
}

func (self BaseResource) GetTransformer() interfaces.TransformerInterface {
    return  self.transformer
}
