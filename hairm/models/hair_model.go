package models

import (
	"hairm/setup"
)

func GetAllHair() ([]Hair, error) {
	var hair Hair
	var err error
	arrHair := []Hair{}

	ctx, client := setup.Init()

	hairs := client.Collection("hairs")

	docSnaps, err := hairs.Documents(ctx).GetAll()
	if err != nil {
		return arrHair, err
	}

	for _, docSnap := range docSnaps {
		err = docSnap.DataTo(&hair)
		if err != nil {
			return arrHair, err
		}

		arrHair = append(arrHair, hair)
	}

	return arrHair, nil
}

func GetHair(id string) (Hair, error) {
	var res Hair
	var err error

	ctx, client := setup.Init()

	hairs := client.Collection("hairs")
	hair := hairs.Doc(id)

	docSnap, err := hair.Get(ctx)
	if err != nil {
		return res, err
	}

	err = docSnap.DataTo(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
