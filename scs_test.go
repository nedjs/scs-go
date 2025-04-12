package main

import (
	"testing"
)

func TestProfiling(t *testing.T) {
	validatedScs(t, []string{"jack", "apple", "maven", "hold", "solid", "mark", "moon", "poor", "spark", "live"})
}

func TestCases(t *testing.T) {
	tests := [][]string{
		{"aba", "bab"}, {"baba"},
		{"aaba", "abaa"}, {"aabaa"},
		{"aaba", "baaa"}, {"baaba"},
		{"aaaa", "abbb"}, {"aaaabbb"},
		{"aabb", "baab"}, {"baabb"},
		{"aaba", "baab"}, {"baaba"},
		{"aabb", "babb"}, {"baabb"},

		{"aaba", "aabb"}, {"aabab"},
		{"aabb", "aaba"}, {"aabba"},
		{"abaa", "baab"}, {"abaab"},
		{"abba", "babb"}, {"babba"},
		{"abba", "bbab"}, {"abbab"},
		{"baab", "aaba"}, {"baaba"},
		{"baab", "abaa"}, {"abaab"},
		{"babb", "abba"}, {"babba"},
		{"bbaa", "bbab"}, {"bbaab"},
		{"bbab", "abba"}, {"abbab"},
		{"bbab", "bbaa"}, {"bbaba"},
		{"abab", "aaab"}, {"abaab"},
		{"abab", "baab"}, {"abaab"},
		{"abba", "abab"}, {"abbab"},
		{"abbb", "abab"}, {"ababb"},
		{"baaa", "baba"}, {"babaa"},
		{"baab", "baba"}, {"baaba"},
		{"baba", "abba"}, {"babba"},
		{"baba", "bbba"}, {"babba"},

		{"baaacbcbc", "bacbcaca"}, {"baaacbcbaca"},
		{"bbabacaa", "cccababab"}, {"cccabbabacaab"},

		{"bcaaacbbbcbdcaddadcacbdddcdcccdadadcbabaccbccdcdcbcaccacbbdcbabb",
			"dddbbdcbccaccbababaacbcbacdddcdabadcacddbacadabdabcdbaaabaccbdaa",
		}, {"dddbbdcbccaaaccbababaacbdcbacddadcdacbadddcacdcccdbadcadcbabdaccbccdcdcbcaaabaccacbbdcbabba"},

		{"atdznrqfwlfbcqkezrltzyeqvqemikzgghxkzenhtapwrmrovwtpzzsyiwongllqmvptwammerobtgmkpowndejvbuwbporfyroknrjoekdgqqlgzxiisweeegxajqlradgcciavbpgqjzwtdetmtallzyukdztoxysggrqkliixnagwzmassthjecvfzmyonglocmvjnxkcwqqvgrzpsswnigjthtkuawirecfuzrbifgwolpnhcapzxwmfhvpfmqapdxgmddsdlhteugqoyepbztspgojbrmpjmwmhnldunskpvwprzrudbmtwdvgyghgprqcdgqjjbyfsujnnssfqvjhnvcotynidziswpzhkdszbblustoxwtlhkowpatbypvkmajumsxqqunlxxvfezayrolwezfzfyzmmneepwshpemynwzyunsxgjflnqmfghsvwpknqhclhrlmnrljwabwpxomwhuhffpfinhnairblcayygghzqmotwrywqayvvgohmujneqlzurxcpnwdipldofyvfdurbsoxdurlofkqnrjomszjimrxbqzyazakkizojwkuzcacnbdifesoiesmkbyffcxhqgqyhwyubtsrqarqagogrnaxuzyggknksrfdrmnoxrctntngdxxechxrsbyhtlbmzgmcqopyixdomhnmvnsafphpkdgndcscbwyhueytaeodlhlzczmpqqmnilliydwtxtpedbncvsqauopbvygqdtcwehffagxmyoalogetacehnbfxlqhklvxfzmrjqofaesvuzfczeuqegwpcmahhpzodsmpvrvkzxxtsdsxwixiraphjlqawxinlwfspdlscdswtgjpoiixbvmpzilxrnpdvigpccnngxmlzoentslzyjjpkxemyiemoluhqifyonbnizcjrlmuylezdkkztcphlmwhnkdguhelqzjgvjtrzofmtpuhifoqnokonhqtzxmimp",
			"xjtuwbmvsdeogmnzorndhmjoqnqjnhmfueifqwleggctttilmfokpgotfykyzdhfafiervrsyuiseumzmymtvsdsowmovagekhevyqhifwevpepgmyhnagjtsciaecswebcuvxoavfgejqrxuvnhvkmolclecqsnsrjmxyokbkesaugbydfsupuqanetgunlqmundxvduqmzidatemaqmzzzfjpgmhyoktbdgpgbmjkhmfjtsxjqbfspedhzrxavhngtnuykpapwluameeqlutkyzyeffmqdsjyklmrxtioawcrvmsthbebdqqrpphncthosljfaeidboyekxezqtzlizqcvvxehrcskstshupglzgmbretpyehtavxegmbtznhpbczdjlzibnouxlxkeiedzoohoxhnhzqqaxdwetyudhyqvdhrggrszqeqkqqnunxqyyagyoptfkolieayokryidtctemtesuhbzczzvhlbbhnufjjocporuzuevofbuevuxhgexmckifntngaohfwqdakyobcooubdvypxjjxeugzdmapyamuwqtnqspsznyszhwqdqjxsmhdlkwkvlkdbjngvdmhvbllqqlcemkqxxdlldcfthjdqkyjrrjqqqpnmmelrwhtyugieuppqqtwychtpjmloxsckhzyitomjzypisxzztdwxhddvtv",
		}, {"axjtuwbmvsdzeogmnzorndhmjoqnqjnhmfwlueifbcqkezrwltzyeqvqemggctttilmfokzpgghxotfykyzendhtfapwfiermrovwtpzzrsyuiwongllqseumvpzmymtvsdsowammerobtvagmekpowndhejvbuwbporfyroknrjoekdgqqlgzxhiisfwevpeepgxmyhnajqlradgjtscciavbpgqjzecswtdetmtallzybcukdztovxysoavfggrejqkliirxuvnagwzmassthjecvfzkmyonglocmvjnxklecwqqvgrzpsswnigsrjthtmxyokuawirbkecfsauzrbifgwolpnhcapzxwmbydfhvsupfmuqapdxnetgunlqmundxvdsuqmzidlhateugmaqoyepbmztspgozzfjbrmpjgmwmhnldunsyokpvwprzrudbmtwbdvgyghgprqcdgqjjbyfsumjnnsskhmfqvjhnvcotynidziswxjqbfspzhkedshzbblustorxwtlhkowpavhngtbnuypvkmpajpwluamsxeeqqunlxxvfeutkyzayrolwezfzfyzmmneepwshpemynwzyunsxgjflnqmfghdsvwpjyknqhclhrlmnrljwabwpxtiomawcrvmsthuhffbebdqqrpfinphncthosljfaeirdblcaoyygghekxezqmotwrywzlizqaycvvgoxehmrcskstshujneqpglzugmbrxcetpnwdipldofyehtavfdurbsoxdurlofkqnrjoegmsbtzjimrxnhpbqczyadjlzakkizbnojwkuzcacnbdifxlxkesoiesmkbyffcdzoohoxhnhzqgqyhaxdwetyubtsrdhyqavdhrqagogrszqeqkqqnaxuznxqyyaggyoptfknolieayoksrfyidrmnoxrtctnemtngdxxechxrsbyuhtlbmzgmcqopyixdomhnmzzvnsafphpkdgndcsclbwybhnueytaefjjodlhlzczmpqqmnilliydwtxtporuzuedbncvsqauopfbvygqdtcwuevuxhffagexmyoalogeckifntngaceohnbfxlwqhdaklyobcooubdvypxfzmrjqofajxesvuzfczeuqegwpcmahhpzodsmpvrvkzxxtsdsxwixiraphjlqyamuwxiqtnlwfqspdlscdznyszhwtgqdqjpoiixbvsmpzilxrnphdlkwkvigpcclkdbjnngxvdmhvblzolqqlcentsmkqxxdlzldcfthjdqkyjrrjqqqpkxenmmelrwhtyugiemoluhppqifqtwyonbnizchtpjrlmuylezdkkztoxscphlmwhnkdguhelqzjgvjyitrzofmtjzypuhifoqnokonhqtsxzztdwxmimphddvtv"},
	}
	assertScsCases(t, tests)
}

