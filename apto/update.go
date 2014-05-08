package apto

func Update() error {
	command := NewCommand()
	command.Update()

	return command.Execute()
}
