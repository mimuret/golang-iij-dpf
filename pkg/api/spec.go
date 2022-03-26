package api

type Object interface {
	DeepCopyObject() Object
}

type Spec interface {
	Object
	GetName() string
	GetGroup() string
	GetPathMethod(Action) (string, string)
}

type ListSpec interface {
	Spec
	Initializer
	GetItems() interface{}
	Len() int
	Index(int) interface{}
}

type CountableListSpec interface {
	ListSpec
	SetCount(int32)
	GetCount() int32
	GetMaxLimit() int32
	AddItem(interface{}) bool
	ClearItems()
}

type Initializer interface {
	Init()
}

func DeepCopySpec(s Spec) Spec {
	if s == nil {
		return nil
	}
	ret, ok := s.DeepCopyObject().(Spec)
	if !ok {
		return nil
	}
	return ret
}

func DeepCopyListSpec(s ListSpec) ListSpec {
	if s == nil {
		return nil
	}
	ret, ok := s.DeepCopyObject().(ListSpec)
	if !ok {
		return nil
	}
	return ret
}

func DeepCopyCountableListSpec(s CountableListSpec) CountableListSpec {
	if s == nil {
		return nil
	}
	ret, ok := s.DeepCopyObject().(CountableListSpec)
	if !ok {
		return nil
	}
	return ret
}
