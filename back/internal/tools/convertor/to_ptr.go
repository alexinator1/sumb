package convertor

func StrPtrIfNotEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func Int64PtrIfNotZero(v uint64) *int64 {
	if v == 0 {
		return nil
	}
	x := int64(v)
	return &x
}	