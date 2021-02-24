package errors

import libErrors "errors"

var (
	INTERNAL_SERVER_ERROR   = libErrors.New("ERR_INTERNAL_SERVER")
	NO_DATA                 = libErrors.New("ERR_NO_DATA")
	LOCATION_DATA_NOT_FOUND = libErrors.New("ERR_LOCATION_DATA_NOT_FOUND")
	INVALID_ID              = libErrors.New("ERR_INVALID_ID")
	INVALID_QUERY_LIMIT     = libErrors.New("ERR_INVALID_QUERY_LIMIT")
	INVALID_QUERY_PARAMS    = libErrors.New("ERR_INVALID_QUERY_PARAMS")
)
