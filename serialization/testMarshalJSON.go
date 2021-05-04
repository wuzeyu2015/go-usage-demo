package main

import "encoding/json"

type Bucket struct {
	FailureDomain string               `json:"failure_domain"`
	Alg           string                  `json:"alg"`
	Hash          string                  `json:"hash"`
	Bucket        map[string]*Bucket	`json:"bucket"`
}

type BucketV struct {
	FailureDomain int               	`json:"failure_domain"`
	Alg           int                  `json:"alg"`
	Hash          int                  `json:"hash"`
	Bucket        map[string]*BucketV	`json:"bucket"`
}

type Stop struct {
	FailureDomain string               `json:"failure_domain"`
	Alg           string                  `json:"alg"`
	Hash          string                  `json:"hash"`
}

type fakeStop Stop

type testStop struct {
	Test1 string               `json:"test_stop"`
	Stops []Stop				`json:"stops"`
}
func main() {
	buckets := Bucket{
		FailureDomain: "1",
		Alg:           "2",
		Hash:          "3",
		Bucket: map[string]*Bucket{"root1":{FailureDomain: "11",
											Alg:           "22",
											Hash:          "33",
											Bucket: make(map[string]*Bucket, 0)},
									"root2":{FailureDomain: "11",
										    Alg:           "22",
										    Hash:          "33",
										    Bucket: make(map[string]*Bucket, 0)}},
	}
	//testStops := testStop{Test1: "testStop",
	//						Stops: []Stop{{
	//							FailureDomain: "111",
	//							Alg:           "222",
	//							Hash:          "333",
	//						}, {
	//							FailureDomain: "111",
	//							Alg:           "222",
	//							Hash:          "333",
	//						}}}
	bytes, err := json.Marshal(buckets)
	if err == nil {
		print(string(bytes))
	}
}

//func (s Bucket) MarshalJSON() ([]byte, error) {
//	b, err := json.Marshal( Stop{FailureDomain: "1",
//		Alg:           "2",
//		Hash:          "3"})
//	if err != nil {
//		print(err)
//	}
//	return b, err
//}

//func (s Stop) MarshalJSON() ([]byte, error) {
//	fake := Stop{
//		FailureDomain: "fffffffff",
//		Alg:           "ffffffff",
//		Hash:          "fffffff",
//	}
//	b, err := json.Marshal(fakeStop(fake))
//	if err != nil {
//		print(err)
//	}
//	return b, err
//}