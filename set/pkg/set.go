package pkg

type StringSet map[string]struct{}

func NewStringSet(size int) StringSet {
	return make(map[string]struct{}, max(0, size))
}

func (set StringSet) Insert(item string) {
	set[item] = struct{}{}
}

func (set StringSet) InsertIfNotZeroValue(item string) {
	if item != "" {
		set[item] = struct{}{}
	}
}

func (set StringSet) InsertIfNotNil(item *string) {
	if item != nil {
		set.InsertIfNotZeroValue(*item)
	}
}

func (set StringSet) GetItems() []string {
	items := make([]string, 0, len(set))
	for item := range set {
		items = append(items, item)
	}
	return items
}
