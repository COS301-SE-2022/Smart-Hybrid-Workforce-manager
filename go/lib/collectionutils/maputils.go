package collectionutils

func MapHasKey[K comparable, V any](_map map[K]V, key K) bool {
	_, ok := _map[key]
	return ok
}
