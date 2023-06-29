package protocol

type Dir struct {
	Id       int64
	HostIp   string
	Domain   string
	Lastseen interface{}
}

type File struct {
	Id       int64
	Name     string
	Url      string
	SiteId   int64
	Lastseen interface{}
}

type Database struct {
	Dirs  []Dir
	Files []File
}
