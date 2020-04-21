package n3errs

import "fmt"

var (
	fEf = fmt.Errorf
)

var (
	FOR_TEST = fEf("FOR TESTING")

	EXTERNAL = fEf("EXTERNAL ERROR")

	NET_TIMEOUT     = fEf("NET TIME OUT")
	NET_NO_RESPONSE = fEf("NO NET RESPONSE")

	INTERNAL          = fEf("INTERNAL ERROR")
	INTERNAL_SCAN_ERR = fEf("INTERNAL ERROR - SCAN")
	INTERNAL_INIT_ERR = fEf("INTERNAL ERROR - INIT")
	INTERNAL_DEADLOCK = fEf("INTERNAL ERROR - DEADLOCK")

	CLI_ARG_ERR        = fEf("CLI ARGUMENT ERROR")
	CLI_FLAG_ERR       = fEf("CLI FLAG ERROR")
	CLI_SUBCMD_ERR     = fEf("CLI SUB COMMAND ERROR")
	CLI_SUBCMD_UNKNOWN = fEf("CLI SUB COMMAND UNKNOWN")

	CFG_INIT_ERR     = fEf("CONFIG INIT ERROR")
	CFG_SIGN_MISSING = fEf("CONFIG KEY SIGN MISSING")

	SRC_SIGN_MISSING = fEf("SOURCE FILE KEY SIGN MISSING")

	PARAM_INVALID      = fEf("INVALID PARAM(S)")
	PARAM_INVALID_JSON = fEf("INVALID JSON PARAM")
	PARAM_INVALID_XML  = fEf("INVALID XML PARAM")
	PARAM_INVALID_CSV  = fEf("INVALID CSV PARAM")
	PARAM_INVALID_FMT  = fEf("INVALID PARAM FORMAT")

	VAR_INVALID = fEf("INVALID VARIABLE(S)")

	DIR_NOT_FOUND = fEf("DIRECTORY IS NOT FOUND")

	FILE_NOT_FOUND = fEf("FILE IS NOT FOUND")
	FILE_EMPTY     = fEf("FILE IS EMPTY OR READ ERROR")

	STR_EMPTY = fEf("EMPTY STRING")
	STR_BLANK = fEf("BLANK STRING")

	XML_INVALID = fEf("INVALID XML")
	XML_NOT_FMT = fEf("XML IS NOT FORMATTED")

	JSON_INVALID       = fEf("INVALID JSON")
	JSON_NOT_FMT       = fEf("JSON IS NOT FORMATTED")
	JSON_ARRAY_INVALID = fEf("INVALID JSON ARRAY")
	JSON_ARRAY_NOT_FMT = fEf("JSON ARRAY IS NOT FORMATTED")

	CSV_COLUMN_HEADER_EMPTY = fEf("CSV HAS EMPTY HEADER")

	MAP_INVALID         = fEf("NOT A MAP")
	MAPS_DIF_KEY_TYPE   = fEf("DIFFERENT KEY TYPE")
	MAPS_DIF_VALUE_TYPE = fEf("DIFFERENT VALUE TYPE")

	SLICE_INVALID         = fEf("NOT A SLICE")
	SLICE_INCORRECT_COUNT = fEf("INCORRECT ELEMENT COUNT")
	SLICES_DIF_LEN        = fEf("DIFFERENT LENGTH")
)
