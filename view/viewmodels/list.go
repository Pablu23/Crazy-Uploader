package viewmodels

type ListItem struct {
	Name string
	Age  int
}

type ListViewmodel struct {
	People []ListItem
}

func NewListViewmodel(people ...ListItem) ListViewmodel {
	return ListViewmodel{
		People: people,
	}
}
