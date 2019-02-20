package interfaces

type ResourceInterface interface {
    GetResourceKey() string
    GetData() interface{}
    GetTransformer() TransformerInterface
}
