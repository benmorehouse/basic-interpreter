package main

type Page interface{
	GetTitle()   string
	GetURL()     string
	GetData()    *interface{}
}

/************* About Page *****************/
type AboutPage struct{
	Title       string
	Url         string
	Data        interface{}
}

func LoadAboutPage() *AboutPage{
	return p := AboutPage{
		Title: "Basic The Programming Language",
		Url: "/about",
	}
}

func (p *AboutPage) GetTitle() string{
	return p.Title
}

func (p *AboutPage) GetURL() string{
	return p.Url
}

func (p *AboutPage) GetData() *interface{}{
	return &p.Data
}
/************* About Page *****************/

/************* Login Page *****************/
type LoginPage struct{
	Title       string
	Url         string
	Data        interface{}
}

func LoadLoginPage() *LoginPage{
	
}

func (p *LoginPage) GetTitle() string{
	return p.Title
}

func (p *LoginPage) GetURL() string{
	return p.Url
}

func (p *LoginPage) GetData() string{
	return p.Data
}
/************* Login Page *****************/

/************* Github Page *****************/
type GithubPage struct{
	Title       string
	Url         string
	Data        string
}

func (p *GithubPage) GetTitle() string{
	return p.Title
}

func (p *GithubPage) GetURL() string{
	return p.Url
}

func (p *GithubPage) GetData() string{
	return p.Data
}
/************* Github Page *****************/

/************* Terminal Page *****************/
type TerminalPage struct{
	Title       string
	Url         string
	Data        string
}

func (p *TerminalPage) GetTitle() string{
	return p.Title
}

func (p *TerminalPage) GetURL() string{
	return p.Url
}

func (p *TerminalPage) GetData() string{
	return p.Data
}
/************* Terminal Page *****************/

