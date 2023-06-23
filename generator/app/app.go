package app

// App is a faker struct for App
type App struct {
}

// Name returns a fake app name for App
func (a App) Name() string {
	return a.Faker.RandomStringElement(appNames)
}

// Version returns a fake app version for App
func (a App) Version() string {
	return a.Faker.Numerify("v#.#.#")
}
