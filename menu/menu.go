package menu

type MenuItem struct {
	Line1    string
	Line2    string
	Children []MenuItem
}

func GetMenuItem(items []MenuItem, position []int) (mi MenuItem) {
	if len(items) == 0 {
		return
	}
	if len(position) == 0 {
		mi = items[0]
		return
	}
	mi = items[position[0]]
	if len(position) == 1 {
		return
	}
	for _, p := range position[1:] {
		mi = mi.Children[p]
	}
	return
}

func GetSiblings(items []MenuItem, position []int) []MenuItem {
	if len(items) == 0 {
		return items
	}
	if len(position) <= 1 {
		return items
	}
	item := items[position[0]]
	siblings := items
	for _, p := range position[1:] {
		siblings = item.Children
		item = item.Children[p]
	}
	return siblings
}

func GetChildren(items []MenuItem, position []int) []MenuItem {
	if len(items) == 0 {
		return []MenuItem{}
	}
	if len(position) == 0 {
		return []MenuItem{}
	}
	item := items[position[0]]
	for _, p := range position[1:] {
		if len(item.Children) <= p {
			return []MenuItem{}
		}
		item = item.Children[p]
	}
	return item.Children
}
