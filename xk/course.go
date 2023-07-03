package xk

import (
	"errors"
	"fdutools-go/captcha"
	"fdutools-go/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
	"strings"
)

type Course struct {
	Name        string        `json:"name"`
	No          string        `json:"no"`
	Code        string        `json:"code"`
	ID          int           `json:"id"`
	CourseID    int           `json:"courseId"`
	Amount      AmountInfo    `json:"amount"`
	ArrangeInfo []ArrangeInfo `json:"arrangeInfo"`
	Credit      float64       `json:"credits"`
	ExamTime    string        `json:"examTime"`
	ExamType    string        `json:"examFormName"`
	Teacher     string        `json:"teachers"`
	Campus      string        `json:"campusName"`
}

type AmountInfo struct {
	Total    int `json:"lc"`
	Selected int `json:"sc"`
}

type ArrangeInfo struct {
	Day   int `json:"weekDay"`
	Start int `json:"startUnit"`
	End   int `json:"endUnit"`
}

type Query struct {
	Name string `json:"courseName"` // 课程名称  eg. 计量经济学
	No   string `json:"lessonNo"`   // 课程序号  eg. ECON130213.01
	Code string `json:"courseCode"` // 课程代码  eg. ECON130213
}

var pattern = regexp.MustCompile(`(\[.+])[\s\S]*?(\{.+})`)

func (xk *XK) QueryCourses(query Query) ([]Course, error) {
	const url = "https://xk.fudan.edu.cn/xk/stdElectCourse!queryLesson.action"
	resp, err := xk.C.Post(
		fmt.Sprintf("%s?profileId=%s", url, xk.User.ProfileID),
		utils.StructToMap(query),
	)
	if err != nil {
		return nil, err
	}
	text := utils.ReadBody(resp.Body)
	match := pattern.FindStringSubmatch(text)
	if len(match) != 3 {
		return nil, fmt.Errorf("unexpected response: %s", text)
	}
	var courses []Course
	err = json.Unmarshal([]byte(utils.NormalizeJSON(match[1])), &courses)
	if err != nil {
		return nil, err
	}
	var amounts map[string]AmountInfo
	err = json.Unmarshal([]byte(utils.NormalizeJSON(match[2])), &amounts)
	if err != nil {
		return nil, err
	}

	for i := range courses {
		amount, ok := amounts[strconv.Itoa(courses[i].ID)]
		if ok {
			courses[i].Amount = amount
		}
	}

	// filter
	if query.Name == "" && query.No == "" && query.Code == "" {
		return courses, nil
	}
	filtered := make([]Course, 0, len(courses))
	for _, course := range courses {
		if (query.Name != "" && course.Name == query.Name) ||
			(query.No != "" && course.No == query.No) ||
			(query.Code != "" && course.Code == query.Code) {
			filtered = append(filtered, course)
		}
	}
	return filtered, nil
}

func (xk *XK) ClassSchedule(courses []Course) [][]string {
	out := make([][]string, 14)
	for i := range out {
		out[i] = make([]string, 7)
	}
	for _, course := range courses {
		schedule := course.ArrangeInfo[0]
		for i := schedule.Start; i <= schedule.End; i++ {
			out[i-1][schedule.Day] = course.Name
		}
	}
	return out
}

// action: true for select, false for drop
func (xk *XK) operateCourse(id int, action bool) error {
	payload := map[string]string{}
	if action {
		payload["optype"] = "true"
		payload["operator0"] = fmt.Sprintf("%d:true:0", id)
	} else {
		payload["optype"] = "false"
		payload["operator0"] = fmt.Sprintf("%d:false", id)
	}
	const url = "https://xk.fudan.edu.cn/xk/stdElectCourse!batchOperator.action"
	resp, err := xk.C.Post(
		fmt.Sprintf("%s?profileId=%s", url, xk.User.ProfileID),
		payload,
	)
	if err != nil {
		return err
	}
	html, err := goquery.NewDocumentFromReader(resp.Body)
	text := regexp.MustCompile(`\s`).ReplaceAllString(html.Find("div").First().Text(), "")
	if strings.Contains(text, "成功") {
		return nil
	} else {
		return errors.New(text)
	}
}

func (xk *XK) OperateCourse(id int, action bool) error {
	err := captcha.DoN(xk.C, 5)
	if err != nil {
		return err
	}
	return xk.operateCourse(id, action)
}

func (xk *XK) selectCourse(id int) error {
	return xk.operateCourse(id, true)
}

func (xk *XK) dropCourse(id int) error {
	return xk.operateCourse(id, false)
}

func (xk *XK) Select(query Query) error {
	courses, err := xk.QueryCourses(query)
	if err != nil {
		return err
	}
	if len(courses) == 0 {
		return fmt.Errorf("no course found")
	}
	slices.SortFunc(courses, func(i, j Course) bool {
		return (i.Amount.Total - i.Amount.Selected) > (j.Amount.Total - j.Amount.Selected)
	})

	cnt := 0
	for cnt < 5 {
		if cnt > 0 {
			captcha.DoN(xk.C, 5)
		}
		err = xk.selectCourse(courses[0].ID)
		if err == nil {
			return err
		}
		cnt++
	}
	return nil
}
