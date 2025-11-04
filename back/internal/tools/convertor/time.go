package convertor

import "time"

func PtrStringToTime(v *string) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	if *v != "" {
		d, err := time.Parse("2006-01-02", *v)
		if err != nil {
			return nil, err
		}
		return &d, nil
	}
	return nil, nil
}