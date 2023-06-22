package functions

type IAppPathGenerator interface {
	CreateAppPath(path string) error
	GenerateAppCode(code string) error
}
