package travelsaver

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/pubsub"
)

type vitaliksMsg struct {
	TS    int
	ID    int
	ADDR  string
	HTTPS string
}

func (r *ReadWriter) PullMsgs() error {

	sch := scheduler{}
	sch.new()

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, r.BQ.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(r.SubID)

	// Receive blocks until the context is cancelled or an error occurs.

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		// for key, value := range msg.Attributes {
		// 	fmt.Printf("%s = %s\n", key, value)
		// }

		fmt.Println("vitalik unpacks:")
		var mv vitaliksMsg

		for key, value := range msg.Attributes {

			switch key {
			case "ID":
				i, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("strconv of ID err: %v\n", err)
				}
				mv.ID = i

			case "TS":
				t, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("strconv of TS err: %v\n", err)
				}
				mv.TS = t
			case "contract":
				mv.ADDR = value
			case "https":
				mv.HTTPS = value
			default:
				fmt.Printf("err fetching keys: %s = %s\n", key, value)

			}
		}

		msg.Ack()

		log.Print("received sub msg: \n")

		_, err := sch.createTask(mv)
		if err != nil {
			fmt.Printf("createTask err: %v\n", err)
		}

	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}

	return nil
}
