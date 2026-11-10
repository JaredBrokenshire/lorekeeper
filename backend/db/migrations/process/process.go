package process

import (
	"lorekeeper/db/migrations/list"
	go_migrations "lorekeeper/pkg/go-migrations"
	"lorekeeper/pkg/go-migrations/store"
)

func Run() {
	go_migrations.Run(List())
}

func List() []store.Migratable {
	return []store.Migratable{
		&list.CreateUsersTable{},
		&list.CreateRolesTable{},
		&list.CreateUserRolesTable{},
		&list.CreateRacesTable{},
		&list.CreateClassesTable{},
		&list.CreateSubclassesTable{},
	}
}
