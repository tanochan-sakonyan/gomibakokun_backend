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

func (t Trashcan) GetLatitude() float64 {
	return t.latitude
}

func (t Trashcan) GetLongitude() float64 {
	return t.longitude
}

func (t Trashcan) GetImage() string {
	return t.image
}

func (t Trashcan) GetTrashType() []string {
	return t.trashType
}

func (t Trashcan) GetNearestBuilding() string {
	return t.nearestBuilding
}

func (t Trashcan) GetNote() string {
	return t.note
}

func (t Trashcan) GetSelectedButton() string {
	return t.selectedButton
}

func NewTrashcan(id string, latitude float64, longitude float64, image string, trashType []string, nearestBuilding string, note string, selectedButton string) (*Trashcan, error) {
	if err := ValidateLatitude(latitude); err != nil {
		return nil, err
	}
	if err := ValidateLongitude(longitude); err != nil {
		return nil, err
	}
	if err := ValidateTrashType(trashType); err != nil {
		return nil, err
	}
	if err := ValidateSelectedButton(selectedButton); err != nil {
		return nil, err
	}

	return &Trashcan{
		id:              id,
		latitude:        latitude,
		longitude:       longitude,
		image:           image,
		trashType:       trashType,
		nearestBuilding: nearestBuilding,
		note:            note,
		selectedButton:  selectedButton,
	}, nil
}
