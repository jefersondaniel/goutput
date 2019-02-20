package output

import "github.com/jefersondaniel/goutput/output/interfaces"
import "github.com/jefersondaniel/goutput/output/resources"

type Scope struct {
    manager *Manager
    resource interfaces.ResourceInterface
}

func (self Scope) GetData() interface{} {
    serializer := self.manager.GetSerializer()
    rawData := self.ExecuteResourceTransformers()
    data := self.SerializeResource(serializer, rawData)

    return data
}

func (self Scope) ExecuteResourceTransformers() interface{} {
    transformer := self.resource.GetTransformer()
    data := self.resource.GetData()

    var transformedData interface{}

    switch self.resource.(type) {
        case *resources.ItemResource:
            transformedData = self.FireTransformer(transformer, data)
        case *resources.CollectionResource:
            transformedData = make([]interface{}, 0)
            for _, item := range data.([]interface{}) {
                transformedData = append(transformedData.([]interface{}), self.FireTransformer(transformer, item))
            }
        default:
            transformedData = nil
    }

    return transformedData
}

func (self Scope) SerializeResource(serializer interfaces.SerializerInterface, data interface{}) interface{} {
    resourceKey := self.resource.GetResourceKey()

    switch self.resource.(type) {
        case *resources.ItemResource:
            return serializer.Item(resourceKey, data)
        case *resources.CollectionResource:
            return serializer.Collection(resourceKey, data)
        default:
            return serializer.Null()
    }
}

func (self Scope) FireTransformer(transformer interfaces.TransformerInterface, data interface{}) interface{} {
    transformedData := transformer.Transform(data)
    return transformedData
}