func TestResult(t *testing.T) {
	values := []string{"jack", "apple", "maven", "hold", "solid", "mark", "moon", "poor", "spark", "live"}
	assertScsLike(t, values, "jmasppholivedcoarkn")
}

func TestBasicCases(t *testing.T) {
	tests := [][]string{
		{"moon", "poor"}, {"mpoonr"},
		{"moon", "spark", "mark", "poor"}, {"mspoonark"},
		{"moon", "poor", "spark", "mark"}, {"mspoonark"},
		{"spark", "apple"}, {"saparkple"},
	}
	assertScsCases(t, tests)
}

func TestRepeats(t *testing.T) {
	tests := [][]string{
		{"aaaa", "aaaaa", "aa"}, {"aaaaa"},
		{"baaa", "daaa"}, {"bdaaa"},
		{"baaaaa", "caaaaa", "daaa"}, {"bcdaaaaa"},
		{"baaaaa", "caaaaa"}, {"bcaaaaa"},
		{"baaa", "caaa", "daa"}, {"bcdaaa"},
		{"baaaa", "caaa"}, {"bcaaaa"},
	}
	assertScsCases(t, tests)
}

// --- Helper testing functions ---

func assertScsCases(t *testing.T, cases [][]string) {
	for i := 0; i < len(cases); i += 2 {
		assertScsLike(t, cases[i], cases[i+1][0])
		assertScsLikeOptions(t, cases[i], cases[i+1][0], ScsOptions{Parallel: true, Deterministic: false})
		assertScsLikeOptions(t, cases[i], cases[i+1][0], ScsOptions{Parallel: false, Deterministic: true})
		assertScsLikeOptions(t, cases[i], cases[i+1][0], ScsOptions{Parallel: true, Deterministic: true})
	}
}

func assertScsLike(t *testing.T, words []string, example string) string {
	return assertScsLikeOptions(t, words, example, ScsOptions{Parallel: false, Deterministic: false})
}

func assertScsLikeOptions(t *testing.T, words []string, example string, options ScsOptions) string {
	result := validatedScsOptions(t, words, options)
	if len(result) != len(example) {
		t.Logf("Expected length: %d, got %d\nWords: %v\nResult: %s", len(example), len(result), words, result)
		t.Fatalf("Invalid SCS result")
	}
	if valid, _ := validate(example, words); !valid {
		t.Fatalf("Invalid test case: example output %s for words %v is not valid", example, words)
	}
	return result
}

func validatedScs(t *testing.T, words []string) string {
	return validatedScsOptions(t, words, ScsOptions{Parallel: false, Deterministic: false})
}

func validatedScsOptions(t *testing.T, words []string, options ScsOptions) string {
	result := scs(words, options)
	if valid, invalid := validate(result, words); !valid {
		t.Logf("Invalid result: %s", result)
		t.Logf("Invalid words: %v", invalid)
		t.Fatalf("Invalid SCS result for words %v", words)
	}
	return result
}
