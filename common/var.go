package common

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	fPt         = fmt.Print
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	fSp         = fmt.Sprint
	fEf         = fmt.Errorf
	sSpl        = strings.Split
	sJoin       = strings.Join
	sCount      = strings.Count
	sReplace    = strings.Replace
	sReplaceAll = strings.ReplaceAll
	sIndex      = strings.Index
	sLastIndex  = strings.LastIndex
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix

	RExpMD5    = regexp.MustCompile("\"[a-f0-9]{32}\"")
	RExpSHA1   = regexp.MustCompile("\"[a-f0-9]{40}\"")
	RExpSHA256 = regexp.MustCompile("\"[a-f0-9]{64}\"")
)
