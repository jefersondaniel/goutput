package interfaces

type TransformerInterface interface {
    Transform(interface{}) map[string]interface{}
}
