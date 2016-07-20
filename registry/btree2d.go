package registry

import "github.com/joeshaw/gengen/generic"

type BTree2D interface {
	Sync(next BTree2D, onAdd, onDel func(key1 generic.T, key2 generic.U))
	GetLayer(key1 generic.T) (SecondaryLayer, bool)
	SetLayer(key1 generic.T, layer SecondaryLayer)
	ForEach(fn func(key generic.T, layer SecondaryLayer) bool)
	ForEach2(key1 generic.T, fn func(key2 generic.U) bool)
	Put(key1 generic.T, key2 generic.U, finalizers ...func())
	Delete(key1 generic.T, key2 generic.U) bool
	Drop(key1 generic.T) bool
}
