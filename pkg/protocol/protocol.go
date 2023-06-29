package protocol

type Site struct {
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
	Sites []Site
	Files []File
}

type CreateSiteRequest struct {
	Id      int64
	HostIp  string
	Domain  string
	Command string
}

type CreateSiteResponse struct {
	Status  string
	Site    Site
	Command string
}

type GetSiteRequest struct {
	Query   string
	Filter  string
	Command string
}

type GetSiteResponse struct {
	Status  string
	Sites   []Site
	Command string
}

type UpdateUrlRequest struct {
	Url      string
	Lastseen string
	Command  string
}

type UpdateSiteResponse struct {
	Status  string
	Site    Site
	Command string
}

type CreateFileRequest struct {
	File
	Command string
}

type CreateFileResponse struct {
	Status  string
	Site    Site
	Command string
}

type GetFileRequest struct {
	Command string
	Query   string
	Filter  string
}

type GetFileResponse struct {
	Command string
	Status  string
	Query   string
	Filter  string
	Files   []File
}

type UpdateFileRequest struct {
	Command  string
	Url      string
	Lastseen interface{}
}

type UpdateFileResponse struct {
	Status  string
	Command string
	File    File
}
