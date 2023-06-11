package grade

import (
	"fdutools-go/fdu"
	"fdutools-go/utils"
	"github.com/goccy/go-json"
)

type Gd struct {
	*fdu.Fdu
	grades []grade
	gpa    float64
}

func New(user *utils.User) *Gd {
	fd := fdu.New(user)
	return &Gd{
		Fdu: fd,
	}
}

type grade struct {
	Code     string
	Year     string
	Semester string
	Name     string
	Credit   float64
	Grade    string
}

func (g *grade) UnmarshalJSON(buf []byte) error {
	tmp := []any{&g.Code, &g.Year, &g.Semester, &g.Name, &g.Credit, &g.Grade}
	return json.Unmarshal(buf, &tmp)
}

type gradeResp struct {
	Data []grade `json:"data"`
}

func (gd *Gd) getGrades() error {
	_, err := gd.C.Get("https://my.fudan.edu.cn")
	if err != nil {
		return err
	}

	const url = "https://my.fudan.edu.cn/data_tables/bks_xx_cj.json"
	resp, err := gd.C.Post(url, map[string]string{
		"start":  "0",
		"length": "100",
	})
	if err != nil {
		return err
	}

	var res gradeResp
	err = utils.UnmarshalResponse(resp.Body, &res)
	if err != nil {
		return err
	}
	gd.grades = res.Data

	return nil
}

func calcGPA(grades []grade) float64 {
	var credits, points float64
	for _, g := range grades {
		if g.Grade != "P" {
			credits += g.Credit
			points += g.Credit * gpaMap[g.Grade]
		}
	}
	return points / credits
}
