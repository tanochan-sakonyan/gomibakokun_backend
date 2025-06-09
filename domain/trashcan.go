package domain

type Trashcan struct {
	id              string
	latitude        float64
	longitude       float64
	image           string
	trashType       []string
	nearestBuilding string
	note            string
	selectedButton  string
}

func (t Trashcan) GetID() string {
	return t.id
}

func (t Trashcan) GetLatitudeAndLongitude() (float64, float64) {
	return t.latitude, t.longitude
}

func NewTrashcan(config *TrashcanConfig) (*Trashcan, error) {
	if err := ValidateTrashcanConfig(config); err != nil {
		return nil, err
	}

	return &Trashcan{
		id:              config.ID,
		latitude:        config.Latitude,
		longitude:       config.Longitude,
		image:           config.Image,
		trashType:       config.TrashType,
		nearestBuilding: config.NearestBuilding,
		note:            config.Note,
		selectedButton:  config.SelectedButton,
	}, nil
}
