package chitanda

type Responsible interface {
	OnRequest() error
}