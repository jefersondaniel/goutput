package output

import "github.com/jefersondaniel/goutput/output/interfaces"
import "github.com/jefersondaniel/goutput/output/serializers"

type Manager struct {
    serializer interfaces.SerializerInterface
    scopeFactory ScopeFactory
}

func CreateManager() *Manager {
    manager := Manager{}
    manager.SetSerializer(serializers.DataArraySerializer{})
    manager.SetScopeFactory(ScopeFactory{})

    return &manager
}

func (self *Manager) SetSerializer(serializer interfaces.SerializerInterface) {
    self.serializer = serializer
}

func (self Manager) GetSerializer() interfaces.SerializerInterface {
    return self.serializer
}

func (self *Manager) SetScopeFactory(scopeFactory ScopeFactory) {
    self.scopeFactory = scopeFactory
}

func (self Manager) GetScopeFactory() ScopeFactory {
    return self.scopeFactory
}

func (self *Manager) CreateData(resource interfaces.ResourceInterface) Scope {
    return self.scopeFactory.CreateScopeFor(self, resource)
}
