package main

import "hairm/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}

// func main() {
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, "cukuran-8a650")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	hairs := client.Collection("hairs")
// 	ny := hairs.Doc("08ba86e0-3593-11ed-b2dd-ef4277669c2e")

// 	docsnap, err := ny.Get(ctx)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	dataMap := docsnap.Data()
// 	fmt.Println(dataMap)
// }
