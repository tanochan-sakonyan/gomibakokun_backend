package interfaces

type CreateTrashcanRequest struct {
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	Image           string   `json:"image"`
	TrashType       []string `json:"trash_type"`
	NearestBuilding string   `json:"nearest_building"`
	Note            string   `json:"note"`
	SelectedButton  string   `json:"selected_button"`
}
