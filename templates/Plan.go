package templates

type Program struct {
	category string
	exercies []struct {
		equipment  []string
		duration   int
		targetArea string
	}
	diet struct {
		meals []struct {
			ingredients []string
			allergies   []string
		}
	}
}
