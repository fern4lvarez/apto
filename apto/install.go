package apto

func Install(args []string) error {
	command, err := unOrInstallCommand(args, "install")
	if err != nil {
		return err
	}

	return command.Execute()
}

func Uninstall(args []string) error {
	command, err := unOrInstallCommand(args, "uninstall")
	if err != nil {
		return err
	}

	return command.Execute()
}

func unOrInstallCommand(args []string, method string) (command *Command, err error) {
	pkgs := args[2:]

	command = NewCommand()

	if method == "install" {
		err = command.Install(pkgs, []string{})
	}

	if method == "uninstall" {
		err = command.Uninstall(pkgs, []string{})
	}

	if err != nil {
		return nil, err
	}

	return command, nil
}
