package apto

func Install(args []string) error {
	command, err := unOrInstallCommand(args, "install", false)
	if err != nil {
		return err
	}

	return command.Execute()
}

func Uninstall(args []string, force bool) error {
	command, err := unOrInstallCommand(args, "uninstall", force)
	if err != nil {
		return err
	}

	return command.Execute()
}

func unOrInstallCommand(args []string, method string, force bool) (command *Command, err error) {
	pkgs := args[1:]

	command = NewCommand()

	if method == "install" {
		err = command.Install(pkgs, []string{})
	}

	if method == "uninstall" {
		err = command.Uninstall(pkgs, []string{}, force)
	}

	if err != nil {
		return nil, err
	}

	return command, nil
}
