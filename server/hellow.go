package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	trippb "server/proto/gen/go"

// 	"github.com/golang/protobuf/proto"
// )

// func main() {
// 	fmt.Println("hellow go!")
// 	trip := trippb.Trip{
// 		Start:       "abc",
// 		End:         "def",
// 		DurationSec: 3600,
// 		FeeCent:     10000,
// 		StartPos: &trippb.Location{
// 			Latitude:  30,
// 			Longitude: 120,
// 		},
// 		EndPos: &trippb.Location{
// 			Latitude:  35,
// 			Longitude: 115,
// 		},
// 		PathLocations: []*trippb.Location{
// 			{
// 				Latitude:  31,
// 				Longitude: 119,
// 			},
// 			{
// 				Latitude:  32,
// 				Longitude: 118,
// 			},
// 		},
// 		Status: trippb.TripStatus_IN_PROGRESS,
// 	}
// 	fmt.Println(&trip)
// 	b, err := proto.Marshal(&trip)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%X\n", b)
// 	j, err := json.Marshal(&trip)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", j)
// 	var trip2 trippb.Trip
// 	proto.Unmarshal(b, &trip2)
// 	fmt.Println(&trip2)
// }
