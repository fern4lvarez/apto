package apto

func Update() error {
	command := NewCommand()
	command.Update()

	return command.Execute()
}

func Upgrade() error {
	command := NewCommand()
	command.Update()
	if err := command.Execute(); err != nil {
		return err
	}

	command = NewCommand()
	command.Upgrade()

	return command.Execute()
}
