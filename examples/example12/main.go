package main

type User struct {
}
type Team struct {
}

type DBEntitiesTypes interface {
	User | Team
}

type GenericResponse[T DBEntitiesTypes] struct {
	dbErr    error
	dbEntity T
}

func (gr GenericResponse[T]) Error() error {
	return gr.dbErr
}
func (gr GenericResponse[T]) Entity() T {
	return gr.dbEntity
}

func parseUser(user User) {
	/*
		...
	*/
}
func parseTeam(team Team) {
	/*
		...
	*/
}

func queryDB[T DBEntitiesTypes]() []GenericResponse[T] {
	return make([]GenericResponse[T], 10)
}

func main() {
	users := queryDB[User]()
	teams := queryDB[Team]()

	for _, u := range users {
		if u.Error() != nil {
			continue
		} else {
			parseUser(u.Entity())
		}
	}
	for _, t := range teams {
		if t.Error() != nil {
			continue
		} else {
			parseTeam(t.Entity())
		}
	}

}
