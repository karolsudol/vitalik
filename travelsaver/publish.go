package travelsaver

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/pubsub"
)

func (r ReadWriter) publishIntervals(ID, TS int) error {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, r.BQ.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic("intervals")
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(r.ContractAddress),
		Attributes: map[string]string{
			"ID":       strconv.Itoa(ID),
			"TS":       strconv.Itoa(TS),
			"contract": r.ContractAddress,
			"https":    r.HTTPS,
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Printf("Published intervals message; msg ID: %v\n", id)
	return nil
}
