package postgresql 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Postgresql struct {

	/*BindIp - Descr: IP address on which the PostgresSQL server will listen.
 Default: 0.0.0.0
*/
	BindIp interface{} `yaml:"bind_ip,omitempty"`

	/*BindPort - Descr: Port on which the PostgresSQL server will listen.
 Default: 5432
*/
	BindPort interface{} `yaml:"bind_port,omitempty"`

	/*Databases - Descr: List of databases to create, along with a role name and password of a role
to create for the database.
 Default: <nil>
*/
	Databases interface{} `yaml:"databases,omitempty"`

	/*MaxConnections - Descr: Maximum number of open database connections to support.
 Default: 500
*/
	MaxConnections interface{} `yaml:"max_connections,omitempty"`

}