package dataaccess

type SQLConnection interface {
	Configure(Hostname *string, Port *int32, Username *string, Password *string, Database *string) error
}
