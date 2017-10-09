package letter

const testVersion = 1

func ConcurrentFrequency(arr []string) FreqMap {
	ch := make(chan FreqMap, len(arr))
	for _, s := range arr {
		go func(str string) {
			ch <- Frequency(str)
		}(s)
	}
	m := FreqMap{}
	for range arr {
		for i, v := range <-ch {
			m[i] += v
		}
	}
	return m
}