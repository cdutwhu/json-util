package n3cfg

import "testing"

func TestRmFileAttrL1(T *testing.T) {
	fPln(
		RmFileAttrL1(
			"../_data/toml/test.toml",
			"./config_rm",
			"Path",
			"WebService",
			"Port",
			"Storage",
			"File",
		),
	)
}
