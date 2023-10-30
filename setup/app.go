package setup

import "database/sql"

type Application struct {
	Env *Env
	DB  *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = app.NewDB()

	return *app
}
