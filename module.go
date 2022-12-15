package plist

import (
	"go.k6.io/k6/js/modules"
)

// Register the extensions on module initialization.
func init() {
	modules.Register("k6/x/plist", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// plist is the exported type
		plist *PList
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu:    vu,
		plist: newPlist(vu),
	}
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.plist,
	}
}

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}
