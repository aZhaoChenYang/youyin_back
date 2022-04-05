package reqstatus

const (
	OK         = 0
	DBERR      = "4001"
	NODATA     = "4002"
	DATAEXIST  = "4003"
	DATAERR    = "4004"
	SESSIONERR = "4101"
	LOGINERR   = "4102"
	PARAMERR   = "4103"
	USERERR    = "4104"
	ROLEERR    = "4105"
	PWDERR     = "4106"
	REQERR     = "4201"
	IPERR      = "4202"
	THIRDERR   = "4301"
	IOERR      = "4302"
	SERVERERR  = "4500"
	UNKOWNERR  = "4501"
)
