package quickstarts

type Quickstart struct {
	ID             string
	Owner          string
	Name           string
	Version        string
	Language       string
	Framework      string
	Tags           []string
	DownloadZipURL string
	GitServer      string
	GitKind        string
}

type QuickstartModel struct {
	Quickstarts map[string]*Quickstart
}

type QuickstartFilter struct {
	Language    string
	Framework   string
	Owner       string
	Text        string
	ProjectName string
	Tags        []string
	AllowML     bool
}

type QuickstartForm struct {
	Quickstart *Quickstart
	Name       string
}
