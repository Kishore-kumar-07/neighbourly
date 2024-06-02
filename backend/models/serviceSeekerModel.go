package models

const (
	 Pending = "pending"
	 Accepted = "accepted"
	 Rejected = "rejected"
	 Completed = "completed"
)

type Service struct {
	SeekerEmail string `json:"seekerEmail"`
	ProviderEmail string `json:"providerEmail"`
	Date string `json:"date"`
	Time string `json:"time"`
	Description string `json:"description"`
	Status string `json:"status"`
}

func (b *Service) SetDefaults() {
	b.SeekerEmail = ""
	b.ProviderEmail = ""
	b.Date = ""
	b.Time = ""
	b.Description = ""
	b.Status = Pending
}