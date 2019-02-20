package output

type Transformer struct {
    TransformFunction func(interface{}) map[string]interface{}
}

func (self Transformer) Transform(data interface{}) map[string]interface{} {
    return self.TransformFunction(data)
}
