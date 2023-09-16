package api

type Author struct{
	Id int
	Name string
	Books []string
}

type Book struct{
	Id int
	Title string
	Description string
	Author string
}
