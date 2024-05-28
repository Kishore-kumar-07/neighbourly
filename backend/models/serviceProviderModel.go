package models

type ServiceProviderModel struct {
	Email string `json:"email" bson:"email"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Experience string `json:"experience" bson:"experience"`
	ServiceDescription string `json:"serviceDescription" bson:"serviceDescription"`
	Rating float64 `json:"rating" bson:"rating"`
}

func (s * ServiceProviderModel) SetDefaults() {
	s.Email = ""
	s.Title = ""
	s.Description = ""
	s.Experience = ""
	s.ServiceDescription = ""
	s.Rating = 0
}