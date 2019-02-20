package output

import "github.com/jefersondaniel/goutput/output/interfaces"

type ScopeFactory struct{}

func (*ScopeFactory) CreateScopeFor(manager *Manager, resource interfaces.ResourceInterface) Scope {
    return Scope{manager, resource}
}

