package setup

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

func Init() (context.Context, *firestore.Client) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cukuran-8a650")
	if err != nil {
		fmt.Println(err.Error())
	}

	return ctx, client
}
