package protocol

type Site struct {
	Id       int64       `json:"id"`
	HostIp   string      `json:"hostip"`
	Domain   string      `json:"domain"`
	Lastseen interface{} `json:"lastseen"`
}

type File struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name"`
	Url      string      `json:"url"`
	SiteId   int64       `json:"siteid"`
	Lastseen interface{} `json:"lastseen"`
}

type Database struct {
	Sites []Site `json:"sites"`
	Files []File `json:"files"`
}

type CreateSiteRequest struct {
	Id      int64  `json:"id"`
	HostIp  string `json:"hostip"`
	Domain  string `json:"domain"`
	Command string `json:"command"`
}

type CreateSiteResponse struct {
	Status  string `json:"status"`
	Site    Site   `json:"site"`
	Command string `json:"command"`
}

type GetSiteRequest struct {
	Query   string `json:"query"`
	Filter  string `json:"filter"`
	Command string `json:"command"`
}

type GetSiteResponse struct {
	Status  string `json:"status"`
	Sites   []Site `json:"sites"`
	Command string `json:"command"`
}

type UpdateUrlRequest struct {
	Url      string `json:"url"`
	Lastseen string `json:"lastseen"`
	Command  string `json:"command"`
}

type UpdateSiteResponse struct {
	Status  string `json:"status"`
	Site    Site   `json:"site"`
	Command string `json:"command"`
}

type CreateFileRequest struct {
	File    File   `json:"file"`
	Command string `json:"command"`
}

type CreateFileResponse struct {
	Status  string `json:"id"`
	Site    Site   `json:"site"`
	Command string `json:"command"`
}

type GetFileRequest struct {
	Command string `json:"command"`
	Query   string `json:"query"`
	Filter  string `json:"filter"`
}

type GetFileResponse struct {
	Command string `json:"id"`
	Status  string `json:"status"`
	Query   string `json:"query"`
	Filter  string `json:"filter"`
	Files   []File `json:"files"`
}

type UpdateFileRequest struct {
	Command  string      `json:"command"`
	Url      string      `json:"url"`
	Lastseen interface{} `json:"lastseen"`
}

type UpdateFileResponse struct {
	Status  string `json:"status"`
	Command string `json:"command"`
	File    File   `json:"file"`
}
