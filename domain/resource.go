package domain

//Resource struct
type Resource struct {
	Firstname       string   `json:"firstname,omitempty" bson:"firstname,omitempty" fake:"{firstname}"` // Any available function all lowercase
	Lastname        string   `json:"lastname,omitempty" bson:"lastname,omitempty" fake:"{lastname}"`    // Can call with parameters
	Gender          string   `json:"gender,omitempty" bson:"gender,omitempty" fake:"{gender}"`
	YearsExperience int      `json:"yearsExperience,omitempty" bson:"yearsExperience,omitempty" fake:"{number:1,11}"`
	Availability    float64  `json:"availability,omitempty" bson:"availability,omitempty"`
	Hours           float64  `json:"hours,omitempty" bson:"hours,omitempty" fake:"{hoursWorked}"` // Comma separated for multiple values
	Skills          []string `json:"skills,omitempty" bson:"skills,omitempty" fake:"{skills}"`
}

type Skill struct {
	Skills []string `json:"skills"`
}
