package cryptipass

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var w2 = []string{"ab", "ac", "ae", "af", "ag", "ah", "ai", "aj", "al", "am", "an", "ao", "ap", "aq", "ar", "as", "at", "au", "av", "aw", "ax", "ba", "bl", "bo", "br", "bu", "ca", "ce", "ch", "ci", "cl", "co", "cr", "cu", "cy", "da", "de", "di", "do", "dr", "du", "dw", "dy", "ea", "eb", "ec", "ed", "ef", "eg", "ei", "ej", "el", "em", "en", "ep", "eq", "er", "es", "et", "eu", "ev", "ex", "fa", "fe", "fi", "fl", "fo", "fr", "ga", "ge", "gi", "gl", "gn", "go", "gr", "gu", "gy", "ha", "he", "hu", "hy", "ib", "ic", "id", "ig", "il", "im", "io", "ip", "ir", "is", "it", "iv", "ja", "je", "ji", "jo", "ju", "ka", "ke", "ki", "kl", "kn", "ko", "kr", "ku", "la", "le", "li", "lu", "ly", "ma", "mi", "mo", "mu", "my", "na", "ne", "ni", "nu", "ny", "oa", "ob", "oc", "og", "oi", "ok", "ol", "om", "on", "op", "os", "ot", "ou", "ov", "ow", "ox", "oy", "oz", "pa", "pe", "ph", "pl", "po", "pr", "ps", "pu", "py", "qu", "ra", "re", "rh", "ri", "ro", "ru", "sa", "sc", "se", "sh", "si", "sk", "sl", "sm", "sn", "sp", "sq", "st", "su", "sw", "sy", "t-", "ta", "th", "ti", "tr", "tu", "tw", "ty", "ud", "ul", "um", "un", "up", "ur", "us", "ut", "va", "ve", "vi", "vo", "wa", "wh", "wi", "wo", "wr", "xb", "xe", "ya", "ye", "yo", "yu", "za", "ze", "zi", "zo"}

