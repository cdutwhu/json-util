// ********** ALL Based On Formatted JSON ********** //

package jkv

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
	n3json "github.com/cdutwhu/n3-util/n3json"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	fEf         = fmt.Errorf
	sSplit      = strings.Split
	sJoin       = strings.Join
	sCount      = strings.Count
	sContains   = strings.Contains
	sReplace    = strings.Replace
	sReplaceAll = strings.ReplaceAll
	sIndex      = strings.Index
	sLastIndex  = strings.LastIndex
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sTrimRight  = strings.TrimRight
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sRepeat     = strings.Repeat

	xin             = cmn.XIn
	trackTime       = cmn.TrackTime
	rmTailFromLastN = cmn.RmTailFromLastN
	rmTailFromLast  = cmn.RmTailFromLast
	rmHeadToLast    = cmn.RmHeadToLast
	failOnErr       = cmn.FailOnErr
	failOnErrWhen   = cmn.FailOnErrWhen
	isJSON          = cmn.IsJSON
	isNumeric       = cmn.IsNumeric
	indent          = cmn.Indent
	iter2Slc        = cmn.Iter2Slc
	iterN           = cmn.N
	hasAnyPrefix    = cmn.HasAnyPrefix
	mapKeys         = cmn.MapKeys
	mapsJoin        = cmn.MapsJoin
	mustWriteFile   = cmn.MustWriteFile
	projectV        = cmn.ProjectV

	maybeJSONArr = n3json.MaybeArr
	splitJSONArr = n3json.SplitArr
	makeJSONArr  = n3json.MakeArr
	fmtJSON      = n3json.Fmt
	fmtJSONFile  = n3json.FmtFile
	fmtInnerJSON = n3json.InnerFmt
)

type (
	b        = byte
	JSONTYPE int
)

var (
	StartTrait = []byte{
		b('"'), // [array : string] OR [object : field]
		// b('{'), // [array : object]
		// b('n'),         // [array : null]
		// b('t'), b('f'), // [array : bool]
		// b('1'), b('2'), b('3'), b('4'), b('5'), b('6'), b('7'), b('8'), b('9'), b('-'), b('0'), // [array : number]
	}

	LF, SP, DQ = byte('\n'), byte(' '), byte('"')
)

var (
	BLANK = " \t\n\r"
	hash  = func(str string) string {
		return "\"" + cmn.SHA1Str(str) + "\""
	}
	// hash     = cmn.SHA1Str
	hashRep = cmn.RepSHA1 // compiled with ""
)

var (
	Trait4Scan = fSf("\n%s", sRepeat(" ", 64))

	StartOfObjArr = func(nGrp ...int) []string {
		starts := make([]string, len(nGrp))
		for i, n := range nGrp {
			spaceLen1 := n*2 + 2
			spaceLen2 := spaceLen1 + 2
			maxSpace := sRepeat(" ", spaceLen2)
			starts[i] = fSf("[\n%s{\n%s", maxSpace[:spaceLen1], maxSpace)
		}
		return starts
	}
	EndOfObjArr = func(nGrp ...int) []string {
		ends := make([]string, len(nGrp))
		for i, n := range nGrp {
			spaceLen1 := n*2 + 2
			spaceLen2 := spaceLen1 - 2
			maxSpace := sRepeat(" ", spaceLen1)
			ends[i] = fSf("\n%s}\n%s]", maxSpace, maxSpace[:spaceLen2])
		}
		return ends
	}
)

const (
	TraitFV    = "\": "
	Trait1EndV = ",\n"
	Trait2EndV = "\n"
	Trait3EndV = "],\n"
	Trait4EndV = "]\n"

	PathLinker = "~~"
	LvlMax     = 20 // init 20 max level in advances
)

// readonly var
var (
	pLinker = PathLinker
)

// JKV :
type JKV struct {
	JSON          string
	LsL12Fields   [][]string          // 2D slice for each Level's each ifield
	lsLvlIPaths   [][]string          // 2D slice for each Level's each ipath
	mPathMAXIdx   map[string]int      //
	mIPathPos     map[string]int      //
	MapIPathValue map[string]string   //
	mIPathOID     map[string]string   //
	mOIDiPath     map[string]string   //
	mOIDObj       map[string]string   //
	mOIDLvl       map[string]int      // from 1 ...
	mOIDType      map[string]JSONTYPE // OID-type is OBJ or ARR|OBJ
	Wrapped       bool                //
}

// ********************************************************** //

// T : JSON line Search Feature.
func T(lvl int) string {
	return Trait4Scan[:2*lvl+1]
}

// PT :
func PT(T string) string {
	return T[:len(T)-2]
}

// NT :
func NT(T string) string {
	return T[:len(T)+2]
}

// TL : get T & L by nchar
func TL(nChar int) (string, int) {
	lvl := (nChar - 1) / 2
	return T(lvl), lvl
}
