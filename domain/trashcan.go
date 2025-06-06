package domain

type Trashcan struct {
	ID              string   `firestore:"ID"              json:"ID"`
	Latitude        float64  `firestore:"latitude"        json:"latitude"`
	Longitude       float64  `firestore:"longitude"       json:"longitude"`
	Image           string   `firestore:"image"           json:"image"`
	TrashType       []string `firestore:"trashType"       json:"trashType"`
	NearestBuilding string   `firestore:"nearestBuilding" json:"nearestBuilding"`
	Note            string   `firestore:"note"            json:"note"`
	SelectedButton  string   `firestore:"selectedButton"  json:"selectedButton"`
}
