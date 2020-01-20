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
	p := AboutPage{
		Title: "Basic The Programming Language",
		Url: "/about",
	}
	return &p
}

func (p *AboutPage) GetTitle() string{
	return p.Title
}

func (p *AboutPage) GetURL() string{
	return p.Url
}

func (p *AboutPage) GetData() (*interface{}){
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
	p := LoginPage{
		Title: "Login Page for the Basic interpreter",
		Url: "/login",
	}
	return &p

}

func (p *LoginPage) GetTitle() string{
	return p.Title
}

func (p *LoginPage) GetURL() string{
	return p.Url
}

func (p *LoginPage) GetData()  (*interface{}){
	return &p.Data
}
/************* Login Page *****************/

/************* Github Page *****************/
type GithubPage struct{
	Title       string
	Url         string
	Data        interface{}
}

func LoadGithubPage() *GithubPage{
	p := GithubPage{
		Title: "Github Page",
		Url: "/github",
	}
	return &p
}

func (p *GithubPage) GetTitle() string{
	return p.Title
}

func (p *GithubPage) GetURL() string{
	return p.Url
}

func (p *GithubPage) GetData() (*interface{}){
	return &p.Data
}
/************* Github Page *****************/

/************* Terminal Page *****************/
type TerminalPage struct{
	Title       string
	Url         string
	Data        interface{}
}

func LoadTerminalPage() *TerminalPage{
	p := TerminalPage{
		Title: "Terminal for the basic interpreter",
		Url: "/terminal",
	}
	return &p
}

func (p *TerminalPage) GetTitle() string{
	return p.Title
}

func (p *TerminalPage) GetURL() string{
	return p.Url
}

func (p *TerminalPage) GetData() (*interface{}){
	return &p.Data
}
/************* Terminal Page *****************/

