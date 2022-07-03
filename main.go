package goPackageTraining

var defalt string = "\u001b[0m"

func Red(text string) string {
	result := "\u001b[34m" + text + defalt
	return result
}
