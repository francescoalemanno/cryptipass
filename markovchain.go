package cryptipass

// THIS FILE HAS BEEN DISTILLED FROM EFF's long word list, without their work this software would not exist.

func PickNext(seed string) (string, float64) {
	L := min(len(seed), 2)
	tok := seed[len(seed)-L:]
	if tok == "nu" {
		r := rng.IntN(134)
		H := 3.8145722608727763
		if 0 <= r && r < 8 {
			return seed + "a", H
		} else if 8 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 18 {
			return seed + "c", H
		} else if 18 <= r && r < 19 {
			return seed + "d", H
		} else if 19 <= r && r < 25 {
			return seed + "e", H
		} else if 25 <= r && r < 28 {
			return seed + "f", H
		} else if 28 <= r && r < 37 {
			return seed + "g", H
		} else if 37 <= r && r < 38 {
			return seed + "h", H
		} else if 38 <= r && r < 40 {
			return seed + "i", H
		} else if 40 <= r && r < 41 {
			return seed + "j", H
		} else if 41 <= r && r < 42 {
			return seed + "k", H
		} else if 42 <= r && r < 47 {
			return seed + "l", H
		} else if 47 <= r && r < 76 {
			return seed + "m", H
		} else if 76 <= r && r < 79 {
			return seed + "n", H
		} else if 79 <= r && r < 83 {
			return seed + "o", H
		} else if 83 <= r && r < 86 {
			return seed + "p", H
		} else if 86 <= r && r < 87 {
			return seed + "q", H
		} else if 87 <= r && r < 94 {
			return seed + "r", H
		} else if 94 <= r && r < 105 {
			return seed + "s", H
		} else if 105 <= r && r < 128 {
			return seed + "t", H
		} else if 128 <= r && r < 129 {
			return seed + "v", H
		} else if 129 <= r && r < 130 {
			return seed + "w", H
		} else if 130 <= r && r < 131 {
			return seed + "x", H
		} else if 131 <= r && r < 134 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mh" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tl" {
		r := rng.IntN(196)
		H := 1.827737619583285
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 90 {
			return seed + "e", H
		} else if 90 <= r && r < 113 {
			return seed + "i", H
		} else if 113 <= r && r < 118 {
			return seed + "o", H
		} else if 118 <= r && r < 119 {
			return seed + "u", H
		} else if 119 <= r && r < 196 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rj" {
		r := rng.IntN(14)
		H := 2.2988252450030506
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sp" {
		r := rng.IntN(326)
		H := 2.8383219475828256
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 90 {
			return seed + "e", H
		} else if 90 <= r && r < 104 {
			return seed + "h", H
		} else if 104 <= r && r < 163 {
			return seed + "i", H
		} else if 163 <= r && r < 197 {
			return seed + "l", H
		} else if 197 <= r && r < 199 {
			return seed + "n", H
		} else if 199 <= r && r < 276 {
			return seed + "o", H
		} else if 276 <= r && r < 310 {
			return seed + "r", H
		} else if 310 <= r && r < 319 {
			return seed + "u", H
		} else if 319 <= r && r < 326 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zn" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ly" {
		r := rng.IntN(72)
		H := 4.061860684444236
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 7 {
			return seed + "d", H
		} else if 7 <= r && r < 11 {
			return seed + "e", H
		} else if 11 <= r && r < 12 {
			return seed + "f", H
		} else if 12 <= r && r < 17 {
			return seed + "g", H
		} else if 17 <= r && r < 18 {
			return seed + "h", H
		} else if 18 <= r && r < 28 {
			return seed + "i", H
		} else if 28 <= r && r < 29 {
			return seed + "j", H
		} else if 29 <= r && r < 30 {
			return seed + "k", H
		} else if 30 <= r && r < 31 {
			return seed + "l", H
		} else if 31 <= r && r < 36 {
			return seed + "m", H
		} else if 36 <= r && r < 37 {
			return seed + "n", H
		} else if 37 <= r && r < 39 {
			return seed + "o", H
		} else if 39 <= r && r < 42 {
			return seed + "p", H
		} else if 42 <= r && r < 43 {
			return seed + "q", H
		} else if 43 <= r && r < 52 {
			return seed + "r", H
		} else if 52 <= r && r < 59 {
			return seed + "s", H
		} else if 59 <= r && r < 60 {
			return seed + "t", H
		} else if 60 <= r && r < 61 {
			return seed + "v", H
		} else if 61 <= r && r < 64 {
			return seed + "w", H
		} else if 64 <= r && r < 65 {
			return seed + "x", H
		} else if 65 <= r && r < 72 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "il" {
		r := rng.IntN(504)
		H := 3.0813012570452143
		if 0 <= r && r < 25 {
			return seed + "a", H
		} else if 25 <= r && r < 29 {
			return seed + "b", H
		} else if 29 <= r && r < 31 {
			return seed + "c", H
		} else if 31 <= r && r < 61 {
			return seed + "d", H
		} else if 61 <= r && r < 144 {
			return seed + "e", H
		} else if 144 <= r && r < 146 {
			return seed + "h", H
		} else if 146 <= r && r < 225 {
			return seed + "i", H
		} else if 225 <= r && r < 227 {
			return seed + "k", H
		} else if 227 <= r && r < 341 {
			return seed + "l", H
		} else if 341 <= r && r < 343 {
			return seed + "m", H
		} else if 343 <= r && r < 345 {
			return seed + "n", H
		} else if 345 <= r && r < 366 {
			return seed + "o", H
		} else if 366 <= r && r < 368 {
			return seed + "r", H
		} else if 368 <= r && r < 372 {
			return seed + "s", H
		} else if 372 <= r && r < 396 {
			return seed + "t", H
		} else if 396 <= r && r < 401 {
			return seed + "u", H
		} else if 401 <= r && r < 403 {
			return seed + "v", H
		} else if 403 <= r && r < 405 {
			return seed + "w", H
		} else if 405 <= r && r < 504 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "te" {
		r := rng.IntN(786)
		H := 2.5146934358999844
		if 0 <= r && r < 14 {
			return seed + "a", H
		} else if 14 <= r && r < 19 {
			return seed + "b", H
		} else if 19 <= r && r < 28 {
			return seed + "c", H
		} else if 28 <= r && r < 235 {
			return seed + "d", H
		} else if 235 <= r && r < 271 {
			return seed + "e", H
		} else if 271 <= r && r < 272 {
			return seed + "f", H
		} else if 272 <= r && r < 281 {
			return seed + "g", H
		} else if 281 <= r && r < 282 {
			return seed + "h", H
		} else if 282 <= r && r < 283 {
			return seed + "j", H
		} else if 283 <= r && r < 284 {
			return seed + "k", H
		} else if 284 <= r && r < 297 {
			return seed + "l", H
		} else if 297 <= r && r < 314 {
			return seed + "m", H
		} else if 314 <= r && r < 409 {
			return seed + "n", H
		} else if 409 <= r && r < 418 {
			return seed + "p", H
		} else if 418 <= r && r < 419 {
			return seed + "q", H
		} else if 419 <= r && r < 756 {
			return seed + "r", H
		} else if 756 <= r && r < 765 {
			return seed + "s", H
		} else if 765 <= r && r < 770 {
			return seed + "t", H
		} else if 770 <= r && r < 771 {
			return seed + "v", H
		} else if 771 <= r && r < 774 {
			return seed + "w", H
		} else if 774 <= r && r < 785 {
			return seed + "x", H
		} else if 785 <= r && r < 786 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "rv" {
		r := rng.IntN(74)
		H := 2.0487714333815044
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 61 {
			return seed + "i", H
		} else if 61 <= r && r < 66 {
			return seed + "o", H
		} else if 66 <= r && r < 67 {
			return seed + "u", H
		} else if 67 <= r && r < 74 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vo" {
		r := rng.IntN(134)
		H := 3.466667151925312
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 17 {
			return seed + "d", H
		} else if 17 <= r && r < 18 {
			return seed + "f", H
		} else if 18 <= r && r < 19 {
			return seed + "g", H
		} else if 19 <= r && r < 20 {
			return seed + "h", H
		} else if 20 <= r && r < 28 {
			return seed + "i", H
		} else if 28 <= r && r < 29 {
			return seed + "j", H
		} else if 29 <= r && r < 38 {
			return seed + "k", H
		} else if 38 <= r && r < 61 {
			return seed + "l", H
		} else if 61 <= r && r < 62 {
			return seed + "m", H
		} else if 62 <= r && r < 63 {
			return seed + "n", H
		} else if 63 <= r && r < 64 {
			return seed + "p", H
		} else if 64 <= r && r < 65 {
			return seed + "q", H
		} else if 65 <= r && r < 98 {
			return seed + "r", H
		} else if 98 <= r && r < 99 {
			return seed + "s", H
		} else if 99 <= r && r < 110 {
			return seed + "t", H
		} else if 110 <= r && r < 122 {
			return seed + "u", H
		} else if 122 <= r && r < 123 {
			return seed + "v", H
		} else if 123 <= r && r < 128 {
			return seed + "w", H
		} else if 128 <= r && r < 129 {
			return seed + "x", H
		} else if 129 <= r && r < 133 {
			return seed + "y", H
		} else if 133 <= r && r < 134 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "tc" {
		r := rng.IntN(124)
		H := 1.2552568881156503
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 14 {
			return seed + "e", H
		} else if 14 <= r && r < 110 {
			return seed + "h", H
		} else if 110 <= r && r < 111 {
			return seed + "i", H
		} else if 111 <= r && r < 113 {
			return seed + "l", H
		} else if 113 <= r && r < 120 {
			return seed + "o", H
		} else if 120 <= r && r < 123 {
			return seed + "u", H
		} else if 123 <= r && r < 124 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "us" {
		r := rng.IntN(380)
		H := 2.862733083228781
		if 0 <= r && r < 29 {
			return seed + "a", H
		} else if 29 <= r && r < 33 {
			return seed + "b", H
		} else if 33 <= r && r < 35 {
			return seed + "c", H
		} else if 35 <= r && r < 118 {
			return seed + "e", H
		} else if 118 <= r && r < 184 {
			return seed + "h", H
		} else if 184 <= r && r < 229 {
			return seed + "i", H
		} else if 229 <= r && r < 247 {
			return seed + "k", H
		} else if 247 <= r && r < 251 {
			return seed + "l", H
		} else if 251 <= r && r < 252 {
			return seed + "o", H
		} else if 252 <= r && r < 260 {
			return seed + "p", H
		} else if 260 <= r && r < 266 {
			return seed + "s", H
		} else if 266 <= r && r < 370 {
			return seed + "t", H
		} else if 370 <= r && r < 375 {
			return seed + "u", H
		} else if 375 <= r && r < 380 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kw" {
		r := rng.IntN(22)
		H := 2.1110957877533387
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 14 {
			return seed + "h", H
		} else if 14 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ix" {
		r := rng.IntN(30)
		H := 2.532380816279281
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 17 {
			return seed + "i", H
		} else if 17 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 28 {
			return seed + "t", H
		} else if 28 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 30 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ji" {
		r := rng.IntN(38)
		H := 3.8382530025147203
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 9 {
			return seed + "g", H
		} else if 9 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 11 {
			return seed + "j", H
		} else if 11 <= r && r < 12 {
			return seed + "k", H
		} else if 12 <= r && r < 13 {
			return seed + "l", H
		} else if 13 <= r && r < 16 {
			return seed + "m", H
		} else if 16 <= r && r < 23 {
			return seed + "n", H
		} else if 23 <= r && r < 24 {
			return seed + "p", H
		} else if 24 <= r && r < 25 {
			return seed + "q", H
		} else if 25 <= r && r < 26 {
			return seed + "r", H
		} else if 26 <= r && r < 27 {
			return seed + "s", H
		} else if 27 <= r && r < 34 {
			return seed + "t", H
		} else if 34 <= r && r < 35 {
			return seed + "v", H
		} else if 35 <= r && r < 36 {
			return seed + "w", H
		} else if 36 <= r && r < 37 {
			return seed + "x", H
		} else if 37 <= r && r < 38 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "or" {
		r := rng.IntN(586)
		H := 3.7747091124267143
		if 0 <= r && r < 51 {
			return seed + "a", H
		} else if 51 <= r && r < 59 {
			return seed + "b", H
		} else if 59 <= r && r < 73 {
			return seed + "c", H
		} else if 73 <= r && r < 99 {
			return seed + "d", H
		} else if 99 <= r && r < 172 {
			return seed + "e", H
		} else if 172 <= r && r < 176 {
			return seed + "f", H
		} else if 176 <= r && r < 182 {
			return seed + "g", H
		} else if 182 <= r && r < 237 {
			return seed + "i", H
		} else if 237 <= r && r < 271 {
			return seed + "k", H
		} else if 271 <= r && r < 275 {
			return seed + "l", H
		} else if 275 <= r && r < 321 {
			return seed + "m", H
		} else if 321 <= r && r < 373 {
			return seed + "n", H
		} else if 373 <= r && r < 390 {
			return seed + "o", H
		} else if 390 <= r && r < 404 {
			return seed + "p", H
		} else if 404 <= r && r < 434 {
			return seed + "r", H
		} else if 434 <= r && r < 458 {
			return seed + "s", H
		} else if 458 <= r && r < 558 {
			return seed + "t", H
		} else if 558 <= r && r < 559 {
			return seed + "u", H
		} else if 559 <= r && r < 561 {
			return seed + "w", H
		} else if 561 <= r && r < 586 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "g" {
		r := rng.IntN(614)
		H := 2.7100044377774277
		if 0 <= r && r < 91 {
			return seed + "a", H
		} else if 91 <= r && r < 158 {
			return seed + "e", H
		} else if 158 <= r && r < 207 {
			return seed + "i", H
		} else if 207 <= r && r < 293 {
			return seed + "l", H
		} else if 293 <= r && r < 297 {
			return seed + "n", H
		} else if 297 <= r && r < 362 {
			return seed + "o", H
		} else if 362 <= r && r < 556 {
			return seed + "r", H
		} else if 556 <= r && r < 611 {
			return seed + "u", H
		} else if 611 <= r && r < 614 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "if" {
		r := rng.IntN(218)
		H := 2.648887261832668
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 44 {
			return seed + "f", H
		} else if 44 <= r && r < 93 {
			return seed + "i", H
		} else if 93 <= r && r < 101 {
			return seed + "l", H
		} else if 101 <= r && r < 108 {
			return seed + "o", H
		} else if 108 <= r && r < 162 {
			return seed + "t", H
		} else if 162 <= r && r < 169 {
			return seed + "u", H
		} else if 169 <= r && r < 218 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lo" {
		r := rng.IntN(426)
		H := 4.053699587749044
		if 0 <= r && r < 34 {
			return seed + "a", H
		} else if 34 <= r && r < 45 {
			return seed + "b", H
		} else if 45 <= r && r < 72 {
			return seed + "c", H
		} else if 72 <= r && r < 81 {
			return seed + "d", H
		} else if 81 <= r && r < 83 {
			return seed + "e", H
		} else if 83 <= r && r < 86 {
			return seed + "f", H
		} else if 86 <= r && r < 139 {
			return seed + "g", H
		} else if 139 <= r && r < 142 {
			return seed + "h", H
		} else if 142 <= r && r < 146 {
			return seed + "i", H
		} else if 146 <= r && r < 147 {
			return seed + "j", H
		} else if 147 <= r && r < 148 {
			return seed + "k", H
		} else if 148 <= r && r < 149 {
			return seed + "l", H
		} else if 149 <= r && r < 154 {
			return seed + "m", H
		} else if 154 <= r && r < 193 {
			return seed + "n", H
		} else if 193 <= r && r < 219 {
			return seed + "o", H
		} else if 219 <= r && r < 242 {
			return seed + "p", H
		} else if 242 <= r && r < 247 {
			return seed + "q", H
		} else if 247 <= r && r < 292 {
			return seed + "r", H
		} else if 292 <= r && r < 313 {
			return seed + "s", H
		} else if 313 <= r && r < 346 {
			return seed + "t", H
		} else if 346 <= r && r < 366 {
			return seed + "u", H
		} else if 366 <= r && r < 381 {
			return seed + "v", H
		} else if 381 <= r && r < 412 {
			return seed + "w", H
		} else if 412 <= r && r < 413 {
			return seed + "x", H
		} else if 413 <= r && r < 425 {
			return seed + "y", H
		} else if 425 <= r && r < 426 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "so" {
		r := rng.IntN(136)
		H := 3.4570528630012376
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 11 {
			return seed + "d", H
		} else if 11 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 20 {
			return seed + "i", H
		} else if 20 <= r && r < 21 {
			return seed + "j", H
		} else if 21 <= r && r < 22 {
			return seed + "k", H
		} else if 22 <= r && r < 55 {
			return seed + "l", H
		} else if 55 <= r && r < 58 {
			return seed + "m", H
		} else if 58 <= r && r < 89 {
			return seed + "n", H
		} else if 89 <= r && r < 91 {
			return seed + "o", H
		} else if 91 <= r && r < 94 {
			return seed + "p", H
		} else if 94 <= r && r < 95 {
			return seed + "q", H
		} else if 95 <= r && r < 118 {
			return seed + "r", H
		} else if 118 <= r && r < 119 {
			return seed + "s", H
		} else if 119 <= r && r < 124 {
			return seed + "t", H
		} else if 124 <= r && r < 128 {
			return seed + "u", H
		} else if 128 <= r && r < 131 {
			return seed + "v", H
		} else if 131 <= r && r < 134 {
			return seed + "w", H
		} else if 134 <= r && r < 135 {
			return seed + "x", H
		} else if 135 <= r && r < 136 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xb" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "at" {
		r := rng.IntN(1018)
		H := 2.988529217959573
		if 0 <= r && r < 39 {
			return seed + "a", H
		} else if 39 <= r && r < 45 {
			return seed + "b", H
		} else if 45 <= r && r < 105 {
			return seed + "c", H
		} else if 105 <= r && r < 470 {
			return seed + "e", H
		} else if 470 <= r && r < 478 {
			return seed + "f", H
		} else if 478 <= r && r < 528 {
			return seed + "h", H
		} else if 528 <= r && r < 707 {
			return seed + "i", H
		} else if 707 <= r && r < 719 {
			return seed + "l", H
		} else if 719 <= r && r < 727 {
			return seed + "n", H
		} else if 727 <= r && r < 820 {
			return seed + "o", H
		} else if 820 <= r && r < 858 {
			return seed + "r", H
		} else if 858 <= r && r < 868 {
			return seed + "s", H
		} else if 868 <= r && r < 964 {
			return seed + "t", H
		} else if 964 <= r && r < 1007 {
			return seed + "u", H
		} else if 1007 <= r && r < 1013 {
			return seed + "w", H
		} else if 1013 <= r && r < 1018 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "a" {
		r := rng.IntN(834)
		H := 3.861741049308229
		if 0 <= r && r < 49 {
			return seed + "b", H
		} else if 49 <= r && r < 110 {
			return seed + "c", H
		} else if 110 <= r && r < 111 {
			return seed + "d", H
		} else if 111 <= r && r < 119 {
			return seed + "e", H
		} else if 119 <= r && r < 166 {
			return seed + "f", H
		} else if 166 <= r && r < 207 {
			return seed + "g", H
		} else if 207 <= r && r < 212 {
			return seed + "h", H
		} else if 212 <= r && r < 218 {
			return seed + "i", H
		} else if 218 <= r && r < 221 {
			return seed + "j", H
		} else if 221 <= r && r < 222 {
			return seed + "k", H
		} else if 222 <= r && r < 289 {
			return seed + "l", H
		} else if 289 <= r && r < 368 {
			return seed + "m", H
		} else if 368 <= r && r < 517 {
			return seed + "n", H
		} else if 517 <= r && r < 519 {
			return seed + "o", H
		} else if 519 <= r && r < 578 {
			return seed + "p", H
		} else if 578 <= r && r < 583 {
			return seed + "q", H
		} else if 583 <= r && r < 654 {
			return seed + "r", H
		} else if 654 <= r && r < 701 {
			return seed + "s", H
		} else if 701 <= r && r < 746 {
			return seed + "t", H
		} else if 746 <= r && r < 784 {
			return seed + "u", H
		} else if 784 <= r && r < 811 {
			return seed + "v", H
		} else if 811 <= r && r < 830 {
			return seed + "w", H
		} else if 830 <= r && r < 833 {
			return seed + "x", H
		} else if 833 <= r && r < 834 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "cj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uk" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zs" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lp" {
		r := rng.IntN(44)
		H := 3.2406667798772903
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 22 {
			return seed + "h", H
		} else if 22 <= r && r < 29 {
			return seed + "i", H
		} else if 29 <= r && r < 33 {
			return seed + "l", H
		} else if 33 <= r && r < 36 {
			return seed + "o", H
		} else if 36 <= r && r < 38 {
			return seed + "r", H
		} else if 38 <= r && r < 42 {
			return seed + "t", H
		} else if 42 <= r && r < 43 {
			return seed + "u", H
		} else if 43 <= r && r < 44 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bj" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "eg" {
		r := rng.IntN(146)
		H := 3.2347050812892975
		if 0 <= r && r < 35 {
			return seed + "a", H
		} else if 35 <= r && r < 37 {
			return seed + "b", H
		} else if 37 <= r && r < 44 {
			return seed + "e", H
		} else if 44 <= r && r < 62 {
			return seed + "g", H
		} else if 62 <= r && r < 83 {
			return seed + "i", H
		} else if 83 <= r && r < 89 {
			return seed + "l", H
		} else if 89 <= r && r < 91 {
			return seed + "m", H
		} else if 91 <= r && r < 95 {
			return seed + "n", H
		} else if 95 <= r && r < 110 {
			return seed + "o", H
		} else if 110 <= r && r < 126 {
			return seed + "r", H
		} else if 126 <= r && r < 139 {
			return seed + "u", H
		} else if 139 <= r && r < 143 {
			return seed + "w", H
		} else if 143 <= r && r < 146 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hb" {
		r := rng.IntN(28)
		H := 1.9194126680510324
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 24 {
			return seed + "o", H
		} else if 24 <= r && r < 27 {
			return seed + "u", H
		} else if 27 <= r && r < 28 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oq" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "aa" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ri" {
		r := rng.IntN(988)
		H := 3.8566232047586393
		if 0 <= r && r < 36 {
			return seed + "a", H
		} else if 36 <= r && r < 69 {
			return seed + "b", H
		} else if 69 <= r && r < 174 {
			return seed + "c", H
		} else if 174 <= r && r < 223 {
			return seed + "d", H
		} else if 223 <= r && r < 289 {
			return seed + "e", H
		} else if 289 <= r && r < 322 {
			return seed + "f", H
		} else if 322 <= r && r < 359 {
			return seed + "g", H
		} else if 359 <= r && r < 360 {
			return seed + "h", H
		} else if 360 <= r && r < 361 {
			return seed + "j", H
		} else if 361 <= r && r < 368 {
			return seed + "k", H
		} else if 368 <= r && r < 409 {
			return seed + "l", H
		} else if 409 <= r && r < 462 {
			return seed + "m", H
		} else if 462 <= r && r < 651 {
			return seed + "n", H
		} else if 651 <= r && r < 677 {
			return seed + "o", H
		} else if 677 <= r && r < 712 {
			return seed + "p", H
		} else if 712 <= r && r < 713 {
			return seed + "q", H
		} else if 713 <= r && r < 714 {
			return seed + "r", H
		} else if 714 <= r && r < 801 {
			return seed + "s", H
		} else if 801 <= r && r < 896 {
			return seed + "t", H
		} else if 896 <= r && r < 906 {
			return seed + "u", H
		} else if 906 <= r && r < 967 {
			return seed + "v", H
		} else if 967 <= r && r < 968 {
			return seed + "w", H
		} else if 968 <= r && r < 971 {
			return seed + "x", H
		} else if 971 <= r && r < 988 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lv" {
		r := rng.IntN(48)
		H := 1.6461654522610805
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 38 {
			return seed + "e", H
		} else if 38 <= r && r < 45 {
			return seed + "i", H
		} else if 45 <= r && r < 46 {
			return seed + "o", H
		} else if 46 <= r && r < 47 {
			return seed + "u", H
		} else if 47 <= r && r < 48 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fr" {
		r := rng.IntN(202)
		H := 2.230355417282172
		if 0 <= r && r < 47 {
			return seed + "a", H
		} else if 47 <= r && r < 114 {
			return seed + "e", H
		} else if 114 <= r && r < 147 {
			return seed + "i", H
		} else if 147 <= r && r < 188 {
			return seed + "o", H
		} else if 188 <= r && r < 199 {
			return seed + "u", H
		} else if 199 <= r && r < 202 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oa" {
		r := rng.IntN(192)
		H := 3.399326801846604
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 53 {
			return seed + "d", H
		} else if 53 <= r && r < 56 {
			return seed + "f", H
		} else if 56 <= r && r < 59 {
			return seed + "g", H
		} else if 59 <= r && r < 60 {
			return seed + "h", H
		} else if 60 <= r && r < 61 {
			return seed + "j", H
		} else if 61 <= r && r < 72 {
			return seed + "k", H
		} else if 72 <= r && r < 79 {
			return seed + "l", H
		} else if 79 <= r && r < 86 {
			return seed + "m", H
		} else if 86 <= r && r < 93 {
			return seed + "n", H
		} else if 93 <= r && r < 94 {
			return seed + "p", H
		} else if 94 <= r && r < 95 {
			return seed + "q", H
		} else if 95 <= r && r < 122 {
			return seed + "r", H
		} else if 122 <= r && r < 149 {
			return seed + "s", H
		} else if 149 <= r && r < 186 {
			return seed + "t", H
		} else if 186 <= r && r < 188 {
			return seed + "u", H
		} else if 188 <= r && r < 189 {
			return seed + "v", H
		} else if 189 <= r && r < 190 {
			return seed + "w", H
		} else if 190 <= r && r < 191 {
			return seed + "x", H
		} else if 191 <= r && r < 192 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "se" {
		r := rng.IntN(418)
		H := 3.745674478657474
		if 0 <= r && r < 12 {
			return seed + "a", H
		} else if 12 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 44 {
			return seed + "c", H
		} else if 44 <= r && r < 109 {
			return seed + "d", H
		} else if 109 <= r && r < 119 {
			return seed + "e", H
		} else if 119 <= r && r < 122 {
			return seed + "f", H
		} else if 122 <= r && r < 125 {
			return seed + "g", H
		} else if 125 <= r && r < 126 {
			return seed + "h", H
		} else if 126 <= r && r < 130 {
			return seed + "i", H
		} else if 130 <= r && r < 131 {
			return seed + "j", H
		} else if 131 <= r && r < 132 {
			return seed + "k", H
		} else if 132 <= r && r < 185 {
			return seed + "l", H
		} else if 185 <= r && r < 204 {
			return seed + "m", H
		} else if 204 <= r && r < 255 {
			return seed + "n", H
		} else if 255 <= r && r < 264 {
			return seed + "p", H
		} else if 264 <= r && r < 271 {
			return seed + "q", H
		} else if 271 <= r && r < 328 {
			return seed + "r", H
		} else if 328 <= r && r < 361 {
			return seed + "s", H
		} else if 361 <= r && r < 390 {
			return seed + "t", H
		} else if 390 <= r && r < 394 {
			return seed + "u", H
		} else if 394 <= r && r < 407 {
			return seed + "v", H
		} else if 407 <= r && r < 410 {
			return seed + "w", H
		} else if 410 <= r && r < 413 {
			return seed + "x", H
		} else if 413 <= r && r < 417 {
			return seed + "y", H
		} else if 417 <= r && r < 418 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "qo" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "sn" {
		r := rng.IntN(102)
		H := 2.0594822608842542
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 28 {
			return seed + "e", H
		} else if 28 <= r && r < 41 {
			return seed + "i", H
		} else if 41 <= r && r < 90 {
			return seed + "o", H
		} else if 90 <= r && r < 101 {
			return seed + "u", H
		} else if 101 <= r && r < 102 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zb" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kk" {
		r := rng.IntN(10)
		H := 2.646439344671015
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 7 {
			return seed + "n", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "al" {
		r := rng.IntN(506)
		H := 3.6964082830610514
		if 0 <= r && r < 43 {
			return seed + "a", H
		} else if 43 <= r && r < 47 {
			return seed + "b", H
		} else if 47 <= r && r < 59 {
			return seed + "c", H
		} else if 59 <= r && r < 67 {
			return seed + "d", H
		} else if 67 <= r && r < 102 {
			return seed + "e", H
		} else if 102 <= r && r < 110 {
			return seed + "f", H
		} else if 110 <= r && r < 114 {
			return seed + "g", H
		} else if 114 <= r && r < 209 {
			return seed + "i", H
		} else if 209 <= r && r < 227 {
			return seed + "k", H
		} else if 227 <= r && r < 337 {
			return seed + "l", H
		} else if 337 <= r && r < 355 {
			return seed + "m", H
		} else if 355 <= r && r < 359 {
			return seed + "n", H
		} else if 359 <= r && r < 392 {
			return seed + "o", H
		} else if 392 <= r && r < 400 {
			return seed + "p", H
		} else if 400 <= r && r < 408 {
			return seed + "r", H
		} else if 408 <= r && r < 422 {
			return seed + "s", H
		} else if 422 <= r && r < 454 {
			return seed + "t", H
		} else if 454 <= r && r < 477 {
			return seed + "u", H
		} else if 477 <= r && r < 489 {
			return seed + "v", H
		} else if 489 <= r && r < 491 {
			return seed + "w", H
		} else if 491 <= r && r < 504 {
			return seed + "y", H
		} else if 504 <= r && r < 506 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "el" {
		r := rng.IntN(468)
		H := 3.49667796316439
		if 0 <= r && r < 31 {
			return seed + "a", H
		} else if 31 <= r && r < 33 {
			return seed + "b", H
		} else if 33 <= r && r < 37 {
			return seed + "c", H
		} else if 37 <= r && r < 67 {
			return seed + "d", H
		} else if 67 <= r && r < 156 {
			return seed + "e", H
		} else if 156 <= r && r < 172 {
			return seed + "f", H
		} else if 172 <= r && r < 261 {
			return seed + "i", H
		} else if 261 <= r && r < 263 {
			return seed + "k", H
		} else if 263 <= r && r < 319 {
			return seed + "l", H
		} else if 319 <= r && r < 323 {
			return seed + "m", H
		} else if 323 <= r && r < 327 {
			return seed + "n", H
		} else if 327 <= r && r < 356 {
			return seed + "o", H
		} else if 356 <= r && r < 372 {
			return seed + "p", H
		} else if 372 <= r && r < 378 {
			return seed + "s", H
		} else if 378 <= r && r < 396 {
			return seed + "t", H
		} else if 396 <= r && r < 411 {
			return seed + "u", H
		} else if 411 <= r && r < 423 {
			return seed + "v", H
		} else if 423 <= r && r < 468 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ag" {
		r := rng.IntN(364)
		H := 2.8615484404812834
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 162 {
			return seed + "e", H
		} else if 162 <= r && r < 164 {
			return seed + "f", H
		} else if 164 <= r && r < 210 {
			return seed + "g", H
		} else if 210 <= r && r < 212 {
			return seed + "h", H
		} else if 212 <= r && r < 245 {
			return seed + "i", H
		} else if 245 <= r && r < 247 {
			return seed + "l", H
		} else if 247 <= r && r < 255 {
			return seed + "m", H
		} else if 255 <= r && r < 281 {
			return seed + "n", H
		} else if 281 <= r && r < 302 {
			return seed + "o", H
		} else if 302 <= r && r < 306 {
			return seed + "p", H
		} else if 306 <= r && r < 338 {
			return seed + "r", H
		} else if 338 <= r && r < 348 {
			return seed + "s", H
		} else if 348 <= r && r < 361 {
			return seed + "u", H
		} else if 361 <= r && r < 363 {
			return seed + "w", H
		} else if 363 <= r && r < 364 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "n" {
		r := rng.IntN(200)
		H := 2.043092467076734
		if 0 <= r && r < 51 {
			return seed + "a", H
		} else if 51 <= r && r < 118 {
			return seed + "e", H
		} else if 118 <= r && r < 143 {
			return seed + "i", H
		} else if 143 <= r && r < 144 {
			return seed + "o", H
		} else if 144 <= r && r < 197 {
			return seed + "u", H
		} else if 197 <= r && r < 200 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "f" {
		r := rng.IntN(622)
		H := 2.554187564301907
		if 0 <= r && r < 109 {
			return seed + "a", H
		} else if 109 <= r && r < 174 {
			return seed + "e", H
		} else if 174 <= r && r < 257 {
			return seed + "i", H
		} else if 257 <= r && r < 383 {
			return seed + "l", H
		} else if 383 <= r && r < 468 {
			return seed + "o", H
		} else if 468 <= r && r < 620 {
			return seed + "r", H
		} else if 620 <= r && r < 621 {
			return seed + "u", H
		} else if 621 <= r && r < 622 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ip" {
		r := rng.IntN(122)
		H := 3.0161565071636254
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 7 {
			return seed + "c", H
		} else if 7 <= r && r < 34 {
			return seed + "e", H
		} else if 34 <= r && r < 36 {
			return seed + "f", H
		} else if 36 <= r && r < 42 {
			return seed + "h", H
		} else if 42 <= r && r < 45 {
			return seed + "i", H
		} else if 45 <= r && r < 51 {
			return seed + "l", H
		} else if 51 <= r && r < 53 {
			return seed + "m", H
		} else if 53 <= r && r < 60 {
			return seed + "o", H
		} else if 60 <= r && r < 100 {
			return seed + "p", H
		} else if 100 <= r && r < 106 {
			return seed + "s", H
		} else if 106 <= r && r < 118 {
			return seed + "t", H
		} else if 118 <= r && r < 121 {
			return seed + "u", H
		} else if 121 <= r && r < 122 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yn" {
		r := rng.IntN(28)
		H := 2.6020475662770886
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 11 {
			return seed + "d", H
		} else if 11 <= r && r < 16 {
			return seed + "e", H
		} else if 16 <= r && r < 17 {
			return seed + "i", H
		} else if 17 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 26 {
			return seed + "t", H
		} else if 26 <= r && r < 27 {
			return seed + "u", H
		} else if 27 <= r && r < 28 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ew" {
		r := rng.IntN(94)
		H := 2.932631803459705
		if 0 <= r && r < 33 {
			return seed + "a", H
		} else if 33 <= r && r < 35 {
			return seed + "d", H
		} else if 35 <= r && r < 48 {
			return seed + "e", H
		} else if 48 <= r && r < 50 {
			return seed + "h", H
		} else if 50 <= r && r < 67 {
			return seed + "i", H
		} else if 67 <= r && r < 71 {
			return seed + "l", H
		} else if 71 <= r && r < 75 {
			return seed + "m", H
		} else if 75 <= r && r < 77 {
			return seed + "n", H
		} else if 77 <= r && r < 84 {
			return seed + "o", H
		} else if 84 <= r && r < 86 {
			return seed + "p", H
		} else if 86 <= r && r < 90 {
			return seed + "r", H
		} else if 90 <= r && r < 91 {
			return seed + "u", H
		} else if 91 <= r && r < 94 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "aj" {
		r := rng.IntN(18)
		H := 2.2349849223695104
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gy" {
		r := rng.IntN(24)
		H := 4.1887218755408675
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 18 {
			return seed + "r", H
		} else if 18 <= r && r < 19 {
			return seed + "s", H
		} else if 19 <= r && r < 20 {
			return seed + "t", H
		} else if 20 <= r && r < 21 {
			return seed + "v", H
		} else if 21 <= r && r < 22 {
			return seed + "w", H
		} else if 22 <= r && r < 23 {
			return seed + "x", H
		} else if 23 <= r && r < 24 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "op" {
		r := rng.IntN(218)
		H := 3.2563225938929627
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 11 {
			return seed + "c", H
		} else if 11 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 72 {
			return seed + "e", H
		} else if 72 <= r && r < 82 {
			return seed + "h", H
		} else if 82 <= r && r < 109 {
			return seed + "i", H
		} else if 109 <= r && r < 111 {
			return seed + "k", H
		} else if 111 <= r && r < 123 {
			return seed + "l", H
		} else if 123 <= r && r < 136 {
			return seed + "o", H
		} else if 136 <= r && r < 178 {
			return seed + "p", H
		} else if 178 <= r && r < 190 {
			return seed + "s", H
		} else if 190 <= r && r < 194 {
			return seed + "t", H
		} else if 194 <= r && r < 207 {
			return seed + "u", H
		} else if 207 <= r && r < 209 {
			return seed + "w", H
		} else if 209 <= r && r < 218 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qu" {
		r := rng.IntN(218)
		H := 2.421738350941232
		if 0 <= r && r < 68 {
			return seed + "a", H
		} else if 68 <= r && r < 69 {
			return seed + "b", H
		} else if 69 <= r && r < 70 {
			return seed + "c", H
		} else if 70 <= r && r < 71 {
			return seed + "d", H
		} else if 71 <= r && r < 119 {
			return seed + "e", H
		} else if 119 <= r && r < 120 {
			return seed + "f", H
		} else if 120 <= r && r < 121 {
			return seed + "g", H
		} else if 121 <= r && r < 122 {
			return seed + "h", H
		} else if 122 <= r && r < 196 {
			return seed + "i", H
		} else if 196 <= r && r < 197 {
			return seed + "j", H
		} else if 197 <= r && r < 198 {
			return seed + "k", H
		} else if 198 <= r && r < 199 {
			return seed + "l", H
		} else if 199 <= r && r < 200 {
			return seed + "m", H
		} else if 200 <= r && r < 201 {
			return seed + "n", H
		} else if 201 <= r && r < 209 {
			return seed + "o", H
		} else if 209 <= r && r < 210 {
			return seed + "p", H
		} else if 210 <= r && r < 211 {
			return seed + "q", H
		} else if 211 <= r && r < 212 {
			return seed + "r", H
		} else if 212 <= r && r < 213 {
			return seed + "s", H
		} else if 213 <= r && r < 214 {
			return seed + "t", H
		} else if 214 <= r && r < 215 {
			return seed + "v", H
		} else if 215 <= r && r < 216 {
			return seed + "w", H
		} else if 216 <= r && r < 217 {
			return seed + "x", H
		} else if 217 <= r && r < 218 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xo" {
		r := rng.IntN(30)
		H := 4.0444138295776115
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 15 {
			return seed + "n", H
		} else if 15 <= r && r < 18 {
			return seed + "p", H
		} else if 18 <= r && r < 19 {
			return seed + "q", H
		} else if 19 <= r && r < 24 {
			return seed + "r", H
		} else if 24 <= r && r < 25 {
			return seed + "s", H
		} else if 25 <= r && r < 26 {
			return seed + "t", H
		} else if 26 <= r && r < 27 {
			return seed + "v", H
		} else if 27 <= r && r < 28 {
			return seed + "w", H
		} else if 28 <= r && r < 29 {
			return seed + "x", H
		} else if 29 <= r && r < 30 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "rg" {
		r := rng.IntN(124)
		H := 2.6508761522487303
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 56 {
			return seed + "e", H
		} else if 56 <= r && r < 75 {
			return seed + "i", H
		} else if 75 <= r && r < 87 {
			return seed + "l", H
		} else if 87 <= r && r < 100 {
			return seed + "o", H
		} else if 100 <= r && r < 108 {
			return seed + "r", H
		} else if 108 <= r && r < 115 {
			return seed + "u", H
		} else if 115 <= r && r < 124 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ln" {
		r := rng.IntN(18)
		H := 1.9399824743635556
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ru" {
		r := rng.IntN(316)
		H := 3.5667344720949834
		if 0 <= r && r < 25 {
			return seed + "b", H
		} else if 25 <= r && r < 48 {
			return seed + "c", H
		} else if 48 <= r && r < 75 {
			return seed + "d", H
		} else if 75 <= r && r < 85 {
			return seed + "e", H
		} else if 85 <= r && r < 96 {
			return seed + "f", H
		} else if 96 <= r && r < 107 {
			return seed + "g", H
		} else if 107 <= r && r < 108 {
			return seed + "h", H
		} else if 108 <= r && r < 114 {
			return seed + "i", H
		} else if 114 <= r && r < 115 {
			return seed + "j", H
		} else if 115 <= r && r < 116 {
			return seed + "k", H
		} else if 116 <= r && r < 123 {
			return seed + "l", H
		} else if 123 <= r && r < 162 {
			return seed + "m", H
		} else if 162 <= r && r < 211 {
			return seed + "n", H
		} else if 211 <= r && r < 224 {
			return seed + "p", H
		} else if 224 <= r && r < 225 {
			return seed + "q", H
		} else if 225 <= r && r < 228 {
			return seed + "r", H
		} else if 228 <= r && r < 295 {
			return seed + "s", H
		} else if 295 <= r && r < 310 {
			return seed + "t", H
		} else if 310 <= r && r < 311 {
			return seed + "v", H
		} else if 311 <= r && r < 312 {
			return seed + "w", H
		} else if 312 <= r && r < 315 {
			return seed + "x", H
		} else if 315 <= r && r < 316 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vp" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ac" {
		r := rng.IntN(562)
		H := 3.0537155907352163
		if 0 <= r && r < 21 {
			return seed + "a", H
		} else if 21 <= r && r < 39 {
			return seed + "c", H
		} else if 39 <= r && r < 106 {
			return seed + "e", H
		} else if 106 <= r && r < 172 {
			return seed + "h", H
		} else if 172 <= r && r < 225 {
			return seed + "i", H
		} else if 225 <= r && r < 387 {
			return seed + "k", H
		} else if 387 <= r && r < 391 {
			return seed + "l", H
		} else if 391 <= r && r < 406 {
			return seed + "o", H
		} else if 406 <= r && r < 412 {
			return seed + "q", H
		} else if 412 <= r && r < 424 {
			return seed + "r", H
		} else if 424 <= r && r < 426 {
			return seed + "s", H
		} else if 426 <= r && r < 530 {
			return seed + "t", H
		} else if 530 <= r && r < 545 {
			return seed + "u", H
		} else if 545 <= r && r < 562 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "t" {
		r := rng.IntN(652)
		H := 2.501876728590702
		if 0 <= r && r < 129 {
			return seed + "a", H
		} else if 129 <= r && r < 130 {
			return seed + "e", H
		} else if 130 <= r && r < 252 {
			return seed + "h", H
		} else if 252 <= r && r < 335 {
			return seed + "i", H
		} else if 335 <= r && r < 336 {
			return seed + "o", H
		} else if 336 <= r && r < 546 {
			return seed + "r", H
		} else if 546 <= r && r < 548 {
			return seed + "s", H
		} else if 548 <= r && r < 599 {
			return seed + "u", H
		} else if 599 <= r && r < 645 {
			return seed + "w", H
		} else if 645 <= r && r < 652 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ws" {
		r := rng.IntN(18)
		H := 2.774401919288771
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 14 {
			return seed + "t", H
		} else if 14 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uu" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zh" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cs" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pr" {
		r := rng.IntN(484)
		H := 1.986523655080605
		if 0 <= r && r < 33 {
			return seed + "a", H
		} else if 33 <= r && r < 174 {
			return seed + "e", H
		} else if 174 <= r && r < 265 {
			return seed + "i", H
		} else if 265 <= r && r < 464 {
			return seed + "o", H
		} else if 464 <= r && r < 479 {
			return seed + "u", H
		} else if 479 <= r && r < 484 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ft" {
		r := rng.IntN(72)
		H := 2.6518176152750605
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 28 {
			return seed + "e", H
		} else if 28 <= r && r < 30 {
			return seed + "h", H
		} else if 30 <= r && r < 45 {
			return seed + "i", H
		} else if 45 <= r && r < 51 {
			return seed + "l", H
		} else if 51 <= r && r < 53 {
			return seed + "n", H
		} else if 53 <= r && r < 56 {
			return seed + "o", H
		} else if 56 <= r && r < 58 {
			return seed + "s", H
		} else if 58 <= r && r < 59 {
			return seed + "u", H
		} else if 59 <= r && r < 61 {
			return seed + "w", H
		} else if 61 <= r && r < 72 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sd" {
		r := rng.IntN(14)
		H := 2.064042639445697
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wl" {
		r := rng.IntN(26)
		H := 1.7529381278020533
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 14 {
			return seed + "e", H
		} else if 14 <= r && r < 23 {
			return seed + "i", H
		} else if 23 <= r && r < 24 {
			return seed + "o", H
		} else if 24 <= r && r < 25 {
			return seed + "u", H
		} else if 25 <= r && r < 26 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sf" {
		r := rng.IntN(20)
		H := 2.6904685707328286
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ot" {
		r := rng.IntN(282)
		H := 3.432576510412063
		if 0 <= r && r < 21 {
			return seed + "a", H
		} else if 21 <= r && r < 27 {
			return seed + "b", H
		} else if 27 <= r && r < 33 {
			return seed + "c", H
		} else if 33 <= r && r < 80 {
			return seed + "e", H
		} else if 80 <= r && r < 82 {
			return seed + "g", H
		} else if 82 <= r && r < 122 {
			return seed + "h", H
		} else if 122 <= r && r < 177 {
			return seed + "i", H
		} else if 177 <= r && r < 187 {
			return seed + "l", H
		} else if 187 <= r && r < 189 {
			return seed + "m", H
		} else if 189 <= r && r < 191 {
			return seed + "n", H
		} else if 191 <= r && r < 210 {
			return seed + "o", H
		} else if 210 <= r && r < 216 {
			return seed + "p", H
		} else if 216 <= r && r < 222 {
			return seed + "r", H
		} else if 222 <= r && r < 230 {
			return seed + "s", H
		} else if 230 <= r && r < 270 {
			return seed + "t", H
		} else if 270 <= r && r < 273 {
			return seed + "u", H
		} else if 273 <= r && r < 277 {
			return seed + "w", H
		} else if 277 <= r && r < 282 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dw" {
		r := rng.IntN(44)
		H := 2.3727787125137962
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 22 {
			return seed + "e", H
		} else if 22 <= r && r < 29 {
			return seed + "i", H
		} else if 29 <= r && r < 40 {
			return seed + "o", H
		} else if 40 <= r && r < 42 {
			return seed + "r", H
		} else if 42 <= r && r < 43 {
			return seed + "u", H
		} else if 43 <= r && r < 44 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ty" {
		r := rng.IntN(60)
		H := 3.552917257559472
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 9 {
			return seed + "g", H
		} else if 9 <= r && r < 12 {
			return seed + "h", H
		} else if 12 <= r && r < 18 {
			return seed + "i", H
		} else if 18 <= r && r < 19 {
			return seed + "j", H
		} else if 19 <= r && r < 22 {
			return seed + "k", H
		} else if 22 <= r && r < 37 {
			return seed + "l", H
		} else if 37 <= r && r < 38 {
			return seed + "m", H
		} else if 38 <= r && r < 39 {
			return seed + "n", H
		} else if 39 <= r && r < 52 {
			return seed + "p", H
		} else if 52 <= r && r < 53 {
			return seed + "q", H
		} else if 53 <= r && r < 54 {
			return seed + "r", H
		} else if 54 <= r && r < 55 {
			return seed + "s", H
		} else if 55 <= r && r < 56 {
			return seed + "t", H
		} else if 56 <= r && r < 57 {
			return seed + "v", H
		} else if 57 <= r && r < 58 {
			return seed + "w", H
		} else if 58 <= r && r < 59 {
			return seed + "x", H
		} else if 59 <= r && r < 60 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "yl" {
		r := rng.IntN(40)
		H := 1.8204796593794925
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 31 {
			return seed + "i", H
		} else if 31 <= r && r < 36 {
			return seed + "o", H
		} else if 36 <= r && r < 39 {
			return seed + "u", H
		} else if 39 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nn" {
		r := rng.IntN(154)
		H := 2.2344996919449143
		if 0 <= r && r < 21 {
			return seed + "a", H
		} else if 21 <= r && r < 84 {
			return seed + "e", H
		} else if 84 <= r && r < 117 {
			return seed + "i", H
		} else if 117 <= r && r < 132 {
			return seed + "o", H
		} else if 132 <= r && r < 137 {
			return seed + "u", H
		} else if 137 <= r && r < 154 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dv" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nw" {
		r := rng.IntN(60)
		H := 2.3476104386001415
		if 0 <= r && r < 17 {
			return seed + "a", H
		} else if 17 <= r && r < 26 {
			return seed + "e", H
		} else if 26 <= r && r < 37 {
			return seed + "i", H
		} else if 37 <= r && r < 54 {
			return seed + "o", H
		} else if 54 <= r && r < 58 {
			return seed + "r", H
		} else if 58 <= r && r < 59 {
			return seed + "u", H
		} else if 59 <= r && r < 60 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yt" {
		r := rng.IntN(28)
		H := 2.4191133005120022
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 14 {
			return seed + "h", H
		} else if 14 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 26 {
			return seed + "o", H
		} else if 26 <= r && r < 27 {
			return seed + "u", H
		} else if 27 <= r && r < 28 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "x" {
		r := rng.IntN(10)
		H := 2.646439344671015
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wu" {
		r := rng.IntN(22)
		H := 4.243300368538956
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 17 {
			return seed + "s", H
		} else if 17 <= r && r < 18 {
			return seed + "t", H
		} else if 18 <= r && r < 19 {
			return seed + "v", H
		} else if 19 <= r && r < 20 {
			return seed + "w", H
		} else if 20 <= r && r < 21 {
			return seed + "x", H
		} else if 21 <= r && r < 22 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "db" {
		r := rng.IntN(34)
		H := 2.267240151565867
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 16 {
			return seed + "e", H
		} else if 16 <= r && r < 19 {
			return seed + "i", H
		} else if 19 <= r && r < 21 {
			return seed + "l", H
		} else if 21 <= r && r < 30 {
			return seed + "o", H
		} else if 30 <= r && r < 32 {
			return seed + "r", H
		} else if 32 <= r && r < 33 {
			return seed + "u", H
		} else if 33 <= r && r < 34 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "iu" {
		r := rng.IntN(56)
		H := 2.3653946626241913
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 46 {
			return seed + "m", H
		} else if 46 <= r && r < 47 {
			return seed + "n", H
		} else if 47 <= r && r < 48 {
			return seed + "p", H
		} else if 48 <= r && r < 49 {
			return seed + "q", H
		} else if 49 <= r && r < 50 {
			return seed + "r", H
		} else if 50 <= r && r < 51 {
			return seed + "s", H
		} else if 51 <= r && r < 52 {
			return seed + "t", H
		} else if 52 <= r && r < 53 {
			return seed + "v", H
		} else if 53 <= r && r < 54 {
			return seed + "w", H
		} else if 54 <= r && r < 55 {
			return seed + "x", H
		} else if 55 <= r && r < 56 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "nc" {
		r := rng.IntN(420)
		H := 2.8684301140450312
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 146 {
			return seed + "e", H
		} else if 146 <= r && r < 232 {
			return seed + "h", H
		} else if 232 <= r && r < 275 {
			return seed + "i", H
		} else if 275 <= r && r < 299 {
			return seed + "l", H
		} else if 299 <= r && r < 338 {
			return seed + "o", H
		} else if 338 <= r && r < 354 {
			return seed + "r", H
		} else if 354 <= r && r < 376 {
			return seed + "t", H
		} else if 376 <= r && r < 387 {
			return seed + "u", H
		} else if 387 <= r && r < 420 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "es" {
		r := rng.IntN(816)
		H := 2.2763699907157355
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 33 {
			return seed + "c", H
		} else if 33 <= r && r < 35 {
			return seed + "d", H
		} else if 35 <= r && r < 72 {
			return seed + "e", H
		} else if 72 <= r && r < 94 {
			return seed + "h", H
		} else if 94 <= r && r < 143 {
			return seed + "i", H
		} else if 143 <= r && r < 153 {
			return seed + "k", H
		} else if 153 <= r && r < 155 {
			return seed + "l", H
		} else if 155 <= r && r < 159 {
			return seed + "m", H
		} else if 159 <= r && r < 180 {
			return seed + "o", H
		} else if 180 <= r && r < 196 {
			return seed + "p", H
		} else if 196 <= r && r < 198 {
			return seed + "q", H
		} else if 198 <= r && r < 662 {
			return seed + "s", H
		} else if 662 <= r && r < 796 {
			return seed + "t", H
		} else if 796 <= r && r < 813 {
			return seed + "u", H
		} else if 813 <= r && r < 816 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "je" {
		r := rng.IntN(52)
		H := 3.744261935633072
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 17 {
			return seed + "e", H
		} else if 17 <= r && r < 18 {
			return seed + "f", H
		} else if 18 <= r && r < 19 {
			return seed + "g", H
		} else if 19 <= r && r < 20 {
			return seed + "h", H
		} else if 20 <= r && r < 21 {
			return seed + "j", H
		} else if 21 <= r && r < 22 {
			return seed + "k", H
		} else if 22 <= r && r < 27 {
			return seed + "l", H
		} else if 27 <= r && r < 28 {
			return seed + "m", H
		} else if 28 <= r && r < 29 {
			return seed + "n", H
		} else if 29 <= r && r < 30 {
			return seed + "p", H
		} else if 30 <= r && r < 31 {
			return seed + "q", H
		} else if 31 <= r && r < 34 {
			return seed + "r", H
		} else if 34 <= r && r < 41 {
			return seed + "s", H
		} else if 41 <= r && r < 48 {
			return seed + "t", H
		} else if 48 <= r && r < 49 {
			return seed + "v", H
		} else if 49 <= r && r < 50 {
			return seed + "w", H
		} else if 50 <= r && r < 51 {
			return seed + "x", H
		} else if 51 <= r && r < 52 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ol" {
		r := rng.IntN(354)
		H := 3.2697718107965503
		if 0 <= r && r < 37 {
			return seed + "a", H
		} else if 37 <= r && r < 87 {
			return seed + "d", H
		} else if 87 <= r && r < 132 {
			return seed + "e", H
		} else if 132 <= r && r < 136 {
			return seed + "f", H
		} else if 136 <= r && r < 177 {
			return seed + "i", H
		} else if 177 <= r && r < 183 {
			return seed + "k", H
		} else if 183 <= r && r < 237 {
			return seed + "l", H
		} else if 237 <= r && r < 300 {
			return seed + "o", H
		} else if 300 <= r && r < 302 {
			return seed + "p", H
		} else if 302 <= r && r < 304 {
			return seed + "s", H
		} else if 304 <= r && r < 316 {
			return seed + "t", H
		} else if 316 <= r && r < 325 {
			return seed + "u", H
		} else if 325 <= r && r < 339 {
			return seed + "v", H
		} else if 339 <= r && r < 354 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xa" {
		r := rng.IntN(38)
		H := 3.9559910942383
		if 0 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 9 {
			return seed + "d", H
		} else if 9 <= r && r < 10 {
			return seed + "f", H
		} else if 10 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 17 {
			return seed + "j", H
		} else if 17 <= r && r < 18 {
			return seed + "k", H
		} else if 18 <= r && r < 21 {
			return seed + "l", H
		} else if 21 <= r && r < 26 {
			return seed + "m", H
		} else if 26 <= r && r < 27 {
			return seed + "n", H
		} else if 27 <= r && r < 28 {
			return seed + "p", H
		} else if 28 <= r && r < 29 {
			return seed + "q", H
		} else if 29 <= r && r < 30 {
			return seed + "r", H
		} else if 30 <= r && r < 31 {
			return seed + "s", H
		} else if 31 <= r && r < 34 {
			return seed + "t", H
		} else if 34 <= r && r < 35 {
			return seed + "v", H
		} else if 35 <= r && r < 36 {
			return seed + "w", H
		} else if 36 <= r && r < 37 {
			return seed + "x", H
		} else if 37 <= r && r < 38 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hs" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 6 {
			return seed + "t", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "st" {
		r := rng.IntN(998)
		H := 2.957473176091378
		if 0 <= r && r < 213 {
			return seed + "a", H
		} else if 213 <= r && r < 221 {
			return seed + "b", H
		} else if 221 <= r && r < 225 {
			return seed + "c", H
		} else if 225 <= r && r < 414 {
			return seed + "e", H
		} else if 414 <= r && r < 424 {
			return seed + "f", H
		} else if 424 <= r && r < 597 {
			return seed + "i", H
		} else if 597 <= r && r < 619 {
			return seed + "l", H
		} else if 619 <= r && r < 629 {
			return seed + "n", H
		} else if 629 <= r && r < 744 {
			return seed + "o", H
		} else if 744 <= r && r < 884 {
			return seed + "r", H
		} else if 884 <= r && r < 949 {
			return seed + "u", H
		} else if 949 <= r && r < 953 {
			return seed + "w", H
		} else if 953 <= r && r < 998 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ab" {
		r := rng.IntN(446)
		H := 2.0906025437154616
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 37 {
			return seed + "b", H
		} else if 37 <= r && r < 43 {
			return seed + "d", H
		} else if 43 <= r && r < 52 {
			return seed + "e", H
		} else if 52 <= r && r < 54 {
			return seed + "g", H
		} else if 54 <= r && r < 85 {
			return seed + "i", H
		} else if 85 <= r && r < 371 {
			return seed + "l", H
		} else if 371 <= r && r < 373 {
			return seed + "m", H
		} else if 373 <= r && r < 375 {
			return seed + "n", H
		} else if 375 <= r && r < 402 {
			return seed + "o", H
		} else if 402 <= r && r < 418 {
			return seed + "r", H
		} else if 418 <= r && r < 438 {
			return seed + "s", H
		} else if 438 <= r && r < 441 {
			return seed + "u", H
		} else if 441 <= r && r < 446 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ca" {
		r := rng.IntN(654)
		H := 3.392069435165686
		if 0 <= r && r < 21 {
			return seed + "b", H
		} else if 21 <= r && r < 32 {
			return seed + "c", H
		} else if 32 <= r && r < 49 {
			return seed + "d", H
		} else if 49 <= r && r < 52 {
			return seed + "f", H
		} else if 52 <= r && r < 57 {
			return seed + "g", H
		} else if 57 <= r && r < 60 {
			return seed + "h", H
		} else if 60 <= r && r < 61 {
			return seed + "j", H
		} else if 61 <= r && r < 70 {
			return seed + "k", H
		} else if 70 <= r && r < 181 {
			return seed + "l", H
		} else if 181 <= r && r < 206 {
			return seed + "m", H
		} else if 206 <= r && r < 275 {
			return seed + "n", H
		} else if 275 <= r && r < 328 {
			return seed + "p", H
		} else if 328 <= r && r < 329 {
			return seed + "q", H
		} else if 329 <= r && r < 472 {
			return seed + "r", H
		} else if 472 <= r && r < 511 {
			return seed + "s", H
		} else if 511 <= r && r < 612 {
			return seed + "t", H
		} else if 612 <= r && r < 630 {
			return seed + "u", H
		} else if 630 <= r && r < 647 {
			return seed + "v", H
		} else if 647 <= r && r < 650 {
			return seed + "w", H
		} else if 650 <= r && r < 651 {
			return seed + "x", H
		} else if 651 <= r && r < 653 {
			return seed + "y", H
		} else if 653 <= r && r < 654 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "wr" {
		r := rng.IntN(64)
		H := 2.2369062935020785
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 43 {
			return seed + "i", H
		} else if 43 <= r && r < 58 {
			return seed + "o", H
		} else if 58 <= r && r < 59 {
			return seed + "u", H
		} else if 59 <= r && r < 64 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vs" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hk" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rf" {
		r := rng.IntN(54)
		H := 2.6938170543971935
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 28 {
			return seed + "e", H
		} else if 28 <= r && r < 35 {
			return seed + "i", H
		} else if 35 <= r && r < 39 {
			return seed + "l", H
		} else if 39 <= r && r < 42 {
			return seed + "o", H
		} else if 42 <= r && r < 44 {
			return seed + "r", H
		} else if 44 <= r && r < 53 {
			return seed + "u", H
		} else if 53 <= r && r < 54 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sa" {
		r := rng.IntN(338)
		H := 3.8606917616050116
		if 0 <= r && r < 35 {
			return seed + "b", H
		} else if 35 <= r && r < 46 {
			return seed + "c", H
		} else if 46 <= r && r < 63 {
			return seed + "d", H
		} else if 63 <= r && r < 78 {
			return seed + "f", H
		} else if 78 <= r && r < 99 {
			return seed + "g", H
		} else if 99 <= r && r < 100 {
			return seed + "h", H
		} else if 100 <= r && r < 112 {
			return seed + "i", H
		} else if 112 <= r && r < 113 {
			return seed + "j", H
		} else if 113 <= r && r < 116 {
			return seed + "k", H
		} else if 116 <= r && r < 177 {
			return seed + "l", H
		} else if 177 <= r && r < 190 {
			return seed + "m", H
		} else if 190 <= r && r < 243 {
			return seed + "n", H
		} else if 243 <= r && r < 250 {
			return seed + "p", H
		} else if 250 <= r && r < 251 {
			return seed + "q", H
		} else if 251 <= r && r < 264 {
			return seed + "r", H
		} else if 264 <= r && r < 273 {
			return seed + "s", H
		} else if 273 <= r && r < 296 {
			return seed + "t", H
		} else if 296 <= r && r < 306 {
			return seed + "u", H
		} else if 306 <= r && r < 323 {
			return seed + "v", H
		} else if 323 <= r && r < 330 {
			return seed + "w", H
		} else if 330 <= r && r < 333 {
			return seed + "x", H
		} else if 333 <= r && r < 337 {
			return seed + "y", H
		} else if 337 <= r && r < 338 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xr" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "du" {
		r := rng.IntN(162)
		H := 3.769464031282425
		if 0 <= r && r < 6 {
			return seed + "a", H
		} else if 6 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 52 {
			return seed + "c", H
		} else if 52 <= r && r < 55 {
			return seed + "d", H
		} else if 55 <= r && r < 63 {
			return seed + "e", H
		} else if 63 <= r && r < 66 {
			return seed + "f", H
		} else if 66 <= r && r < 69 {
			return seed + "g", H
		} else if 69 <= r && r < 72 {
			return seed + "h", H
		} else if 72 <= r && r < 76 {
			return seed + "i", H
		} else if 76 <= r && r < 77 {
			return seed + "j", H
		} else if 77 <= r && r < 80 {
			return seed + "k", H
		} else if 80 <= r && r < 97 {
			return seed + "l", H
		} else if 97 <= r && r < 104 {
			return seed + "m", H
		} else if 104 <= r && r < 105 {
			return seed + "n", H
		} else if 105 <= r && r < 107 {
			return seed + "o", H
		} else if 107 <= r && r < 118 {
			return seed + "p", H
		} else if 118 <= r && r < 119 {
			return seed + "q", H
		} else if 119 <= r && r < 140 {
			return seed + "r", H
		} else if 140 <= r && r < 151 {
			return seed + "s", H
		} else if 151 <= r && r < 156 {
			return seed + "t", H
		} else if 156 <= r && r < 159 {
			return seed + "v", H
		} else if 159 <= r && r < 160 {
			return seed + "w", H
		} else if 160 <= r && r < 161 {
			return seed + "x", H
		} else if 161 <= r && r < 162 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "rh" {
		r := rng.IntN(28)
		H := 2.0897532141176747
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 25 {
			return seed + "u", H
		} else if 25 <= r && r < 28 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ux" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hy" {
		r := rng.IntN(62)
		H := 3.1008100776460776
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 19 {
			return seed + "d", H
		} else if 19 <= r && r < 20 {
			return seed + "f", H
		} else if 20 <= r && r < 21 {
			return seed + "g", H
		} else if 21 <= r && r < 22 {
			return seed + "h", H
		} else if 22 <= r && r < 23 {
			return seed + "j", H
		} else if 23 <= r && r < 24 {
			return seed + "k", H
		} else if 24 <= r && r < 25 {
			return seed + "l", H
		} else if 25 <= r && r < 28 {
			return seed + "m", H
		} else if 28 <= r && r < 29 {
			return seed + "n", H
		} else if 29 <= r && r < 52 {
			return seed + "p", H
		} else if 52 <= r && r < 53 {
			return seed + "q", H
		} else if 53 <= r && r < 54 {
			return seed + "r", H
		} else if 54 <= r && r < 57 {
			return seed + "s", H
		} else if 57 <= r && r < 58 {
			return seed + "t", H
		} else if 58 <= r && r < 59 {
			return seed + "v", H
		} else if 59 <= r && r < 60 {
			return seed + "w", H
		} else if 60 <= r && r < 61 {
			return seed + "x", H
		} else if 61 <= r && r < 62 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "jy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fb" {
		r := rng.IntN(10)
		H := 2.3709505944546683
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xc" {
		r := rng.IntN(46)
		H := 2.857013796380975
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 14 {
			return seed + "e", H
		} else if 14 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 31 {
			return seed + "l", H
		} else if 31 <= r && r < 32 {
			return seed + "o", H
		} else if 32 <= r && r < 36 {
			return seed + "r", H
		} else if 36 <= r && r < 45 {
			return seed + "u", H
		} else if 45 <= r && r < 46 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ue" {
		r := rng.IntN(98)
		H := 3.7256236802012
		if 0 <= r && r < 8 {
			return seed + "a", H
		} else if 8 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 21 {
			return seed + "d", H
		} else if 21 <= r && r < 29 {
			return seed + "e", H
		} else if 29 <= r && r < 32 {
			return seed + "f", H
		} else if 32 <= r && r < 33 {
			return seed + "g", H
		} else if 33 <= r && r < 34 {
			return seed + "h", H
		} else if 34 <= r && r < 35 {
			return seed + "j", H
		} else if 35 <= r && r < 36 {
			return seed + "k", H
		} else if 36 <= r && r < 55 {
			return seed + "l", H
		} else if 55 <= r && r < 56 {
			return seed + "m", H
		} else if 56 <= r && r < 73 {
			return seed + "n", H
		} else if 73 <= r && r < 74 {
			return seed + "p", H
		} else if 74 <= r && r < 75 {
			return seed + "q", H
		} else if 75 <= r && r < 80 {
			return seed + "r", H
		} else if 80 <= r && r < 89 {
			return seed + "s", H
		} else if 89 <= r && r < 92 {
			return seed + "t", H
		} else if 92 <= r && r < 94 {
			return seed + "u", H
		} else if 94 <= r && r < 95 {
			return seed + "v", H
		} else if 95 <= r && r < 96 {
			return seed + "w", H
		} else if 96 <= r && r < 97 {
			return seed + "x", H
		} else if 97 <= r && r < 98 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "eb" {
		r := rng.IntN(100)
		H := 2.947172585876227
		if 0 <= r && r < 21 {
			return seed + "a", H
		} else if 21 <= r && r < 25 {
			return seed + "b", H
		} else if 25 <= r && r < 30 {
			return seed + "e", H
		} else if 30 <= r && r < 37 {
			return seed + "i", H
		} else if 37 <= r && r < 45 {
			return seed + "l", H
		} else if 45 <= r && r < 66 {
			return seed + "o", H
		} else if 66 <= r && r < 76 {
			return seed + "r", H
		} else if 76 <= r && r < 80 {
			return seed + "t", H
		} else if 80 <= r && r < 99 {
			return seed + "u", H
		} else if 99 <= r && r < 100 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jn" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ld" {
		r := rng.IntN(72)
		H := 3.2315574085830168
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 9 {
			return seed + "c", H
		} else if 9 <= r && r < 28 {
			return seed + "e", H
		} else if 28 <= r && r < 34 {
			return seed + "f", H
		} else if 34 <= r && r < 36 {
			return seed + "h", H
		} else if 36 <= r && r < 43 {
			return seed + "i", H
		} else if 43 <= r && r < 55 {
			return seed + "l", H
		} else if 55 <= r && r < 57 {
			return seed + "m", H
		} else if 57 <= r && r < 59 {
			return seed + "n", H
		} else if 59 <= r && r < 66 {
			return seed + "o", H
		} else if 66 <= r && r < 68 {
			return seed + "s", H
		} else if 68 <= r && r < 69 {
			return seed + "u", H
		} else if 69 <= r && r < 72 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nl" {
		r := rng.IntN(110)
		H := 2.3509349645884354
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 71 {
			return seed + "i", H
		} else if 71 <= r && r < 90 {
			return seed + "o", H
		} else if 90 <= r && r < 95 {
			return seed + "u", H
		} else if 95 <= r && r < 110 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jr" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dd" {
		r := rng.IntN(102)
		H := 2.424492990791112
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 34 {
			return seed + "e", H
		} else if 34 <= r && r < 38 {
			return seed + "h", H
		} else if 38 <= r && r < 55 {
			return seed + "i", H
		} else if 55 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 92 {
			return seed + "o", H
		} else if 92 <= r && r < 94 {
			return seed + "r", H
		} else if 94 <= r && r < 95 {
			return seed + "u", H
		} else if 95 <= r && r < 102 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "k" {
		r := rng.IntN(116)
		H := 2.3015080389560936
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 34 {
			return seed + "e", H
		} else if 34 <= r && r < 89 {
			return seed + "i", H
		} else if 89 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 101 {
			return seed + "n", H
		} else if 101 <= r && r < 108 {
			return seed + "o", H
		} else if 108 <= r && r < 110 {
			return seed + "r", H
		} else if 110 <= r && r < 115 {
			return seed + "u", H
		} else if 115 <= r && r < 116 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pf" {
		r := rng.IntN(16)
		H := 2.4237949406953985
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 10 {
			return seed + "r", H
		} else if 10 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cl" {
		r := rng.IntN(252)
		H := 2.303542326198559
		if 0 <= r && r < 73 {
			return seed + "a", H
		} else if 73 <= r && r < 130 {
			return seed + "e", H
		} else if 130 <= r && r < 171 {
			return seed + "i", H
		} else if 171 <= r && r < 210 {
			return seed + "o", H
		} else if 210 <= r && r < 251 {
			return seed + "u", H
		} else if 251 <= r && r < 252 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "da" {
		r := rng.IntN(234)
		H := 3.717771766639873
		if 0 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 24 {
			return seed + "c", H
		} else if 24 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 34 {
			return seed + "f", H
		} else if 34 <= r && r < 39 {
			return seed + "g", H
		} else if 39 <= r && r < 40 {
			return seed + "h", H
		} else if 40 <= r && r < 52 {
			return seed + "i", H
		} else if 52 <= r && r < 53 {
			return seed + "j", H
		} else if 53 <= r && r < 54 {
			return seed + "k", H
		} else if 54 <= r && r < 65 {
			return seed + "l", H
		} else if 65 <= r && r < 68 {
			return seed + "m", H
		} else if 68 <= r && r < 107 {
			return seed + "n", H
		} else if 107 <= r && r < 108 {
			return seed + "p", H
		} else if 108 <= r && r < 109 {
			return seed + "q", H
		} else if 109 <= r && r < 142 {
			return seed + "r", H
		} else if 142 <= r && r < 149 {
			return seed + "s", H
		} else if 149 <= r && r < 178 {
			return seed + "t", H
		} else if 178 <= r && r < 184 {
			return seed + "u", H
		} else if 184 <= r && r < 187 {
			return seed + "v", H
		} else if 187 <= r && r < 192 {
			return seed + "w", H
		} else if 192 <= r && r < 193 {
			return seed + "x", H
		} else if 193 <= r && r < 229 {
			return seed + "y", H
		} else if 229 <= r && r < 234 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ef" {
		r := rng.IntN(188)
		H := 3.1258405597039096
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 38 {
			return seed + "e", H
		} else if 38 <= r && r < 46 {
			return seed + "f", H
		} else if 46 <= r && r < 79 {
			return seed + "i", H
		} else if 79 <= r && r < 99 {
			return seed + "l", H
		} else if 99 <= r && r < 122 {
			return seed + "o", H
		} else if 122 <= r && r < 136 {
			return seed + "r", H
		} else if 136 <= r && r < 148 {
			return seed + "t", H
		} else if 148 <= r && r < 183 {
			return seed + "u", H
		} else if 183 <= r && r < 188 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "io" {
		r := rng.IntN(394)
		H := 1.7314586053179288
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 17 {
			return seed + "j", H
		} else if 17 <= r && r < 18 {
			return seed + "k", H
		} else if 18 <= r && r < 31 {
			return seed + "l", H
		} else if 31 <= r && r < 34 {
			return seed + "m", H
		} else if 34 <= r && r < 315 {
			return seed + "n", H
		} else if 315 <= r && r < 316 {
			return seed + "p", H
		} else if 316 <= r && r < 317 {
			return seed + "q", H
		} else if 317 <= r && r < 330 {
			return seed + "r", H
		} else if 330 <= r && r < 331 {
			return seed + "s", H
		} else if 331 <= r && r < 342 {
			return seed + "t", H
		} else if 342 <= r && r < 388 {
			return seed + "u", H
		} else if 388 <= r && r < 389 {
			return seed + "v", H
		} else if 389 <= r && r < 390 {
			return seed + "w", H
		} else if 390 <= r && r < 393 {
			return seed + "x", H
		} else if 393 <= r && r < 394 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "wo" {
		r := rng.IntN(154)
		H := 3.0200652985123204
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 9 {
			return seed + "d", H
		} else if 9 <= r && r < 12 {
			return seed + "f", H
		} else if 12 <= r && r < 13 {
			return seed + "g", H
		} else if 13 <= r && r < 14 {
			return seed + "h", H
		} else if 14 <= r && r < 15 {
			return seed + "j", H
		} else if 15 <= r && r < 20 {
			return seed + "k", H
		} else if 20 <= r && r < 25 {
			return seed + "l", H
		} else if 25 <= r && r < 42 {
			return seed + "m", H
		} else if 42 <= r && r < 45 {
			return seed + "n", H
		} else if 45 <= r && r < 65 {
			return seed + "o", H
		} else if 65 <= r && r < 66 {
			return seed + "p", H
		} else if 66 <= r && r < 67 {
			return seed + "q", H
		} else if 67 <= r && r < 134 {
			return seed + "r", H
		} else if 134 <= r && r < 135 {
			return seed + "s", H
		} else if 135 <= r && r < 136 {
			return seed + "t", H
		} else if 136 <= r && r < 140 {
			return seed + "u", H
		} else if 140 <= r && r < 147 {
			return seed + "v", H
		} else if 147 <= r && r < 152 {
			return seed + "w", H
		} else if 152 <= r && r < 153 {
			return seed + "x", H
		} else if 153 <= r && r < 154 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "w" {
		r := rng.IntN(316)
		H := 2.2898450930569996
		if 0 <= r && r < 85 {
			return seed + "a", H
		} else if 85 <= r && r < 86 {
			return seed + "e", H
		} else if 86 <= r && r < 122 {
			return seed + "h", H
		} else if 122 <= r && r < 217 {
			return seed + "i", H
		} else if 217 <= r && r < 274 {
			return seed + "o", H
		} else if 274 <= r && r < 314 {
			return seed + "r", H
		} else if 314 <= r && r < 315 {
			return seed + "u", H
		} else if 315 <= r && r < 316 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "aq" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gw" {
		r := rng.IntN(16)
		H := 2.5550365325772657
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hf" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sh" {
		r := rng.IntN(472)
		H := 3.040508314254535
		if 0 <= r && r < 83 {
			return seed + "a", H
		} else if 83 <= r && r < 93 {
			return seed + "b", H
		} else if 93 <= r && r < 101 {
			return seed + "c", H
		} else if 101 <= r && r < 103 {
			return seed + "d", H
		} else if 103 <= r && r < 188 {
			return seed + "e", H
		} else if 188 <= r && r < 190 {
			return seed + "h", H
		} else if 190 <= r && r < 277 {
			return seed + "i", H
		} else if 277 <= r && r < 289 {
			return seed + "l", H
		} else if 289 <= r && r < 291 {
			return seed + "n", H
		} else if 291 <= r && r < 388 {
			return seed + "o", H
		} else if 388 <= r && r < 390 {
			return seed + "p", H
		} else if 390 <= r && r < 426 {
			return seed + "r", H
		} else if 426 <= r && r < 428 {
			return seed + "s", H
		} else if 428 <= r && r < 430 {
			return seed + "t", H
		} else if 430 <= r && r < 451 {
			return seed + "u", H
		} else if 451 <= r && r < 472 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bw" {
		r := rng.IntN(12)
		H := 2.396240625180289
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ib" {
		r := rng.IntN(108)
		H := 2.6964249538282297
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 17 {
			return seed + "c", H
		} else if 17 <= r && r < 40 {
			return seed + "e", H
		} else if 40 <= r && r < 47 {
			return seed + "i", H
		} else if 47 <= r && r < 85 {
			return seed + "l", H
		} else if 85 <= r && r < 88 {
			return seed + "o", H
		} else if 88 <= r && r < 94 {
			return seed + "r", H
		} else if 94 <= r && r < 107 {
			return seed + "u", H
		} else if 107 <= r && r < 108 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ou" {
		r := rng.IntN(536)
		H := 2.6212991922232844
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 28 {
			return seed + "c", H
		} else if 28 <= r && r < 35 {
			return seed + "d", H
		} else if 35 <= r && r < 36 {
			return seed + "f", H
		} else if 36 <= r && r < 51 {
			return seed + "g", H
		} else if 51 <= r && r < 52 {
			return seed + "h", H
		} else if 52 <= r && r < 53 {
			return seed + "j", H
		} else if 53 <= r && r < 54 {
			return seed + "k", H
		} else if 54 <= r && r < 59 {
			return seed + "l", H
		} else if 59 <= r && r < 60 {
			return seed + "m", H
		} else if 60 <= r && r < 191 {
			return seed + "n", H
		} else if 191 <= r && r < 204 {
			return seed + "p", H
		} else if 204 <= r && r < 205 {
			return seed + "q", H
		} else if 205 <= r && r < 238 {
			return seed + "r", H
		} else if 238 <= r && r < 385 {
			return seed + "s", H
		} else if 385 <= r && r < 532 {
			return seed + "t", H
		} else if 532 <= r && r < 533 {
			return seed + "v", H
		} else if 533 <= r && r < 534 {
			return seed + "w", H
		} else if 534 <= r && r < 535 {
			return seed + "x", H
		} else if 535 <= r && r < 536 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "qx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sc" {
		r := rng.IntN(274)
		H := 2.551187202536169
		if 0 <= r && r < 81 {
			return seed + "a", H
		} else if 81 <= r && r < 98 {
			return seed + "e", H
		} else if 98 <= r && r < 116 {
			return seed + "h", H
		} else if 116 <= r && r < 125 {
			return seed + "i", H
		} else if 125 <= r && r < 127 {
			return seed + "l", H
		} else if 127 <= r && r < 194 {
			return seed + "o", H
		} else if 194 <= r && r < 246 {
			return seed + "r", H
		} else if 246 <= r && r < 273 {
			return seed + "u", H
		} else if 273 <= r && r < 274 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "iw" {
		r := rng.IntN(10)
		H := 2.3709505944546683
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ml" {
		r := rng.IntN(32)
		H := 2.2226057415937133
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 16 {
			return seed + "e", H
		} else if 16 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 24 {
			return seed + "o", H
		} else if 24 <= r && r < 25 {
			return seed + "u", H
		} else if 25 <= r && r < 32 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dm" {
		r := rng.IntN(24)
		H := 2.0093984370492777
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 23 {
			return seed + "u", H
		} else if 23 <= r && r < 24 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ic" {
		r := rng.IntN(402)
		H := 3.10865538070483
		if 0 <= r && r < 83 {
			return seed + "a", H
		} else if 83 <= r && r < 142 {
			return seed + "e", H
		} else if 142 <= r && r < 150 {
			return seed + "h", H
		} else if 150 <= r && r < 205 {
			return seed + "i", H
		} else if 205 <= r && r < 279 {
			return seed + "k", H
		} else if 279 <= r && r < 289 {
			return seed + "l", H
		} else if 289 <= r && r < 316 {
			return seed + "o", H
		} else if 316 <= r && r < 318 {
			return seed + "r", H
		} else if 318 <= r && r < 352 {
			return seed + "s", H
		} else if 352 <= r && r < 382 {
			return seed + "t", H
		} else if 382 <= r && r < 391 {
			return seed + "u", H
		} else if 391 <= r && r < 402 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ro" {
		r := rng.IntN(796)
		H := 4.181153167717222
		if 0 <= r && r < 40 {
			return seed + "a", H
		} else if 40 <= r && r < 77 {
			return seed + "b", H
		} else if 77 <= r && r < 122 {
			return seed + "c", H
		} else if 122 <= r && r < 143 {
			return seed + "d", H
		} else if 143 <= r && r < 160 {
			return seed + "f", H
		} else if 160 <= r && r < 189 {
			return seed + "g", H
		} else if 189 <= r && r < 190 {
			return seed + "h", H
		} else if 190 <= r && r < 206 {
			return seed + "i", H
		} else if 206 <= r && r < 209 {
			return seed + "j", H
		} else if 209 <= r && r < 218 {
			return seed + "k", H
		} else if 218 <= r && r < 257 {
			return seed + "l", H
		} else if 257 <= r && r < 290 {
			return seed + "m", H
		} else if 290 <= r && r < 369 {
			return seed + "n", H
		} else if 369 <= r && r < 431 {
			return seed + "o", H
		} else if 431 <= r && r < 484 {
			return seed + "p", H
		} else if 484 <= r && r < 485 {
			return seed + "q", H
		} else if 485 <= r && r < 496 {
			return seed + "r", H
		} else if 496 <= r && r < 539 {
			return seed + "s", H
		} else if 539 <= r && r < 574 {
			return seed + "t", H
		} else if 574 <= r && r < 664 {
			return seed + "u", H
		} else if 664 <= r && r < 713 {
			return seed + "v", H
		} else if 713 <= r && r < 776 {
			return seed + "w", H
		} else if 776 <= r && r < 787 {
			return seed + "x", H
		} else if 787 <= r && r < 789 {
			return seed + "y", H
		} else if 789 <= r && r < 796 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tw" {
		r := rng.IntN(84)
		H := 2.0441830918840447
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 36 {
			return seed + "e", H
		} else if 36 <= r && r < 38 {
			return seed + "h", H
		} else if 38 <= r && r < 75 {
			return seed + "i", H
		} else if 75 <= r && r < 82 {
			return seed + "o", H
		} else if 82 <= r && r < 83 {
			return seed + "u", H
		} else if 83 <= r && r < 84 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "et" {
		r := rng.IntN(318)
		H := 3.255158026113046
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 29 {
			return seed + "b", H
		} else if 29 <= r && r < 37 {
			return seed + "c", H
		} else if 37 <= r && r < 39 {
			return seed + "d", H
		} else if 39 <= r && r < 104 {
			return seed + "e", H
		} else if 104 <= r && r < 108 {
			return seed + "f", H
		} else if 108 <= r && r < 138 {
			return seed + "h", H
		} else if 138 <= r && r < 195 {
			return seed + "i", H
		} else if 195 <= r && r < 218 {
			return seed + "o", H
		} else if 218 <= r && r < 246 {
			return seed + "r", H
		} else if 246 <= r && r < 248 {
			return seed + "s", H
		} else if 248 <= r && r < 294 {
			return seed + "t", H
		} else if 294 <= r && r < 305 {
			return seed + "u", H
		} else if 305 <= r && r < 316 {
			return seed + "y", H
		} else if 316 <= r && r < 318 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "cr" {
		r := rng.IntN(342)
		H := 2.4028604667342477
		if 0 <= r && r < 85 {
			return seed + "a", H
		} else if 85 <= r && r < 168 {
			return seed + "e", H
		} else if 168 <= r && r < 227 {
			return seed + "i", H
		} else if 227 <= r && r < 276 {
			return seed + "o", H
		} else if 276 <= r && r < 333 {
			return seed + "u", H
		} else if 333 <= r && r < 342 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wa" {
		r := rng.IntN(260)
		H := 3.5988728999558517
		if 0 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 14 {
			return seed + "c", H
		} else if 14 <= r && r < 21 {
			return seed + "d", H
		} else if 21 <= r && r < 26 {
			return seed + "f", H
		} else if 26 <= r && r < 39 {
			return seed + "g", H
		} else if 39 <= r && r < 40 {
			return seed + "h", H
		} else if 40 <= r && r < 42 {
			return seed + "i", H
		} else if 42 <= r && r < 43 {
			return seed + "j", H
		} else if 43 <= r && r < 54 {
			return seed + "k", H
		} else if 54 <= r && r < 75 {
			return seed + "l", H
		} else if 75 <= r && r < 78 {
			return seed + "m", H
		} else if 78 <= r && r < 91 {
			return seed + "n", H
		} else if 91 <= r && r < 94 {
			return seed + "p", H
		} else if 94 <= r && r < 95 {
			return seed + "q", H
		} else if 95 <= r && r < 146 {
			return seed + "r", H
		} else if 146 <= r && r < 193 {
			return seed + "s", H
		} else if 193 <= r && r < 206 {
			return seed + "t", H
		} else if 206 <= r && r < 215 {
			return seed + "v", H
		} else if 215 <= r && r < 220 {
			return seed + "w", H
		} else if 220 <= r && r < 221 {
			return seed + "x", H
		} else if 221 <= r && r < 259 {
			return seed + "y", H
		} else if 259 <= r && r < 260 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "nh" {
		r := rng.IntN(52)
		H := 2.323934450134503
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 26 {
			return seed + "e", H
		} else if 26 <= r && r < 33 {
			return seed + "i", H
		} else if 33 <= r && r < 46 {
			return seed + "o", H
		} else if 46 <= r && r < 51 {
			return seed + "u", H
		} else if 51 <= r && r < 52 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sb" {
		r := rng.IntN(18)
		H := 2.3516441151533924
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "y" {
		r := rng.IntN(74)
		H := 3.5273227992092866
		if 0 <= r && r < 12 {
			return seed + "a", H
		} else if 12 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 14 {
			return seed + "c", H
		} else if 14 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 35 {
			return seed + "e", H
		} else if 35 <= r && r < 36 {
			return seed + "f", H
		} else if 36 <= r && r < 37 {
			return seed + "g", H
		} else if 37 <= r && r < 38 {
			return seed + "h", H
		} else if 38 <= r && r < 46 {
			return seed + "i", H
		} else if 46 <= r && r < 47 {
			return seed + "j", H
		} else if 47 <= r && r < 48 {
			return seed + "k", H
		} else if 48 <= r && r < 49 {
			return seed + "l", H
		} else if 49 <= r && r < 50 {
			return seed + "m", H
		} else if 50 <= r && r < 51 {
			return seed + "n", H
		} else if 51 <= r && r < 63 {
			return seed + "o", H
		} else if 63 <= r && r < 64 {
			return seed + "p", H
		} else if 64 <= r && r < 65 {
			return seed + "q", H
		} else if 65 <= r && r < 66 {
			return seed + "r", H
		} else if 66 <= r && r < 67 {
			return seed + "s", H
		} else if 67 <= r && r < 68 {
			return seed + "t", H
		} else if 68 <= r && r < 70 {
			return seed + "u", H
		} else if 70 <= r && r < 71 {
			return seed + "v", H
		} else if 71 <= r && r < 72 {
			return seed + "w", H
		} else if 72 <= r && r < 73 {
			return seed + "x", H
		} else if 73 <= r && r < 74 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "an" {
		r := rng.IntN(1112)
		H := 3.4189585104251776
		if 0 <= r && r < 39 {
			return seed + "a", H
		} else if 39 <= r && r < 41 {
			return seed + "b", H
		} else if 41 <= r && r < 153 {
			return seed + "c", H
		} else if 153 <= r && r < 403 {
			return seed + "d", H
		} else if 403 <= r && r < 432 {
			return seed + "e", H
		} else if 432 <= r && r < 434 {
			return seed + "f", H
		} else if 434 <= r && r < 540 {
			return seed + "g", H
		} else if 540 <= r && r < 552 {
			return seed + "h", H
		} else if 552 <= r && r < 625 {
			return seed + "i", H
		} else if 625 <= r && r < 627 {
			return seed + "j", H
		} else if 627 <= r && r < 705 {
			return seed + "k", H
		} else if 705 <= r && r < 719 {
			return seed + "l", H
		} else if 719 <= r && r < 721 {
			return seed + "m", H
		} else if 721 <= r && r < 785 {
			return seed + "n", H
		} else if 785 <= r && r < 812 {
			return seed + "o", H
		} else if 812 <= r && r < 814 {
			return seed + "p", H
		} else if 814 <= r && r < 818 {
			return seed + "q", H
		} else if 818 <= r && r < 842 {
			return seed + "s", H
		} else if 842 <= r && r < 1068 {
			return seed + "t", H
		} else if 1068 <= r && r < 1075 {
			return seed + "u", H
		} else if 1075 <= r && r < 1077 {
			return seed + "v", H
		} else if 1077 <= r && r < 1110 {
			return seed + "y", H
		} else if 1110 <= r && r < 1112 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mv" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qf" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bi" {
		r := rng.IntN(164)
		H := 3.578780422738712
		if 0 <= r && r < 8 {
			return seed + "a", H
		} else if 8 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 20 {
			return seed + "c", H
		} else if 20 <= r && r < 35 {
			return seed + "d", H
		} else if 35 <= r && r < 43 {
			return seed + "e", H
		} else if 43 <= r && r < 44 {
			return seed + "f", H
		} else if 44 <= r && r < 49 {
			return seed + "g", H
		} else if 49 <= r && r < 50 {
			return seed + "h", H
		} else if 50 <= r && r < 51 {
			return seed + "j", H
		} else if 51 <= r && r < 52 {
			return seed + "k", H
		} else if 52 <= r && r < 83 {
			return seed + "l", H
		} else if 83 <= r && r < 84 {
			return seed + "m", H
		} else if 84 <= r && r < 109 {
			return seed + "n", H
		} else if 109 <= r && r < 113 {
			return seed + "o", H
		} else if 113 <= r && r < 114 {
			return seed + "p", H
		} else if 114 <= r && r < 115 {
			return seed + "q", H
		} else if 115 <= r && r < 124 {
			return seed + "r", H
		} else if 124 <= r && r < 129 {
			return seed + "s", H
		} else if 129 <= r && r < 158 {
			return seed + "t", H
		} else if 158 <= r && r < 159 {
			return seed + "v", H
		} else if 159 <= r && r < 160 {
			return seed + "w", H
		} else if 160 <= r && r < 161 {
			return seed + "x", H
		} else if 161 <= r && r < 164 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bf" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 5 {
			return seed + "l", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cy" {
		r := rng.IntN(46)
		H := 3.309865472705015
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 20 {
			return seed + "c", H
		} else if 20 <= r && r < 21 {
			return seed + "d", H
		} else if 21 <= r && r < 22 {
			return seed + "f", H
		} else if 22 <= r && r < 23 {
			return seed + "g", H
		} else if 23 <= r && r < 24 {
			return seed + "h", H
		} else if 24 <= r && r < 25 {
			return seed + "j", H
		} else if 25 <= r && r < 26 {
			return seed + "k", H
		} else if 26 <= r && r < 29 {
			return seed + "l", H
		} else if 29 <= r && r < 32 {
			return seed + "m", H
		} else if 32 <= r && r < 33 {
			return seed + "n", H
		} else if 33 <= r && r < 34 {
			return seed + "p", H
		} else if 34 <= r && r < 35 {
			return seed + "q", H
		} else if 35 <= r && r < 36 {
			return seed + "r", H
		} else if 36 <= r && r < 37 {
			return seed + "s", H
		} else if 37 <= r && r < 42 {
			return seed + "t", H
		} else if 42 <= r && r < 43 {
			return seed + "v", H
		} else if 43 <= r && r < 44 {
			return seed + "w", H
		} else if 44 <= r && r < 45 {
			return seed + "x", H
		} else if 45 <= r && r < 46 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ka" {
		r := rng.IntN(56)
		H := 3.805823267289439
		if 0 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 14 {
			return seed + "c", H
		} else if 14 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 16 {
			return seed + "f", H
		} else if 16 <= r && r < 21 {
			return seed + "g", H
		} else if 21 <= r && r < 22 {
			return seed + "h", H
		} else if 22 <= r && r < 23 {
			return seed + "j", H
		} else if 23 <= r && r < 24 {
			return seed + "k", H
		} else if 24 <= r && r < 31 {
			return seed + "l", H
		} else if 31 <= r && r < 32 {
			return seed + "m", H
		} else if 32 <= r && r < 35 {
			return seed + "n", H
		} else if 35 <= r && r < 36 {
			return seed + "p", H
		} else if 36 <= r && r < 37 {
			return seed + "q", H
		} else if 37 <= r && r < 44 {
			return seed + "r", H
		} else if 44 <= r && r < 45 {
			return seed + "s", H
		} else if 45 <= r && r < 50 {
			return seed + "t", H
		} else if 50 <= r && r < 51 {
			return seed + "v", H
		} else if 51 <= r && r < 52 {
			return seed + "w", H
		} else if 52 <= r && r < 53 {
			return seed + "x", H
		} else if 53 <= r && r < 55 {
			return seed + "y", H
		} else if 55 <= r && r < 56 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ko" {
		r := rng.IntN(26)
		H := 4.363713275750188
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 15 {
			return seed + "o", H
		} else if 15 <= r && r < 16 {
			return seed + "p", H
		} else if 16 <= r && r < 17 {
			return seed + "q", H
		} else if 17 <= r && r < 18 {
			return seed + "r", H
		} else if 18 <= r && r < 21 {
			return seed + "s", H
		} else if 21 <= r && r < 22 {
			return seed + "t", H
		} else if 22 <= r && r < 23 {
			return seed + "v", H
		} else if 23 <= r && r < 24 {
			return seed + "w", H
		} else if 24 <= r && r < 25 {
			return seed + "x", H
		} else if 25 <= r && r < 26 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lr" {
		r := rng.IntN(18)
		H := 2.3516441151533924
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "um" {
		r := rng.IntN(204)
		H := 2.8434929925855936
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 63 {
			return seed + "b", H
		} else if 63 <= r && r < 65 {
			return seed + "d", H
		} else if 65 <= r && r < 94 {
			return seed + "e", H
		} else if 94 <= r && r < 113 {
			return seed + "i", H
		} else if 113 <= r && r < 143 {
			return seed + "m", H
		} else if 143 <= r && r < 145 {
			return seed + "n", H
		} else if 145 <= r && r < 156 {
			return seed + "o", H
		} else if 156 <= r && r < 196 {
			return seed + "p", H
		} else if 196 <= r && r < 200 {
			return seed + "s", H
		} else if 200 <= r && r < 201 {
			return seed + "u", H
		} else if 201 <= r && r < 203 {
			return seed + "v", H
		} else if 203 <= r && r < 204 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wb" {
		r := rng.IntN(20)
		H := 2.3427376486136673
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sy" {
		r := rng.IntN(56)
		H := 3.5233623527669695
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 9 {
			return seed + "d", H
		} else if 9 <= r && r < 10 {
			return seed + "f", H
		} else if 10 <= r && r < 11 {
			return seed + "g", H
		} else if 11 <= r && r < 12 {
			return seed + "h", H
		} else if 12 <= r && r < 13 {
			return seed + "j", H
		} else if 13 <= r && r < 14 {
			return seed + "k", H
		} else if 14 <= r && r < 15 {
			return seed + "l", H
		} else if 15 <= r && r < 24 {
			return seed + "m", H
		} else if 24 <= r && r < 39 {
			return seed + "n", H
		} else if 39 <= r && r < 40 {
			return seed + "p", H
		} else if 40 <= r && r < 41 {
			return seed + "q", H
		} else if 41 <= r && r < 44 {
			return seed + "r", H
		} else if 44 <= r && r < 51 {
			return seed + "s", H
		} else if 51 <= r && r < 52 {
			return seed + "t", H
		} else if 52 <= r && r < 53 {
			return seed + "v", H
		} else if 53 <= r && r < 54 {
			return seed + "w", H
		} else if 54 <= r && r < 55 {
			return seed + "x", H
		} else if 55 <= r && r < 56 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hn" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "eo" {
		r := rng.IntN(66)
		H := 3.873831139409895
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 17 {
			return seed + "g", H
		} else if 17 <= r && r < 18 {
			return seed + "h", H
		} else if 18 <= r && r < 19 {
			return seed + "j", H
		} else if 19 <= r && r < 20 {
			return seed + "k", H
		} else if 20 <= r && r < 33 {
			return seed + "l", H
		} else if 33 <= r && r < 38 {
			return seed + "m", H
		} else if 38 <= r && r < 43 {
			return seed + "n", H
		} else if 43 <= r && r < 46 {
			return seed + "p", H
		} else if 46 <= r && r < 47 {
			return seed + "q", H
		} else if 47 <= r && r < 54 {
			return seed + "r", H
		} else if 54 <= r && r < 55 {
			return seed + "s", H
		} else if 55 <= r && r < 58 {
			return seed + "t", H
		} else if 58 <= r && r < 60 {
			return seed + "u", H
		} else if 60 <= r && r < 63 {
			return seed + "v", H
		} else if 63 <= r && r < 64 {
			return seed + "w", H
		} else if 64 <= r && r < 65 {
			return seed + "x", H
		} else if 65 <= r && r < 66 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pe" {
		r := rng.IntN(506)
		H := 3.029971408127494
		if 0 <= r && r < 28 {
			return seed + "a", H
		} else if 28 <= r && r < 33 {
			return seed + "b", H
		} else if 33 <= r && r < 66 {
			return seed + "c", H
		} else if 66 <= r && r < 133 {
			return seed + "d", H
		} else if 133 <= r && r < 143 {
			return seed + "e", H
		} else if 143 <= r && r < 144 {
			return seed + "f", H
		} else if 144 <= r && r < 149 {
			return seed + "g", H
		} else if 149 <= r && r < 150 {
			return seed + "h", H
		} else if 150 <= r && r < 151 {
			return seed + "j", H
		} else if 151 <= r && r < 152 {
			return seed + "k", H
		} else if 152 <= r && r < 179 {
			return seed + "l", H
		} else if 179 <= r && r < 180 {
			return seed + "m", H
		} else if 180 <= r && r < 271 {
			return seed + "n", H
		} else if 271 <= r && r < 274 {
			return seed + "p", H
		} else if 274 <= r && r < 275 {
			return seed + "q", H
		} else if 275 <= r && r < 446 {
			return seed + "r", H
		} else if 446 <= r && r < 461 {
			return seed + "s", H
		} else if 461 <= r && r < 496 {
			return seed + "t", H
		} else if 496 <= r && r < 497 {
			return seed + "v", H
		} else if 497 <= r && r < 500 {
			return seed + "w", H
		} else if 500 <= r && r < 501 {
			return seed + "x", H
		} else if 501 <= r && r < 506 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "c" {
		r := rng.IntN(1508)
		H := 2.6778392592781093
		if 0 <= r && r < 331 {
			return seed + "a", H
		} else if 331 <= r && r < 364 {
			return seed + "e", H
		} else if 364 <= r && r < 558 {
			return seed + "h", H
		} else if 558 <= r && r < 597 {
			return seed + "i", H
		} else if 597 <= r && r < 723 {
			return seed + "l", H
		} else if 723 <= r && r < 1172 {
			return seed + "o", H
		} else if 1172 <= r && r < 1386 {
			return seed + "r", H
		} else if 1386 <= r && r < 1491 {
			return seed + "u", H
		} else if 1491 <= r && r < 1508 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ii" {
		r := rng.IntN(22)
		H := 4.243300368538956
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 17 {
			return seed + "s", H
		} else if 17 <= r && r < 18 {
			return seed + "t", H
		} else if 18 <= r && r < 19 {
			return seed + "v", H
		} else if 19 <= r && r < 20 {
			return seed + "w", H
		} else if 20 <= r && r < 21 {
			return seed + "x", H
		} else if 21 <= r && r < 22 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "cn" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oy" {
		r := rng.IntN(66)
		H := 3.996929211068931
		if 0 <= r && r < 10 {
			return seed + "a", H
		} else if 10 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 25 {
			return seed + "e", H
		} else if 25 <= r && r < 28 {
			return seed + "f", H
		} else if 28 <= r && r < 29 {
			return seed + "g", H
		} else if 29 <= r && r < 30 {
			return seed + "h", H
		} else if 30 <= r && r < 34 {
			return seed + "i", H
		} else if 34 <= r && r < 35 {
			return seed + "j", H
		} else if 35 <= r && r < 36 {
			return seed + "k", H
		} else if 36 <= r && r < 39 {
			return seed + "l", H
		} else if 39 <= r && r < 42 {
			return seed + "m", H
		} else if 42 <= r && r < 45 {
			return seed + "n", H
		} else if 45 <= r && r < 51 {
			return seed + "o", H
		} else if 51 <= r && r < 52 {
			return seed + "p", H
		} else if 52 <= r && r < 53 {
			return seed + "q", H
		} else if 53 <= r && r < 56 {
			return seed + "r", H
		} else if 56 <= r && r < 61 {
			return seed + "s", H
		} else if 61 <= r && r < 62 {
			return seed + "t", H
		} else if 62 <= r && r < 63 {
			return seed + "v", H
		} else if 63 <= r && r < 64 {
			return seed + "w", H
		} else if 64 <= r && r < 65 {
			return seed + "x", H
		} else if 65 <= r && r < 66 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zm" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lu" {
		r := rng.IntN(274)
		H := 3.715400662435354
		if 0 <= r && r < 6 {
			return seed + "a", H
		} else if 6 <= r && r < 19 {
			return seed + "b", H
		} else if 19 <= r && r < 40 {
			return seed + "c", H
		} else if 40 <= r && r < 59 {
			return seed + "d", H
		} else if 59 <= r && r < 75 {
			return seed + "e", H
		} else if 75 <= r && r < 78 {
			return seed + "f", H
		} else if 78 <= r && r < 91 {
			return seed + "g", H
		} else if 91 <= r && r < 92 {
			return seed + "h", H
		} else if 92 <= r && r < 94 {
			return seed + "i", H
		} else if 94 <= r && r < 95 {
			return seed + "j", H
		} else if 95 <= r && r < 98 {
			return seed + "k", H
		} else if 98 <= r && r < 101 {
			return seed + "l", H
		} else if 101 <= r && r < 128 {
			return seed + "m", H
		} else if 128 <= r && r < 153 {
			return seed + "n", H
		} else if 153 <= r && r < 154 {
			return seed + "p", H
		} else if 154 <= r && r < 155 {
			return seed + "q", H
		} else if 155 <= r && r < 178 {
			return seed + "r", H
		} else if 178 <= r && r < 233 {
			return seed + "s", H
		} else if 233 <= r && r < 262 {
			return seed + "t", H
		} else if 262 <= r && r < 263 {
			return seed + "v", H
		} else if 263 <= r && r < 264 {
			return seed + "w", H
		} else if 264 <= r && r < 273 {
			return seed + "x", H
		} else if 273 <= r && r < 274 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "jd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ja" {
		r := rng.IntN(76)
		H := 4.093929448442851
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 24 {
			return seed + "i", H
		} else if 24 <= r && r < 25 {
			return seed + "j", H
		} else if 25 <= r && r < 26 {
			return seed + "k", H
		} else if 26 <= r && r < 29 {
			return seed + "l", H
		} else if 29 <= r && r < 34 {
			return seed + "m", H
		} else if 34 <= r && r < 39 {
			return seed + "n", H
		} else if 39 <= r && r < 40 {
			return seed + "p", H
		} else if 40 <= r && r < 41 {
			return seed + "q", H
		} else if 41 <= r && r < 48 {
			return seed + "r", H
		} else if 48 <= r && r < 51 {
			return seed + "s", H
		} else if 51 <= r && r < 52 {
			return seed + "t", H
		} else if 52 <= r && r < 56 {
			return seed + "u", H
		} else if 56 <= r && r < 59 {
			return seed + "v", H
		} else if 59 <= r && r < 68 {
			return seed + "w", H
		} else if 68 <= r && r < 69 {
			return seed + "x", H
		} else if 69 <= r && r < 73 {
			return seed + "y", H
		} else if 73 <= r && r < 76 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "jp" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kd" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 6 {
			return seed + "r", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ha" {
		r := rng.IntN(534)
		H := 3.5587087219908575
		if 0 <= r && r < 21 {
			return seed + "b", H
		} else if 21 <= r && r < 38 {
			return seed + "c", H
		} else if 38 <= r && r < 53 {
			return seed + "d", H
		} else if 53 <= r && r < 60 {
			return seed + "f", H
		} else if 60 <= r && r < 65 {
			return seed + "g", H
		} else if 65 <= r && r < 66 {
			return seed + "h", H
		} else if 66 <= r && r < 78 {
			return seed + "i", H
		} else if 78 <= r && r < 79 {
			return seed + "j", H
		} else if 79 <= r && r < 94 {
			return seed + "k", H
		} else if 94 <= r && r < 117 {
			return seed + "l", H
		} else if 117 <= r && r < 148 {
			return seed + "m", H
		} else if 148 <= r && r < 281 {
			return seed + "n", H
		} else if 281 <= r && r < 283 {
			return seed + "o", H
		} else if 283 <= r && r < 318 {
			return seed + "p", H
		} else if 318 <= r && r < 319 {
			return seed + "q", H
		} else if 319 <= r && r < 426 {
			return seed + "r", H
		} else if 426 <= r && r < 459 {
			return seed + "s", H
		} else if 459 <= r && r < 498 {
			return seed + "t", H
		} else if 498 <= r && r < 504 {
			return seed + "u", H
		} else if 504 <= r && r < 511 {
			return seed + "v", H
		} else if 511 <= r && r < 518 {
			return seed + "w", H
		} else if 518 <= r && r < 519 {
			return seed + "x", H
		} else if 519 <= r && r < 534 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "yz" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "au" {
		r := rng.IntN(158)
		H := 3.419904940570912
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 18 {
			return seed + "c", H
		} else if 18 <= r && r < 41 {
			return seed + "d", H
		} else if 41 <= r && r < 42 {
			return seed + "f", H
		} else if 42 <= r && r < 53 {
			return seed + "g", H
		} else if 53 <= r && r < 54 {
			return seed + "h", H
		} else if 54 <= r && r < 55 {
			return seed + "j", H
		} else if 55 <= r && r < 56 {
			return seed + "k", H
		} else if 56 <= r && r < 61 {
			return seed + "l", H
		} else if 61 <= r && r < 62 {
			return seed + "m", H
		} else if 62 <= r && r < 89 {
			return seed + "n", H
		} else if 89 <= r && r < 92 {
			return seed + "p", H
		} else if 92 <= r && r < 93 {
			return seed + "q", H
		} else if 93 <= r && r < 100 {
			return seed + "r", H
		} else if 100 <= r && r < 119 {
			return seed + "s", H
		} else if 119 <= r && r < 150 {
			return seed + "t", H
		} else if 150 <= r && r < 153 {
			return seed + "v", H
		} else if 153 <= r && r < 154 {
			return seed + "w", H
		} else if 154 <= r && r < 155 {
			return seed + "x", H
		} else if 155 <= r && r < 158 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ze" {
		r := rng.IntN(74)
		H := 3.4343985850816123
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 17 {
			return seed + "d", H
		} else if 17 <= r && r < 18 {
			return seed + "f", H
		} else if 18 <= r && r < 19 {
			return seed + "g", H
		} else if 19 <= r && r < 20 {
			return seed + "h", H
		} else if 20 <= r && r < 21 {
			return seed + "j", H
		} else if 21 <= r && r < 22 {
			return seed + "k", H
		} else if 22 <= r && r < 27 {
			return seed + "l", H
		} else if 27 <= r && r < 28 {
			return seed + "m", H
		} else if 28 <= r && r < 37 {
			return seed + "n", H
		} else if 37 <= r && r < 40 {
			return seed + "p", H
		} else if 40 <= r && r < 41 {
			return seed + "q", H
		} else if 41 <= r && r < 64 {
			return seed + "r", H
		} else if 64 <= r && r < 69 {
			return seed + "s", H
		} else if 69 <= r && r < 70 {
			return seed + "t", H
		} else if 70 <= r && r < 71 {
			return seed + "v", H
		} else if 71 <= r && r < 72 {
			return seed + "w", H
		} else if 72 <= r && r < 73 {
			return seed + "x", H
		} else if 73 <= r && r < 74 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bz" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yg" {
		r := rng.IntN(16)
		H := 2.7806390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 7 {
			return seed + "l", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uh" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "it" {
		r := rng.IntN(538)
		H := 3.097094866991805
		if 0 <= r && r < 77 {
			return seed + "a", H
		} else if 77 <= r && r < 103 {
			return seed + "c", H
		} else if 103 <= r && r < 190 {
			return seed + "e", H
		} else if 190 <= r && r < 200 {
			return seed + "h", H
		} else if 200 <= r && r < 257 {
			return seed + "i", H
		} else if 257 <= r && r < 265 {
			return seed + "l", H
		} else if 265 <= r && r < 267 {
			return seed + "m", H
		} else if 267 <= r && r < 290 {
			return seed + "o", H
		} else if 290 <= r && r < 300 {
			return seed + "r", H
		} else if 300 <= r && r < 304 {
			return seed + "s", H
		} else if 304 <= r && r < 350 {
			return seed + "t", H
		} else if 350 <= r && r < 379 {
			return seed + "u", H
		} else if 379 <= r && r < 530 {
			return seed + "y", H
		} else if 530 <= r && r < 538 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "q" {
		r := rng.IntN(80)
		H := 0.48241057254747405
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 79 {
			return seed + "u", H
		} else if 79 <= r && r < 80 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pl" {
		r := rng.IntN(362)
		H := 2.3617855747195007
		if 0 <= r && r < 127 {
			return seed + "a", H
		} else if 127 <= r && r < 200 {
			return seed + "e", H
		} else if 200 <= r && r < 263 {
			return seed + "i", H
		} else if 263 <= r && r < 310 {
			return seed + "o", H
		} else if 310 <= r && r < 333 {
			return seed + "u", H
		} else if 333 <= r && r < 362 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "di" {
		r := rng.IntN(596)
		H := 3.544927826239836
		if 0 <= r && r < 48 {
			return seed + "a", H
		} else if 48 <= r && r < 57 {
			return seed + "b", H
		} else if 57 <= r && r < 92 {
			return seed + "c", H
		} else if 92 <= r && r < 95 {
			return seed + "d", H
		} else if 95 <= r && r < 113 {
			return seed + "e", H
		} else if 113 <= r && r < 128 {
			return seed + "f", H
		} else if 128 <= r && r < 141 {
			return seed + "g", H
		} else if 141 <= r && r < 142 {
			return seed + "h", H
		} else if 142 <= r && r < 143 {
			return seed + "j", H
		} else if 143 <= r && r < 144 {
			return seed + "k", H
		} else if 144 <= r && r < 173 {
			return seed + "l", H
		} else if 173 <= r && r < 192 {
			return seed + "m", H
		} else if 192 <= r && r < 319 {
			return seed + "n", H
		} else if 319 <= r && r < 333 {
			return seed + "o", H
		} else if 333 <= r && r < 342 {
			return seed + "p", H
		} else if 342 <= r && r < 343 {
			return seed + "q", H
		} else if 343 <= r && r < 358 {
			return seed + "r", H
		} else if 358 <= r && r < 501 {
			return seed + "s", H
		} else if 501 <= r && r < 542 {
			return seed + "t", H
		} else if 542 <= r && r < 544 {
			return seed + "u", H
		} else if 544 <= r && r < 577 {
			return seed + "v", H
		} else if 577 <= r && r < 580 {
			return seed + "w", H
		} else if 580 <= r && r < 583 {
			return seed + "x", H
		} else if 583 <= r && r < 596 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "oh" {
		r := rng.IntN(16)
		H := 2.349601752714581
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pw" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ba" {
		r := rng.IntN(418)
		H := 3.356096104222986
		if 0 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 112 {
			return seed + "c", H
		} else if 112 <= r && r < 123 {
			return seed + "d", H
		} else if 123 <= r && r < 128 {
			return seed + "f", H
		} else if 128 <= r && r < 161 {
			return seed + "g", H
		} else if 161 <= r && r < 162 {
			return seed + "h", H
		} else if 162 <= r && r < 163 {
			return seed + "j", H
		} else if 163 <= r && r < 176 {
			return seed + "k", H
		} else if 176 <= r && r < 215 {
			return seed + "l", H
		} else if 215 <= r && r < 218 {
			return seed + "m", H
		} else if 218 <= r && r < 271 {
			return seed + "n", H
		} else if 271 <= r && r < 272 {
			return seed + "p", H
		} else if 272 <= r && r < 273 {
			return seed + "q", H
		} else if 273 <= r && r < 332 {
			return seed + "r", H
		} else if 332 <= r && r < 359 {
			return seed + "s", H
		} else if 359 <= r && r < 408 {
			return seed + "t", H
		} else if 408 <= r && r < 410 {
			return seed + "u", H
		} else if 410 <= r && r < 411 {
			return seed + "v", H
		} else if 411 <= r && r < 412 {
			return seed + "w", H
		} else if 412 <= r && r < 413 {
			return seed + "x", H
		} else if 413 <= r && r < 415 {
			return seed + "y", H
		} else if 415 <= r && r < 418 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bm" {
		r := rng.IntN(18)
		H := 2.1690354219421173
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kf" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ww" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yi" {
		r := rng.IntN(70)
		H := 2.434380669242648
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 7 {
			return seed + "e", H
		} else if 7 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 9 {
			return seed + "g", H
		} else if 9 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 11 {
			return seed + "j", H
		} else if 11 <= r && r < 12 {
			return seed + "k", H
		} else if 12 <= r && r < 13 {
			return seed + "l", H
		} else if 13 <= r && r < 14 {
			return seed + "m", H
		} else if 14 <= r && r < 59 {
			return seed + "n", H
		} else if 59 <= r && r < 62 {
			return seed + "p", H
		} else if 62 <= r && r < 63 {
			return seed + "q", H
		} else if 63 <= r && r < 64 {
			return seed + "r", H
		} else if 64 <= r && r < 65 {
			return seed + "s", H
		} else if 65 <= r && r < 66 {
			return seed + "t", H
		} else if 66 <= r && r < 67 {
			return seed + "v", H
		} else if 67 <= r && r < 68 {
			return seed + "w", H
		} else if 68 <= r && r < 69 {
			return seed + "x", H
		} else if 69 <= r && r < 70 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ya" {
		r := rng.IntN(58)
		H := 3.6804218636058126
		if 0 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 17 {
			return seed + "g", H
		} else if 17 <= r && r < 20 {
			return seed + "h", H
		} else if 20 <= r && r < 21 {
			return seed + "j", H
		} else if 21 <= r && r < 22 {
			return seed + "k", H
		} else if 22 <= r && r < 27 {
			return seed + "l", H
		} else if 27 <= r && r < 30 {
			return seed + "m", H
		} else if 30 <= r && r < 33 {
			return seed + "n", H
		} else if 33 <= r && r < 36 {
			return seed + "p", H
		} else if 36 <= r && r < 37 {
			return seed + "q", H
		} else if 37 <= r && r < 50 {
			return seed + "r", H
		} else if 50 <= r && r < 51 {
			return seed + "s", H
		} else if 51 <= r && r < 52 {
			return seed + "t", H
		} else if 52 <= r && r < 53 {
			return seed + "v", H
		} else if 53 <= r && r < 56 {
			return seed + "w", H
		} else if 56 <= r && r < 57 {
			return seed + "x", H
		} else if 57 <= r && r < 58 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "rt" {
		r := rng.IntN(250)
		H := 3.4313421831089777
		if 0 <= r && r < 37 {
			return seed + "a", H
		} else if 37 <= r && r < 41 {
			return seed + "c", H
		} else if 41 <= r && r < 90 {
			return seed + "e", H
		} else if 90 <= r && r < 92 {
			return seed + "f", H
		} else if 92 <= r && r < 116 {
			return seed + "h", H
		} else if 116 <= r && r < 151 {
			return seed + "i", H
		} else if 151 <= r && r < 175 {
			return seed + "l", H
		} else if 175 <= r && r < 185 {
			return seed + "n", H
		} else if 185 <= r && r < 198 {
			return seed + "o", H
		} else if 198 <= r && r < 202 {
			return seed + "r", H
		} else if 202 <= r && r < 212 {
			return seed + "s", H
		} else if 212 <= r && r < 229 {
			return seed + "u", H
		} else if 229 <= r && r < 233 {
			return seed + "w", H
		} else if 233 <= r && r < 246 {
			return seed + "y", H
		} else if 246 <= r && r < 250 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ny" {
		r := rng.IntN(48)
		H := 4.0298541066775275
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 11 {
			return seed + "j", H
		} else if 11 <= r && r < 12 {
			return seed + "k", H
		} else if 12 <= r && r < 15 {
			return seed + "l", H
		} else if 15 <= r && r < 22 {
			return seed + "m", H
		} else if 22 <= r && r < 23 {
			return seed + "n", H
		} else if 23 <= r && r < 27 {
			return seed + "o", H
		} else if 27 <= r && r < 30 {
			return seed + "p", H
		} else if 30 <= r && r < 31 {
			return seed + "q", H
		} else if 31 <= r && r < 32 {
			return seed + "r", H
		} else if 32 <= r && r < 33 {
			return seed + "s", H
		} else if 33 <= r && r < 38 {
			return seed + "t", H
		} else if 38 <= r && r < 39 {
			return seed + "v", H
		} else if 39 <= r && r < 44 {
			return seed + "w", H
		} else if 44 <= r && r < 47 {
			return seed + "x", H
		} else if 47 <= r && r < 48 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hi" {
		r := rng.IntN(278)
		H := 3.16844114290633
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 24 {
			return seed + "c", H
		} else if 24 <= r && r < 27 {
			return seed + "d", H
		} else if 27 <= r && r < 39 {
			return seed + "e", H
		} else if 39 <= r && r < 54 {
			return seed + "f", H
		} else if 54 <= r && r < 57 {
			return seed + "g", H
		} else if 57 <= r && r < 60 {
			return seed + "h", H
		} else if 60 <= r && r < 61 {
			return seed + "j", H
		} else if 61 <= r && r < 62 {
			return seed + "k", H
		} else if 62 <= r && r < 95 {
			return seed + "l", H
		} else if 95 <= r && r < 106 {
			return seed + "m", H
		} else if 106 <= r && r < 215 {
			return seed + "n", H
		} else if 215 <= r && r < 234 {
			return seed + "p", H
		} else if 234 <= r && r < 235 {
			return seed + "q", H
		} else if 235 <= r && r < 256 {
			return seed + "r", H
		} else if 256 <= r && r < 263 {
			return seed + "s", H
		} else if 263 <= r && r < 268 {
			return seed + "t", H
		} else if 268 <= r && r < 275 {
			return seed + "v", H
		} else if 275 <= r && r < 276 {
			return seed + "w", H
		} else if 276 <= r && r < 277 {
			return seed + "x", H
		} else if 277 <= r && r < 278 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "tv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "by" {
		r := rng.IntN(26)
		H := 4.071034795964157
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 17 {
			return seed + "s", H
		} else if 17 <= r && r < 22 {
			return seed + "t", H
		} else if 22 <= r && r < 23 {
			return seed + "v", H
		} else if 23 <= r && r < 24 {
			return seed + "w", H
		} else if 24 <= r && r < 25 {
			return seed + "x", H
		} else if 25 <= r && r < 26 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ig" {
		r := rng.IntN(224)
		H := 3.0099270594724468
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 42 {
			return seed + "e", H
		} else if 42 <= r && r < 60 {
			return seed + "g", H
		} else if 60 <= r && r < 144 {
			return seed + "h", H
		} else if 144 <= r && r < 155 {
			return seed + "i", H
		} else if 155 <= r && r < 159 {
			return seed + "l", H
		} else if 159 <= r && r < 163 {
			return seed + "m", H
		} else if 163 <= r && r < 185 {
			return seed + "n", H
		} else if 185 <= r && r < 198 {
			return seed + "o", H
		} else if 198 <= r && r < 202 {
			return seed + "r", H
		} else if 202 <= r && r < 204 {
			return seed + "s", H
		} else if 204 <= r && r < 219 {
			return seed + "u", H
		} else if 219 <= r && r < 222 {
			return seed + "y", H
		} else if 222 <= r && r < 224 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "rc" {
		r := rng.IntN(112)
		H := 2.8666613149672338
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 62 {
			return seed + "h", H
		} else if 62 <= r && r < 71 {
			return seed + "i", H
		} else if 71 <= r && r < 77 {
			return seed + "l", H
		} else if 77 <= r && r < 94 {
			return seed + "o", H
		} else if 94 <= r && r < 96 {
			return seed + "r", H
		} else if 96 <= r && r < 100 {
			return seed + "t", H
		} else if 100 <= r && r < 111 {
			return seed + "u", H
		} else if 111 <= r && r < 112 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dc" {
		r := rng.IntN(32)
		H := 2.4581887399951703
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 17 {
			return seed + "l", H
		} else if 17 <= r && r < 26 {
			return seed + "o", H
		} else if 26 <= r && r < 28 {
			return seed + "r", H
		} else if 28 <= r && r < 31 {
			return seed + "u", H
		} else if 31 <= r && r < 32 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fo" {
		r := rng.IntN(192)
		H := 3.4062880235360544
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 17 {
			return seed + "e", H
		} else if 17 <= r && r < 18 {
			return seed + "f", H
		} else if 18 <= r && r < 23 {
			return seed + "g", H
		} else if 23 <= r && r < 24 {
			return seed + "h", H
		} else if 24 <= r && r < 28 {
			return seed + "i", H
		} else if 28 <= r && r < 29 {
			return seed + "j", H
		} else if 29 <= r && r < 30 {
			return seed + "k", H
		} else if 30 <= r && r < 57 {
			return seed + "l", H
		} else if 57 <= r && r < 58 {
			return seed + "m", H
		} else if 58 <= r && r < 69 {
			return seed + "n", H
		} else if 69 <= r && r < 121 {
			return seed + "o", H
		} else if 121 <= r && r < 122 {
			return seed + "p", H
		} else if 122 <= r && r < 123 {
			return seed + "q", H
		} else if 123 <= r && r < 162 {
			return seed + "r", H
		} else if 162 <= r && r < 167 {
			return seed + "s", H
		} else if 167 <= r && r < 168 {
			return seed + "t", H
		} else if 168 <= r && r < 182 {
			return seed + "u", H
		} else if 182 <= r && r < 183 {
			return seed + "v", H
		} else if 183 <= r && r < 186 {
			return seed + "w", H
		} else if 186 <= r && r < 189 {
			return seed + "x", H
		} else if 189 <= r && r < 191 {
			return seed + "y", H
		} else if 191 <= r && r < 192 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "o" {
		r := rng.IntN(512)
		H := 3.106224254173011
		if 0 <= r && r < 8 {
			return seed + "a", H
		} else if 8 <= r && r < 63 {
			return seed + "b", H
		} else if 63 <= r && r < 86 {
			return seed + "c", H
		} else if 86 <= r && r < 87 {
			return seed + "d", H
		} else if 87 <= r && r < 88 {
			return seed + "f", H
		} else if 88 <= r && r < 91 {
			return seed + "g", H
		} else if 91 <= r && r < 92 {
			return seed + "h", H
		} else if 92 <= r && r < 98 {
			return seed + "i", H
		} else if 98 <= r && r < 99 {
			return seed + "j", H
		} else if 99 <= r && r < 102 {
			return seed + "k", H
		} else if 102 <= r && r < 109 {
			return seed + "l", H
		} else if 109 <= r && r < 122 {
			return seed + "m", H
		} else if 122 <= r && r < 153 {
			return seed + "n", H
		} else if 153 <= r && r < 159 {
			return seed + "o", H
		} else if 159 <= r && r < 198 {
			return seed + "p", H
		} else if 198 <= r && r < 199 {
			return seed + "q", H
		} else if 199 <= r && r < 200 {
			return seed + "r", H
		} else if 200 <= r && r < 203 {
			return seed + "s", H
		} else if 203 <= r && r < 208 {
			return seed + "t", H
		} else if 208 <= r && r < 314 {
			return seed + "u", H
		} else if 314 <= r && r < 489 {
			return seed + "v", H
		} else if 489 <= r && r < 492 {
			return seed + "w", H
		} else if 492 <= r && r < 507 {
			return seed + "x", H
		} else if 507 <= r && r < 509 {
			return seed + "y", H
		} else if 509 <= r && r < 512 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bg" {
		r := rng.IntN(10)
		H := 2.521928094887362
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 8 {
			return seed + "r", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oi" {
		r := rng.IntN(152)
		H := 2.8426470939738255
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 32 {
			return seed + "f", H
		} else if 32 <= r && r < 33 {
			return seed + "g", H
		} else if 33 <= r && r < 34 {
			return seed + "h", H
		} else if 34 <= r && r < 35 {
			return seed + "j", H
		} else if 35 <= r && r < 36 {
			return seed + "k", H
		} else if 36 <= r && r < 77 {
			return seed + "l", H
		} else if 77 <= r && r < 78 {
			return seed + "m", H
		} else if 78 <= r && r < 127 {
			return seed + "n", H
		} else if 127 <= r && r < 128 {
			return seed + "p", H
		} else if 128 <= r && r < 129 {
			return seed + "q", H
		} else if 129 <= r && r < 132 {
			return seed + "r", H
		} else if 132 <= r && r < 145 {
			return seed + "s", H
		} else if 145 <= r && r < 148 {
			return seed + "t", H
		} else if 148 <= r && r < 149 {
			return seed + "v", H
		} else if 149 <= r && r < 150 {
			return seed + "w", H
		} else if 150 <= r && r < 151 {
			return seed + "x", H
		} else if 151 <= r && r < 152 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "nb" {
		r := rng.IntN(56)
		H := 2.84515317626941
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 23 {
			return seed + "i", H
		} else if 23 <= r && r < 29 {
			return seed + "l", H
		} else if 29 <= r && r < 38 {
			return seed + "o", H
		} else if 38 <= r && r < 46 {
			return seed + "r", H
		} else if 46 <= r && r < 55 {
			return seed + "u", H
		} else if 55 <= r && r < 56 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vr" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "aw" {
		r := rng.IntN(76)
		H := 3.139016800131057
		if 0 <= r && r < 25 {
			return seed + "a", H
		} else if 25 <= r && r < 27 {
			return seed + "d", H
		} else if 27 <= r && r < 34 {
			return seed + "e", H
		} else if 34 <= r && r < 38 {
			return seed + "f", H
		} else if 38 <= r && r < 40 {
			return seed + "h", H
		} else if 40 <= r && r < 43 {
			return seed + "i", H
		} else if 43 <= r && r < 47 {
			return seed + "k", H
		} else if 47 <= r && r < 59 {
			return seed + "l", H
		} else if 59 <= r && r < 67 {
			return seed + "n", H
		} else if 67 <= r && r < 70 {
			return seed + "o", H
		} else if 70 <= r && r < 72 {
			return seed + "r", H
		} else if 72 <= r && r < 74 {
			return seed + "s", H
		} else if 74 <= r && r < 75 {
			return seed + "u", H
		} else if 75 <= r && r < 76 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vt" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mt" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gs" {
		r := rng.IntN(22)
		H := 2.6171895725927143
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 20 {
			return seed + "t", H
		} else if 20 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qm" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qt" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ug" {
		r := rng.IntN(74)
		H := 3.0400915499077836
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 30 {
			return seed + "g", H
		} else if 30 <= r && r < 48 {
			return seed + "h", H
		} else if 48 <= r && r < 51 {
			return seed + "i", H
		} else if 51 <= r && r < 55 {
			return seed + "l", H
		} else if 55 <= r && r < 57 {
			return seed + "m", H
		} else if 57 <= r && r < 63 {
			return seed + "n", H
		} else if 63 <= r && r < 66 {
			return seed + "o", H
		} else if 66 <= r && r < 68 {
			return seed + "s", H
		} else if 68 <= r && r < 73 {
			return seed + "u", H
		} else if 73 <= r && r < 74 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zp" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ec" {
		r := rng.IntN(384)
		H := 3.047441345171078
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 29 {
			return seed + "c", H
		} else if 29 <= r && r < 62 {
			return seed + "e", H
		} else if 62 <= r && r < 72 {
			return seed + "h", H
		} else if 72 <= r && r < 105 {
			return seed + "i", H
		} else if 105 <= r && r < 131 {
			return seed + "k", H
		} else if 131 <= r && r < 157 {
			return seed + "l", H
		} else if 157 <= r && r < 210 {
			return seed + "o", H
		} else if 210 <= r && r < 224 {
			return seed + "r", H
		} else if 224 <= r && r < 352 {
			return seed + "t", H
		} else if 352 <= r && r < 375 {
			return seed + "u", H
		} else if 375 <= r && r < 384 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hu" {
		r := rng.IntN(180)
		H := 3.5105727498827575
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 25 {
			return seed + "d", H
		} else if 25 <= r && r < 34 {
			return seed + "f", H
		} else if 34 <= r && r < 39 {
			return seed + "g", H
		} else if 39 <= r && r < 40 {
			return seed + "h", H
		} else if 40 <= r && r < 41 {
			return seed + "j", H
		} else if 41 <= r && r < 42 {
			return seed + "k", H
		} else if 42 <= r && r < 49 {
			return seed + "l", H
		} else if 49 <= r && r < 92 {
			return seed + "m", H
		} else if 92 <= r && r < 121 {
			return seed + "n", H
		} else if 121 <= r && r < 126 {
			return seed + "p", H
		} else if 126 <= r && r < 127 {
			return seed + "q", H
		} else if 127 <= r && r < 152 {
			return seed + "r", H
		} else if 152 <= r && r < 167 {
			return seed + "s", H
		} else if 167 <= r && r < 176 {
			return seed + "t", H
		} else if 176 <= r && r < 177 {
			return seed + "v", H
		} else if 177 <= r && r < 178 {
			return seed + "w", H
		} else if 178 <= r && r < 179 {
			return seed + "x", H
		} else if 179 <= r && r < 180 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "si" {
		r := rng.IntN(512)
		H := 3.806000496322532
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 25 {
			return seed + "b", H
		} else if 25 <= r && r < 36 {
			return seed + "c", H
		} else if 36 <= r && r < 79 {
			return seed + "d", H
		} else if 79 <= r && r < 89 {
			return seed + "e", H
		} else if 89 <= r && r < 94 {
			return seed + "f", H
		} else if 94 <= r && r < 117 {
			return seed + "g", H
		} else if 117 <= r && r < 118 {
			return seed + "h", H
		} else if 118 <= r && r < 119 {
			return seed + "j", H
		} else if 119 <= r && r < 120 {
			return seed + "k", H
		} else if 120 <= r && r < 163 {
			return seed + "l", H
		} else if 163 <= r && r < 182 {
			return seed + "m", H
		} else if 182 <= r && r < 269 {
			return seed + "n", H
		} else if 269 <= r && r < 333 {
			return seed + "o", H
		} else if 333 <= r && r < 338 {
			return seed + "p", H
		} else if 338 <= r && r < 339 {
			return seed + "q", H
		} else if 339 <= r && r < 344 {
			return seed + "r", H
		} else if 344 <= r && r < 377 {
			return seed + "s", H
		} else if 377 <= r && r < 428 {
			return seed + "t", H
		} else if 428 <= r && r < 430 {
			return seed + "u", H
		} else if 430 <= r && r < 475 {
			return seed + "v", H
		} else if 475 <= r && r < 476 {
			return seed + "w", H
		} else if 476 <= r && r < 489 {
			return seed + "x", H
		} else if 489 <= r && r < 512 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "oc" {
		r := rng.IntN(188)
		H := 2.8885064616784883
		if 0 <= r && r < 25 {
			return seed + "a", H
		} else if 25 <= r && r < 41 {
			return seed + "c", H
		} else if 41 <= r && r < 54 {
			return seed + "e", H
		} else if 54 <= r && r < 58 {
			return seed + "h", H
		} else if 58 <= r && r < 71 {
			return seed + "i", H
		} else if 71 <= r && r < 145 {
			return seed + "k", H
		} else if 145 <= r && r < 147 {
			return seed + "l", H
		} else if 147 <= r && r < 152 {
			return seed + "o", H
		} else if 152 <= r && r < 166 {
			return seed + "r", H
		} else if 166 <= r && r < 176 {
			return seed + "t", H
		} else if 176 <= r && r < 185 {
			return seed + "u", H
		} else if 185 <= r && r < 188 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fm" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wt" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "h" {
		r := rng.IntN(504)
		H := 1.7651454355562342
		if 0 <= r && r < 245 {
			return seed + "a", H
		} else if 245 <= r && r < 370 {
			return seed + "e", H
		} else if 370 <= r && r < 371 {
			return seed + "i", H
		} else if 371 <= r && r < 372 {
			return seed + "o", H
		} else if 372 <= r && r < 469 {
			return seed + "u", H
		} else if 469 <= r && r < 504 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ex" {
		r := rng.IntN(210)
		H := 2.888340763033017
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 51 {
			return seed + "c", H
		} else if 51 <= r && r < 68 {
			return seed + "e", H
		} else if 68 <= r && r < 70 {
			return seed + "f", H
		} else if 70 <= r && r < 76 {
			return seed + "h", H
		} else if 76 <= r && r < 93 {
			return seed + "i", H
		} else if 93 <= r && r < 102 {
			return seed + "o", H
		} else if 102 <= r && r < 162 {
			return seed + "p", H
		} else if 162 <= r && r < 164 {
			return seed + "q", H
		} else if 164 <= r && r < 204 {
			return seed + "t", H
		} else if 204 <= r && r < 209 {
			return seed + "u", H
		} else if 209 <= r && r < 210 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ax" {
		r := rng.IntN(20)
		H := 2.5261207468426807
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 18 {
			return seed + "s", H
		} else if 18 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ra" {
		r := rng.IntN(962)
		H := 3.987292018054268
		if 0 <= r && r < 39 {
			return seed + "b", H
		} else if 39 <= r && r < 112 {
			return seed + "c", H
		} else if 112 <= r && r < 179 {
			return seed + "d", H
		} else if 179 <= r && r < 181 {
			return seed + "e", H
		} else if 181 <= r && r < 210 {
			return seed + "f", H
		} else if 210 <= r && r < 271 {
			return seed + "g", H
		} else if 271 <= r && r < 272 {
			return seed + "h", H
		} else if 272 <= r && r < 340 {
			return seed + "i", H
		} else if 340 <= r && r < 341 {
			return seed + "j", H
		} else if 341 <= r && r < 350 {
			return seed + "k", H
		} else if 350 <= r && r < 413 {
			return seed + "l", H
		} else if 413 <= r && r < 478 {
			return seed + "m", H
		} else if 478 <= r && r < 617 {
			return seed + "n", H
		} else if 617 <= r && r < 619 {
			return seed + "o", H
		} else if 619 <= r && r < 666 {
			return seed + "p", H
		} else if 666 <= r && r < 667 {
			return seed + "q", H
		} else if 667 <= r && r < 686 {
			return seed + "r", H
		} else if 686 <= r && r < 731 {
			return seed + "s", H
		} else if 731 <= r && r < 844 {
			return seed + "t", H
		} else if 844 <= r && r < 848 {
			return seed + "u", H
		} else if 848 <= r && r < 899 {
			return seed + "v", H
		} else if 899 <= r && r < 916 {
			return seed + "w", H
		} else if 916 <= r && r < 919 {
			return seed + "x", H
		} else if 919 <= r && r < 949 {
			return seed + "y", H
		} else if 949 <= r && r < 962 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "gf" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gm" {
		r := rng.IntN(22)
		H := 1.8658566174572233
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 19 {
			return seed + "i", H
		} else if 19 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fl" {
		r := rng.IntN(214)
		H := 2.1966786393098277
		if 0 <= r && r < 95 {
			return seed + "a", H
		} else if 95 <= r && r < 130 {
			return seed + "e", H
		} else if 130 <= r && r < 163 {
			return seed + "i", H
		} else if 163 <= r && r < 188 {
			return seed + "o", H
		} else if 188 <= r && r < 195 {
			return seed + "u", H
		} else if 195 <= r && r < 214 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zt" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oz" {
		r := rng.IntN(38)
		H := 2.2444108662888076
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 16 {
			return seed + "e", H
		} else if 16 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 26 {
			return seed + "o", H
		} else if 26 <= r && r < 27 {
			return seed + "u", H
		} else if 27 <= r && r < 38 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rb" {
		r := rng.IntN(100)
		H := 2.78277089044991
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 28 {
			return seed + "e", H
		} else if 28 <= r && r < 51 {
			return seed + "i", H
		} else if 51 <= r && r < 59 {
			return seed + "l", H
		} else if 59 <= r && r < 84 {
			return seed + "o", H
		} else if 84 <= r && r < 86 {
			return seed + "r", H
		} else if 86 <= r && r < 88 {
			return seed + "s", H
		} else if 88 <= r && r < 93 {
			return seed + "u", H
		} else if 93 <= r && r < 100 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xq" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "am" {
		r := rng.IntN(362)
		H := 3.196553751895435
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 65 {
			return seed + "b", H
		} else if 65 <= r && r < 67 {
			return seed + "c", H
		} else if 67 <= r && r < 148 {
			return seed + "e", H
		} else if 148 <= r && r < 215 {
			return seed + "i", H
		} else if 215 <= r && r < 225 {
			return seed + "l", H
		} else if 225 <= r && r < 245 {
			return seed + "m", H
		} else if 245 <= r && r < 249 {
			return seed + "n", H
		} else if 249 <= r && r < 270 {
			return seed + "o", H
		} else if 270 <= r && r < 330 {
			return seed + "p", H
		} else if 330 <= r && r < 334 {
			return seed + "r", H
		} else if 334 <= r && r < 340 {
			return seed + "s", H
		} else if 340 <= r && r < 342 {
			return seed + "t", H
		} else if 342 <= r && r < 357 {
			return seed + "u", H
		} else if 357 <= r && r < 362 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nj" {
		r := rng.IntN(30)
		H := 1.84397154729945
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 30 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bt" {
		r := rng.IntN(34)
		H := 2.903967813087766
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 19 {
			return seed + "l", H
		} else if 19 <= r && r < 24 {
			return seed + "o", H
		} else if 24 <= r && r < 28 {
			return seed + "r", H
		} else if 28 <= r && r < 31 {
			return seed + "u", H
		} else if 31 <= r && r < 34 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dt" {
		r := rng.IntN(10)
		H := 2.5219280948873615
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jh" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xf" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tb" {
		r := rng.IntN(40)
		H := 2.327515039439817
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 14 {
			return seed + "e", H
		} else if 14 <= r && r < 19 {
			return seed + "i", H
		} else if 19 <= r && r < 34 {
			return seed + "o", H
		} else if 34 <= r && r < 36 {
			return seed + "r", H
		} else if 36 <= r && r < 39 {
			return seed + "u", H
		} else if 39 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fh" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "im" {
		r := rng.IntN(368)
		H := 2.9232182198034034
		if 0 <= r && r < 49 {
			return seed + "a", H
		} else if 49 <= r && r < 61 {
			return seed + "b", H
		} else if 61 <= r && r < 118 {
			return seed + "e", H
		} else if 118 <= r && r < 151 {
			return seed + "i", H
		} else if 151 <= r && r < 155 {
			return seed + "l", H
		} else if 155 <= r && r < 217 {
			return seed + "m", H
		} else if 217 <= r && r < 223 {
			return seed + "n", H
		} else if 223 <= r && r < 234 {
			return seed + "o", H
		} else if 234 <= r && r < 344 {
			return seed + "p", H
		} else if 344 <= r && r < 350 {
			return seed + "s", H
		} else if 350 <= r && r < 361 {
			return seed + "u", H
		} else if 361 <= r && r < 363 {
			return seed + "w", H
		} else if 363 <= r && r < 368 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tz" {
		r := rng.IntN(16)
		H := 2.0461796919474975
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yo" {
		r := rng.IntN(48)
		H := 3.9457346197190524
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 13 {
			return seed + "g", H
		} else if 13 <= r && r < 14 {
			return seed + "h", H
		} else if 14 <= r && r < 15 {
			return seed + "j", H
		} else if 15 <= r && r < 16 {
			return seed + "k", H
		} else if 16 <= r && r < 17 {
			return seed + "l", H
		} else if 17 <= r && r < 18 {
			return seed + "m", H
		} else if 18 <= r && r < 29 {
			return seed + "n", H
		} else if 29 <= r && r < 30 {
			return seed + "p", H
		} else if 30 <= r && r < 31 {
			return seed + "q", H
		} else if 31 <= r && r < 34 {
			return seed + "r", H
		} else if 34 <= r && r < 35 {
			return seed + "s", H
		} else if 35 <= r && r < 36 {
			return seed + "t", H
		} else if 36 <= r && r < 38 {
			return seed + "u", H
		} else if 38 <= r && r < 41 {
			return seed + "v", H
		} else if 41 <= r && r < 42 {
			return seed + "w", H
		} else if 42 <= r && r < 43 {
			return seed + "x", H
		} else if 43 <= r && r < 47 {
			return seed + "y", H
		} else if 47 <= r && r < 48 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "tm" {
		r := rng.IntN(22)
		H := 2.294629213844555
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lf" {
		r := rng.IntN(32)
		H := 2.3715123682506185
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 19 {
			return seed + "i", H
		} else if 19 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 24 {
			return seed + "r", H
		} else if 24 <= r && r < 31 {
			return seed + "u", H
		} else if 31 <= r && r < 32 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "u" {
		r := rng.IntN(972)
		H := 1.0333361736970286
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 17 {
			return seed + "l", H
		} else if 17 <= r && r < 24 {
			return seed + "m", H
		} else if 24 <= r && r < 845 {
			return seed + "n", H
		} else if 845 <= r && r < 918 {
			return seed + "p", H
		} else if 918 <= r && r < 919 {
			return seed + "q", H
		} else if 919 <= r && r < 938 {
			return seed + "r", H
		} else if 938 <= r && r < 955 {
			return seed + "s", H
		} else if 955 <= r && r < 968 {
			return seed + "t", H
		} else if 968 <= r && r < 969 {
			return seed + "v", H
		} else if 969 <= r && r < 970 {
			return seed + "w", H
		} else if 970 <= r && r < 971 {
			return seed + "x", H
		} else if 971 <= r && r < 972 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "as" {
		r := rng.IntN(566)
		H := 3.0381050209049225
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 31 {
			return seed + "c", H
		} else if 31 <= r && r < 86 {
			return seed + "e", H
		} else if 86 <= r && r < 198 {
			return seed + "h", H
		} else if 198 <= r && r < 259 {
			return seed + "i", H
		} else if 259 <= r && r < 275 {
			return seed + "k", H
		} else if 275 <= r && r < 277 {
			return seed + "l", H
		} else if 277 <= r && r < 285 {
			return seed + "m", H
		} else if 285 <= r && r < 294 {
			return seed + "o", H
		} else if 294 <= r && r < 320 {
			return seed + "p", H
		} else if 320 <= r && r < 322 {
			return seed + "q", H
		} else if 322 <= r && r < 396 {
			return seed + "s", H
		} else if 396 <= r && r < 554 {
			return seed + "t", H
		} else if 554 <= r && r < 561 {
			return seed + "u", H
		} else if 561 <= r && r < 566 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lg" {
		r := rng.IntN(16)
		H := 2.382856063692049
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "md" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 6 {
			return seed + "r", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ms" {
		r := rng.IntN(24)
		H := 2.7772925846688996
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 4 {
			return seed + "h", H
		} else if 4 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 18 {
			return seed + "t", H
		} else if 18 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 24 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xi" {
		r := rng.IntN(72)
		H := 3.827405609382162
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 25 {
			return seed + "d", H
		} else if 25 <= r && r < 28 {
			return seed + "f", H
		} else if 28 <= r && r < 31 {
			return seed + "g", H
		} else if 31 <= r && r < 32 {
			return seed + "h", H
		} else if 32 <= r && r < 33 {
			return seed + "j", H
		} else if 33 <= r && r < 34 {
			return seed + "k", H
		} else if 34 <= r && r < 37 {
			return seed + "l", H
		} else if 37 <= r && r < 44 {
			return seed + "m", H
		} else if 44 <= r && r < 49 {
			return seed + "n", H
		} else if 49 <= r && r < 51 {
			return seed + "o", H
		} else if 51 <= r && r < 52 {
			return seed + "p", H
		} else if 52 <= r && r < 53 {
			return seed + "q", H
		} else if 53 <= r && r < 56 {
			return seed + "r", H
		} else if 56 <= r && r < 65 {
			return seed + "s", H
		} else if 65 <= r && r < 68 {
			return seed + "t", H
		} else if 68 <= r && r < 69 {
			return seed + "v", H
		} else if 69 <= r && r < 70 {
			return seed + "w", H
		} else if 70 <= r && r < 71 {
			return seed + "x", H
		} else if 71 <= r && r < 72 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ks" {
		r := rng.IntN(40)
		H := 2.930240672752193
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 15 {
			return seed + "k", H
		} else if 15 <= r && r < 19 {
			return seed + "l", H
		} else if 19 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 24 {
			return seed + "p", H
		} else if 24 <= r && r < 38 {
			return seed + "t", H
		} else if 38 <= r && r < 39 {
			return seed + "u", H
		} else if 39 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fs" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ju" {
		r := rng.IntN(104)
		H := 3.7141403417989767
		if 0 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 17 {
			return seed + "d", H
		} else if 17 <= r && r < 18 {
			return seed + "f", H
		} else if 18 <= r && r < 25 {
			return seed + "g", H
		} else if 25 <= r && r < 26 {
			return seed + "h", H
		} else if 26 <= r && r < 32 {
			return seed + "i", H
		} else if 32 <= r && r < 35 {
			return seed + "j", H
		} else if 35 <= r && r < 38 {
			return seed + "k", H
		} else if 38 <= r && r < 41 {
			return seed + "l", H
		} else if 41 <= r && r < 48 {
			return seed + "m", H
		} else if 48 <= r && r < 65 {
			return seed + "n", H
		} else if 65 <= r && r < 66 {
			return seed + "p", H
		} else if 66 <= r && r < 67 {
			return seed + "q", H
		} else if 67 <= r && r < 84 {
			return seed + "r", H
		} else if 84 <= r && r < 97 {
			return seed + "s", H
		} else if 97 <= r && r < 98 {
			return seed + "t", H
		} else if 98 <= r && r < 101 {
			return seed + "v", H
		} else if 101 <= r && r < 102 {
			return seed + "w", H
		} else if 102 <= r && r < 103 {
			return seed + "x", H
		} else if 103 <= r && r < 104 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vh" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xh" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lk" {
		r := rng.IntN(14)
		H := 2.064042639445697
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ik" {
		r := rng.IntN(74)
		H := 1.3058231327446612
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 60 {
			return seed + "e", H
		} else if 60 <= r && r < 69 {
			return seed + "i", H
		} else if 69 <= r && r < 70 {
			return seed + "o", H
		} else if 70 <= r && r < 73 {
			return seed + "u", H
		} else if 73 <= r && r < 74 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tg" {
		r := rng.IntN(12)
		H := 2.6258145836939115
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 10 {
			return seed + "r", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qn" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "e" {
		r := rng.IntN(816)
		H := 3.508461416301628
		if 0 <= r && r < 74 {
			return seed + "a", H
		} else if 74 <= r && r < 81 {
			return seed + "b", H
		} else if 81 <= r && r < 106 {
			return seed + "c", H
		} else if 106 <= r && r < 125 {
			return seed + "d", H
		} else if 125 <= r && r < 127 {
			return seed + "e", H
		} else if 127 <= r && r < 136 {
			return seed + "f", H
		} else if 136 <= r && r < 153 {
			return seed + "g", H
		} else if 153 <= r && r < 154 {
			return seed + "h", H
		} else if 154 <= r && r < 156 {
			return seed + "i", H
		} else if 156 <= r && r < 159 {
			return seed + "j", H
		} else if 159 <= r && r < 160 {
			return seed + "k", H
		} else if 160 <= r && r < 231 {
			return seed + "l", H
		} else if 231 <= r && r < 310 {
			return seed + "m", H
		} else if 310 <= r && r < 485 {
			return seed + "n", H
		} else if 485 <= r && r < 506 {
			return seed + "p", H
		} else if 506 <= r && r < 523 {
			return seed + "q", H
		} else if 523 <= r && r < 546 {
			return seed + "r", H
		} else if 546 <= r && r < 589 {
			return seed + "s", H
		} else if 589 <= r && r < 604 {
			return seed + "t", H
		} else if 604 <= r && r < 606 {
			return seed + "u", H
		} else if 606 <= r && r < 649 {
			return seed + "v", H
		} else if 649 <= r && r < 650 {
			return seed + "w", H
		} else if 650 <= r && r < 815 {
			return seed + "x", H
		} else if 815 <= r && r < 816 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pb" {
		r := rng.IntN(14)
		H := 2.2988252450030506
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wd" {
		r := rng.IntN(24)
		H := 2.5314536459234778
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 15 {
			return seed + "l", H
		} else if 15 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 22 {
			return seed + "r", H
		} else if 22 <= r && r < 23 {
			return seed + "u", H
		} else if 23 <= r && r < 24 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pm" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ao" {
		r := rng.IntN(26)
		H := 4.151798852506845
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 18 {
			return seed + "r", H
		} else if 18 <= r && r < 21 {
			return seed + "s", H
		} else if 21 <= r && r < 22 {
			return seed + "t", H
		} else if 22 <= r && r < 23 {
			return seed + "v", H
		} else if 23 <= r && r < 24 {
			return seed + "w", H
		} else if 24 <= r && r < 25 {
			return seed + "x", H
		} else if 25 <= r && r < 26 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vu" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mp" {
		r := rng.IntN(292)
		H := 3.2807035237847755
		if 0 <= r && r < 31 {
			return seed + "a", H
		} else if 31 <= r && r < 33 {
			return seed + "b", H
		} else if 33 <= r && r < 72 {
			return seed + "e", H
		} else if 72 <= r && r < 74 {
			return seed + "f", H
		} else if 74 <= r && r < 90 {
			return seed + "h", H
		} else if 90 <= r && r < 121 {
			return seed + "i", H
		} else if 121 <= r && r < 179 {
			return seed + "l", H
		} else if 179 <= r && r < 181 {
			return seed + "n", H
		} else if 181 <= r && r < 222 {
			return seed + "o", H
		} else if 222 <= r && r < 246 {
			return seed + "r", H
		} else if 246 <= r && r < 254 {
			return seed + "s", H
		} else if 254 <= r && r < 276 {
			return seed + "t", H
		} else if 276 <= r && r < 291 {
			return seed + "u", H
		} else if 291 <= r && r < 292 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ia" {
		r := rng.IntN(232)
		H := 3.116890149742443
		if 0 <= r && r < 37 {
			return seed + "b", H
		} else if 37 <= r && r < 44 {
			return seed + "c", H
		} else if 44 <= r && r < 45 {
			return seed + "d", H
		} else if 45 <= r && r < 46 {
			return seed + "f", H
		} else if 46 <= r && r < 55 {
			return seed + "g", H
		} else if 55 <= r && r < 56 {
			return seed + "h", H
		} else if 56 <= r && r < 57 {
			return seed + "j", H
		} else if 57 <= r && r < 58 {
			return seed + "k", H
		} else if 58 <= r && r < 111 {
			return seed + "l", H
		} else if 111 <= r && r < 116 {
			return seed + "m", H
		} else if 116 <= r && r < 161 {
			return seed + "n", H
		} else if 161 <= r && r < 166 {
			return seed + "p", H
		} else if 166 <= r && r < 167 {
			return seed + "q", H
		} else if 167 <= r && r < 186 {
			return seed + "r", H
		} else if 186 <= r && r < 191 {
			return seed + "s", H
		} else if 191 <= r && r < 228 {
			return seed + "t", H
		} else if 228 <= r && r < 229 {
			return seed + "v", H
		} else if 229 <= r && r < 230 {
			return seed + "w", H
		} else if 230 <= r && r < 231 {
			return seed + "x", H
		} else if 231 <= r && r < 232 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "za" {
		r := rng.IntN(44)
		H := 3.595298361893484
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 10 {
			return seed + "c", H
		} else if 10 <= r && r < 11 {
			return seed + "d", H
		} else if 11 <= r && r < 12 {
			return seed + "f", H
		} else if 12 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 17 {
			return seed + "j", H
		} else if 17 <= r && r < 18 {
			return seed + "k", H
		} else if 18 <= r && r < 19 {
			return seed + "l", H
		} else if 19 <= r && r < 20 {
			return seed + "m", H
		} else if 20 <= r && r < 21 {
			return seed + "n", H
		} else if 21 <= r && r < 24 {
			return seed + "p", H
		} else if 24 <= r && r < 25 {
			return seed + "q", H
		} else if 25 <= r && r < 38 {
			return seed + "r", H
		} else if 38 <= r && r < 39 {
			return seed + "s", H
		} else if 39 <= r && r < 40 {
			return seed + "t", H
		} else if 40 <= r && r < 41 {
			return seed + "v", H
		} else if 41 <= r && r < 42 {
			return seed + "w", H
		} else if 42 <= r && r < 43 {
			return seed + "x", H
		} else if 43 <= r && r < 44 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fa" {
		r := rng.IntN(180)
		H := 3.5357511168441373
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 54 {
			return seed + "c", H
		} else if 54 <= r && r < 59 {
			return seed + "d", H
		} else if 59 <= r && r < 60 {
			return seed + "f", H
		} else if 60 <= r && r < 61 {
			return seed + "g", H
		} else if 61 <= r && r < 62 {
			return seed + "h", H
		} else if 62 <= r && r < 70 {
			return seed + "i", H
		} else if 70 <= r && r < 71 {
			return seed + "j", H
		} else if 71 <= r && r < 72 {
			return seed + "k", H
		} else if 72 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 104 {
			return seed + "m", H
		} else if 104 <= r && r < 131 {
			return seed + "n", H
		} else if 131 <= r && r < 132 {
			return seed + "p", H
		} else if 132 <= r && r < 133 {
			return seed + "q", H
		} else if 133 <= r && r < 140 {
			return seed + "r", H
		} else if 140 <= r && r < 155 {
			return seed + "s", H
		} else if 155 <= r && r < 158 {
			return seed + "t", H
		} else if 158 <= r && r < 162 {
			return seed + "u", H
		} else if 162 <= r && r < 173 {
			return seed + "v", H
		} else if 173 <= r && r < 174 {
			return seed + "w", H
		} else if 174 <= r && r < 177 {
			return seed + "x", H
		} else if 177 <= r && r < 180 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "af" {
		r := rng.IntN(124)
		H := 2.385424396085014
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 22 {
			return seed + "e", H
		} else if 22 <= r && r < 72 {
			return seed + "f", H
		} else if 72 <= r && r < 73 {
			return seed + "i", H
		} else if 73 <= r && r < 79 {
			return seed + "l", H
		} else if 79 <= r && r < 81 {
			return seed + "n", H
		} else if 81 <= r && r < 84 {
			return seed + "o", H
		} else if 84 <= r && r < 88 {
			return seed + "r", H
		} else if 88 <= r && r < 122 {
			return seed + "t", H
		} else if 122 <= r && r < 123 {
			return seed + "u", H
		} else if 123 <= r && r < 124 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bn" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "li" {
		r := rng.IntN(970)
		H := 3.5979335281871765
		if 0 <= r && r < 42 {
			return seed + "a", H
		} else if 42 <= r && r < 61 {
			return seed + "b", H
		} else if 61 <= r && r < 148 {
			return seed + "c", H
		} else if 148 <= r && r < 171 {
			return seed + "d", H
		} else if 171 <= r && r < 199 {
			return seed + "e", H
		} else if 199 <= r && r < 236 {
			return seed + "f", H
		} else if 236 <= r && r < 295 {
			return seed + "g", H
		} else if 295 <= r && r < 296 {
			return seed + "h", H
		} else if 296 <= r && r < 297 {
			return seed + "j", H
		} else if 297 <= r && r < 356 {
			return seed + "k", H
		} else if 356 <= r && r < 363 {
			return seed + "l", H
		} else if 363 <= r && r < 400 {
			return seed + "m", H
		} else if 400 <= r && r < 715 {
			return seed + "n", H
		} else if 715 <= r && r < 735 {
			return seed + "o", H
		} else if 735 <= r && r < 756 {
			return seed + "p", H
		} else if 756 <= r && r < 765 {
			return seed + "q", H
		} else if 765 <= r && r < 772 {
			return seed + "r", H
		} else if 772 <= r && r < 827 {
			return seed + "s", H
		} else if 827 <= r && r < 898 {
			return seed + "t", H
		} else if 898 <= r && r < 902 {
			return seed + "u", H
		} else if 902 <= r && r < 935 {
			return seed + "v", H
		} else if 935 <= r && r < 936 {
			return seed + "w", H
		} else if 936 <= r && r < 939 {
			return seed + "x", H
		} else if 939 <= r && r < 970 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "kj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ss" {
		r := rng.IntN(148)
		H := 2.833316861688887
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 17 {
			return seed + "c", H
		} else if 17 <= r && r < 56 {
			return seed + "e", H
		} else if 56 <= r && r < 58 {
			return seed + "f", H
		} else if 58 <= r && r < 97 {
			return seed + "i", H
		} else if 97 <= r && r < 103 {
			return seed + "l", H
		} else if 103 <= r && r < 114 {
			return seed + "o", H
		} else if 114 <= r && r < 118 {
			return seed + "p", H
		} else if 118 <= r && r < 137 {
			return seed + "u", H
		} else if 137 <= r && r < 139 {
			return seed + "w", H
		} else if 139 <= r && r < 148 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jt" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "p" {
		r := rng.IntN(1162)
		H := 2.605231600876813
		if 0 <= r && r < 273 {
			return seed + "a", H
		} else if 273 <= r && r < 410 {
			return seed + "e", H
		} else if 410 <= r && r < 434 {
			return seed + "h", H
		} else if 434 <= r && r < 435 {
			return seed + "i", H
		} else if 435 <= r && r < 547 {
			return seed + "l", H
		} else if 547 <= r && r < 702 {
			return seed + "o", H
		} else if 702 <= r && r < 1034 {
			return seed + "r", H
		} else if 1034 <= r && r < 1036 {
			return seed + "s", H
		} else if 1036 <= r && r < 1155 {
			return seed + "u", H
		} else if 1155 <= r && r < 1162 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tn" {
		r := rng.IntN(42)
		H := 1.719549195651597
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 35 {
			return seed + "i", H
		} else if 35 <= r && r < 38 {
			return seed + "o", H
		} else if 38 <= r && r < 41 {
			return seed + "u", H
		} else if 41 <= r && r < 42 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zi" {
		r := rng.IntN(86)
		H := 2.908737225471599
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 7 {
			return seed + "e", H
		} else if 7 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 11 {
			return seed + "g", H
		} else if 11 <= r && r < 12 {
			return seed + "h", H
		} else if 12 <= r && r < 13 {
			return seed + "j", H
		} else if 13 <= r && r < 14 {
			return seed + "k", H
		} else if 14 <= r && r < 23 {
			return seed + "l", H
		} else if 23 <= r && r < 24 {
			return seed + "m", H
		} else if 24 <= r && r < 65 {
			return seed + "n", H
		} else if 65 <= r && r < 76 {
			return seed + "p", H
		} else if 76 <= r && r < 77 {
			return seed + "q", H
		} else if 77 <= r && r < 78 {
			return seed + "r", H
		} else if 78 <= r && r < 79 {
			return seed + "s", H
		} else if 79 <= r && r < 82 {
			return seed + "t", H
		} else if 82 <= r && r < 83 {
			return seed + "v", H
		} else if 83 <= r && r < 84 {
			return seed + "w", H
		} else if 84 <= r && r < 85 {
			return seed + "x", H
		} else if 85 <= r && r < 86 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "gd" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "of" {
		r := rng.IntN(58)
		H := 3.1639348173800292
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 32 {
			return seed + "f", H
		} else if 32 <= r && r < 39 {
			return seed + "i", H
		} else if 39 <= r && r < 44 {
			return seed + "o", H
		} else if 44 <= r && r < 46 {
			return seed + "r", H
		} else if 46 <= r && r < 48 {
			return seed + "s", H
		} else if 48 <= r && r < 52 {
			return seed + "t", H
		} else if 52 <= r && r < 55 {
			return seed + "u", H
		} else if 55 <= r && r < 58 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cb" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "is" {
		r := rng.IntN(684)
		H := 3.446733002570817
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 27 {
			return seed + "b", H
		} else if 27 <= r && r < 53 {
			return seed + "c", H
		} else if 53 <= r && r < 57 {
			return seed + "d", H
		} else if 57 <= r && r < 116 {
			return seed + "e", H
		} else if 116 <= r && r < 122 {
			return seed + "f", H
		} else if 122 <= r && r < 124 {
			return seed + "g", H
		} else if 124 <= r && r < 252 {
			return seed + "h", H
		} else if 252 <= r && r < 299 {
			return seed + "i", H
		} else if 299 <= r && r < 301 {
			return seed + "j", H
		} else if 301 <= r && r < 317 {
			return seed + "k", H
		} else if 317 <= r && r < 333 {
			return seed + "l", H
		} else if 333 <= r && r < 393 {
			return seed + "m", H
		} else if 393 <= r && r < 428 {
			return seed + "o", H
		} else if 428 <= r && r < 472 {
			return seed + "p", H
		} else if 472 <= r && r < 476 {
			return seed + "r", H
		} else if 476 <= r && r < 508 {
			return seed + "s", H
		} else if 508 <= r && r < 676 {
			return seed + "t", H
		} else if 676 <= r && r < 677 {
			return seed + "u", H
		} else if 677 <= r && r < 679 {
			return seed + "w", H
		} else if 679 <= r && r < 684 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cp" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tu" {
		r := rng.IntN(290)
		H := 3.4629809233173017
		if 0 <= r && r < 26 {
			return seed + "a", H
		} else if 26 <= r && r < 43 {
			return seed + "b", H
		} else if 43 <= r && r < 54 {
			return seed + "c", H
		} else if 54 <= r && r < 77 {
			return seed + "d", H
		} else if 77 <= r && r < 81 {
			return seed + "e", H
		} else if 81 <= r && r < 92 {
			return seed + "f", H
		} else if 92 <= r && r < 95 {
			return seed + "g", H
		} else if 95 <= r && r < 96 {
			return seed + "h", H
		} else if 96 <= r && r < 100 {
			return seed + "i", H
		} else if 100 <= r && r < 101 {
			return seed + "j", H
		} else if 101 <= r && r < 102 {
			return seed + "k", H
		} else if 102 <= r && r < 107 {
			return seed + "l", H
		} else if 107 <= r && r < 126 {
			return seed + "m", H
		} else if 126 <= r && r < 143 {
			return seed + "n", H
		} else if 143 <= r && r < 145 {
			return seed + "o", H
		} else if 145 <= r && r < 156 {
			return seed + "p", H
		} else if 156 <= r && r < 157 {
			return seed + "q", H
		} else if 157 <= r && r < 258 {
			return seed + "r", H
		} else if 258 <= r && r < 271 {
			return seed + "s", H
		} else if 271 <= r && r < 284 {
			return seed + "t", H
		} else if 284 <= r && r < 285 {
			return seed + "v", H
		} else if 285 <= r && r < 286 {
			return seed + "w", H
		} else if 286 <= r && r < 289 {
			return seed + "x", H
		} else if 289 <= r && r < 290 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pc" {
		r := rng.IntN(18)
		H := 2.2096909728231857
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tt" {
		r := rng.IntN(260)
		H := 2.3621821342637057
		if 0 <= r && r < 21 {
			return seed + "a", H
		} else if 21 <= r && r < 140 {
			return seed + "e", H
		} else if 140 <= r && r < 142 {
			return seed + "h", H
		} else if 142 <= r && r < 185 {
			return seed + "i", H
		} else if 185 <= r && r < 217 {
			return seed + "l", H
		} else if 217 <= r && r < 238 {
			return seed + "o", H
		} else if 238 <= r && r < 242 {
			return seed + "r", H
		} else if 242 <= r && r < 245 {
			return seed + "u", H
		} else if 245 <= r && r < 260 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xl" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sl" {
		r := rng.IntN(176)
		H := 2.334161366430973
		if 0 <= r && r < 49 {
			return seed + "a", H
		} else if 49 <= r && r < 76 {
			return seed + "e", H
		} else if 76 <= r && r < 129 {
			return seed + "i", H
		} else if 129 <= r && r < 156 {
			return seed + "o", H
		} else if 156 <= r && r < 167 {
			return seed + "u", H
		} else if 167 <= r && r < 176 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dr" {
		r := rng.IntN(200)
		H := 2.283013977766452
		if 0 <= r && r < 55 {
			return seed + "a", H
		} else if 55 <= r && r < 102 {
			return seed + "e", H
		} else if 102 <= r && r < 145 {
			return seed + "i", H
		} else if 145 <= r && r < 186 {
			return seed + "o", H
		} else if 186 <= r && r < 195 {
			return seed + "u", H
		} else if 195 <= r && r < 200 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "av" {
		r := rng.IntN(198)
		H := 2.0047183594164117
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 110 {
			return seed + "e", H
		} else if 110 <= r && r < 167 {
			return seed + "i", H
		} else if 167 <= r && r < 190 {
			return seed + "o", H
		} else if 190 <= r && r < 191 {
			return seed + "u", H
		} else if 191 <= r && r < 198 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tf" {
		r := rng.IntN(40)
		H := 2.0496422893170583
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 13 {
			return seed + "l", H
		} else if 13 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 39 {
			return seed + "u", H
		} else if 39 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "i" {
		r := rng.IntN(250)
		H := 2.875211201949062
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 41 {
			return seed + "d", H
		} else if 41 <= r && r < 42 {
			return seed + "f", H
		} else if 42 <= r && r < 51 {
			return seed + "g", H
		} else if 51 <= r && r < 52 {
			return seed + "h", H
		} else if 52 <= r && r < 53 {
			return seed + "j", H
		} else if 53 <= r && r < 54 {
			return seed + "k", H
		} else if 54 <= r && r < 61 {
			return seed + "l", H
		} else if 61 <= r && r < 180 {
			return seed + "m", H
		} else if 180 <= r && r < 181 {
			return seed + "n", H
		} else if 181 <= r && r < 187 {
			return seed + "o", H
		} else if 187 <= r && r < 194 {
			return seed + "p", H
		} else if 194 <= r && r < 195 {
			return seed + "q", H
		} else if 195 <= r && r < 214 {
			return seed + "r", H
		} else if 214 <= r && r < 231 {
			return seed + "s", H
		} else if 231 <= r && r < 242 {
			return seed + "t", H
		} else if 242 <= r && r < 247 {
			return seed + "v", H
		} else if 247 <= r && r < 248 {
			return seed + "w", H
		} else if 248 <= r && r < 249 {
			return seed + "x", H
		} else if 249 <= r && r < 250 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pa" {
		r := rng.IntN(448)
		H := 3.6412524865520517
		if 0 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 56 {
			return seed + "c", H
		} else if 56 <= r && r < 75 {
			return seed + "d", H
		} else if 75 <= r && r < 76 {
			return seed + "f", H
		} else if 76 <= r && r < 85 {
			return seed + "g", H
		} else if 85 <= r && r < 86 {
			return seed + "h", H
		} else if 86 <= r && r < 104 {
			return seed + "i", H
		} else if 104 <= r && r < 107 {
			return seed + "j", H
		} else if 107 <= r && r < 108 {
			return seed + "k", H
		} else if 108 <= r && r < 129 {
			return seed + "l", H
		} else if 129 <= r && r < 138 {
			return seed + "m", H
		} else if 138 <= r && r < 191 {
			return seed + "n", H
		} else if 191 <= r && r < 206 {
			return seed + "p", H
		} else if 206 <= r && r < 207 {
			return seed + "q", H
		} else if 207 <= r && r < 304 {
			return seed + "r", H
		} else if 304 <= r && r < 355 {
			return seed + "s", H
		} else if 355 <= r && r < 394 {
			return seed + "t", H
		} else if 394 <= r && r < 396 {
			return seed + "u", H
		} else if 396 <= r && r < 411 {
			return seed + "v", H
		} else if 411 <= r && r < 414 {
			return seed + "w", H
		} else if 414 <= r && r < 415 {
			return seed + "x", H
		} else if 415 <= r && r < 447 {
			return seed + "y", H
		} else if 447 <= r && r < 448 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ah" {
		r := rng.IntN(20)
		H := 2.077235445408308
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rw" {
		r := rng.IntN(22)
		H := 2.4198513730338034
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 17 {
			return seed + "i", H
		} else if 17 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 20 {
			return seed + "r", H
		} else if 20 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "td" {
		r := rng.IntN(16)
		H := 1.9197367178034825
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dy" {
		r := rng.IntN(40)
		H := 3.9837834090041
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 12 {
			return seed + "i", H
		} else if 12 <= r && r < 13 {
			return seed + "j", H
		} else if 13 <= r && r < 14 {
			return seed + "k", H
		} else if 14 <= r && r < 17 {
			return seed + "l", H
		} else if 17 <= r && r < 20 {
			return seed + "m", H
		} else if 20 <= r && r < 27 {
			return seed + "n", H
		} else if 27 <= r && r < 28 {
			return seed + "p", H
		} else if 28 <= r && r < 29 {
			return seed + "q", H
		} else if 29 <= r && r < 30 {
			return seed + "r", H
		} else if 30 <= r && r < 35 {
			return seed + "s", H
		} else if 35 <= r && r < 36 {
			return seed + "t", H
		} else if 36 <= r && r < 37 {
			return seed + "v", H
		} else if 37 <= r && r < 38 {
			return seed + "w", H
		} else if 38 <= r && r < 39 {
			return seed + "x", H
		} else if 39 <= r && r < 40 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "gk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ep" {
		r := rng.IntN(192)
		H := 3.326250252829282
		if 0 <= r && r < 23 {
			return seed + "a", H
		} else if 23 <= r && r < 38 {
			return seed + "e", H
		} else if 38 <= r && r < 44 {
			return seed + "h", H
		} else if 44 <= r && r < 69 {
			return seed + "i", H
		} else if 69 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 93 {
			return seed + "n", H
		} else if 93 <= r && r < 106 {
			return seed + "o", H
		} else if 106 <= r && r < 110 {
			return seed + "p", H
		} else if 110 <= r && r < 136 {
			return seed + "r", H
		} else if 136 <= r && r < 142 {
			return seed + "s", H
		} else if 142 <= r && r < 174 {
			return seed + "t", H
		} else if 174 <= r && r < 191 {
			return seed + "u", H
		} else if 191 <= r && r < 192 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ee" {
		r := rng.IntN(272)
		H := 3.892148428429431
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 22 {
			return seed + "c", H
		} else if 22 <= r && r < 63 {
			return seed + "d", H
		} else if 63 <= r && r < 70 {
			return seed + "f", H
		} else if 70 <= r && r < 73 {
			return seed + "g", H
		} else if 73 <= r && r < 76 {
			return seed + "h", H
		} else if 76 <= r && r < 82 {
			return seed + "i", H
		} else if 82 <= r && r < 83 {
			return seed + "j", H
		} else if 83 <= r && r < 90 {
			return seed + "k", H
		} else if 90 <= r && r < 109 {
			return seed + "l", H
		} else if 109 <= r && r < 122 {
			return seed + "m", H
		} else if 122 <= r && r < 163 {
			return seed + "n", H
		} else if 163 <= r && r < 190 {
			return seed + "p", H
		} else if 190 <= r && r < 191 {
			return seed + "q", H
		} else if 191 <= r && r < 208 {
			return seed + "r", H
		} else if 208 <= r && r < 215 {
			return seed + "s", H
		} else if 215 <= r && r < 236 {
			return seed + "t", H
		} else if 236 <= r && r < 239 {
			return seed + "v", H
		} else if 239 <= r && r < 246 {
			return seed + "w", H
		} else if 246 <= r && r < 249 {
			return seed + "x", H
		} else if 249 <= r && r < 272 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ph" {
		r := rng.IntN(118)
		H := 2.55866166483463
		if 0 <= r && r < 25 {
			return seed + "a", H
		} else if 25 <= r && r < 50 {
			return seed + "e", H
		} else if 50 <= r && r < 63 {
			return seed + "i", H
		} else if 63 <= r && r < 65 {
			return seed + "l", H
		} else if 65 <= r && r < 100 {
			return seed + "o", H
		} else if 100 <= r && r < 108 {
			return seed + "r", H
		} else if 108 <= r && r < 111 {
			return seed + "u", H
		} else if 111 <= r && r < 118 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jm" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "me" {
		r := rng.IntN(282)
		H := 3.0131674517504305
		if 0 <= r && r < 10 {
			return seed + "a", H
		} else if 10 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 59 {
			return seed + "d", H
		} else if 59 <= r && r < 60 {
			return seed + "f", H
		} else if 60 <= r && r < 65 {
			return seed + "g", H
		} else if 65 <= r && r < 66 {
			return seed + "h", H
		} else if 66 <= r && r < 67 {
			return seed + "j", H
		} else if 67 <= r && r < 68 {
			return seed + "k", H
		} else if 68 <= r && r < 85 {
			return seed + "l", H
		} else if 85 <= r && r < 88 {
			return seed + "m", H
		} else if 88 <= r && r < 165 {
			return seed + "n", H
		} else if 165 <= r && r < 167 {
			return seed + "o", H
		} else if 167 <= r && r < 168 {
			return seed + "p", H
		} else if 168 <= r && r < 169 {
			return seed + "q", H
		} else if 169 <= r && r < 234 {
			return seed + "r", H
		} else if 234 <= r && r < 253 {
			return seed + "s", H
		} else if 253 <= r && r < 276 {
			return seed + "t", H
		} else if 276 <= r && r < 279 {
			return seed + "v", H
		} else if 279 <= r && r < 280 {
			return seed + "w", H
		} else if 280 <= r && r < 281 {
			return seed + "x", H
		} else if 281 <= r && r < 282 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "yk" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gt" {
		r := rng.IntN(10)
		H := 2.5219280948873615
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ku" {
		r := rng.IntN(30)
		H := 4.0444138295776115
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 15 {
			return seed + "n", H
		} else if 15 <= r && r < 20 {
			return seed + "p", H
		} else if 20 <= r && r < 21 {
			return seed + "q", H
		} else if 21 <= r && r < 22 {
			return seed + "r", H
		} else if 22 <= r && r < 25 {
			return seed + "s", H
		} else if 25 <= r && r < 26 {
			return seed + "t", H
		} else if 26 <= r && r < 27 {
			return seed + "v", H
		} else if 27 <= r && r < 28 {
			return seed + "w", H
		} else if 28 <= r && r < 29 {
			return seed + "x", H
		} else if 29 <= r && r < 30 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "og" {
		r := rng.IntN(126)
		H := 3.2279806860360964
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 30 {
			return seed + "e", H
		} else if 30 <= r && r < 46 {
			return seed + "g", H
		} else if 46 <= r && r < 59 {
			return seed + "i", H
		} else if 59 <= r && r < 63 {
			return seed + "l", H
		} else if 63 <= r && r < 71 {
			return seed + "n", H
		} else if 71 <= r && r < 74 {
			return seed + "o", H
		} else if 74 <= r && r < 90 {
			return seed + "r", H
		} else if 90 <= r && r < 103 {
			return seed + "u", H
		} else if 103 <= r && r < 105 {
			return seed + "w", H
		} else if 105 <= r && r < 126 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rs" {
		r := rng.IntN(164)
		H := 3.0278165791036242
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 60 {
			return seed + "e", H
		} else if 60 <= r && r < 76 {
			return seed + "h", H
		} else if 76 <= r && r < 101 {
			return seed + "i", H
		} else if 101 <= r && r < 105 {
			return seed + "l", H
		} else if 105 <= r && r < 107 {
			return seed + "m", H
		} else if 107 <= r && r < 109 {
			return seed + "n", H
		} else if 109 <= r && r < 116 {
			return seed + "o", H
		} else if 116 <= r && r < 118 {
			return seed + "p", H
		} else if 118 <= r && r < 146 {
			return seed + "t", H
		} else if 146 <= r && r < 161 {
			return seed + "u", H
		} else if 161 <= r && r < 163 {
			return seed + "w", H
		} else if 163 <= r && r < 164 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nk" {
		r := rng.IntN(128)
		H := 3.0441306178518825
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 36 {
			return seed + "e", H
		} else if 36 <= r && r < 40 {
			return seed + "h", H
		} else if 40 <= r && r < 75 {
			return seed + "i", H
		} else if 75 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 95 {
			return seed + "m", H
		} else if 95 <= r && r < 105 {
			return seed + "n", H
		} else if 105 <= r && r < 106 {
			return seed + "o", H
		} else if 106 <= r && r < 108 {
			return seed + "r", H
		} else if 108 <= r && r < 114 {
			return seed + "s", H
		} else if 114 <= r && r < 115 {
			return seed + "u", H
		} else if 115 <= r && r < 128 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zr" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "r" {
		r := rng.IntN(1032)
		H := 1.5835331069827565
		if 0 <= r && r < 113 {
			return seed + "a", H
		} else if 113 <= r && r < 804 {
			return seed + "e", H
		} else if 804 <= r && r < 806 {
			return seed + "h", H
		} else if 806 <= r && r < 895 {
			return seed + "i", H
		} else if 895 <= r && r < 974 {
			return seed + "o", H
		} else if 974 <= r && r < 1031 {
			return seed + "u", H
		} else if 1031 <= r && r < 1032 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lc" {
		r := rng.IntN(26)
		H := 2.669180336866731
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 18 {
			return seed + "r", H
		} else if 18 <= r && r < 25 {
			return seed + "u", H
		} else if 25 <= r && r < 26 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "br" {
		r := rng.IntN(180)
		H := 2.2576987496979575
		if 0 <= r && r < 23 {
			return seed + "a", H
		} else if 23 <= r && r < 60 {
			return seed + "e", H
		} else if 60 <= r && r < 107 {
			return seed + "i", H
		} else if 107 <= r && r < 160 {
			return seed + "o", H
		} else if 160 <= r && r < 179 {
			return seed + "u", H
		} else if 179 <= r && r < 180 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hm" {
		r := rng.IntN(10)
		H := 2.3709505944546683
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gl" {
		r := rng.IntN(230)
		H := 2.4537258030541063
		if 0 <= r && r < 39 {
			return seed + "a", H
		} else if 39 <= r && r < 110 {
			return seed + "e", H
		} else if 110 <= r && r < 147 {
			return seed + "i", H
		} else if 147 <= r && r < 180 {
			return seed + "o", H
		} else if 180 <= r && r < 195 {
			return seed + "u", H
		} else if 195 <= r && r < 230 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kr" {
		r := rng.IntN(16)
		H := 2.1774212838293647
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pp" {
		r := rng.IntN(234)
		H := 2.4784142461690775
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 92 {
			return seed + "e", H
		} else if 92 <= r && r < 145 {
			return seed + "i", H
		} else if 145 <= r && r < 175 {
			return seed + "l", H
		} else if 175 <= r && r < 192 {
			return seed + "o", H
		} else if 192 <= r && r < 210 {
			return seed + "r", H
		} else if 210 <= r && r < 212 {
			return seed + "s", H
		} else if 212 <= r && r < 213 {
			return seed + "u", H
		} else if 213 <= r && r < 234 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yf" {
		r := rng.IntN(16)
		H := 2.5550365325772653
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 7 {
			return seed + "l", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "em" {
		r := rng.IntN(276)
		H := 2.928439834918273
		if 0 <= r && r < 31 {
			return seed + "a", H
		} else if 31 <= r && r < 69 {
			return seed + "b", H
		} else if 69 <= r && r < 71 {
			return seed + "c", H
		} else if 71 <= r && r < 116 {
			return seed + "e", H
		} else if 116 <= r && r < 173 {
			return seed + "i", H
		} else if 173 <= r && r < 177 {
			return seed + "l", H
		} else if 177 <= r && r < 179 {
			return seed + "n", H
		} else if 179 <= r && r < 220 {
			return seed + "o", H
		} else if 220 <= r && r < 264 {
			return seed + "p", H
		} else if 264 <= r && r < 266 {
			return seed + "s", H
		} else if 266 <= r && r < 271 {
			return seed + "u", H
		} else if 271 <= r && r < 276 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "df" {
		r := rng.IntN(30)
		H := 2.14197056912214
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 17 {
			return seed + "i", H
		} else if 17 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 30 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "iy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mc" {
		r := rng.IntN(12)
		H := 2.625814583693911
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ni" {
		r := rng.IntN(424)
		H := 3.5893541021010518
		if 0 <= r && r < 20 {
			return seed + "a", H
		} else if 20 <= r && r < 23 {
			return seed + "b", H
		} else if 23 <= r && r < 66 {
			return seed + "c", H
		} else if 66 <= r && r < 69 {
			return seed + "d", H
		} else if 69 <= r && r < 71 {
			return seed + "e", H
		} else if 71 <= r && r < 96 {
			return seed + "f", H
		} else if 96 <= r && r < 101 {
			return seed + "g", H
		} else if 101 <= r && r < 102 {
			return seed + "h", H
		} else if 102 <= r && r < 103 {
			return seed + "j", H
		} else if 103 <= r && r < 104 {
			return seed + "k", H
		} else if 104 <= r && r < 113 {
			return seed + "l", H
		} else if 113 <= r && r < 136 {
			return seed + "m", H
		} else if 136 <= r && r < 241 {
			return seed + "n", H
		} else if 241 <= r && r < 259 {
			return seed + "o", H
		} else if 259 <= r && r < 270 {
			return seed + "p", H
		} else if 270 <= r && r < 273 {
			return seed + "q", H
		} else if 273 <= r && r < 274 {
			return seed + "r", H
		} else if 274 <= r && r < 331 {
			return seed + "s", H
		} else if 331 <= r && r < 380 {
			return seed + "t", H
		} else if 380 <= r && r < 390 {
			return seed + "u", H
		} else if 390 <= r && r < 401 {
			return seed + "v", H
		} else if 401 <= r && r < 402 {
			return seed + "w", H
		} else if 402 <= r && r < 405 {
			return seed + "x", H
		} else if 405 <= r && r < 424 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vm" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ho" {
		r := rng.IntN(274)
		H := 3.7407894410304046
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 7 {
			return seed + "d", H
		} else if 7 <= r && r < 13 {
			return seed + "e", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 17 {
			return seed + "g", H
		} else if 17 <= r && r < 18 {
			return seed + "h", H
		} else if 18 <= r && r < 20 {
			return seed + "i", H
		} else if 20 <= r && r < 21 {
			return seed + "j", H
		} else if 21 <= r && r < 26 {
			return seed + "k", H
		} else if 26 <= r && r < 51 {
			return seed + "l", H
		} else if 51 <= r && r < 56 {
			return seed + "m", H
		} else if 56 <= r && r < 89 {
			return seed + "n", H
		} else if 89 <= r && r < 123 {
			return seed + "o", H
		} else if 123 <= r && r < 136 {
			return seed + "p", H
		} else if 136 <= r && r < 137 {
			return seed + "q", H
		} else if 137 <= r && r < 178 {
			return seed + "r", H
		} else if 178 <= r && r < 187 {
			return seed + "s", H
		} else if 187 <= r && r < 198 {
			return seed + "t", H
		} else if 198 <= r && r < 226 {
			return seed + "u", H
		} else if 226 <= r && r < 233 {
			return seed + "v", H
		} else if 233 <= r && r < 270 {
			return seed + "w", H
		} else if 270 <= r && r < 271 {
			return seed + "x", H
		} else if 271 <= r && r < 273 {
			return seed + "y", H
		} else if 273 <= r && r < 274 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fp" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qp" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ub" {
		r := rng.IntN(184)
		H := 3.741521090236069
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 49 {
			return seed + "b", H
		} else if 49 <= r && r < 51 {
			return seed + "c", H
		} else if 51 <= r && r < 59 {
			return seed + "d", H
		} else if 59 <= r && r < 66 {
			return seed + "e", H
		} else if 66 <= r && r < 68 {
			return seed + "f", H
		} else if 68 <= r && r < 70 {
			return seed + "g", H
		} else if 70 <= r && r < 74 {
			return seed + "h", H
		} else if 74 <= r && r < 85 {
			return seed + "i", H
		} else if 85 <= r && r < 87 {
			return seed + "j", H
		} else if 87 <= r && r < 107 {
			return seed + "l", H
		} else if 107 <= r && r < 117 {
			return seed + "m", H
		} else if 117 <= r && r < 118 {
			return seed + "o", H
		} else if 118 <= r && r < 126 {
			return seed + "p", H
		} else if 126 <= r && r < 130 {
			return seed + "r", H
		} else if 130 <= r && r < 154 {
			return seed + "s", H
		} else if 154 <= r && r < 170 {
			return seed + "t", H
		} else if 170 <= r && r < 175 {
			return seed + "u", H
		} else if 175 <= r && r < 179 {
			return seed + "w", H
		} else if 179 <= r && r < 182 {
			return seed + "y", H
		} else if 182 <= r && r < 184 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bo" {
		r := rng.IntN(326)
		H := 3.8483893325938223
		if 0 <= r && r < 36 {
			return seed + "a", H
		} else if 36 <= r && r < 51 {
			return seed + "b", H
		} else if 51 <= r && r < 52 {
			return seed + "c", H
		} else if 52 <= r && r < 67 {
			return seed + "d", H
		} else if 67 <= r && r < 69 {
			return seed + "e", H
		} else if 69 <= r && r < 72 {
			return seed + "f", H
		} else if 72 <= r && r < 79 {
			return seed + "g", H
		} else if 79 <= r && r < 80 {
			return seed + "h", H
		} else if 80 <= r && r < 82 {
			return seed + "i", H
		} else if 82 <= r && r < 85 {
			return seed + "j", H
		} else if 85 <= r && r < 88 {
			return seed + "k", H
		} else if 88 <= r && r < 101 {
			return seed + "l", H
		} else if 101 <= r && r < 102 {
			return seed + "m", H
		} else if 102 <= r && r < 143 {
			return seed + "n", H
		} else if 143 <= r && r < 195 {
			return seed + "o", H
		} else if 195 <= r && r < 196 {
			return seed + "p", H
		} else if 196 <= r && r < 197 {
			return seed + "q", H
		} else if 197 <= r && r < 226 {
			return seed + "r", H
		} else if 226 <= r && r < 233 {
			return seed + "s", H
		} else if 233 <= r && r < 254 {
			return seed + "t", H
		} else if 254 <= r && r < 278 {
			return seed + "u", H
		} else if 278 <= r && r < 281 {
			return seed + "v", H
		} else if 281 <= r && r < 288 {
			return seed + "w", H
		} else if 288 <= r && r < 321 {
			return seed + "x", H
		} else if 321 <= r && r < 325 {
			return seed + "y", H
		} else if 325 <= r && r < 326 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hd" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ta" {
		r := rng.IntN(668)
		H := 3.7092486779779312
		if 0 <= r && r < 79 {
			return seed + "b", H
		} else if 79 <= r && r < 122 {
			return seed + "c", H
		} else if 122 <= r && r < 131 {
			return seed + "d", H
		} else if 131 <= r && r < 133 {
			return seed + "e", H
		} else if 133 <= r && r < 138 {
			return seed + "f", H
		} else if 138 <= r && r < 183 {
			return seed + "g", H
		} else if 183 <= r && r < 184 {
			return seed + "h", H
		} else if 184 <= r && r < 242 {
			return seed + "i", H
		} else if 242 <= r && r < 243 {
			return seed + "j", H
		} else if 243 <= r && r < 264 {
			return seed + "k", H
		} else if 264 <= r && r < 349 {
			return seed + "l", H
		} else if 349 <= r && r < 370 {
			return seed + "m", H
		} else if 370 <= r && r < 437 {
			return seed + "n", H
		} else if 437 <= r && r < 460 {
			return seed + "p", H
		} else if 460 <= r && r < 461 {
			return seed + "q", H
		} else if 461 <= r && r < 552 {
			return seed + "r", H
		} else if 552 <= r && r < 577 {
			return seed + "s", H
		} else if 577 <= r && r < 650 {
			return seed + "t", H
		} else if 650 <= r && r < 654 {
			return seed + "u", H
		} else if 654 <= r && r < 657 {
			return seed + "v", H
		} else if 657 <= r && r < 660 {
			return seed + "w", H
		} else if 660 <= r && r < 663 {
			return seed + "x", H
		} else if 663 <= r && r < 667 {
			return seed + "y", H
		} else if 667 <= r && r < 668 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bl" {
		r := rng.IntN(592)
		H := 1.8451178309343863
		if 0 <= r && r < 39 {
			return seed + "a", H
		} else if 39 <= r && r < 394 {
			return seed + "e", H
		} else if 394 <= r && r < 467 {
			return seed + "i", H
		} else if 467 <= r && r < 498 {
			return seed + "o", H
		} else if 498 <= r && r < 521 {
			return seed + "u", H
		} else if 521 <= r && r < 592 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qb" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ts" {
		r := rng.IntN(50)
		H := 3.24456920311317
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 3 {
			return seed + "c", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 23 {
			return seed + "i", H
		} else if 23 <= r && r < 27 {
			return seed + "k", H
		} else if 27 <= r && r < 33 {
			return seed + "m", H
		} else if 33 <= r && r < 38 {
			return seed + "o", H
		} else if 38 <= r && r < 40 {
			return seed + "p", H
		} else if 40 <= r && r < 45 {
			return seed + "u", H
		} else if 45 <= r && r < 50 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ql" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yu" {
		r := rng.IntN(22)
		H := 4.243300368538956
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 17 {
			return seed + "s", H
		} else if 17 <= r && r < 18 {
			return seed + "t", H
		} else if 18 <= r && r < 19 {
			return seed + "v", H
		} else if 19 <= r && r < 20 {
			return seed + "w", H
		} else if 20 <= r && r < 21 {
			return seed + "x", H
		} else if 21 <= r && r < 22 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "hp" {
		r := rng.IntN(10)
		H := 2.646439344671015
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 8 {
			return seed + "r", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gb" {
		r := rng.IntN(14)
		H := 2.2988252450030515
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yp" {
		r := rng.IntN(56)
		H := 2.880738916786822
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 24 {
			return seed + "h", H
		} else if 24 <= r && r < 27 {
			return seed + "i", H
		} else if 27 <= r && r < 29 {
			return seed + "l", H
		} else if 29 <= r && r < 41 {
			return seed + "n", H
		} else if 41 <= r && r < 46 {
			return seed + "o", H
		} else if 46 <= r && r < 52 {
			return seed + "t", H
		} else if 52 <= r && r < 55 {
			return seed + "u", H
		} else if 55 <= r && r < 56 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "kg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ez" {
		r := rng.IntN(34)
		H := 2.30747737374971
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 25 {
			return seed + "i", H
		} else if 25 <= r && r < 28 {
			return seed + "o", H
		} else if 28 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 32 {
			return seed + "y", H
		} else if 32 <= r && r < 34 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "nf" {
		r := rng.IntN(108)
		H := 2.646816746847481
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 55 {
			return seed + "i", H
		} else if 55 <= r && r < 63 {
			return seed + "l", H
		} else if 63 <= r && r < 84 {
			return seed + "o", H
		} else if 84 <= r && r < 94 {
			return seed + "r", H
		} else if 94 <= r && r < 107 {
			return seed + "u", H
		} else if 107 <= r && r < 108 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hc" {
		r := rng.IntN(14)
		H := 2.692380602454975
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zu" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "m" {
		r := rng.IntN(596)
		H := 1.634377172951412
		if 0 <= r && r < 261 {
			return seed + "a", H
		} else if 261 <= r && r < 262 {
			return seed + "e", H
		} else if 262 <= r && r < 263 {
			return seed + "i", H
		} else if 263 <= r && r < 478 {
			return seed + "o", H
		} else if 478 <= r && r < 585 {
			return seed + "u", H
		} else if 585 <= r && r < 596 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ok" {
		r := rng.IntN(78)
		H := 1.7253375086830145
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 52 {
			return seed + "e", H
		} else if 52 <= r && r < 67 {
			return seed + "i", H
		} else if 67 <= r && r < 68 {
			return seed + "o", H
		} else if 68 <= r && r < 71 {
			return seed + "u", H
		} else if 71 <= r && r < 78 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fn" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lb" {
		r := rng.IntN(20)
		H := 2.690468570732828
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rl" {
		r := rng.IntN(100)
		H := 2.3816873745792693
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 34 {
			return seed + "e", H
		} else if 34 <= r && r < 65 {
			return seed + "i", H
		} else if 65 <= r && r < 78 {
			return seed + "o", H
		} else if 78 <= r && r < 79 {
			return seed + "u", H
		} else if 79 <= r && r < 100 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sj" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yw" {
		r := rng.IntN(20)
		H := 2.3199730940219747
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 12 {
			return seed + "h", H
		} else if 12 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bh" {
		r := rng.IntN(10)
		H := 2.3709505944546683
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ch" {
		r := rng.IntN(400)
		H := 2.74840130371753
		if 0 <= r && r < 107 {
			return seed + "a", H
		} else if 107 <= r && r < 117 {
			return seed + "b", H
		} else if 117 <= r && r < 230 {
			return seed + "e", H
		} else if 230 <= r && r < 301 {
			return seed + "i", H
		} else if 301 <= r && r < 303 {
			return seed + "k", H
		} else if 303 <= r && r < 313 {
			return seed + "l", H
		} else if 313 <= r && r < 317 {
			return seed + "m", H
		} else if 317 <= r && r < 321 {
			return seed + "n", H
		} else if 321 <= r && r < 358 {
			return seed + "o", H
		} else if 358 <= r && r < 362 {
			return seed + "r", H
		} else if 362 <= r && r < 364 {
			return seed + "t", H
		} else if 364 <= r && r < 385 {
			return seed + "u", H
		} else if 385 <= r && r < 387 {
			return seed + "w", H
		} else if 387 <= r && r < 400 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ll" {
		r := rng.IntN(280)
		H := 2.9574223279561234
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 29 {
			return seed + "b", H
		} else if 29 <= r && r < 33 {
			return seed + "d", H
		} else if 33 <= r && r < 94 {
			return seed + "e", H
		} else if 94 <= r && r < 102 {
			return seed + "f", H
		} else if 102 <= r && r < 106 {
			return seed + "h", H
		} else if 106 <= r && r < 165 {
			return seed + "i", H
		} else if 165 <= r && r < 169 {
			return seed + "n", H
		} else if 169 <= r && r < 190 {
			return seed + "o", H
		} else if 190 <= r && r < 194 {
			return seed + "p", H
		} else if 194 <= r && r < 196 {
			return seed + "r", H
		} else if 196 <= r && r < 202 {
			return seed + "s", H
		} else if 202 <= r && r < 209 {
			return seed + "u", H
		} else if 209 <= r && r < 211 {
			return seed + "w", H
		} else if 211 <= r && r < 280 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mr" {
		r := rng.IntN(14)
		H := 2.2988252450030515
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xy" {
		r := rng.IntN(24)
		H := 4.1887218755408675
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 14 {
			return seed + "m", H
		} else if 14 <= r && r < 15 {
			return seed + "n", H
		} else if 15 <= r && r < 16 {
			return seed + "p", H
		} else if 16 <= r && r < 17 {
			return seed + "q", H
		} else if 17 <= r && r < 18 {
			return seed + "r", H
		} else if 18 <= r && r < 19 {
			return seed + "s", H
		} else if 19 <= r && r < 20 {
			return seed + "t", H
		} else if 20 <= r && r < 21 {
			return seed + "v", H
		} else if 21 <= r && r < 22 {
			return seed + "w", H
		} else if 22 <= r && r < 23 {
			return seed + "x", H
		} else if 23 <= r && r < 24 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zy" {
		r := rng.IntN(22)
		H := 4.243300368538956
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 17 {
			return seed + "s", H
		} else if 17 <= r && r < 18 {
			return seed + "t", H
		} else if 18 <= r && r < 19 {
			return seed + "v", H
		} else if 19 <= r && r < 20 {
			return seed + "w", H
		} else if 20 <= r && r < 21 {
			return seed + "x", H
		} else if 21 <= r && r < 22 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "km" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pg" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 6 {
			return seed + "r", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oo" {
		r := rng.IntN(410)
		H := 3.614128962015573
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 33 {
			return seed + "e", H
		} else if 33 <= r && r < 58 {
			return seed + "f", H
		} else if 58 <= r && r < 69 {
			return seed + "g", H
		} else if 69 <= r && r < 70 {
			return seed + "h", H
		} else if 70 <= r && r < 74 {
			return seed + "i", H
		} else if 74 <= r && r < 75 {
			return seed + "j", H
		} else if 75 <= r && r < 130 {
			return seed + "k", H
		} else if 130 <= r && r < 155 {
			return seed + "l", H
		} else if 155 <= r && r < 194 {
			return seed + "m", H
		} else if 194 <= r && r < 245 {
			return seed + "n", H
		} else if 245 <= r && r < 266 {
			return seed + "p", H
		} else if 266 <= r && r < 267 {
			return seed + "q", H
		} else if 267 <= r && r < 294 {
			return seed + "r", H
		} else if 294 <= r && r < 307 {
			return seed + "s", H
		} else if 307 <= r && r < 388 {
			return seed + "t", H
		} else if 388 <= r && r < 395 {
			return seed + "v", H
		} else if 395 <= r && r < 396 {
			return seed + "w", H
		} else if 396 <= r && r < 397 {
			return seed + "x", H
		} else if 397 <= r && r < 410 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lw" {
		r := rng.IntN(12)
		H := 2.4508257945180882
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mu" {
		r := rng.IntN(180)
		H := 3.384494902744932
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 18 {
			return seed + "f", H
		} else if 18 <= r && r < 29 {
			return seed + "g", H
		} else if 29 <= r && r < 30 {
			return seed + "h", H
		} else if 30 <= r && r < 31 {
			return seed + "j", H
		} else if 31 <= r && r < 32 {
			return seed + "k", H
		} else if 32 <= r && r < 61 {
			return seed + "l", H
		} else if 61 <= r && r < 78 {
			return seed + "m", H
		} else if 78 <= r && r < 89 {
			return seed + "n", H
		} else if 89 <= r && r < 92 {
			return seed + "p", H
		} else if 92 <= r && r < 93 {
			return seed + "q", H
		} else if 93 <= r && r < 106 {
			return seed + "r", H
		} else if 106 <= r && r < 151 {
			return seed + "s", H
		} else if 151 <= r && r < 174 {
			return seed + "t", H
		} else if 174 <= r && r < 175 {
			return seed + "v", H
		} else if 175 <= r && r < 176 {
			return seed + "w", H
		} else if 176 <= r && r < 177 {
			return seed + "x", H
		} else if 177 <= r && r < 180 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ut" {
		r := rng.IntN(310)
		H := 3.855768705190133
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 31 {
			return seed + "b", H
		} else if 31 <= r && r < 41 {
			return seed + "c", H
		} else if 41 <= r && r < 47 {
			return seed + "d", H
		} else if 47 <= r && r < 102 {
			return seed + "e", H
		} else if 102 <= r && r < 108 {
			return seed + "f", H
		} else if 108 <= r && r < 112 {
			return seed + "g", H
		} else if 112 <= r && r < 128 {
			return seed + "h", H
		} else if 128 <= r && r < 173 {
			return seed + "i", H
		} else if 173 <= r && r < 189 {
			return seed + "l", H
		} else if 189 <= r && r < 197 {
			return seed + "m", H
		} else if 197 <= r && r < 199 {
			return seed + "n", H
		} else if 199 <= r && r < 220 {
			return seed + "o", H
		} else if 220 <= r && r < 228 {
			return seed + "p", H
		} else if 228 <= r && r < 240 {
			return seed + "r", H
		} else if 240 <= r && r < 264 {
			return seed + "s", H
		} else if 264 <= r && r < 294 {
			return seed + "t", H
		} else if 294 <= r && r < 299 {
			return seed + "u", H
		} else if 299 <= r && r < 305 {
			return seed + "w", H
		} else if 305 <= r && r < 310 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kb" {
		r := rng.IntN(20)
		H := 2.1822287189138017
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uo" {
		r := rng.IntN(38)
		H := 3.622964198657942
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 24 {
			return seed + "t", H
		} else if 24 <= r && r < 34 {
			return seed + "u", H
		} else if 34 <= r && r < 35 {
			return seed + "v", H
		} else if 35 <= r && r < 36 {
			return seed + "w", H
		} else if 36 <= r && r < 37 {
			return seed + "x", H
		} else if 37 <= r && r < 38 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ct" {
		r := rng.IntN(242)
		H := 2.615098213435302
		if 0 <= r && r < 27 {
			return seed + "a", H
		} else if 27 <= r && r < 60 {
			return seed + "e", H
		} else if 60 <= r && r < 62 {
			return seed + "f", H
		} else if 62 <= r && r < 151 {
			return seed + "i", H
		} else if 151 <= r && r < 159 {
			return seed + "l", H
		} else if 159 <= r && r < 161 {
			return seed + "m", H
		} else if 161 <= r && r < 208 {
			return seed + "o", H
		} else if 208 <= r && r < 214 {
			return seed + "r", H
		} else if 214 <= r && r < 220 {
			return seed + "s", H
		} else if 220 <= r && r < 241 {
			return seed + "u", H
		} else if 241 <= r && r < 242 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uc" {
		r := rng.IntN(180)
		H := 2.9259061484383784
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 15 {
			return seed + "c", H
		} else if 15 <= r && r < 34 {
			return seed + "e", H
		} else if 34 <= r && r < 60 {
			return seed + "h", H
		} else if 60 <= r && r < 71 {
			return seed + "i", H
		} else if 71 <= r && r < 133 {
			return seed + "k", H
		} else if 133 <= r && r < 139 {
			return seed + "l", H
		} else if 139 <= r && r < 144 {
			return seed + "o", H
		} else if 144 <= r && r < 146 {
			return seed + "r", H
		} else if 146 <= r && r < 172 {
			return seed + "t", H
		} else if 172 <= r && r < 177 {
			return seed + "u", H
		} else if 177 <= r && r < 180 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "np" {
		r := rng.IntN(48)
		H := 2.4908632937969326
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 27 {
			return seed + "i", H
		} else if 27 <= r && r < 39 {
			return seed + "l", H
		} else if 39 <= r && r < 44 {
			return seed + "o", H
		} else if 44 <= r && r < 46 {
			return seed + "r", H
		} else if 46 <= r && r < 47 {
			return seed + "u", H
		} else if 47 <= r && r < 48 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ow" {
		r := rng.IntN(210)
		H := 3.64443982610508
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 21 {
			return seed + "b", H
		} else if 21 <= r && r < 25 {
			return seed + "c", H
		} else if 25 <= r && r < 43 {
			return seed + "d", H
		} else if 43 <= r && r < 76 {
			return seed + "e", H
		} else if 76 <= r && r < 82 {
			return seed + "f", H
		} else if 82 <= r && r < 84 {
			return seed + "g", H
		} else if 84 <= r && r < 103 {
			return seed + "i", H
		} else if 103 <= r && r < 121 {
			return seed + "l", H
		} else if 121 <= r && r < 127 {
			return seed + "m", H
		} else if 127 <= r && r < 171 {
			return seed + "n", H
		} else if 171 <= r && r < 174 {
			return seed + "o", H
		} else if 174 <= r && r < 180 {
			return seed + "p", H
		} else if 180 <= r && r < 184 {
			return seed + "r", H
		} else if 184 <= r && r < 198 {
			return seed + "s", H
		} else if 198 <= r && r < 200 {
			return seed + "t", H
		} else if 200 <= r && r < 201 {
			return seed + "u", H
		} else if 201 <= r && r < 205 {
			return seed + "w", H
		} else if 205 <= r && r < 210 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "la" {
		r := rng.IntN(806)
		H := 3.793900755827733
		if 0 <= r && r < 25 {
			return seed + "b", H
		} else if 25 <= r && r < 68 {
			return seed + "c", H
		} else if 68 <= r && r < 103 {
			return seed + "d", H
		} else if 103 <= r && r < 104 {
			return seed + "f", H
		} else if 104 <= r && r < 127 {
			return seed + "g", H
		} else if 127 <= r && r < 130 {
			return seed + "h", H
		} else if 130 <= r && r < 158 {
			return seed + "i", H
		} else if 158 <= r && r < 159 {
			return seed + "j", H
		} else if 159 <= r && r < 170 {
			return seed + "k", H
		} else if 170 <= r && r < 171 {
			return seed + "l", H
		} else if 171 <= r && r < 212 {
			return seed + "m", H
		} else if 212 <= r && r < 341 {
			return seed + "n", H
		} else if 341 <= r && r < 378 {
			return seed + "p", H
		} else if 378 <= r && r < 379 {
			return seed + "q", H
		} else if 379 <= r && r < 464 {
			return seed + "r", H
		} else if 464 <= r && r < 541 {
			return seed + "s", H
		} else if 541 <= r && r < 678 {
			return seed + "t", H
		} else if 678 <= r && r < 702 {
			return seed + "u", H
		} else if 702 <= r && r < 717 {
			return seed + "v", H
		} else if 717 <= r && r < 730 {
			return seed + "w", H
		} else if 730 <= r && r < 737 {
			return seed + "x", H
		} else if 737 <= r && r < 785 {
			return seed + "y", H
		} else if 785 <= r && r < 806 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xn" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ge" {
		r := rng.IntN(284)
		H := 3.255777563088015
		if 0 <= r && r < 6 {
			return seed + "a", H
		} else if 6 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 12 {
			return seed + "c", H
		} else if 12 <= r && r < 69 {
			return seed + "d", H
		} else if 69 <= r && r < 79 {
			return seed + "e", H
		} else if 79 <= r && r < 82 {
			return seed + "f", H
		} else if 82 <= r && r < 83 {
			return seed + "g", H
		} else if 83 <= r && r < 84 {
			return seed + "h", H
		} else if 84 <= r && r < 86 {
			return seed + "i", H
		} else if 86 <= r && r < 87 {
			return seed + "j", H
		} else if 87 <= r && r < 88 {
			return seed + "k", H
		} else if 88 <= r && r < 99 {
			return seed + "l", H
		} else if 99 <= r && r < 102 {
			return seed + "m", H
		} else if 102 <= r && r < 159 {
			return seed + "n", H
		} else if 159 <= r && r < 173 {
			return seed + "o", H
		} else if 173 <= r && r < 174 {
			return seed + "p", H
		} else if 174 <= r && r < 175 {
			return seed + "q", H
		} else if 175 <= r && r < 244 {
			return seed + "r", H
		} else if 244 <= r && r < 257 {
			return seed + "s", H
		} else if 257 <= r && r < 278 {
			return seed + "t", H
		} else if 278 <= r && r < 279 {
			return seed + "v", H
		} else if 279 <= r && r < 280 {
			return seed + "w", H
		} else if 280 <= r && r < 281 {
			return seed + "x", H
		} else if 281 <= r && r < 283 {
			return seed + "y", H
		} else if 283 <= r && r < 284 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bc" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ar" {
		r := rng.IntN(1032)
		H := 4.030274813835981
		if 0 <= r && r < 65 {
			return seed + "a", H
		} else if 65 <= r && r < 97 {
			return seed + "b", H
		} else if 97 <= r && r < 131 {
			return seed + "c", H
		} else if 131 <= r && r < 289 {
			return seed + "d", H
		} else if 289 <= r && r < 362 {
			return seed + "e", H
		} else if 362 <= r && r < 378 {
			return seed + "f", H
		} else if 378 <= r && r < 428 {
			return seed + "g", H
		} else if 428 <= r && r < 430 {
			return seed + "h", H
		} else if 430 <= r && r < 515 {
			return seed + "i", H
		} else if 515 <= r && r < 545 {
			return seed + "k", H
		} else if 545 <= r && r < 587 {
			return seed + "l", H
		} else if 587 <= r && r < 669 {
			return seed + "m", H
		} else if 669 <= r && r < 703 {
			return seed + "n", H
		} else if 703 <= r && r < 728 {
			return seed + "o", H
		} else if 728 <= r && r < 754 {
			return seed + "p", H
		} else if 754 <= r && r < 820 {
			return seed + "r", H
		} else if 820 <= r && r < 842 {
			return seed + "s", H
		} else if 842 <= r && r < 948 {
			return seed + "t", H
		} else if 948 <= r && r < 949 {
			return seed + "u", H
		} else if 949 <= r && r < 963 {
			return seed + "v", H
		} else if 963 <= r && r < 969 {
			return seed + "w", H
		} else if 969 <= r && r < 971 {
			return seed + "x", H
		} else if 971 <= r && r < 1032 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bp" {
		r := rng.IntN(14)
		H := 2.692380602454975
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 12 {
			return seed + "r", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xp" {
		r := rng.IntN(66)
		H := 2.6227085054880908
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 26 {
			return seed + "e", H
		} else if 26 <= r && r < 33 {
			return seed + "i", H
		} else if 33 <= r && r < 47 {
			return seed + "l", H
		} else if 47 <= r && r < 60 {
			return seed + "o", H
		} else if 60 <= r && r < 62 {
			return seed + "r", H
		} else if 62 <= r && r < 65 {
			return seed + "u", H
		} else if 65 <= r && r < 66 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "py" {
		r := rng.IntN(30)
		H := 3.9348486136508467
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 22 {
			return seed + "r", H
		} else if 22 <= r && r < 23 {
			return seed + "s", H
		} else if 23 <= r && r < 26 {
			return seed + "t", H
		} else if 26 <= r && r < 27 {
			return seed + "v", H
		} else if 27 <= r && r < 28 {
			return seed + "w", H
		} else if 28 <= r && r < 29 {
			return seed + "x", H
		} else if 29 <= r && r < 30 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "to" {
		r := rng.IntN(420)
		H := 3.1094255099001855
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 21 {
			return seed + "d", H
		} else if 21 <= r && r < 23 {
			return seed + "e", H
		} else if 23 <= r && r < 26 {
			return seed + "f", H
		} else if 26 <= r && r < 29 {
			return seed + "g", H
		} else if 29 <= r && r < 30 {
			return seed + "h", H
		} else if 30 <= r && r < 34 {
			return seed + "i", H
		} else if 34 <= r && r < 35 {
			return seed + "j", H
		} else if 35 <= r && r < 38 {
			return seed + "k", H
		} else if 38 <= r && r < 47 {
			return seed + "l", H
		} else if 47 <= r && r < 84 {
			return seed + "m", H
		} else if 84 <= r && r < 153 {
			return seed + "n", H
		} else if 153 <= r && r < 173 {
			return seed + "o", H
		} else if 173 <= r && r < 214 {
			return seed + "p", H
		} else if 214 <= r && r < 215 {
			return seed + "q", H
		} else if 215 <= r && r < 380 {
			return seed + "r", H
		} else if 380 <= r && r < 383 {
			return seed + "s", H
		} else if 383 <= r && r < 388 {
			return seed + "t", H
		} else if 388 <= r && r < 400 {
			return seed + "u", H
		} else if 400 <= r && r < 403 {
			return seed + "v", H
		} else if 403 <= r && r < 412 {
			return seed + "w", H
		} else if 412 <= r && r < 417 {
			return seed + "x", H
		} else if 417 <= r && r < 420 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ff" {
		r := rng.IntN(132)
		H := 2.755119832048778
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 38 {
			return seed + "e", H
		} else if 38 <= r && r < 67 {
			return seed + "i", H
		} else if 67 <= r && r < 95 {
			return seed + "l", H
		} else if 95 <= r && r < 97 {
			return seed + "n", H
		} else if 97 <= r && r < 108 {
			return seed + "o", H
		} else if 108 <= r && r < 114 {
			return seed + "r", H
		} else if 114 <= r && r < 123 {
			return seed + "u", H
		} else if 123 <= r && r < 132 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qe" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "wk" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 7 {
			return seed + "w", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ce" {
		r := rng.IntN(196)
		H := 3.660624401482931
		if 0 <= r && r < 8 {
			return seed + "a", H
		} else if 8 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 14 {
			return seed + "c", H
		} else if 14 <= r && r < 41 {
			return seed + "d", H
		} else if 41 <= r && r < 47 {
			return seed + "e", H
		} else if 47 <= r && r < 50 {
			return seed + "f", H
		} else if 50 <= r && r < 51 {
			return seed + "g", H
		} else if 51 <= r && r < 52 {
			return seed + "h", H
		} else if 52 <= r && r < 60 {
			return seed + "i", H
		} else if 60 <= r && r < 61 {
			return seed + "j", H
		} else if 61 <= r && r < 62 {
			return seed + "k", H
		} else if 62 <= r && r < 85 {
			return seed + "l", H
		} else if 85 <= r && r < 92 {
			return seed + "m", H
		} else if 92 <= r && r < 119 {
			return seed + "n", H
		} else if 119 <= r && r < 132 {
			return seed + "p", H
		} else if 132 <= r && r < 133 {
			return seed + "q", H
		} else if 133 <= r && r < 168 {
			return seed + "r", H
		} else if 168 <= r && r < 185 {
			return seed + "s", H
		} else if 185 <= r && r < 192 {
			return seed + "t", H
		} else if 192 <= r && r < 193 {
			return seed + "v", H
		} else if 193 <= r && r < 194 {
			return seed + "w", H
		} else if 194 <= r && r < 195 {
			return seed + "x", H
		} else if 195 <= r && r < 196 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "in" {
		r := rng.IntN(2106)
		H := 2.112655114453052
		if 0 <= r && r < 53 {
			return seed + "a", H
		} else if 53 <= r && r < 81 {
			return seed + "c", H
		} else if 81 <= r && r < 131 {
			return seed + "d", H
		} else if 131 <= r && r < 428 {
			return seed + "e", H
		} else if 428 <= r && r < 436 {
			return seed + "f", H
		} else if 436 <= r && r < 1766 {
			return seed + "g", H
		} else if 1766 <= r && r < 1768 {
			return seed + "h", H
		} else if 1768 <= r && r < 1835 {
			return seed + "i", H
		} else if 1835 <= r && r < 1839 {
			return seed + "j", H
		} else if 1839 <= r && r < 1899 {
			return seed + "k", H
		} else if 1899 <= r && r < 1913 {
			return seed + "l", H
		} else if 1913 <= r && r < 1943 {
			return seed + "n", H
		} else if 1943 <= r && r < 1962 {
			return seed + "o", H
		} else if 1962 <= r && r < 1964 {
			return seed + "p", H
		} else if 1964 <= r && r < 1992 {
			return seed + "s", H
		} else if 1992 <= r && r < 2076 {
			return seed + "t", H
		} else if 2076 <= r && r < 2083 {
			return seed + "u", H
		} else if 2083 <= r && r < 2085 {
			return seed + "v", H
		} else if 2085 <= r && r < 2087 {
			return seed + "w", H
		} else if 2087 <= r && r < 2091 {
			return seed + "x", H
		} else if 2091 <= r && r < 2106 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ga" {
		r := rng.IntN(216)
		H := 3.7068928602172515
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 10 {
			return seed + "c", H
		} else if 10 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 17 {
			return seed + "e", H
		} else if 17 <= r && r < 20 {
			return seed + "f", H
		} else if 20 <= r && r < 33 {
			return seed + "g", H
		} else if 33 <= r && r < 36 {
			return seed + "h", H
		} else if 36 <= r && r < 44 {
			return seed + "i", H
		} else if 44 <= r && r < 45 {
			return seed + "j", H
		} else if 45 <= r && r < 46 {
			return seed + "k", H
		} else if 46 <= r && r < 81 {
			return seed + "l", H
		} else if 81 <= r && r < 96 {
			return seed + "m", H
		} else if 96 <= r && r < 121 {
			return seed + "n", H
		} else if 121 <= r && r < 124 {
			return seed + "p", H
		} else if 124 <= r && r < 125 {
			return seed + "q", H
		} else if 125 <= r && r < 156 {
			return seed + "r", H
		} else if 156 <= r && r < 159 {
			return seed + "s", H
		} else if 159 <= r && r < 196 {
			return seed + "t", H
		} else if 196 <= r && r < 202 {
			return seed + "u", H
		} else if 202 <= r && r < 205 {
			return seed + "v", H
		} else if 205 <= r && r < 208 {
			return seed + "w", H
		} else if 208 <= r && r < 209 {
			return seed + "x", H
		} else if 209 <= r && r < 216 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ds" {
		r := rng.IntN(46)
		H := 3.268157478625395
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 7 {
			return seed + "c", H
		} else if 7 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 18 {
			return seed + "h", H
		} else if 18 <= r && r < 23 {
			return seed + "i", H
		} else if 23 <= r && r < 25 {
			return seed + "l", H
		} else if 25 <= r && r < 29 {
			return seed + "m", H
		} else if 29 <= r && r < 32 {
			return seed + "o", H
		} else if 32 <= r && r < 44 {
			return seed + "t", H
		} else if 44 <= r && r < 45 {
			return seed + "u", H
		} else if 45 <= r && r < 46 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ps" {
		r := rng.IntN(60)
		H := 2.64308241646309
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 5 {
			return seed + "c", H
		} else if 5 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 33 {
			return seed + "i", H
		} else if 33 <= r && r < 34 {
			return seed + "o", H
		} else if 34 <= r && r < 50 {
			return seed + "t", H
		} else if 50 <= r && r < 53 {
			return seed + "u", H
		} else if 53 <= r && r < 55 {
			return seed + "w", H
		} else if 55 <= r && r < 60 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rr" {
		r := rng.IntN(188)
		H := 2.41853444523844
		if 0 <= r && r < 33 {
			return seed + "a", H
		} else if 33 <= r && r < 68 {
			return seed + "e", H
		} else if 68 <= r && r < 123 {
			return seed + "i", H
		} else if 123 <= r && r < 154 {
			return seed + "o", H
		} else if 154 <= r && r < 161 {
			return seed + "u", H
		} else if 161 <= r && r < 188 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xs" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fu" {
		r := rng.IntN(142)
		H := 2.366297063987337
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 5 {
			return seed + "e", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 94 {
			return seed + "m", H
		} else if 94 <= r && r < 99 {
			return seed + "n", H
		} else if 99 <= r && r < 100 {
			return seed + "p", H
		} else if 100 <= r && r < 101 {
			return seed + "q", H
		} else if 101 <= r && r < 108 {
			return seed + "r", H
		} else if 108 <= r && r < 133 {
			return seed + "s", H
		} else if 133 <= r && r < 138 {
			return seed + "t", H
		} else if 138 <= r && r < 139 {
			return seed + "v", H
		} else if 139 <= r && r < 140 {
			return seed + "w", H
		} else if 140 <= r && r < 141 {
			return seed + "x", H
		} else if 141 <= r && r < 142 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "wx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "l" {
		r := rng.IntN(394)
		H := 2.0984645167702127
		if 0 <= r && r < 127 {
			return seed + "a", H
		} else if 127 <= r && r < 192 {
			return seed + "e", H
		} else if 192 <= r && r < 305 {
			return seed + "i", H
		} else if 305 <= r && r < 306 {
			return seed + "o", H
		} else if 306 <= r && r < 383 {
			return seed + "u", H
		} else if 383 <= r && r < 394 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ns" {
		r := rng.IntN(262)
		H := 3.4714061826531477
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 29 {
			return seed + "c", H
		} else if 29 <= r && r < 78 {
			return seed + "e", H
		} else if 78 <= r && r < 82 {
			return seed + "f", H
		} else if 82 <= r && r < 102 {
			return seed + "h", H
		} else if 102 <= r && r < 131 {
			return seed + "i", H
		} else if 131 <= r && r < 133 {
			return seed + "k", H
		} else if 133 <= r && r < 141 {
			return seed + "l", H
		} else if 141 <= r && r < 147 {
			return seed + "m", H
		} else if 147 <= r && r < 151 {
			return seed + "n", H
		} else if 151 <= r && r < 174 {
			return seed + "o", H
		} else if 174 <= r && r < 184 {
			return seed + "p", H
		} else if 184 <= r && r < 228 {
			return seed + "t", H
		} else if 228 <= r && r < 255 {
			return seed + "u", H
		} else if 255 <= r && r < 261 {
			return seed + "w", H
		} else if 261 <= r && r < 262 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "un" {
		r := rng.IntN(1172)
		H := 4.028381558484818
		if 0 <= r && r < 37 {
			return seed + "a", H
		} else if 37 <= r && r < 77 {
			return seed + "b", H
		} else if 77 <= r && r < 213 {
			return seed + "c", H
		} else if 213 <= r && r < 397 {
			return seed + "d", H
		} else if 397 <= r && r < 442 {
			return seed + "e", H
		} else if 442 <= r && r < 482 {
			return seed + "f", H
		} else if 482 <= r && r < 528 {
			return seed + "g", H
		} else if 528 <= r && r < 554 {
			return seed + "h", H
		} else if 554 <= r && r < 617 {
			return seed + "i", H
		} else if 617 <= r && r < 619 {
			return seed + "j", H
		} else if 619 <= r && r < 655 {
			return seed + "k", H
		} else if 655 <= r && r < 707 {
			return seed + "l", H
		} else if 707 <= r && r < 739 {
			return seed + "m", H
		} else if 739 <= r && r < 767 {
			return seed + "n", H
		} else if 767 <= r && r < 772 {
			return seed + "o", H
		} else if 772 <= r && r < 808 {
			return seed + "p", H
		} else if 808 <= r && r < 810 {
			return seed + "q", H
		} else if 810 <= r && r < 854 {
			return seed + "r", H
		} else if 854 <= r && r < 962 {
			return seed + "s", H
		} else if 962 <= r && r < 1096 {
			return seed + "t", H
		} else if 1096 <= r && r < 1103 {
			return seed + "u", H
		} else if 1103 <= r && r < 1121 {
			return seed + "v", H
		} else if 1121 <= r && r < 1169 {
			return seed + "w", H
		} else if 1169 <= r && r < 1170 {
			return seed + "y", H
		} else if 1170 <= r && r < 1172 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ie" {
		r := rng.IntN(292)
		H := 3.2957006777417726
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 10 {
			return seed + "c", H
		} else if 10 <= r && r < 85 {
			return seed + "d", H
		} else if 85 <= r && r < 96 {
			return seed + "f", H
		} else if 96 <= r && r < 97 {
			return seed + "g", H
		} else if 97 <= r && r < 98 {
			return seed + "h", H
		} else if 98 <= r && r < 99 {
			return seed + "j", H
		} else if 99 <= r && r < 102 {
			return seed + "k", H
		} else if 102 <= r && r < 121 {
			return seed + "l", H
		} else if 121 <= r && r < 124 {
			return seed + "m", H
		} else if 124 <= r && r < 163 {
			return seed + "n", H
		} else if 163 <= r && r < 164 {
			return seed + "p", H
		} else if 164 <= r && r < 165 {
			return seed + "q", H
		} else if 165 <= r && r < 220 {
			return seed + "r", H
		} else if 220 <= r && r < 239 {
			return seed + "s", H
		} else if 239 <= r && r < 252 {
			return seed + "t", H
		} else if 252 <= r && r < 254 {
			return seed + "u", H
		} else if 254 <= r && r < 273 {
			return seed + "v", H
		} else if 273 <= r && r < 290 {
			return seed + "w", H
		} else if 290 <= r && r < 291 {
			return seed + "x", H
		} else if 291 <= r && r < 292 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dn" {
		r := rng.IntN(30)
		H := 1.2803301792260264
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 24 {
			return seed + "e", H
		} else if 24 <= r && r < 25 {
			return seed + "i", H
		} else if 25 <= r && r < 28 {
			return seed + "o", H
		} else if 28 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 30 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "iv" {
		r := rng.IntN(328)
		H := 1.652086893609666
		if 0 <= r && r < 45 {
			return seed + "a", H
		} else if 45 <= r && r < 228 {
			return seed + "e", H
		} else if 228 <= r && r < 309 {
			return seed + "i", H
		} else if 309 <= r && r < 324 {
			return seed + "o", H
		} else if 324 <= r && r < 325 {
			return seed + "u", H
		} else if 325 <= r && r < 328 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yr" {
		r := rng.IntN(34)
		H := 2.1660870601063933
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 17 {
			return seed + "i", H
		} else if 17 <= r && r < 28 {
			return seed + "o", H
		} else if 28 <= r && r < 33 {
			return seed + "u", H
		} else if 33 <= r && r < 34 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lh" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ae" {
		r := rng.IntN(32)
		H := 3.6622301466508205
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "m", H
		} else if 12 <= r && r < 13 {
			return seed + "n", H
		} else if 13 <= r && r < 14 {
			return seed + "p", H
		} else if 14 <= r && r < 15 {
			return seed + "q", H
		} else if 15 <= r && r < 26 {
			return seed + "r", H
		} else if 26 <= r && r < 27 {
			return seed + "s", H
		} else if 27 <= r && r < 28 {
			return seed + "t", H
		} else if 28 <= r && r < 29 {
			return seed + "v", H
		} else if 29 <= r && r < 30 {
			return seed + "w", H
		} else if 30 <= r && r < 31 {
			return seed + "x", H
		} else if 31 <= r && r < 32 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ne" {
		r := rng.IntN(672)
		H := 3.190526089028604
		if 0 <= r && r < 30 {
			return seed + "a", H
		} else if 30 <= r && r < 37 {
			return seed + "b", H
		} else if 37 <= r && r < 44 {
			return seed + "c", H
		} else if 44 <= r && r < 123 {
			return seed + "d", H
		} else if 123 <= r && r < 137 {
			return seed + "e", H
		} else if 137 <= r && r < 138 {
			return seed + "f", H
		} else if 138 <= r && r < 155 {
			return seed + "g", H
		} else if 155 <= r && r < 158 {
			return seed + "h", H
		} else if 158 <= r && r < 159 {
			return seed + "j", H
		} else if 159 <= r && r < 160 {
			return seed + "k", H
		} else if 160 <= r && r < 181 {
			return seed + "l", H
		} else if 181 <= r && r < 192 {
			return seed + "m", H
		} else if 192 <= r && r < 213 {
			return seed + "n", H
		} else if 213 <= r && r < 215 {
			return seed + "o", H
		} else if 215 <= r && r < 218 {
			return seed + "p", H
		} else if 218 <= r && r < 221 {
			return seed + "q", H
		} else if 221 <= r && r < 312 {
			return seed + "r", H
		} else if 312 <= r && r < 579 {
			return seed + "s", H
		} else if 579 <= r && r < 610 {
			return seed + "t", H
		} else if 610 <= r && r < 624 {
			return seed + "u", H
		} else if 624 <= r && r < 631 {
			return seed + "v", H
		} else if 631 <= r && r < 648 {
			return seed + "w", H
		} else if 648 <= r && r < 659 {
			return seed + "x", H
		} else if 659 <= r && r < 671 {
			return seed + "y", H
		} else if 671 <= r && r < 672 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gp" {
		r := rng.IntN(12)
		H := 2.625814583693911
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 7 {
			return seed + "l", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ma" {
		r := rng.IntN(554)
		H := 3.348309073249998
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 28 {
			return seed + "c", H
		} else if 28 <= r && r < 37 {
			return seed + "d", H
		} else if 37 <= r && r < 38 {
			return seed + "f", H
		} else if 38 <= r && r < 81 {
			return seed + "g", H
		} else if 81 <= r && r < 84 {
			return seed + "h", H
		} else if 84 <= r && r < 92 {
			return seed + "i", H
		} else if 92 <= r && r < 101 {
			return seed + "j", H
		} else if 101 <= r && r < 122 {
			return seed + "k", H
		} else if 122 <= r && r < 147 {
			return seed + "l", H
		} else if 147 <= r && r < 156 {
			return seed + "m", H
		} else if 156 <= r && r < 307 {
			return seed + "n", H
		} else if 307 <= r && r < 312 {
			return seed + "p", H
		} else if 312 <= r && r < 313 {
			return seed + "q", H
		} else if 313 <= r && r < 398 {
			return seed + "r", H
		} else if 398 <= r && r < 431 {
			return seed + "s", H
		} else if 431 <= r && r < 530 {
			return seed + "t", H
		} else if 530 <= r && r < 532 {
			return seed + "u", H
		} else if 532 <= r && r < 535 {
			return seed + "v", H
		} else if 535 <= r && r < 536 {
			return seed + "w", H
		} else if 536 <= r && r < 541 {
			return seed + "x", H
		} else if 541 <= r && r < 549 {
			return seed + "y", H
		} else if 549 <= r && r < 554 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "qv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ve" {
		r := rng.IntN(574)
		H := 2.235089534302184
		if 0 <= r && r < 6 {
			return seed + "a", H
		} else if 6 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 43 {
			return seed + "d", H
		} else if 43 <= r && r < 45 {
			return seed + "e", H
		} else if 45 <= r && r < 46 {
			return seed + "f", H
		} else if 46 <= r && r < 51 {
			return seed + "g", H
		} else if 51 <= r && r < 54 {
			return seed + "h", H
		} else if 54 <= r && r < 58 {
			return seed + "i", H
		} else if 58 <= r && r < 59 {
			return seed + "j", H
		} else if 59 <= r && r < 60 {
			return seed + "k", H
		} else if 60 <= r && r < 93 {
			return seed + "l", H
		} else if 93 <= r && r < 96 {
			return seed + "m", H
		} else if 96 <= r && r < 183 {
			return seed + "n", H
		} else if 183 <= r && r < 184 {
			return seed + "p", H
		} else if 184 <= r && r < 185 {
			return seed + "q", H
		} else if 185 <= r && r < 524 {
			return seed + "r", H
		} else if 524 <= r && r < 543 {
			return seed + "s", H
		} else if 543 <= r && r < 560 {
			return seed + "t", H
		} else if 560 <= r && r < 561 {
			return seed + "v", H
		} else if 561 <= r && r < 564 {
			return seed + "w", H
		} else if 564 <= r && r < 567 {
			return seed + "x", H
		} else if 567 <= r && r < 573 {
			return seed + "y", H
		} else if 573 <= r && r < 574 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zl" {
		r := rng.IntN(32)
		H := 1.7150939152501616
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 27 {
			return seed + "i", H
		} else if 27 <= r && r < 28 {
			return seed + "o", H
		} else if 28 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 32 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jl" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nz" {
		r := rng.IntN(20)
		H := 2.2833830982290135
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ck" {
		r := rng.IntN(276)
		H := 3.4140722500944016
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 17 {
			return seed + "b", H
		} else if 17 <= r && r < 19 {
			return seed + "d", H
		} else if 19 <= r && r < 80 {
			return seed + "e", H
		} else if 80 <= r && r < 86 {
			return seed + "f", H
		} else if 86 <= r && r < 88 {
			return seed + "h", H
		} else if 88 <= r && r < 131 {
			return seed + "i", H
		} else if 131 <= r && r < 133 {
			return seed + "k", H
		} else if 133 <= r && r < 175 {
			return seed + "l", H
		} else if 175 <= r && r < 183 {
			return seed + "n", H
		} else if 183 <= r && r < 186 {
			return seed + "o", H
		} else if 186 <= r && r < 192 {
			return seed + "p", H
		} else if 192 <= r && r < 196 {
			return seed + "r", H
		} else if 196 <= r && r < 234 {
			return seed + "s", H
		} else if 234 <= r && r < 242 {
			return seed + "t", H
		} else if 242 <= r && r < 249 {
			return seed + "u", H
		} else if 249 <= r && r < 257 {
			return seed + "w", H
		} else if 257 <= r && r < 276 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uf" {
		r := rng.IntN(72)
		H := 0.6292303920303968
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 68 {
			return seed + "f", H
		} else if 68 <= r && r < 69 {
			return seed + "i", H
		} else if 69 <= r && r < 70 {
			return seed + "o", H
		} else if 70 <= r && r < 71 {
			return seed + "u", H
		} else if 71 <= r && r < 72 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ap" {
		r := rng.IntN(292)
		H := 2.9142100268241307
		if 0 <= r && r < 17 {
			return seed + "a", H
		} else if 17 <= r && r < 21 {
			return seed + "d", H
		} else if 21 <= r && r < 66 {
			return seed + "e", H
		} else if 66 <= r && r < 82 {
			return seed + "h", H
		} else if 82 <= r && r < 95 {
			return seed + "i", H
		} else if 95 <= r && r < 97 {
			return seed + "k", H
		} else if 97 <= r && r < 105 {
			return seed + "l", H
		} else if 105 <= r && r < 107 {
			return seed + "n", H
		} else if 107 <= r && r < 114 {
			return seed + "o", H
		} else if 114 <= r && r < 234 {
			return seed + "p", H
		} else if 234 <= r && r < 244 {
			return seed + "r", H
		} else if 244 <= r && r < 266 {
			return seed + "s", H
		} else if 266 <= r && r < 286 {
			return seed + "t", H
		} else if 286 <= r && r < 289 {
			return seed + "u", H
		} else if 289 <= r && r < 292 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gi" {
		r := rng.IntN(224)
		H := 3.309053563139037
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 28 {
			return seed + "c", H
		} else if 28 <= r && r < 41 {
			return seed + "d", H
		} else if 41 <= r && r < 47 {
			return seed + "e", H
		} else if 47 <= r && r < 50 {
			return seed + "f", H
		} else if 50 <= r && r < 65 {
			return seed + "g", H
		} else if 65 <= r && r < 66 {
			return seed + "h", H
		} else if 66 <= r && r < 67 {
			return seed + "j", H
		} else if 67 <= r && r < 68 {
			return seed + "k", H
		} else if 68 <= r && r < 85 {
			return seed + "l", H
		} else if 85 <= r && r < 90 {
			return seed + "m", H
		} else if 90 <= r && r < 177 {
			return seed + "n", H
		} else if 177 <= r && r < 179 {
			return seed + "o", H
		} else if 179 <= r && r < 180 {
			return seed + "p", H
		} else if 180 <= r && r < 181 {
			return seed + "q", H
		} else if 181 <= r && r < 186 {
			return seed + "r", H
		} else if 186 <= r && r < 203 {
			return seed + "s", H
		} else if 203 <= r && r < 204 {
			return seed + "t", H
		} else if 204 <= r && r < 215 {
			return seed + "v", H
		} else if 215 <= r && r < 216 {
			return seed + "w", H
		} else if 216 <= r && r < 217 {
			return seed + "x", H
		} else if 217 <= r && r < 224 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qs" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ki" {
		r := rng.IntN(278)
		H := 2.47511759884632
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 11 {
			return seed + "d", H
		} else if 11 <= r && r < 21 {
			return seed + "e", H
		} else if 21 <= r && r < 22 {
			return seed + "f", H
		} else if 22 <= r && r < 23 {
			return seed + "g", H
		} else if 23 <= r && r < 24 {
			return seed + "h", H
		} else if 24 <= r && r < 26 {
			return seed + "i", H
		} else if 26 <= r && r < 27 {
			return seed + "j", H
		} else if 27 <= r && r < 28 {
			return seed + "k", H
		} else if 28 <= r && r < 63 {
			return seed + "l", H
		} else if 63 <= r && r < 76 {
			return seed + "m", H
		} else if 76 <= r && r < 235 {
			return seed + "n", H
		} else if 235 <= r && r < 242 {
			return seed + "p", H
		} else if 242 <= r && r < 243 {
			return seed + "q", H
		} else if 243 <= r && r < 250 {
			return seed + "r", H
		} else if 250 <= r && r < 261 {
			return seed + "s", H
		} else if 261 <= r && r < 272 {
			return seed + "t", H
		} else if 272 <= r && r < 273 {
			return seed + "v", H
		} else if 273 <= r && r < 276 {
			return seed + "w", H
		} else if 276 <= r && r < 277 {
			return seed + "x", H
		} else if 277 <= r && r < 278 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "wn" {
		r := rng.IntN(14)
		H := 2.645593314451147
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 9 {
			return seed + "n", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mm" {
		r := rng.IntN(144)
		H := 2.4746985376948505
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 58 {
			return seed + "e", H
		} else if 58 <= r && r < 81 {
			return seed + "i", H
		} else if 81 <= r && r < 114 {
			return seed + "o", H
		} else if 114 <= r && r < 125 {
			return seed + "u", H
		} else if 125 <= r && r < 144 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wf" {
		r := rng.IntN(16)
		H := 2.5550365325772657
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ye" {
		r := rng.IntN(68)
		H := 3.64442751173858
		if 0 <= r && r < 12 {
			return seed + "a", H
		} else if 12 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 14 {
			return seed + "c", H
		} else if 14 <= r && r < 23 {
			return seed + "d", H
		} else if 23 <= r && r < 27 {
			return seed + "e", H
		} else if 27 <= r && r < 28 {
			return seed + "f", H
		} else if 28 <= r && r < 29 {
			return seed + "g", H
		} else if 29 <= r && r < 30 {
			return seed + "h", H
		} else if 30 <= r && r < 31 {
			return seed + "j", H
		} else if 31 <= r && r < 32 {
			return seed + "k", H
		} else if 32 <= r && r < 37 {
			return seed + "l", H
		} else if 37 <= r && r < 38 {
			return seed + "m", H
		} else if 38 <= r && r < 41 {
			return seed + "n", H
		} else if 41 <= r && r < 42 {
			return seed + "p", H
		} else if 42 <= r && r < 43 {
			return seed + "q", H
		} else if 43 <= r && r < 58 {
			return seed + "r", H
		} else if 58 <= r && r < 63 {
			return seed + "s", H
		} else if 63 <= r && r < 64 {
			return seed + "t", H
		} else if 64 <= r && r < 65 {
			return seed + "v", H
		} else if 65 <= r && r < 66 {
			return seed + "w", H
		} else if 66 <= r && r < 67 {
			return seed + "x", H
		} else if 67 <= r && r < 68 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xu" {
		r := rng.IntN(28)
		H := 4.151478922893308
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 7 {
			return seed + "d", H
		} else if 7 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 9 {
			return seed + "g", H
		} else if 9 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 11 {
			return seed + "j", H
		} else if 11 <= r && r < 12 {
			return seed + "k", H
		} else if 12 <= r && r < 13 {
			return seed + "l", H
		} else if 13 <= r && r < 14 {
			return seed + "m", H
		} else if 14 <= r && r < 15 {
			return seed + "n", H
		} else if 15 <= r && r < 16 {
			return seed + "p", H
		} else if 16 <= r && r < 17 {
			return seed + "q", H
		} else if 17 <= r && r < 22 {
			return seed + "r", H
		} else if 22 <= r && r < 23 {
			return seed + "s", H
		} else if 23 <= r && r < 24 {
			return seed + "t", H
		} else if 24 <= r && r < 25 {
			return seed + "v", H
		} else if 25 <= r && r < 26 {
			return seed + "w", H
		} else if 26 <= r && r < 27 {
			return seed + "x", H
		} else if 27 <= r && r < 28 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "er" {
		r := rng.IntN(938)
		H := 4.231329828166286
		if 0 <= r && r < 77 {
			return seed + "a", H
		} else if 77 <= r && r < 121 {
			return seed + "b", H
		} else if 121 <= r && r < 155 {
			return seed + "c", H
		} else if 155 <= r && r < 175 {
			return seed + "d", H
		} else if 175 <= r && r < 252 {
			return seed + "e", H
		} else if 252 <= r && r < 276 {
			return seed + "f", H
		} else if 276 <= r && r < 312 {
			return seed + "g", H
		} else if 312 <= r && r < 330 {
			return seed + "h", H
		} else if 330 <= r && r < 427 {
			return seed + "i", H
		} else if 427 <= r && r < 435 {
			return seed + "j", H
		} else if 435 <= r && r < 443 {
			return seed + "k", H
		} else if 443 <= r && r < 475 {
			return seed + "l", H
		} else if 475 <= r && r < 511 {
			return seed + "m", H
		} else if 511 <= r && r < 547 {
			return seed + "n", H
		} else if 547 <= r && r < 594 {
			return seed + "o", H
		} else if 594 <= r && r < 620 {
			return seed + "p", H
		} else if 620 <= r && r < 666 {
			return seed + "r", H
		} else if 666 <= r && r < 762 {
			return seed + "s", H
		} else if 762 <= r && r < 828 {
			return seed + "t", H
		} else if 828 <= r && r < 839 {
			return seed + "u", H
		} else if 839 <= r && r < 873 {
			return seed + "v", H
		} else if 873 <= r && r < 881 {
			return seed + "w", H
		} else if 881 <= r && r < 938 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pd" {
		r := rng.IntN(14)
		H := 2.064042639445697
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hh" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ui" {
		r := rng.IntN(166)
		H := 3.499233981643169
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 18 {
			return seed + "c", H
		} else if 18 <= r && r < 33 {
			return seed + "d", H
		} else if 33 <= r && r < 35 {
			return seed + "e", H
		} else if 35 <= r && r < 36 {
			return seed + "f", H
		} else if 36 <= r && r < 41 {
			return seed + "g", H
		} else if 41 <= r && r < 42 {
			return seed + "h", H
		} else if 42 <= r && r < 43 {
			return seed + "j", H
		} else if 43 <= r && r < 44 {
			return seed + "k", H
		} else if 44 <= r && r < 59 {
			return seed + "l", H
		} else if 59 <= r && r < 60 {
			return seed + "m", H
		} else if 60 <= r && r < 79 {
			return seed + "n", H
		} else if 79 <= r && r < 82 {
			return seed + "p", H
		} else if 82 <= r && r < 83 {
			return seed + "q", H
		} else if 83 <= r && r < 100 {
			return seed + "r", H
		} else if 100 <= r && r < 121 {
			return seed + "s", H
		} else if 121 <= r && r < 156 {
			return seed + "t", H
		} else if 156 <= r && r < 161 {
			return seed + "v", H
		} else if 161 <= r && r < 162 {
			return seed + "w", H
		} else if 162 <= r && r < 163 {
			return seed + "x", H
		} else if 163 <= r && r < 166 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "om" {
		r := rng.IntN(250)
		H := 2.9507116839110523
		if 0 <= r && r < 39 {
			return seed + "a", H
		} else if 39 <= r && r < 51 {
			return seed + "b", H
		} else if 51 <= r && r < 96 {
			return seed + "e", H
		} else if 96 <= r && r < 100 {
			return seed + "f", H
		} else if 100 <= r && r < 143 {
			return seed + "i", H
		} else if 143 <= r && r < 169 {
			return seed + "m", H
		} else if 169 <= r && r < 171 {
			return seed + "n", H
		} else if 171 <= r && r < 178 {
			return seed + "o", H
		} else if 178 <= r && r < 236 {
			return seed + "p", H
		} else if 236 <= r && r < 238 {
			return seed + "r", H
		} else if 238 <= r && r < 240 {
			return seed + "s", H
		} else if 240 <= r && r < 241 {
			return seed + "u", H
		} else if 241 <= r && r < 250 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hl" {
		r := rng.IntN(40)
		H := 2.1930689652207995
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 23 {
			return seed + "u", H
		} else if 23 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sm" {
		r := rng.IntN(112)
		H := 2.1869546638292445
		if 0 <= r && r < 37 {
			return seed + "a", H
		} else if 37 <= r && r < 44 {
			return seed + "e", H
		} else if 44 <= r && r < 69 {
			return seed + "i", H
		} else if 69 <= r && r < 98 {
			return seed + "o", H
		} else if 98 <= r && r < 111 {
			return seed + "u", H
		} else if 111 <= r && r < 112 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "v" {
		r := rng.IntN(270)
		H := 1.9781579547549006
		if 0 <= r && r < 69 {
			return seed + "a", H
		} else if 69 <= r && r < 138 {
			return seed + "e", H
		} else if 138 <= r && r < 233 {
			return seed + "i", H
		} else if 233 <= r && r < 268 {
			return seed + "o", H
		} else if 268 <= r && r < 269 {
			return seed + "u", H
		} else if 269 <= r && r < 270 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fi" {
		r := rng.IntN(302)
		H := 3.500452407026575
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 22 {
			return seed + "c", H
		} else if 22 <= r && r < 45 {
			return seed + "d", H
		} else if 45 <= r && r < 91 {
			return seed + "e", H
		} else if 91 <= r && r < 100 {
			return seed + "f", H
		} else if 100 <= r && r < 115 {
			return seed + "g", H
		} else if 115 <= r && r < 116 {
			return seed + "h", H
		} else if 116 <= r && r < 117 {
			return seed + "j", H
		} else if 117 <= r && r < 118 {
			return seed + "k", H
		} else if 118 <= r && r < 155 {
			return seed + "l", H
		} else if 155 <= r && r < 156 {
			return seed + "m", H
		} else if 156 <= r && r < 225 {
			return seed + "n", H
		} else if 225 <= r && r < 226 {
			return seed + "p", H
		} else if 226 <= r && r < 227 {
			return seed + "q", H
		} else if 227 <= r && r < 246 {
			return seed + "r", H
		} else if 246 <= r && r < 273 {
			return seed + "s", H
		} else if 273 <= r && r < 286 {
			return seed + "t", H
		} else if 286 <= r && r < 289 {
			return seed + "v", H
		} else if 289 <= r && r < 290 {
			return seed + "w", H
		} else if 290 <= r && r < 301 {
			return seed + "x", H
		} else if 301 <= r && r < 302 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ih" {
		r := rng.IntN(10)
		H := 2.3709505944546683
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ea" {
		r := rng.IntN(612)
		H := 3.376745698000607
		if 0 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 48 {
			return seed + "c", H
		} else if 48 <= r && r < 151 {
			return seed + "d", H
		} else if 151 <= r && r < 158 {
			return seed + "f", H
		} else if 158 <= r && r < 163 {
			return seed + "g", H
		} else if 163 <= r && r < 166 {
			return seed + "h", H
		} else if 166 <= r && r < 167 {
			return seed + "j", H
		} else if 167 <= r && r < 186 {
			return seed + "k", H
		} else if 186 <= r && r < 233 {
			return seed + "l", H
		} else if 233 <= r && r < 272 {
			return seed + "m", H
		} else if 272 <= r && r < 287 {
			return seed + "n", H
		} else if 287 <= r && r < 298 {
			return seed + "p", H
		} else if 298 <= r && r < 299 {
			return seed + "q", H
		} else if 299 <= r && r < 418 {
			return seed + "r", H
		} else if 418 <= r && r < 491 {
			return seed + "s", H
		} else if 491 <= r && r < 590 {
			return seed + "t", H
		} else if 590 <= r && r < 605 {
			return seed + "v", H
		} else if 605 <= r && r < 610 {
			return seed + "w", H
		} else if 610 <= r && r < 611 {
			return seed + "x", H
		} else if 611 <= r && r < 612 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xt" {
		r := rng.IntN(46)
		H := 2.3694062351215366
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 22 {
			return seed + "h", H
		} else if 22 <= r && r < 29 {
			return seed + "i", H
		} else if 29 <= r && r < 32 {
			return seed + "o", H
		} else if 32 <= r && r < 42 {
			return seed + "r", H
		} else if 42 <= r && r < 43 {
			return seed + "u", H
		} else if 43 <= r && r < 46 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ev" {
		r := rng.IntN(218)
		H := 2.0506013620233894
		if 0 <= r && r < 35 {
			return seed + "a", H
		} else if 35 <= r && r < 120 {
			return seed + "e", H
		} else if 120 <= r && r < 181 {
			return seed + "i", H
		} else if 181 <= r && r < 212 {
			return seed + "o", H
		} else if 212 <= r && r < 214 {
			return seed + "r", H
		} else if 214 <= r && r < 215 {
			return seed + "u", H
		} else if 215 <= r && r < 218 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nq" {
		r := rng.IntN(14)
		H := 1.7695459925589747
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vb" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tp" {
		r := rng.IntN(20)
		H := 2.7232196723355075
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 16 {
			return seed + "r", H
		} else if 16 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ej" {
		r := rng.IntN(16)
		H := 2.0461796919474975
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mo" {
		r := rng.IntN(426)
		H := 3.753088087266153
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 30 {
			return seed + "c", H
		} else if 30 <= r && r < 51 {
			return seed + "d", H
		} else if 51 <= r && r < 52 {
			return seed + "f", H
		} else if 52 <= r && r < 57 {
			return seed + "g", H
		} else if 57 <= r && r < 58 {
			return seed + "h", H
		} else if 58 <= r && r < 68 {
			return seed + "i", H
		} else if 68 <= r && r < 69 {
			return seed + "j", H
		} else if 69 <= r && r < 80 {
			return seed + "k", H
		} else if 80 <= r && r < 103 {
			return seed + "l", H
		} else if 103 <= r && r < 110 {
			return seed + "m", H
		} else if 110 <= r && r < 193 {
			return seed + "n", H
		} else if 193 <= r && r < 225 {
			return seed + "o", H
		} else if 225 <= r && r < 228 {
			return seed + "p", H
		} else if 228 <= r && r < 229 {
			return seed + "q", H
		} else if 229 <= r && r < 292 {
			return seed + "r", H
		} else if 292 <= r && r < 317 {
			return seed + "s", H
		} else if 317 <= r && r < 358 {
			return seed + "t", H
		} else if 358 <= r && r < 390 {
			return seed + "u", H
		} else if 390 <= r && r < 419 {
			return seed + "v", H
		} else if 419 <= r && r < 424 {
			return seed + "w", H
		} else if 424 <= r && r < 425 {
			return seed + "x", H
		} else if 425 <= r && r < 426 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "uz" {
		r := rng.IntN(20)
		H := 1.9332062193464954
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		} else if 8 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "fv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qr" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cq" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ay" {
		r := rng.IntN(126)
		H := 4.198083408872458
		if 0 <= r && r < 6 {
			return seed + "a", H
		} else if 6 <= r && r < 19 {
			return seed + "b", H
		} else if 19 <= r && r < 24 {
			return seed + "c", H
		} else if 24 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 43 {
			return seed + "e", H
		} else if 43 <= r && r < 50 {
			return seed + "f", H
		} else if 50 <= r && r < 53 {
			return seed + "g", H
		} else if 53 <= r && r < 56 {
			return seed + "h", H
		} else if 56 <= r && r < 66 {
			return seed + "i", H
		} else if 66 <= r && r < 67 {
			return seed + "j", H
		} else if 67 <= r && r < 68 {
			return seed + "k", H
		} else if 68 <= r && r < 75 {
			return seed + "l", H
		} else if 75 <= r && r < 84 {
			return seed + "m", H
		} else if 84 <= r && r < 85 {
			return seed + "n", H
		} else if 85 <= r && r < 89 {
			return seed + "o", H
		} else if 89 <= r && r < 94 {
			return seed + "p", H
		} else if 94 <= r && r < 95 {
			return seed + "q", H
		} else if 95 <= r && r < 102 {
			return seed + "r", H
		} else if 102 <= r && r < 113 {
			return seed + "s", H
		} else if 113 <= r && r < 120 {
			return seed + "t", H
		} else if 120 <= r && r < 121 {
			return seed + "v", H
		} else if 121 <= r && r < 124 {
			return seed + "w", H
		} else if 124 <= r && r < 125 {
			return seed + "x", H
		} else if 125 <= r && r < 126 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "my" {
		r := rng.IntN(32)
		H := 3.6622301466508205
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 25 {
			return seed + "s", H
		} else if 25 <= r && r < 28 {
			return seed + "t", H
		} else if 28 <= r && r < 29 {
			return seed + "v", H
		} else if 29 <= r && r < 30 {
			return seed + "w", H
		} else if 30 <= r && r < 31 {
			return seed + "x", H
		} else if 31 <= r && r < 32 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pk" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rn" {
		r := rng.IntN(104)
		H := 3.0248321500013837
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 23 {
			return seed + "b", H
		} else if 23 <= r && r < 25 {
			return seed + "c", H
		} else if 25 <= r && r < 54 {
			return seed + "e", H
		} else if 54 <= r && r < 60 {
			return seed + "f", H
		} else if 60 <= r && r < 62 {
			return seed + "h", H
		} else if 62 <= r && r < 85 {
			return seed + "i", H
		} else if 85 <= r && r < 87 {
			return seed + "l", H
		} else if 87 <= r && r < 89 {
			return seed + "m", H
		} else if 89 <= r && r < 91 {
			return seed + "n", H
		} else if 91 <= r && r < 96 {
			return seed + "o", H
		} else if 96 <= r && r < 98 {
			return seed + "s", H
		} else if 98 <= r && r < 101 {
			return seed + "u", H
		} else if 101 <= r && r < 104 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jb" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gg" {
		r := rng.IntN(124)
		H := 2.6293806113478104
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 38 {
			return seed + "e", H
		} else if 38 <= r && r < 65 {
			return seed + "i", H
		} else if 65 <= r && r < 103 {
			return seed + "l", H
		} else if 103 <= r && r < 105 {
			return seed + "n", H
		} else if 105 <= r && r < 108 {
			return seed + "o", H
		} else if 108 <= r && r < 110 {
			return seed + "p", H
		} else if 110 <= r && r < 112 {
			return seed + "r", H
		} else if 112 <= r && r < 114 {
			return seed + "s", H
		} else if 114 <= r && r < 115 {
			return seed + "u", H
		} else if 115 <= r && r < 124 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xm" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bs" {
		r := rng.IntN(76)
		H := 2.7319949617023767
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 9 {
			return seed + "c", H
		} else if 9 <= r && r < 30 {
			return seed + "e", H
		} else if 30 <= r && r < 41 {
			return seed + "i", H
		} else if 41 <= r && r < 43 {
			return seed + "l", H
		} else if 43 <= r && r < 56 {
			return seed + "o", H
		} else if 56 <= r && r < 70 {
			return seed + "t", H
		} else if 70 <= r && r < 73 {
			return seed + "u", H
		} else if 73 <= r && r < 76 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "be" {
		r := rng.IntN(158)
		H := 3.156325099409102
		if 0 <= r && r < 12 {
			return seed + "a", H
		} else if 12 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 18 {
			return seed + "c", H
		} else if 18 <= r && r < 51 {
			return seed + "d", H
		} else if 51 <= r && r < 55 {
			return seed + "e", H
		} else if 55 <= r && r < 56 {
			return seed + "f", H
		} else if 56 <= r && r < 57 {
			return seed + "g", H
		} else if 57 <= r && r < 58 {
			return seed + "h", H
		} else if 58 <= r && r < 59 {
			return seed + "j", H
		} else if 59 <= r && r < 60 {
			return seed + "k", H
		} else if 60 <= r && r < 75 {
			return seed + "l", H
		} else if 75 <= r && r < 76 {
			return seed + "m", H
		} else if 76 <= r && r < 81 {
			return seed + "n", H
		} else if 81 <= r && r < 82 {
			return seed + "p", H
		} else if 82 <= r && r < 83 {
			return seed + "q", H
		} else if 83 <= r && r < 138 {
			return seed + "r", H
		} else if 138 <= r && r < 143 {
			return seed + "s", H
		} else if 143 <= r && r < 150 {
			return seed + "t", H
		} else if 150 <= r && r < 151 {
			return seed + "v", H
		} else if 151 <= r && r < 152 {
			return seed + "w", H
		} else if 152 <= r && r < 153 {
			return seed + "x", H
		} else if 153 <= r && r < 155 {
			return seed + "y", H
		} else if 155 <= r && r < 158 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ir" {
		r := rng.IntN(226)
		H := 3.234912888337366
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 23 {
			return seed + "c", H
		} else if 23 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 110 {
			return seed + "e", H
		} else if 110 <= r && r < 137 {
			return seed + "i", H
		} else if 137 <= r && r < 143 {
			return seed + "k", H
		} else if 143 <= r && r < 149 {
			return seed + "l", H
		} else if 149 <= r && r < 159 {
			return seed + "m", H
		} else if 159 <= r && r < 162 {
			return seed + "o", H
		} else if 162 <= r && r < 166 {
			return seed + "p", H
		} else if 166 <= r && r < 180 {
			return seed + "r", H
		} else if 180 <= r && r < 190 {
			return seed + "s", H
		} else if 190 <= r && r < 214 {
			return seed + "t", H
		} else if 214 <= r && r < 221 {
			return seed + "u", H
		} else if 221 <= r && r < 226 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bu" {
		r := rng.IntN(198)
		H := 3.5606989772697664
		if 0 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 26 {
			return seed + "c", H
		} else if 26 <= r && r < 37 {
			return seed + "d", H
		} else if 37 <= r && r < 50 {
			return seed + "f", H
		} else if 50 <= r && r < 57 {
			return seed + "g", H
		} else if 57 <= r && r < 58 {
			return seed + "h", H
		} else if 58 <= r && r < 64 {
			return seed + "i", H
		} else if 64 <= r && r < 65 {
			return seed + "j", H
		} else if 65 <= r && r < 66 {
			return seed + "k", H
		} else if 66 <= r && r < 115 {
			return seed + "l", H
		} else if 115 <= r && r < 118 {
			return seed + "m", H
		} else if 118 <= r && r < 145 {
			return seed + "n", H
		} else if 145 <= r && r < 148 {
			return seed + "p", H
		} else if 148 <= r && r < 149 {
			return seed + "q", H
		} else if 149 <= r && r < 164 {
			return seed + "r", H
		} else if 164 <= r && r < 181 {
			return seed + "s", H
		} else if 181 <= r && r < 192 {
			return seed + "t", H
		} else if 192 <= r && r < 193 {
			return seed + "v", H
		} else if 193 <= r && r < 194 {
			return seed + "w", H
		} else if 194 <= r && r < 195 {
			return seed + "x", H
		} else if 195 <= r && r < 198 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "le" {
		r := rng.IntN(656)
		H := 3.6842108008877057
		if 0 <= r && r < 36 {
			return seed + "a", H
		} else if 36 <= r && r < 41 {
			return seed + "b", H
		} else if 41 <= r && r < 84 {
			return seed + "c", H
		} else if 84 <= r && r < 177 {
			return seed + "d", H
		} else if 177 <= r && r < 197 {
			return seed + "e", H
		} else if 197 <= r && r < 204 {
			return seed + "f", H
		} else if 204 <= r && r < 239 {
			return seed + "g", H
		} else if 239 <= r && r < 242 {
			return seed + "h", H
		} else if 242 <= r && r < 244 {
			return seed + "i", H
		} else if 244 <= r && r < 245 {
			return seed + "j", H
		} else if 245 <= r && r < 246 {
			return seed + "k", H
		} else if 246 <= r && r < 249 {
			return seed + "l", H
		} else if 249 <= r && r < 264 {
			return seed + "m", H
		} else if 264 <= r && r < 309 {
			return seed + "n", H
		} else if 309 <= r && r < 311 {
			return seed + "o", H
		} else if 311 <= r && r < 320 {
			return seed + "p", H
		} else if 320 <= r && r < 321 {
			return seed + "q", H
		} else if 321 <= r && r < 398 {
			return seed + "r", H
		} else if 398 <= r && r < 537 {
			return seed + "s", H
		} else if 537 <= r && r < 596 {
			return seed + "t", H
		} else if 596 <= r && r < 602 {
			return seed + "u", H
		} else if 602 <= r && r < 629 {
			return seed + "v", H
		} else if 629 <= r && r < 632 {
			return seed + "w", H
		} else if 632 <= r && r < 645 {
			return seed + "x", H
		} else if 645 <= r && r < 655 {
			return seed + "y", H
		} else if 655 <= r && r < 656 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "kh" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cf" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sk" {
		r := rng.IntN(128)
		H := 2.129438550927799
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 36 {
			return seed + "e", H
		} else if 36 <= r && r < 101 {
			return seed + "i", H
		} else if 101 <= r && r < 103 {
			return seed + "l", H
		} else if 103 <= r && r < 105 {
			return seed + "n", H
		} else if 105 <= r && r < 106 {
			return seed + "o", H
		} else if 106 <= r && r < 108 {
			return seed + "t", H
		} else if 108 <= r && r < 109 {
			return seed + "u", H
		} else if 109 <= r && r < 111 {
			return seed + "w", H
		} else if 111 <= r && r < 128 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zo" {
		r := rng.IntN(46)
		H := 3.8875101959226344
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 7 {
			return seed + "d", H
		} else if 7 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 9 {
			return seed + "g", H
		} else if 9 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 12 {
			return seed + "i", H
		} else if 12 <= r && r < 13 {
			return seed + "j", H
		} else if 13 <= r && r < 14 {
			return seed + "k", H
		} else if 14 <= r && r < 15 {
			return seed + "l", H
		} else if 15 <= r && r < 18 {
			return seed + "m", H
		} else if 18 <= r && r < 27 {
			return seed + "n", H
		} else if 27 <= r && r < 37 {
			return seed + "o", H
		} else if 37 <= r && r < 38 {
			return seed + "p", H
		} else if 38 <= r && r < 39 {
			return seed + "q", H
		} else if 39 <= r && r < 40 {
			return seed + "r", H
		} else if 40 <= r && r < 41 {
			return seed + "s", H
		} else if 41 <= r && r < 42 {
			return seed + "t", H
		} else if 42 <= r && r < 43 {
			return seed + "v", H
		} else if 43 <= r && r < 44 {
			return seed + "w", H
		} else if 44 <= r && r < 45 {
			return seed + "x", H
		} else if 45 <= r && r < 46 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ud" {
		r := rng.IntN(140)
		H := 2.523897290317083
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 76 {
			return seed + "e", H
		} else if 76 <= r && r < 94 {
			return seed + "g", H
		} else if 94 <= r && r < 127 {
			return seed + "i", H
		} else if 127 <= r && r < 134 {
			return seed + "o", H
		} else if 134 <= r && r < 136 {
			return seed + "s", H
		} else if 136 <= r && r < 137 {
			return seed + "u", H
		} else if 137 <= r && r < 140 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rm" {
		r := rng.IntN(140)
		H := 2.9374634997521354
		if 0 <= r && r < 35 {
			return seed + "a", H
		} else if 35 <= r && r < 37 {
			return seed + "b", H
		} else if 37 <= r && r < 39 {
			return seed + "c", H
		} else if 39 <= r && r < 62 {
			return seed + "e", H
		} else if 62 <= r && r < 66 {
			return seed + "f", H
		} else if 66 <= r && r < 68 {
			return seed + "h", H
		} else if 68 <= r && r < 91 {
			return seed + "i", H
		} else if 91 <= r && r < 99 {
			return seed + "l", H
		} else if 99 <= r && r < 128 {
			return seed + "o", H
		} else if 128 <= r && r < 130 {
			return seed + "p", H
		} else if 130 <= r && r < 132 {
			return seed + "r", H
		} else if 132 <= r && r < 137 {
			return seed + "u", H
		} else if 137 <= r && r < 140 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gh" {
		r := rng.IntN(100)
		H := 0.816909740308999
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 96 {
			return seed + "t", H
		} else if 96 <= r && r < 97 {
			return seed + "u", H
		} else if 97 <= r && r < 100 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "z" {
		r := rng.IntN(50)
		H := 2.0248395683071
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 31 {
			return seed + "i", H
		} else if 31 <= r && r < 48 {
			return seed + "o", H
		} else if 48 <= r && r < 49 {
			return seed + "u", H
		} else if 49 <= r && r < 50 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "s" {
		r := rng.IntN(2180)
		H := 3.620539902377829
		if 0 <= r && r < 185 {
			return seed + "a", H
		} else if 185 <= r && r < 355 {
			return seed + "c", H
		} else if 355 <= r && r < 500 {
			return seed + "e", H
		} else if 500 <= r && r < 730 {
			return seed + "h", H
		} else if 730 <= r && r < 841 {
			return seed + "i", H
		} else if 841 <= r && r < 917 {
			return seed + "k", H
		} else if 917 <= r && r < 1031 {
			return seed + "l", H
		} else if 1031 <= r && r < 1091 {
			return seed + "m", H
		} else if 1091 <= r && r < 1181 {
			return seed + "n", H
		} else if 1181 <= r && r < 1182 {
			return seed + "o", H
		} else if 1182 <= r && r < 1390 {
			return seed + "p", H
		} else if 1390 <= r && r < 1436 {
			return seed + "q", H
		} else if 1436 <= r && r < 1850 {
			return seed + "t", H
		} else if 1850 <= r && r < 2081 {
			return seed + "u", H
		} else if 2081 <= r && r < 2151 {
			return seed + "w", H
		} else if 2151 <= r && r < 2180 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "j" {
		r := rng.IntN(198)
		H := 2.1260339399730337
		if 0 <= r && r < 53 {
			return seed + "a", H
		} else if 53 <= r && r < 68 {
			return seed + "e", H
		} else if 68 <= r && r < 85 {
			return seed + "i", H
		} else if 85 <= r && r < 124 {
			return seed + "o", H
		} else if 124 <= r && r < 197 {
			return seed + "u", H
		} else if 197 <= r && r < 198 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nv" {
		r := rng.IntN(50)
		H := 2.1372563073635367
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 22 {
			return seed + "e", H
		} else if 22 <= r && r < 41 {
			return seed + "i", H
		} else if 41 <= r && r < 46 {
			return seed + "o", H
		} else if 46 <= r && r < 47 {
			return seed + "u", H
		} else if 47 <= r && r < 50 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "de" {
		r := rng.IntN(800)
		H := 3.681186220948558
		if 0 <= r && r < 32 {
			return seed + "a", H
		} else if 32 <= r && r < 51 {
			return seed + "b", H
		} else if 51 <= r && r < 112 {
			return seed + "c", H
		} else if 112 <= r && r < 195 {
			return seed + "d", H
		} else if 195 <= r && r < 207 {
			return seed + "e", H
		} else if 207 <= r && r < 276 {
			return seed + "f", H
		} else if 276 <= r && r < 285 {
			return seed + "g", H
		} else if 285 <= r && r < 288 {
			return seed + "h", H
		} else if 288 <= r && r < 290 {
			return seed + "i", H
		} else if 290 <= r && r < 293 {
			return seed + "j", H
		} else if 293 <= r && r < 294 {
			return seed + "k", H
		} else if 294 <= r && r < 347 {
			return seed + "l", H
		} else if 347 <= r && r < 370 {
			return seed + "m", H
		} else if 370 <= r && r < 449 {
			return seed + "n", H
		} else if 449 <= r && r < 457 {
			return seed + "o", H
		} else if 457 <= r && r < 492 {
			return seed + "p", H
		} else if 492 <= r && r < 493 {
			return seed + "q", H
		} else if 493 <= r && r < 688 {
			return seed + "r", H
		} else if 688 <= r && r < 731 {
			return seed + "s", H
		} else if 731 <= r && r < 760 {
			return seed + "t", H
		} else if 760 <= r && r < 762 {
			return seed + "u", H
		} else if 762 <= r && r < 793 {
			return seed + "v", H
		} else if 793 <= r && r < 794 {
			return seed + "w", H
		} else if 794 <= r && r < 799 {
			return seed + "x", H
		} else if 799 <= r && r < 800 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "wi" {
		r := rng.IntN(244)
		H := 3.432966611417366
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 23 {
			return seed + "d", H
		} else if 23 <= r && r < 29 {
			return seed + "e", H
		} else if 29 <= r && r < 40 {
			return seed + "f", H
		} else if 40 <= r && r < 45 {
			return seed + "g", H
		} else if 45 <= r && r < 46 {
			return seed + "h", H
		} else if 46 <= r && r < 47 {
			return seed + "j", H
		} else if 47 <= r && r < 50 {
			return seed + "k", H
		} else if 50 <= r && r < 87 {
			return seed + "l", H
		} else if 87 <= r && r < 100 {
			return seed + "m", H
		} else if 100 <= r && r < 167 {
			return seed + "n", H
		} else if 167 <= r && r < 172 {
			return seed + "p", H
		} else if 172 <= r && r < 173 {
			return seed + "q", H
		} else if 173 <= r && r < 194 {
			return seed + "r", H
		} else if 194 <= r && r < 223 {
			return seed + "s", H
		} else if 223 <= r && r < 234 {
			return seed + "t", H
		} else if 234 <= r && r < 237 {
			return seed + "v", H
		} else if 237 <= r && r < 238 {
			return seed + "w", H
		} else if 238 <= r && r < 239 {
			return seed + "x", H
		} else if 239 <= r && r < 244 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "xe" {
		r := rng.IntN(50)
		H := 3.5384810883997098
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 13 {
			return seed + "d", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 17 {
			return seed + "j", H
		} else if 17 <= r && r < 18 {
			return seed + "k", H
		} else if 18 <= r && r < 19 {
			return seed + "l", H
		} else if 19 <= r && r < 26 {
			return seed + "m", H
		} else if 26 <= r && r < 29 {
			return seed + "n", H
		} else if 29 <= r && r < 30 {
			return seed + "p", H
		} else if 30 <= r && r < 31 {
			return seed + "q", H
		} else if 31 <= r && r < 42 {
			return seed + "r", H
		} else if 42 <= r && r < 45 {
			return seed + "s", H
		} else if 45 <= r && r < 46 {
			return seed + "t", H
		} else if 46 <= r && r < 47 {
			return seed + "v", H
		} else if 47 <= r && r < 48 {
			return seed + "w", H
		} else if 48 <= r && r < 49 {
			return seed + "x", H
		} else if 49 <= r && r < 50 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "" {
		r := rng.IntN(15578)
		H := 4.214986884965409
		if 0 <= r && r < 815 {
			return seed + "a", H
		} else if 815 <= r && r < 1540 {
			return seed + "b", H
		} else if 1540 <= r && r < 3043 {
			return seed + "c", H
		} else if 3043 <= r && r < 4138 {
			return seed + "d", H
		} else if 4138 <= r && r < 4935 {
			return seed + "e", H
		} else if 4935 <= r && r < 5552 {
			return seed + "f", H
		} else if 5552 <= r && r < 6161 {
			return seed + "g", H
		} else if 6161 <= r && r < 6660 {
			return seed + "h", H
		} else if 6660 <= r && r < 6891 {
			return seed + "i", H
		} else if 6891 <= r && r < 7084 {
			return seed + "j", H
		} else if 7084 <= r && r < 7195 {
			return seed + "k", H
		} else if 7195 <= r && r < 7584 {
			return seed + "l", H
		} else if 7584 <= r && r < 8175 {
			return seed + "m", H
		} else if 8175 <= r && r < 8370 {
			return seed + "n", H
		} else if 8370 <= r && r < 8863 {
			return seed + "o", H
		} else if 8863 <= r && r < 10020 {
			return seed + "p", H
		} else if 10020 <= r && r < 10095 {
			return seed + "q", H
		} else if 10095 <= r && r < 11122 {
			return seed + "r", H
		} else if 11122 <= r && r < 13297 {
			return seed + "s", H
		} else if 13297 <= r && r < 13944 {
			return seed + "t", H
		} else if 13944 <= r && r < 14897 {
			return seed + "u", H
		} else if 14897 <= r && r < 15162 {
			return seed + "v", H
		} else if 15162 <= r && r < 15473 {
			return seed + "w", H
		} else if 15473 <= r && r < 15478 {
			return seed + "x", H
		} else if 15478 <= r && r < 15533 {
			return seed + "y", H
		} else if 15533 <= r && r < 15578 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "d" {
		r := rng.IntN(1100)
		H := 2.420515893221568
		if 0 <= r && r < 101 {
			return seed + "a", H
		} else if 101 <= r && r < 502 {
			return seed + "e", H
		} else if 502 <= r && r < 779 {
			return seed + "i", H
		} else if 779 <= r && r < 878 {
			return seed + "o", H
		} else if 878 <= r && r < 1004 {
			return seed + "r", H
		} else if 1004 <= r && r < 1075 {
			return seed + "u", H
		} else if 1075 <= r && r < 1089 {
			return seed + "w", H
		} else if 1089 <= r && r < 1100 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wg" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ua" {
		r := rng.IntN(174)
		H := 3.2972216027926184
		if 0 <= r && r < 11 {
			return seed + "b", H
		} else if 11 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 27 {
			return seed + "d", H
		} else if 27 <= r && r < 28 {
			return seed + "f", H
		} else if 28 <= r && r < 31 {
			return seed + "g", H
		} else if 31 <= r && r < 34 {
			return seed + "h", H
		} else if 34 <= r && r < 40 {
			return seed + "i", H
		} else if 40 <= r && r < 41 {
			return seed + "j", H
		} else if 41 <= r && r < 46 {
			return seed + "k", H
		} else if 46 <= r && r < 97 {
			return seed + "l", H
		} else if 97 <= r && r < 98 {
			return seed + "m", H
		} else if 98 <= r && r < 107 {
			return seed + "n", H
		} else if 107 <= r && r < 108 {
			return seed + "p", H
		} else if 108 <= r && r < 109 {
			return seed + "q", H
		} else if 109 <= r && r < 134 {
			return seed + "r", H
		} else if 134 <= r && r < 137 {
			return seed + "s", H
		} else if 137 <= r && r < 168 {
			return seed + "t", H
		} else if 168 <= r && r < 171 {
			return seed + "v", H
		} else if 171 <= r && r < 172 {
			return seed + "w", H
		} else if 172 <= r && r < 173 {
			return seed + "x", H
		} else if 173 <= r && r < 174 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "co" {
		r := rng.IntN(718)
		H := 3.5540498028797405
		if 0 <= r && r < 32 {
			return seed + "a", H
		} else if 32 <= r && r < 41 {
			return seed + "b", H
		} else if 41 <= r && r < 46 {
			return seed + "c", H
		} else if 46 <= r && r < 57 {
			return seed + "d", H
		} else if 57 <= r && r < 63 {
			return seed + "e", H
		} else if 63 <= r && r < 70 {
			return seed + "f", H
		} else if 70 <= r && r < 79 {
			return seed + "g", H
		} else if 79 <= r && r < 86 {
			return seed + "h", H
		} else if 86 <= r && r < 92 {
			return seed + "i", H
		} else if 92 <= r && r < 93 {
			return seed + "j", H
		} else if 93 <= r && r < 96 {
			return seed + "k", H
		} else if 96 <= r && r < 153 {
			return seed + "l", H
		} else if 153 <= r && r < 258 {
			return seed + "m", H
		} else if 258 <= r && r < 445 {
			return seed + "n", H
		} else if 445 <= r && r < 461 {
			return seed + "o", H
		} else if 461 <= r && r < 488 {
			return seed + "p", H
		} else if 488 <= r && r < 489 {
			return seed + "q", H
		} else if 489 <= r && r < 592 {
			return seed + "r", H
		} else if 592 <= r && r < 613 {
			return seed + "s", H
		} else if 613 <= r && r < 630 {
			return seed + "t", H
		} else if 630 <= r && r < 686 {
			return seed + "u", H
		} else if 686 <= r && r < 703 {
			return seed + "v", H
		} else if 703 <= r && r < 706 {
			return seed + "w", H
		} else if 706 <= r && r < 707 {
			return seed + "x", H
		} else if 707 <= r && r < 711 {
			return seed + "y", H
		} else if 711 <= r && r < 718 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "gz" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ry" {
		r := rng.IntN(44)
		H := 3.9896353808405456
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 7 {
			return seed + "d", H
		} else if 7 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 9 {
			return seed + "g", H
		} else if 9 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 18 {
			return seed + "i", H
		} else if 18 <= r && r < 19 {
			return seed + "j", H
		} else if 19 <= r && r < 20 {
			return seed + "k", H
		} else if 20 <= r && r < 21 {
			return seed + "l", H
		} else if 21 <= r && r < 22 {
			return seed + "m", H
		} else if 22 <= r && r < 23 {
			return seed + "n", H
		} else if 23 <= r && r < 25 {
			return seed + "o", H
		} else if 25 <= r && r < 32 {
			return seed + "p", H
		} else if 32 <= r && r < 33 {
			return seed + "q", H
		} else if 33 <= r && r < 34 {
			return seed + "r", H
		} else if 34 <= r && r < 37 {
			return seed + "s", H
		} else if 37 <= r && r < 38 {
			return seed + "t", H
		} else if 38 <= r && r < 39 {
			return seed + "v", H
		} else if 39 <= r && r < 42 {
			return seed + "w", H
		} else if 42 <= r && r < 43 {
			return seed + "x", H
		} else if 43 <= r && r < 44 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "dg" {
		r := rng.IntN(72)
		H := 1.6487333932419523
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 48 {
			return seed + "e", H
		} else if 48 <= r && r < 59 {
			return seed + "i", H
		} else if 59 <= r && r < 60 {
			return seed + "o", H
		} else if 60 <= r && r < 62 {
			return seed + "r", H
		} else if 62 <= r && r < 65 {
			return seed + "u", H
		} else if 65 <= r && r < 72 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dp" {
		r := rng.IntN(24)
		H := 2.6320139422512874
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 22 {
			return seed + "r", H
		} else if 22 <= r && r < 23 {
			return seed + "u", H
		} else if 23 <= r && r < 24 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oj" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "uj" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "b" {
		r := rng.IntN(730)
		H := 2.2653853023696926
		if 0 <= r && r < 253 {
			return seed + "a", H
		} else if 253 <= r && r < 254 {
			return seed + "e", H
		} else if 254 <= r && r < 255 {
			return seed + "i", H
		} else if 255 <= r && r < 361 {
			return seed + "l", H
		} else if 361 <= r && r < 504 {
			return seed + "o", H
		} else if 504 <= r && r < 618 {
			return seed + "r", H
		} else if 618 <= r && r < 729 {
			return seed + "u", H
		} else if 729 <= r && r < 730 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ur" {
		r := rng.IntN(494)
		H := 3.7817560048348255
		if 0 <= r && r < 41 {
			return seed + "a", H
		} else if 41 <= r && r < 59 {
			return seed + "b", H
		} else if 59 <= r && r < 71 {
			return seed + "c", H
		} else if 71 <= r && r < 83 {
			return seed + "d", H
		} else if 83 <= r && r < 210 {
			return seed + "e", H
		} else if 210 <= r && r < 224 {
			return seed + "f", H
		} else if 224 <= r && r < 250 {
			return seed + "g", H
		} else if 250 <= r && r < 305 {
			return seed + "i", H
		} else if 305 <= r && r < 313 {
			return seed + "k", H
		} else if 313 <= r && r < 331 {
			return seed + "l", H
		} else if 331 <= r && r < 335 {
			return seed + "m", H
		} else if 335 <= r && r < 355 {
			return seed + "n", H
		} else if 355 <= r && r < 372 {
			return seed + "o", H
		} else if 372 <= r && r < 388 {
			return seed + "p", H
		} else if 388 <= r && r < 416 {
			return seed + "r", H
		} else if 416 <= r && r < 446 {
			return seed + "s", H
		} else if 446 <= r && r < 462 {
			return seed + "t", H
		} else if 462 <= r && r < 467 {
			return seed + "u", H
		} else if 467 <= r && r < 487 {
			return seed + "v", H
		} else if 487 <= r && r < 494 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zz" {
		r := rng.IntN(46)
		H := 1.9348185196451597
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 39 {
			return seed + "l", H
		} else if 39 <= r && r < 40 {
			return seed + "o", H
		} else if 40 <= r && r < 41 {
			return seed + "u", H
		} else if 41 <= r && r < 46 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ht" {
		r := rng.IntN(44)
		H := 2.865884828083969
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 20 {
			return seed + "f", H
		} else if 20 <= r && r < 25 {
			return seed + "i", H
		} else if 25 <= r && r < 31 {
			return seed + "l", H
		} else if 31 <= r && r < 33 {
			return seed + "n", H
		} else if 33 <= r && r < 34 {
			return seed + "o", H
		} else if 34 <= r && r < 36 {
			return seed + "r", H
		} else if 36 <= r && r < 39 {
			return seed + "u", H
		} else if 39 <= r && r < 41 {
			return seed + "w", H
		} else if 41 <= r && r < 44 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wc" {
		r := rng.IntN(10)
		H := 2.160964047443681
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ei" {
		r := rng.IntN(62)
		H := 3.74238148993631
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 11 {
			return seed + "g", H
		} else if 11 <= r && r < 12 {
			return seed + "h", H
		} else if 12 <= r && r < 13 {
			return seed + "j", H
		} else if 13 <= r && r < 14 {
			return seed + "k", H
		} else if 14 <= r && r < 19 {
			return seed + "l", H
		} else if 19 <= r && r < 22 {
			return seed + "m", H
		} else if 22 <= r && r < 31 {
			return seed + "n", H
		} else if 31 <= r && r < 32 {
			return seed + "p", H
		} else if 32 <= r && r < 33 {
			return seed + "q", H
		} else if 33 <= r && r < 34 {
			return seed + "r", H
		} else if 34 <= r && r < 41 {
			return seed + "s", H
		} else if 41 <= r && r < 50 {
			return seed + "t", H
		} else if 50 <= r && r < 57 {
			return seed + "v", H
		} else if 57 <= r && r < 58 {
			return seed + "w", H
		} else if 58 <= r && r < 59 {
			return seed + "x", H
		} else if 59 <= r && r < 62 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "gn" {
		r := rng.IntN(80)
		H := 2.1071429380220454
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 46 {
			return seed + "e", H
		} else if 46 <= r && r < 63 {
			return seed + "i", H
		} else if 63 <= r && r < 78 {
			return seed + "o", H
		} else if 78 <= r && r < 79 {
			return seed + "u", H
		} else if 79 <= r && r < 80 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jo" {
		r := rng.IntN(84)
		H := 3.699024065880204
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 9 {
			return seed + "d", H
		} else if 9 <= r && r < 10 {
			return seed + "f", H
		} else if 10 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 18 {
			return seed + "h", H
		} else if 18 <= r && r < 30 {
			return seed + "i", H
		} else if 30 <= r && r < 31 {
			return seed + "j", H
		} else if 31 <= r && r < 36 {
			return seed + "k", H
		} else if 36 <= r && r < 43 {
			return seed + "l", H
		} else if 43 <= r && r < 44 {
			return seed + "m", H
		} else if 44 <= r && r < 45 {
			return seed + "n", H
		} else if 45 <= r && r < 46 {
			return seed + "p", H
		} else if 46 <= r && r < 47 {
			return seed + "q", H
		} else if 47 <= r && r < 52 {
			return seed + "r", H
		} else if 52 <= r && r < 53 {
			return seed + "s", H
		} else if 53 <= r && r < 56 {
			return seed + "t", H
		} else if 56 <= r && r < 59 {
			return seed + "v", H
		} else if 59 <= r && r < 60 {
			return seed + "w", H
		} else if 60 <= r && r < 61 {
			return seed + "x", H
		} else if 61 <= r && r < 83 {
			return seed + "y", H
		} else if 83 <= r && r < 84 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lm" {
		r := rng.IntN(24)
		H := 2.402439983737665
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 24 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rx" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "do" {
		r := rng.IntN(226)
		H := 3.895791793421256
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 18 {
			return seed + "c", H
		} else if 18 <= r && r < 23 {
			return seed + "d", H
		} else if 23 <= r && r < 25 {
			return seed + "e", H
		} else if 25 <= r && r < 28 {
			return seed + "f", H
		} else if 28 <= r && r < 35 {
			return seed + "g", H
		} else if 35 <= r && r < 36 {
			return seed + "h", H
		} else if 36 <= r && r < 42 {
			return seed + "i", H
		} else if 42 <= r && r < 43 {
			return seed + "j", H
		} else if 43 <= r && r < 46 {
			return seed + "k", H
		} else if 46 <= r && r < 63 {
			return seed + "l", H
		} else if 63 <= r && r < 90 {
			return seed + "m", H
		} else if 90 <= r && r < 109 {
			return seed + "n", H
		} else if 109 <= r && r < 137 {
			return seed + "o", H
		} else if 137 <= r && r < 138 {
			return seed + "p", H
		} else if 138 <= r && r < 139 {
			return seed + "q", H
		} else if 139 <= r && r < 170 {
			return seed + "r", H
		} else if 170 <= r && r < 177 {
			return seed + "s", H
		} else if 177 <= r && r < 182 {
			return seed + "t", H
		} else if 182 <= r && r < 186 {
			return seed + "u", H
		} else if 186 <= r && r < 189 {
			return seed + "v", H
		} else if 189 <= r && r < 218 {
			return seed + "w", H
		} else if 218 <= r && r < 221 {
			return seed + "x", H
		} else if 221 <= r && r < 226 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "ak" {
		r := rng.IntN(116)
		H := 1.2585437040266554
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 86 {
			return seed + "e", H
		} else if 86 <= r && r < 109 {
			return seed + "i", H
		} else if 109 <= r && r < 110 {
			return seed + "o", H
		} else if 110 <= r && r < 111 {
			return seed + "u", H
		} else if 111 <= r && r < 116 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kn" {
		r := rng.IntN(42)
		H := 2.1073126561406172
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 20 {
			return seed + "e", H
		} else if 20 <= r && r < 27 {
			return seed + "i", H
		} else if 27 <= r && r < 40 {
			return seed + "o", H
		} else if 40 <= r && r < 41 {
			return seed + "u", H
		} else if 41 <= r && r < 42 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yc" {
		r := rng.IntN(34)
		H := 2.1632517235279303
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 10 {
			return seed + "h", H
		} else if 10 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 29 {
			return seed + "l", H
		} else if 29 <= r && r < 32 {
			return seed + "o", H
		} else if 32 <= r && r < 33 {
			return seed + "u", H
		} else if 33 <= r && r < 34 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ky" {
		r := rng.IntN(36)
		H := 4.09689193246757
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 9 {
			return seed + "d", H
		} else if 9 <= r && r < 10 {
			return seed + "f", H
		} else if 10 <= r && r < 11 {
			return seed + "g", H
		} else if 11 <= r && r < 12 {
			return seed + "h", H
		} else if 12 <= r && r < 13 {
			return seed + "j", H
		} else if 13 <= r && r < 14 {
			return seed + "k", H
		} else if 14 <= r && r < 19 {
			return seed + "l", H
		} else if 19 <= r && r < 20 {
			return seed + "m", H
		} else if 20 <= r && r < 21 {
			return seed + "n", H
		} else if 21 <= r && r < 24 {
			return seed + "p", H
		} else if 24 <= r && r < 25 {
			return seed + "q", H
		} else if 25 <= r && r < 28 {
			return seed + "r", H
		} else if 28 <= r && r < 29 {
			return seed + "s", H
		} else if 29 <= r && r < 30 {
			return seed + "t", H
		} else if 30 <= r && r < 31 {
			return seed + "v", H
		} else if 31 <= r && r < 34 {
			return seed + "w", H
		} else if 34 <= r && r < 35 {
			return seed + "x", H
		} else if 35 <= r && r < 36 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "qk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hr" {
		r := rng.IntN(92)
		H := 2.2872612279546445
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 53 {
			return seed + "i", H
		} else if 53 <= r && r < 82 {
			return seed + "o", H
		} else if 82 <= r && r < 91 {
			return seed + "u", H
		} else if 91 <= r && r < 92 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ij" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zf" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ti" {
		r := rng.IntN(1020)
		H := 3.5429641475360745
		if 0 <= r && r < 12 {
			return seed + "a", H
		} else if 12 <= r && r < 17 {
			return seed + "b", H
		} else if 17 <= r && r < 148 {
			return seed + "c", H
		} else if 148 <= r && r < 165 {
			return seed + "d", H
		} else if 165 <= r && r < 187 {
			return seed + "e", H
		} else if 187 <= r && r < 234 {
			return seed + "f", H
		} else if 234 <= r && r < 259 {
			return seed + "g", H
		} else if 259 <= r && r < 262 {
			return seed + "h", H
		} else if 262 <= r && r < 263 {
			return seed + "j", H
		} else if 263 <= r && r < 264 {
			return seed + "k", H
		} else if 264 <= r && r < 313 {
			return seed + "l", H
		} else if 313 <= r && r < 356 {
			return seed + "m", H
		} else if 356 <= r && r < 593 {
			return seed + "n", H
		} else if 593 <= r && r < 773 {
			return seed + "o", H
		} else if 773 <= r && r < 796 {
			return seed + "p", H
		} else if 796 <= r && r < 803 {
			return seed + "q", H
		} else if 803 <= r && r < 826 {
			return seed + "r", H
		} else if 826 <= r && r < 867 {
			return seed + "s", H
		} else if 867 <= r && r < 916 {
			return seed + "t", H
		} else if 916 <= r && r < 997 {
			return seed + "v", H
		} else if 997 <= r && r < 998 {
			return seed + "w", H
		} else if 998 <= r && r < 999 {
			return seed + "x", H
		} else if 999 <= r && r < 1020 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "id" {
		r := rng.IntN(304)
		H := 2.4697126333393777
		if 0 <= r && r < 21 {
			return seed + "a", H
		} else if 21 <= r && r < 23 {
			return seed + "b", H
		} else if 23 <= r && r < 45 {
			return seed + "d", H
		} else if 45 <= r && r < 202 {
			return seed + "e", H
		} else if 202 <= r && r < 220 {
			return seed + "g", H
		} else if 220 <= r && r < 261 {
			return seed + "i", H
		} else if 261 <= r && r < 275 {
			return seed + "l", H
		} else if 275 <= r && r < 279 {
			return seed + "n", H
		} else if 279 <= r && r < 286 {
			return seed + "o", H
		} else if 286 <= r && r < 288 {
			return seed + "s", H
		} else if 288 <= r && r < 290 {
			return seed + "t", H
		} else if 290 <= r && r < 297 {
			return seed + "u", H
		} else if 297 <= r && r < 304 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qa" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "up" {
		r := rng.IntN(178)
		H := 3.7462676644761825
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 17 {
			return seed + "c", H
		} else if 17 <= r && r < 19 {
			return seed + "d", H
		} else if 19 <= r && r < 44 {
			return seed + "e", H
		} else if 44 <= r && r < 46 {
			return seed + "f", H
		} else if 46 <= r && r < 48 {
			return seed + "g", H
		} else if 48 <= r && r < 58 {
			return seed + "h", H
		} else if 58 <= r && r < 69 {
			return seed + "i", H
		} else if 69 <= r && r < 85 {
			return seed + "l", H
		} else if 85 <= r && r < 90 {
			return seed + "o", H
		} else if 90 <= r && r < 112 {
			return seed + "p", H
		} else if 112 <= r && r < 128 {
			return seed + "r", H
		} else if 128 <= r && r < 146 {
			return seed + "s", H
		} else if 146 <= r && r < 166 {
			return seed + "t", H
		} else if 166 <= r && r < 167 {
			return seed + "u", H
		} else if 167 <= r && r < 171 {
			return seed + "w", H
		} else if 171 <= r && r < 178 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ey" {
		r := rng.IntN(38)
		H := 4.306494947369446
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 9 {
			return seed + "d", H
		} else if 9 <= r && r < 11 {
			return seed + "e", H
		} else if 11 <= r && r < 12 {
			return seed + "f", H
		} else if 12 <= r && r < 13 {
			return seed + "g", H
		} else if 13 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 17 {
			return seed + "j", H
		} else if 17 <= r && r < 18 {
			return seed + "k", H
		} else if 18 <= r && r < 21 {
			return seed + "l", H
		} else if 21 <= r && r < 24 {
			return seed + "m", H
		} else if 24 <= r && r < 25 {
			return seed + "n", H
		} else if 25 <= r && r < 27 {
			return seed + "o", H
		} else if 27 <= r && r < 28 {
			return seed + "p", H
		} else if 28 <= r && r < 29 {
			return seed + "q", H
		} else if 29 <= r && r < 30 {
			return seed + "r", H
		} else if 30 <= r && r < 31 {
			return seed + "s", H
		} else if 31 <= r && r < 32 {
			return seed + "t", H
		} else if 32 <= r && r < 33 {
			return seed + "v", H
		} else if 33 <= r && r < 36 {
			return seed + "w", H
		} else if 36 <= r && r < 37 {
			return seed + "x", H
		} else if 37 <= r && r < 38 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "js" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ai" {
		r := rng.IntN(284)
		H := 2.656586649075426
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 31 {
			return seed + "d", H
		} else if 31 <= r && r < 32 {
			return seed + "f", H
		} else if 32 <= r && r < 33 {
			return seed + "g", H
		} else if 33 <= r && r < 34 {
			return seed + "h", H
		} else if 34 <= r && r < 35 {
			return seed + "j", H
		} else if 35 <= r && r < 38 {
			return seed + "k", H
		} else if 38 <= r && r < 103 {
			return seed + "l", H
		} else if 103 <= r && r < 120 {
			return seed + "m", H
		} else if 120 <= r && r < 233 {
			return seed + "n", H
		} else if 233 <= r && r < 234 {
			return seed + "p", H
		} else if 234 <= r && r < 235 {
			return seed + "q", H
		} else if 235 <= r && r < 264 {
			return seed + "r", H
		} else if 264 <= r && r < 275 {
			return seed + "s", H
		} else if 275 <= r && r < 280 {
			return seed + "t", H
		} else if 280 <= r && r < 281 {
			return seed + "v", H
		} else if 281 <= r && r < 282 {
			return seed + "w", H
		} else if 282 <= r && r < 283 {
			return seed + "x", H
		} else if 283 <= r && r < 284 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vf" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ov" {
		r := rng.IntN(326)
		H := 1.1724551119979107
		if 0 <= r && r < 29 {
			return seed + "a", H
		} else if 29 <= r && r < 280 {
			return seed + "e", H
		} else if 280 <= r && r < 313 {
			return seed + "i", H
		} else if 313 <= r && r < 320 {
			return seed + "o", H
		} else if 320 <= r && r < 321 {
			return seed + "u", H
		} else if 321 <= r && r < 326 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qy" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "uv" {
		r := rng.IntN(12)
		H := 1.9473387961875537
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 9 {
			return seed + "i", H
		} else if 9 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gu" {
		r := rng.IntN(158)
		H := 3.7017030768889736
		if 0 <= r && r < 14 {
			return seed + "a", H
		} else if 14 <= r && r < 15 {
			return seed + "b", H
		} else if 15 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 17 {
			return seed + "d", H
		} else if 17 <= r && r < 33 {
			return seed + "e", H
		} else if 33 <= r && r < 34 {
			return seed + "f", H
		} else if 34 <= r && r < 35 {
			return seed + "g", H
		} else if 35 <= r && r < 36 {
			return seed + "h", H
		} else if 36 <= r && r < 58 {
			return seed + "i", H
		} else if 58 <= r && r < 59 {
			return seed + "j", H
		} else if 59 <= r && r < 60 {
			return seed + "k", H
		} else if 60 <= r && r < 87 {
			return seed + "l", H
		} else if 87 <= r && r < 100 {
			return seed + "m", H
		} else if 100 <= r && r < 103 {
			return seed + "n", H
		} else if 103 <= r && r < 105 {
			return seed + "o", H
		} else if 105 <= r && r < 108 {
			return seed + "p", H
		} else if 108 <= r && r < 109 {
			return seed + "q", H
		} else if 109 <= r && r < 128 {
			return seed + "r", H
		} else if 128 <= r && r < 143 {
			return seed + "s", H
		} else if 143 <= r && r < 150 {
			return seed + "t", H
		} else if 150 <= r && r < 151 {
			return seed + "v", H
		} else if 151 <= r && r < 152 {
			return seed + "w", H
		} else if 152 <= r && r < 153 {
			return seed + "x", H
		} else if 153 <= r && r < 155 {
			return seed + "y", H
		} else if 155 <= r && r < 158 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "go" {
		r := rng.IntN(154)
		H := 3.8030358808874998
		if 0 <= r && r < 6 {
			return seed + "a", H
		} else if 6 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 11 {
			return seed + "d", H
		} else if 11 <= r && r < 13 {
			return seed + "e", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 17 {
			return seed + "g", H
		} else if 17 <= r && r < 18 {
			return seed + "h", H
		} else if 18 <= r && r < 24 {
			return seed + "i", H
		} else if 24 <= r && r < 25 {
			return seed + "j", H
		} else if 25 <= r && r < 26 {
			return seed + "k", H
		} else if 26 <= r && r < 41 {
			return seed + "l", H
		} else if 41 <= r && r < 44 {
			return seed + "m", H
		} else if 44 <= r && r < 77 {
			return seed + "n", H
		} else if 77 <= r && r < 95 {
			return seed + "o", H
		} else if 95 <= r && r < 98 {
			return seed + "p", H
		} else if 98 <= r && r < 99 {
			return seed + "q", H
		} else if 99 <= r && r < 116 {
			return seed + "r", H
		} else if 116 <= r && r < 123 {
			return seed + "s", H
		} else if 123 <= r && r < 138 {
			return seed + "t", H
		} else if 138 <= r && r < 146 {
			return seed + "u", H
		} else if 146 <= r && r < 149 {
			return seed + "v", H
		} else if 149 <= r && r < 152 {
			return seed + "w", H
		} else if 152 <= r && r < 153 {
			return seed + "x", H
		} else if 153 <= r && r < 154 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "sq" {
		r := rng.IntN(56)
		H := 0.6413961284764207
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 55 {
			return seed + "u", H
		} else if 55 <= r && r < 56 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "xx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wh" {
		r := rng.IntN(56)
		H := 2.26098623064271
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 24 {
			return seed + "e", H
		} else if 24 <= r && r < 39 {
			return seed + "i", H
		} else if 39 <= r && r < 52 {
			return seed + "o", H
		} else if 52 <= r && r < 53 {
			return seed + "u", H
		} else if 53 <= r && r < 56 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "va" {
		r := rng.IntN(236)
		H := 3.3332780140899363
		if 0 <= r && r < 23 {
			return seed + "b", H
		} else if 23 <= r && r < 40 {
			return seed + "c", H
		} else if 40 <= r && r < 43 {
			return seed + "d", H
		} else if 43 <= r && r < 44 {
			return seed + "f", H
		} else if 44 <= r && r < 63 {
			return seed + "g", H
		} else if 63 <= r && r < 64 {
			return seed + "h", H
		} else if 64 <= r && r < 68 {
			return seed + "i", H
		} else if 68 <= r && r < 69 {
			return seed + "j", H
		} else if 69 <= r && r < 70 {
			return seed + "k", H
		} else if 70 <= r && r < 127 {
			return seed + "l", H
		} else if 127 <= r && r < 128 {
			return seed + "m", H
		} else if 128 <= r && r < 149 {
			return seed + "n", H
		} else if 149 <= r && r < 154 {
			return seed + "p", H
		} else if 154 <= r && r < 155 {
			return seed + "q", H
		} else if 155 <= r && r < 180 {
			return seed + "r", H
		} else if 180 <= r && r < 193 {
			return seed + "s", H
		} else if 193 <= r && r < 232 {
			return seed + "t", H
		} else if 232 <= r && r < 233 {
			return seed + "v", H
		} else if 233 <= r && r < 234 {
			return seed + "w", H
		} else if 234 <= r && r < 235 {
			return seed + "x", H
		} else if 235 <= r && r < 236 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pu" {
		r := rng.IntN(198)
		H := 3.2823217078353077
		if 0 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 11 {
			return seed + "d", H
		} else if 11 <= r && r < 13 {
			return seed + "e", H
		} else if 13 <= r && r < 14 {
			return seed + "f", H
		} else if 14 <= r && r < 19 {
			return seed + "g", H
		} else if 19 <= r && r < 20 {
			return seed + "h", H
		} else if 20 <= r && r < 21 {
			return seed + "j", H
		} else if 21 <= r && r < 22 {
			return seed + "k", H
		} else if 22 <= r && r < 59 {
			return seed + "l", H
		} else if 59 <= r && r < 66 {
			return seed + "m", H
		} else if 66 <= r && r < 81 {
			return seed + "n", H
		} else if 81 <= r && r < 88 {
			return seed + "p", H
		} else if 88 <= r && r < 89 {
			return seed + "q", H
		} else if 89 <= r && r < 142 {
			return seed + "r", H
		} else if 142 <= r && r < 167 {
			return seed + "s", H
		} else if 167 <= r && r < 190 {
			return seed + "t", H
		} else if 190 <= r && r < 191 {
			return seed + "v", H
		} else if 191 <= r && r < 192 {
			return seed + "w", H
		} else if 192 <= r && r < 193 {
			return seed + "x", H
		} else if 193 <= r && r < 198 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "pn" {
		r := rng.IntN(28)
		H := 1.7302312027417304
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 26 {
			return seed + "o", H
		} else if 26 <= r && r < 27 {
			return seed + "u", H
		} else if 27 <= r && r < 28 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ul" {
		r := rng.IntN(286)
		H := 3.3765542424051302
		if 0 <= r && r < 67 {
			return seed + "a", H
		} else if 67 <= r && r < 73 {
			return seed + "b", H
		} else if 73 <= r && r < 75 {
			return seed + "c", H
		} else if 75 <= r && r < 77 {
			return seed + "d", H
		} else if 77 <= r && r < 102 {
			return seed + "e", H
		} else if 102 <= r && r < 114 {
			return seed + "f", H
		} else if 114 <= r && r < 120 {
			return seed + "g", H
		} else if 120 <= r && r < 133 {
			return seed + "i", H
		} else if 133 <= r && r < 139 {
			return seed + "k", H
		} else if 139 <= r && r < 199 {
			return seed + "l", H
		} else if 199 <= r && r < 203 {
			return seed + "m", H
		} else if 203 <= r && r < 206 {
			return seed + "o", H
		} else if 206 <= r && r < 222 {
			return seed + "p", H
		} else if 222 <= r && r < 236 {
			return seed + "s", H
		} else if 236 <= r && r < 272 {
			return seed + "t", H
		} else if 272 <= r && r < 277 {
			return seed + "u", H
		} else if 277 <= r && r < 279 {
			return seed + "v", H
		} else if 279 <= r && r < 286 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "we" {
		r := rng.IntN(136)
		H := 3.360525984295571
		if 0 <= r && r < 16 {
			return seed + "a", H
		} else if 16 <= r && r < 19 {
			return seed + "b", H
		} else if 19 <= r && r < 20 {
			return seed + "c", H
		} else if 20 <= r && r < 37 {
			return seed + "d", H
		} else if 37 <= r && r < 53 {
			return seed + "e", H
		} else if 53 <= r && r < 54 {
			return seed + "f", H
		} else if 54 <= r && r < 55 {
			return seed + "g", H
		} else if 55 <= r && r < 56 {
			return seed + "h", H
		} else if 56 <= r && r < 58 {
			return seed + "i", H
		} else if 58 <= r && r < 59 {
			return seed + "j", H
		} else if 59 <= r && r < 60 {
			return seed + "k", H
		} else if 60 <= r && r < 79 {
			return seed + "l", H
		} else if 79 <= r && r < 80 {
			return seed + "m", H
		} else if 80 <= r && r < 87 {
			return seed + "n", H
		} else if 87 <= r && r < 90 {
			return seed + "p", H
		} else if 90 <= r && r < 91 {
			return seed + "q", H
		} else if 91 <= r && r < 128 {
			return seed + "r", H
		} else if 128 <= r && r < 131 {
			return seed + "s", H
		} else if 131 <= r && r < 132 {
			return seed + "t", H
		} else if 132 <= r && r < 133 {
			return seed + "v", H
		} else if 133 <= r && r < 134 {
			return seed + "w", H
		} else if 134 <= r && r < 135 {
			return seed + "x", H
		} else if 135 <= r && r < 136 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "re" {
		r := rng.IntN(1508)
		H := 4.0906412540230175
		if 0 <= r && r < 176 {
			return seed + "a", H
		} else if 176 <= r && r < 209 {
			return seed + "b", H
		} else if 209 <= r && r < 328 {
			return seed + "c", H
		} else if 328 <= r && r < 441 {
			return seed + "d", H
		} else if 441 <= r && r < 557 {
			return seed + "e", H
		} else if 557 <= r && r < 638 {
			return seed + "f", H
		} else if 638 <= r && r < 683 {
			return seed + "g", H
		} else if 683 <= r && r < 696 {
			return seed + "h", H
		} else if 696 <= r && r < 704 {
			return seed + "i", H
		} else if 704 <= r && r < 711 {
			return seed + "j", H
		} else if 711 <= r && r < 716 {
			return seed + "k", H
		} else if 716 <= r && r < 803 {
			return seed + "l", H
		} else if 803 <= r && r < 866 {
			return seed + "m", H
		} else if 866 <= r && r < 953 {
			return seed + "n", H
		} else if 953 <= r && r < 969 {
			return seed + "o", H
		} else if 969 <= r && r < 1058 {
			return seed + "p", H
		} else if 1058 <= r && r < 1069 {
			return seed + "q", H
		} else if 1069 <= r && r < 1110 {
			return seed + "r", H
		} else if 1110 <= r && r < 1281 {
			return seed + "s", H
		} else if 1281 <= r && r < 1384 {
			return seed + "t", H
		} else if 1384 <= r && r < 1392 {
			return seed + "u", H
		} else if 1392 <= r && r < 1459 {
			return seed + "v", H
		} else if 1459 <= r && r < 1502 {
			return seed + "w", H
		} else if 1502 <= r && r < 1505 {
			return seed + "x", H
		} else if 1505 <= r && r < 1507 {
			return seed + "y", H
		} else if 1507 <= r && r < 1508 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "bv" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ym" {
		r := rng.IntN(44)
		H := 2.43696831036105
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 26 {
			return seed + "e", H
		} else if 26 <= r && r < 27 {
			return seed + "i", H
		} else if 27 <= r && r < 32 {
			return seed + "o", H
		} else if 32 <= r && r < 42 {
			return seed + "p", H
		} else if 42 <= r && r < 43 {
			return seed + "u", H
		} else if 43 <= r && r < 44 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yb" {
		r := rng.IntN(30)
		H := 2.6261444493561785
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 15 {
			return seed + "i", H
		} else if 15 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 26 {
			return seed + "r", H
		} else if 26 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 30 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sg" {
		r := rng.IntN(8)
		H := 2.75
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 6 {
			return seed + "r", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mn" {
		r := rng.IntN(22)
		H := 2.0532635745669556
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 19 {
			return seed + "i", H
		} else if 19 <= r && r < 20 {
			return seed + "o", H
		} else if 20 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ed" {
		r := rng.IntN(154)
		H := 2.9246637849234993
		if 0 <= r && r < 17 {
			return seed + "a", H
		} else if 17 <= r && r < 25 {
			return seed + "d", H
		} else if 25 <= r && r < 42 {
			return seed + "e", H
		} else if 42 <= r && r < 58 {
			return seed + "g", H
		} else if 58 <= r && r < 109 {
			return seed + "i", H
		} else if 109 <= r && r < 113 {
			return seed + "l", H
		} else if 113 <= r && r < 122 {
			return seed + "o", H
		} else if 122 <= r && r < 124 {
			return seed + "r", H
		} else if 124 <= r && r < 126 {
			return seed + "t", H
		} else if 126 <= r && r < 147 {
			return seed + "u", H
		} else if 147 <= r && r < 154 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "os" {
		r := rng.IntN(262)
		H := 2.888197718364442
		if 0 <= r && r < 15 {
			return seed + "a", H
		} else if 15 <= r && r < 82 {
			return seed + "e", H
		} else if 82 <= r && r < 88 {
			return seed + "h", H
		} else if 88 <= r && r < 123 {
			return seed + "i", H
		} else if 123 <= r && r < 125 {
			return seed + "l", H
		} else if 125 <= r && r < 133 {
			return seed + "m", H
		} else if 133 <= r && r < 136 {
			return seed + "o", H
		} else if 136 <= r && r < 144 {
			return seed + "p", H
		} else if 144 <= r && r < 180 {
			return seed + "s", H
		} else if 180 <= r && r < 246 {
			return seed + "t", H
		} else if 246 <= r && r < 255 {
			return seed + "u", H
		} else if 255 <= r && r < 262 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bb" {
		r := rng.IntN(110)
		H := 2.268763662285358
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 36 {
			return seed + "e", H
		} else if 36 <= r && r < 51 {
			return seed + "i", H
		} else if 51 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 96 {
			return seed + "o", H
		} else if 96 <= r && r < 97 {
			return seed + "u", H
		} else if 97 <= r && r < 110 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "eq" {
		r := rng.IntN(40)
		H := 0.8338054550605167
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 39 {
			return seed + "u", H
		} else if 39 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tj" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yd" {
		r := rng.IntN(30)
		H := 1.9600115303896288
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 28 {
			return seed + "r", H
		} else if 28 <= r && r < 29 {
			return seed + "u", H
		} else if 29 <= r && r < 30 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "sr" {
		r := rng.IntN(10)
		H := 2.3709505944546683
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 10 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nt" {
		r := rng.IntN(444)
		H := 3.0863001506476593
		if 0 <= r && r < 59 {
			return seed + "a", H
		} else if 59 <= r && r < 61 {
			return seed + "d", H
		} else if 61 <= r && r < 154 {
			return seed + "e", H
		} else if 154 <= r && r < 178 {
			return seed + "h", H
		} else if 178 <= r && r < 283 {
			return seed + "i", H
		} else if 283 <= r && r < 315 {
			return seed + "l", H
		} else if 315 <= r && r < 317 {
			return seed + "m", H
		} else if 317 <= r && r < 342 {
			return seed + "o", H
		} else if 342 <= r && r < 392 {
			return seed + "r", H
		} else if 392 <= r && r < 404 {
			return seed + "s", H
		} else if 404 <= r && r < 419 {
			return seed + "u", H
		} else if 419 <= r && r < 423 {
			return seed + "w", H
		} else if 423 <= r && r < 444 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "eh" {
		r := rng.IntN(32)
		H := 2.405945909665798
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 14 {
			return seed + "e", H
		} else if 14 <= r && r < 21 {
			return seed + "i", H
		} else if 21 <= r && r < 26 {
			return seed + "o", H
		} else if 26 <= r && r < 27 {
			return seed + "u", H
		} else if 27 <= r && r < 32 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vg" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mi" {
		r := rng.IntN(320)
		H := 3.1214578243231497
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 46 {
			return seed + "c", H
		} else if 46 <= r && r < 55 {
			return seed + "d", H
		} else if 55 <= r && r < 59 {
			return seed + "e", H
		} else if 59 <= r && r < 66 {
			return seed + "f", H
		} else if 66 <= r && r < 71 {
			return seed + "g", H
		} else if 71 <= r && r < 72 {
			return seed + "h", H
		} else if 72 <= r && r < 73 {
			return seed + "j", H
		} else if 73 <= r && r < 74 {
			return seed + "k", H
		} else if 74 <= r && r < 101 {
			return seed + "l", H
		} else if 101 <= r && r < 104 {
			return seed + "m", H
		} else if 104 <= r && r < 213 {
			return seed + "n", H
		} else if 213 <= r && r < 214 {
			return seed + "p", H
		} else if 214 <= r && r < 215 {
			return seed + "q", H
		} else if 215 <= r && r < 218 {
			return seed + "r", H
		} else if 218 <= r && r < 265 {
			return seed + "s", H
		} else if 265 <= r && r < 302 {
			return seed + "t", H
		} else if 302 <= r && r < 306 {
			return seed + "u", H
		} else if 306 <= r && r < 307 {
			return seed + "v", H
		} else if 307 <= r && r < 308 {
			return seed + "w", H
		} else if 308 <= r && r < 315 {
			return seed + "x", H
		} else if 315 <= r && r < 320 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pi" {
		r := rng.IntN(278)
		H := 3.265926837188238
		if 0 <= r && r < 8 {
			return seed + "a", H
		} else if 8 <= r && r < 9 {
			return seed + "b", H
		} else if 9 <= r && r < 26 {
			return seed + "c", H
		} else if 26 <= r && r < 39 {
			return seed + "d", H
		} else if 39 <= r && r < 61 {
			return seed + "e", H
		} else if 61 <= r && r < 64 {
			return seed + "f", H
		} else if 64 <= r && r < 65 {
			return seed + "g", H
		} else if 65 <= r && r < 66 {
			return seed + "h", H
		} else if 66 <= r && r < 67 {
			return seed + "j", H
		} else if 67 <= r && r < 68 {
			return seed + "k", H
		} else if 68 <= r && r < 101 {
			return seed + "l", H
		} else if 101 <= r && r < 102 {
			return seed + "m", H
		} else if 102 <= r && r < 201 {
			return seed + "n", H
		} else if 201 <= r && r < 211 {
			return seed + "o", H
		} else if 211 <= r && r < 218 {
			return seed + "p", H
		} else if 218 <= r && r < 219 {
			return seed + "q", H
		} else if 219 <= r && r < 248 {
			return seed + "r", H
		} else if 248 <= r && r < 261 {
			return seed + "s", H
		} else if 261 <= r && r < 272 {
			return seed + "t", H
		} else if 272 <= r && r < 274 {
			return seed + "u", H
		} else if 274 <= r && r < 275 {
			return seed + "v", H
		} else if 275 <= r && r < 276 {
			return seed + "w", H
		} else if 276 <= r && r < 277 {
			return seed + "x", H
		} else if 277 <= r && r < 278 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "rd" {
		r := rng.IntN(110)
		H := 3.5418583350404824
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 7 {
			return seed + "b", H
		} else if 7 <= r && r < 13 {
			return seed + "c", H
		} else if 13 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 38 {
			return seed + "e", H
		} else if 38 <= r && r < 42 {
			return seed + "h", H
		} else if 42 <= r && r < 59 {
			return seed + "i", H
		} else if 59 <= r && r < 69 {
			return seed + "l", H
		} else if 69 <= r && r < 71 {
			return seed + "n", H
		} else if 71 <= r && r < 80 {
			return seed + "o", H
		} else if 80 <= r && r < 90 {
			return seed + "r", H
		} else if 90 <= r && r < 94 {
			return seed + "s", H
		} else if 94 <= r && r < 99 {
			return seed + "u", H
		} else if 99 <= r && r < 105 {
			return seed + "w", H
		} else if 105 <= r && r < 110 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "th" {
		r := rng.IntN(264)
		H := 2.889653563732246
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 21 {
			return seed + "b", H
		} else if 21 <= r && r < 90 {
			return seed + "e", H
		} else if 90 <= r && r < 151 {
			return seed + "i", H
		} else if 151 <= r && r < 161 {
			return seed + "l", H
		} else if 161 <= r && r < 163 {
			return seed + "m", H
		} else if 163 <= r && r < 194 {
			return seed + "o", H
		} else if 194 <= r && r < 196 {
			return seed + "p", H
		} else if 196 <= r && r < 234 {
			return seed + "r", H
		} else if 234 <= r && r < 245 {
			return seed + "u", H
		} else if 245 <= r && r < 249 {
			return seed + "w", H
		} else if 249 <= r && r < 264 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "eu" {
		r := rng.IntN(60)
		H := 3.581670189696645
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 9 {
			return seed + "j", H
		} else if 9 <= r && r < 10 {
			return seed + "k", H
		} else if 10 <= r && r < 11 {
			return seed + "l", H
		} else if 11 <= r && r < 20 {
			return seed + "m", H
		} else if 20 <= r && r < 25 {
			return seed + "n", H
		} else if 25 <= r && r < 28 {
			return seed + "p", H
		} else if 28 <= r && r < 29 {
			return seed + "q", H
		} else if 29 <= r && r < 44 {
			return seed + "r", H
		} else if 44 <= r && r < 51 {
			return seed + "s", H
		} else if 51 <= r && r < 56 {
			return seed + "t", H
		} else if 56 <= r && r < 57 {
			return seed + "v", H
		} else if 57 <= r && r < 58 {
			return seed + "w", H
		} else if 58 <= r && r < 59 {
			return seed + "x", H
		} else if 59 <= r && r < 60 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "sw" {
		r := rng.IntN(90)
		H := 2.1484725096748734
		if 0 <= r && r < 17 {
			return seed + "a", H
		} else if 17 <= r && r < 36 {
			return seed + "e", H
		} else if 36 <= r && r < 69 {
			return seed + "i", H
		} else if 69 <= r && r < 86 {
			return seed + "o", H
		} else if 86 <= r && r < 89 {
			return seed + "u", H
		} else if 89 <= r && r < 90 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vn" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "pt" {
		r := rng.IntN(96)
		H := 2.3759844824318206
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 14 {
			return seed + "h", H
		} else if 14 <= r && r < 57 {
			return seed + "i", H
		} else if 57 <= r && r < 63 {
			return seed + "l", H
		} else if 63 <= r && r < 78 {
			return seed + "o", H
		} else if 78 <= r && r < 93 {
			return seed + "u", H
		} else if 93 <= r && r < 96 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "az" {
		r := rng.IntN(74)
		H := 2.339592248416986
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 26 {
			return seed + "e", H
		} else if 26 <= r && r < 53 {
			return seed + "i", H
		} else if 53 <= r && r < 56 {
			return seed + "o", H
		} else if 56 <= r && r < 57 {
			return seed + "u", H
		} else if 57 <= r && r < 64 {
			return seed + "y", H
		} else if 64 <= r && r < 74 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "mf" {
		r := rng.IntN(14)
		H := 2.298825245003051
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "oe" {
		r := rng.IntN(44)
		H := 4.192417846776969
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 6 {
			return seed + "f", H
		} else if 6 <= r && r < 7 {
			return seed + "g", H
		} else if 7 <= r && r < 8 {
			return seed + "h", H
		} else if 8 <= r && r < 10 {
			return seed + "i", H
		} else if 10 <= r && r < 11 {
			return seed + "j", H
		} else if 11 <= r && r < 12 {
			return seed + "k", H
		} else if 12 <= r && r < 13 {
			return seed + "l", H
		} else if 13 <= r && r < 16 {
			return seed + "m", H
		} else if 16 <= r && r < 19 {
			return seed + "n", H
		} else if 19 <= r && r < 20 {
			return seed + "p", H
		} else if 20 <= r && r < 21 {
			return seed + "q", H
		} else if 21 <= r && r < 26 {
			return seed + "r", H
		} else if 26 <= r && r < 31 {
			return seed + "s", H
		} else if 31 <= r && r < 34 {
			return seed + "t", H
		} else if 34 <= r && r < 37 {
			return seed + "v", H
		} else if 37 <= r && r < 38 {
			return seed + "w", H
		} else if 38 <= r && r < 41 {
			return seed + "x", H
		} else if 41 <= r && r < 43 {
			return seed + "y", H
		} else if 43 <= r && r < 44 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "dk" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ls" {
		r := rng.IntN(38)
		H := 2.2529025504107514
		if 0 <= r && r < 7 {
			return seed + "a", H
		} else if 7 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 31 {
			return seed + "i", H
		} else if 31 <= r && r < 32 {
			return seed + "o", H
		} else if 32 <= r && r < 36 {
			return seed + "t", H
		} else if 36 <= r && r < 37 {
			return seed + "u", H
		} else if 37 <= r && r < 38 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nr" {
		r := rng.IntN(60)
		H := 2.2601271410304116
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 43 {
			return seed + "i", H
		} else if 43 <= r && r < 52 {
			return seed + "o", H
		} else if 52 <= r && r < 59 {
			return seed + "u", H
		} else if 59 <= r && r < 60 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "gr" {
		r := rng.IntN(324)
		H := 2.0670808934740346
		if 0 <= r && r < 153 {
			return seed + "a", H
		} else if 153 <= r && r < 210 {
			return seed + "e", H
		} else if 210 <= r && r < 251 {
			return seed + "i", H
		} else if 251 <= r && r < 298 {
			return seed + "o", H
		} else if 298 <= r && r < 321 {
			return seed + "u", H
		} else if 321 <= r && r < 324 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "su" {
		r := rng.IntN(352)
		H := 3.3839603063306103
		if 0 <= r && r < 22 {
			return seed + "a", H
		} else if 22 <= r && r < 113 {
			return seed + "b", H
		} else if 113 <= r && r < 120 {
			return seed + "c", H
		} else if 120 <= r && r < 127 {
			return seed + "d", H
		} else if 127 <= r && r < 135 {
			return seed + "e", H
		} else if 135 <= r && r < 148 {
			return seed + "f", H
		} else if 148 <= r && r < 153 {
			return seed + "g", H
		} else if 153 <= r && r < 154 {
			return seed + "h", H
		} else if 154 <= r && r < 174 {
			return seed + "i", H
		} else if 174 <= r && r < 175 {
			return seed + "j", H
		} else if 175 <= r && r < 176 {
			return seed + "k", H
		} else if 176 <= r && r < 201 {
			return seed + "l", H
		} else if 201 <= r && r < 216 {
			return seed + "m", H
		} else if 216 <= r && r < 217 {
			return seed + "n", H
		} else if 217 <= r && r < 219 {
			return seed + "o", H
		} else if 219 <= r && r < 256 {
			return seed + "p", H
		} else if 256 <= r && r < 257 {
			return seed + "q", H
		} else if 257 <= r && r < 328 {
			return seed + "r", H
		} else if 328 <= r && r < 347 {
			return seed + "s", H
		} else if 347 <= r && r < 348 {
			return seed + "t", H
		} else if 348 <= r && r < 349 {
			return seed + "v", H
		} else if 349 <= r && r < 350 {
			return seed + "w", H
		} else if 350 <= r && r < 351 {
			return seed + "x", H
		} else if 351 <= r && r < 352 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "vl" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "no" {
		r := rng.IntN(222)
		H := 3.815429709608293
		if 0 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 4 {
			return seed + "c", H
		} else if 4 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 7 {
			return seed + "e", H
		} else if 7 <= r && r < 8 {
			return seed + "f", H
		} else if 8 <= r && r < 15 {
			return seed + "g", H
		} else if 15 <= r && r < 16 {
			return seed + "h", H
		} else if 16 <= r && r < 18 {
			return seed + "i", H
		} else if 18 <= r && r < 19 {
			return seed + "j", H
		} else if 19 <= r && r < 20 {
			return seed + "k", H
		} else if 20 <= r && r < 35 {
			return seed + "l", H
		} else if 35 <= r && r < 50 {
			return seed + "m", H
		} else if 50 <= r && r < 55 {
			return seed + "n", H
		} else if 55 <= r && r < 61 {
			return seed + "o", H
		} else if 61 <= r && r < 76 {
			return seed + "p", H
		} else if 76 <= r && r < 77 {
			return seed + "q", H
		} else if 77 <= r && r < 104 {
			return seed + "r", H
		} else if 104 <= r && r < 119 {
			return seed + "s", H
		} else if 119 <= r && r < 150 {
			return seed + "t", H
		} else if 150 <= r && r < 166 {
			return seed + "u", H
		} else if 166 <= r && r < 173 {
			return seed + "v", H
		} else if 173 <= r && r < 212 {
			return seed + "w", H
		} else if 212 <= r && r < 219 {
			return seed + "x", H
		} else if 219 <= r && r < 221 {
			return seed + "y", H
		} else if 221 <= r && r < 222 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "iz" {
		r := rng.IntN(144)
		H := 1.4895321425420336
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 110 {
			return seed + "e", H
		} else if 110 <= r && r < 119 {
			return seed + "i", H
		} else if 119 <= r && r < 121 {
			return seed + "m", H
		} else if 121 <= r && r < 122 {
			return seed + "o", H
		} else if 122 <= r && r < 123 {
			return seed + "u", H
		} else if 123 <= r && r < 124 {
			return seed + "y", H
		} else if 124 <= r && r < 144 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "px" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "lt" {
		r := rng.IntN(78)
		H := 2.9756469650995516
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 18 {
			return seed + "e", H
		} else if 18 <= r && r < 24 {
			return seed + "h", H
		} else if 24 <= r && r < 47 {
			return seed + "i", H
		} else if 47 <= r && r < 50 {
			return seed + "o", H
		} else if 50 <= r && r < 58 {
			return seed + "r", H
		} else if 58 <= r && r < 60 {
			return seed + "t", H
		} else if 60 <= r && r < 65 {
			return seed + "u", H
		} else if 65 <= r && r < 74 {
			return seed + "y", H
		} else if 74 <= r && r < 78 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "en" {
		r := rng.IntN(888)
		H := 3.2545072685557557
		if 0 <= r && r < 35 {
			return seed + "a", H
		} else if 35 <= r && r < 137 {
			return seed + "c", H
		} else if 137 <= r && r < 275 {
			return seed + "d", H
		} else if 275 <= r && r < 352 {
			return seed + "e", H
		} else if 352 <= r && r < 358 {
			return seed + "f", H
		} else if 358 <= r && r < 398 {
			return seed + "g", H
		} else if 398 <= r && r < 402 {
			return seed + "h", H
		} else if 402 <= r && r < 435 {
			return seed + "i", H
		} else if 435 <= r && r < 445 {
			return seed + "j", H
		} else if 445 <= r && r < 447 {
			return seed + "k", H
		} else if 447 <= r && r < 457 {
			return seed + "l", H
		} else if 457 <= r && r < 473 {
			return seed + "n", H
		} else if 473 <= r && r < 488 {
			return seed + "o", H
		} else if 488 <= r && r < 490 {
			return seed + "p", H
		} else if 490 <= r && r < 492 {
			return seed + "q", H
		} else if 492 <= r && r < 500 {
			return seed + "r", H
		} else if 500 <= r && r < 550 {
			return seed + "s", H
		} else if 550 <= r && r < 846 {
			return seed + "t", H
		} else if 846 <= r && r < 861 {
			return seed + "u", H
		} else if 861 <= r && r < 877 {
			return seed + "v", H
		} else if 877 <= r && r < 882 {
			return seed + "y", H
		} else if 882 <= r && r < 888 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "jw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dh" {
		r := rng.IntN(18)
		H := 2.351644115153392
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 16 {
			return seed + "o", H
		} else if 16 <= r && r < 17 {
			return seed + "u", H
		} else if 17 <= r && r < 18 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kp" {
		r := rng.IntN(12)
		H := 2.396240625180289
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "po" {
		r := rng.IntN(398)
		H := 3.62035000088673
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 3 {
			return seed + "b", H
		} else if 3 <= r && r < 8 {
			return seed + "c", H
		} else if 8 <= r && r < 15 {
			return seed + "d", H
		} else if 15 <= r && r < 19 {
			return seed + "e", H
		} else if 19 <= r && r < 22 {
			return seed + "f", H
		} else if 22 <= r && r < 25 {
			return seed + "g", H
		} else if 25 <= r && r < 26 {
			return seed + "h", H
		} else if 26 <= r && r < 60 {
			return seed + "i", H
		} else if 60 <= r && r < 61 {
			return seed + "j", H
		} else if 61 <= r && r < 74 {
			return seed + "k", H
		} else if 74 <= r && r < 107 {
			return seed + "l", H
		} else if 107 <= r && r < 108 {
			return seed + "m", H
		} else if 108 <= r && r < 133 {
			return seed + "n", H
		} else if 133 <= r && r < 149 {
			return seed + "o", H
		} else if 149 <= r && r < 170 {
			return seed + "p", H
		} else if 170 <= r && r < 171 {
			return seed + "q", H
		} else if 171 <= r && r < 232 {
			return seed + "r", H
		} else if 232 <= r && r < 327 {
			return seed + "s", H
		} else if 327 <= r && r < 348 {
			return seed + "t", H
		} else if 348 <= r && r < 374 {
			return seed + "u", H
		} else if 374 <= r && r < 375 {
			return seed + "v", H
		} else if 375 <= r && r < 394 {
			return seed + "w", H
		} else if 394 <= r && r < 397 {
			return seed + "x", H
		} else if 397 <= r && r < 398 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "qi" {
		r := rng.IntN(20)
		H := 4.321928094887363
		if 0 <= r && r < 1 {
			return seed + "b", H
		} else if 1 <= r && r < 2 {
			return seed + "c", H
		} else if 2 <= r && r < 3 {
			return seed + "d", H
		} else if 3 <= r && r < 4 {
			return seed + "f", H
		} else if 4 <= r && r < 5 {
			return seed + "g", H
		} else if 5 <= r && r < 6 {
			return seed + "h", H
		} else if 6 <= r && r < 7 {
			return seed + "j", H
		} else if 7 <= r && r < 8 {
			return seed + "k", H
		} else if 8 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 10 {
			return seed + "m", H
		} else if 10 <= r && r < 11 {
			return seed + "n", H
		} else if 11 <= r && r < 12 {
			return seed + "p", H
		} else if 12 <= r && r < 13 {
			return seed + "q", H
		} else if 13 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "s", H
		} else if 15 <= r && r < 16 {
			return seed + "t", H
		} else if 16 <= r && r < 17 {
			return seed + "v", H
		} else if 17 <= r && r < 18 {
			return seed + "w", H
		} else if 18 <= r && r < 19 {
			return seed + "x", H
		} else if 19 <= r && r < 20 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "lz" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "na" {
		r := rng.IntN(342)
		H := 3.811150472403008
		if 0 <= r && r < 19 {
			return seed + "b", H
		} else if 19 <= r && r < 36 {
			return seed + "c", H
		} else if 36 <= r && r < 45 {
			return seed + "d", H
		} else if 45 <= r && r < 49 {
			return seed + "e", H
		} else if 49 <= r && r < 52 {
			return seed + "f", H
		} else if 52 <= r && r < 73 {
			return seed + "g", H
		} else if 73 <= r && r < 74 {
			return seed + "h", H
		} else if 74 <= r && r < 82 {
			return seed + "i", H
		} else if 82 <= r && r < 83 {
			return seed + "j", H
		} else if 83 <= r && r < 86 {
			return seed + "k", H
		} else if 86 <= r && r < 133 {
			return seed + "l", H
		} else if 133 <= r && r < 158 {
			return seed + "m", H
		} else if 158 <= r && r < 193 {
			return seed + "n", H
		} else if 193 <= r && r < 216 {
			return seed + "p", H
		} else if 216 <= r && r < 217 {
			return seed + "q", H
		} else if 217 <= r && r < 246 {
			return seed + "r", H
		} else if 246 <= r && r < 259 {
			return seed + "s", H
		} else if 259 <= r && r < 320 {
			return seed + "t", H
		} else if 320 <= r && r < 326 {
			return seed + "u", H
		} else if 326 <= r && r < 333 {
			return seed + "v", H
		} else if 333 <= r && r < 338 {
			return seed + "w", H
		} else if 338 <= r && r < 339 {
			return seed + "x", H
		} else if 339 <= r && r < 342 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "cu" {
		r := rng.IntN(268)
		H := 3.1337687759458137
		if 0 <= r && r < 2 {
			return seed + "a", H
		} else if 2 <= r && r < 13 {
			return seed + "b", H
		} else if 13 <= r && r < 16 {
			return seed + "c", H
		} else if 16 <= r && r < 23 {
			return seed + "d", H
		} else if 23 <= r && r < 29 {
			return seed + "e", H
		} else if 29 <= r && r < 36 {
			return seed + "f", H
		} else if 36 <= r && r < 37 {
			return seed + "g", H
		} else if 37 <= r && r < 38 {
			return seed + "h", H
		} else if 38 <= r && r < 39 {
			return seed + "j", H
		} else if 39 <= r && r < 40 {
			return seed + "k", H
		} else if 40 <= r && r < 93 {
			return seed + "l", H
		} else if 93 <= r && r < 100 {
			return seed + "m", H
		} else if 100 <= r && r < 101 {
			return seed + "n", H
		} else if 101 <= r && r < 126 {
			return seed + "p", H
		} else if 126 <= r && r < 127 {
			return seed + "q", H
		} else if 127 <= r && r < 204 {
			return seed + "r", H
		} else if 204 <= r && r < 245 {
			return seed + "s", H
		} else if 245 <= r && r < 264 {
			return seed + "t", H
		} else if 264 <= r && r < 265 {
			return seed + "v", H
		} else if 265 <= r && r < 266 {
			return seed + "w", H
		} else if 266 <= r && r < 267 {
			return seed + "x", H
		} else if 267 <= r && r < 268 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fe" {
		r := rng.IntN(188)
		H := 3.643555367846094
		if 0 <= r && r < 4 {
			return seed + "a", H
		} else if 4 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 28 {
			return seed + "c", H
		} else if 28 <= r && r < 45 {
			return seed + "d", H
		} else if 45 <= r && r < 59 {
			return seed + "e", H
		} else if 59 <= r && r < 60 {
			return seed + "f", H
		} else if 60 <= r && r < 63 {
			return seed + "g", H
		} else if 63 <= r && r < 66 {
			return seed + "h", H
		} else if 66 <= r && r < 68 {
			return seed + "i", H
		} else if 68 <= r && r < 69 {
			return seed + "j", H
		} else if 69 <= r && r < 70 {
			return seed + "k", H
		} else if 70 <= r && r < 79 {
			return seed + "l", H
		} else if 79 <= r && r < 90 {
			return seed + "m", H
		} else if 90 <= r && r < 113 {
			return seed + "n", H
		} else if 113 <= r && r < 114 {
			return seed + "p", H
		} else if 114 <= r && r < 115 {
			return seed + "q", H
		} else if 115 <= r && r < 156 {
			return seed + "r", H
		} else if 156 <= r && r < 173 {
			return seed + "s", H
		} else if 173 <= r && r < 180 {
			return seed + "t", H
		} else if 180 <= r && r < 183 {
			return seed + "v", H
		} else if 183 <= r && r < 186 {
			return seed + "w", H
		} else if 186 <= r && r < 187 {
			return seed + "x", H
		} else if 187 <= r && r < 188 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zx" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qc" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rk" {
		r := rng.IntN(40)
		H := 2.7101275052218416
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 14 {
			return seed + "e", H
		} else if 14 <= r && r < 25 {
			return seed + "i", H
		} else if 25 <= r && r < 29 {
			return seed + "n", H
		} else if 29 <= r && r < 30 {
			return seed + "o", H
		} else if 30 <= r && r < 32 {
			return seed + "r", H
		} else if 32 <= r && r < 33 {
			return seed + "u", H
		} else if 33 <= r && r < 35 {
			return seed + "w", H
		} else if 35 <= r && r < 40 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "hw" {
		r := rng.IntN(12)
		H := 2.221251836004466
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 10 {
			return seed + "o", H
		} else if 10 <= r && r < 11 {
			return seed + "u", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mw" {
		r := rng.IntN(8)
		H := 2.4056390622295662
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 4 {
			return seed + "e", H
		} else if 4 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 6 {
			return seed + "o", H
		} else if 6 <= r && r < 7 {
			return seed + "u", H
		} else if 7 <= r && r < 8 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ke" {
		r := rng.IntN(246)
		H := 3.3743265196340344
		if 0 <= r && r < 5 {
			return seed + "b", H
		} else if 5 <= r && r < 6 {
			return seed + "c", H
		} else if 6 <= r && r < 55 {
			return seed + "d", H
		} else if 55 <= r && r < 65 {
			return seed + "e", H
		} else if 65 <= r && r < 66 {
			return seed + "f", H
		} else if 66 <= r && r < 69 {
			return seed + "g", H
		} else if 69 <= r && r < 72 {
			return seed + "h", H
		} else if 72 <= r && r < 73 {
			return seed + "j", H
		} else if 73 <= r && r < 74 {
			return seed + "k", H
		} else if 74 <= r && r < 87 {
			return seed + "l", H
		} else if 87 <= r && r < 90 {
			return seed + "m", H
		} else if 90 <= r && r < 121 {
			return seed + "n", H
		} else if 121 <= r && r < 123 {
			return seed + "o", H
		} else if 123 <= r && r < 128 {
			return seed + "p", H
		} else if 128 <= r && r < 129 {
			return seed + "q", H
		} else if 129 <= r && r < 190 {
			return seed + "r", H
		} else if 190 <= r && r < 201 {
			return seed + "s", H
		} else if 201 <= r && r < 228 {
			return seed + "t", H
		} else if 228 <= r && r < 229 {
			return seed + "v", H
		} else if 229 <= r && r < 240 {
			return seed + "w", H
		} else if 240 <= r && r < 241 {
			return seed + "x", H
		} else if 241 <= r && r < 245 {
			return seed + "y", H
		} else if 245 <= r && r < 246 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "yv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "rp" {
		r := rng.IntN(88)
		H := 2.976398251818003
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 26 {
			return seed + "e", H
		} else if 26 <= r && r < 34 {
			return seed + "h", H
		} else if 34 <= r && r < 45 {
			return seed + "i", H
		} else if 45 <= r && r < 61 {
			return seed + "l", H
		} else if 61 <= r && r < 63 {
			return seed + "n", H
		} else if 63 <= r && r < 80 {
			return seed + "o", H
		} else if 80 <= r && r < 84 {
			return seed + "r", H
		} else if 84 <= r && r < 85 {
			return seed + "u", H
		} else if 85 <= r && r < 88 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "jf" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wm" {
		r := rng.IntN(16)
		H := 1.9197367178034825
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 12 {
			return seed + "e", H
		} else if 12 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cc" {
		r := rng.IntN(48)
		H := 2.1502269866058414
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 11 {
			return seed + "i", H
		} else if 11 <= r && r < 15 {
			return seed + "l", H
		} else if 15 <= r && r < 22 {
			return seed + "o", H
		} else if 22 <= r && r < 47 {
			return seed + "u", H
		} else if 47 <= r && r < 48 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "iq" {
		r := rng.IntN(22)
		H := 1.3009376049438532
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 21 {
			return seed + "u", H
		} else if 21 <= r && r < 22 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ox" {
		r := rng.IntN(50)
		H := 2.153437628858099
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 3 {
			return seed + "c", H
		} else if 3 <= r && r < 8 {
			return seed + "e", H
		} else if 8 <= r && r < 10 {
			return seed + "f", H
		} else if 10 <= r && r < 37 {
			return seed + "i", H
		} else if 37 <= r && r < 39 {
			return seed + "l", H
		} else if 39 <= r && r < 40 {
			return seed + "o", H
		} else if 40 <= r && r < 41 {
			return seed + "u", H
		} else if 41 <= r && r < 50 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cw" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ad" {
		r := rng.IntN(304)
		H := 3.4228170682971975
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 17 {
			return seed + "b", H
		} else if 17 <= r && r < 21 {
			return seed + "c", H
		} else if 21 <= r && r < 57 {
			return seed + "d", H
		} else if 57 <= r && r < 138 {
			return seed + "e", H
		} else if 138 <= r && r < 144 {
			return seed + "f", H
		} else if 144 <= r && r < 148 {
			return seed + "g", H
		} else if 148 <= r && r < 207 {
			return seed + "i", H
		} else if 207 <= r && r < 233 {
			return seed + "l", H
		} else if 233 <= r && r < 237 {
			return seed + "m", H
		} else if 237 <= r && r < 243 {
			return seed + "n", H
		} else if 243 <= r && r < 254 {
			return seed + "o", H
		} else if 254 <= r && r < 260 {
			return seed + "p", H
		} else if 260 <= r && r < 266 {
			return seed + "r", H
		} else if 266 <= r && r < 278 {
			return seed + "s", H
		} else if 278 <= r && r < 283 {
			return seed + "u", H
		} else if 283 <= r && r < 285 {
			return seed + "v", H
		} else if 285 <= r && r < 291 {
			return seed + "w", H
		} else if 291 <= r && r < 304 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "on" {
		r := rng.IntN(610)
		H := 3.751377767467918
		if 0 <= r && r < 39 {
			return seed + "a", H
		} else if 39 <= r && r < 43 {
			return seed + "b", H
		} else if 43 <= r && r < 77 {
			return seed + "c", H
		} else if 77 <= r && r < 123 {
			return seed + "d", H
		} else if 123 <= r && r < 242 {
			return seed + "e", H
		} else if 242 <= r && r < 282 {
			return seed + "f", H
		} else if 282 <= r && r < 338 {
			return seed + "g", H
		} else if 338 <= r && r < 395 {
			return seed + "i", H
		} else if 395 <= r && r < 401 {
			return seed + "j", H
		} else if 401 <= r && r < 403 {
			return seed + "k", H
		} else if 403 <= r && r < 415 {
			return seed + "l", H
		} else if 415 <= r && r < 421 {
			return seed + "n", H
		} else if 421 <= r && r < 454 {
			return seed + "o", H
		} else if 454 <= r && r < 456 {
			return seed + "r", H
		} else if 456 <= r && r < 512 {
			return seed + "s", H
		} else if 512 <= r && r < 556 {
			return seed + "t", H
		} else if 556 <= r && r < 565 {
			return seed + "u", H
		} else if 565 <= r && r < 571 {
			return seed + "v", H
		} else if 571 <= r && r < 575 {
			return seed + "w", H
		} else if 575 <= r && r < 606 {
			return seed + "y", H
		} else if 606 <= r && r < 610 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "od" {
		r := rng.IntN(106)
		H := 2.742215346081511
		if 0 <= r && r < 3 {
			return seed + "a", H
		} else if 3 <= r && r < 5 {
			return seed + "d", H
		} else if 5 <= r && r < 32 {
			return seed + "e", H
		} else if 32 <= r && r < 38 {
			return seed + "g", H
		} else if 38 <= r && r < 63 {
			return seed + "i", H
		} else if 63 <= r && r < 67 {
			return seed + "l", H
		} else if 67 <= r && r < 74 {
			return seed + "o", H
		} else if 74 <= r && r < 89 {
			return seed + "u", H
		} else if 89 <= r && r < 106 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ob" {
		r := rng.IntN(144)
		H := 3.3250729565478627
		if 0 <= r && r < 11 {
			return seed + "a", H
		} else if 11 <= r && r < 33 {
			return seed + "b", H
		} else if 33 <= r && r < 35 {
			return seed + "c", H
		} else if 35 <= r && r < 54 {
			return seed + "e", H
		} else if 54 <= r && r < 77 {
			return seed + "i", H
		} else if 77 <= r && r < 79 {
			return seed + "j", H
		} else if 79 <= r && r < 91 {
			return seed + "l", H
		} else if 91 <= r && r < 93 {
			return seed + "n", H
		} else if 93 <= r && r < 98 {
			return seed + "o", H
		} else if 98 <= r && r < 126 {
			return seed + "s", H
		} else if 126 <= r && r < 134 {
			return seed + "t", H
		} else if 134 <= r && r < 137 {
			return seed + "u", H
		} else if 137 <= r && r < 139 {
			return seed + "v", H
		} else if 139 <= r && r < 141 {
			return seed + "w", H
		} else if 141 <= r && r < 144 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "cm" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ek" {
		r := rng.IntN(12)
		H := 2.8553885422075336
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 7 {
			return seed + "k", H
		} else if 7 <= r && r < 8 {
			return seed + "o", H
		} else if 8 <= r && r < 9 {
			return seed + "u", H
		} else if 9 <= r && r < 11 {
			return seed + "w", H
		} else if 11 <= r && r < 12 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ci" {
		r := rng.IntN(284)
		H := 3.682916826317365
		if 0 <= r && r < 20 {
			return seed + "a", H
		} else if 20 <= r && r < 23 {
			return seed + "b", H
		} else if 23 <= r && r < 24 {
			return seed + "c", H
		} else if 24 <= r && r < 49 {
			return seed + "d", H
		} else if 49 <= r && r < 65 {
			return seed + "e", H
		} else if 65 <= r && r < 76 {
			return seed + "f", H
		} else if 76 <= r && r < 77 {
			return seed + "g", H
		} else if 77 <= r && r < 78 {
			return seed + "h", H
		} else if 78 <= r && r < 79 {
			return seed + "j", H
		} else if 79 <= r && r < 80 {
			return seed + "k", H
		} else if 80 <= r && r < 95 {
			return seed + "l", H
		} else if 95 <= r && r < 102 {
			return seed + "m", H
		} else if 102 <= r && r < 165 {
			return seed + "n", H
		} else if 165 <= r && r < 181 {
			return seed + "o", H
		} else if 181 <= r && r < 188 {
			return seed + "p", H
		} else if 188 <= r && r < 189 {
			return seed + "q", H
		} else if 189 <= r && r < 202 {
			return seed + "r", H
		} else if 202 <= r && r < 225 {
			return seed + "s", H
		} else if 225 <= r && r < 268 {
			return seed + "t", H
		} else if 268 <= r && r < 270 {
			return seed + "u", H
		} else if 270 <= r && r < 279 {
			return seed + "v", H
		} else if 279 <= r && r < 280 {
			return seed + "w", H
		} else if 280 <= r && r < 281 {
			return seed + "x", H
		} else if 281 <= r && r < 284 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "fk" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "zq" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nd" {
		r := rng.IntN(520)
		H := 3.5341282044255693
		if 0 <= r && r < 43 {
			return seed + "a", H
		} else if 43 <= r && r < 61 {
			return seed + "b", H
		} else if 61 <= r && r < 71 {
			return seed + "c", H
		} else if 71 <= r && r < 73 {
			return seed + "d", H
		} else if 73 <= r && r < 224 {
			return seed + "e", H
		} else if 224 <= r && r < 236 {
			return seed + "f", H
		} else if 236 <= r && r < 240 {
			return seed + "g", H
		} else if 240 <= r && r < 242 {
			return seed + "h", H
		} else if 242 <= r && r < 299 {
			return seed + "i", H
		} else if 299 <= r && r < 301 {
			return seed + "k", H
		} else if 301 <= r && r < 355 {
			return seed + "l", H
		} else if 355 <= r && r < 367 {
			return seed + "m", H
		} else if 367 <= r && r < 377 {
			return seed + "n", H
		} else if 377 <= r && r < 416 {
			return seed + "o", H
		} else if 416 <= r && r < 428 {
			return seed + "p", H
		} else if 428 <= r && r < 450 {
			return seed + "r", H
		} else if 450 <= r && r < 476 {
			return seed + "s", H
		} else if 476 <= r && r < 497 {
			return seed + "u", H
		} else if 497 <= r && r < 509 {
			return seed + "w", H
		} else if 509 <= r && r < 520 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "tr" {
		r := rng.IntN(540)
		H := 2.4036181959394227
		if 0 <= r && r < 151 {
			return seed + "a", H
		} else if 151 <= r && r < 222 {
			return seed + "e", H
		} else if 222 <= r && r < 367 {
			return seed + "i", H
		} else if 367 <= r && r < 446 {
			return seed + "o", H
		} else if 446 <= r && r < 513 {
			return seed + "u", H
		} else if 513 <= r && r < 540 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "vi" {
		r := rng.IntN(398)
		H := 3.5285972892084723
		if 0 <= r && r < 30 {
			return seed + "a", H
		} else if 30 <= r && r < 33 {
			return seed + "b", H
		} else if 33 <= r && r < 52 {
			return seed + "c", H
		} else if 52 <= r && r < 87 {
			return seed + "d", H
		} else if 87 <= r && r < 107 {
			return seed + "e", H
		} else if 107 <= r && r < 108 {
			return seed + "f", H
		} else if 108 <= r && r < 115 {
			return seed + "g", H
		} else if 115 <= r && r < 116 {
			return seed + "h", H
		} else if 116 <= r && r < 117 {
			return seed + "j", H
		} else if 117 <= r && r < 118 {
			return seed + "k", H
		} else if 118 <= r && r < 137 {
			return seed + "l", H
		} else if 137 <= r && r < 138 {
			return seed + "m", H
		} else if 138 <= r && r < 229 {
			return seed + "n", H
		} else if 229 <= r && r < 255 {
			return seed + "o", H
		} else if 255 <= r && r < 258 {
			return seed + "p", H
		} else if 258 <= r && r < 259 {
			return seed + "q", H
		} else if 259 <= r && r < 272 {
			return seed + "r", H
		} else if 272 <= r && r < 337 {
			return seed + "s", H
		} else if 337 <= r && r < 370 {
			return seed + "t", H
		} else if 370 <= r && r < 393 {
			return seed + "v", H
		} else if 393 <= r && r < 394 {
			return seed + "w", H
		} else if 394 <= r && r < 397 {
			return seed + "x", H
		} else if 397 <= r && r < 398 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else if tok == "zd" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "dl" {
		r := rng.IntN(170)
		H := 1.993779496847317
		if 0 <= r && r < 9 {
			return seed + "a", H
		} else if 9 <= r && r < 84 {
			return seed + "e", H
		} else if 84 <= r && r < 115 {
			return seed + "i", H
		} else if 115 <= r && r < 126 {
			return seed + "o", H
		} else if 126 <= r && r < 127 {
			return seed + "u", H
		} else if 127 <= r && r < 170 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "mb" {
		r := rng.IntN(156)
		H := 2.527166228400974
		if 0 <= r && r < 17 {
			return seed + "a", H
		} else if 17 <= r && r < 46 {
			return seed + "e", H
		} else if 46 <= r && r < 63 {
			return seed + "i", H
		} else if 63 <= r && r < 123 {
			return seed + "l", H
		} else if 123 <= r && r < 125 {
			return seed + "n", H
		} else if 125 <= r && r < 142 {
			return seed + "o", H
		} else if 142 <= r && r < 146 {
			return seed + "r", H
		} else if 146 <= r && r < 155 {
			return seed + "u", H
		} else if 155 <= r && r < 156 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wv" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kz" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "qh" {
		r := rng.IntN(6)
		H := 2.584962500721156
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 4 {
			return seed + "o", H
		} else if 4 <= r && r < 5 {
			return seed + "u", H
		} else if 5 <= r && r < 6 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "bd" {
		r := rng.IntN(20)
		H := 2.421127473337187
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 19 {
			return seed + "u", H
		} else if 19 <= r && r < 20 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kl" {
		r := rng.IntN(68)
		H := 1.8852666081078917
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 42 {
			return seed + "e", H
		} else if 42 <= r && r < 55 {
			return seed + "i", H
		} else if 55 <= r && r < 58 {
			return seed + "o", H
		} else if 58 <= r && r < 59 {
			return seed + "u", H
		} else if 59 <= r && r < 68 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "kt" {
		r := rng.IntN(16)
		H := 2.4237949406953985
		if 0 <= r && r < 5 {
			return seed + "a", H
		} else if 5 <= r && r < 6 {
			return seed + "e", H
		} else if 6 <= r && r < 7 {
			return seed + "i", H
		} else if 7 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 14 {
			return seed + "r", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ys" {
		r := rng.IntN(42)
		H := 2.303679936757534
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 10 {
			return seed + "e", H
		} else if 10 <= r && r < 13 {
			return seed + "i", H
		} else if 13 <= r && r < 17 {
			return seed + "l", H
		} else if 17 <= r && r < 18 {
			return seed + "o", H
		} else if 18 <= r && r < 20 {
			return seed + "p", H
		} else if 20 <= r && r < 40 {
			return seed + "t", H
		} else if 40 <= r && r < 41 {
			return seed + "u", H
		} else if 41 <= r && r < 42 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "wp" {
		r := rng.IntN(14)
		H := 2.5566567074628237
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 5 {
			return seed + "i", H
		} else if 5 <= r && r < 9 {
			return seed + "l", H
		} else if 9 <= r && r < 12 {
			return seed + "o", H
		} else if 12 <= r && r < 13 {
			return seed + "u", H
		} else if 13 <= r && r < 14 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "yh" {
		r := rng.IntN(16)
		H := 1.6216407621868583
		if 0 <= r && r < 1 {
			return seed + "a", H
		} else if 1 <= r && r < 2 {
			return seed + "e", H
		} else if 2 <= r && r < 3 {
			return seed + "i", H
		} else if 3 <= r && r < 14 {
			return seed + "o", H
		} else if 14 <= r && r < 15 {
			return seed + "u", H
		} else if 15 <= r && r < 16 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "nm" {
		r := rng.IntN(42)
		H := 1.983490219992925
		if 0 <= r && r < 19 {
			return seed + "a", H
		} else if 19 <= r && r < 22 {
			return seed + "e", H
		} else if 22 <= r && r < 29 {
			return seed + "i", H
		} else if 29 <= r && r < 40 {
			return seed + "o", H
		} else if 40 <= r && r < 41 {
			return seed + "u", H
		} else if 41 <= r && r < 42 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "ng" {
		r := rng.IntN(296)
		H := 3.1881383083472175
		if 0 <= r && r < 13 {
			return seed + "a", H
		} else if 13 <= r && r < 17 {
			return seed + "b", H
		} else if 17 <= r && r < 19 {
			return seed + "d", H
		} else if 19 <= r && r < 98 {
			return seed + "e", H
		} else if 98 <= r && r < 100 {
			return seed + "f", H
		} else if 100 <= r && r < 102 {
			return seed + "h", H
		} else if 102 <= r && r < 125 {
			return seed + "i", H
		} else if 125 <= r && r < 191 {
			return seed + "l", H
		} else if 191 <= r && r < 193 {
			return seed + "m", H
		} else if 193 <= r && r < 197 {
			return seed + "n", H
		} else if 197 <= r && r < 218 {
			return seed + "o", H
		} else if 218 <= r && r < 248 {
			return seed + "r", H
		} else if 248 <= r && r < 260 {
			return seed + "s", H
		} else if 260 <= r && r < 264 {
			return seed + "t", H
		} else if 264 <= r && r < 285 {
			return seed + "u", H
		} else if 285 <= r && r < 287 {
			return seed + "w", H
		} else if 287 <= r && r < 296 {
			return seed + "y", H
		}
		panic("unexpected rand num")
	} else if tok == "he" {
		r := rng.IntN(474)
		H := 3.296399457568567
		if 0 <= r && r < 106 {
			return seed + "a", H
		} else if 106 <= r && r < 107 {
			return seed + "b", H
		} else if 107 <= r && r < 112 {
			return seed + "c", H
		} else if 112 <= r && r < 171 {
			return seed + "d", H
		} else if 171 <= r && r < 193 {
			return seed + "e", H
		} else if 193 <= r && r < 202 {
			return seed + "f", H
		} else if 202 <= r && r < 203 {
			return seed + "g", H
		} else if 203 <= r && r < 204 {
			return seed + "h", H
		} else if 204 <= r && r < 205 {
			return seed + "j", H
		} else if 205 <= r && r < 206 {
			return seed + "k", H
		} else if 206 <= r && r < 245 {
			return seed + "l", H
		} else if 245 <= r && r < 272 {
			return seed + "m", H
		} else if 272 <= r && r < 293 {
			return seed + "n", H
		} else if 293 <= r && r < 299 {
			return seed + "o", H
		} else if 299 <= r && r < 300 {
			return seed + "p", H
		} else if 300 <= r && r < 301 {
			return seed + "q", H
		} else if 301 <= r && r < 414 {
			return seed + "r", H
		} else if 414 <= r && r < 447 {
			return seed + "s", H
		} else if 447 <= r && r < 452 {
			return seed + "t", H
		} else if 452 <= r && r < 457 {
			return seed + "v", H
		} else if 457 <= r && r < 468 {
			return seed + "w", H
		} else if 468 <= r && r < 473 {
			return seed + "x", H
		} else if 473 <= r && r < 474 {
			return seed + "z", H
		}
		panic("unexpected rand num")
	} else {
		panic("unexpected token " + tok)
	}
}

func PickLength() (int, float64) {
	r := rng.IntN(4911)
	H := 2.240647562674072
	if 0 <= r && r < 82 {
		return 3, H
	} else if 82 <= r && r < 550 {
		return 4, H
	} else if 550 <= r && r < 1477 {
		return 5, H
	} else if 1477 <= r && r < 2850 {
		return 6, H
	} else if 2850 <= r && r < 4441 {
		return 7, H
	} else if 4441 <= r && r < 4911 {
		return 8, H
	}
	panic("unexpected rand num")
}
