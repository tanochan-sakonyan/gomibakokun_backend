package usecase

type TrashcanInput struct {
	Latitude        float64
	Longitude       float64
	Image           string
	TrashType       []string
	NearestBuilding string
	Note            string
	SelectedButton  string
}
