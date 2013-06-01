package cookoo

type Params struct {
	storage map[string]interface{}
}

// Create a new Params instance, initialized with the given map.
// Note that the given map is actually used (not copied).
func NewParamsWithValues (initialVals map[string]interface{}) *Params {
	p := new(Params)
	p.Init(initialVals)
	return p
}

func NewParams (size int) *Params {
	p := new(Params)
	p.storage = make(map[string]interface{}, size);
	return p;
}

func (p *Params) Init(initialValues map[string]interface{}) {
	p.storage = initialValues
}

func (p *Params) set(name string, value interface{}) bool {
	_, ret := p.storage[name]
	p.storage[name] = value
	return ret;
}

// Check if a parameter exists, and return it if found.
func (p *Params) Has(name string) (value interface{}, ok bool) {
	value, ok = p.storage[name]
	return
}

// Get a parameter value, or return the default value.
func (p *Params) Get(name string, defaultValue interface{}) interface{} {
	val, ok := p.Has(name)
	if ok {
		return val
	}
	return defaultValue
}

// Require that a given list of parameters are present.
// If they are all present, ok = true. Otherwise, ok = false and the 
// `missing` array contains a list of missing params.
func (p *Params) Requires(paramNames ...string) (ok bool, missing []string) {
	missing = make([]string, 0, len(p.storage))
	for _, val := range paramNames {
		_, ok := p.storage[val]
		if !ok {
			missing = append(missing, val)
		}
	}
	ok = len(missing) == 0
	return
}

// Given a name and a validation function, return a valid value.
// If the value is not valid, ok = false.
func (p *Params) Validate(name string, validator func(interface{})bool) (value interface{}, ok bool) {
	value, ok = p.storage[name]
	if !ok {
		return
	}

	if !validator(value.(interface{})) {
		// XXX: For safety, we set a failed value to nil.
		value = nil
		ok = false
	}
	return
}