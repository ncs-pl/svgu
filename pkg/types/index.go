package types

import "sync"

// Index is the global object representing the Starlark configuration.
type Index struct {
	Domain  string
	Modules map[string]Module
	// internal
	lock sync.Mutex
}

// SetDomain sets the domain of the index.
func (i *Index) SetDomain(d string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Domain = d
}

// AddModule adds a module to the index.
func (i *Index) AddModule(n string, m Module) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Modules[n] = m
}

// GetModule returns a module from the index.
func (i *Index) GetModule(n string) Module {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.Modules[n]
}

// RemoveModule removes a module from the index.
func (i *Index) RemoveModule(n string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	delete(i.Modules, n)
}

// CheckModule checks if a module is in the index.
func (i *Index) CheckModule(n string) bool {
	i.lock.Lock()
	defer i.lock.Unlock()
	_, ok := i.Modules[n]
	return ok
}
