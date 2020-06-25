package main

func main() {

}

// Installer checks if required files/dependencies exists, download when necessary
type Installer interface{
	EnsureRequirements() error
}

// Runner starts/stops the world
type Runner interface {
	Run() error
	Stop() error
}