var tm = map[string]string{"nk": "@abehilmnrsy", "nu": "abcegilmnoprstz", "tw": "aehio", "et": "@abcdefhiorsuyz", "tl": "aeioy", "mh": "o", "lc": "ahioru", "cr": "aeiouy", "wa": "bcdfgiklmnprstvwy", "nh": "aeiou", "rj": "eou", "hm": "@ae", "sb": "aeou", "gl": "aeiouy", "br": "aeiou", "an": "@abcdefghijklmopqstuvyz", "sp": "@aehilnoruy", "mv": "e", "kr": "eoy", "bi": "@acdeglnorstz", "tp": "aloru", "ly": "@aegmoprswz", "bf": "l", "te": "@abcdglmnprstwx", "il": "@abcdehikmnorstuvwy", "cy": "@clmt", "ka": "@bcglnrty", "lr": "iouy", "rv": "aeioy", "ej": "eo", "yf": "ilou", "ko": "@as", "mo": "@abcdgiklmnprstuvw", "uz": "e", "vo": "ciklrtuwy", "em": "@abceilnopsuy", "tc": "ahlou", "df": "aiou", "us": "@abcehiklptuy", "um": "@abdeinopsv", "wb": "aeio", "sy": "@bcmnrs", "mc": "eho", "hn": "@ae", "cq": "u", "kw": "aho", "ni": "@abcdefglmnopqstuvxz", "ho": "@bcegiklmnprstuvwy", "ix": "@aefit", "pe": "@abcdglnprstwz", "eo": "@cdglmnprtuv", "ay": "@abcdefghlmoprstw", "my": "@st", "oy": "@aeflmnors", "zm": "o", "rn": "@abcefhilmosuy", "pk": "i", "lu": "abcdegiklmnrstx", "ji": "fgmnt", "or": "@abcdefgiklmnopstwy", "ub": "@acdefghijlmprstuwyz", "bo": "@abdefgijklnrstuvwxy", "bs": "@ceilotuy", "ja": "@bcilmnrsuvwyz", "hd": "a", "if": "eilotuy", "ta": "@bcdefgiklmnprstuvwxy", "be": "@acdlnrstyz", "kd": "r", "lo": "@abcdefghimnpqrstuvwy", "bu": "bcdgilmnprstz", "bl": "aeiouy", "le": "@abcdfghilmnoprstuvwxy", "kh": "ao", "ir": "@acdeiklmopstuy", "ha": "@bcdfgiklmnoprstuvwz", "yz": "e", "au": "bcdglnprstvz", "ze": "@abdlnprs", "bz": "e", "sk": "@abeilntwy", "so": "@abcdfilmnprtuvw", "zo": "adimn", "yg": "elor", "ud": "@aegiosy", "xb": "o", "uh": "@", "at": "@abcefhilnorsuwy", "ts": "@cehikmopuy", "uk": "e", "rm": "@abcefhilopruy", "gh": "@aity", "it": "@acehilmorsuyz", "pl": "aeiouy", "yu": "m", "nv": "aeioy", "di": "@abcdefglmnoprstuvwxz", "lp": "@aefhilort", "bj": "e", "de": "@abcdfghijlmnoprstuvx", "wi": "@cdefgklmnprstvz", "xe": "@dmnrs", "eg": "@abeilmnoruwy", "fy": "@", "oh": "aen", "-y": "o", "hp": "ir", "pw": "ai", "ba": "@bcdfgklmnrstuyz", "bm": "aei", "kf": "i", "hb": "aou", "gb": "aeo", "oq": "u", "ri": "@abcdefgklmnopstuvxz", "lv": "aei", "ya": "@bghlmnprw", "fr": "aeiouy", "ua": "@bcdghiklnrstv", "co": "@abcdefghiklmnprstuvwyz", "gz": "a", "ny": "@bhlmoptwx", "sn": "aeiou", "se": "@acdfgilmnpqrstuvwxy", "rt": "@acefhilnorsuwyz", "wg": "i", "ry": "@bdopsw", "dg": "eiruy", "oa": "@bcdfgklmnrstu", "hi": "@acdefghlmnprstv", "dp": "ahior", "yp": "aehilnotu", "oj": "e", "al": "@abcdefgikmnoprstuvwyz", "uy": "@", "uj": "i", "ur": "@abcdefgiklmnopstuvy", "ez": "aeioy", "by": "@ht", "p-": "d", "ht": "@efilnruwy", "ig": "@aehilmnorsuyz", "wc": "a", "nf": "aeiloru", "ei": "@glmnstvz", "rc": "aehilortu", "gn": "@aeio", "jo": "@bcghiklrtvy", "dc": "aloru", "fo": "acdegilnrsuwxy", "lm": "@aeioy", "rx": "i", "hc": "ahl", "bg": "r", "do": "@acdefgiklmnrstuvwxz", "ak": "@aeiy", "kn": "aeio", "ok": "@aeiuy", "fn": "e", "-d": "o", "lb": "@aeioru", "el": "@abcdefikmnopstuvy", "yc": "ahlo", "oi": "cdlnrst", "ky": "@adlprw", "ag": "@aefhilmnoprsuw", "nb": "aeiloru", "rl": "@adeioy", "sj": "o", "ip": "@acefhilmostu", "yn": "adeot", "vr": "o", "yw": "ahio", "ew": "@adehilmnopry", "aw": "@adefhiklnors", "hr": "aeiou", "bh": "eo", "aj": "aeo", "gy": "@mr", "mt": "@", "ti": "@abcdefghlmnopqrstvz", "id": "@abegilnostuy", "gs": "@ahit", "op": "-@abcehiklostuwy", "ch": "@abeiklmnortuwy", "qu": "aeio", "up": "@abcdefghilorstwy", "xo": "dnpr", "ey": "@abehlmow", "mr": "aeo", "zy": "@m", "xy": "@gm", "rg": "aeiloruy", "km": "a", "ug": "@aehilmnosu", "ln": "@eu", "ai": "@cdklmnrst", "hu": "abcdglmnprst", "ec": "aehiklortuy", "si": "abcdefglmnoprstuvxz", "ru": "@bcdegilmnprstx", "oc": "aehiklortuy", "pg": "r", "lw": "ah", "wt": "i", "mu": "@cdglmnprstz", "ac": "@aehikloqrtuy", "ov": "aeioy", "uv": "e", "ut": "@abcdefghilmnoprsuwy", "gu": "aeilmnoprstyz", "ex": "@acefhiopqtu", "ax": "@aeios", "ra": "@bcdefgiklmnoprstuvwxyz", "kb": "aeio", "uo": "@tu", "gf": "u", "uc": "aehiklortuy", "go": "@adegilmnprstuvw", "sq": "u", "ct": "@aefilmorsu", "gm": "@ae", "np": "aeilor", "wh": "aeioy", "ow": "@abcdefgilmnoprsty", "yx": "@", "ws": "@ehitu", "la": "@bcdghikmnprstuvwxyz", "va": "@bcdgilnprst", "fl": "aeiouy", "pu": "bcdeglmnprstz", "pn": "eo", "oz": "aeioy", "rb": "@aeilorsuy", "ul": "@abcdefgikmopstuvy", "xq": "u", "am": "@abceilnoprstuy", "nj": "aou", "we": "abdilnprs", "bt": "aeiloruy", "dt": "h", "pr": "aeiouy", "ge": "@abcdfilmnorsty", "bc": "a", "xf": "o", "tb": "aeioru", "re": "@abcdfghijklmnopqrstuvwxy", "bv": "i", "ym": "@abeop", "im": "@abeilnopsuwy", "yb": "aeioru", "sg": "r", "yo": "-@dfgnruvy", "ar": "@abcdefghiklmnopstvwxy", "tz": "@ey", "tm": "aeou", "lf": "@aioru", "xp": "aeiloru", "ft": "@ehilnoswy", "bp": "alr", "wl": "@ei", "sd": "ao", "sf": "eioruy", "mn": "aei", "py": "@grt", "as": "@acehiklmopqtuy", "lg": "eiou", "to": "@abcdefgiklmnprstuvwxz", "ed": "@aegilortuy", "os": "@aehilmoptuy", "ot": "@abceghilmnoprsuwy", "md": "r", "dw": "aeior", "ty": "@cfhklp", "eq": "u", "yl": "eiou", "yd": "air", "sr": "eu", "nt": "@adehilmorsuwy", "wk": "@w", "ms": "@hiotuy", "eh": "aeioy", "xi": "acdfglmnorst", "ce": "@abcdfilmnprst", "in": "@acdefghijklopstuvwxy", "ks": "@ahiklpt", "-t": "i", "mi": "@acdefglmnrstuxz", "ju": "bdgijklmnrsv", "fs": "@", "ga": "@bcdefghilmnprstuvwz", "pi": "acdeflnoprstu", "ds": "@acefhilmot", "ps": "@aceituwy", "xh": "au", "rd": "@abcehilnorsuwy", "dv": "i", "lk": "@ae", "ik": "aeiu", "tg": "eor", "nw": "aeior", "yt": "ehio", "th": "@abeilmopruwy", "eu": "cmnprst", "sw": "aeiou", "xs": "e", "o-": "y", "pb": "aeo", "pt": "@aehilouy", "az": "aeioy", "wd": "@elor", "nx": "@", "mf": "ouy", "wu": "n", "pm": "e", "oe": "@dimnrstvxy", "iu": "m", "db": "ailor", "ao": "krs", "fu": "elmnrst", "dk": "i", "nc": "aehilortuy", "es": "@acdehiklmopqtuy", "ls": "@aeit", "nr": "aeiou", "je": "clrst", "ia": "@bcglmnprst", "ol": "@adefikopstuvy", "ns": "@acefhiklmnoptuw", "za": "@bcgpr", "su": "@abcdegilmoprs", "mp": "@abefhilnorstu", "fa": "@bcdilmnrstuvxz", "af": "@aelnort", "iz": "@aeim", "no": "@begilmnprstuvwxy", "gr": "aeiouy", "un": "@abcdefghijklmopqrstuvwz", "lt": "-@aehioruyz", "en": "@acdefghijklopqrstuvyz", "xa": "bcglmt", "hs": "t", "st": "@abcefilnoruwy", "wr": "aeioy", "ab": "@adegilmnorsuy", "ca": "@bcdfghklmnprstuvwy", "bn": "eo", "li": "@abcdefgklmnopqrstuvxz", "ie": "@cdfklmnrstuvw", "dh": "aeio", "kp": "aeo", "hk": "i", "dn": "eo", "iv": "aeioy", "yr": "aiou", "po": "acdefgiklnprstuwx", "t-": "st", "lh": "o", "rf": "@abeiloru", "ae": "@kr", "-s": "h", "ne": "@abcdghlmnopqrstuvwxy", "lz": "o", "sa": "@bcdfgiklmnprstuvwxy", "na": "@bcdefgiklmnprstuvwz", "gp": "ilo", "cu": "abcdelmprst", "ma": "@bcdghijklmnprstuvxyz", "du": "abcdeghiklmoprstv", "rh": "aeuy", "fe": "@acdghilmnrstvw", "ve": "@adghilmnrstwxy", "ux": "@eu", "hy": "@bdmps", "tn": "aeiou", "zi": "ceglnpt", "rk": "@aeinrwy", "fb": "ao", "nz": "aeiy", "hw": "ao", "mw": "e", "ck": "@abdefhilnoprstuwy", "xc": "aehilru", "ue": "@abdflnrstu", "gd": "o", "ke": "@bdghlmnoprstwy", "eb": "@aeilortu", "of": "abeiorstuy", "ap": "@adehiklnorstuy", "ld": "@acefhilmnosy", "nl": "aeiouy", "is": "@abcdefghijklmoprtwy", "gi": "abcdefglmnorsvz", "tu": "@abcdegilmnoprstx", "pc": "aho", "rp": "@aehilnory", "xl": "i", "ki": "cdelmnprstw", "wm": "ae", "wy": "@", "sl": "aeiouy", "ox": "@cefily", "iq": "u", "wn": "@eiy", "ad": "@abcefgilmnoprsuvwy", "on": "@abcdefgijklorstuvwyz", "od": "@aegilouy", "pf": "iru", "ob": "@aceijlnostuvwy", "cl": "aeiou", "wf": "ailu", "dr": "aeiouy", "ef": "@aeilortuy", "io": "@cdlmnrtux", "ek": "@iw", "da": "@bcdfgilmnrstuvwyz", "wo": "bfklmnruvw", "av": "aeioy", "ci": "abdeflmnoprstuvz", "tf": "ilou", "ye": "@adlnrs", "pa": "@bcdgijlmnprstuvwy", "nd": "@abcefghiklmnoprsuwy", "ah": "@eou", "vy": "@", "rw": "aeir", "tr": "aeiouy", "vi": "abcdeglnoprstvx", "xu": "abr", "er": "@abcdefghijklmnopstuvwy", "gw": "aeho", "aq": "u", "pd": "ao", "ui": "cdeglnprstvz", "dl": "aeioy", "td": "ao", "om": "@abefinoprsy", "bw": "aeo", "dy": "@blmns", "mb": "@aeilnoru", "sh": "@abcdeilnoprstuy", "ep": "@aehilnorstu", "ib": "@aceiloru", "ou": "bcdglnprst", "hl": "aeioy", "ph": "@aeiloruy", "kt": "aor", "kl": "aeioy", "bd": "ioru", "sm": "@aeiou", "ys": "@eilpt", "ml": "aeioy", "wp": "ilo", "iw": "io", "sc": "aehiloru", "fi": "@abcdefglnrstvx", "ih": "eu", "yh": "o", "me": "@adglmnorstv", "yk": "e", "dm": "aei", "gt": "h", "ku": "@dnps", "nm": "aeio", "og": "@aeilnoruwy", "ng": "@abdefhilmnorstuwy", "rs": "@adehilmnoptuw", "ic": "@aehiklortuy", "xt": "@ehiory", "ea": "@bcdfghklmnprstvw", "ro": "@abcdfgijklmnprstuvwxyz", "he": "@acdflmnorstvwx", "ev": "aeiory", "nq": "u"}

func cryptoInt64[T int | int64 | uint | uint64](max T) int64 {
	LEN := big.NewInt(int64(max))
	nBig, err := rand.Int(rand.Reader, LEN)
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

func NewPassphrase(words uint64) string {
	tokens := []string{}
	for range words {
		tokens = append(tokens, GenMixWord())
	}
	return strings.Join(tokens, ".")
}

func GenMixWord() string {
	D := []int{3, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7}
	// this mixture is certified for entropy > 20.25 bits
	n := cryptoInt64(len(D))
	return GenWord(D[n])
}

func GenWord(n int) string {
repeat:
	s := w2[cryptoInt64(len(w2))]
	for {
		next := tm[s[len(s)-2:]]
		if len(next) == 0 {
			goto repeat
		}
		c := next[cryptoInt64(len(next))]
		if c == '@' {
			break
		}
		s += string(c)
		if len(s) > n {
			goto repeat
		}
	}
	if len(s) < n {
		goto repeat
	}
	return s
}
