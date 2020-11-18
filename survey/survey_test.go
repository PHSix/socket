package survey


import (
  "testing"
  "fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

func TestSur(t *testing.T){
	answers := &model{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
	}
}
