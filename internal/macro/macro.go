package macro

func Fetch(code string) ([]byte, error) {
	source := `
<macro code="test"></macro>

`

	return []byte(source), nil
}
