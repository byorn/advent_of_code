package util

type Set struct {
	items map[Item]bool
}

func (set Set) Add(item Item) {
	if _, ok := set.items[item]; ok == false {
		set.items[item] = true
	}
}

func (set Set) Contains(item Item) bool {
	return set.items[item]
}

func (set Set) Get(key Item) Item {
	return set.items[key]
}

func (set Set) Delete(key Item) {
	delete(set.items, key)
}
