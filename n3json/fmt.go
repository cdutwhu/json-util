package n3json

import (
	"encoding/json"
	"io/ioutil"
)

var (
	mSpace = map[int]string{
		1: " ",
		2: "  ",
		3: "   ",
		4: "    ",
	}

	// SPACE is indent space
	SPACE = mSpace[2]
)

// func dumpMap(fmtJSON *string, space string, m map[string]interface{}) {
// 	for key, value := range m {
// 		chkSM := false

// 		switch value.(type) {
// 		case bool,
// 			int, int8, int16, int32, int64,
// 			uint, uint8, uint16, uint32, uint64,
// 			float32, float64:
// 			*fmtJSON += fSf("%v\"%v\": %v,\n", space, key, value)
// 		case nil:
// 			*fmtJSON += fSf("%v\"%v\": %v,\n", space, key, "null")
// 		case string:
// 			*fmtJSON += fSf("%v\"%v\": \"%v\",\n", space, key, value)
// 		default:
// 			chkSM = true
// 		}

// 		if chkSM {
// 			v := reflect.ValueOf(value)
// 			switch v.Kind() {
// 			// element is another slice is not allowed !
// 			case reflect.Slice:
// 				*fmtJSON += fSf("%v\"%v\": [\n", space, key)
// 				for i := 0; i < v.Len(); i++ {
// 					IEle := v.Index(i).Interface()
// 					switch IEle.(type) {
// 					case bool,
// 						int, int8, int16, int32, int64,
// 						uint, uint8, uint16, uint32, uint64,
// 						float32, float64:
// 						*fmtJSON += fSf("%v%v,\n", space+SPACE, IEle)
// 					case nil:
// 						*fmtJSON += fSf("%v%v,\n", space+SPACE, "null")
// 					case string:
// 						*fmtJSON += fSf("%v\"%v\",\n", space+SPACE, IEle)
// 					default:
// 						*fmtJSON += fSf("%v{\n", space+SPACE)
// 						dumpMap(fmtJSON, space+SPACE+SPACE, IEle.(map[string]interface{}))
// 						*fmtJSON += fSf("%v},\n", space+SPACE)
// 					}
// 				}
// 				*fmtJSON += fSf("%v],\n", space)

// 			case reflect.Map:
// 				*fmtJSON += fSf("%v\"%v\": {\n", space, key)
// 				dumpMap(fmtJSON, space+SPACE, v.Interface().(map[string]interface{}))
// 				*fmtJSON += fSf("%v},\n", space)

// 			default:
// 				panic(fSf("Missing Type @ %v", v.Interface()))
// 			}
// 		}
// 	}
// }

// // FmtJSON :
// func FmtJSON(jsonStr string, nSpace int) string {

// 	// recursively iterate map[string]interface{}, formatted with error comma
// 	SPACE = mSpace[nSpace]
// 	jsonMap := make(map[string]interface{})
// 	failOnErr("%v", json.Unmarshal([]byte(jsonStr), &jsonMap))

// 	fmtJSON := fSf("{\n")
// 	dumpMap(&fmtJSON, SPACE, jsonMap)
// 	fmtJSON += fSf("}")

// 	// remove last item's comma
// 	posCommaGrp := [][]int{}
// 	posGrp := regexp.MustCompile(`,\n[ ]*[\]\}]`).FindAllIndex([]byte(fmtJSON), -1)
// 	for _, pos := range posGrp {
// 		posCommaGrp = append(posCommaGrp, []int{pos[0], pos[0] + 1})
// 	}
// 	return replByPosGrp(fmtJSON, posCommaGrp, []string{""})

//  // use ReplByPosGrp instead of below code
// 	// 	in, out := []rune(fmtJSON), []rune{}
// 	// 	posGrp := regexp.MustCompile(`,\n[ ]*[\]\}]`).FindAllIndex([]byte(fmtJSON), -1)
// 	// NEXT:
// 	// 	for p, c := range in {
// 	// 		for _, posPair := range posGrp {
// 	// 			if p == posPair[0] {
// 	// 				continue NEXT
// 	// 			}
// 	// 		}
// 	// 		out = append(out, c)
// 	// 	}
// 	// 	return string(out)
// }

// ---------------------------------------------------------- //

// Fmt :
func Fmt(jsonstr string, nSpace int) string {
	if MaybeArr(jsonstr) {
		return FmtArr(jsonstr, nSpace)
	}

	jsonMap := make(map[string]interface{})
	failOnErr("%v", json.Unmarshal([]byte(jsonstr), &jsonMap))

	bytes, err := json.MarshalIndent(&jsonMap, "", mSpace[nSpace])
	failOnErr("%v", err)
	return string(bytes)
}

// FmtArr :
func FmtArr(jsonArr string, nSpace int) string {
	jsonGrp := []string{}
	for _, json := range SplitArr(jsonArr, nSpace) {
		jsonGrp = append(jsonGrp, json)
	}
	if len(jsonGrp) > 0 {
		return MakeArr(jsonGrp...)
	}
	return ""
}

// FmtFile :
func FmtFile(filename string, nSpace int) string {
	bytes, err := ioutil.ReadFile(filename)
	failOnErr("%v", err)
	if json := string(bytes); isJSON(json) {
		return Fmt(json, nSpace)
	}
	return ""
}
