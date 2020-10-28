package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jksch/cc/internal/models"
)

type App struct {
	Addr string
}

// LoadCourses will request all courses.
func (a *App) LoadCourses() (courses []models.Course, err error) {
	var res *http.Response
	res, err = http.Get(a.Addr + "/api/courses")
	if err != nil {
		return nil, fmt.Errorf("could not request courses, %v", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("could not request courses, server responded with: %d", res.StatusCode)
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&courses); err != nil {
		return nil, fmt.Errorf("could not decode courses JSON, %v", err)
	}
	return courses, nil
}

// SaveCourse will save a given course and return it's ID.
func (a *App) SaveCourse(course models.Course) (ID int64, err error) {
	buf := bytes.NewBuffer(nil)
	_ = json.NewEncoder(buf).Encode(&course) // Type models.Course can always be encoded.

	res, err := http.Post(a.Addr+"/api/courses", "application-json", buf)
	if err != nil {
		return 0, fmt.Errorf("could not save course, %v", err)
	}
	if res.StatusCode != 201 {
		return 0, fmt.Errorf("could not save course, server responded with %d", res.StatusCode)
	}
	b, _ := ioutil.ReadAll(res.Body)
	ID, err = strconv.ParseInt(string(bytes.Trim(b, " \n\t")), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not extract saved courses ID from response, %v", err)
	}
	return ID, nil
}
