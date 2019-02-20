package interfaces

type SerializerInterface interface {
    Collection(resourceKey string, data interface{}) interface{}
    Item(resourceKey string, data interface{}) interface{}
    Null() interface{}
}
