package components

type ListItem struct {
	Name string
	Age  int
}

type ListComponent struct {
	People []ListItem
}

func NewListComponent(people ...ListItem) ListComponent {
	return ListComponent{
		People: people,
	}
}
