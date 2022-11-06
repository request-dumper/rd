package logging

//goland:noinspection SpellCheckingInspection
func getTestData() map[string][]string {
	rtn := make(map[string][]string)
	rtn["Firstkey"] = []string{"firstKey-firstValue", "firstKey-secondValue"}
	rtn["Secondkey"] = []string{"secondKey-firstValue"}
	return rtn
}
