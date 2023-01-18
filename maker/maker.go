package maker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/pubsub"
)

type Maker struct {
	SubID     string
	ProjectID string
	LogInfo   *log.Logger
}

func (m Maker) PullMsgs() error {

	sch := scheduler{}
	sch.new()

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, m.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(m.SubID)

	// Receive blocks until the context is cancelled or an error occurs.

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		m.LogInfo.Printf("Got message: %q\n", string(msg.Data))
		// for key, value := range msg.Attributes {
		// 	fmt.Printf("%s = %s\n", key, value)
		// }

		m.LogInfo.Panicln("vitalik unpacks:")
		var mv vitaliksMsg

		for key, value := range msg.Attributes {

			switch key {
			case "ID":
				i, err := strconv.Atoi(value)
				if err != nil {
					m.LogInfo.Printf("strconv of ID err: %v\n", err)
				}
				mv.ID = i

			case "TS":
				t, err := strconv.Atoi(value)
				if err != nil {
					m.LogInfo.Printf("strconv of TS err: %v\n", err)
				}
				mv.TS = t
			case "contract":
				mv.ADDR = value
			case "https":
				mv.HTTPS = value
			default:
				m.LogInfo.Printf("err fetching keys: %s = %s\n", key, value)

			}
		}

		msg.Ack()

		m.LogInfo.Print("received sub msg: \n")

		_, err := sch.createTask(mv, m.LogInfo)
		if err != nil {
			m.LogInfo.Printf("createTask err: %v\n", err)
		}

	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}

	return nil
}

func prettyPrint(d ...interface{}) {
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// func prettyPrint(i interface{}) string {
// 	s, _ := json.MarshalIndent(i, "", "\t")
// 	return string(s)
// }
