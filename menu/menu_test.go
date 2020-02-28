package menu

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetChildren(t *testing.T) {
	testCases := []struct {
		desc     string
		menu     []MenuItem
		position []int
		expected []MenuItem
	}{
		{
			desc: "can get the children of a top level menu item",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
					Children: []MenuItem{
						MenuItem{
							Line1: "AA",
						},
					},
				},
			},
			position: []int{0},
			expected: []MenuItem{
				MenuItem{
					Line1: "AA",
				},
			},
		},
		{
			desc: "can get the children of a 2nd level menu item",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
					Children: []MenuItem{
						MenuItem{
							Line1: "AA",
						},
					},
				},
				MenuItem{
					Line1: "B",
					Children: []MenuItem{
						MenuItem{
							Line1: "BA",
						},
						MenuItem{
							Line1: "BB",
							Children: []MenuItem{
								MenuItem{
									Line1: "BBA",
								},
							},
						},
					},
				},
			},
			position: []int{1, 1},
			expected: []MenuItem{
				MenuItem{
					Line1: "BBA",
				},
			},
		},
		{
			desc: "if we're out of bounds, we get an empty list",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
					Children: []MenuItem{
						MenuItem{
							Line1: "AA",
						},
					},
				},
				MenuItem{
					Line1: "B",
					Children: []MenuItem{
						MenuItem{
							Line1: "BA",
						},
						MenuItem{
							Line1: "BB",
							Children: []MenuItem{
								MenuItem{
									Line1: "BBA",
								},
							},
						},
					},
				},
			},
			position: []int{1, 2},
			expected: []MenuItem{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := GetChildren(tC.menu, tC.position)
			if !cmp.Equal(actual, tC.expected) {
				t.Errorf(cmp.Diff(tC.expected, actual))
			}
		})
	}
}

func TestGetSiblings(t *testing.T) {
	testCases := []struct {
		desc     string
		menu     []MenuItem
		position []int
		expected []MenuItem
	}{
		{
			desc: "can get the siblings of a top level menu item",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
				},
				MenuItem{
					Line1: "B",
				},
			},
			position: []int{1},
			expected: []MenuItem{
				MenuItem{
					Line1: "A",
				},
				MenuItem{
					Line1: "B",
				},
			},
		},
		{
			desc: "can get the siblings of a 2nd level menu item",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
				},
				MenuItem{
					Line1: "B",
					Children: []MenuItem{
						MenuItem{
							Line1: "BA",
						},
						MenuItem{
							Line1: "BB",
						},
					},
				},
			},
			position: []int{1, 1},
			expected: []MenuItem{
				MenuItem{
					Line1: "BA",
				},
				MenuItem{
					Line1: "BB",
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := GetSiblings(tC.menu, tC.position)
			if !cmp.Equal(actual, tC.expected) {
				t.Errorf(cmp.Diff(tC.expected, actual))
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	testCases := []struct {
		desc     string
		menu     []MenuItem
		position []int
		expected MenuItem
	}{
		{
			desc: "can get a top level menu item",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
				},
				MenuItem{
					Line1: "B",
				},
			},
			position: []int{1},
			expected: MenuItem{
				Line1: "B",
			},
		},
		{
			desc: "can geta 2nd level menu item",
			menu: []MenuItem{
				MenuItem{
					Line1: "A",
				},
				MenuItem{
					Line1: "B",
					Children: []MenuItem{
						MenuItem{
							Line1: "BA",
						},
						MenuItem{
							Line1: "BB",
						},
					},
				},
			},
			position: []int{1, 1},
			expected: MenuItem{
				Line1: "BB",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := GetMenuItem(tC.menu, tC.position)
			if !cmp.Equal(actual, tC.expected) {
				t.Errorf(cmp.Diff(tC.expected, actual))
			}
		})
	}
}
