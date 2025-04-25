package validators

import "errors"

func Page(pageN int32) error {
	if pageN < 1 {
		return errors.New("invalid page")
	}
	return nil
}
