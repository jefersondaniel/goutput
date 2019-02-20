package serializers

type DataArraySerializer struct {}

func (DataArraySerializer) Collection(resourceKey string, data interface{}) interface{} {
    if resourceKey == "data" {
        return map[string]interface{}{
            "data": data,
        }
    }

    return data
}

func (DataArraySerializer) Item(resourceKey string, data interface{}) interface{} {
    if resourceKey == "data" {
        return map[string]interface{}{
            "data": data,
        }
    }

    return data
}

func (DataArraySerializer) Null() interface{} {
    return make([]interface{}, 0)
}
