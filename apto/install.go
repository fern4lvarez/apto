package apto

func Install(args []string) error {
	command, err := installCommand(args)
	if err != nil {
		return err
	}

	return Execute(command)
}

func installCommand(args []string) (*Command, error) {
	pkgs := args[2:]

	command := NewCommand()
	err := command.Install(pkgs, []string{})
	if err != nil {
		return nil, err
	}

	return command, nil
}
