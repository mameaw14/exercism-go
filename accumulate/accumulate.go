package accumulate

const testVersion = 1

func Accumulate(collection []string, f func(string) string) []string {
	for i, v := range collection {
		collection[i] = f(v)
	}
	return collection
}
