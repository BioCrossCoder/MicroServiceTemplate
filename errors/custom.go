package errors

func CustomBizErr(tag, cause string) *BizErr {
	return &BizErr{
		Tag:     tag,
		Message: customMsgs[tag],
		Cause:   cause,
	}
}

var customMsgs = map[string]string{}
