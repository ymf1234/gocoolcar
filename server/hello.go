package main

import (
	trippb "coolcar/proto/gen/go"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "asc",
		End:         "fff",
		DurationSec: 10,
		FeeCent:     10,
		StartPos: &trippb.Location{
			Latitude:  33.11,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:  35,
			Longitude: 115,
		},
		PathLocations: []*trippb.Location{
			{
				Latitude:  35,
				Longitude: 115,
			},
			{
				Latitude:  32,
				Longitude: 105,
			},
		},
		Status: trippb.TripStatus_FINISHED,
	}
	fmt.Println(&trip)

	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)

	var trip2 trippb.Trip
	err2 := proto.Unmarshal(b, &trip2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(&trip2)

	b2, err3 := json.Marshal(&trip2)

	if err3 != nil {
		panic(err3)
	}

	fmt.Printf("%s\n", b2)
}
