package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"

	"github.com/bhupeshbhatia/domain/resource"
	"github.com/brianvoe/gofakeit/v6"
)

func genListOfSkills() *resource.Skill {

	var skills resource.Skill

	// Open our jsonFile
	jsonFile, err := os.Open("listOfSkills.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic("Unable to read ListOfSkills.json file")
	}
	fmt.Println("Successfully Opened List of Skills.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic("Unable to read jsonFile")
	}

	json.Unmarshal(byteValue, &skills)

	return &skills

}

func GenFakeData() *resource.Resource {
	gofakeit.AddFuncLookup("hoursWorked", gofakeit.Info{
		Category:    "custom",
		Description: "Hours available during the week",
		Example:     "30.5",
		Output:      "float64",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return (math.Round(gofakeit.Float64Range(1, 40))), nil
		},
	})

	gofakeit.AddFuncLookup("skills", gofakeit.Info{
		Category:    "custom",
		Description: "Job skills",
		Example:     "MS Word",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {

			// listOfSkills := []string{"MS Word", "MS Excel", "Workday", "ServiceNow", "Custom Design", "Database", "UML", "Health", "Indesign", "Data governance", "Cloud Security", "Vacation", "Compensation", "Payment"}

			listOfSkills := genListOfSkills()
			var skillArr []string
			length := rand.Intn(len(listOfSkills.Skills))

			for i := 0; i <= length; i++ {
				skill := listOfSkills.Skills[rand.Intn(len(listOfSkills.Skills))]

				skillArr = append(skillArr, skill)
			}

			return gofakeit.RandomString(skillArr), nil
		},
	})

	r := resource.Resource{}
	gofakeit.Struct(&r)
	r.Availability = 40 - r.Hours
	return &r
}
