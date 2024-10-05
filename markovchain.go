package cryptipass

import "strings"

// THIS FILE HAS BEEN DISTILLED FROM EFF's long word list, without their work this software would not exist.

// PickNext returns the next character to the current seed string based on specific rules.
// It evaluates the last two characters of the seed to decide on the next character, either
// through predefined cases (such as 'mh' becoming 'mho') or randomly selecting a new character
// with associated entropy.
//
//   - Parameters:
//     seed (string): The current string being used to generate the passphrase.
//
//   - Return values:
//     string: The updated string after appending the next character.
//     float64: The entropy contributed by the character selection process.
func PickNext(seed string) (string, float64) {
	L := min(len(seed), 2)
	tok := strings.ToLower(seed[len(seed)-L:])
retry:
	switch tok {
	case `mh`:
		return `o`, 0.0
	case `mv`:
		return `e`, 0.0
	case `bf`:
		return `l`, 0.0
	case `cq`:
		return `u`, 0.0
	case `ii`:
		return `n`, 0.0
	case `zm`:
		return `o`, 0.0
	case `pk`:
		return `i`, 0.0
	case `hd`:
		return `a`, 0.0
	case `kd`:
		return `r`, 0.0
	case `yz`:
		return `e`, 0.0
	case `bz`:
		return `e`, 0.0
	case `xb`:
		return `o`, 0.0
	case `uk`:
		return `e`, 0.0
	case `q`:
		return `u`, 0.0
	case `bj`:
		return `e`, 0.0
	case `yu`:
		return `m`, 0.0
	case `kf`:
		return `i`, 0.0
	case `ww`:
		return `o`, 0.0
	case `oq`:
		return `u`, 0.0
	case `wg`:
		return `i`, 0.0
	case `gz`:
		return `a`, 0.0
	case `oj`:
		return `e`, 0.0
	case `uj`:
		return `i`, 0.0
	case `wc`:
		return `a`, 0.0
	case `rx`:
		return `i`, 0.0
	case `bg`:
		return `r`, 0.0
	case `fn`:
		return `e`, 0.0
	case `sj`:
		return `o`, 0.0
	case `vr`:
		return `o`, 0.0
	case `zy`:
		return `m`, 0.0
	case `km`:
		return `a`, 0.0
	case `pg`:
		return `r`, 0.0
	case `wt`:
		return `i`, 0.0
	case `uv`:
		return `e`, 0.0
	case `gf`:
		return `u`, 0.0
	case `sq`:
		return `u`, 0.0
	case `xq`:
		return `u`, 0.0
	case `cs`:
		return `i`, 0.0
	case `dt`:
		return `h`, 0.0
	case `bc`:
		return `a`, 0.0
	case `xf`:
		return `o`, 0.0
	case `bv`:
		return `i`, 0.0
	case `sg`:
		return `r`, 0.0
	case `md`:
		return `r`, 0.0
	case `eq`:
		return `u`, 0.0
	case `wk`:
		return `w`, 0.0
	case `dv`:
		return `i`, 0.0
	case `xs`:
		return `e`, 0.0
	case `wu`:
		return `n`, 0.0
	case `pm`:
		return `e`, 0.0
	case `iu`:
		return `m`, 0.0
	case `dk`:
		return `i`, 0.0
	case `hs`:
		return `t`, 0.0
	case `hk`:
		return `i`, 0.0
	case `lh`:
		return `o`, 0.0
	case `lz`:
		return `o`, 0.0
	case `mw`:
		return `e`, 0.0
	case `gd`:
		return `o`, 0.0
	case `uf`:
		return `f`, 0.0
	case `xl`:
		return `i`, 0.0
	case `iq`:
		return `u`, 0.0
	case `aq`:
		return `u`, 0.0
	case `hh`:
		return `o`, 0.0
	case `yh`:
		return `o`, 0.0
	case `yk`:
		return `e`, 0.0
	case `gt`:
		return `h`, 0.0
	case `nq`:
		return `u`, 0.0
	case `hm`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `a`, H
		case r < 2:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ej`:
		H := 0.9709505944546686
		r := rng.IntN(5)
		switch {
		case r < 3:
			return `o`, H
		case r < 5:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `uz`:
		H := 0.5916727785823275
		r := rng.IntN(7)
		switch {
		case r < 6:
			return `z`, H
		case r < 7:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `hn`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `e`, H
		case r < 3:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `my`:
		H := 0.6500224216483541
		r := rng.IntN(6)
		switch {
		case r < 5:
			return `s`, H
		case r < 6:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `kh`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `o`, H
		case r < 3:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `hp`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `i`, H
		case r < 2:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `pw`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `kk`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `n`, H
		case r < 2:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `by`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `t`, H
		case r < 3:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `bh`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `e`, H
		case r < 2:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gy`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `r`, H
		case r < 2:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `xy`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `g`, H
		case r < 2:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ln`:
		H := 0.9182958340544896
		r := rng.IntN(6)
		switch {
		case r < 4:
			return `e`, H
		case r < 6:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lw`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `uo`:
		H := 0.9910760598382222
		r := rng.IntN(9)
		switch {
		case r < 5:
			return `u`, H
		case r < 9:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `gm`:
		H := 1.0
		r := rng.IntN(8)
		switch {
		case r < 4:
			return `a`, H
		case r < 8:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `pn`:
		H := 0.9940302114769566
		r := rng.IntN(11)
		switch {
		case r < 6:
			return `o`, H
		case r < 11:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `tz`:
		H := 0.9709505944546686
		r := rng.IntN(5)
		switch {
		case r < 3:
			return `y`, H
		case r < 5:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `wl`:
		H := 0.9709505944546686
		r := rng.IntN(10)
		switch {
		case r < 6:
			return `e`, H
		case r < 10:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `sd`:
		H := 0.8112781244591328
		r := rng.IntN(4)
		switch {
		case r < 3:
			return `a`, H
		case r < 4:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `sr`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `u`, H
		case r < 2:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `xh`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lk`:
		H := 0.8112781244591328
		r := rng.IntN(4)
		switch {
		case r < 3:
			return `a`, H
		case r < 4:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `x`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `e`, H
		case r < 2:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `bn`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `o`, H
		case r < 3:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `dn`:
		H := 0.41381685030363374
		r := rng.IntN(12)
		switch {
		case r < 11:
			return `e`, H
		case r < 12:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ae`:
		H := 0.6500224216483541
		r := rng.IntN(6)
		switch {
		case r < 5:
			return `r`, H
		case r < 6:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ux`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `u`, H
		case r < 3:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `fb`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `a`, H
		case r < 2:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `hw`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return `o`, H
		case r < 3:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `wm`:
		H := 0.7219280948873623
		r := rng.IntN(5)
		switch {
		case r < 4:
			return `a`, H
		case r < 5:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `pd`:
		H := 0.8112781244591328
		r := rng.IntN(4)
		switch {
		case r < 3:
			return `o`, H
		case r < 4:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `td`:
		H := 0.7219280948873623
		r := rng.IntN(5)
		switch {
		case r < 4:
			return `o`, H
		case r < 5:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `iw`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `i`, H
		case r < 2:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ih`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return `u`, H
		case r < 2:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `rj`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `u`, H
		case r < 3:
			return `e`, H
		case r < 4:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `kr`:
		H := 1.3709505944546687
		r := rng.IntN(5)
		switch {
		case r < 3:
			return `o`, H
		case r < 4:
			return `y`, H
		case r < 5:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ko`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `a`, H
		case r < 2:
			return `s`, H
		case r < 3:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `mc`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `h`, H
		case r < 2:
			return `e`, H
		case r < 3:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `kw`:
		H := 1.2987949406953985
		r := rng.IntN(8)
		switch {
		case r < 5:
			return `a`, H
		case r < 7:
			return `o`, H
		case r < 8:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `oh`:
		H := 1.3709505944546687
		r := rng.IntN(5)
		switch {
		case r < 3:
			return `e`, H
		case r < 4:
			return `n`, H
		case r < 5:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `bm`:
		H := 1.4591479170272448
		r := rng.IntN(6)
		switch {
		case r < 3:
			return `e`, H
		case r < 5:
			return `i`, H
		case r < 6:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `hb`:
		H := 1.3485878960124222
		r := rng.IntN(11)
		switch {
		case r < 5:
			return `a`, H
		case r < 10:
			return `o`, H
		case r < 11:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `gb`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `o`, H
		case r < 3:
			return `a`, H
		case r < 4:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `lv`:
		H := 1.2467052127325735
		r := rng.IntN(21)
		switch {
		case r < 14:
			return `e`, H
		case r < 18:
			return `a`, H
		case r < 21:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `hc`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `h`, H
		case r < 4:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `aj`:
		H := 1.584962500721156
		r := rng.IntN(6)
		switch {
		case r < 2:
			return `a`, H
		case r < 4:
			return `e`, H
		case r < 6:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `mr`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `o`, H
		case r < 3:
			return `a`, H
		case r < 4:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `nj`:
		H := 1.280672129520887
		r := rng.IntN(12)
		switch {
		case r < 7:
			return `o`, H
		case r < 11:
			return `u`, H
		case r < 12:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `bp`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `l`, H
		case r < 4:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `mn`:
		H := 1.4056390622295665
		r := rng.IntN(8)
		switch {
		case r < 4:
			return `e`, H
		case r < 7:
			return `i`, H
		case r < 8:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `py`:
		H := 1.3709505944546687
		r := rng.IntN(5)
		switch {
		case r < 3:
			return `r`, H
		case r < 4:
			return `g`, H
		case r < 5:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `yd`:
		H := 1.1887218755408673
		r := rng.IntN(12)
		switch {
		case r < 8:
			return `r`, H
		case r < 11:
			return `a`, H
		case r < 12:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `tg`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `e`, H
		case r < 2:
			return `r`, H
		case r < 3:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `pb`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `e`, H
		case r < 3:
			return `a`, H
		case r < 4:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `mf`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `u`, H
		case r < 3:
			return `y`, H
		case r < 4:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ao`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `s`, H
		case r < 2:
			return `r`, H
		case r < 3:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `kp`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `a`, H
		case r < 2:
			return `e`, H
		case r < 3:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gp`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `i`, H
		case r < 2:
			return `l`, H
		case r < 3:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `zl`:
		H := 1.1401156785146092
		r := rng.IntN(13)
		switch {
		case r < 9:
			return `e`, H
		case r < 12:
			return `i`, H
		case r < 13:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `pc`:
		H := 1.2516291673878228
		r := rng.IntN(6)
		switch {
		case r < 4:
			return `o`, H
		case r < 5:
			return `h`, H
		case r < 6:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `pf`:
		H := 1.5219280948873621
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `i`, H
		case r < 4:
			return `u`, H
		case r < 5:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ek`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `w`, H
		case r < 2:
			return `i`, H
		case r < 3:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ah`:
		H := 1.3787834934861756
		r := rng.IntN(7)
		switch {
		case r < 4:
			return `o`, H
		case r < 6:
			return `e`, H
		case r < 7:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `xu`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `r`, H
		case r < 3:
			return `a`, H
		case r < 4:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `bw`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return `a`, H
		case r < 2:
			return `e`, H
		case r < 3:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `kt`:
		H := 1.5219280948873621
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `a`, H
		case r < 4:
			return `o`, H
		case r < 5:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `wp`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return `l`, H
		case r < 3:
			return `i`, H
		case r < 4:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `dm`:
		H := 1.3921472236645345
		r := rng.IntN(9)
		switch {
		case r < 4:
			return `a`, H
		case r < 8:
			return `i`, H
		case r < 9:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `sb`:
		H := 1.9182958340544893
		r := rng.IntN(6)
		switch {
		case r < 2:
			return `a`, H
		case r < 4:
			return `e`, H
		case r < 5:
			return `u`, H
		case r < 6:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `cy`:
		H := 1.3520301017579528
		r := rng.IntN(13)
		switch {
		case r < 9:
			return `c`, H
		case r < 11:
			return `t`, H
		case r < 12:
			return `l`, H
		case r < 13:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `lr`:
		H := 1.9182958340544893
		r := rng.IntN(6)
		switch {
		case r < 2:
			return `y`, H
		case r < 4:
			return `i`, H
		case r < 5:
			return `u`, H
		case r < 6:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `yf`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `u`, H
		case r < 3:
			return `i`, H
		case r < 4:
			return `l`, H
		case r < 5:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `df`:
		H := 1.7841591278514217
		r := rng.IntN(12)
		switch {
		case r < 5:
			return `i`, H
		case r < 9:
			return `u`, H
		case r < 11:
			return `a`, H
		case r < 12:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `wb`:
		H := 1.9502120649147472
		r := rng.IntN(7)
		switch {
		case r < 2:
			return `a`, H
		case r < 4:
			return `i`, H
		case r < 6:
			return `o`, H
		case r < 7:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `yg`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `r`, H
		case r < 3:
			return `l`, H
		case r < 4:
			return `e`, H
		case r < 5:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gh`:
		H := 0.44363159167676336
		r := rng.IntN(47)
		switch {
		case r < 44:
			return `t`, H
		case r < 45:
			return `a`, H
		case r < 46:
			return `y`, H
		case r < 47:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `z`:
		H := 1.7702834614222898
		r := rng.IntN(22)
		switch {
		case r < 8:
			return `o`, H
		case r < 15:
			return `e`, H
		case r < 21:
			return `i`, H
		case r < 22:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `yi`:
		H := 0.719556365373903
		r := rng.IntN(25)
		switch {
		case r < 22:
			return `n`, H
		case r < 23:
			return `d`, H
		case r < 24:
			return `e`, H
		case r < 25:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `zz`:
		H := 1.478897902987479
		r := rng.IntN(20)
		switch {
		case r < 13:
			return `l`, H
		case r < 16:
			return `i`, H
		case r < 18:
			return `a`, H
		case r < 20:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `gn`:
		H := 1.9584661817581863
		r := rng.IntN(37)
		switch {
		case r < 13:
			return `e`, H
		case r < 22:
			return `a`, H
		case r < 30:
			return `i`, H
		case r < 37:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ak`:
		H := 1.0592982590819013
		r := rng.IntN(55)
		switch {
		case r < 41:
			return `e`, H
		case r < 52:
			return `i`, H
		case r < 54:
			return `y`, H
		case r < 55:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `kn`:
		H := 1.8412501702815551
		r := rng.IntN(18)
		switch {
		case r < 7:
			return `e`, H
		case r < 13:
			return `o`, H
		case r < 16:
			return `i`, H
		case r < 18:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `yc`:
		H := 1.4838317068446891
		r := rng.IntN(14)
		switch {
		case r < 9:
			return `l`, H
		case r < 11:
			return `h`, H
		case r < 13:
			return `a`, H
		case r < 14:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `m`:
		H := 1.596319592368159
		r := rng.IntN(295)
		switch {
		case r < 130:
			return `a`, H
		case r < 237:
			return `o`, H
		case r < 290:
			return `u`, H
		case r < 295:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `yw`:
		H := 1.6644977792004614
		r := rng.IntN(7)
		switch {
		case r < 4:
			return `a`, H
		case r < 5:
			return `h`, H
		case r < 6:
			return `i`, H
		case r < 7:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gs`:
		H := 1.811278124459133
		r := rng.IntN(8)
		switch {
		case r < 3:
			return `h`, H
		case r < 6:
			return `t`, H
		case r < 7:
			return `a`, H
		case r < 8:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `qu`:
		H := 1.7428655983817265
		r := rng.IntN(99)
		switch {
		case r < 37:
			return `i`, H
		case r < 71:
			return `a`, H
		case r < 95:
			return `e`, H
		case r < 99:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `xo`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `r`, H
		case r < 3:
			return `n`, H
		case r < 4:
			return `d`, H
		case r < 5:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `h`:
		H := 1.7259650053109339
		r := rng.IntN(249)
		switch {
		case r < 122:
			return `a`, H
		case r < 184:
			return `e`, H
		case r < 232:
			return `u`, H
		case r < 249:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `kb`:
		H := 1.6644977792004614
		r := rng.IntN(7)
		switch {
		case r < 4:
			return `o`, H
		case r < 5:
			return `a`, H
		case r < 6:
			return `i`, H
		case r < 7:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `tm`:
		H := 1.9056390622295665
		r := rng.IntN(8)
		switch {
		case r < 3:
			return `e`, H
		case r < 5:
			return `a`, H
		case r < 7:
			return `o`, H
		case r < 8:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lg`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `e`, H
		case r < 3:
			return `i`, H
		case r < 4:
			return `u`, H
		case r < 5:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `yl`:
		H := 1.4516607643577668
		r := rng.IntN(17)
		switch {
		case r < 11:
			return `i`, H
		case r < 14:
			return `e`, H
		case r < 16:
			return `o`, H
		case r < 17:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ik`:
		H := 1.0174051189440556
		r := rng.IntN(34)
		switch {
		case r < 27:
			return `e`, H
		case r < 31:
			return `i`, H
		case r < 33:
			return `a`, H
		case r < 34:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `yt`:
		H := 1.9362600275315274
		r := rng.IntN(11)
		switch {
		case r < 4:
			return `h`, H
		case r < 7:
			return `i`, H
		case r < 9:
			return `e`, H
		case r < 11:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `wd`:
		H := 1.8365916681089791
		r := rng.IntN(9)
		switch {
		case r < 4:
			return `e`, H
		case r < 6:
			return `l`, H
		case r < 8:
			return `r`, H
		case r < 9:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ls`:
		H := 1.8828560636920488
		r := rng.IntN(16)
		switch {
		case r < 6:
			return `i`, H
		case r < 11:
			return `e`, H
		case r < 14:
			return `a`, H
		case r < 16:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `dh`:
		H := 1.9182958340544893
		r := rng.IntN(6)
		switch {
		case r < 2:
			return `i`, H
		case r < 4:
			return `e`, H
		case r < 5:
			return `a`, H
		case r < 6:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `yr`:
		H := 1.8631205685666312
		r := rng.IntN(14)
		switch {
		case r < 5:
			return `i`, H
		case r < 10:
			return `o`, H
		case r < 12:
			return `a`, H
		case r < 14:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rh`:
		H := 1.676737030052133
		r := rng.IntN(11)
		switch {
		case r < 5:
			return `e`, H
		case r < 9:
			return `a`, H
		case r < 10:
			return `u`, H
		case r < 11:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `nz`:
		H := 1.8423709931771088
		r := rng.IntN(7)
		switch {
		case r < 3:
			return `i`, H
		case r < 5:
			return `y`, H
		case r < 6:
			return `a`, H
		case r < 7:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `wn`:
		H := 2.0
		r := rng.IntN(4)
		switch {
		case r < 1:
			return `n`, H
		case r < 2:
			return `y`, H
		case r < 3:
			return `i`, H
		case r < 4:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `wf`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `i`, H
		case r < 3:
			return `a`, H
		case r < 4:
			return `l`, H
		case r < 5:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `tf`:
		H := 1.6589797521242051
		r := rng.IntN(17)
		switch {
		case r < 9:
			return `u`, H
		case r < 13:
			return `i`, H
		case r < 16:
			return `o`, H
		case r < 17:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `rw`:
		H := 1.9056390622295665
		r := rng.IntN(8)
		switch {
		case r < 3:
			return `i`, H
		case r < 5:
			return `a`, H
		case r < 7:
			return `e`, H
		case r < 8:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `gw`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `h`, H
		case r < 4:
			return `e`, H
		case r < 5:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `bd`:
		H := 1.8423709931771088
		r := rng.IntN(7)
		switch {
		case r < 3:
			return `o`, H
		case r < 5:
			return `u`, H
		case r < 6:
			return `i`, H
		case r < 7:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `v`:
		H := 1.919393986516576
		r := rng.IntN(132)
		switch {
		case r < 47:
			return `i`, H
		case r < 81:
			return `a`, H
		case r < 115:
			return `e`, H
		case r < 132:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ku`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return `p`, H
		case r < 3:
			return `n`, H
		case r < 4:
			return `d`, H
		case r < 5:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `nm`:
		H := 1.6758220575766962
		r := rng.IntN(18)
		switch {
		case r < 9:
			return `a`, H
		case r < 14:
			return `o`, H
		case r < 17:
			return `i`, H
		case r < 18:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `tw`:
		H := 1.8832406052365576
		r := rng.IntN(39)
		switch {
		case r < 18:
			return `i`, H
		case r < 28:
			return `e`, H
		case r < 35:
			return `a`, H
		case r < 38:
			return `o`, H
		case r < 39:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `tl`:
		H := 1.7570627627948316
		r := rng.IntN(95)
		switch {
		case r < 39:
			return `e`, H
		case r < 77:
			return `y`, H
		case r < 88:
			return `i`, H
		case r < 93:
			return `a`, H
		case r < 95:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `br`:
		H := 2.2142644095851374
		r := rng.IntN(87)
		switch {
		case r < 26:
			return `o`, H
		case r < 49:
			return `i`, H
		case r < 67:
			return `e`, H
		case r < 78:
			return `a`, H
		case r < 87:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `nh`:
		H := 2.2068570640942182
		r := rng.IntN(23)
		switch {
		case r < 6:
			return `a`, H
		case r < 12:
			return `e`, H
		case r < 18:
			return `o`, H
		case r < 21:
			return `i`, H
		case r < 23:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `y`:
		H := 2.0793752444772653
		r := rng.IntN(27)
		switch {
		case r < 10:
			return `e`, H
		case r < 16:
			return `a`, H
		case r < 22:
			return `o`, H
		case r < 26:
			return `i`, H
		case r < 27:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `tp`:
		H := 2.235926350629033
		r := rng.IntN(7)
		switch {
		case r < 2:
			return `a`, H
		case r < 4:
			return `o`, H
		case r < 5:
			return `l`, H
		case r < 6:
			return `r`, H
		case r < 7:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rv`:
		H := 1.915924784727185
		r := rng.IntN(34)
		switch {
		case r < 14:
			return `i`, H
		case r < 26:
			return `e`, H
		case r < 29:
			return `a`, H
		case r < 32:
			return `y`, H
		case r < 34:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `tc`:
		H := 0.9954758876481817
		r := rng.IntN(59)
		switch {
		case r < 48:
			return `h`, H
		case r < 54:
			return `a`, H
		case r < 57:
			return `o`, H
		case r < 58:
			return `l`, H
		case r < 59:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ix`:
		H := 2.054585169337799
		r := rng.IntN(12)
		switch {
		case r < 5:
			return `t`, H
		case r < 8:
			return `e`, H
		case r < 10:
			return `a`, H
		case r < 11:
			return `f`, H
		case r < 12:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `ji`:
		H := 2.113283334294875
		r := rng.IntN(9)
		switch {
		case r < 3:
			return `n`, H
		case r < 6:
			return `t`, H
		case r < 7:
			return `f`, H
		case r < 8:
			return `g`, H
		case r < 9:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `j`:
		H := 2.077766251414667
		r := rng.IntN(96)
		switch {
		case r < 36:
			return `u`, H
		case r < 62:
			return `a`, H
		case r < 81:
			return `o`, H
		case r < 89:
			return `i`, H
		case r < 96:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `nv`:
		H := 1.962354483657865
		r := rng.IntN(22)
		switch {
		case r < 9:
			return `i`, H
		case r < 16:
			return `e`, H
		case r < 19:
			return `a`, H
		case r < 21:
			return `o`, H
		case r < 22:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `xe`:
		H := 2.0419460322060456
		r := rng.IntN(15)
		switch {
		case r < 5:
			return `d`, H
		case r < 10:
			return `r`, H
		case r < 13:
			return `m`, H
		case r < 14:
			return `n`, H
		case r < 15:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `sn`:
		H := 1.9664658234492332
		r := rng.IntN(48)
		switch {
		case r < 24:
			return `o`, H
		case r < 33:
			return `a`, H
		case r < 39:
			return `i`, H
		case r < 44:
			return `u`, H
		case r < 48:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `dg`:
		H := 1.3957137264520418
		r := rng.IntN(33)
		switch {
		case r < 23:
			return `e`, H
		case r < 28:
			return `i`, H
		case r < 31:
			return `y`, H
		case r < 32:
			return `r`, H
		case r < 33:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `dp`:
		H := 2.197159723424149
		r := rng.IntN(9)
		switch {
		case r < 3:
			return `i`, H
		case r < 5:
			return `a`, H
		case r < 7:
			return `o`, H
		case r < 8:
			return `h`, H
		case r < 9:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `b`:
		H := 2.2296839672243163
		r := rng.IntN(362)
		switch {
		case r < 126:
			return `a`, H
		case r < 197:
			return `o`, H
		case r < 254:
			return `r`, H
		case r < 309:
			return `u`, H
		case r < 362:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `dc`:
		H := 2.038159681645953
		r := rng.IntN(13)
		switch {
		case r < 5:
			return `a`, H
		case r < 9:
			return `o`, H
		case r < 11:
			return `l`, H
		case r < 12:
			return `r`, H
		case r < 13:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lm`:
		H := 2.197159723424149
		r := rng.IntN(9)
		switch {
		case r < 3:
			return `a`, H
		case r < 5:
			return `i`, H
		case r < 7:
			return `o`, H
		case r < 8:
			return `y`, H
		case r < 9:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ok`:
		H := 1.546361197390146
		r := rng.IntN(36)
		switch {
		case r < 23:
			return `e`, H
		case r < 30:
			return `i`, H
		case r < 33:
			return `y`, H
		case r < 35:
			return `a`, H
		case r < 36:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `n`:
		H := 1.9834964157590933
		r := rng.IntN(97)
		switch {
		case r < 33:
			return `e`, H
		case r < 59:
			return `u`, H
		case r < 84:
			return `a`, H
		case r < 96:
			return `i`, H
		case r < 97:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `yn`:
		H := 2.1867043459100244
		r := rng.IntN(11)
		switch {
		case r < 4:
			return `a`, H
		case r < 6:
			return `e`, H
		case r < 8:
			return `t`, H
		case r < 10:
			return `o`, H
		case r < 11:
			return `d`, H
		default:
			panic("unexpected rand num")
		}
	case `hr`:
		H := 2.2129220962815555
		r := rng.IntN(43)
		switch {
		case r < 14:
			return `o`, H
		case r < 24:
			return `i`, H
		case r < 32:
			return `e`, H
		case r < 39:
			return `a`, H
		case r < 43:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ov`:
		H := 1.1045475199265447
		r := rng.IntN(160)
		switch {
		case r < 125:
			return `e`, H
		case r < 141:
			return `i`, H
		case r < 155:
			return `a`, H
		case r < 158:
			return `o`, H
		case r < 160:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ax`:
		H := 2.1280852788913944
		r := rng.IntN(7)
		switch {
		case r < 3:
			return `i`, H
		case r < 4:
			return `s`, H
		case r < 5:
			return `a`, H
		case r < 6:
			return `e`, H
		case r < 7:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `wh`:
		H := 2.131346433249389
		r := rng.IntN(25)
		switch {
		case r < 7:
			return `i`, H
		case r < 14:
			return `e`, H
		case r < 20:
			return `o`, H
		case r < 24:
			return `a`, H
		case r < 25:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ws`:
		H := 2.2516291673878226
		r := rng.IntN(6)
		switch {
		case r < 2:
			return `i`, H
		case r < 3:
			return `h`, H
		case r < 4:
			return `e`, H
		case r < 5:
			return `t`, H
		case r < 6:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `oz`:
		H := 2.0550365325772657
		r := rng.IntN(16)
		switch {
		case r < 6:
			return `e`, H
		case r < 11:
			return `y`, H
		case r < 13:
			return `i`, H
		case r < 15:
			return `o`, H
		case r < 16:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `ym`:
		H := 2.104292989590925
		r := rng.IntN(19)
		switch {
		case r < 6:
			return `e`, H
		case r < 11:
			return `a`, H
		case r < 16:
			return `p`, H
		case r < 18:
			return `o`, H
		case r < 19:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `lf`:
		H := 2.075869296697727
		r := rng.IntN(13)
		switch {
		case r < 5:
			return `i`, H
		case r < 8:
			return `a`, H
		case r < 11:
			return `u`, H
		case r < 12:
			return `r`, H
		case r < 13:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `dw`:
		H := 2.1645691489449588
		r := rng.IntN(19)
		switch {
		case r < 5:
			return `a`, H
		case r < 10:
			return `e`, H
		case r < 15:
			return `o`, H
		case r < 18:
			return `i`, H
		case r < 19:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `eh`:
		H := 2.2577560641285177
		r := rng.IntN(13)
		switch {
		case r < 4:
			return `e`, H
		case r < 7:
			return `i`, H
		case r < 9:
			return `y`, H
		case r < 11:
			return `a`, H
		case r < 13:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `nw`:
		H := 2.17675266977692
		r := rng.IntN(27)
		switch {
		case r < 8:
			return `a`, H
		case r < 16:
			return `o`, H
		case r < 21:
			return `i`, H
		case r < 25:
			return `e`, H
		case r < 27:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `sw`:
		H := 2.046381112945884
		r := rng.IntN(42)
		switch {
		case r < 16:
			return `i`, H
		case r < 25:
			return `e`, H
		case r < 33:
			return `a`, H
		case r < 41:
			return `o`, H
		case r < 42:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `db`:
		H := 1.8322488896002307
		r := rng.IntN(14)
		switch {
		case r < 7:
			return `a`, H
		case r < 11:
			return `o`, H
		case r < 12:
			return `l`, H
		case r < 13:
			return `r`, H
		case r < 14:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `l`:
		H := 2.0699571885736554
		r := rng.IntN(194)
		switch {
		case r < 63:
			return `a`, H
		case r < 119:
			return `i`, H
		case r < 157:
			return `u`, H
		case r < 189:
			return `e`, H
		case r < 194:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `nr`:
		H := 2.1468032139556708
		r := rng.IntN(27)
		switch {
		case r < 11:
			return `e`, H
		case r < 16:
			return `i`, H
		case r < 20:
			return `a`, H
		case r < 24:
			return `o`, H
		case r < 27:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `za`:
		H := 1.896240625180289
		r := rng.IntN(12)
		switch {
		case r < 6:
			return `r`, H
		case r < 9:
			return `b`, H
		case r < 10:
			return `g`, H
		case r < 11:
			return `c`, H
		case r < 12:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `iz`:
		H := 1.3054506823200551
		r := rng.IntN(69)
		switch {
		case r < 50:
			return `e`, H
		case r < 60:
			return `z`, H
		case r < 64:
			return `a`, H
		case r < 68:
			return `i`, H
		case r < 69:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `wr`:
		H := 2.1229853836154926
		r := rng.IntN(29)
		switch {
		case r < 11:
			return `i`, H
		case r < 18:
			return `o`, H
		case r < 23:
			return `e`, H
		case r < 27:
			return `a`, H
		case r < 29:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `iv`:
		H := 1.5989577717916381
		r := rng.IntN(161)
		switch {
		case r < 91:
			return `e`, H
		case r < 131:
			return `i`, H
		case r < 153:
			return `a`, H
		case r < 160:
			return `o`, H
		case r < 161:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `hy`:
		H := 1.6444492199495462
		r := rng.IntN(21)
		switch {
		case r < 11:
			return `p`, H
		case r < 18:
			return `d`, H
		case r < 19:
			return `s`, H
		case r < 20:
			return `m`, H
		case r < 21:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `tn`:
		H := 1.3862740938959681
		r := rng.IntN(18)
		switch {
		case r < 13:
			return `e`, H
		case r < 15:
			return `a`, H
		case r < 16:
			return `i`, H
		case r < 17:
			return `u`, H
		case r < 18:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `cl`:
		H := 2.273319510050957
		r := rng.IntN(123)
		switch {
		case r < 36:
			return `a`, H
		case r < 64:
			return `e`, H
		case r < 84:
			return `i`, H
		case r < 104:
			return `u`, H
		case r < 123:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `av`:
		H := 1.9476716502007156
		r := rng.IntN(96)
		switch {
		case r < 41:
			return `e`, H
		case r < 69:
			return `i`, H
		case r < 82:
			return `a`, H
		case r < 93:
			return `o`, H
		case r < 96:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `w`:
		H := 2.220342743810504
		r := rng.IntN(155)
		switch {
		case r < 47:
			return `i`, H
		case r < 89:
			return `a`, H
		case r < 117:
			return `o`, H
		case r < 137:
			return `r`, H
		case r < 155:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `dl`:
		H := 1.9282661532133314
		r := rng.IntN(82)
		switch {
		case r < 37:
			return `e`, H
		case r < 58:
			return `y`, H
		case r < 73:
			return `i`, H
		case r < 78:
			return `o`, H
		case r < 82:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `hl`:
		H := 1.9986525468781666
		r := rng.IntN(17)
		switch {
		case r < 8:
			return `y`, H
		case r < 11:
			return `i`, H
		case r < 14:
			return `o`, H
		case r < 16:
			return `e`, H
		case r < 17:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `kl`:
		H := 1.710053567474341
		r := rng.IntN(31)
		switch {
		case r < 18:
			return `e`, H
		case r < 24:
			return `i`, H
		case r < 28:
			return `y`, H
		case r < 30:
			return `a`, H
		case r < 31:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ys`:
		H := 1.7688538376160001
		r := rng.IntN(18)
		switch {
		case r < 10:
			return `t`, H
		case r < 14:
			return `e`, H
		case r < 16:
			return `l`, H
		case r < 17:
			return `p`, H
		case r < 18:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `ml`:
		H := 1.987773371487984
		r := rng.IntN(13)
		switch {
		case r < 6:
			return `e`, H
		case r < 9:
			return `y`, H
		case r < 11:
			return `i`, H
		case r < 12:
			return `a`, H
		case r < 13:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `sm`:
		H := 2.1119613706025886
		r := rng.IntN(53)
		switch {
		case r < 18:
			return `a`, H
		case r < 32:
			return `o`, H
		case r < 44:
			return `i`, H
		case r < 50:
			return `u`, H
		case r < 53:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `r`:
		H := 1.5668118547237875
		r := rng.IntN(513)
		switch {
		case r < 345:
			return `e`, H
		case r < 401:
			return `a`, H
		case r < 445:
			return `i`, H
		case r < 484:
			return `o`, H
		case r < 512:
			return `u`, H
		case r < 513:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `cr`:
		H := 2.394320692780727
		r := rng.IntN(168)
		switch {
		case r < 42:
			return `a`, H
		case r < 83:
			return `e`, H
		case r < 112:
			return `i`, H
		case r < 140:
			return `u`, H
		case r < 164:
			return `o`, H
		case r < 168:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `lc`:
		H := 2.3709505944546683
		r := rng.IntN(10)
		switch {
		case r < 3:
			return `u`, H
		case r < 6:
			return `o`, H
		case r < 7:
			return `h`, H
		case r < 8:
			return `a`, H
		case r < 9:
			return `i`, H
		case r < 10:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `gl`:
		H := 2.4463501798243668
		r := rng.IntN(112)
		switch {
		case r < 35:
			return `e`, H
		case r < 54:
			return `a`, H
		case r < 72:
			return `i`, H
		case r < 89:
			return `y`, H
		case r < 105:
			return `o`, H
		case r < 112:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sy`:
		H := 2.2584598927441624
		r := rng.IntN(18)
		switch {
		case r < 7:
			return `n`, H
		case r < 11:
			return `m`, H
		case r < 14:
			return `s`, H
		case r < 16:
			return `c`, H
		case r < 17:
			return `r`, H
		case r < 18:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `bl`:
		H := 1.8303332451927128
		r := rng.IntN(293)
		switch {
		case r < 177:
			return `e`, H
		case r < 213:
			return `i`, H
		case r < 248:
			return `y`, H
		case r < 267:
			return `a`, H
		case r < 282:
			return `o`, H
		case r < 293:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `zo`:
		H := 2.1920058354921066
		r := rng.IntN(13)
		switch {
		case r < 5:
			return `o`, H
		case r < 9:
			return `n`, H
		case r < 10:
			return `a`, H
		case r < 11:
			return `i`, H
		case r < 12:
			return `d`, H
		case r < 13:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `pl`:
		H := 2.3540218435644924
		r := rng.IntN(178)
		switch {
		case r < 63:
			return `a`, H
		case r < 99:
			return `e`, H
		case r < 130:
			return `i`, H
		case r < 153:
			return `o`, H
		case r < 167:
			return `y`, H
		case r < 178:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `fr`:
		H := 2.200881953295572
		r := rng.IntN(98)
		switch {
		case r < 33:
			return `e`, H
		case r < 56:
			return `a`, H
		case r < 76:
			return `o`, H
		case r < 92:
			return `i`, H
		case r < 97:
			return `u`, H
		case r < 98:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ez`:
		H := 2.064042639445697
		r := rng.IntN(14)
		switch {
		case r < 7:
			return `e`, H
		case r < 10:
			return `i`, H
		case r < 11:
			return `y`, H
		case r < 12:
			return `a`, H
		case r < 13:
			return `o`, H
		case r < 14:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `lb`:
		H := 2.5216406363433186
		r := rng.IntN(7)
		switch {
		case r < 2:
			return `i`, H
		case r < 3:
			return `a`, H
		case r < 4:
			return `u`, H
		case r < 5:
			return `r`, H
		case r < 6:
			return `o`, H
		case r < 7:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ky`:
		H := 2.5
		r := rng.IntN(8)
		switch {
		case r < 2:
			return `a`, H
		case r < 4:
			return `l`, H
		case r < 5:
			return `w`, H
		case r < 6:
			return `d`, H
		case r < 7:
			return `p`, H
		case r < 8:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `rl`:
		H := 2.317113611025748
		r := rng.IntN(47)
		switch {
		case r < 15:
			return `i`, H
		case r < 25:
			return `y`, H
		case r < 35:
			return `e`, H
		case r < 41:
			return `o`, H
		case r < 46:
			return `a`, H
		case r < 47:
			return `d`, H
		default:
			panic("unexpected rand num")
		}
	case `f`:
		H := 2.5255137638602134
		r := rng.IntN(308)
		switch {
		case r < 76:
			return `r`, H
		case r < 139:
			return `l`, H
		case r < 193:
			return `a`, H
		case r < 235:
			return `o`, H
		case r < 276:
			return `i`, H
		case r < 308:
			return `e`, H
		default:
			panic("unexpected rand num")
		}
	case `np`:
		H := 2.2417812580773235
		r := rng.IntN(21)
		switch {
		case r < 7:
			return `a`, H
		case r < 13:
			return `l`, H
		case r < 17:
			return `i`, H
		case r < 19:
			return `o`, H
		case r < 20:
			return `e`, H
		case r < 21:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `fl`:
		H := 2.1729752291612567
		r := rng.IntN(104)
		switch {
		case r < 47:
			return `a`, H
		case r < 64:
			return `e`, H
		case r < 80:
			return `i`, H
		case r < 92:
			return `o`, H
		case r < 101:
			return `y`, H
		case r < 104:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `pr`:
		H := 1.9672224968373695
		r := rng.IntN(239)
		switch {
		case r < 99:
			return `o`, H
		case r < 169:
			return `e`, H
		case r < 214:
			return `i`, H
		case r < 230:
			return `a`, H
		case r < 237:
			return `u`, H
		case r < 239:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `tb`:
		H := 2.1309260807303376
		r := rng.IntN(17)
		switch {
		case r < 7:
			return `o`, H
		case r < 12:
			return `a`, H
		case r < 14:
			return `i`, H
		case r < 15:
			return `e`, H
		case r < 16:
			return `r`, H
		case r < 17:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `yb`:
		H := 2.4591479170272446
		r := rng.IntN(12)
		switch {
		case r < 3:
			return `a`, H
		case r < 6:
			return `o`, H
		case r < 8:
			return `e`, H
		case r < 10:
			return `r`, H
		case r < 11:
			return `i`, H
		case r < 12:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sf`:
		H := 2.5216406363433186
		r := rng.IntN(7)
		switch {
		case r < 2:
			return `i`, H
		case r < 3:
			return `y`, H
		case r < 4:
			return `u`, H
		case r < 5:
			return `r`, H
		case r < 6:
			return `e`, H
		case r < 7:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `bb`:
		H := 2.1580708663986847
		r := rng.IntN(52)
		switch {
		case r < 20:
			return `l`, H
		case r < 35:
			return `e`, H
		case r < 42:
			return `i`, H
		case r < 48:
			return `y`, H
		case r < 50:
			return `a`, H
		case r < 52:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ms`:
		H := 2.4193819456463714
		r := rng.IntN(9)
		switch {
		case r < 3:
			return `t`, H
		case r < 5:
			return `i`, H
		case r < 6:
			return `h`, H
		case r < 7:
			return `y`, H
		case r < 8:
			return `u`, H
		case r < 9:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `nn`:
		H := 2.203360784206382
		r := rng.IntN(74)
		switch {
		case r < 31:
			return `e`, H
		case r < 47:
			return `i`, H
		case r < 57:
			return `a`, H
		case r < 65:
			return `y`, H
		case r < 72:
			return `o`, H
		case r < 74:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rr`:
		H := 2.4052364684207306
		r := rng.IntN(91)
		switch {
		case r < 27:
			return `i`, H
		case r < 44:
			return `e`, H
		case r < 60:
			return `a`, H
		case r < 75:
			return `o`, H
		case r < 88:
			return `y`, H
		case r < 91:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `az`:
		H := 2.212331758086258
		r := rng.IntN(34)
		switch {
		case r < 13:
			return `i`, H
		case r < 22:
			return `e`, H
		case r < 27:
			return `z`, H
		case r < 30:
			return `y`, H
		case r < 33:
			return `a`, H
		case r < 34:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gr`:
		H := 2.042551738105775
		r := rng.IntN(159)
		switch {
		case r < 76:
			return `a`, H
		case r < 104:
			return `e`, H
		case r < 127:
			return `o`, H
		case r < 147:
			return `i`, H
		case r < 158:
			return `u`, H
		case r < 159:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `je`:
		H := 2.4300365325772657
		r := rng.IntN(16)
		switch {
		case r < 5:
			return `c`, H
		case r < 8:
			return `s`, H
		case r < 11:
			return `t`, H
		case r < 13:
			return `l`, H
		case r < 15:
			return `e`, H
		case r < 16:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `xa`:
		H := 2.5032583347756456
		r := rng.IntN(9)
		switch {
		case r < 2:
			return `g`, H
		case r < 4:
			return `m`, H
		case r < 6:
			return `b`, H
		case r < 7:
			return `c`, H
		case r < 8:
			return `l`, H
		case r < 9:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `nl`:
		H := 2.3212021480233496
		r := rng.IntN(52)
		switch {
		case r < 19:
			return `i`, H
		case r < 29:
			return `e`, H
		case r < 38:
			return `o`, H
		case r < 45:
			return `y`, H
		case r < 50:
			return `a`, H
		case r < 52:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sl`:
		H := 2.314430867489686
		r := rng.IntN(85)
		switch {
		case r < 26:
			return `i`, H
		case r < 50:
			return `a`, H
		case r < 63:
			return `e`, H
		case r < 76:
			return `o`, H
		case r < 81:
			return `u`, H
		case r < 85:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `cc`:
		H := 1.9268680175017443
		r := rng.IntN(21)
		switch {
		case r < 12:
			return `u`, H
		case r < 15:
			return `o`, H
		case r < 17:
			return `l`, H
		case r < 19:
			return `e`, H
		case r < 20:
			return `a`, H
		case r < 21:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `ox`:
		H := 1.8182626942811975
		r := rng.IntN(22)
		switch {
		case r < 13:
			return `i`, H
		case r < 17:
			return `y`, H
		case r < 19:
			return `e`, H
		case r < 20:
			return `f`, H
		case r < 21:
			return `c`, H
		case r < 22:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `mm`:
		H := 2.4645728472008237
		r := rng.IntN(69)
		switch {
		case r < 19:
			return `e`, H
		case r < 35:
			return `o`, H
		case r < 46:
			return `i`, H
		case r < 55:
			return `y`, H
		case r < 64:
			return `a`, H
		case r < 69:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `dr`:
		H := 2.258664230601291
		r := rng.IntN(97)
		switch {
		case r < 27:
			return `a`, H
		case r < 50:
			return `e`, H
		case r < 71:
			return `i`, H
		case r < 91:
			return `o`, H
		case r < 95:
			return `u`, H
		case r < 97:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `tr`:
		H := 2.3991891290867966
		r := rng.IntN(267)
		switch {
		case r < 75:
			return `a`, H
		case r < 147:
			return `i`, H
		case r < 186:
			return `o`, H
		case r < 221:
			return `e`, H
		case r < 254:
			return `u`, H
		case r < 267:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `dy`:
		H := 2.446439344671015
		r := rng.IntN(10)
		switch {
		case r < 3:
			return `n`, H
		case r < 5:
			return `s`, H
		case r < 7:
			return `i`, H
		case r < 8:
			return `l`, H
		case r < 9:
			return `m`, H
		case r < 10:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `xt`:
		H := 2.077235445408308
		r := rng.IntN(20)
		switch {
		case r < 9:
			return `e`, H
		case r < 14:
			return `r`, H
		case r < 17:
			return `i`, H
		case r < 18:
			return `h`, H
		case r < 19:
			return `y`, H
		case r < 20:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ev`:
		H := 1.9942008615410574
		r := rng.IntN(106)
		switch {
		case r < 42:
			return `e`, H
		case r < 72:
			return `i`, H
		case r < 89:
			return `a`, H
		case r < 104:
			return `o`, H
		case r < 105:
			return `y`, H
		case r < 106:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ry`:
		H := 2.5220552088742005
		r := rng.IntN(12)
		switch {
		case r < 4:
			return `i`, H
		case r < 7:
			return `p`, H
		case r < 8:
			return `w`, H
		case r < 9:
			return `s`, H
		case r < 10:
			return `d`, H
		case r < 11:
			return `o`, H
		case r < 12:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `nf`:
		H := 2.585460534063065
		r := rng.IntN(51)
		switch {
		case r < 17:
			return `i`, H
		case r < 27:
			return `o`, H
		case r < 33:
			return `u`, H
		case r < 38:
			return `a`, H
		case r < 43:
			return `r`, H
		case r < 47:
			return `e`, H
		case r < 51:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `oi`:
		H := 2.2369637280222663
		r := rng.IntN(66)
		switch {
		case r < 24:
			return `n`, H
		case r < 44:
			return `l`, H
		case r < 51:
			return `d`, H
		case r < 58:
			return `c`, H
		case r < 64:
			return `s`, H
		case r < 65:
			return `r`, H
		case r < 66:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `nb`:
		H := 2.759079570624175
		r := rng.IntN(25)
		switch {
		case r < 5:
			return `e`, H
		case r < 9:
			return `r`, H
		case r < 13:
			return `o`, H
		case r < 17:
			return `u`, H
		case r < 20:
			return `a`, H
		case r < 23:
			return `l`, H
		case r < 25:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `xp`:
		H := 2.5096856969120642
		r := rng.IntN(30)
		switch {
		case r < 8:
			return `e`, H
		case r < 15:
			return `l`, H
		case r < 21:
			return `o`, H
		case r < 25:
			return `a`, H
		case r < 28:
			return `i`, H
		case r < 29:
			return `r`, H
		case r < 30:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ty`:
		H := 2.3261207468426806
		r := rng.IntN(20)
		switch {
		case r < 7:
			return `l`, H
		case r < 13:
			return `p`, H
		case r < 16:
			return `i`, H
		case r < 17:
			return `f`, H
		case r < 18:
			return `h`, H
		case r < 19:
			return `c`, H
		case r < 20:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ks`:
		H := 2.460904932167797
		r := rng.IntN(17)
		switch {
		case r < 7:
			return `t`, H
		case r < 9:
			return `h`, H
		case r < 11:
			return `a`, H
		case r < 13:
			return `l`, H
		case r < 15:
			return `p`, H
		case r < 16:
			return `i`, H
		case r < 17:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `eu`:
		H := 2.501609497059027
		r := rng.IntN(20)
		switch {
		case r < 7:
			return `r`, H
		case r < 11:
			return `m`, H
		case r < 14:
			return `s`, H
		case r < 16:
			return `n`, H
		case r < 18:
			return `t`, H
		case r < 19:
			return `c`, H
		case r < 20:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `fu`:
		H := 1.5921871522135118
		r := rng.IntN(61)
		switch {
		case r < 40:
			return `l`, H
		case r < 52:
			return `s`, H
		case r < 55:
			return `r`, H
		case r < 57:
			return `n`, H
		case r < 59:
			return `t`, H
		case r < 60:
			return `e`, H
		case r < 61:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `zi`:
		H := 1.8308122292922697
		r := rng.IntN(33)
		switch {
		case r < 20:
			return `n`, H
		case r < 25:
			return `p`, H
		case r < 29:
			return `l`, H
		case r < 30:
			return `g`, H
		case r < 31:
			return `c`, H
		case r < 32:
			return `e`, H
		case r < 33:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `rk`:
		H := 2.4863286677871854
		r := rng.IntN(17)
		switch {
		case r < 5:
			return `i`, H
		case r < 10:
			return `e`, H
		case r < 12:
			return `n`, H
		case r < 14:
			return `y`, H
		case r < 15:
			return `w`, H
		case r < 16:
			return `r`, H
		case r < 17:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `xc`:
		H := 2.6659573209491745
		r := rng.IntN(20)
		switch {
		case r < 5:
			return `l`, H
		case r < 9:
			return `u`, H
		case r < 12:
			return `a`, H
		case r < 15:
			return `e`, H
		case r < 17:
			return `i`, H
		case r < 19:
			return `r`, H
		case r < 20:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `dd`:
		H := 2.279251222260265
		r := rng.IntN(48)
		switch {
		case r < 18:
			return `l`, H
		case r < 31:
			return `e`, H
		case r < 39:
			return `i`, H
		case r < 42:
			return `y`, H
		case r < 45:
			return `a`, H
		case r < 47:
			return `h`, H
		case r < 48:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ye`:
		H := 2.5365766899407323
		r := rng.IntN(24)
		switch {
		case r < 7:
			return `r`, H
		case r < 13:
			return `a`, H
		case r < 17:
			return `d`, H
		case r < 19:
			return `s`, H
		case r < 21:
			return `e`, H
		case r < 23:
			return `l`, H
		case r < 24:
			return `n`, H
		default:
			panic("unexpected rand num")
		}
	case `pp`:
		H := 2.4303217508328046
		r := rng.IntN(114)
		switch {
		case r < 43:
			return `e`, H
		case r < 69:
			return `i`, H
		case r < 84:
			return `l`, H
		case r < 94:
			return `y`, H
		case r < 103:
			return `r`, H
		case r < 111:
			return `o`, H
		case r < 113:
			return `a`, H
		case r < 114:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ka`:
		H := 2.7744019192887706
		r := rng.IntN(18)
		switch {
		case r < 5:
			return `b`, H
		case r < 8:
			return `l`, H
		case r < 11:
			return `r`, H
		case r < 13:
			return `g`, H
		case r < 15:
			return `t`, H
		case r < 16:
			return `n`, H
		case r < 17:
			return `c`, H
		case r < 18:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `bs`:
		H := 2.615277848886663
		r := rng.IntN(35)
		switch {
		case r < 10:
			return `e`, H
		case r < 17:
			return `t`, H
		case r < 23:
			return `o`, H
		case r < 28:
			return `i`, H
		case r < 32:
			return `c`, H
		case r < 33:
			return `u`, H
		case r < 34:
			return `l`, H
		case r < 35:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `if`:
		H := 2.5951312665142345
		r := rng.IntN(106)
		switch {
		case r < 27:
			return `t`, H
		case r < 51:
			return `y`, H
		case r < 75:
			return `i`, H
		case r < 88:
			return `f`, H
		case r < 96:
			return `e`, H
		case r < 100:
			return `l`, H
		case r < 103:
			return `u`, H
		case r < 106:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ze`:
		H := 2.471058306628392
		r := rng.IntN(27)
		switch {
		case r < 11:
			return `r`, H
		case r < 16:
			return `d`, H
		case r < 20:
			return `n`, H
		case r < 22:
			return `l`, H
		case r < 24:
			return `s`, H
		case r < 25:
			return `a`, H
		case r < 26:
			return `p`, H
		case r < 27:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ud`:
		H := 2.436743733204307
		r := rng.IntN(67)
		switch {
		case r < 22:
			return `e`, H
		case r < 38:
			return `i`, H
		case r < 50:
			return `d`, H
		case r < 59:
			return `g`, H
		case r < 62:
			return `a`, H
		case r < 65:
			return `o`, H
		case r < 66:
			return `y`, H
		case r < 67:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `d`:
		H := 2.414563621434015
		r := rng.IntN(547)
		switch {
		case r < 200:
			return `e`, H
		case r < 338:
			return `i`, H
		case r < 401:
			return `r`, H
		case r < 451:
			return `a`, H
		case r < 500:
			return `o`, H
		case r < 535:
			return `u`, H
		case r < 542:
			return `w`, H
		case r < 547:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ei`:
		H := 2.8559049224696933
		r := rng.IntN(21)
		switch {
		case r < 4:
			return `n`, H
		case r < 8:
			return `t`, H
		case r < 11:
			return `v`, H
		case r < 14:
			return `s`, H
		case r < 17:
			return `g`, H
		case r < 19:
			return `l`, H
		case r < 20:
			return `m`, H
		case r < 21:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ey`:
		H := 2.94770277922009
		r := rng.IntN(9)
		switch {
		case r < 2:
			return `a`, H
		case r < 3:
			return `w`, H
		case r < 4:
			return `h`, H
		case r < 5:
			return `l`, H
		case r < 6:
			return `e`, H
		case r < 7:
			return `o`, H
		case r < 8:
			return `m`, H
		case r < 9:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `rg`:
		H := 2.6225466508345625
		r := rng.IntN(59)
		switch {
		case r < 23:
			return `e`, H
		case r < 32:
			return `i`, H
		case r < 38:
			return `l`, H
		case r < 44:
			return `o`, H
		case r < 48:
			return `y`, H
		case r < 52:
			return `r`, H
		case r < 56:
			return `a`, H
		case r < 59:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `t`:
		H := 2.467515412367413
		r := rng.IntN(323)
		switch {
		case r < 105:
			return `r`, H
		case r < 169:
			return `a`, H
		case r < 230:
			return `h`, H
		case r < 271:
			return `i`, H
		case r < 296:
			return `u`, H
		case r < 319:
			return `w`, H
		case r < 322:
			return `y`, H
		case r < 323:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `bt`:
		H := 2.8073549220576046
		r := rng.IntN(14)
		switch {
		case r < 4:
			return `l`, H
		case r < 6:
			return `a`, H
		case r < 8:
			return `r`, H
		case r < 10:
			return `o`, H
		case r < 11:
			return `u`, H
		case r < 12:
			return `e`, H
		case r < 13:
			return `i`, H
		case r < 14:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `yo`:
		H := 2.692380602454975
		r := rng.IntN(14)
		switch {
		case r < 5:
			return `n`, H
		case r < 7:
			return `y`, H
		case r < 9:
			return `g`, H
		case r < 10:
			return `f`, H
		case r < 11:
			return `v`, H
		case r < 12:
			return `u`, H
		case r < 13:
			return `r`, H
		case r < 14:
			return `d`, H
		default:
			panic("unexpected rand num")
		}
	case `u`:
		H := 0.8648744526550217
		r := rng.IntN(476)
		switch {
		case r < 410:
			return `n`, H
		case r < 446:
			return `p`, H
		case r < 455:
			return `r`, H
		case r < 463:
			return `s`, H
		case r < 469:
			return `t`, H
		case r < 472:
			return `l`, H
		case r < 475:
			return `m`, H
		case r < 476:
			return `d`, H
		default:
			panic("unexpected rand num")
		}
	case `ps`:
		H := 2.489655670358648
		r := rng.IntN(27)
		switch {
		case r < 8:
			return `t`, H
		case r < 15:
			return `e`, H
		case r < 21:
			return `i`, H
		case r < 23:
			return `y`, H
		case r < 24:
			return `w`, H
		case r < 25:
			return `u`, H
		case r < 26:
			return `a`, H
		case r < 27:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `pt`:
		H := 2.3129328787410666
		r := rng.IntN(45)
		switch {
		case r < 21:
			return `i`, H
		case r < 28:
			return `u`, H
		case r < 35:
			return `o`, H
		case r < 38:
			return `l`, H
		case r < 41:
			return `e`, H
		case r < 43:
			return `a`, H
		case r < 44:
			return `h`, H
		case r < 45:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `af`:
		H := 2.1819962604162924
		r := rng.IntN(59)
		switch {
		case r < 25:
			return `f`, H
		case r < 42:
			return `t`, H
		case r < 50:
			return `e`, H
		case r < 53:
			return `l`, H
		case r < 55:
			return `a`, H
		case r < 57:
			return `r`, H
		case r < 58:
			return `n`, H
		case r < 59:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `rf`:
		H := 2.5833333333333335
		r := rng.IntN(24)
		switch {
		case r < 9:
			return `e`, H
		case r < 13:
			return `u`, H
		case r < 16:
			return `a`, H
		case r < 19:
			return `i`, H
		case r < 21:
			return `l`, H
		case r < 22:
			return `r`, H
		case r < 23:
			return `o`, H
		case r < 24:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `k`:
		H := 2.2102710078071865
		r := rng.IntN(55)
		switch {
		case r < 27:
			return `i`, H
		case r < 38:
			return `e`, H
		case r < 43:
			return `n`, H
		case r < 48:
			return `a`, H
		case r < 51:
			return `o`, H
		case r < 53:
			return `u`, H
		case r < 54:
			return `l`, H
		case r < 55:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `mb`:
		H := 2.4620228668017208
		r := rng.IntN(75)
		switch {
		case r < 30:
			return `l`, H
		case r < 44:
			return `e`, H
		case r < 52:
			return `a`, H
		case r < 60:
			return `i`, H
		case r < 68:
			return `o`, H
		case r < 72:
			return `u`, H
		case r < 74:
			return `r`, H
		case r < 75:
			return `n`, H
		default:
			panic("unexpected rand num")
		}
	case `ph`:
		H := 2.525379371610213
		r := rng.IntN(56)
		switch {
		case r < 17:
			return `o`, H
		case r < 29:
			return `a`, H
		case r < 41:
			return `e`, H
		case r < 47:
			return `i`, H
		case r < 51:
			return `r`, H
		case r < 54:
			return `y`, H
		case r < 55:
			return `l`, H
		case r < 56:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sc`:
		H := 2.5124428937801095
		r := rng.IntN(134)
		switch {
		case r < 40:
			return `a`, H
		case r < 73:
			return `o`, H
		case r < 99:
			return `r`, H
		case r < 112:
			return `u`, H
		case r < 121:
			return `h`, H
		case r < 129:
			return `e`, H
		case r < 133:
			return `i`, H
		case r < 134:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `vo`:
		H := 2.871054088933562
		r := rng.IntN(57)
		switch {
		case r < 16:
			return `r`, H
		case r < 27:
			return `l`, H
		case r < 34:
			return `c`, H
		case r < 40:
			return `u`, H
		case r < 45:
			return `t`, H
		case r < 49:
			return `i`, H
		case r < 53:
			return `k`, H
		case r < 55:
			return `w`, H
		case r < 57:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `c`:
		H := 2.671815163549163
		r := rng.IntN(751)
		switch {
		case r < 224:
			return `o`, H
		case r < 389:
			return `a`, H
		case r < 496:
			return `r`, H
		case r < 593:
			return `h`, H
		case r < 656:
			return `l`, H
		case r < 708:
			return `u`, H
		case r < 727:
			return `i`, H
		case r < 743:
			return `e`, H
		case r < 751:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `g`:
		H := 2.6968078160957982
		r := rng.IntN(304)
		switch {
		case r < 97:
			return `r`, H
		case r < 142:
			return `a`, H
		case r < 185:
			return `l`, H
		case r < 218:
			return `e`, H
		case r < 250:
			return `o`, H
		case r < 277:
			return `u`, H
		case r < 301:
			return `i`, H
		case r < 303:
			return `n`, H
		case r < 304:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `sk`:
		H := 2.0077378472545835
		r := rng.IntN(61)
		switch {
		case r < 32:
			return `i`, H
		case r < 46:
			return `e`, H
		case r < 54:
			return `y`, H
		case r < 56:
			return `a`, H
		case r < 57:
			return `n`, H
		case r < 58:
			return `w`, H
		case r < 59:
			return `l`, H
		case r < 60:
			return `t`, H
		case r < 61:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `lp`:
		H := 3.0503018554349826
		r := rng.IntN(19)
		switch {
		case r < 4:
			return `h`, H
		case r < 7:
			return `i`, H
		case r < 9:
			return `f`, H
		case r < 11:
			return `a`, H
		case r < 13:
			return `e`, H
		case r < 15:
			return `l`, H
		case r < 17:
			return `t`, H
		case r < 18:
			return `r`, H
		case r < 19:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ya`:
		H := 2.7153266987718103
		r := rng.IntN(19)
		switch {
		case r < 6:
			return `r`, H
		case r < 11:
			return `b`, H
		case r < 13:
			return `l`, H
		case r < 14:
			return `n`, H
		case r < 15:
			return `w`, H
		case r < 16:
			return `g`, H
		case r < 17:
			return `h`, H
		case r < 18:
			return `p`, H
		case r < 19:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ny`:
		H := 3.039148671903072
		r := rng.IntN(14)
		switch {
		case r < 3:
			return `m`, H
		case r < 5:
			return `w`, H
		case r < 7:
			return `t`, H
		case r < 9:
			return `o`, H
		case r < 10:
			return `h`, H
		case r < 11:
			return `x`, H
		case r < 12:
			return `l`, H
		case r < 13:
			return `p`, H
		case r < 14:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `yp`:
		H := 2.7132696895151076
		r := rng.IntN(25)
		switch {
		case r < 8:
			return `e`, H
		case r < 14:
			return `n`, H
		case r < 17:
			return `t`, H
		case r < 19:
			return `h`, H
		case r < 21:
			return `o`, H
		case r < 22:
			return `a`, H
		case r < 23:
			return `l`, H
		case r < 24:
			return `i`, H
		case r < 25:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ht`:
		H := 2.6292492238560348
		r := rng.IntN(19)
		switch {
		case r < 8:
			return `e`, H
		case r < 11:
			return `l`, H
		case r < 13:
			return `i`, H
		case r < 14:
			return `n`, H
		case r < 15:
			return `f`, H
		case r < 16:
			return `w`, H
		case r < 17:
			return `y`, H
		case r < 18:
			return `r`, H
		case r < 19:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rc`:
		H := 2.8029839504926644
		r := rng.IntN(53)
		switch {
		case r < 15:
			return `h`, H
		case r < 26:
			return `e`, H
		case r < 34:
			return `o`, H
		case r < 39:
			return `u`, H
		case r < 43:
			return `a`, H
		case r < 47:
			return `i`, H
		case r < 50:
			return `l`, H
		case r < 52:
			return `t`, H
		case r < 53:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ai`:
		H := 2.315377442431913
		r := rng.IntN(132)
		switch {
		case r < 56:
			return `n`, H
		case r < 88:
			return `l`, H
		case r < 102:
			return `r`, H
		case r < 115:
			return `d`, H
		case r < 123:
			return `m`, H
		case r < 128:
			return `s`, H
		case r < 130:
			return `t`, H
		case r < 131:
			return `k`, H
		case r < 132:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `rb`:
		H := 2.7675793883891737
		r := rng.IntN(47)
		switch {
		case r < 12:
			return `o`, H
		case r < 23:
			return `i`, H
		case r < 30:
			return `a`, H
		case r < 36:
			return `e`, H
		case r < 40:
			return `l`, H
		case r < 43:
			return `y`, H
		case r < 45:
			return `u`, H
		case r < 46:
			return `r`, H
		case r < 47:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ft`:
		H := 2.4952508233936745
		r := rng.IntN(33)
		switch {
		case r < 13:
			return `e`, H
		case r < 20:
			return `i`, H
		case r < 25:
			return `y`, H
		case r < 28:
			return `l`, H
		case r < 29:
			return `n`, H
		case r < 30:
			return `w`, H
		case r < 31:
			return `h`, H
		case r < 32:
			return `s`, H
		case r < 33:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ff`:
		H := 2.723780466448913
		r := rng.IntN(63)
		switch {
		case r < 16:
			return `e`, H
		case r < 30:
			return `i`, H
		case r < 44:
			return `l`, H
		case r < 49:
			return `o`, H
		case r < 53:
			return `u`, H
		case r < 57:
			return `y`, H
		case r < 60:
			return `r`, H
		case r < 62:
			return `a`, H
		case r < 63:
			return `n`, H
		default:
			panic("unexpected rand num")
		}
	case `p`:
		H := 2.5923931878457345
		r := rng.IntN(578)
		switch {
		case r < 166:
			return `r`, H
		case r < 302:
			return `a`, H
		case r < 379:
			return `o`, H
		case r < 447:
			return `e`, H
		case r < 506:
			return `u`, H
		case r < 562:
			return `l`, H
		case r < 574:
			return `h`, H
		case r < 577:
			return `y`, H
		case r < 578:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `eb`:
		H := 2.887941726415332
		r := rng.IntN(47)
		switch {
		case r < 10:
			return `a`, H
		case r < 20:
			return `o`, H
		case r < 29:
			return `u`, H
		case r < 34:
			return `r`, H
		case r < 38:
			return `l`, H
		case r < 41:
			return `i`, H
		case r < 43:
			return `t`, H
		case r < 45:
			return `e`, H
		case r < 47:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `tt`:
		H := 2.3319551423402523
		r := rng.IntN(127)
		switch {
		case r < 59:
			return `e`, H
		case r < 80:
			return `i`, H
		case r < 96:
			return `l`, H
		case r < 106:
			return `a`, H
		case r < 116:
			return `o`, H
		case r < 123:
			return `y`, H
		case r < 125:
			return `r`, H
		case r < 126:
			return `h`, H
		case r < 127:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rp`:
		H := 2.9031815050305063
		r := rng.IntN(41)
		switch {
		case r < 8:
			return `l`, H
		case r < 16:
			return `o`, H
		case r < 22:
			return `e`, H
		case r < 28:
			return `a`, H
		case r < 33:
			return `i`, H
		case r < 37:
			return `h`, H
		case r < 39:
			return `r`, H
		case r < 40:
			return `n`, H
		case r < 41:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `od`:
		H := 2.7181256737103596
		r := rng.IntN(50)
		switch {
		case r < 13:
			return `e`, H
		case r < 25:
			return `i`, H
		case r < 33:
			return `y`, H
		case r < 40:
			return `u`, H
		case r < 43:
			return `g`, H
		case r < 46:
			return `o`, H
		case r < 48:
			return `l`, H
		case r < 49:
			return `a`, H
		case r < 50:
			return `d`, H
		default:
			panic("unexpected rand num")
		}
	case `io`:
		H := 1.3890034152520174
		r := rng.IntN(187)
		switch {
		case r < 140:
			return `n`, H
		case r < 163:
			return `u`, H
		case r < 169:
			return `r`, H
		case r < 175:
			return `l`, H
		case r < 180:
			return `t`, H
		case r < 183:
			return `c`, H
		case r < 185:
			return `d`, H
		case r < 186:
			return `x`, H
		case r < 187:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ib`:
		H := 2.5862770862160236
		r := rng.IntN(51)
		switch {
		case r < 19:
			return `l`, H
		case r < 30:
			return `e`, H
		case r < 36:
			return `u`, H
		case r < 41:
			return `b`, H
		case r < 44:
			return `i`, H
		case r < 47:
			return `r`, H
		case r < 49:
			return `a`, H
		case r < 50:
			return `o`, H
		case r < 51:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `sp`:
		H := 2.8235716352955578
		r := rng.IntN(160)
		switch {
		case r < 38:
			return `o`, H
		case r < 73:
			return `e`, H
		case r < 102:
			return `i`, H
		case r < 119:
			return `l`, H
		case r < 136:
			return `r`, H
		case r < 145:
			return `a`, H
		case r < 152:
			return `h`, H
		case r < 156:
			return `u`, H
		case r < 159:
			return `y`, H
		case r < 160:
			return `n`, H
		default:
			panic("unexpected rand num")
		}
	case `oy`:
		H := 2.9638096525384383
		r := rng.IntN(23)
		switch {
		case r < 6:
			return `e`, H
		case r < 11:
			return `a`, H
		case r < 14:
			return `o`, H
		case r < 16:
			return `i`, H
		case r < 18:
			return `s`, H
		case r < 19:
			return `n`, H
		case r < 20:
			return `f`, H
		case r < 21:
			return `l`, H
		case r < 22:
			return `m`, H
		case r < 23:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ts`:
		H := 3.135821824148036
		r := rng.IntN(22)
		switch {
		case r < 5:
			return `h`, H
		case r < 8:
			return `i`, H
		case r < 11:
			return `m`, H
		case r < 13:
			return `y`, H
		case r < 15:
			return `u`, H
		case r < 17:
			return `k`, H
		case r < 19:
			return `o`, H
		case r < 20:
			return `c`, H
		case r < 21:
			return `e`, H
		case r < 22:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ct`:
		H := 2.5852558586975434
		r := rng.IntN(118)
		switch {
		case r < 44:
			return `i`, H
		case r < 67:
			return `o`, H
		case r < 83:
			return `e`, H
		case r < 96:
			return `a`, H
		case r < 106:
			return `u`, H
		case r < 110:
			return `l`, H
		case r < 113:
			return `s`, H
		case r < 116:
			return `r`, H
		case r < 117:
			return `f`, H
		case r < 118:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `we`:
		H := 2.748621296142989
		r := rng.IntN(58)
		switch {
		case r < 18:
			return `r`, H
		case r < 27:
			return `l`, H
		case r < 35:
			return `a`, H
		case r < 43:
			return `e`, H
		case r < 51:
			return `d`, H
		case r < 54:
			return `n`, H
		case r < 55:
			return `i`, H
		case r < 56:
			return `s`, H
		case r < 57:
			return `b`, H
		case r < 58:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ds`:
		H := 3.0464393446710147
		r := rng.IntN(20)
		switch {
		case r < 6:
			return `t`, H
		case r < 8:
			return `h`, H
		case r < 10:
			return `i`, H
		case r < 12:
			return `e`, H
		case r < 14:
			return `m`, H
		case r < 16:
			return `c`, H
		case r < 17:
			return `f`, H
		case r < 18:
			return `a`, H
		case r < 19:
			return `o`, H
		case r < 20:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `oe`:
		H := 3.2516291673878226
		r := rng.IntN(12)
		switch {
		case r < 2:
			return `s`, H
		case r < 4:
			return `r`, H
		case r < 5:
			return `n`, H
		case r < 6:
			return `x`, H
		case r < 7:
			return `d`, H
		case r < 8:
			return `y`, H
		case r < 9:
			return `t`, H
		case r < 10:
			return `m`, H
		case r < 11:
			return `i`, H
		case r < 12:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `nc`:
		H := 2.8557990432188514
		r := rng.IntN(207)
		switch {
		case r < 66:
			return `e`, H
		case r < 109:
			return `h`, H
		case r < 130:
			return `i`, H
		case r < 149:
			return `o`, H
		case r < 165:
			return `y`, H
		case r < 177:
			return `l`, H
		case r < 188:
			return `t`, H
		case r < 196:
			return `r`, H
		case r < 202:
			return `a`, H
		case r < 207:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ia`:
		H := 2.7871615786315345
		r := rng.IntN(106)
		switch {
		case r < 26:
			return `l`, H
		case r < 48:
			return `n`, H
		case r < 66:
			return `t`, H
		case r < 84:
			return `b`, H
		case r < 93:
			return `r`, H
		case r < 97:
			return `g`, H
		case r < 100:
			return `c`, H
		case r < 102:
			return `s`, H
		case r < 104:
			return `m`, H
		case r < 106:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `lt`:
		H := 2.9388580483450712
		r := rng.IntN(36)
		switch {
		case r < 11:
			return `i`, H
		case r < 17:
			return `e`, H
		case r < 21:
			return `y`, H
		case r < 25:
			return `r`, H
		case r < 28:
			return `h`, H
		case r < 30:
			return `u`, H
		case r < 32:
			return `a`, H
		case r < 34:
			return `z`, H
		case r < 35:
			return `o`, H
		case r < 36:
			return `t`, H
		default:
			panic("unexpected rand num")
		}
	case `ef`:
		H := 3.1182289091344035
		r := rng.IntN(91)
		switch {
		case r < 17:
			return `u`, H
		case r < 33:
			return `i`, H
		case r < 45:
			return `e`, H
		case r < 56:
			return `o`, H
		case r < 66:
			return `l`, H
		case r < 73:
			return `r`, H
		case r < 79:
			return `a`, H
		case r < 85:
			return `t`, H
		case r < 89:
			return `f`, H
		case r < 91:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ou`:
		H := 2.433480709472623
		r := rng.IntN(258)
		switch {
		case r < 73:
			return `s`, H
		case r < 146:
			return `t`, H
		case r < 211:
			return `n`, H
		case r < 227:
			return `r`, H
		case r < 237:
			return `c`, H
		case r < 244:
			return `g`, H
		case r < 250:
			return `p`, H
		case r < 253:
			return `d`, H
		case r < 256:
			return `b`, H
		case r < 258:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `nk`:
		H := 2.955870327431456
		r := rng.IntN(61)
		switch {
		case r < 17:
			return `i`, H
		case r < 30:
			return `e`, H
		case r < 38:
			return `l`, H
		case r < 44:
			return `y`, H
		case r < 49:
			return `n`, H
		case r < 52:
			return `s`, H
		case r < 54:
			return `h`, H
		case r < 56:
			return `a`, H
		case r < 58:
			return `m`, H
		case r < 60:
			return `b`, H
		case r < 61:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ly`:
		H := 3.272769892034794
		r := rng.IntN(26)
		switch {
		case r < 5:
			return `i`, H
		case r < 9:
			return `r`, H
		case r < 12:
			return `s`, H
		case r < 15:
			return `z`, H
		case r < 17:
			return `e`, H
		case r < 19:
			return `a`, H
		case r < 21:
			return `m`, H
		case r < 23:
			return `g`, H
		case r < 24:
			return `w`, H
		case r < 25:
			return `o`, H
		case r < 26:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `um`:
		H := 2.7569798111220467
		r := rng.IntN(99)
		switch {
		case r < 29:
			return `b`, H
		case r < 49:
			return `p`, H
		case r < 64:
			return `m`, H
		case r < 78:
			return `e`, H
		case r < 87:
			return `i`, H
		case r < 92:
			return `o`, H
		case r < 94:
			return `s`, H
		case r < 96:
			return `a`, H
		case r < 97:
			return `n`, H
		case r < 98:
			return `d`, H
		case r < 99:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `eo`:
		H := 3.1748858687242363
		r := rng.IntN(23)
		switch {
		case r < 6:
			return `l`, H
		case r < 9:
			return `r`, H
		case r < 12:
			return `c`, H
		case r < 14:
			return `n`, H
		case r < 16:
			return `d`, H
		case r < 18:
			return `m`, H
		case r < 19:
			return `t`, H
		case r < 20:
			return `p`, H
		case r < 21:
			return `g`, H
		case r < 22:
			return `v`, H
		case r < 23:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `gg`:
		H := 2.5365464726603135
		r := rng.IntN(59)
		switch {
		case r < 19:
			return `l`, H
		case r < 34:
			return `e`, H
		case r < 47:
			return `i`, H
		case r < 51:
			return `y`, H
		case r < 53:
			return `a`, H
		case r < 54:
			return `n`, H
		case r < 55:
			return `o`, H
		case r < 56:
			return `r`, H
		case r < 57:
			return `s`, H
		case r < 58:
			return `p`, H
		case r < 59:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `be`:
		H := 2.6259431593969045
		r := rng.IntN(69)
		switch {
		case r < 27:
			return `r`, H
		case r < 43:
			return `d`, H
		case r < 50:
			return `l`, H
		case r < 56:
			return `a`, H
		case r < 59:
			return `t`, H
		case r < 61:
			return `n`, H
		case r < 63:
			return `e`, H
		case r < 65:
			return `s`, H
		case r < 67:
			return `c`, H
		case r < 68:
			return `y`, H
		case r < 69:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `jo`:
		H := 2.9275496777656036
		r := rng.IntN(32)
		switch {
		case r < 11:
			return `y`, H
		case r < 17:
			return `i`, H
		case r < 20:
			return `l`, H
		case r < 22:
			return `r`, H
		case r < 24:
			return `k`, H
		case r < 26:
			return `c`, H
		case r < 28:
			return `g`, H
		case r < 29:
			return `h`, H
		case r < 30:
			return `t`, H
		case r < 31:
			return `v`, H
		case r < 32:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ug`:
		H := 2.8739810481273578
		r := rng.IntN(34)
		switch {
		case r < 10:
			return `g`, H
		case r < 19:
			return `h`, H
		case r < 22:
			return `n`, H
		case r < 25:
			return `a`, H
		case r < 27:
			return `u`, H
		case r < 29:
			return `l`, H
		case r < 30:
			return `e`, H
		case r < 31:
			return `o`, H
		case r < 32:
			return `s`, H
		case r < 33:
			return `i`, H
		case r < 34:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ex`:
		H := 2.823204116207561
		r := rng.IntN(102)
		switch {
		case r < 30:
			return `p`, H
		case r < 50:
			return `t`, H
		case r < 69:
			return `c`, H
		case r < 77:
			return `e`, H
		case r < 85:
			return `i`, H
		case r < 91:
			return `a`, H
		case r < 95:
			return `o`, H
		case r < 98:
			return `h`, H
		case r < 100:
			return `u`, H
		case r < 101:
			return `f`, H
		case r < 102:
			return `q`, H
		default:
			panic("unexpected rand num")
		}
	case `va`:
		H := 3.0359223229700754
		r := rng.IntN(108)
		switch {
		case r < 28:
			return `l`, H
		case r < 47:
			return `t`, H
		case r < 59:
			return `r`, H
		case r < 70:
			return `b`, H
		case r < 80:
			return `n`, H
		case r < 89:
			return `g`, H
		case r < 97:
			return `c`, H
		case r < 103:
			return `s`, H
		case r < 105:
			return `i`, H
		case r < 107:
			return `p`, H
		case r < 108:
			return `d`, H
		default:
			panic("unexpected rand num")
		}
	case `ed`:
		H := 2.9112023765018518
		r := rng.IntN(74)
		switch {
		case r < 25:
			return `i`, H
		case r < 35:
			return `u`, H
		case r < 43:
			return `e`, H
		case r < 51:
			return `a`, H
		case r < 59:
			return `g`, H
		case r < 63:
			return `d`, H
		case r < 67:
			return `o`, H
		case r < 70:
			return `y`, H
		case r < 72:
			return `l`, H
		case r < 73:
			return `t`, H
		case r < 74:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ss`:
		H := 2.826274120140731
		r := rng.IntN(71)
		switch {
		case r < 19:
			return `e`, H
		case r < 38:
			return `i`, H
		case r < 47:
			return `u`, H
		case r < 54:
			return `a`, H
		case r < 59:
			return `o`, H
		case r < 63:
			return `y`, H
		case r < 66:
			return `l`, H
		case r < 68:
			return `p`, H
		case r < 69:
			return `f`, H
		case r < 70:
			return `w`, H
		case r < 71:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ue`:
		H := 3.066701577503766
		r := rng.IntN(39)
		switch {
		case r < 9:
			return `l`, H
		case r < 17:
			return `n`, H
		case r < 21:
			return `d`, H
		case r < 25:
			return `e`, H
		case r < 29:
			return `s`, H
		case r < 33:
			return `a`, H
		case r < 35:
			return `r`, H
		case r < 36:
			return `f`, H
		case r < 37:
			return `t`, H
		case r < 38:
			return `u`, H
		case r < 39:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `of`:
		H := 3.1173143543437773
		r := rng.IntN(26)
		switch {
		case r < 7:
			return `f`, H
		case r < 11:
			return `e`, H
		case r < 14:
			return `i`, H
		case r < 17:
			return `a`, H
		case r < 19:
			return `o`, H
		case r < 21:
			return `t`, H
		case r < 22:
			return `s`, H
		case r < 23:
			return `y`, H
		case r < 24:
			return `r`, H
		case r < 25:
			return `u`, H
		case r < 26:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `wo`:
		H := 2.4661778189349435
		r := rng.IntN(67)
		switch {
		case r < 33:
			return `r`, H
		case r < 43:
			return `o`, H
		case r < 51:
			return `m`, H
		case r < 54:
			return `v`, H
		case r < 57:
			return `b`, H
		case r < 59:
			return `w`, H
		case r < 61:
			return `k`, H
		case r < 63:
			return `u`, H
		case r < 65:
			return `l`, H
		case r < 66:
			return `n`, H
		case r < 67:
			return `f`, H
		default:
			panic("unexpected rand num")
		}
	case `me`:
		H := 2.700171980219928
		r := rng.IntN(131)
		switch {
		case r < 38:
			return `n`, H
		case r < 70:
			return `r`, H
		case r < 93:
			return `d`, H
		case r < 104:
			return `t`, H
		case r < 113:
			return `s`, H
		case r < 121:
			return `l`, H
		case r < 126:
			return `a`, H
		case r < 128:
			return `g`, H
		case r < 129:
			return `o`, H
		case r < 130:
			return `m`, H
		case r < 131:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `og`:
		H := 3.2145272645029523
		r := rng.IntN(60)
		switch {
		case r < 10:
			return `y`, H
		case r < 18:
			return `r`, H
		case r < 26:
			return `g`, H
		case r < 33:
			return `e`, H
		case r < 40:
			return `a`, H
		case r < 46:
			return `i`, H
		case r < 52:
			return `u`, H
		case r < 56:
			return `n`, H
		case r < 58:
			return `l`, H
		case r < 59:
			return `w`, H
		case r < 60:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `bi`:
		H := 3.167327740603032
		r := rng.IntN(72)
		switch {
		case r < 15:
			return `l`, H
		case r < 29:
			return `t`, H
		case r < 41:
			return `n`, H
		case r < 48:
			return `d`, H
		case r < 53:
			return `c`, H
		case r < 57:
			return `e`, H
		case r < 61:
			return `r`, H
		case r < 65:
			return `a`, H
		case r < 67:
			return `o`, H
		case r < 69:
			return `s`, H
		case r < 71:
			return `g`, H
		case r < 72:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `em`:
		H := 2.909604688717815
		r := rng.IntN(135)
		switch {
		case r < 28:
			return `i`, H
		case r < 50:
			return `e`, H
		case r < 72:
			return `p`, H
		case r < 92:
			return `o`, H
		case r < 111:
			return `b`, H
		case r < 126:
			return `a`, H
		case r < 128:
			return `y`, H
		case r < 130:
			return `l`, H
		case r < 132:
			return `u`, H
		case r < 133:
			return `n`, H
		case r < 134:
			return `s`, H
		case r < 135:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `au`:
		H := 3.046087347807784
		r := rng.IntN(69)
		switch {
		case r < 15:
			return `t`, H
		case r < 28:
			return `n`, H
		case r < 39:
			return `d`, H
		case r < 48:
			return `s`, H
		case r < 55:
			return `c`, H
		case r < 60:
			return `g`, H
		case r < 63:
			return `r`, H
		case r < 65:
			return `l`, H
		case r < 66:
			return `p`, H
		case r < 67:
			return `z`, H
		case r < 68:
			return `v`, H
		case r < 69:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `aw`:
		H := 2.978305516512272
		r := rng.IntN(35)
		switch {
		case r < 12:
			return `a`, H
		case r < 18:
			return `l`, H
		case r < 22:
			return `n`, H
		case r < 25:
			return `e`, H
		case r < 27:
			return `f`, H
		case r < 29:
			return `k`, H
		case r < 30:
			return `d`, H
		case r < 31:
			return `o`, H
		case r < 32:
			return `h`, H
		case r < 33:
			return `s`, H
		case r < 34:
			return `i`, H
		case r < 35:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ew`:
		H := 2.8625074707625155
		r := rng.IntN(44)
		switch {
		case r < 16:
			return `a`, H
		case r < 24:
			return `i`, H
		case r < 30:
			return `e`, H
		case r < 33:
			return `o`, H
		case r < 35:
			return `r`, H
		case r < 37:
			return `m`, H
		case r < 39:
			return `l`, H
		case r < 40:
			return `n`, H
		case r < 41:
			return `d`, H
		case r < 42:
			return `h`, H
		case r < 43:
			return `y`, H
		case r < 44:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ec`:
		H := 3.0339701786495312
		r := rng.IntN(189)
		switch {
		case r < 64:
			return `t`, H
		case r < 90:
			return `o`, H
		case r < 106:
			return `e`, H
		case r < 122:
			return `i`, H
		case r < 135:
			return `k`, H
		case r < 148:
			return `a`, H
		case r < 161:
			return `l`, H
		case r < 172:
			return `u`, H
		case r < 179:
			return `r`, H
		case r < 184:
			return `h`, H
		case r < 188:
			return `y`, H
		case r < 189:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `oc`:
		H := 2.8369155250951916
		r := rng.IntN(91)
		switch {
		case r < 37:
			return `k`, H
		case r < 49:
			return `a`, H
		case r < 57:
			return `c`, H
		case r < 64:
			return `r`, H
		case r < 70:
			return `e`, H
		case r < 76:
			return `i`, H
		case r < 81:
			return `t`, H
		case r < 85:
			return `u`, H
		case r < 87:
			return `o`, H
		case r < 89:
			return `h`, H
		case r < 90:
			return `y`, H
		case r < 91:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `mu`:
		H := 3.0506400538703033
		r := rng.IntN(80)
		switch {
		case r < 22:
			return `s`, H
		case r < 36:
			return `l`, H
		case r < 47:
			return `t`, H
		case r < 55:
			return `m`, H
		case r < 61:
			return `r`, H
		case r < 66:
			return `n`, H
		case r < 71:
			return `g`, H
		case r < 74:
			return `d`, H
		case r < 77:
			return `c`, H
		case r < 78:
			return `f`, H
		case r < 79:
			return `p`, H
		case r < 80:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `uc`:
		H := 2.8631704063613026
		r := rng.IntN(87)
		switch {
		case r < 31:
			return `k`, H
		case r < 44:
			return `h`, H
		case r < 57:
			return `t`, H
		case r < 66:
			return `e`, H
		case r < 71:
			return `i`, H
		case r < 75:
			return `a`, H
		case r < 78:
			return `c`, H
		case r < 81:
			return `l`, H
		case r < 83:
			return `o`, H
		case r < 85:
			return `u`, H
		case r < 86:
			return `y`, H
		case r < 87:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `os`:
		H := 2.8544000807530825
		r := rng.IntN(128)
		switch {
		case r < 33:
			return `e`, H
		case r < 66:
			return `t`, H
		case r < 84:
			return `s`, H
		case r < 101:
			return `i`, H
		case r < 108:
			return `a`, H
		case r < 112:
			return `p`, H
		case r < 116:
			return `m`, H
		case r < 120:
			return `u`, H
		case r < 123:
			return `h`, H
		case r < 126:
			return `y`, H
		case r < 127:
			return `o`, H
		case r < 128:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `xi`:
		H := 3.1329440449809582
		r := rng.IntN(26)
		switch {
		case r < 8:
			return `d`, H
		case r < 12:
			return `s`, H
		case r < 15:
			return `m`, H
		case r < 17:
			return `n`, H
		case r < 19:
			return `c`, H
		case r < 20:
			return `f`, H
		case r < 21:
			return `o`, H
		case r < 22:
			return `r`, H
		case r < 23:
			return `t`, H
		case r < 24:
			return `a`, H
		case r < 25:
			return `g`, H
		case r < 26:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `ju`:
		H := 3.2165060421774707
		r := rng.IntN(42)
		switch {
		case r < 8:
			return `n`, H
		case r < 16:
			return `r`, H
		case r < 22:
			return `s`, H
		case r < 27:
			return `d`, H
		case r < 30:
			return `i`, H
		case r < 33:
			return `m`, H
		case r < 36:
			return `g`, H
		case r < 38:
			return `b`, H
		case r < 39:
			return `j`, H
		case r < 40:
			return `k`, H
		case r < 41:
			return `v`, H
		case r < 42:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `th`:
		H := 2.880261933124179
		r := rng.IntN(129)
		switch {
		case r < 34:
			return `e`, H
		case r < 64:
			return `i`, H
		case r < 83:
			return `r`, H
		case r < 98:
			return `o`, H
		case r < 107:
			return `a`, H
		case r < 114:
			return `y`, H
		case r < 119:
			return `l`, H
		case r < 124:
			return `u`, H
		case r < 126:
			return `w`, H
		case r < 127:
			return `p`, H
		case r < 128:
			return `m`, H
		case r < 129:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `cu`:
		H := 2.8392825556691563
		r := rng.IntN(124)
		switch {
		case r < 38:
			return `r`, H
		case r < 64:
			return `l`, H
		case r < 84:
			return `s`, H
		case r < 96:
			return `p`, H
		case r < 105:
			return `t`, H
		case r < 110:
			return `b`, H
		case r < 113:
			return `f`, H
		case r < 116:
			return `d`, H
		case r < 119:
			return `e`, H
		case r < 122:
			return `m`, H
		case r < 123:
			return `c`, H
		case r < 124:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `ld`:
		H := 3.133525936753556
		r := rng.IntN(33)
		switch {
		case r < 9:
			return `e`, H
		case r < 15:
			return `l`, H
		case r < 18:
			return `f`, H
		case r < 21:
			return `o`, H
		case r < 24:
			return `i`, H
		case r < 27:
			return `c`, H
		case r < 28:
			return `n`, H
		case r < 29:
			return `h`, H
		case r < 30:
			return `s`, H
		case r < 31:
			return `y`, H
		case r < 32:
			return `a`, H
		case r < 33:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ki`:
		H := 2.1171650938789814
		r := rng.IntN(129)
		switch {
		case r < 79:
			return `n`, H
		case r < 96:
			return `l`, H
		case r < 102:
			return `m`, H
		case r < 107:
			return `e`, H
		case r < 112:
			return `s`, H
		case r < 117:
			return `t`, H
		case r < 120:
			return `r`, H
		case r < 123:
			return `p`, H
		case r < 125:
			return `d`, H
		case r < 127:
			return `c`, H
		case r < 128:
			return `w`, H
		case r < 129:
			return `i`, H
		default:
			panic("unexpected rand num")
		}
	case `i`:
		H := 2.5497605472988316
		r := rng.IntN(115)
		switch {
		case r < 59:
			return `m`, H
		case r < 71:
			return `d`, H
		case r < 80:
			return `r`, H
		case r < 88:
			return `s`, H
		case r < 94:
			return `c`, H
		case r < 99:
			return `t`, H
		case r < 103:
			return `g`, H
		case r < 106:
			return `o`, H
		case r < 109:
			return `p`, H
		case r < 112:
			return `l`, H
		case r < 114:
			return `v`, H
		case r < 115:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ui`:
		H := 3.141352734607186
		r := rng.IntN(73)
		switch {
		case r < 17:
			return `t`, H
		case r < 27:
			return `s`, H
		case r < 36:
			return `n`, H
		case r < 44:
			return `r`, H
		case r < 52:
			return `c`, H
		case r < 59:
			return `d`, H
		case r < 66:
			return `l`, H
		case r < 68:
			return `v`, H
		case r < 70:
			return `g`, H
		case r < 71:
			return `e`, H
		case r < 72:
			return `p`, H
		case r < 73:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `om`:
		H := 2.9118142888347975
		r := rng.IntN(122)
		switch {
		case r < 29:
			return `p`, H
		case r < 51:
			return `e`, H
		case r < 72:
			return `i`, H
		case r < 91:
			return `a`, H
		case r < 104:
			return `m`, H
		case r < 110:
			return `b`, H
		case r < 114:
			return `y`, H
		case r < 117:
			return `o`, H
		case r < 119:
			return `f`, H
		case r < 120:
			return `n`, H
		case r < 121:
			return `r`, H
		case r < 122:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ep`:
		H := 3.292881430273992
		r := rng.IntN(93)
		switch {
		case r < 16:
			return `t`, H
		case r < 29:
			return `r`, H
		case r < 41:
			return `i`, H
		case r < 52:
			return `a`, H
		case r < 63:
			return `l`, H
		case r < 71:
			return `u`, H
		case r < 78:
			return `e`, H
		case r < 84:
			return `o`, H
		case r < 87:
			return `h`, H
		case r < 90:
			return `s`, H
		case r < 92:
			return `p`, H
		case r < 93:
			return `n`, H
		default:
			panic("unexpected rand num")
		}
	case `ic`:
		H := 3.0986089299037656
		r := rng.IntN(198)
		switch {
		case r < 41:
			return `a`, H
		case r < 78:
			return `k`, H
		case r < 107:
			return `e`, H
		case r < 134:
			return `i`, H
		case r < 151:
			return `s`, H
		case r < 166:
			return `t`, H
		case r < 179:
			return `o`, H
		case r < 184:
			return `y`, H
		case r < 189:
			return `l`, H
		case r < 193:
			return `h`, H
		case r < 197:
			return `u`, H
		case r < 198:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `us`:
		H := 2.823674243280801
		r := rng.IntN(187)
		switch {
		case r < 52:
			return `t`, H
		case r < 93:
			return `e`, H
		case r < 126:
			return `h`, H
		case r < 148:
			return `i`, H
		case r < 162:
			return `a`, H
		case r < 171:
			return `k`, H
		case r < 175:
			return `p`, H
		case r < 178:
			return `s`, H
		case r < 180:
			return `y`, H
		case r < 182:
			return `l`, H
		case r < 184:
			return `u`, H
		case r < 186:
			return `b`, H
		case r < 187:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ja`:
		H := 3.494680368408909
		r := rng.IntN(28)
		switch {
		case r < 4:
			return `w`, H
		case r < 8:
			return `i`, H
		case r < 12:
			return `c`, H
		case r < 15:
			return `r`, H
		case r < 17:
			return `n`, H
		case r < 19:
			return `y`, H
		case r < 21:
			return `m`, H
		case r < 23:
			return `u`, H
		case r < 24:
			return `s`, H
		case r < 25:
			return `z`, H
		case r < 26:
			return `v`, H
		case r < 27:
			return `l`, H
		case r < 28:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `rm`:
		H := 2.9183678050834025
		r := rng.IntN(67)
		switch {
		case r < 17:
			return `a`, H
		case r < 31:
			return `o`, H
		case r < 42:
			return `e`, H
		case r < 53:
			return `i`, H
		case r < 57:
			return `l`, H
		case r < 59:
			return `f`, H
		case r < 61:
			return `u`, H
		case r < 62:
			return `h`, H
		case r < 63:
			return `y`, H
		case r < 64:
			return `r`, H
		case r < 65:
			return `c`, H
		case r < 66:
			return `p`, H
		case r < 67:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `eg`:
		H := 3.2163328922854326
		r := rng.IntN(70)
		switch {
		case r < 17:
			return `a`, H
		case r < 27:
			return `i`, H
		case r < 36:
			return `g`, H
		case r < 44:
			return `r`, H
		case r < 51:
			return `o`, H
		case r < 57:
			return `u`, H
		case r < 60:
			return `e`, H
		case r < 63:
			return `l`, H
		case r < 65:
			return `n`, H
		case r < 67:
			return `w`, H
		case r < 68:
			return `y`, H
		case r < 69:
			return `m`, H
		case r < 70:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ua`:
		H := 2.9201179414976686
		r := rng.IntN(77)
		switch {
		case r < 25:
			return `l`, H
		case r < 40:
			return `t`, H
		case r < 52:
			return `r`, H
		case r < 57:
			return `d`, H
		case r < 62:
			return `b`, H
		case r < 66:
			return `n`, H
		case r < 69:
			return `i`, H
		case r < 71:
			return `k`, H
		case r < 73:
			return `c`, H
		case r < 74:
			return `h`, H
		case r < 75:
			return `s`, H
		case r < 76:
			return `g`, H
		case r < 77:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `oa`:
		H := 3.0807669557620305
		r := rng.IntN(86)
		switch {
		case r < 18:
			return `d`, H
		case r < 36:
			return `t`, H
		case r < 49:
			return `s`, H
		case r < 62:
			return `r`, H
		case r < 68:
			return `c`, H
		case r < 73:
			return `k`, H
		case r < 76:
			return `n`, H
		case r < 79:
			return `m`, H
		case r < 82:
			return `l`, H
		case r < 83:
			return `f`, H
		case r < 84:
			return `g`, H
		case r < 85:
			return `u`, H
		case r < 86:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ip`:
		H := 2.9084361115661546
		r := rng.IntN(58)
		switch {
		case r < 20:
			return `p`, H
		case r < 33:
			return `e`, H
		case r < 39:
			return `t`, H
		case r < 42:
			return `o`, H
		case r < 45:
			return `h`, H
		case r < 48:
			return `s`, H
		case r < 51:
			return `l`, H
		case r < 53:
			return `a`, H
		case r < 54:
			return `f`, H
		case r < 55:
			return `i`, H
		case r < 56:
			return `c`, H
		case r < 57:
			return `m`, H
		case r < 58:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `id`:
		H := 2.437836737262423
		r := rng.IntN(149)
		switch {
		case r < 78:
			return `e`, H
		case r < 98:
			return `i`, H
		case r < 109:
			return `d`, H
		case r < 119:
			return `a`, H
		case r < 128:
			return `g`, H
		case r < 135:
			return `l`, H
		case r < 138:
			return `o`, H
		case r < 141:
			return `y`, H
		case r < 144:
			return `u`, H
		case r < 146:
			return `n`, H
		case r < 147:
			return `s`, H
		case r < 148:
			return `t`, H
		case r < 149:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `hu`:
		H := 3.2003975417055566
		r := rng.IntN(80)
		switch {
		case r < 21:
			return `m`, H
		case r < 35:
			return `n`, H
		case r < 47:
			return `r`, H
		case r < 54:
			return `s`, H
		case r < 58:
			return `f`, H
		case r < 62:
			return `d`, H
		case r < 66:
			return `t`, H
		case r < 69:
			return `c`, H
		case r < 72:
			return `l`, H
		case r < 74:
			return `a`, H
		case r < 76:
			return `p`, H
		case r < 78:
			return `g`, H
		case r < 80:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `gu`:
		H := 3.2352429992546594
		r := rng.IntN(69)
		switch {
		case r < 13:
			return `l`, H
		case r < 24:
			return `i`, H
		case r < 33:
			return `r`, H
		case r < 41:
			return `e`, H
		case r < 48:
			return `s`, H
		case r < 55:
			return `a`, H
		case r < 61:
			return `m`, H
		case r < 64:
			return `t`, H
		case r < 65:
			return `n`, H
		case r < 66:
			return `o`, H
		case r < 67:
			return `y`, H
		case r < 68:
			return `p`, H
		case r < 69:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `pu`:
		H := 2.953209386359383
		r := rng.IntN(89)
		switch {
		case r < 26:
			return `r`, H
		case r < 44:
			return `l`, H
		case r < 56:
			return `s`, H
		case r < 67:
			return `t`, H
		case r < 74:
			return `n`, H
		case r < 77:
			return `p`, H
		case r < 80:
			return `m`, H
		case r < 82:
			return `z`, H
		case r < 84:
			return `g`, H
		case r < 86:
			return `b`, H
		case r < 87:
			return `d`, H
		case r < 88:
			return `e`, H
		case r < 89:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `im`:
		H := 2.9007368337224446
		r := rng.IntN(181)
		switch {
		case r < 55:
			return `p`, H
		case r < 86:
			return `m`, H
		case r < 114:
			return `e`, H
		case r < 138:
			return `a`, H
		case r < 154:
			return `i`, H
		case r < 160:
			return `b`, H
		case r < 165:
			return `o`, H
		case r < 170:
			return `u`, H
		case r < 173:
			return `n`, H
		case r < 176:
			return `s`, H
		case r < 178:
			return `y`, H
		case r < 180:
			return `l`, H
		case r < 181:
			return `w`, H
		default:
			panic("unexpected rand num")
		}
	case `nt`:
		H := 3.0811605328891254
		r := rng.IntN(219)
		switch {
		case r < 52:
			return `i`, H
		case r < 98:
			return `e`, H
		case r < 127:
			return `a`, H
		case r < 152:
			return `r`, H
		case r < 168:
			return `l`, H
		case r < 180:
			return `o`, H
		case r < 192:
			return `h`, H
		case r < 202:
			return `y`, H
		case r < 209:
			return `u`, H
		case r < 215:
			return `s`, H
		case r < 217:
			return `w`, H
		case r < 218:
			return `d`, H
		case r < 219:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `pi`:
		H := 2.9688147163910608
		r := rng.IntN(129)
		switch {
		case r < 49:
			return `n`, H
		case r < 65:
			return `l`, H
		case r < 79:
			return `r`, H
		case r < 90:
			return `e`, H
		case r < 98:
			return `c`, H
		case r < 104:
			return `d`, H
		case r < 110:
			return `s`, H
		case r < 115:
			return `o`, H
		case r < 120:
			return `t`, H
		case r < 124:
			return `a`, H
		case r < 127:
			return `p`, H
		case r < 128:
			return `f`, H
		case r < 129:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `mp`:
		H := 3.2580382343434384
		r := rng.IntN(143)
		switch {
		case r < 29:
			return `l`, H
		case r < 49:
			return `o`, H
		case r < 68:
			return `e`, H
		case r < 83:
			return `i`, H
		case r < 98:
			return `a`, H
		case r < 110:
			return `r`, H
		case r < 121:
			return `t`, H
		case r < 129:
			return `h`, H
		case r < 136:
			return `u`, H
		case r < 140:
			return `s`, H
		case r < 141:
			return `n`, H
		case r < 142:
			return `f`, H
		case r < 143:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `st`:
		H := 2.956505271065868
		r := rng.IntN(496)
		switch {
		case r < 106:
			return `a`, H
		case r < 200:
			return `e`, H
		case r < 286:
			return `i`, H
		case r < 356:
			return `r`, H
		case r < 413:
			return `o`, H
		case r < 445:
			return `u`, H
		case r < 467:
			return `y`, H
		case r < 478:
			return `l`, H
		case r < 483:
			return `n`, H
		case r < 488:
			return `f`, H
		case r < 492:
			return `b`, H
		case r < 494:
			return `w`, H
		case r < 496:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ie`:
		H := 3.0687146768399054
		r := rng.IntN(136)
		switch {
		case r < 37:
			return `d`, H
		case r < 64:
			return `r`, H
		case r < 83:
			return `n`, H
		case r < 92:
			return `s`, H
		case r < 101:
			return `v`, H
		case r < 110:
			return `l`, H
		case r < 118:
			return `w`, H
		case r < 124:
			return `t`, H
		case r < 129:
			return `f`, H
		case r < 133:
			return `c`, H
		case r < 134:
			return `k`, H
		case r < 135:
			return `m`, H
		case r < 136:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rs`:
		H := 2.9825091489647586
		r := rng.IntN(79)
		switch {
		case r < 23:
			return `e`, H
		case r < 37:
			return `t`, H
		case r < 49:
			return `i`, H
		case r < 57:
			return `h`, H
		case r < 64:
			return `u`, H
		case r < 69:
			return `a`, H
		case r < 72:
			return `o`, H
		case r < 74:
			return `l`, H
		case r < 75:
			return `n`, H
		case r < 76:
			return `w`, H
		case r < 77:
			return `d`, H
		case r < 78:
			return `p`, H
		case r < 79:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `pe`:
		H := 2.860807140148557
		r := rng.IntN(243)
		switch {
		case r < 85:
			return `r`, H
		case r < 130:
			return `n`, H
		case r < 163:
			return `d`, H
		case r < 180:
			return `t`, H
		case r < 196:
			return `c`, H
		case r < 210:
			return `a`, H
		case r < 223:
			return `l`, H
		case r < 230:
			return `s`, H
		case r < 235:
			return `e`, H
		case r < 237:
			return `z`, H
		case r < 239:
			return `g`, H
		case r < 241:
			return `b`, H
		case r < 242:
			return `w`, H
		case r < 243:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `rn`:
		H := 2.9893847578098818
		r := rng.IntN(49)
		switch {
		case r < 14:
			return `e`, H
		case r < 25:
			return `i`, H
		case r < 34:
			return `a`, H
		case r < 37:
			return `f`, H
		case r < 39:
			return `o`, H
		case r < 41:
			return `b`, H
		case r < 42:
			return `n`, H
		case r < 43:
			return `h`, H
		case r < 44:
			return `s`, H
		case r < 45:
			return `y`, H
		case r < 46:
			return `c`, H
		case r < 47:
			return `m`, H
		case r < 48:
			return `l`, H
		case r < 49:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `bu`:
		H := 3.292370012087556
		r := rng.IntN(89)
		switch {
		case r < 24:
			return `l`, H
		case r < 37:
			return `n`, H
		case r < 46:
			return `c`, H
		case r < 54:
			return `s`, H
		case r < 61:
			return `r`, H
		case r < 67:
			return `f`, H
		case r < 72:
			return `d`, H
		case r < 77:
			return `t`, H
		case r < 80:
			return `i`, H
		case r < 83:
			return `g`, H
		case r < 86:
			return `b`, H
		case r < 87:
			return `p`, H
		case r < 88:
			return `m`, H
		case r < 89:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `it`:
		H := 3.0957275161189104
		r := rng.IntN(266)
		switch {
		case r < 75:
			return `y`, H
		case r < 118:
			return `e`, H
		case r < 156:
			return `a`, H
		case r < 184:
			return `i`, H
		case r < 207:
			return `t`, H
		case r < 221:
			return `u`, H
		case r < 234:
			return `c`, H
		case r < 245:
			return `o`, H
		case r < 250:
			return `h`, H
		case r < 255:
			return `r`, H
		case r < 259:
			return `z`, H
		case r < 263:
			return `l`, H
		case r < 265:
			return `s`, H
		case r < 266:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ig`:
		H := 2.973794367913592
		r := rng.IntN(109)
		switch {
		case r < 42:
			return `h`, H
		case r < 55:
			return `a`, H
		case r < 66:
			return `n`, H
		case r < 75:
			return `g`, H
		case r < 82:
			return `e`, H
		case r < 89:
			return `u`, H
		case r < 95:
			return `o`, H
		case r < 100:
			return `i`, H
		case r < 102:
			return `r`, H
		case r < 104:
			return `m`, H
		case r < 106:
			return `l`, H
		case r < 107:
			return `s`, H
		case r < 108:
			return `y`, H
		case r < 109:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ch`:
		H := 2.7422368334249807
		r := rng.IntN(197)
		switch {
		case r < 56:
			return `e`, H
		case r < 109:
			return `a`, H
		case r < 144:
			return `i`, H
		case r < 162:
			return `o`, H
		case r < 172:
			return `u`, H
		case r < 178:
			return `y`, H
		case r < 183:
			return `l`, H
		case r < 188:
			return `b`, H
		case r < 190:
			return `n`, H
		case r < 192:
			return `r`, H
		case r < 194:
			return `m`, H
		case r < 195:
			return `w`, H
		case r < 196:
			return `t`, H
		case r < 197:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ac`:
		H := 3.038335326308152
		r := rng.IntN(278)
		switch {
		case r < 81:
			return `k`, H
		case r < 133:
			return `t`, H
		case r < 166:
			return `e`, H
		case r < 199:
			return `h`, H
		case r < 225:
			return `i`, H
		case r < 235:
			return `a`, H
		case r < 244:
			return `c`, H
		case r < 252:
			return `y`, H
		case r < 259:
			return `o`, H
		case r < 266:
			return `u`, H
		case r < 272:
			return `r`, H
		case r < 275:
			return `q`, H
		case r < 277:
			return `l`, H
		case r < 278:
			return `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ce`:
		H := 3.3511760926287546
		r := rng.IntN(88)
		switch {
		case r < 17:
			return `r`, H
		case r < 30:
			return `n`, H
		case r < 43:
			return `d`, H
		case r < 54:
			return `l`, H
		case r < 62:
			return `s`, H
		case r < 68:
			return `p`, H
		case r < 72:
			return `i`, H
		case r < 76:
			return `a`, H
		case r < 79:
			return `e`, H
		case r < 82:
			return `t`, H
		case r < 85:
			return `m`, H
		case r < 86:
			return `f`, H
		case r < 87:
			return `c`, H
		case r < 88:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ol`:
		H := 3.263456186229182
		r := rng.IntN(174)
		switch {
		case r < 31:
			return `o`, H
		case r < 58:
			return `l`, H
		case r < 83:
			return `d`, H
		case r < 105:
			return `e`, H
		case r < 125:
			return `i`, H
		case r < 143:
			return `a`, H
		case r < 150:
			return `y`, H
		case r < 157:
			return `v`, H
		case r < 163:
			return `t`, H
		case r < 167:
			return `u`, H
		case r < 170:
			return `k`, H
		case r < 172:
			return `f`, H
		case r < 173:
			return `s`, H
		case r < 174:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `su`:
		H := 3.1629317012719462
		r := rng.IntN(166)
		switch {
		case r < 45:
			return `b`, H
		case r < 80:
			return `r`, H
		case r < 98:
			return `p`, H
		case r < 110:
			return `l`, H
		case r < 121:
			return `a`, H
		case r < 131:
			return `i`, H
		case r < 140:
			return `s`, H
		case r < 147:
			return `m`, H
		case r < 153:
			return `f`, H
		case r < 157:
			return `e`, H
		case r < 160:
			return `d`, H
		case r < 163:
			return `c`, H
		case r < 165:
			return `g`, H
		case r < 166:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `fa`:
		H := 3.2099575789421095
		r := rng.IntN(80)
		switch {
		case r < 23:
			return `c`, H
		case r < 36:
			return `n`, H
		case r < 45:
			return `l`, H
		case r < 52:
			return `s`, H
		case r < 58:
			return `m`, H
		case r < 63:
			return `v`, H
		case r < 67:
			return `i`, H
		case r < 70:
			return `r`, H
		case r < 73:
			return `b`, H
		case r < 75:
			return `d`, H
		case r < 77:
			return `u`, H
		case r < 78:
			return `t`, H
		case r < 79:
			return `z`, H
		case r < 80:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ab`:
		H := 2.0438133422038423
		r := rng.IntN(220)
		switch {
		case r < 143:
			return `l`, H
		case r < 158:
			return `i`, H
		case r < 171:
			return `o`, H
		case r < 184:
			return `b`, H
		case r < 194:
			return `s`, H
		case r < 202:
			return `r`, H
		case r < 207:
			return `a`, H
		case r < 211:
			return `e`, H
		case r < 214:
			return `d`, H
		case r < 216:
			return `y`, H
		case r < 217:
			return `n`, H
		case r < 218:
			return `m`, H
		case r < 219:
			return `g`, H
		case r < 220:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `fi`:
		H := 3.287172025028723
		r := rng.IntN(141)
		switch {
		case r < 34:
			return `n`, H
		case r < 57:
			return `e`, H
		case r < 75:
			return `l`, H
		case r < 88:
			return `s`, H
		case r < 99:
			return `d`, H
		case r < 108:
			return `r`, H
		case r < 115:
			return `c`, H
		case r < 122:
			return `g`, H
		case r < 128:
			return `t`, H
		case r < 133:
			return `x`, H
		case r < 137:
			return `f`, H
		case r < 139:
			return `a`, H
		case r < 140:
			return `v`, H
		case r < 141:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `et`:
		H := 3.247370579505989
		r := rng.IntN(156)
		switch {
		case r < 32:
			return `e`, H
		case r < 60:
			return `i`, H
		case r < 83:
			return `t`, H
		case r < 98:
			return `h`, H
		case r < 112:
			return `r`, H
		case r < 125:
			return `a`, H
		case r < 136:
			return `o`, H
		case r < 141:
			return `y`, H
		case r < 146:
			return `u`, H
		case r < 150:
			return `c`, H
		case r < 152:
			return `f`, H
		case r < 153:
			return `d`, H
		case r < 154:
			return `s`, H
		case r < 155:
			return `z`, H
		case r < 156:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `te`:
		H := 2.3840358108683493
		r := rng.IntN(383)
		switch {
		case r < 168:
			return `r`, H
		case r < 271:
			return `d`, H
		case r < 318:
			return `n`, H
		case r < 336:
			return `e`, H
		case r < 344:
			return `m`, H
		case r < 351:
			return `a`, H
		case r < 357:
			return `l`, H
		case r < 362:
			return `x`, H
		case r < 366:
			return `s`, H
		case r < 370:
			return `c`, H
		case r < 374:
			return `p`, H
		case r < 378:
			return `g`, H
		case r < 380:
			return `t`, H
		case r < 382:
			return `b`, H
		case r < 383:
			return `w`, H
		default:
			panic("unexpected rand num")
		}
	case `ir`:
		H := 3.203988838355143
		r := rng.IntN(110)
		switch {
		case r < 39:
			return `e`, H
		case r < 52:
			return `i`, H
		case r < 64:
			return `t`, H
		case r < 71:
			return `r`, H
		case r < 77:
			return `c`, H
		case r < 82:
			return `s`, H
		case r < 87:
			return `a`, H
		case r < 92:
			return `m`, H
		case r < 96:
			return `d`, H
		case r < 99:
			return `k`, H
		case r < 102:
			return `u`, H
		case r < 105:
			return `l`, H
		case r < 107:
			return `y`, H
		case r < 109:
			return `p`, H
		case r < 110:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `s`:
		H := 3.614984946742652
		r := rng.IntN(1087)
		switch {
		case r < 207:
			return `t`, H
		case r < 322:
			return `h`, H
		case r < 437:
			return `u`, H
		case r < 541:
			return `p`, H
		case r < 633:
			return `a`, H
		case r < 718:
			return `c`, H
		case r < 790:
			return `e`, H
		case r < 847:
			return `l`, H
		case r < 902:
			return `i`, H
		case r < 947:
			return `n`, H
		case r < 985:
			return `k`, H
		case r < 1020:
			return `w`, H
		case r < 1050:
			return `m`, H
		case r < 1073:
			return `q`, H
		case r < 1087:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `wi`:
		H := 3.2044098392197946
		r := rng.IntN(112)
		switch {
		case r < 33:
			return `n`, H
		case r < 51:
			return `l`, H
		case r < 65:
			return `s`, H
		case r < 75:
			return `r`, H
		case r < 83:
			return `d`, H
		case r < 89:
			return `m`, H
		case r < 94:
			return `f`, H
		case r < 99:
			return `t`, H
		case r < 102:
			return `e`, H
		case r < 104:
			return `c`, H
		case r < 106:
			return `p`, H
		case r < 108:
			return `z`, H
		case r < 110:
			return `g`, H
		case r < 111:
			return `k`, H
		case r < 112:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ba`:
		H := 3.20045667779631
		r := rng.IntN(199)
		switch {
		case r < 48:
			return `c`, H
		case r < 77:
			return `r`, H
		case r < 103:
			return `n`, H
		case r < 127:
			return `t`, H
		case r < 146:
			return `l`, H
		case r < 162:
			return `g`, H
		case r < 175:
			return `s`, H
		case r < 182:
			return `b`, H
		case r < 188:
			return `k`, H
		case r < 193:
			return `d`, H
		case r < 195:
			return `f`, H
		case r < 196:
			return `y`, H
		case r < 197:
			return `m`, H
		case r < 198:
			return `z`, H
		case r < 199:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rt`:
		H := 3.4318639824861625
		r := rng.IntN(122)
		switch {
		case r < 24:
			return `e`, H
		case r < 42:
			return `a`, H
		case r < 59:
			return `i`, H
		case r < 71:
			return `h`, H
		case r < 83:
			return `l`, H
		case r < 91:
			return `u`, H
		case r < 97:
			return `o`, H
		case r < 103:
			return `y`, H
		case r < 108:
			return `n`, H
		case r < 113:
			return `s`, H
		case r < 115:
			return `w`, H
		case r < 117:
			return `r`, H
		case r < 119:
			return `c`, H
		case r < 121:
			return `z`, H
		case r < 122:
			return `f`, H
		default:
			panic("unexpected rand num")
		}
	case `hi`:
		H := 2.918107201430624
		r := rng.IntN(129)
		switch {
		case r < 54:
			return `n`, H
		case r < 70:
			return `l`, H
		case r < 80:
			return `r`, H
		case r < 90:
			return `c`, H
		case r < 99:
			return `p`, H
		case r < 106:
			return `f`, H
		case r < 112:
			return `e`, H
		case r < 117:
			return `m`, H
		case r < 120:
			return `s`, H
		case r < 123:
			return `v`, H
		case r < 125:
			return `t`, H
		case r < 126:
			return `d`, H
		case r < 127:
			return `h`, H
		case r < 128:
			return `a`, H
		case r < 129:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `fo`:
		H := 2.981045025728019
		r := rng.IntN(86)
		switch {
		case r < 26:
			return `o`, H
		case r < 45:
			return `r`, H
		case r < 58:
			return `l`, H
		case r < 65:
			return `u`, H
		case r < 70:
			return `n`, H
		case r < 74:
			return `c`, H
		case r < 76:
			return `s`, H
		case r < 78:
			return `i`, H
		case r < 80:
			return `g`, H
		case r < 81:
			return `w`, H
		case r < 82:
			return `d`, H
		case r < 83:
			return `e`, H
		case r < 84:
			return `y`, H
		case r < 85:
			return `a`, H
		case r < 86:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ag`:
		H := 2.8283474600883034
		r := rng.IntN(179)
		switch {
		case r < 76:
			return `e`, H
		case r < 99:
			return `g`, H
		case r < 115:
			return `i`, H
		case r < 131:
			return `r`, H
		case r < 144:
			return `n`, H
		case r < 154:
			return `o`, H
		case r < 160:
			return `u`, H
		case r < 165:
			return `s`, H
		case r < 169:
			return `a`, H
		case r < 173:
			return `m`, H
		case r < 175:
			return `p`, H
		case r < 176:
			return `f`, H
		case r < 177:
			return `w`, H
		case r < 178:
			return `h`, H
		case r < 179:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `ll`:
		H := 2.9528008740663543
		r := rng.IntN(137)
		switch {
		case r < 34:
			return `y`, H
		case r < 64:
			return `e`, H
		case r < 93:
			return `i`, H
		case r < 106:
			return `a`, H
		case r < 116:
			return `o`, H
		case r < 120:
			return `f`, H
		case r < 123:
			return `s`, H
		case r < 126:
			return `u`, H
		case r < 128:
			return `n`, H
		case r < 130:
			return `d`, H
		case r < 132:
			return `h`, H
		case r < 134:
			return `p`, H
		case r < 135:
			return `w`, H
		case r < 136:
			return `r`, H
		case r < 137:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ru`:
		H := 3.385847293011174
		r := rng.IntN(148)
		switch {
		case r < 33:
			return `s`, H
		case r < 57:
			return `n`, H
		case r < 76:
			return `m`, H
		case r < 89:
			return `d`, H
		case r < 101:
			return `b`, H
		case r < 112:
			return `c`, H
		case r < 119:
			return `t`, H
		case r < 125:
			return `p`, H
		case r < 130:
			return `f`, H
		case r < 135:
			return `e`, H
		case r < 140:
			return `g`, H
		case r < 143:
			return `i`, H
		case r < 146:
			return `l`, H
		case r < 147:
			return `r`, H
		case r < 148:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `am`:
		H := 3.1846111780304867
		r := rng.IntN(178)
		switch {
		case r < 40:
			return `e`, H
		case r < 73:
			return `i`, H
		case r < 103:
			return `p`, H
		case r < 122:
			return `b`, H
		case r < 135:
			return `a`, H
		case r < 145:
			return `o`, H
		case r < 155:
			return `m`, H
		case r < 162:
			return `u`, H
		case r < 167:
			return `l`, H
		case r < 170:
			return `s`, H
		case r < 172:
			return `n`, H
		case r < 174:
			return `y`, H
		case r < 176:
			return `r`, H
		case r < 177:
			return `t`, H
		case r < 178:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ge`:
		H := 2.964543201961648
		r := rng.IntN(132)
		switch {
		case r < 34:
			return `r`, H
		case r < 62:
			return `n`, H
		case r < 90:
			return `d`, H
		case r < 100:
			return `t`, H
		case r < 107:
			return `o`, H
		case r < 113:
			return `s`, H
		case r < 118:
			return `e`, H
		case r < 123:
			return `l`, H
		case r < 126:
			return `a`, H
		case r < 127:
			return `f`, H
		case r < 128:
			return `y`, H
		case r < 129:
			return `i`, H
		case r < 130:
			return `c`, H
		case r < 131:
			return `m`, H
		case r < 132:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `as`:
		H := 3.0133424407336613
		r := rng.IntN(280)
		switch {
		case r < 79:
			return `t`, H
		case r < 135:
			return `h`, H
		case r < 172:
			return `s`, H
		case r < 202:
			return `i`, H
		case r < 229:
			return `e`, H
		case r < 242:
			return `p`, H
		case r < 253:
			return `c`, H
		case r < 261:
			return `k`, H
		case r < 265:
			return `o`, H
		case r < 269:
			return `a`, H
		case r < 273:
			return `m`, H
		case r < 276:
			return `u`, H
		case r < 278:
			return `y`, H
		case r < 279:
			return `q`, H
		case r < 280:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `mi`:
		H := 2.8826950013624906
		r := rng.IntN(150)
		switch {
		case r < 54:
			return `n`, H
		case r < 77:
			return `s`, H
		case r < 97:
			return `c`, H
		case r < 115:
			return `t`, H
		case r < 128:
			return `l`, H
		case r < 132:
			return `d`, H
		case r < 135:
			return `f`, H
		case r < 138:
			return `x`, H
		case r < 140:
			return `e`, H
		case r < 142:
			return `a`, H
		case r < 144:
			return `z`, H
		case r < 146:
			return `g`, H
		case r < 148:
			return `u`, H
		case r < 149:
			return `r`, H
		case r < 150:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `rd`:
		H := 3.531539645867806
		r := rng.IntN(52)
		switch {
		case r < 11:
			return `e`, H
		case r < 19:
			return `i`, H
		case r < 24:
			return `r`, H
		case r < 29:
			return `l`, H
		case r < 33:
			return `o`, H
		case r < 36:
			return `w`, H
		case r < 39:
			return `c`, H
		case r < 41:
			return `h`, H
		case r < 43:
			return `s`, H
		case r < 45:
			return `y`, H
		case r < 47:
			return `a`, H
		case r < 49:
			return `u`, H
		case r < 50:
			return `n`, H
		case r < 51:
			return `d`, H
		case r < 52:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ns`:
		H := 3.452447729749734
		r := rng.IntN(128)
		switch {
		case r < 24:
			return `e`, H
		case r < 46:
			return `t`, H
		case r < 60:
			return `i`, H
		case r < 73:
			return `u`, H
		case r < 84:
			return `o`, H
		case r < 94:
			return `h`, H
		case r < 103:
			return `a`, H
		case r < 108:
			return `c`, H
		case r < 113:
			return `p`, H
		case r < 117:
			return `l`, H
		case r < 120:
			return `w`, H
		case r < 123:
			return `m`, H
		case r < 125:
			return `n`, H
		case r < 127:
			return `f`, H
		case r < 128:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `fe`:
		H := 3.3380070663648165
		r := rng.IntN(84)
		switch {
		case r < 20:
			return `r`, H
		case r < 31:
			return `n`, H
		case r < 42:
			return `c`, H
		case r < 50:
			return `d`, H
		case r < 58:
			return `s`, H
		case r < 65:
			return `e`, H
		case r < 70:
			return `m`, H
		case r < 74:
			return `l`, H
		case r < 77:
			return `t`, H
		case r < 79:
			return `a`, H
		case r < 80:
			return `w`, H
		case r < 81:
			return `h`, H
		case r < 82:
			return `i`, H
		case r < 83:
			return `v`, H
		case r < 84:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ve`:
		H := 2.0357494674844125
		r := rng.IntN(277)
		switch {
		case r < 169:
			return `r`, H
		case r < 212:
			return `n`, H
		case r < 229:
			return `d`, H
		case r < 245:
			return `l`, H
		case r < 254:
			return `s`, H
		case r < 262:
			return `t`, H
		case r < 265:
			return `y`, H
		case r < 268:
			return `a`, H
		case r < 270:
			return `i`, H
		case r < 272:
			return `g`, H
		case r < 273:
			return `w`, H
		case r < 274:
			return `e`, H
		case r < 275:
			return `h`, H
		case r < 276:
			return `m`, H
		case r < 277:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ke`:
		H := 3.1007666652190577
		r := rng.IntN(113)
		switch {
		case r < 30:
			return `r`, H
		case r < 54:
			return `d`, H
		case r < 69:
			return `n`, H
		case r < 82:
			return `t`, H
		case r < 88:
			return `l`, H
		case r < 93:
			return `w`, H
		case r < 98:
			return `e`, H
		case r < 103:
			return `s`, H
		case r < 105:
			return `y`, H
		case r < 107:
			return `p`, H
		case r < 109:
			return `b`, H
		case r < 110:
			return `o`, H
		case r < 111:
			return `h`, H
		case r < 112:
			return `m`, H
		case r < 113:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ap`:
		H := 2.8684848816641035
		r := rng.IntN(143)
		switch {
		case r < 60:
			return `p`, H
		case r < 82:
			return `e`, H
		case r < 93:
			return `s`, H
		case r < 103:
			return `t`, H
		case r < 111:
			return `h`, H
		case r < 119:
			return `a`, H
		case r < 125:
			return `i`, H
		case r < 130:
			return `r`, H
		case r < 134:
			return `l`, H
		case r < 137:
			return `o`, H
		case r < 139:
			return `d`, H
		case r < 140:
			return `n`, H
		case r < 141:
			return `k`, H
		case r < 142:
			return `y`, H
		case r < 143:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `gi`:
		H := 3.019193877281908
		r := rng.IntN(102)
		switch {
		case r < 43:
			return `n`, H
		case r < 51:
			return `s`, H
		case r < 59:
			return `l`, H
		case r < 66:
			return `g`, H
		case r < 72:
			return `d`, H
		case r < 78:
			return `c`, H
		case r < 84:
			return `b`, H
		case r < 89:
			return `v`, H
		case r < 92:
			return `e`, H
		case r < 95:
			return `z`, H
		case r < 97:
			return `r`, H
		case r < 99:
			return `m`, H
		case r < 100:
			return `f`, H
		case r < 101:
			return `o`, H
		case r < 102:
			return `a`, H
		default:
			panic("unexpected rand num")
		}
	case `ob`:
		H := 3.2815794566249576
		r := rng.IntN(69)
		switch {
		case r < 14:
			return `s`, H
		case r < 25:
			return `i`, H
		case r < 36:
			return `b`, H
		case r < 45:
			return `e`, H
		case r < 51:
			return `l`, H
		case r < 56:
			return `a`, H
		case r < 60:
			return `t`, H
		case r < 62:
			return `o`, H
		case r < 63:
			return `n`, H
		case r < 64:
			return `w`, H
		case r < 65:
			return `j`, H
		case r < 66:
			return `y`, H
		case r < 67:
			return `c`, H
		case r < 68:
			return `v`, H
		case r < 69:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `vi`:
		H := 3.350995470822189
		r := rng.IntN(189)
		switch {
		case r < 45:
			return `n`, H
		case r < 77:
			return `s`, H
		case r < 94:
			return `d`, H
		case r < 110:
			return `t`, H
		case r < 125:
			return `a`, H
		case r < 138:
			return `o`, H
		case r < 149:
			return `v`, H
		case r < 159:
			return `e`, H
		case r < 168:
			return `c`, H
		case r < 177:
			return `l`, H
		case r < 183:
			return `r`, H
		case r < 186:
			return `g`, H
		case r < 187:
			return `p`, H
		case r < 188:
			return `x`, H
		case r < 189:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `he`:
		H := 3.1299537542181337
		r := rng.IntN(227)
		switch {
		case r < 56:
			return `r`, H
		case r < 109:
			return `a`, H
		case r < 138:
			return `d`, H
		case r < 157:
			return `l`, H
		case r < 173:
			return `s`, H
		case r < 186:
			return `m`, H
		case r < 197:
			return `e`, H
		case r < 207:
			return `n`, H
		case r < 212:
			return `w`, H
		case r < 216:
			return `f`, H
		case r < 219:
			return `o`, H
		case r < 221:
			return `t`, H
		case r < 223:
			return `c`, H
		case r < 225:
			return `v`, H
		case r < 227:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `nu`:
		H := 3.4253282651341452
		r := rng.IntN(57)
		switch {
		case r < 14:
			return `m`, H
		case r < 25:
			return `t`, H
		case r < 30:
			return `s`, H
		case r < 34:
			return `a`, H
		case r < 38:
			return `g`, H
		case r < 41:
			return `e`, H
		case r < 44:
			return `r`, H
		case r < 47:
			return `c`, H
		case r < 49:
			return `o`, H
		case r < 51:
			return `l`, H
		case r < 52:
			return `n`, H
		case r < 53:
			return `f`, H
		case r < 54:
			return `i`, H
		case r < 55:
			return `p`, H
		case r < 56:
			return `z`, H
		case r < 57:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `lu`:
		H := 3.5210705994541325
		r := rng.IntN(127)
		switch {
		case r < 27:
			return `s`, H
		case r < 41:
			return `t`, H
		case r < 54:
			return `m`, H
		case r < 66:
			return `n`, H
		case r < 77:
			return `r`, H
		case r < 87:
			return `c`, H
		case r < 96:
			return `d`, H
		case r < 104:
			return `e`, H
		case r < 110:
			return `g`, H
		case r < 116:
			return `b`, H
		case r < 120:
			return `x`, H
		case r < 123:
			return `a`, H
		case r < 124:
			return `f`, H
		case r < 125:
			return `i`, H
		case r < 126:
			return `k`, H
		case r < 127:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `so`:
		H := 2.98458260307347
		r := rng.IntN(58)
		switch {
		case r < 16:
			return `l`, H
		case r < 31:
			return `n`, H
		case r < 42:
			return `r`, H
		case r < 44:
			return `i`, H
		case r < 46:
			return `t`, H
		case r < 48:
			return `u`, H
		case r < 49:
			return `f`, H
		case r < 50:
			return `w`, H
		case r < 51:
			return `d`, H
		case r < 52:
			return `o`, H
		case r < 53:
			return `a`, H
		case r < 54:
			return `c`, H
		case r < 55:
			return `p`, H
		case r < 56:
			return `m`, H
		case r < 57:
			return `v`, H
		case r < 58:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `at`:
		H := 2.98197939609681
		r := rng.IntN(506)
		switch {
		case r < 182:
			return `e`, H
		case r < 271:
			return `i`, H
		case r < 319:
			return `t`, H
		case r < 365:
			return `o`, H
		case r < 395:
			return `c`, H
		case r < 420:
			return `h`, H
		case r < 441:
			return `u`, H
		case r < 460:
			return `r`, H
		case r < 479:
			return `a`, H
		case r < 485:
			return `l`, H
		case r < 490:
			return `s`, H
		case r < 494:
			return `n`, H
		case r < 498:
			return `f`, H
		case r < 501:
			return `w`, H
		case r < 504:
			return `b`, H
		case r < 506:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `op`:
		H := 3.240850086008522
		r := rng.IntN(106)
		switch {
		case r < 29:
			return `e`, H
		case r < 50:
			return `p`, H
		case r < 63:
			return `i`, H
		case r < 69:
			return `o`, H
		case r < 75:
			return `s`, H
		case r < 81:
			return `l`, H
		case r < 87:
			return `u`, H
		case r < 92:
			return `h`, H
		case r < 96:
			return `y`, H
		case r < 99:
			return `a`, H
		case r < 101:
			return `t`, H
		case r < 102:
			return `w`, H
		case r < 103:
			return `d`, H
		case r < 104:
			return `k`, H
		case r < 105:
			return `c`, H
		case r < 106:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `oo`:
		H := 3.489791824124952
		r := rng.IntN(195)
		switch {
		case r < 40:
			return `t`, H
		case r < 67:
			return `k`, H
		case r < 92:
			return `n`, H
		case r < 111:
			return `m`, H
		case r < 124:
			return `d`, H
		case r < 137:
			return `r`, H
		case r < 149:
			return `f`, H
		case r < 161:
			return `l`, H
		case r < 171:
			return `p`, H
		case r < 177:
			return `s`, H
		case r < 183:
			return `z`, H
		case r < 188:
			return `g`, H
		case r < 191:
			return `v`, H
		case r < 193:
			return `i`, H
		case r < 194:
			return `e`, H
		case r < 195:
			return `c`, H
		default:
			panic("unexpected rand num")
		}
	case `go`:
		H := 3.4079256616775075
		r := rng.IntN(67)
		switch {
		case r < 16:
			return `n`, H
		case r < 25:
			return `o`, H
		case r < 33:
			return `r`, H
		case r < 40:
			return `t`, H
		case r < 47:
			return `l`, H
		case r < 51:
			return `u`, H
		case r < 54:
			return `s`, H
		case r < 57:
			return `i`, H
		case r < 60:
			return `a`, H
		case r < 61:
			return `w`, H
		case r < 62:
			return `d`, H
		case r < 63:
			return `e`, H
		case r < 64:
			return `p`, H
		case r < 65:
			return `m`, H
		case r < 66:
			return `g`, H
		case r < 67:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `es`:
		H := 2.2514347845496214
		r := rng.IntN(405)
		switch {
		case r < 232:
			return `s`, H
		case r < 299:
			return `t`, H
		case r < 323:
			return `i`, H
		case r < 341:
			return `e`, H
		case r < 352:
			return `h`, H
		case r < 362:
			return `o`, H
		case r < 372:
			return `c`, H
		case r < 380:
			return `p`, H
		case r < 388:
			return `u`, H
		case r < 394:
			return `a`, H
		case r < 399:
			return `k`, H
		case r < 401:
			return `m`, H
		case r < 402:
			return `d`, H
		case r < 403:
			return `y`, H
		case r < 404:
			return `q`, H
		case r < 405:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `ci`:
		H := 3.465666807234371
		r := rng.IntN(132)
		switch {
		case r < 31:
			return `n`, H
		case r < 52:
			return `t`, H
		case r < 64:
			return `d`, H
		case r < 75:
			return `s`, H
		case r < 85:
			return `a`, H
		case r < 93:
			return `e`, H
		case r < 101:
			return `o`, H
		case r < 108:
			return `l`, H
		case r < 114:
			return `r`, H
		case r < 119:
			return `f`, H
		case r < 123:
			return `v`, H
		case r < 126:
			return `p`, H
		case r < 129:
			return `m`, H
		case r < 130:
			return `z`, H
		case r < 131:
			return `u`, H
		case r < 132:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `sh`:
		H := 3.039460068107811
		r := rng.IntN(233)
		switch {
		case r < 48:
			return `o`, H
		case r < 91:
			return `i`, H
		case r < 133:
			return `e`, H
		case r < 174:
			return `a`, H
		case r < 192:
			return `r`, H
		case r < 202:
			return `y`, H
		case r < 212:
			return `u`, H
		case r < 218:
			return `l`, H
		case r < 223:
			return `b`, H
		case r < 227:
			return `c`, H
		case r < 228:
			return `n`, H
		case r < 229:
			return `d`, H
		case r < 230:
			return `h`, H
		case r < 231:
			return `s`, H
		case r < 232:
			return `t`, H
		case r < 233:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ea`:
		H := 3.286394666608718
		r := rng.IntN(296)
		switch {
		case r < 59:
			return `r`, H
		case r < 110:
			return `d`, H
		case r < 159:
			return `t`, H
		case r < 195:
			return `s`, H
		case r < 218:
			return `l`, H
		case r < 237:
			return `m`, H
		case r < 255:
			return `c`, H
		case r < 264:
			return `k`, H
		case r < 271:
			return `n`, H
		case r < 278:
			return `v`, H
		case r < 283:
			return `p`, H
		case r < 288:
			return `b`, H
		case r < 291:
			return `f`, H
		case r < 293:
			return `w`, H
		case r < 295:
			return `g`, H
		case r < 296:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `wa`:
		H := 3.402126204758526
		r := rng.IntN(120)
		switch {
		case r < 25:
			return `r`, H
		case r < 48:
			return `s`, H
		case r < 67:
			return `y`, H
		case r < 77:
			return `l`, H
		case r < 83:
			return `n`, H
		case r < 89:
			return `t`, H
		case r < 95:
			return `g`, H
		case r < 100:
			return `k`, H
		case r < 105:
			return `b`, H
		case r < 109:
			return `v`, H
		case r < 112:
			return `d`, H
		case r < 114:
			return `f`, H
		case r < 116:
			return `w`, H
		case r < 117:
			return `i`, H
		case r < 118:
			return `c`, H
		case r < 119:
			return `p`, H
		case r < 120:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ay`:
		H := 3.9021061909687025
		r := rng.IntN(53)
		switch {
		case r < 6:
			return `e`, H
		case r < 12:
			return `b`, H
		case r < 17:
			return `s`, H
		case r < 22:
			return `i`, H
		case r < 26:
			return `m`, H
		case r < 29:
			return `f`, H
		case r < 32:
			return `d`, H
		case r < 35:
			return `r`, H
		case r < 38:
			return `t`, H
		case r < 41:
			return `a`, H
		case r < 44:
			return `l`, H
		case r < 46:
			return `o`, H
		case r < 48:
			return `c`, H
		case r < 50:
			return `p`, H
		case r < 51:
			return `w`, H
		case r < 52:
			return `h`, H
		case r < 53:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `up`:
		H := 3.696077868346799
		r := rng.IntN(86)
		switch {
		case r < 12:
			return `e`, H
		case r < 23:
			return `p`, H
		case r < 33:
			return `t`, H
		case r < 42:
			return `s`, H
		case r < 50:
			return `r`, H
		case r < 58:
			return `l`, H
		case r < 63:
			return `h`, H
		case r < 68:
			return `i`, H
		case r < 72:
			return `c`, H
		case r < 75:
			return `y`, H
		case r < 77:
			return `w`, H
		case r < 79:
			return `o`, H
		case r < 81:
			return `a`, H
		case r < 83:
			return `b`, H
		case r < 84:
			return `f`, H
		case r < 85:
			return `d`, H
		case r < 86:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `no`:
		H := 3.578556177289845
		r := rng.IntN(101)
		switch {
		case r < 19:
			return `w`, H
		case r < 34:
			return `t`, H
		case r < 47:
			return `r`, H
		case r < 55:
			return `u`, H
		case r < 62:
			return `s`, H
		case r < 69:
			return `p`, H
		case r < 76:
			return `m`, H
		case r < 83:
			return `l`, H
		case r < 86:
			return `o`, H
		case r < 89:
			return `g`, H
		case r < 92:
			return `x`, H
		case r < 95:
			return `v`, H
		case r < 97:
			return `n`, H
		case r < 98:
			return `e`, H
		case r < 99:
			return `y`, H
		case r < 100:
			return `i`, H
		case r < 101:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `tu`:
		H := 3.249185448350641
		r := rng.IntN(135)
		switch {
		case r < 50:
			return `r`, H
		case r < 63:
			return `a`, H
		case r < 74:
			return `d`, H
		case r < 83:
			return `m`, H
		case r < 91:
			return `n`, H
		case r < 99:
			return `b`, H
		case r < 105:
			return `s`, H
		case r < 111:
			return `t`, H
		case r < 116:
			return `f`, H
		case r < 121:
			return `c`, H
		case r < 126:
			return `p`, H
		case r < 128:
			return `e`, H
		case r < 130:
			return `i`, H
		case r < 132:
			return `l`, H
		case r < 133:
			return `o`, H
		case r < 134:
			return `g`, H
		case r < 135:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `da`:
		H := 3.488123022570739
		r := rng.IntN(107)
		switch {
		case r < 19:
			return `n`, H
		case r < 37:
			return `y`, H
		case r < 53:
			return `r`, H
		case r < 67:
			return `t`, H
		case r < 74:
			return `b`, H
		case r < 80:
			return `i`, H
		case r < 85:
			return `l`, H
		case r < 89:
			return `c`, H
		case r < 92:
			return `d`, H
		case r < 95:
			return `s`, H
		case r < 98:
			return `u`, H
		case r < 100:
			return `w`, H
		case r < 102:
			return `z`, H
		case r < 104:
			return `g`, H
		case r < 105:
			return `f`, H
		case r < 106:
			return `m`, H
		case r < 107:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `pa`:
		H := 3.51517726223104
		r := rng.IntN(214)
		switch {
		case r < 48:
			return `r`, H
		case r < 74:
			return `n`, H
		case r < 99:
			return `s`, H
		case r < 120:
			return `c`, H
		case r < 139:
			return `t`, H
		case r < 155:
			return `y`, H
		case r < 165:
			return `l`, H
		case r < 174:
			return `d`, H
		case r < 183:
			return `i`, H
		case r < 190:
			return `p`, H
		case r < 197:
			return `v`, H
		case r < 203:
			return `b`, H
		case r < 207:
			return `m`, H
		case r < 211:
			return `g`, H
		case r < 212:
			return `w`, H
		case r < 213:
			return `j`, H
		case r < 214:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ng`:
		H := 3.17500625470462
		r := rng.IntN(145)
		switch {
		case r < 39:
			return `e`, H
		case r < 72:
			return `l`, H
		case r < 87:
			return `r`, H
		case r < 98:
			return `i`, H
		case r < 108:
			return `o`, H
		case r < 118:
			return `u`, H
		case r < 124:
			return `s`, H
		case r < 130:
			return `a`, H
		case r < 134:
			return `y`, H
		case r < 136:
			return `n`, H
		case r < 138:
			return `t`, H
		case r < 140:
			return `b`, H
		case r < 141:
			return `f`, H
		case r < 142:
			return `w`, H
		case r < 143:
			return `d`, H
		case r < 144:
			return `h`, H
		case r < 145:
			return `m`, H
		default:
			panic("unexpected rand num")
		}
	case `mo`:
		H := 3.623419066149199
		r := rng.IntN(203)
		switch {
		case r < 41:
			return `n`, H
		case r < 72:
			return `r`, H
		case r < 92:
			return `t`, H
		case r < 108:
			return `o`, H
		case r < 124:
			return `u`, H
		case r < 138:
			return `v`, H
		case r < 150:
			return `s`, H
		case r < 161:
			return `l`, H
		case r < 171:
			return `d`, H
		case r < 178:
			return `c`, H
		case r < 183:
			return `i`, H
		case r < 188:
			return `k`, H
		case r < 193:
			return `b`, H
		case r < 196:
			return `m`, H
		case r < 198:
			return `w`, H
		case r < 200:
			return `a`, H
		case r < 202:
			return `g`, H
		case r < 203:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ho`:
		H := 3.5246233511381844
		r := rng.IntN(127)
		switch {
		case r < 20:
			return `r`, H
		case r < 38:
			return `w`, H
		case r < 55:
			return `o`, H
		case r < 71:
			return `n`, H
		case r < 85:
			return `u`, H
		case r < 97:
			return `l`, H
		case r < 103:
			return `p`, H
		case r < 108:
			return `t`, H
		case r < 112:
			return `s`, H
		case r < 115:
			return `e`, H
		case r < 118:
			return `v`, H
		case r < 120:
			return `k`, H
		case r < 122:
			return `m`, H
		case r < 123:
			return `y`, H
		case r < 124:
			return `i`, H
		case r < 125:
			return `c`, H
		case r < 126:
			return `g`, H
		case r < 127:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `el`:
		H := 3.4963439564868266
		r := rng.IntN(231)
		switch {
		case r < 44:
			return `e`, H
		case r < 88:
			return `i`, H
		case r < 116:
			return `l`, H
		case r < 138:
			return `y`, H
		case r < 153:
			return `d`, H
		case r < 168:
			return `a`, H
		case r < 182:
			return `o`, H
		case r < 191:
			return `t`, H
		case r < 199:
			return `f`, H
		case r < 207:
			return `p`, H
		case r < 214:
			return `u`, H
		case r < 220:
			return `v`, H
		case r < 223:
			return `s`, H
		case r < 225:
			return `n`, H
		case r < 227:
			return `c`, H
		case r < 229:
			return `m`, H
		case r < 230:
			return `k`, H
		case r < 231:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ow`:
		H := 3.593817601124209
		r := rng.IntN(102)
		switch {
		case r < 22:
			return `n`, H
		case r < 38:
			return `e`, H
		case r < 47:
			return `d`, H
		case r < 56:
			return `i`, H
		case r < 65:
			return `l`, H
		case r < 72:
			return `s`, H
		case r < 79:
			return `b`, H
		case r < 82:
			return `f`, H
		case r < 85:
			return `a`, H
		case r < 88:
			return `p`, H
		case r < 91:
			return `m`, H
		case r < 93:
			return `w`, H
		case r < 95:
			return `r`, H
		case r < 97:
			return `y`, H
		case r < 99:
			return `c`, H
		case r < 100:
			return `o`, H
		case r < 101:
			return `t`, H
		case r < 102:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ul`:
		H := 3.3475651004206015
		r := rng.IntN(140)
		switch {
		case r < 33:
			return `a`, H
		case r < 63:
			return `l`, H
		case r < 81:
			return `t`, H
		case r < 93:
			return `e`, H
		case r < 101:
			return `p`, H
		case r < 108:
			return `s`, H
		case r < 114:
			return `f`, H
		case r < 120:
			return `i`, H
		case r < 123:
			return `y`, H
		case r < 126:
			return `k`, H
		case r < 129:
			return `g`, H
		case r < 132:
			return `b`, H
		case r < 134:
			return `m`, H
		case r < 136:
			return `u`, H
		case r < 137:
			return `d`, H
		case r < 138:
			return `o`, H
		case r < 139:
			return `c`, H
		case r < 140:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ot`:
		H := 3.4151997626610338
		r := rng.IntN(138)
		switch {
		case r < 27:
			return `i`, H
		case r < 50:
			return `e`, H
		case r < 70:
			return `h`, H
		case r < 90:
			return `t`, H
		case r < 100:
			return `a`, H
		case r < 109:
			return `o`, H
		case r < 114:
			return `l`, H
		case r < 118:
			return `s`, H
		case r < 121:
			return `r`, H
		case r < 124:
			return `c`, H
		case r < 127:
			return `p`, H
		case r < 130:
			return `b`, H
		case r < 132:
			return `w`, H
		case r < 134:
			return `y`, H
		case r < 135:
			return `n`, H
		case r < 136:
			return `m`, H
		case r < 137:
			return `g`, H
		case r < 138:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ca`:
		H := 3.301798879691526
		r := rng.IntN(317)
		switch {
		case r < 71:
			return `r`, H
		case r < 126:
			return `l`, H
		case r < 176:
			return `t`, H
		case r < 210:
			return `n`, H
		case r < 236:
			return `p`, H
		case r < 255:
			return `s`, H
		case r < 267:
			return `m`, H
		case r < 277:
			return `b`, H
		case r < 286:
			return `u`, H
		case r < 294:
			return `d`, H
		case r < 302:
			return `v`, H
		case r < 307:
			return `c`, H
		case r < 311:
			return `k`, H
		case r < 313:
			return `g`, H
		case r < 314:
			return `f`, H
		case r < 315:
			return `w`, H
		case r < 316:
			return `h`, H
		case r < 317:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `po`:
		H := 3.4558028582947937
		r := rng.IntN(189)
		switch {
		case r < 47:
			return `s`, H
		case r < 77:
			return `r`, H
		case r < 94:
			return `i`, H
		case r < 110:
			return `l`, H
		case r < 123:
			return `u`, H
		case r < 135:
			return `n`, H
		case r < 145:
			return `t`, H
		case r < 155:
			return `p`, H
		case r < 164:
			return `w`, H
		case r < 172:
			return `o`, H
		case r < 178:
			return `k`, H
		case r < 181:
			return `d`, H
		case r < 183:
			return `e`, H
		case r < 185:
			return `c`, H
		case r < 186:
			return `f`, H
		case r < 187:
			return `a`, H
		case r < 188:
			return `g`, H
		case r < 189:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `du`:
		H := 3.4682415026057893
		r := rng.IntN(71)
		switch {
		case r < 20:
			return `c`, H
		case r < 30:
			return `r`, H
		case r < 38:
			return `l`, H
		case r < 43:
			return `s`, H
		case r < 48:
			return `p`, H
		case r < 52:
			return `e`, H
		case r < 55:
			return `a`, H
		case r < 58:
			return `m`, H
		case r < 60:
			return `i`, H
		case r < 62:
			return `t`, H
		case r < 64:
			return `b`, H
		case r < 65:
			return `f`, H
		case r < 66:
			return `d`, H
		case r < 67:
			return `o`, H
		case r < 68:
			return `h`, H
		case r < 69:
			return `k`, H
		case r < 70:
			return `g`, H
		case r < 71:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ck`:
		H := 3.3939366348296245
		r := rng.IntN(135)
		switch {
		case r < 30:
			return `e`, H
		case r < 51:
			return `i`, H
		case r < 72:
			return `l`, H
		case r < 91:
			return `s`, H
		case r < 100:
			return `y`, H
		case r < 104:
			return `n`, H
		case r < 108:
			return `w`, H
		case r < 112:
			return `t`, H
		case r < 116:
			return `a`, H
		case r < 120:
			return `b`, H
		case r < 123:
			return `f`, H
		case r < 126:
			return `p`, H
		case r < 129:
			return `u`, H
		case r < 131:
			return `r`, H
		case r < 132:
			return `d`, H
		case r < 133:
			return `o`, H
		case r < 134:
			return `h`, H
		case r < 135:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `il`:
		H := 3.0714574670306756
		r := rng.IntN(249)
		switch {
		case r < 57:
			return `l`, H
		case r < 106:
			return `y`, H
		case r < 147:
			return `e`, H
		case r < 186:
			return `i`, H
		case r < 201:
			return `d`, H
		case r < 213:
			return `t`, H
		case r < 225:
			return `a`, H
		case r < 235:
			return `o`, H
		case r < 237:
			return `s`, H
		case r < 239:
			return `u`, H
		case r < 241:
			return `b`, H
		case r < 242:
			return `n`, H
		case r < 243:
			return `w`, H
		case r < 244:
			return `h`, H
		case r < 245:
			return `r`, H
		case r < 246:
			return `k`, H
		case r < 247:
			return `c`, H
		case r < 248:
			return `m`, H
		case r < 249:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ni`:
		H := 3.452831684393273
		r := rng.IntN(202)
		switch {
		case r < 52:
			return `n`, H
		case r < 80:
			return `s`, H
		case r < 104:
			return `t`, H
		case r < 125:
			return `c`, H
		case r < 137:
			return `f`, H
		case r < 148:
			return `m`, H
		case r < 158:
			return `a`, H
		case r < 167:
			return `o`, H
		case r < 176:
			return `z`, H
		case r < 181:
			return `p`, H
		case r < 186:
			return `v`, H
		case r < 191:
			return `u`, H
		case r < 195:
			return `l`, H
		case r < 197:
			return `g`, H
		case r < 198:
			return `d`, H
		case r < 199:
			return `e`, H
		case r < 200:
			return `q`, H
		case r < 201:
			return `x`, H
		case r < 202:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `or`:
		H := 3.7615134537945343
		r := rng.IntN(290)
		switch {
		case r < 50:
			return `t`, H
		case r < 86:
			return `e`, H
		case r < 113:
			return `i`, H
		case r < 139:
			return `n`, H
		case r < 164:
			return `a`, H
		case r < 187:
			return `m`, H
		case r < 204:
			return `k`, H
		case r < 219:
			return `r`, H
		case r < 232:
			return `d`, H
		case r < 244:
			return `y`, H
		case r < 256:
			return `s`, H
		case r < 264:
			return `o`, H
		case r < 271:
			return `c`, H
		case r < 278:
			return `p`, H
		case r < 282:
			return `b`, H
		case r < 285:
			return `g`, H
		case r < 287:
			return `f`, H
		case r < 289:
			return `l`, H
		case r < 290:
			return `w`, H
		default:
			panic("unexpected rand num")
		}
	case `ha`:
		H := 3.468397105262385
		r := rng.IntN(257)
		switch {
		case r < 66:
			return `n`, H
		case r < 119:
			return `r`, H
		case r < 138:
			return `t`, H
		case r < 155:
			return `p`, H
		case r < 171:
			return `s`, H
		case r < 186:
			return `m`, H
		case r < 197:
			return `l`, H
		case r < 207:
			return `b`, H
		case r < 215:
			return `c`, H
		case r < 222:
			return `d`, H
		case r < 229:
			return `k`, H
		case r < 236:
			return `z`, H
		case r < 242:
			return `i`, H
		case r < 245:
			return `f`, H
		case r < 248:
			return `w`, H
		case r < 251:
			return `v`, H
		case r < 254:
			return `u`, H
		case r < 256:
			return `g`, H
		case r < 257:
			return `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ri`:
		H := 3.8044324336764634
		r := rng.IntN(484)
		switch {
		case r < 94:
			return `n`, H
		case r < 146:
			return `c`, H
		case r < 193:
			return `t`, H
		case r < 236:
			return `s`, H
		case r < 269:
			return `e`, H
		case r < 299:
			return `v`, H
		case r < 325:
			return `m`, H
		case r < 349:
			return `d`, H
		case r < 369:
			return `l`, H
		case r < 387:
			return `a`, H
		case r < 405:
			return `g`, H
		case r < 422:
			return `p`, H
		case r < 438:
			return `f`, H
		case r < 454:
			return `b`, H
		case r < 467:
			return `o`, H
		case r < 475:
			return `z`, H
		case r < 480:
			return `u`, H
		case r < 483:
			return `k`, H
		case r < 484:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `o`:
		H := 2.9453946587990396
		r := rng.IntN(246)
		switch {
		case r < 87:
			return `v`, H
		case r < 140:
			return `u`, H
		case r < 167:
			return `b`, H
		case r < 186:
			return `p`, H
		case r < 201:
			return `n`, H
		case r < 212:
			return `c`, H
		case r < 219:
			return `x`, H
		case r < 225:
			return `m`, H
		case r < 229:
			return `a`, H
		case r < 232:
			return `o`, H
		case r < 235:
			return `i`, H
		case r < 238:
			return `l`, H
		case r < 240:
			return `t`, H
		case r < 241:
			return `w`, H
		case r < 242:
			return `s`, H
		case r < 243:
			return `k`, H
		case r < 244:
			return `y`, H
		case r < 245:
			return `z`, H
		case r < 246:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ti`:
		H := 3.485360725030509
		r := rng.IntN(500)
		switch {
		case r < 118:
			return `n`, H
		case r < 208:
			return `o`, H
		case r < 273:
			return `c`, H
		case r < 313:
			return `v`, H
		case r < 337:
			return `t`, H
		case r < 361:
			return `l`, H
		case r < 384:
			return `f`, H
		case r < 405:
			return `m`, H
		case r < 425:
			return `s`, H
		case r < 437:
			return `g`, H
		case r < 448:
			return `e`, H
		case r < 459:
			return `r`, H
		case r < 470:
			return `p`, H
		case r < 480:
			return `z`, H
		case r < 488:
			return `d`, H
		case r < 494:
			return `a`, H
		case r < 497:
			return `q`, H
		case r < 499:
			return `b`, H
		case r < 500:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `si`:
		H := 3.7124503830288225
		r := rng.IntN(246)
		switch {
		case r < 43:
			return `n`, H
		case r < 75:
			return `o`, H
		case r < 100:
			return `t`, H
		case r < 122:
			return `v`, H
		case r < 143:
			return `d`, H
		case r < 164:
			return `l`, H
		case r < 180:
			return `s`, H
		case r < 191:
			return `z`, H
		case r < 202:
			return `g`, H
		case r < 212:
			return `b`, H
		case r < 221:
			return `m`, H
		case r < 227:
			return `x`, H
		case r < 232:
			return `e`, H
		case r < 237:
			return `c`, H
		case r < 239:
			return `f`, H
		case r < 241:
			return `r`, H
		case r < 243:
			return `a`, H
		case r < 245:
			return `p`, H
		case r < 246:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `la`:
		H := 3.734800815977329
		r := rng.IntN(393)
		switch {
		case r < 68:
			return `t`, H
		case r < 132:
			return `n`, H
		case r < 174:
			return `r`, H
		case r < 212:
			return `s`, H
		case r < 236:
			return `y`, H
		case r < 257:
			return `c`, H
		case r < 277:
			return `m`, H
		case r < 295:
			return `p`, H
		case r < 312:
			return `d`, H
		case r < 326:
			return `i`, H
		case r < 338:
			return `u`, H
		case r < 350:
			return `b`, H
		case r < 361:
			return `g`, H
		case r < 371:
			return `z`, H
		case r < 378:
			return `v`, H
		case r < 384:
			return `w`, H
		case r < 389:
			return `k`, H
		case r < 392:
			return `x`, H
		case r < 393:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `ga`:
		H := 3.4964001580415376
		r := rng.IntN(98)
		switch {
		case r < 18:
			return `t`, H
		case r < 35:
			return `l`, H
		case r < 50:
			return `r`, H
		case r < 62:
			return `n`, H
		case r < 69:
			return `m`, H
		case r < 75:
			return `g`, H
		case r < 79:
			return `i`, H
		case r < 82:
			return `z`, H
		case r < 85:
			return `u`, H
		case r < 88:
			return `b`, H
		case r < 90:
			return `d`, H
		case r < 91:
			return `f`, H
		case r < 92:
			return `w`, H
		case r < 93:
			return `e`, H
		case r < 94:
			return `h`, H
		case r < 95:
			return `s`, H
		case r < 96:
			return `c`, H
		case r < 97:
			return `p`, H
		case r < 98:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `sa`:
		H := 3.7462923073871974
		r := rng.IntN(159)
		switch {
		case r < 30:
			return `l`, H
		case r < 56:
			return `n`, H
		case r < 73:
			return `b`, H
		case r < 84:
			return `t`, H
		case r < 94:
			return `g`, H
		case r < 102:
			return `d`, H
		case r < 110:
			return `v`, H
		case r < 117:
			return `f`, H
		case r < 123:
			return `i`, H
		case r < 129:
			return `r`, H
		case r < 135:
			return `m`, H
		case r < 140:
			return `c`, H
		case r < 145:
			return `u`, H
		case r < 149:
			return `s`, H
		case r < 152:
			return `w`, H
		case r < 155:
			return `p`, H
		case r < 157:
			return `y`, H
		case r < 158:
			return `k`, H
		case r < 159:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `na`:
		H := 3.6919110777928674
		r := rng.IntN(161)
		switch {
		case r < 30:
			return `t`, H
		case r < 53:
			return `l`, H
		case r < 70:
			return `n`, H
		case r < 84:
			return `r`, H
		case r < 96:
			return `m`, H
		case r < 107:
			return `p`, H
		case r < 117:
			return `g`, H
		case r < 126:
			return `b`, H
		case r < 134:
			return `c`, H
		case r < 140:
			return `s`, H
		case r < 144:
			return `d`, H
		case r < 148:
			return `i`, H
		case r < 151:
			return `v`, H
		case r < 154:
			return `u`, H
		case r < 156:
			return `w`, H
		case r < 158:
			return `e`, H
		case r < 159:
			return `f`, H
		case r < 160:
			return `k`, H
		case r < 161:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ad`:
		H := 3.4091611715594885
		r := rng.IntN(149)
		switch {
		case r < 40:
			return `e`, H
		case r < 69:
			return `i`, H
		case r < 87:
			return `d`, H
		case r < 100:
			return `l`, H
		case r < 106:
			return `s`, H
		case r < 112:
			return `y`, H
		case r < 117:
			return `o`, H
		case r < 122:
			return `a`, H
		case r < 125:
			return `n`, H
		case r < 128:
			return `f`, H
		case r < 131:
			return `w`, H
		case r < 134:
			return `r`, H
		case r < 137:
			return `p`, H
		case r < 140:
			return `b`, H
		case r < 142:
			return `c`, H
		case r < 144:
			return `m`, H
		case r < 146:
			return `g`, H
		case r < 148:
			return `u`, H
		case r < 149:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ub`:
		H := 3.6756598041227804
		r := rng.IntN(89)
		switch {
		case r < 21:
			return `b`, H
		case r < 33:
			return `s`, H
		case r < 43:
			return `l`, H
		case r < 51:
			return `t`, H
		case r < 56:
			return `i`, H
		case r < 61:
			return `m`, H
		case r < 65:
			return `d`, H
		case r < 69:
			return `p`, H
		case r < 72:
			return `e`, H
		case r < 75:
			return `a`, H
		case r < 77:
			return `w`, H
		case r < 79:
			return `h`, H
		case r < 81:
			return `r`, H
		case r < 83:
			return `u`, H
		case r < 84:
			return `f`, H
		case r < 85:
			return `y`, H
		case r < 86:
			return `j`, H
		case r < 87:
			return `c`, H
		case r < 88:
			return `z`, H
		case r < 89:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `bo`:
		H := 3.6735456800695463
		r := rng.IntN(153)
		switch {
		case r < 26:
			return `o`, H
		case r < 46:
			return `n`, H
		case r < 64:
			return `a`, H
		case r < 80:
			return `x`, H
		case r < 94:
			return `r`, H
		case r < 106:
			return `u`, H
		case r < 116:
			return `t`, H
		case r < 123:
			return `d`, H
		case r < 130:
			return `b`, H
		case r < 136:
			return `l`, H
		case r < 139:
			return `w`, H
		case r < 142:
			return `s`, H
		case r < 145:
			return `g`, H
		case r < 147:
			return `y`, H
		case r < 148:
			return `f`, H
		case r < 149:
			return `e`, H
		case r < 150:
			return `j`, H
		case r < 151:
			return `i`, H
		case r < 152:
			return `k`, H
		case r < 153:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ta`:
		H := 3.632864868080895
		r := rng.IntN(324)
		switch {
		case r < 45:
			return `r`, H
		case r < 87:
			return `l`, H
		case r < 126:
			return `b`, H
		case r < 162:
			return `t`, H
		case r < 195:
			return `n`, H
		case r < 224:
			return `i`, H
		case r < 246:
			return `g`, H
		case r < 267:
			return `c`, H
		case r < 279:
			return `s`, H
		case r < 290:
			return `p`, H
		case r < 300:
			return `k`, H
		case r < 310:
			return `m`, H
		case r < 314:
			return `d`, H
		case r < 316:
			return `f`, H
		case r < 318:
			return `y`, H
		case r < 320:
			return `u`, H
		case r < 321:
			return `w`, H
		case r < 322:
			return `e`, H
		case r < 323:
			return `v`, H
		case r < 324:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `di`:
		H := 3.4494359973163755
		r := rng.IntN(288)
		switch {
		case r < 71:
			return `s`, H
		case r < 134:
			return `n`, H
		case r < 158:
			return `a`, H
		case r < 178:
			return `t`, H
		case r < 195:
			return `c`, H
		case r < 211:
			return `v`, H
		case r < 225:
			return `l`, H
		case r < 234:
			return `e`, H
		case r < 243:
			return `m`, H
		case r < 250:
			return `f`, H
		case r < 257:
			return `o`, H
		case r < 264:
			return `r`, H
		case r < 270:
			return `z`, H
		case r < 276:
			return `g`, H
		case r < 280:
			return `p`, H
		case r < 284:
			return `b`, H
		case r < 285:
			return `w`, H
		case r < 286:
			return `d`, H
		case r < 287:
			return `x`, H
		case r < 288:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `se`:
		H := 3.6228884091123996
		r := rng.IntN(199)
		switch {
		case r < 32:
			return `d`, H
		case r < 60:
			return `r`, H
		case r < 86:
			return `l`, H
		case r < 111:
			return `n`, H
		case r < 127:
			return `s`, H
		case r < 142:
			return `c`, H
		case r < 156:
			return `t`, H
		case r < 165:
			return `m`, H
		case r < 171:
			return `a`, H
		case r < 177:
			return `v`, H
		case r < 182:
			return `e`, H
		case r < 186:
			return `p`, H
		case r < 189:
			return `q`, H
		case r < 191:
			return `y`, H
		case r < 193:
			return `i`, H
		case r < 195:
			return `u`, H
		case r < 196:
			return `f`, H
		case r < 197:
			return `w`, H
		case r < 198:
			return `g`, H
		case r < 199:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ur`:
		H := 3.773670379758159
		r := rng.IntN(244)
		switch {
		case r < 63:
			return `e`, H
		case r < 90:
			return `i`, H
		case r < 110:
			return `a`, H
		case r < 125:
			return `s`, H
		case r < 139:
			return `r`, H
		case r < 152:
			return `g`, H
		case r < 162:
			return `n`, H
		case r < 172:
			return `v`, H
		case r < 181:
			return `l`, H
		case r < 190:
			return `b`, H
		case r < 198:
			return `o`, H
		case r < 206:
			return `t`, H
		case r < 214:
			return `p`, H
		case r < 221:
			return `f`, H
		case r < 227:
			return `d`, H
		case r < 233:
			return `c`, H
		case r < 237:
			return `k`, H
		case r < 240:
			return `y`, H
		case r < 242:
			return `m`, H
		case r < 244:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `do`:
		H := 3.6985346528002148
		r := rng.IntN(103)
		switch {
		case r < 15:
			return `r`, H
		case r < 29:
			return `w`, H
		case r < 43:
			return `o`, H
		case r < 56:
			return `m`, H
		case r < 65:
			return `n`, H
		case r < 73:
			return `l`, H
		case r < 80:
			return `c`, H
		case r < 83:
			return `s`, H
		case r < 86:
			return `i`, H
		case r < 89:
			return `g`, H
		case r < 91:
			return `d`, H
		case r < 93:
			return `t`, H
		case r < 95:
			return `z`, H
		case r < 97:
			return `u`, H
		case r < 98:
			return `f`, H
		case r < 99:
			return `e`, H
		case r < 100:
			return `k`, H
		case r < 101:
			return `a`, H
		case r < 102:
			return `v`, H
		case r < 103:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ut`:
		H := 3.8482247103998266
		r := rng.IntN(152)
		switch {
		case r < 27:
			return `e`, H
		case r < 49:
			return `i`, H
		case r < 64:
			return `t`, H
		case r < 76:
			return `s`, H
		case r < 86:
			return `o`, H
		case r < 95:
			return `a`, H
		case r < 103:
			return `h`, H
		case r < 111:
			return `l`, H
		case r < 117:
			return `r`, H
		case r < 123:
			return `b`, H
		case r < 128:
			return `c`, H
		case r < 132:
			return `p`, H
		case r < 136:
			return `m`, H
		case r < 139:
			return `f`, H
		case r < 142:
			return `w`, H
		case r < 145:
			return `d`, H
		case r < 147:
			return `y`, H
		case r < 149:
			return `g`, H
		case r < 151:
			return `u`, H
		case r < 152:
			return `n`, H
		default:
			panic("unexpected rand num")
		}
	case `e`:
		H := 3.440729337470061
		r := rng.IntN(398)
		switch {
		case r < 87:
			return `n`, H
		case r < 169:
			return `x`, H
		case r < 208:
			return `m`, H
		case r < 245:
			return `a`, H
		case r < 280:
			return `l`, H
		case r < 301:
			return `s`, H
		case r < 322:
			return `v`, H
		case r < 334:
			return `c`, H
		case r < 345:
			return `r`, H
		case r < 355:
			return `p`, H
		case r < 364:
			return `d`, H
		case r < 372:
			return `q`, H
		case r < 380:
			return `g`, H
		case r < 387:
			return `t`, H
		case r < 391:
			return `f`, H
		case r < 394:
			return `b`, H
		case r < 395:
			return `e`, H
		case r < 396:
			return `j`, H
		case r < 397:
			return `i`, H
		case r < 398:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ma`:
		H := 3.24717090849481
		r := rng.IntN(267)
		switch {
		case r < 75:
			return `n`, H
		case r < 124:
			return `t`, H
		case r < 166:
			return `r`, H
		case r < 187:
			return `g`, H
		case r < 203:
			return `s`, H
		case r < 215:
			return `l`, H
		case r < 225:
			return `k`, H
		case r < 235:
			return `c`, H
		case r < 239:
			return `d`, H
		case r < 243:
			return `i`, H
		case r < 247:
			return `y`, H
		case r < 251:
			return `j`, H
		case r < 255:
			return `m`, H
		case r < 258:
			return `b`, H
		case r < 260:
			return `p`, H
		case r < 262:
			return `z`, H
		case r < 264:
			return `x`, H
		case r < 265:
			return `h`, H
		case r < 266:
			return `v`, H
		case r < 267:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `is`:
		H := 3.425586749295823
		r := rng.IntN(339)
		switch {
		case r < 84:
			return `t`, H
		case r < 148:
			return `h`, H
		case r < 178:
			return `m`, H
		case r < 207:
			return `e`, H
		case r < 230:
			return `i`, H
		case r < 252:
			return `p`, H
		case r < 269:
			return `o`, H
		case r < 285:
			return `s`, H
		case r < 298:
			return `c`, H
		case r < 307:
			return `a`, H
		case r < 315:
			return `k`, H
		case r < 323:
			return `l`, H
		case r < 327:
			return `b`, H
		case r < 330:
			return `f`, H
		case r < 332:
			return `d`, H
		case r < 334:
			return `y`, H
		case r < 336:
			return `r`, H
		case r < 337:
			return `w`, H
		case r < 338:
			return `j`, H
		case r < 339:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `nd`:
		H := 3.5314039953405065
		r := rng.IntN(257)
		switch {
		case r < 75:
			return `e`, H
		case r < 103:
			return `i`, H
		case r < 130:
			return `l`, H
		case r < 151:
			return `a`, H
		case r < 170:
			return `o`, H
		case r < 183:
			return `s`, H
		case r < 194:
			return `r`, H
		case r < 204:
			return `u`, H
		case r < 213:
			return `b`, H
		case r < 219:
			return `f`, H
		case r < 225:
			return `w`, H
		case r < 231:
			return `p`, H
		case r < 237:
			return `m`, H
		case r < 242:
			return `n`, H
		case r < 247:
			return `y`, H
		case r < 252:
			return `c`, H
		case r < 254:
			return `g`, H
		case r < 255:
			return `d`, H
		case r < 256:
			return `h`, H
		case r < 257:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ee`:
		H := 3.7833598092940597
		r := rng.IntN(126)
		switch {
		case r < 20:
			return `n`, H
		case r < 40:
			return `d`, H
		case r < 53:
			return `p`, H
		case r < 64:
			return `z`, H
		case r < 74:
			return `t`, H
		case r < 83:
			return `l`, H
		case r < 91:
			return `r`, H
		case r < 97:
			return `m`, H
		case r < 102:
			return `b`, H
		case r < 105:
			return `f`, H
		case r < 108:
			return `w`, H
		case r < 111:
			return `s`, H
		case r < 114:
			return `i`, H
		case r < 117:
			return `k`, H
		case r < 120:
			return `c`, H
		case r < 122:
			return `a`, H
		case r < 123:
			return `h`, H
		case r < 124:
			return `x`, H
		case r < 125:
			return `v`, H
		case r < 126:
			return `g`, H
		default:
			panic("unexpected rand num")
		}
	case `lo`:
		H := 3.9481220227019564
		r := rng.IntN(203)
		switch {
		case r < 26:
			return `g`, H
		case r < 48:
			return `r`, H
		case r < 67:
			return `n`, H
		case r < 84:
			return `a`, H
		case r < 100:
			return `t`, H
		case r < 115:
			return `w`, H
		case r < 128:
			return `o`, H
		case r < 141:
			return `c`, H
		case r < 152:
			return `p`, H
		case r < 162:
			return `s`, H
		case r < 172:
			return `u`, H
		case r < 179:
			return `v`, H
		case r < 185:
			return `y`, H
		case r < 190:
			return `b`, H
		case r < 194:
			return `d`, H
		case r < 196:
			return `i`, H
		case r < 198:
			return `q`, H
		case r < 200:
			return `m`, H
		case r < 201:
			return `f`, H
		case r < 202:
			return `e`, H
		case r < 203:
			return `h`, H
		default:
			panic("unexpected rand num")
		}
	case `a`:
		H := 3.8126739941897543
		r := rng.IntN(407)
		switch {
		case r < 74:
			return `n`, H
		case r < 113:
			return `m`, H
		case r < 148:
			return `r`, H
		case r < 181:
			return `l`, H
		case r < 211:
			return `c`, H
		case r < 240:
			return `p`, H
		case r < 264:
			return `b`, H
		case r < 287:
			return `f`, H
		case r < 310:
			return `s`, H
		case r < 332:
			return `t`, H
		case r < 352:
			return `g`, H
		case r < 371:
			return `u`, H
		case r < 384:
			return `v`, H
		case r < 393:
			return `w`, H
		case r < 397:
			return `e`, H
		case r < 400:
			return `i`, H
		case r < 402:
			return `h`, H
		case r < 404:
			return `q`, H
		case r < 405:
			return `o`, H
		case r < 406:
			return `j`, H
		case r < 407:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `de`:
		H := 3.6173342588664883
		r := rng.IntN(390)
		switch {
		case r < 97:
			return `r`, H
		case r < 138:
			return `d`, H
		case r < 177:
			return `n`, H
		case r < 211:
			return `f`, H
		case r < 241:
			return `c`, H
		case r < 267:
			return `l`, H
		case r < 288:
			return `s`, H
		case r < 305:
			return `p`, H
		case r < 321:
			return `a`, H
		case r < 336:
			return `v`, H
		case r < 350:
			return `t`, H
		case r < 361:
			return `m`, H
		case r < 370:
			return `b`, H
		case r < 376:
			return `e`, H
		case r < 380:
			return `o`, H
		case r < 384:
			return `g`, H
		case r < 386:
			return `x`, H
		case r < 387:
			return `h`, H
		case r < 388:
			return `j`, H
		case r < 389:
			return `i`, H
		case r < 390:
			return `u`, H
		default:
			panic("unexpected rand num")
		}
	case `in`:
		H := 2.1017950678134936
		r := rng.IntN(1050)
		switch {
		case r < 665:
			return `g`, H
		case r < 813:
			return `e`, H
		case r < 855:
			return `t`, H
		case r < 888:
			return `i`, H
		case r < 918:
			return `k`, H
		case r < 944:
			return `a`, H
		case r < 969:
			return `d`, H
		case r < 984:
			return `n`, H
		case r < 998:
			return `s`, H
		case r < 1012:
			return `c`, H
		case r < 1021:
			return `o`, H
		case r < 1028:
			return `y`, H
		case r < 1035:
			return `l`, H
		case r < 1039:
			return `f`, H
		case r < 1042:
			return `u`, H
		case r < 1044:
			return `j`, H
		case r < 1046:
			return `x`, H
		case r < 1047:
			return `w`, H
		case r < 1048:
			return `h`, H
		case r < 1049:
			return `p`, H
		case r < 1050:
			return `v`, H
		default:
			panic("unexpected rand num")
		}
	case `li`:
		H := 3.546418837133385
		r := rng.IntN(475)
		switch {
		case r < 157:
			return `n`, H
		case r < 200:
			return `c`, H
		case r < 235:
			return `t`, H
		case r < 264:
			return `k`, H
		case r < 293:
			return `g`, H
		case r < 320:
			return `s`, H
		case r < 341:
			return `a`, H
		case r < 359:
			return `f`, H
		case r < 377:
			return `m`, H
		case r < 393:
			return `v`, H
		case r < 408:
			return `z`, H
		case r < 422:
			return `e`, H
		case r < 433:
			return `d`, H
		case r < 443:
			return `o`, H
		case r < 453:
			return `p`, H
		case r < 462:
			return `b`, H
		case r < 466:
			return `q`, H
		case r < 469:
			return `r`, H
		case r < 472:
			return `l`, H
		case r < 474:
			return `u`, H
		case r < 475:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ne`:
		H := 3.0862784397322383
		r := rng.IntN(326)
		switch {
		case r < 133:
			return `s`, H
		case r < 178:
			return `r`, H
		case r < 217:
			return `d`, H
		case r < 232:
			return `t`, H
		case r < 247:
			return `a`, H
		case r < 257:
			return `n`, H
		case r < 267:
			return `l`, H
		case r < 275:
			return `w`, H
		case r < 283:
			return `g`, H
		case r < 290:
			return `e`, H
		case r < 297:
			return `u`, H
		case r < 303:
			return `y`, H
		case r < 308:
			return `m`, H
		case r < 313:
			return `x`, H
		case r < 316:
			return `c`, H
		case r < 319:
			return `v`, H
		case r < 322:
			return `b`, H
		case r < 323:
			return `o`, H
		case r < 324:
			return `h`, H
		case r < 325:
			return `q`, H
		case r < 326:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `on`:
		H := 3.7481496950737694
		r := rng.IntN(302)
		switch {
		case r < 59:
			return `e`, H
		case r < 87:
			return `i`, H
		case r < 115:
			return `s`, H
		case r < 143:
			return `g`, H
		case r < 166:
			return `d`, H
		case r < 188:
			return `t`, H
		case r < 208:
			return `f`, H
		case r < 227:
			return `a`, H
		case r < 244:
			return `c`, H
		case r < 260:
			return `o`, H
		case r < 275:
			return `y`, H
		case r < 281:
			return `l`, H
		case r < 285:
			return `u`, H
		case r < 288:
			return `n`, H
		case r < 291:
			return `j`, H
		case r < 294:
			return `v`, H
		case r < 296:
			return `w`, H
		case r < 298:
			return `z`, H
		case r < 300:
			return `b`, H
		case r < 301:
			return `k`, H
		case r < 302:
			return `r`, H
		default:
			panic("unexpected rand num")
		}
	case `le`:
		H := 3.601021056216224
		r := rng.IntN(318)
		switch {
		case r < 69:
			return `s`, H
		case r < 115:
			return `d`, H
		case r < 153:
			return `r`, H
		case r < 182:
			return `t`, H
		case r < 204:
			return `n`, H
		case r < 225:
			return `c`, H
		case r < 243:
			return `a`, H
		case r < 260:
			return `g`, H
		case r < 273:
			return `v`, H
		case r < 283:
			return `e`, H
		case r < 290:
			return `m`, H
		case r < 296:
			return `x`, H
		case r < 301:
			return `y`, H
		case r < 305:
			return `p`, H
		case r < 308:
			return `f`, H
		case r < 311:
			return `u`, H
		case r < 313:
			return `b`, H
		case r < 314:
			return `w`, H
		case r < 315:
			return `o`, H
		case r < 316:
			return `h`, H
		case r < 317:
			return `i`, H
		case r < 318:
			return `l`, H
		default:
			panic("unexpected rand num")
		}
	case `al`:
		H := 3.693547477659404
		r := rng.IntN(250)
		switch {
		case r < 55:
			return `l`, H
		case r < 102:
			return `i`, H
		case r < 123:
			return `a`, H
		case r < 140:
			return `e`, H
		case r < 156:
			return `o`, H
		case r < 172:
			return `t`, H
		case r < 183:
			return `u`, H
		case r < 192:
			return `k`, H
		case r < 201:
			return `m`, H
		case r < 208:
			return `s`, H
		case r < 214:
			return `y`, H
		case r < 220:
			return `c`, H
		case r < 226:
			return `v`, H
		case r < 230:
			return `f`, H
		case r < 234:
			return `d`, H
		case r < 238:
			return `r`, H
		case r < 242:
			return `p`, H
		case r < 244:
			return `n`, H
		case r < 246:
			return `g`, H
		case r < 248:
			return `b`, H
		case r < 249:
			return `w`, H
		case r < 250:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ra`:
		H := 3.9497641727639463
		r := rng.IntN(471)
		switch {
		case r < 69:
			return `n`, H
		case r < 125:
			return `t`, H
		case r < 161:
			return `c`, H
		case r < 195:
			return `i`, H
		case r < 228:
			return `d`, H
		case r < 260:
			return `m`, H
		case r < 291:
			return `l`, H
		case r < 321:
			return `g`, H
		case r < 346:
			return `v`, H
		case r < 369:
			return `p`, H
		case r < 391:
			return `s`, H
		case r < 410:
			return `b`, H
		case r < 425:
			return `y`, H
		case r < 439:
			return `f`, H
		case r < 448:
			return `r`, H
		case r < 456:
			return `w`, H
		case r < 462:
			return `z`, H
		case r < 466:
			return `k`, H
		case r < 468:
			return `u`, H
		case r < 469:
			return `e`, H
		case r < 470:
			return `o`, H
		case r < 471:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ar`:
		H := 4.022304287988487
		r := rng.IntN(513)
		switch {
		case r < 79:
			return `d`, H
		case r < 132:
			return `t`, H
		case r < 174:
			return `i`, H
		case r < 215:
			return `m`, H
		case r < 251:
			return `e`, H
		case r < 284:
			return `r`, H
		case r < 316:
			return `a`, H
		case r < 346:
			return `y`, H
		case r < 371:
			return `g`, H
		case r < 392:
			return `l`, H
		case r < 409:
			return `n`, H
		case r < 426:
			return `c`, H
		case r < 442:
			return `b`, H
		case r < 457:
			return `k`, H
		case r < 470:
			return `p`, H
		case r < 482:
			return `o`, H
		case r < 493:
			return `s`, H
		case r < 501:
			return `f`, H
		case r < 508:
			return `v`, H
		case r < 511:
			return `w`, H
		case r < 512:
			return `h`, H
		case r < 513:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case `to`:
		H := 2.943211643112022
		r := rng.IntN(200)
		switch {
		case r < 82:
			return `r`, H
		case r < 116:
			return `n`, H
		case r < 136:
			return `p`, H
		case r < 154:
			return `m`, H
		case r < 164:
			return `o`, H
		case r < 170:
			return `u`, H
		case r < 175:
			return `c`, H
		case r < 179:
			return `w`, H
		case r < 183:
			return `l`, H
		case r < 185:
			return `d`, H
		case r < 187:
			return `i`, H
		case r < 189:
			return `t`, H
		case r < 191:
			return `x`, H
		case r < 192:
			return `f`, H
		case r < 193:
			return `e`, H
		case r < 194:
			return `s`, H
		case r < 195:
			return `k`, H
		case r < 196:
			return `a`, H
		case r < 197:
			return `z`, H
		case r < 198:
			return `g`, H
		case r < 199:
			return `v`, H
		case r < 200:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `en`:
		H := 3.239836470529042
		r := rng.IntN(441)
		switch {
		case r < 148:
			return `t`, H
		case r < 217:
			return `d`, H
		case r < 268:
			return `c`, H
		case r < 306:
			return `e`, H
		case r < 331:
			return `s`, H
		case r < 351:
			return `g`, H
		case r < 368:
			return `a`, H
		case r < 384:
			return `i`, H
		case r < 392:
			return `n`, H
		case r < 400:
			return `v`, H
		case r < 407:
			return `o`, H
		case r < 414:
			return `u`, H
		case r < 419:
			return `j`, H
		case r < 424:
			return `l`, H
		case r < 428:
			return `r`, H
		case r < 431:
			return `f`, H
		case r < 434:
			return `z`, H
		case r < 436:
			return `h`, H
		case r < 438:
			return `y`, H
		case r < 439:
			return `q`, H
		case r < 440:
			return `k`, H
		case r < 441:
			return `p`, H
		default:
			panic("unexpected rand num")
		}
	case `an`:
		H := 3.4086078047375628
		r := rng.IntN(553)
		switch {
		case r < 125:
			return `d`, H
		case r < 238:
			return `t`, H
		case r < 294:
			return `c`, H
		case r < 347:
			return `g`, H
		case r < 386:
			return `k`, H
		case r < 422:
			return `i`, H
		case r < 454:
			return `n`, H
		case r < 473:
			return `a`, H
		case r < 489:
			return `y`, H
		case r < 503:
			return `e`, H
		case r < 516:
			return `o`, H
		case r < 528:
			return `s`, H
		case r < 535:
			return `l`, H
		case r < 541:
			return `h`, H
		case r < 544:
			return `u`, H
		case r < 546:
			return `q`, H
		case r < 547:
			return `f`, H
		case r < 548:
			return `j`, H
		case r < 549:
			return `p`, H
		case r < 550:
			return `m`, H
		case r < 551:
			return `z`, H
		case r < 552:
			return `v`, H
		case r < 553:
			return `b`, H
		default:
			panic("unexpected rand num")
		}
	case `co`:
		H := 3.4735036327365343
		r := rng.IntN(349)
		switch {
		case r < 93:
			return `n`, H
		case r < 145:
			return `m`, H
		case r < 196:
			return `r`, H
		case r < 224:
			return `u`, H
		case r < 252:
			return `l`, H
		case r < 268:
			return `a`, H
		case r < 281:
			return `p`, H
		case r < 291:
			return `s`, H
		case r < 299:
			return `o`, H
		case r < 307:
			return `t`, H
		case r < 315:
			return `v`, H
		case r < 320:
			return `d`, H
		case r < 324:
			return `g`, H
		case r < 328:
			return `b`, H
		case r < 331:
			return `f`, H
		case r < 334:
			return `e`, H
		case r < 337:
			return `h`, H
		case r < 340:
			return `i`, H
		case r < 343:
			return `z`, H
		case r < 345:
			return `y`, H
		case r < 347:
			return `c`, H
		case r < 348:
			return `w`, H
		case r < 349:
			return `k`, H
		default:
			panic("unexpected rand num")
		}
	case `er`:
		H := 4.231365586934958
		r := rng.IntN(466)
		switch {
		case r < 48:
			return `s`, H
		case r < 96:
			return `i`, H
		case r < 134:
			return `e`, H
		case r < 172:
			return `a`, H
		case r < 205:
			return `t`, H
		case r < 233:
			return `y`, H
		case r < 256:
			return `o`, H
		case r < 279:
			return `r`, H
		case r < 301:
			return `b`, H
		case r < 319:
			return `n`, H
		case r < 337:
			return `m`, H
		case r < 355:
			return `g`, H
		case r < 372:
			return `c`, H
		case r < 389:
			return `v`, H
		case r < 405:
			return `l`, H
		case r < 418:
			return `p`, H
		case r < 430:
			return `f`, H
		case r < 440:
			return `d`, H
		case r < 449:
			return `h`, H
		case r < 454:
			return `u`, H
		case r < 458:
			return `w`, H
		case r < 462:
			return `k`, H
		case r < 466:
			return `j`, H
		default:
			panic("unexpected rand num")
		}
	case `ro`:
		H := 4.143180722133236
		r := rng.IntN(388)
		switch {
		case r < 45:
			return `u`, H
		case r < 84:
			return `n`, H
		case r < 115:
			return `w`, H
		case r < 146:
			return `o`, H
		case r < 172:
			return `p`, H
		case r < 196:
			return `v`, H
		case r < 218:
			return `c`, H
		case r < 239:
			return `s`, H
		case r < 259:
			return `a`, H
		case r < 278:
			return `l`, H
		case r < 296:
			return `b`, H
		case r < 313:
			return `t`, H
		case r < 329:
			return `m`, H
		case r < 343:
			return `g`, H
		case r < 353:
			return `d`, H
		case r < 361:
			return `f`, H
		case r < 369:
			return `i`, H
		case r < 374:
			return `r`, H
		case r < 379:
			return `x`, H
		case r < 383:
			return `k`, H
		case r < 386:
			return `z`, H
		case r < 387:
			return `j`, H
		case r < 388:
			return `y`, H
		default:
			panic("unexpected rand num")
		}
	case `un`:
		H := 4.013869088306411
		r := rng.IntN(583)
		switch {
		case r < 92:
			return `d`, H
		case r < 160:
			return `c`, H
		case r < 227:
			return `t`, H
		case r < 281:
			return `s`, H
		case r < 312:
			return `i`, H
		case r < 338:
			return `l`, H
		case r < 362:
			return `w`, H
		case r < 385:
			return `g`, H
		case r < 407:
			return `e`, H
		case r < 429:
			return `r`, H
		case r < 449:
			return `f`, H
		case r < 469:
			return `b`, H
		case r < 487:
			return `k`, H
		case r < 505:
			return `a`, H
		case r < 523:
			return `p`, H
		case r < 539:
			return `m`, H
		case r < 553:
			return `n`, H
		case r < 566:
			return `h`, H
		case r < 575:
			return `v`, H
		case r < 578:
			return `u`, H
		case r < 580:
			return `o`, H
		case r < 581:
			return `j`, H
		case r < 582:
			return `q`, H
		case r < 583:
			return `z`, H
		default:
			panic("unexpected rand num")
		}
	case `re`:
		H := 4.070035666531293
		r := rng.IntN(744)
		switch {
		case r < 88:
			return `a`, H
		case r < 173:
			return `s`, H
		case r < 232:
			return `c`, H
		case r < 290:
			return `e`, H
		case r < 346:
			return `d`, H
		case r < 397:
			return `t`, H
		case r < 441:
			return `p`, H
		case r < 484:
			return `n`, H
		case r < 527:
			return `l`, H
		case r < 567:
			return `f`, H
		case r < 600:
			return `v`, H
		case r < 631:
			return `m`, H
		case r < 653:
			return `g`, H
		case r < 674:
			return `w`, H
		case r < 694:
			return `r`, H
		case r < 710:
			return `b`, H
		case r < 718:
			return `o`, H
		case r < 724:
			return `h`, H
		case r < 729:
			return `q`, H
		case r < 733:
			return `i`, H
		case r < 737:
			return `u`, H
		case r < 740:
			return `j`, H
		case r < 742:
			return `k`, H
		case r < 743:
			return `y`, H
		case r < 744:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	case ``:
		H := 4.2128960043149775
		r := rng.IntN(7776)
		switch {
		case r < 1087:
			return `s`, H
		case r < 1838:
			return `c`, H
		case r < 2416:
			return `p`, H
		case r < 2963:
			return `d`, H
		case r < 3476:
			return `r`, H
		case r < 3952:
			return `u`, H
		case r < 4359:
			return `a`, H
		case r < 4757:
			return `e`, H
		case r < 5119:
			return `b`, H
		case r < 5442:
			return `t`, H
		case r < 5750:
			return `f`, H
		case r < 6054:
			return `g`, H
		case r < 6349:
			return `m`, H
		case r < 6598:
			return `h`, H
		case r < 6844:
			return `o`, H
		case r < 7038:
			return `l`, H
		case r < 7193:
			return `w`, H
		case r < 7325:
			return `v`, H
		case r < 7440:
			return `i`, H
		case r < 7537:
			return `n`, H
		case r < 7633:
			return `j`, H
		case r < 7688:
			return `k`, H
		case r < 7725:
			return `q`, H
		case r < 7752:
			return `y`, H
		case r < 7774:
			return `z`, H
		case r < 7776:
			return `x`, H
		default:
			panic("unexpected rand num")
		}
	default:
		tok = tok[1:]
		goto retry
	}
}

// PickLength returns a random word length and its associated entropy.
// The word length is chosen based on a weighted random selection,
// with lengths ranging from 3 to 9 characters.
//
//   - Return values:
//     int: The selected word length in characters.
//     float64: The entropy contributed by the length selection process.
func PickLength() (int, float64) {
	r := rng.IntN(7776)
	H := 2.5404552183901115
	if r < 1779 {
		return 8, H
	} else if r < 3370 {
		return 7, H
	} else if r < 4926 {
		return 9, H
	} else if r < 6299 {
		return 6, H
	} else if r < 7226 {
		return 5, H
	} else if r < 7694 {
		return 4, H
	} else if r < 7776 {
		return 3, H
	}
	panic("unexpected rand num")
}
