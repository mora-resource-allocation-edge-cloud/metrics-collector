package utils

func Must(err error) error {
	if err != nil {
		panic(err)
	}
	return err
}
