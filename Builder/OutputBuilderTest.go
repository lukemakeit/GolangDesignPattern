package Builder

import "fmt"

func OutoutBuilderTest()  {
	ssList:=make([]StaffSalary,3)
	ssList[0]=StaffSalary{
		ID:1,
		Name:"luke",
		Salary:90.8,
	}
	ssList[1]=StaffSalary{
		ID:2,
		Name:"leo",
		Salary:95.8,
	}
	ssList[2]=StaffSalary{
		ID:3,
		Name:"专家",
		Salary:99.8,
	}
	tb:=TxtBuilder{
		result:"",
	}
	hb:=HtmlBuilder{
		result:"",
	}
	TxtDirec:=Director{
		Bd:&tb,
	}
	TxtDirec.Generate(ssList)
	fmt.Println(tb.result)

	HtmlDirec:=Director{
		Bd:&hb,
	}
	HtmlDirec.Generate(ssList)
	fmt.Println(hb.result)
}
