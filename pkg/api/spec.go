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
