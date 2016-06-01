package main

type SiteSettings struct {
	Name             string
	SiteID           int
	StartUrl         string
	BaseUrl          string
	PagePattern      string
	PageTrigger      string
	PageSelector     string
	PageSize         string
	ItemPatternFrom  string
	ItemPatternUntil string
	ItemLength       int
}

type PageSelectors struct {
	Code        string
	Url         string
	Title       string
	Description string
	JobType     string
	Travel      string
	Salary      string
	City        string
	State       string
	PostalCode  string
	Country     string
}

type Settings struct {
	Site      SiteSettings
	Selectors PageSelectors
}

type Link struct {
	Metadata Settings
	Url      string
}

type Page struct {
	Metadata Settings
	Url      string
	Html     []byte
}

type Job struct {
	Metadata    Settings
	SiteID      int
	Code        string
	Url         string
	Title       string
	Slug        string
	Description string
	JobType     string
	Travel      string
	Salary      string
	City        string
	State       string
	PostalCode  string
	Country     string
	Content     string
	Html        string
}
