package device

type CommandRunner interface {
	RunCommand(cmd string) (any, error)
}
