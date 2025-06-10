package domain

import (
	"slices"
)

func ValidateTrashcanConfig(config *TrashcanConfig) error {
	if config.ID == "" {
		return ErrInvalidInput
	}

	if err := ValidateLatitude(config.Latitude); err != nil {
		return err
	}

	if err := ValidateLongitude(config.Longitude); err != nil {
		return err
	}

	if err := ValidateTrashType(config.TrashType); err != nil {
		return err
	}

	if err := ValidateSelectedButton(config.SelectedButton); err != nil {
		return err
	}

	return nil
}

func ValidateLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return ErrInvalidInput
	}
	return nil
}

func ValidateLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return ErrInvalidInput
	}
	return nil
}

func ValidateTrashType(trashType []string) error {
	availableTypes := []string{
		"burnable",
		"unburnable",
		"pet_bottle",
		"bottle",
		"can",
		"plastic",
		"other",
		"ashtray",
		"everything",
	}

	for _, t := range trashType {
		if !slices.Contains(availableTypes, t) {
			return ErrInvalidInput
		}
	}

	return nil
}

func ValidateSelectedButton(selectedButton string) error {
	availableButtons := []string{
		"insideGate",
		"outside",
		"insideBuilding",
	}

	if !slices.Contains(availableButtons, selectedButton) {
		return ErrInvalidInput
	}

	return nil
}
