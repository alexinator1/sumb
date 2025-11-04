package convertor

func PtrToInt64(v *uint64) int64 {
	if v == nil {
		return 0
	}
	return int64(*v)
}