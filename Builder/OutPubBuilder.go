package Builder

import (
	"fmt"
	"strconv"
)

type StaffSalary struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Salary float64 `json:"salary"`
}

type Builder interface {
	BuildHeader() string
	BuildBody(ssList []StaffSalary) string
	BuildFooter() string
	GetResult(header string,body string,footer string)
}

type TxtBuilder struct {
	result string
}
func (tb *TxtBuilder)BuildHeader() string {
	return "Hi,vito gg,下面是大家的薪资情况"
}
func (tb *TxtBuilder)BuildBody(ssList []StaffSalary) string {
	var ret string
	for _,v:=range ssList{
		ret=ret+fmt.Sprintf("id:%d\tname:%s\tsalary:%.2f\n",v.ID,v.Name,v.Salary)
	}
	return ret
}
func (tb *TxtBuilder)BuildFooter() string {
	return "**系统自动提供**"
}

func (tb *TxtBuilder)GetResult(header string,body string,footer string)  {
	tb.result=fmt.Sprintf("%s\n%s\n%s",header,body,footer)
}

type HtmlBuilder struct {
	result string
}

func (hb *HtmlBuilder)BuildHeader() string  {
	return "<H2>Hi,vito gg,下面是大家的薪资情况</h2>"
}
func (hb *HtmlBuilder)BuildBody(ssList []StaffSalary) string {
	var ret string
	ret="<table>\n<tbody>"
	for _,v:=range ssList{
		ret=ret+"<tr>"
		ret=ret+"<td>"+strconv.Itoa(v.ID)+"</td>"
		ret=ret+"<td>"+v.Name+"</td>"
		ret=ret+"<td>"+fmt.Sprintf("%.2",v.Salary)+"</td>"
		ret=ret+"</tr>\n"
	}
	ret=ret+"</tbody>\n</table>"
	return ret
}
func (hb *HtmlBuilder)BuildFooter() string {
	return `
<footer>
<div>*该邮件由系统自动推送,请勿回复*</div>
</footer>
`
}
func (hb *HtmlBuilder)GetResult(header string,body string,footer string) {
	hb.result=fmt.Sprintf(`
<html>
<head>
<meta charset="utf-8">
</head>
<body>
<div id="test">
	%s
	%s
	%s
</div>
</body>
</html>
`,header,body,footer)
}

type Director struct {
	Bd Builder
}

func (d *Director)Generate(ssList []StaffSalary) {
	header:=d.Bd.BuildHeader()
	body:=d.Bd.BuildBody(ssList)
	footer:=d.Bd.BuildFooter()
	d.Bd.GetResult(header,body,footer)
}