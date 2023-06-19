package fdu

import (
	"fdutools-go/utils"
	"github.com/goccy/go-json"
	"log"
)

type Grade struct {
	Code     string  `json:"code"`
	Year     string  `json:"year"`
	Semester string  `json:"semester"`
	Name     string  `json:"name"`
	Credit   float64 `json:"credit"`
	Grade    string  `json:"grade"`
}

func (g *Grade) UnmarshalJSON(buf []byte) error {
	tmp := []any{&g.Code, &g.Year, &g.Semester, &g.Name, &g.Credit, &g.Grade}
	return json.Unmarshal(buf, &tmp)
}

type gradeResp struct {
	Data []Grade `json:"data"`
}

func (f *Fdu) getGrades() ([]Grade, error) {
	_, err := f.C.Get("https://my.fudan.edu.cn")
	if err != nil {
		return nil, FDUErr(ErrNetwork, err.Error())
	}

	const url = "https://my.fudan.edu.cn/data_tables/bks_xx_cj.json"
	resp, err := f.C.Post(url, map[string]string{
		"start":  "0",
		"length": "100",
	})
	if err != nil {
		return nil, FDUErr(ErrNetwork, err.Error())
	}

	var res gradeResp
	err = utils.UnmarshalResponse(resp.Body, &res)
	if err != nil {
		return nil, FDUErr(ErrParseJSON, err.Error())
	}

	return res.Data, nil
}

var gpaMap = map[string]float64{
	"A":  4.0,
	"A-": 3.7,
	"B+": 3.3,
	"B":  3.0,
	"B-": 2.7,
	"C+": 2.3,
	"C":  2.0,
	"C-": 1.7,
	"D":  1.3,
	"D-": 1.0,
	"F":  0,
	"P":  0,
}

func calcGPA(grades []Grade) float64 {
	var credits, points float64
	for _, g := range grades {
		if g.Grade != "P" {
			credits += g.Credit
			points += g.Credit * gpaMap[g.Grade]
		}
	}
	return points / credits
}

type Grades struct {
	Grades []Grade `json:"grades"`
	GPA    float64 `json:"gpa"`
}

func (f *Fdu) GetGrades() (Grades, error) {
	var out Grades
	log.Println("Getting grades")
	grades, err := f.getGrades()
	if err != nil {
		return out, err
	}
	out.Grades = grades
	out.GPA = calcGPA(grades)
	return out, nil
}
