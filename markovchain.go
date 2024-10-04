package cryptipass

// THIS FILE HAS BEEN DISTILLED FROM EFF's long word list, without their work this software would not exist.

// PickNext appends the next character to the current seed string based on specific rules.
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
	tok := seed[len(seed)-L:]
retry:
	switch tok {
	case `mh`:
		return seed + `o`, 0.0
	case `mv`:
		return seed + `e`, 0.0
	case `bf`:
		return seed + `l`, 0.0
	case `cq`:
		return seed + `u`, 0.0
	case `ii`:
		return seed + `n`, 0.0
	case `zm`:
		return seed + `o`, 0.0
	case `pk`:
		return seed + `i`, 0.0
	case `hd`:
		return seed + `a`, 0.0
	case `kd`:
		return seed + `r`, 0.0
	case `yz`:
		return seed + `e`, 0.0
	case `bz`:
		return seed + `e`, 0.0
	case `xb`:
		return seed + `o`, 0.0
	case `uk`:
		return seed + `e`, 0.0
	case `q`:
		return seed + `u`, 0.0
	case `bj`:
		return seed + `e`, 0.0
	case `yu`:
		return seed + `m`, 0.0
	case `kf`:
		return seed + `i`, 0.0
	case `ww`:
		return seed + `o`, 0.0
	case `oq`:
		return seed + `u`, 0.0
	case `wg`:
		return seed + `i`, 0.0
	case `gz`:
		return seed + `a`, 0.0
	case `oj`:
		return seed + `e`, 0.0
	case `uj`:
		return seed + `i`, 0.0
	case `wc`:
		return seed + `a`, 0.0
	case `rx`:
		return seed + `i`, 0.0
	case `bg`:
		return seed + `r`, 0.0
	case `fn`:
		return seed + `e`, 0.0
	case `sj`:
		return seed + `o`, 0.0
	case `vr`:
		return seed + `o`, 0.0
	case `zy`:
		return seed + `m`, 0.0
	case `km`:
		return seed + `a`, 0.0
	case `pg`:
		return seed + `r`, 0.0
	case `wt`:
		return seed + `i`, 0.0
	case `uv`:
		return seed + `e`, 0.0
	case `gf`:
		return seed + `u`, 0.0
	case `sq`:
		return seed + `u`, 0.0
	case `xq`:
		return seed + `u`, 0.0
	case `cs`:
		return seed + `i`, 0.0
	case `dt`:
		return seed + `h`, 0.0
	case `bc`:
		return seed + `a`, 0.0
	case `xf`:
		return seed + `o`, 0.0
	case `bv`:
		return seed + `i`, 0.0
	case `sg`:
		return seed + `r`, 0.0
	case `md`:
		return seed + `r`, 0.0
	case `eq`:
		return seed + `u`, 0.0
	case `wk`:
		return seed + `w`, 0.0
	case `dv`:
		return seed + `i`, 0.0
	case `xs`:
		return seed + `e`, 0.0
	case `wu`:
		return seed + `n`, 0.0
	case `pm`:
		return seed + `e`, 0.0
	case `iu`:
		return seed + `m`, 0.0
	case `dk`:
		return seed + `i`, 0.0
	case `hs`:
		return seed + `t`, 0.0
	case `hk`:
		return seed + `i`, 0.0
	case `lh`:
		return seed + `o`, 0.0
	case `lz`:
		return seed + `o`, 0.0
	case `mw`:
		return seed + `e`, 0.0
	case `gd`:
		return seed + `o`, 0.0
	case `uf`:
		return seed + `f`, 0.0
	case `xl`:
		return seed + `i`, 0.0
	case `iq`:
		return seed + `u`, 0.0
	case `aq`:
		return seed + `u`, 0.0
	case `hh`:
		return seed + `o`, 0.0
	case `yh`:
		return seed + `o`, 0.0
	case `yk`:
		return seed + `e`, 0.0
	case `gt`:
		return seed + `h`, 0.0
	case `nq`:
		return seed + `u`, 0.0
	case `hm`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `a`, H
		case r < 2:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ej`:
		H := 0.9709505944546686
		r := rng.IntN(5)
		switch {
		case r < 3:
			return seed + `o`, H
		case r < 5:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `uz`:
		H := 0.5916727785823275
		r := rng.IntN(7)
		switch {
		case r < 6:
			return seed + `z`, H
		case r < 7:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `hn`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `e`, H
		case r < 3:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `my`:
		H := 0.6500224216483541
		r := rng.IntN(6)
		switch {
		case r < 5:
			return seed + `s`, H
		case r < 6:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `kh`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `o`, H
		case r < 3:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `hp`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `i`, H
		case r < 2:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `pw`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `kk`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `n`, H
		case r < 2:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `by`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `t`, H
		case r < 3:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `bh`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `e`, H
		case r < 2:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gy`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `r`, H
		case r < 2:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `xy`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `g`, H
		case r < 2:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ln`:
		H := 0.9182958340544896
		r := rng.IntN(6)
		switch {
		case r < 4:
			return seed + `e`, H
		case r < 6:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lw`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `uo`:
		H := 0.9910760598382222
		r := rng.IntN(9)
		switch {
		case r < 5:
			return seed + `u`, H
		case r < 9:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `gm`:
		H := 1.0
		r := rng.IntN(8)
		switch {
		case r < 4:
			return seed + `a`, H
		case r < 8:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `pn`:
		H := 0.9940302114769566
		r := rng.IntN(11)
		switch {
		case r < 6:
			return seed + `o`, H
		case r < 11:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `tz`:
		H := 0.9709505944546686
		r := rng.IntN(5)
		switch {
		case r < 3:
			return seed + `y`, H
		case r < 5:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `wl`:
		H := 0.9709505944546686
		r := rng.IntN(10)
		switch {
		case r < 6:
			return seed + `e`, H
		case r < 10:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `sd`:
		H := 0.8112781244591328
		r := rng.IntN(4)
		switch {
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `sr`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `u`, H
		case r < 2:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `xh`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lk`:
		H := 0.8112781244591328
		r := rng.IntN(4)
		switch {
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `x`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `e`, H
		case r < 2:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `bn`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `o`, H
		case r < 3:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `dn`:
		H := 0.41381685030363374
		r := rng.IntN(12)
		switch {
		case r < 11:
			return seed + `e`, H
		case r < 12:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ae`:
		H := 0.6500224216483541
		r := rng.IntN(6)
		switch {
		case r < 5:
			return seed + `r`, H
		case r < 6:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ux`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `u`, H
		case r < 3:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `fb`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `a`, H
		case r < 2:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `hw`:
		H := 0.9182958340544896
		r := rng.IntN(3)
		switch {
		case r < 2:
			return seed + `o`, H
		case r < 3:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `wm`:
		H := 0.7219280948873623
		r := rng.IntN(5)
		switch {
		case r < 4:
			return seed + `a`, H
		case r < 5:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `pd`:
		H := 0.8112781244591328
		r := rng.IntN(4)
		switch {
		case r < 3:
			return seed + `o`, H
		case r < 4:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `td`:
		H := 0.7219280948873623
		r := rng.IntN(5)
		switch {
		case r < 4:
			return seed + `o`, H
		case r < 5:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `iw`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `i`, H
		case r < 2:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ih`:
		H := 1.0
		r := rng.IntN(2)
		switch {
		case r < 1:
			return seed + `u`, H
		case r < 2:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `rj`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `u`, H
		case r < 3:
			return seed + `e`, H
		case r < 4:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `kr`:
		H := 1.3709505944546687
		r := rng.IntN(5)
		switch {
		case r < 3:
			return seed + `o`, H
		case r < 4:
			return seed + `y`, H
		case r < 5:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ko`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `a`, H
		case r < 2:
			return seed + `s`, H
		case r < 3:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `mc`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `h`, H
		case r < 2:
			return seed + `e`, H
		case r < 3:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `kw`:
		H := 1.2987949406953985
		r := rng.IntN(8)
		switch {
		case r < 5:
			return seed + `a`, H
		case r < 7:
			return seed + `o`, H
		case r < 8:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `oh`:
		H := 1.3709505944546687
		r := rng.IntN(5)
		switch {
		case r < 3:
			return seed + `e`, H
		case r < 4:
			return seed + `n`, H
		case r < 5:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `bm`:
		H := 1.4591479170272448
		r := rng.IntN(6)
		switch {
		case r < 3:
			return seed + `e`, H
		case r < 5:
			return seed + `i`, H
		case r < 6:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `hb`:
		H := 1.3485878960124222
		r := rng.IntN(11)
		switch {
		case r < 5:
			return seed + `a`, H
		case r < 10:
			return seed + `o`, H
		case r < 11:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `gb`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `o`, H
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `lv`:
		H := 1.2467052127325735
		r := rng.IntN(21)
		switch {
		case r < 14:
			return seed + `e`, H
		case r < 18:
			return seed + `a`, H
		case r < 21:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `hc`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `h`, H
		case r < 4:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `aj`:
		H := 1.584962500721156
		r := rng.IntN(6)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 4:
			return seed + `e`, H
		case r < 6:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `mr`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `o`, H
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `nj`:
		H := 1.280672129520887
		r := rng.IntN(12)
		switch {
		case r < 7:
			return seed + `o`, H
		case r < 11:
			return seed + `u`, H
		case r < 12:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `bp`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `l`, H
		case r < 4:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `mn`:
		H := 1.4056390622295665
		r := rng.IntN(8)
		switch {
		case r < 4:
			return seed + `e`, H
		case r < 7:
			return seed + `i`, H
		case r < 8:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `py`:
		H := 1.3709505944546687
		r := rng.IntN(5)
		switch {
		case r < 3:
			return seed + `r`, H
		case r < 4:
			return seed + `g`, H
		case r < 5:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `yd`:
		H := 1.1887218755408673
		r := rng.IntN(12)
		switch {
		case r < 8:
			return seed + `r`, H
		case r < 11:
			return seed + `a`, H
		case r < 12:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `tg`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `e`, H
		case r < 2:
			return seed + `r`, H
		case r < 3:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `pb`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `e`, H
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `mf`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `u`, H
		case r < 3:
			return seed + `y`, H
		case r < 4:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ao`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `s`, H
		case r < 2:
			return seed + `r`, H
		case r < 3:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `kp`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `a`, H
		case r < 2:
			return seed + `e`, H
		case r < 3:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gp`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `i`, H
		case r < 2:
			return seed + `l`, H
		case r < 3:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `zl`:
		H := 1.1401156785146092
		r := rng.IntN(13)
		switch {
		case r < 9:
			return seed + `e`, H
		case r < 12:
			return seed + `i`, H
		case r < 13:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `pc`:
		H := 1.2516291673878228
		r := rng.IntN(6)
		switch {
		case r < 4:
			return seed + `o`, H
		case r < 5:
			return seed + `h`, H
		case r < 6:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `pf`:
		H := 1.5219280948873621
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `i`, H
		case r < 4:
			return seed + `u`, H
		case r < 5:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ek`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `w`, H
		case r < 2:
			return seed + `i`, H
		case r < 3:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ah`:
		H := 1.3787834934861756
		r := rng.IntN(7)
		switch {
		case r < 4:
			return seed + `o`, H
		case r < 6:
			return seed + `e`, H
		case r < 7:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `xu`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `r`, H
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `bw`:
		H := 1.584962500721156
		r := rng.IntN(3)
		switch {
		case r < 1:
			return seed + `a`, H
		case r < 2:
			return seed + `e`, H
		case r < 3:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `kt`:
		H := 1.5219280948873621
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 4:
			return seed + `o`, H
		case r < 5:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `wp`:
		H := 1.5
		r := rng.IntN(4)
		switch {
		case r < 2:
			return seed + `l`, H
		case r < 3:
			return seed + `i`, H
		case r < 4:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `dm`:
		H := 1.3921472236645345
		r := rng.IntN(9)
		switch {
		case r < 4:
			return seed + `a`, H
		case r < 8:
			return seed + `i`, H
		case r < 9:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `sb`:
		H := 1.9182958340544893
		r := rng.IntN(6)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 4:
			return seed + `e`, H
		case r < 5:
			return seed + `u`, H
		case r < 6:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `cy`:
		H := 1.3520301017579528
		r := rng.IntN(13)
		switch {
		case r < 9:
			return seed + `c`, H
		case r < 11:
			return seed + `t`, H
		case r < 12:
			return seed + `l`, H
		case r < 13:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `lr`:
		H := 1.9182958340544893
		r := rng.IntN(6)
		switch {
		case r < 2:
			return seed + `y`, H
		case r < 4:
			return seed + `i`, H
		case r < 5:
			return seed + `u`, H
		case r < 6:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `yf`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `u`, H
		case r < 3:
			return seed + `i`, H
		case r < 4:
			return seed + `l`, H
		case r < 5:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `df`:
		H := 1.7841591278514217
		r := rng.IntN(12)
		switch {
		case r < 5:
			return seed + `i`, H
		case r < 9:
			return seed + `u`, H
		case r < 11:
			return seed + `a`, H
		case r < 12:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `wb`:
		H := 1.9502120649147472
		r := rng.IntN(7)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 4:
			return seed + `i`, H
		case r < 6:
			return seed + `o`, H
		case r < 7:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `yg`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `r`, H
		case r < 3:
			return seed + `l`, H
		case r < 4:
			return seed + `e`, H
		case r < 5:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gh`:
		H := 0.44363159167676336
		r := rng.IntN(47)
		switch {
		case r < 44:
			return seed + `t`, H
		case r < 45:
			return seed + `a`, H
		case r < 46:
			return seed + `y`, H
		case r < 47:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `z`:
		H := 1.7702834614222898
		r := rng.IntN(22)
		switch {
		case r < 8:
			return seed + `o`, H
		case r < 15:
			return seed + `e`, H
		case r < 21:
			return seed + `i`, H
		case r < 22:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `yi`:
		H := 0.719556365373903
		r := rng.IntN(25)
		switch {
		case r < 22:
			return seed + `n`, H
		case r < 23:
			return seed + `d`, H
		case r < 24:
			return seed + `e`, H
		case r < 25:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `zz`:
		H := 1.478897902987479
		r := rng.IntN(20)
		switch {
		case r < 13:
			return seed + `l`, H
		case r < 16:
			return seed + `i`, H
		case r < 18:
			return seed + `a`, H
		case r < 20:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `gn`:
		H := 1.9584661817581863
		r := rng.IntN(37)
		switch {
		case r < 13:
			return seed + `e`, H
		case r < 22:
			return seed + `a`, H
		case r < 30:
			return seed + `i`, H
		case r < 37:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ak`:
		H := 1.0592982590819013
		r := rng.IntN(55)
		switch {
		case r < 41:
			return seed + `e`, H
		case r < 52:
			return seed + `i`, H
		case r < 54:
			return seed + `y`, H
		case r < 55:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `kn`:
		H := 1.8412501702815551
		r := rng.IntN(18)
		switch {
		case r < 7:
			return seed + `e`, H
		case r < 13:
			return seed + `o`, H
		case r < 16:
			return seed + `i`, H
		case r < 18:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `yc`:
		H := 1.4838317068446891
		r := rng.IntN(14)
		switch {
		case r < 9:
			return seed + `l`, H
		case r < 11:
			return seed + `h`, H
		case r < 13:
			return seed + `a`, H
		case r < 14:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `m`:
		H := 1.596319592368159
		r := rng.IntN(295)
		switch {
		case r < 130:
			return seed + `a`, H
		case r < 237:
			return seed + `o`, H
		case r < 290:
			return seed + `u`, H
		case r < 295:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `yw`:
		H := 1.6644977792004614
		r := rng.IntN(7)
		switch {
		case r < 4:
			return seed + `a`, H
		case r < 5:
			return seed + `h`, H
		case r < 6:
			return seed + `i`, H
		case r < 7:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gs`:
		H := 1.811278124459133
		r := rng.IntN(8)
		switch {
		case r < 3:
			return seed + `h`, H
		case r < 6:
			return seed + `t`, H
		case r < 7:
			return seed + `a`, H
		case r < 8:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `qu`:
		H := 1.7428655983817265
		r := rng.IntN(99)
		switch {
		case r < 37:
			return seed + `i`, H
		case r < 71:
			return seed + `a`, H
		case r < 95:
			return seed + `e`, H
		case r < 99:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `xo`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `r`, H
		case r < 3:
			return seed + `n`, H
		case r < 4:
			return seed + `d`, H
		case r < 5:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `h`:
		H := 1.7259650053109339
		r := rng.IntN(249)
		switch {
		case r < 122:
			return seed + `a`, H
		case r < 184:
			return seed + `e`, H
		case r < 232:
			return seed + `u`, H
		case r < 249:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `kb`:
		H := 1.6644977792004614
		r := rng.IntN(7)
		switch {
		case r < 4:
			return seed + `o`, H
		case r < 5:
			return seed + `a`, H
		case r < 6:
			return seed + `i`, H
		case r < 7:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `tm`:
		H := 1.9056390622295665
		r := rng.IntN(8)
		switch {
		case r < 3:
			return seed + `e`, H
		case r < 5:
			return seed + `a`, H
		case r < 7:
			return seed + `o`, H
		case r < 8:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lg`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `e`, H
		case r < 3:
			return seed + `i`, H
		case r < 4:
			return seed + `u`, H
		case r < 5:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `yl`:
		H := 1.4516607643577668
		r := rng.IntN(17)
		switch {
		case r < 11:
			return seed + `i`, H
		case r < 14:
			return seed + `e`, H
		case r < 16:
			return seed + `o`, H
		case r < 17:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ik`:
		H := 1.0174051189440556
		r := rng.IntN(34)
		switch {
		case r < 27:
			return seed + `e`, H
		case r < 31:
			return seed + `i`, H
		case r < 33:
			return seed + `a`, H
		case r < 34:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `yt`:
		H := 1.9362600275315274
		r := rng.IntN(11)
		switch {
		case r < 4:
			return seed + `h`, H
		case r < 7:
			return seed + `i`, H
		case r < 9:
			return seed + `e`, H
		case r < 11:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `wd`:
		H := 1.8365916681089791
		r := rng.IntN(9)
		switch {
		case r < 4:
			return seed + `e`, H
		case r < 6:
			return seed + `l`, H
		case r < 8:
			return seed + `r`, H
		case r < 9:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ls`:
		H := 1.8828560636920488
		r := rng.IntN(16)
		switch {
		case r < 6:
			return seed + `i`, H
		case r < 11:
			return seed + `e`, H
		case r < 14:
			return seed + `a`, H
		case r < 16:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `dh`:
		H := 1.9182958340544893
		r := rng.IntN(6)
		switch {
		case r < 2:
			return seed + `i`, H
		case r < 4:
			return seed + `e`, H
		case r < 5:
			return seed + `a`, H
		case r < 6:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `yr`:
		H := 1.8631205685666312
		r := rng.IntN(14)
		switch {
		case r < 5:
			return seed + `i`, H
		case r < 10:
			return seed + `o`, H
		case r < 12:
			return seed + `a`, H
		case r < 14:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rh`:
		H := 1.676737030052133
		r := rng.IntN(11)
		switch {
		case r < 5:
			return seed + `e`, H
		case r < 9:
			return seed + `a`, H
		case r < 10:
			return seed + `u`, H
		case r < 11:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `nz`:
		H := 1.8423709931771088
		r := rng.IntN(7)
		switch {
		case r < 3:
			return seed + `i`, H
		case r < 5:
			return seed + `y`, H
		case r < 6:
			return seed + `a`, H
		case r < 7:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `wn`:
		H := 2.0
		r := rng.IntN(4)
		switch {
		case r < 1:
			return seed + `n`, H
		case r < 2:
			return seed + `y`, H
		case r < 3:
			return seed + `i`, H
		case r < 4:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `wf`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `i`, H
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `l`, H
		case r < 5:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `tf`:
		H := 1.6589797521242051
		r := rng.IntN(17)
		switch {
		case r < 9:
			return seed + `u`, H
		case r < 13:
			return seed + `i`, H
		case r < 16:
			return seed + `o`, H
		case r < 17:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `rw`:
		H := 1.9056390622295665
		r := rng.IntN(8)
		switch {
		case r < 3:
			return seed + `i`, H
		case r < 5:
			return seed + `a`, H
		case r < 7:
			return seed + `e`, H
		case r < 8:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `gw`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `h`, H
		case r < 4:
			return seed + `e`, H
		case r < 5:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `bd`:
		H := 1.8423709931771088
		r := rng.IntN(7)
		switch {
		case r < 3:
			return seed + `o`, H
		case r < 5:
			return seed + `u`, H
		case r < 6:
			return seed + `i`, H
		case r < 7:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `v`:
		H := 1.919393986516576
		r := rng.IntN(132)
		switch {
		case r < 47:
			return seed + `i`, H
		case r < 81:
			return seed + `a`, H
		case r < 115:
			return seed + `e`, H
		case r < 132:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ku`:
		H := 1.9219280948873623
		r := rng.IntN(5)
		switch {
		case r < 2:
			return seed + `p`, H
		case r < 3:
			return seed + `n`, H
		case r < 4:
			return seed + `d`, H
		case r < 5:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `nm`:
		H := 1.6758220575766962
		r := rng.IntN(18)
		switch {
		case r < 9:
			return seed + `a`, H
		case r < 14:
			return seed + `o`, H
		case r < 17:
			return seed + `i`, H
		case r < 18:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `tw`:
		H := 1.8832406052365576
		r := rng.IntN(39)
		switch {
		case r < 18:
			return seed + `i`, H
		case r < 28:
			return seed + `e`, H
		case r < 35:
			return seed + `a`, H
		case r < 38:
			return seed + `o`, H
		case r < 39:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `tl`:
		H := 1.7570627627948316
		r := rng.IntN(95)
		switch {
		case r < 39:
			return seed + `e`, H
		case r < 77:
			return seed + `y`, H
		case r < 88:
			return seed + `i`, H
		case r < 93:
			return seed + `a`, H
		case r < 95:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `br`:
		H := 2.2142644095851374
		r := rng.IntN(87)
		switch {
		case r < 26:
			return seed + `o`, H
		case r < 49:
			return seed + `i`, H
		case r < 67:
			return seed + `e`, H
		case r < 78:
			return seed + `a`, H
		case r < 87:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `nh`:
		H := 2.2068570640942182
		r := rng.IntN(23)
		switch {
		case r < 6:
			return seed + `a`, H
		case r < 12:
			return seed + `e`, H
		case r < 18:
			return seed + `o`, H
		case r < 21:
			return seed + `i`, H
		case r < 23:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `y`:
		H := 2.0793752444772653
		r := rng.IntN(27)
		switch {
		case r < 10:
			return seed + `e`, H
		case r < 16:
			return seed + `a`, H
		case r < 22:
			return seed + `o`, H
		case r < 26:
			return seed + `i`, H
		case r < 27:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `tp`:
		H := 2.235926350629033
		r := rng.IntN(7)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 4:
			return seed + `o`, H
		case r < 5:
			return seed + `l`, H
		case r < 6:
			return seed + `r`, H
		case r < 7:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rv`:
		H := 1.915924784727185
		r := rng.IntN(34)
		switch {
		case r < 14:
			return seed + `i`, H
		case r < 26:
			return seed + `e`, H
		case r < 29:
			return seed + `a`, H
		case r < 32:
			return seed + `y`, H
		case r < 34:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `tc`:
		H := 0.9954758876481817
		r := rng.IntN(59)
		switch {
		case r < 48:
			return seed + `h`, H
		case r < 54:
			return seed + `a`, H
		case r < 57:
			return seed + `o`, H
		case r < 58:
			return seed + `l`, H
		case r < 59:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ix`:
		H := 2.054585169337799
		r := rng.IntN(12)
		switch {
		case r < 5:
			return seed + `t`, H
		case r < 8:
			return seed + `e`, H
		case r < 10:
			return seed + `a`, H
		case r < 11:
			return seed + `f`, H
		case r < 12:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `ji`:
		H := 2.113283334294875
		r := rng.IntN(9)
		switch {
		case r < 3:
			return seed + `n`, H
		case r < 6:
			return seed + `t`, H
		case r < 7:
			return seed + `f`, H
		case r < 8:
			return seed + `g`, H
		case r < 9:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `j`:
		H := 2.077766251414667
		r := rng.IntN(96)
		switch {
		case r < 36:
			return seed + `u`, H
		case r < 62:
			return seed + `a`, H
		case r < 81:
			return seed + `o`, H
		case r < 89:
			return seed + `i`, H
		case r < 96:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `nv`:
		H := 1.962354483657865
		r := rng.IntN(22)
		switch {
		case r < 9:
			return seed + `i`, H
		case r < 16:
			return seed + `e`, H
		case r < 19:
			return seed + `a`, H
		case r < 21:
			return seed + `o`, H
		case r < 22:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `xe`:
		H := 2.0419460322060456
		r := rng.IntN(15)
		switch {
		case r < 5:
			return seed + `d`, H
		case r < 10:
			return seed + `r`, H
		case r < 13:
			return seed + `m`, H
		case r < 14:
			return seed + `n`, H
		case r < 15:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `sn`:
		H := 1.9664658234492332
		r := rng.IntN(48)
		switch {
		case r < 24:
			return seed + `o`, H
		case r < 33:
			return seed + `a`, H
		case r < 39:
			return seed + `i`, H
		case r < 44:
			return seed + `u`, H
		case r < 48:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `dg`:
		H := 1.3957137264520418
		r := rng.IntN(33)
		switch {
		case r < 23:
			return seed + `e`, H
		case r < 28:
			return seed + `i`, H
		case r < 31:
			return seed + `y`, H
		case r < 32:
			return seed + `r`, H
		case r < 33:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `dp`:
		H := 2.197159723424149
		r := rng.IntN(9)
		switch {
		case r < 3:
			return seed + `i`, H
		case r < 5:
			return seed + `a`, H
		case r < 7:
			return seed + `o`, H
		case r < 8:
			return seed + `h`, H
		case r < 9:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `b`:
		H := 2.2296839672243163
		r := rng.IntN(362)
		switch {
		case r < 126:
			return seed + `a`, H
		case r < 197:
			return seed + `o`, H
		case r < 254:
			return seed + `r`, H
		case r < 309:
			return seed + `u`, H
		case r < 362:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `dc`:
		H := 2.038159681645953
		r := rng.IntN(13)
		switch {
		case r < 5:
			return seed + `a`, H
		case r < 9:
			return seed + `o`, H
		case r < 11:
			return seed + `l`, H
		case r < 12:
			return seed + `r`, H
		case r < 13:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `lm`:
		H := 2.197159723424149
		r := rng.IntN(9)
		switch {
		case r < 3:
			return seed + `a`, H
		case r < 5:
			return seed + `i`, H
		case r < 7:
			return seed + `o`, H
		case r < 8:
			return seed + `y`, H
		case r < 9:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ok`:
		H := 1.546361197390146
		r := rng.IntN(36)
		switch {
		case r < 23:
			return seed + `e`, H
		case r < 30:
			return seed + `i`, H
		case r < 33:
			return seed + `y`, H
		case r < 35:
			return seed + `a`, H
		case r < 36:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `n`:
		H := 1.9834964157590933
		r := rng.IntN(97)
		switch {
		case r < 33:
			return seed + `e`, H
		case r < 59:
			return seed + `u`, H
		case r < 84:
			return seed + `a`, H
		case r < 96:
			return seed + `i`, H
		case r < 97:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `yn`:
		H := 2.1867043459100244
		r := rng.IntN(11)
		switch {
		case r < 4:
			return seed + `a`, H
		case r < 6:
			return seed + `e`, H
		case r < 8:
			return seed + `t`, H
		case r < 10:
			return seed + `o`, H
		case r < 11:
			return seed + `d`, H
		default:
			panic("unexpected rand num")
		}
	case `hr`:
		H := 2.2129220962815555
		r := rng.IntN(43)
		switch {
		case r < 14:
			return seed + `o`, H
		case r < 24:
			return seed + `i`, H
		case r < 32:
			return seed + `e`, H
		case r < 39:
			return seed + `a`, H
		case r < 43:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ov`:
		H := 1.1045475199265447
		r := rng.IntN(160)
		switch {
		case r < 125:
			return seed + `e`, H
		case r < 141:
			return seed + `i`, H
		case r < 155:
			return seed + `a`, H
		case r < 158:
			return seed + `o`, H
		case r < 160:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ax`:
		H := 2.1280852788913944
		r := rng.IntN(7)
		switch {
		case r < 3:
			return seed + `i`, H
		case r < 4:
			return seed + `s`, H
		case r < 5:
			return seed + `a`, H
		case r < 6:
			return seed + `e`, H
		case r < 7:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `wh`:
		H := 2.131346433249389
		r := rng.IntN(25)
		switch {
		case r < 7:
			return seed + `i`, H
		case r < 14:
			return seed + `e`, H
		case r < 20:
			return seed + `o`, H
		case r < 24:
			return seed + `a`, H
		case r < 25:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ws`:
		H := 2.2516291673878226
		r := rng.IntN(6)
		switch {
		case r < 2:
			return seed + `i`, H
		case r < 3:
			return seed + `h`, H
		case r < 4:
			return seed + `e`, H
		case r < 5:
			return seed + `t`, H
		case r < 6:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `oz`:
		H := 2.0550365325772657
		r := rng.IntN(16)
		switch {
		case r < 6:
			return seed + `e`, H
		case r < 11:
			return seed + `y`, H
		case r < 13:
			return seed + `i`, H
		case r < 15:
			return seed + `o`, H
		case r < 16:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `ym`:
		H := 2.104292989590925
		r := rng.IntN(19)
		switch {
		case r < 6:
			return seed + `e`, H
		case r < 11:
			return seed + `a`, H
		case r < 16:
			return seed + `p`, H
		case r < 18:
			return seed + `o`, H
		case r < 19:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `lf`:
		H := 2.075869296697727
		r := rng.IntN(13)
		switch {
		case r < 5:
			return seed + `i`, H
		case r < 8:
			return seed + `a`, H
		case r < 11:
			return seed + `u`, H
		case r < 12:
			return seed + `r`, H
		case r < 13:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `dw`:
		H := 2.1645691489449588
		r := rng.IntN(19)
		switch {
		case r < 5:
			return seed + `a`, H
		case r < 10:
			return seed + `e`, H
		case r < 15:
			return seed + `o`, H
		case r < 18:
			return seed + `i`, H
		case r < 19:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `eh`:
		H := 2.2577560641285177
		r := rng.IntN(13)
		switch {
		case r < 4:
			return seed + `e`, H
		case r < 7:
			return seed + `i`, H
		case r < 9:
			return seed + `y`, H
		case r < 11:
			return seed + `a`, H
		case r < 13:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `nw`:
		H := 2.17675266977692
		r := rng.IntN(27)
		switch {
		case r < 8:
			return seed + `a`, H
		case r < 16:
			return seed + `o`, H
		case r < 21:
			return seed + `i`, H
		case r < 25:
			return seed + `e`, H
		case r < 27:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `sw`:
		H := 2.046381112945884
		r := rng.IntN(42)
		switch {
		case r < 16:
			return seed + `i`, H
		case r < 25:
			return seed + `e`, H
		case r < 33:
			return seed + `a`, H
		case r < 41:
			return seed + `o`, H
		case r < 42:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `db`:
		H := 1.8322488896002307
		r := rng.IntN(14)
		switch {
		case r < 7:
			return seed + `a`, H
		case r < 11:
			return seed + `o`, H
		case r < 12:
			return seed + `l`, H
		case r < 13:
			return seed + `r`, H
		case r < 14:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `l`:
		H := 2.0699571885736554
		r := rng.IntN(194)
		switch {
		case r < 63:
			return seed + `a`, H
		case r < 119:
			return seed + `i`, H
		case r < 157:
			return seed + `u`, H
		case r < 189:
			return seed + `e`, H
		case r < 194:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `nr`:
		H := 2.1468032139556708
		r := rng.IntN(27)
		switch {
		case r < 11:
			return seed + `e`, H
		case r < 16:
			return seed + `i`, H
		case r < 20:
			return seed + `a`, H
		case r < 24:
			return seed + `o`, H
		case r < 27:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `za`:
		H := 1.896240625180289
		r := rng.IntN(12)
		switch {
		case r < 6:
			return seed + `r`, H
		case r < 9:
			return seed + `b`, H
		case r < 10:
			return seed + `g`, H
		case r < 11:
			return seed + `c`, H
		case r < 12:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `iz`:
		H := 1.3054506823200551
		r := rng.IntN(69)
		switch {
		case r < 50:
			return seed + `e`, H
		case r < 60:
			return seed + `z`, H
		case r < 64:
			return seed + `a`, H
		case r < 68:
			return seed + `i`, H
		case r < 69:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `wr`:
		H := 2.1229853836154926
		r := rng.IntN(29)
		switch {
		case r < 11:
			return seed + `i`, H
		case r < 18:
			return seed + `o`, H
		case r < 23:
			return seed + `e`, H
		case r < 27:
			return seed + `a`, H
		case r < 29:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `iv`:
		H := 1.5989577717916381
		r := rng.IntN(161)
		switch {
		case r < 91:
			return seed + `e`, H
		case r < 131:
			return seed + `i`, H
		case r < 153:
			return seed + `a`, H
		case r < 160:
			return seed + `o`, H
		case r < 161:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `hy`:
		H := 1.6444492199495462
		r := rng.IntN(21)
		switch {
		case r < 11:
			return seed + `p`, H
		case r < 18:
			return seed + `d`, H
		case r < 19:
			return seed + `s`, H
		case r < 20:
			return seed + `m`, H
		case r < 21:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `tn`:
		H := 1.3862740938959681
		r := rng.IntN(18)
		switch {
		case r < 13:
			return seed + `e`, H
		case r < 15:
			return seed + `a`, H
		case r < 16:
			return seed + `i`, H
		case r < 17:
			return seed + `u`, H
		case r < 18:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `cl`:
		H := 2.273319510050957
		r := rng.IntN(123)
		switch {
		case r < 36:
			return seed + `a`, H
		case r < 64:
			return seed + `e`, H
		case r < 84:
			return seed + `i`, H
		case r < 104:
			return seed + `u`, H
		case r < 123:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `av`:
		H := 1.9476716502007156
		r := rng.IntN(96)
		switch {
		case r < 41:
			return seed + `e`, H
		case r < 69:
			return seed + `i`, H
		case r < 82:
			return seed + `a`, H
		case r < 93:
			return seed + `o`, H
		case r < 96:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `w`:
		H := 2.220342743810504
		r := rng.IntN(155)
		switch {
		case r < 47:
			return seed + `i`, H
		case r < 89:
			return seed + `a`, H
		case r < 117:
			return seed + `o`, H
		case r < 137:
			return seed + `r`, H
		case r < 155:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `dl`:
		H := 1.9282661532133314
		r := rng.IntN(82)
		switch {
		case r < 37:
			return seed + `e`, H
		case r < 58:
			return seed + `y`, H
		case r < 73:
			return seed + `i`, H
		case r < 78:
			return seed + `o`, H
		case r < 82:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `hl`:
		H := 1.9986525468781666
		r := rng.IntN(17)
		switch {
		case r < 8:
			return seed + `y`, H
		case r < 11:
			return seed + `i`, H
		case r < 14:
			return seed + `o`, H
		case r < 16:
			return seed + `e`, H
		case r < 17:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `kl`:
		H := 1.710053567474341
		r := rng.IntN(31)
		switch {
		case r < 18:
			return seed + `e`, H
		case r < 24:
			return seed + `i`, H
		case r < 28:
			return seed + `y`, H
		case r < 30:
			return seed + `a`, H
		case r < 31:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ys`:
		H := 1.7688538376160001
		r := rng.IntN(18)
		switch {
		case r < 10:
			return seed + `t`, H
		case r < 14:
			return seed + `e`, H
		case r < 16:
			return seed + `l`, H
		case r < 17:
			return seed + `p`, H
		case r < 18:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `ml`:
		H := 1.987773371487984
		r := rng.IntN(13)
		switch {
		case r < 6:
			return seed + `e`, H
		case r < 9:
			return seed + `y`, H
		case r < 11:
			return seed + `i`, H
		case r < 12:
			return seed + `a`, H
		case r < 13:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `sm`:
		H := 2.1119613706025886
		r := rng.IntN(53)
		switch {
		case r < 18:
			return seed + `a`, H
		case r < 32:
			return seed + `o`, H
		case r < 44:
			return seed + `i`, H
		case r < 50:
			return seed + `u`, H
		case r < 53:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `r`:
		H := 1.5668118547237875
		r := rng.IntN(513)
		switch {
		case r < 345:
			return seed + `e`, H
		case r < 401:
			return seed + `a`, H
		case r < 445:
			return seed + `i`, H
		case r < 484:
			return seed + `o`, H
		case r < 512:
			return seed + `u`, H
		case r < 513:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `cr`:
		H := 2.394320692780727
		r := rng.IntN(168)
		switch {
		case r < 42:
			return seed + `a`, H
		case r < 83:
			return seed + `e`, H
		case r < 112:
			return seed + `i`, H
		case r < 140:
			return seed + `u`, H
		case r < 164:
			return seed + `o`, H
		case r < 168:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `lc`:
		H := 2.3709505944546683
		r := rng.IntN(10)
		switch {
		case r < 3:
			return seed + `u`, H
		case r < 6:
			return seed + `o`, H
		case r < 7:
			return seed + `h`, H
		case r < 8:
			return seed + `a`, H
		case r < 9:
			return seed + `i`, H
		case r < 10:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `gl`:
		H := 2.4463501798243668
		r := rng.IntN(112)
		switch {
		case r < 35:
			return seed + `e`, H
		case r < 54:
			return seed + `a`, H
		case r < 72:
			return seed + `i`, H
		case r < 89:
			return seed + `y`, H
		case r < 105:
			return seed + `o`, H
		case r < 112:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sy`:
		H := 2.2584598927441624
		r := rng.IntN(18)
		switch {
		case r < 7:
			return seed + `n`, H
		case r < 11:
			return seed + `m`, H
		case r < 14:
			return seed + `s`, H
		case r < 16:
			return seed + `c`, H
		case r < 17:
			return seed + `r`, H
		case r < 18:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `bl`:
		H := 1.8303332451927128
		r := rng.IntN(293)
		switch {
		case r < 177:
			return seed + `e`, H
		case r < 213:
			return seed + `i`, H
		case r < 248:
			return seed + `y`, H
		case r < 267:
			return seed + `a`, H
		case r < 282:
			return seed + `o`, H
		case r < 293:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `zo`:
		H := 2.1920058354921066
		r := rng.IntN(13)
		switch {
		case r < 5:
			return seed + `o`, H
		case r < 9:
			return seed + `n`, H
		case r < 10:
			return seed + `a`, H
		case r < 11:
			return seed + `i`, H
		case r < 12:
			return seed + `d`, H
		case r < 13:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `pl`:
		H := 2.3540218435644924
		r := rng.IntN(178)
		switch {
		case r < 63:
			return seed + `a`, H
		case r < 99:
			return seed + `e`, H
		case r < 130:
			return seed + `i`, H
		case r < 153:
			return seed + `o`, H
		case r < 167:
			return seed + `y`, H
		case r < 178:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `fr`:
		H := 2.200881953295572
		r := rng.IntN(98)
		switch {
		case r < 33:
			return seed + `e`, H
		case r < 56:
			return seed + `a`, H
		case r < 76:
			return seed + `o`, H
		case r < 92:
			return seed + `i`, H
		case r < 97:
			return seed + `u`, H
		case r < 98:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ez`:
		H := 2.064042639445697
		r := rng.IntN(14)
		switch {
		case r < 7:
			return seed + `e`, H
		case r < 10:
			return seed + `i`, H
		case r < 11:
			return seed + `y`, H
		case r < 12:
			return seed + `a`, H
		case r < 13:
			return seed + `o`, H
		case r < 14:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `lb`:
		H := 2.5216406363433186
		r := rng.IntN(7)
		switch {
		case r < 2:
			return seed + `i`, H
		case r < 3:
			return seed + `a`, H
		case r < 4:
			return seed + `u`, H
		case r < 5:
			return seed + `r`, H
		case r < 6:
			return seed + `o`, H
		case r < 7:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `ky`:
		H := 2.5
		r := rng.IntN(8)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 4:
			return seed + `l`, H
		case r < 5:
			return seed + `w`, H
		case r < 6:
			return seed + `d`, H
		case r < 7:
			return seed + `p`, H
		case r < 8:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `rl`:
		H := 2.317113611025748
		r := rng.IntN(47)
		switch {
		case r < 15:
			return seed + `i`, H
		case r < 25:
			return seed + `y`, H
		case r < 35:
			return seed + `e`, H
		case r < 41:
			return seed + `o`, H
		case r < 46:
			return seed + `a`, H
		case r < 47:
			return seed + `d`, H
		default:
			panic("unexpected rand num")
		}
	case `f`:
		H := 2.5255137638602134
		r := rng.IntN(308)
		switch {
		case r < 76:
			return seed + `r`, H
		case r < 139:
			return seed + `l`, H
		case r < 193:
			return seed + `a`, H
		case r < 235:
			return seed + `o`, H
		case r < 276:
			return seed + `i`, H
		case r < 308:
			return seed + `e`, H
		default:
			panic("unexpected rand num")
		}
	case `np`:
		H := 2.2417812580773235
		r := rng.IntN(21)
		switch {
		case r < 7:
			return seed + `a`, H
		case r < 13:
			return seed + `l`, H
		case r < 17:
			return seed + `i`, H
		case r < 19:
			return seed + `o`, H
		case r < 20:
			return seed + `e`, H
		case r < 21:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `fl`:
		H := 2.1729752291612567
		r := rng.IntN(104)
		switch {
		case r < 47:
			return seed + `a`, H
		case r < 64:
			return seed + `e`, H
		case r < 80:
			return seed + `i`, H
		case r < 92:
			return seed + `o`, H
		case r < 101:
			return seed + `y`, H
		case r < 104:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `pr`:
		H := 1.9672224968373695
		r := rng.IntN(239)
		switch {
		case r < 99:
			return seed + `o`, H
		case r < 169:
			return seed + `e`, H
		case r < 214:
			return seed + `i`, H
		case r < 230:
			return seed + `a`, H
		case r < 237:
			return seed + `u`, H
		case r < 239:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `tb`:
		H := 2.1309260807303376
		r := rng.IntN(17)
		switch {
		case r < 7:
			return seed + `o`, H
		case r < 12:
			return seed + `a`, H
		case r < 14:
			return seed + `i`, H
		case r < 15:
			return seed + `e`, H
		case r < 16:
			return seed + `r`, H
		case r < 17:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `yb`:
		H := 2.4591479170272446
		r := rng.IntN(12)
		switch {
		case r < 3:
			return seed + `a`, H
		case r < 6:
			return seed + `o`, H
		case r < 8:
			return seed + `e`, H
		case r < 10:
			return seed + `r`, H
		case r < 11:
			return seed + `i`, H
		case r < 12:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sf`:
		H := 2.5216406363433186
		r := rng.IntN(7)
		switch {
		case r < 2:
			return seed + `i`, H
		case r < 3:
			return seed + `y`, H
		case r < 4:
			return seed + `u`, H
		case r < 5:
			return seed + `r`, H
		case r < 6:
			return seed + `e`, H
		case r < 7:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `bb`:
		H := 2.1580708663986847
		r := rng.IntN(52)
		switch {
		case r < 20:
			return seed + `l`, H
		case r < 35:
			return seed + `e`, H
		case r < 42:
			return seed + `i`, H
		case r < 48:
			return seed + `y`, H
		case r < 50:
			return seed + `a`, H
		case r < 52:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ms`:
		H := 2.4193819456463714
		r := rng.IntN(9)
		switch {
		case r < 3:
			return seed + `t`, H
		case r < 5:
			return seed + `i`, H
		case r < 6:
			return seed + `h`, H
		case r < 7:
			return seed + `y`, H
		case r < 8:
			return seed + `u`, H
		case r < 9:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `nn`:
		H := 2.203360784206382
		r := rng.IntN(74)
		switch {
		case r < 31:
			return seed + `e`, H
		case r < 47:
			return seed + `i`, H
		case r < 57:
			return seed + `a`, H
		case r < 65:
			return seed + `y`, H
		case r < 72:
			return seed + `o`, H
		case r < 74:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rr`:
		H := 2.4052364684207306
		r := rng.IntN(91)
		switch {
		case r < 27:
			return seed + `i`, H
		case r < 44:
			return seed + `e`, H
		case r < 60:
			return seed + `a`, H
		case r < 75:
			return seed + `o`, H
		case r < 88:
			return seed + `y`, H
		case r < 91:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `az`:
		H := 2.212331758086258
		r := rng.IntN(34)
		switch {
		case r < 13:
			return seed + `i`, H
		case r < 22:
			return seed + `e`, H
		case r < 27:
			return seed + `z`, H
		case r < 30:
			return seed + `y`, H
		case r < 33:
			return seed + `a`, H
		case r < 34:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `gr`:
		H := 2.042551738105775
		r := rng.IntN(159)
		switch {
		case r < 76:
			return seed + `a`, H
		case r < 104:
			return seed + `e`, H
		case r < 127:
			return seed + `o`, H
		case r < 147:
			return seed + `i`, H
		case r < 158:
			return seed + `u`, H
		case r < 159:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `je`:
		H := 2.4300365325772657
		r := rng.IntN(16)
		switch {
		case r < 5:
			return seed + `c`, H
		case r < 8:
			return seed + `s`, H
		case r < 11:
			return seed + `t`, H
		case r < 13:
			return seed + `l`, H
		case r < 15:
			return seed + `e`, H
		case r < 16:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `xa`:
		H := 2.5032583347756456
		r := rng.IntN(9)
		switch {
		case r < 2:
			return seed + `g`, H
		case r < 4:
			return seed + `m`, H
		case r < 6:
			return seed + `b`, H
		case r < 7:
			return seed + `c`, H
		case r < 8:
			return seed + `l`, H
		case r < 9:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `nl`:
		H := 2.3212021480233496
		r := rng.IntN(52)
		switch {
		case r < 19:
			return seed + `i`, H
		case r < 29:
			return seed + `e`, H
		case r < 38:
			return seed + `o`, H
		case r < 45:
			return seed + `y`, H
		case r < 50:
			return seed + `a`, H
		case r < 52:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sl`:
		H := 2.314430867489686
		r := rng.IntN(85)
		switch {
		case r < 26:
			return seed + `i`, H
		case r < 50:
			return seed + `a`, H
		case r < 63:
			return seed + `e`, H
		case r < 76:
			return seed + `o`, H
		case r < 81:
			return seed + `u`, H
		case r < 85:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `cc`:
		H := 1.9268680175017443
		r := rng.IntN(21)
		switch {
		case r < 12:
			return seed + `u`, H
		case r < 15:
			return seed + `o`, H
		case r < 17:
			return seed + `l`, H
		case r < 19:
			return seed + `e`, H
		case r < 20:
			return seed + `a`, H
		case r < 21:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `ox`:
		H := 1.8182626942811975
		r := rng.IntN(22)
		switch {
		case r < 13:
			return seed + `i`, H
		case r < 17:
			return seed + `y`, H
		case r < 19:
			return seed + `e`, H
		case r < 20:
			return seed + `f`, H
		case r < 21:
			return seed + `c`, H
		case r < 22:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `mm`:
		H := 2.4645728472008237
		r := rng.IntN(69)
		switch {
		case r < 19:
			return seed + `e`, H
		case r < 35:
			return seed + `o`, H
		case r < 46:
			return seed + `i`, H
		case r < 55:
			return seed + `y`, H
		case r < 64:
			return seed + `a`, H
		case r < 69:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `dr`:
		H := 2.258664230601291
		r := rng.IntN(97)
		switch {
		case r < 27:
			return seed + `a`, H
		case r < 50:
			return seed + `e`, H
		case r < 71:
			return seed + `i`, H
		case r < 91:
			return seed + `o`, H
		case r < 95:
			return seed + `u`, H
		case r < 97:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `tr`:
		H := 2.3991891290867966
		r := rng.IntN(267)
		switch {
		case r < 75:
			return seed + `a`, H
		case r < 147:
			return seed + `i`, H
		case r < 186:
			return seed + `o`, H
		case r < 221:
			return seed + `e`, H
		case r < 254:
			return seed + `u`, H
		case r < 267:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `dy`:
		H := 2.446439344671015
		r := rng.IntN(10)
		switch {
		case r < 3:
			return seed + `n`, H
		case r < 5:
			return seed + `s`, H
		case r < 7:
			return seed + `i`, H
		case r < 8:
			return seed + `l`, H
		case r < 9:
			return seed + `m`, H
		case r < 10:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `xt`:
		H := 2.077235445408308
		r := rng.IntN(20)
		switch {
		case r < 9:
			return seed + `e`, H
		case r < 14:
			return seed + `r`, H
		case r < 17:
			return seed + `i`, H
		case r < 18:
			return seed + `h`, H
		case r < 19:
			return seed + `y`, H
		case r < 20:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ev`:
		H := 1.9942008615410574
		r := rng.IntN(106)
		switch {
		case r < 42:
			return seed + `e`, H
		case r < 72:
			return seed + `i`, H
		case r < 89:
			return seed + `a`, H
		case r < 104:
			return seed + `o`, H
		case r < 105:
			return seed + `y`, H
		case r < 106:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ry`:
		H := 2.5220552088742005
		r := rng.IntN(12)
		switch {
		case r < 4:
			return seed + `i`, H
		case r < 7:
			return seed + `p`, H
		case r < 8:
			return seed + `w`, H
		case r < 9:
			return seed + `s`, H
		case r < 10:
			return seed + `d`, H
		case r < 11:
			return seed + `o`, H
		case r < 12:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `nf`:
		H := 2.585460534063065
		r := rng.IntN(51)
		switch {
		case r < 17:
			return seed + `i`, H
		case r < 27:
			return seed + `o`, H
		case r < 33:
			return seed + `u`, H
		case r < 38:
			return seed + `a`, H
		case r < 43:
			return seed + `r`, H
		case r < 47:
			return seed + `e`, H
		case r < 51:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `oi`:
		H := 2.2369637280222663
		r := rng.IntN(66)
		switch {
		case r < 24:
			return seed + `n`, H
		case r < 44:
			return seed + `l`, H
		case r < 51:
			return seed + `d`, H
		case r < 58:
			return seed + `c`, H
		case r < 64:
			return seed + `s`, H
		case r < 65:
			return seed + `r`, H
		case r < 66:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `nb`:
		H := 2.759079570624175
		r := rng.IntN(25)
		switch {
		case r < 5:
			return seed + `e`, H
		case r < 9:
			return seed + `r`, H
		case r < 13:
			return seed + `o`, H
		case r < 17:
			return seed + `u`, H
		case r < 20:
			return seed + `a`, H
		case r < 23:
			return seed + `l`, H
		case r < 25:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `xp`:
		H := 2.5096856969120642
		r := rng.IntN(30)
		switch {
		case r < 8:
			return seed + `e`, H
		case r < 15:
			return seed + `l`, H
		case r < 21:
			return seed + `o`, H
		case r < 25:
			return seed + `a`, H
		case r < 28:
			return seed + `i`, H
		case r < 29:
			return seed + `r`, H
		case r < 30:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ty`:
		H := 2.3261207468426806
		r := rng.IntN(20)
		switch {
		case r < 7:
			return seed + `l`, H
		case r < 13:
			return seed + `p`, H
		case r < 16:
			return seed + `i`, H
		case r < 17:
			return seed + `f`, H
		case r < 18:
			return seed + `h`, H
		case r < 19:
			return seed + `c`, H
		case r < 20:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ks`:
		H := 2.460904932167797
		r := rng.IntN(17)
		switch {
		case r < 7:
			return seed + `t`, H
		case r < 9:
			return seed + `h`, H
		case r < 11:
			return seed + `a`, H
		case r < 13:
			return seed + `l`, H
		case r < 15:
			return seed + `p`, H
		case r < 16:
			return seed + `i`, H
		case r < 17:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `eu`:
		H := 2.501609497059027
		r := rng.IntN(20)
		switch {
		case r < 7:
			return seed + `r`, H
		case r < 11:
			return seed + `m`, H
		case r < 14:
			return seed + `s`, H
		case r < 16:
			return seed + `n`, H
		case r < 18:
			return seed + `t`, H
		case r < 19:
			return seed + `c`, H
		case r < 20:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `fu`:
		H := 1.5921871522135118
		r := rng.IntN(61)
		switch {
		case r < 40:
			return seed + `l`, H
		case r < 52:
			return seed + `s`, H
		case r < 55:
			return seed + `r`, H
		case r < 57:
			return seed + `n`, H
		case r < 59:
			return seed + `t`, H
		case r < 60:
			return seed + `e`, H
		case r < 61:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `zi`:
		H := 1.8308122292922697
		r := rng.IntN(33)
		switch {
		case r < 20:
			return seed + `n`, H
		case r < 25:
			return seed + `p`, H
		case r < 29:
			return seed + `l`, H
		case r < 30:
			return seed + `g`, H
		case r < 31:
			return seed + `c`, H
		case r < 32:
			return seed + `e`, H
		case r < 33:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `rk`:
		H := 2.4863286677871854
		r := rng.IntN(17)
		switch {
		case r < 5:
			return seed + `i`, H
		case r < 10:
			return seed + `e`, H
		case r < 12:
			return seed + `n`, H
		case r < 14:
			return seed + `y`, H
		case r < 15:
			return seed + `w`, H
		case r < 16:
			return seed + `r`, H
		case r < 17:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `xc`:
		H := 2.6659573209491745
		r := rng.IntN(20)
		switch {
		case r < 5:
			return seed + `l`, H
		case r < 9:
			return seed + `u`, H
		case r < 12:
			return seed + `a`, H
		case r < 15:
			return seed + `e`, H
		case r < 17:
			return seed + `i`, H
		case r < 19:
			return seed + `r`, H
		case r < 20:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `dd`:
		H := 2.279251222260265
		r := rng.IntN(48)
		switch {
		case r < 18:
			return seed + `l`, H
		case r < 31:
			return seed + `e`, H
		case r < 39:
			return seed + `i`, H
		case r < 42:
			return seed + `y`, H
		case r < 45:
			return seed + `a`, H
		case r < 47:
			return seed + `h`, H
		case r < 48:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ye`:
		H := 2.5365766899407323
		r := rng.IntN(24)
		switch {
		case r < 7:
			return seed + `r`, H
		case r < 13:
			return seed + `a`, H
		case r < 17:
			return seed + `d`, H
		case r < 19:
			return seed + `s`, H
		case r < 21:
			return seed + `e`, H
		case r < 23:
			return seed + `l`, H
		case r < 24:
			return seed + `n`, H
		default:
			panic("unexpected rand num")
		}
	case `pp`:
		H := 2.4303217508328046
		r := rng.IntN(114)
		switch {
		case r < 43:
			return seed + `e`, H
		case r < 69:
			return seed + `i`, H
		case r < 84:
			return seed + `l`, H
		case r < 94:
			return seed + `y`, H
		case r < 103:
			return seed + `r`, H
		case r < 111:
			return seed + `o`, H
		case r < 113:
			return seed + `a`, H
		case r < 114:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ka`:
		H := 2.7744019192887706
		r := rng.IntN(18)
		switch {
		case r < 5:
			return seed + `b`, H
		case r < 8:
			return seed + `l`, H
		case r < 11:
			return seed + `r`, H
		case r < 13:
			return seed + `g`, H
		case r < 15:
			return seed + `t`, H
		case r < 16:
			return seed + `n`, H
		case r < 17:
			return seed + `c`, H
		case r < 18:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `bs`:
		H := 2.615277848886663
		r := rng.IntN(35)
		switch {
		case r < 10:
			return seed + `e`, H
		case r < 17:
			return seed + `t`, H
		case r < 23:
			return seed + `o`, H
		case r < 28:
			return seed + `i`, H
		case r < 32:
			return seed + `c`, H
		case r < 33:
			return seed + `u`, H
		case r < 34:
			return seed + `l`, H
		case r < 35:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `if`:
		H := 2.5951312665142345
		r := rng.IntN(106)
		switch {
		case r < 27:
			return seed + `t`, H
		case r < 51:
			return seed + `y`, H
		case r < 75:
			return seed + `i`, H
		case r < 88:
			return seed + `f`, H
		case r < 96:
			return seed + `e`, H
		case r < 100:
			return seed + `l`, H
		case r < 103:
			return seed + `u`, H
		case r < 106:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ze`:
		H := 2.471058306628392
		r := rng.IntN(27)
		switch {
		case r < 11:
			return seed + `r`, H
		case r < 16:
			return seed + `d`, H
		case r < 20:
			return seed + `n`, H
		case r < 22:
			return seed + `l`, H
		case r < 24:
			return seed + `s`, H
		case r < 25:
			return seed + `a`, H
		case r < 26:
			return seed + `p`, H
		case r < 27:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ud`:
		H := 2.436743733204307
		r := rng.IntN(67)
		switch {
		case r < 22:
			return seed + `e`, H
		case r < 38:
			return seed + `i`, H
		case r < 50:
			return seed + `d`, H
		case r < 59:
			return seed + `g`, H
		case r < 62:
			return seed + `a`, H
		case r < 65:
			return seed + `o`, H
		case r < 66:
			return seed + `y`, H
		case r < 67:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `d`:
		H := 2.414563621434015
		r := rng.IntN(547)
		switch {
		case r < 200:
			return seed + `e`, H
		case r < 338:
			return seed + `i`, H
		case r < 401:
			return seed + `r`, H
		case r < 451:
			return seed + `a`, H
		case r < 500:
			return seed + `o`, H
		case r < 535:
			return seed + `u`, H
		case r < 542:
			return seed + `w`, H
		case r < 547:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ei`:
		H := 2.8559049224696933
		r := rng.IntN(21)
		switch {
		case r < 4:
			return seed + `n`, H
		case r < 8:
			return seed + `t`, H
		case r < 11:
			return seed + `v`, H
		case r < 14:
			return seed + `s`, H
		case r < 17:
			return seed + `g`, H
		case r < 19:
			return seed + `l`, H
		case r < 20:
			return seed + `m`, H
		case r < 21:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ey`:
		H := 2.94770277922009
		r := rng.IntN(9)
		switch {
		case r < 2:
			return seed + `a`, H
		case r < 3:
			return seed + `w`, H
		case r < 4:
			return seed + `h`, H
		case r < 5:
			return seed + `l`, H
		case r < 6:
			return seed + `e`, H
		case r < 7:
			return seed + `o`, H
		case r < 8:
			return seed + `m`, H
		case r < 9:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `rg`:
		H := 2.6225466508345625
		r := rng.IntN(59)
		switch {
		case r < 23:
			return seed + `e`, H
		case r < 32:
			return seed + `i`, H
		case r < 38:
			return seed + `l`, H
		case r < 44:
			return seed + `o`, H
		case r < 48:
			return seed + `y`, H
		case r < 52:
			return seed + `r`, H
		case r < 56:
			return seed + `a`, H
		case r < 59:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `t`:
		H := 2.467515412367413
		r := rng.IntN(323)
		switch {
		case r < 105:
			return seed + `r`, H
		case r < 169:
			return seed + `a`, H
		case r < 230:
			return seed + `h`, H
		case r < 271:
			return seed + `i`, H
		case r < 296:
			return seed + `u`, H
		case r < 319:
			return seed + `w`, H
		case r < 322:
			return seed + `y`, H
		case r < 323:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `bt`:
		H := 2.8073549220576046
		r := rng.IntN(14)
		switch {
		case r < 4:
			return seed + `l`, H
		case r < 6:
			return seed + `a`, H
		case r < 8:
			return seed + `r`, H
		case r < 10:
			return seed + `o`, H
		case r < 11:
			return seed + `u`, H
		case r < 12:
			return seed + `e`, H
		case r < 13:
			return seed + `i`, H
		case r < 14:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `yo`:
		H := 2.692380602454975
		r := rng.IntN(14)
		switch {
		case r < 5:
			return seed + `n`, H
		case r < 7:
			return seed + `y`, H
		case r < 9:
			return seed + `g`, H
		case r < 10:
			return seed + `f`, H
		case r < 11:
			return seed + `v`, H
		case r < 12:
			return seed + `u`, H
		case r < 13:
			return seed + `r`, H
		case r < 14:
			return seed + `d`, H
		default:
			panic("unexpected rand num")
		}
	case `u`:
		H := 0.8648744526550217
		r := rng.IntN(476)
		switch {
		case r < 410:
			return seed + `n`, H
		case r < 446:
			return seed + `p`, H
		case r < 455:
			return seed + `r`, H
		case r < 463:
			return seed + `s`, H
		case r < 469:
			return seed + `t`, H
		case r < 472:
			return seed + `l`, H
		case r < 475:
			return seed + `m`, H
		case r < 476:
			return seed + `d`, H
		default:
			panic("unexpected rand num")
		}
	case `ps`:
		H := 2.489655670358648
		r := rng.IntN(27)
		switch {
		case r < 8:
			return seed + `t`, H
		case r < 15:
			return seed + `e`, H
		case r < 21:
			return seed + `i`, H
		case r < 23:
			return seed + `y`, H
		case r < 24:
			return seed + `w`, H
		case r < 25:
			return seed + `u`, H
		case r < 26:
			return seed + `a`, H
		case r < 27:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `pt`:
		H := 2.3129328787410666
		r := rng.IntN(45)
		switch {
		case r < 21:
			return seed + `i`, H
		case r < 28:
			return seed + `u`, H
		case r < 35:
			return seed + `o`, H
		case r < 38:
			return seed + `l`, H
		case r < 41:
			return seed + `e`, H
		case r < 43:
			return seed + `a`, H
		case r < 44:
			return seed + `h`, H
		case r < 45:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `af`:
		H := 2.1819962604162924
		r := rng.IntN(59)
		switch {
		case r < 25:
			return seed + `f`, H
		case r < 42:
			return seed + `t`, H
		case r < 50:
			return seed + `e`, H
		case r < 53:
			return seed + `l`, H
		case r < 55:
			return seed + `a`, H
		case r < 57:
			return seed + `r`, H
		case r < 58:
			return seed + `n`, H
		case r < 59:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `rf`:
		H := 2.5833333333333335
		r := rng.IntN(24)
		switch {
		case r < 9:
			return seed + `e`, H
		case r < 13:
			return seed + `u`, H
		case r < 16:
			return seed + `a`, H
		case r < 19:
			return seed + `i`, H
		case r < 21:
			return seed + `l`, H
		case r < 22:
			return seed + `r`, H
		case r < 23:
			return seed + `o`, H
		case r < 24:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `k`:
		H := 2.2102710078071865
		r := rng.IntN(55)
		switch {
		case r < 27:
			return seed + `i`, H
		case r < 38:
			return seed + `e`, H
		case r < 43:
			return seed + `n`, H
		case r < 48:
			return seed + `a`, H
		case r < 51:
			return seed + `o`, H
		case r < 53:
			return seed + `u`, H
		case r < 54:
			return seed + `l`, H
		case r < 55:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `mb`:
		H := 2.4620228668017208
		r := rng.IntN(75)
		switch {
		case r < 30:
			return seed + `l`, H
		case r < 44:
			return seed + `e`, H
		case r < 52:
			return seed + `a`, H
		case r < 60:
			return seed + `i`, H
		case r < 68:
			return seed + `o`, H
		case r < 72:
			return seed + `u`, H
		case r < 74:
			return seed + `r`, H
		case r < 75:
			return seed + `n`, H
		default:
			panic("unexpected rand num")
		}
	case `ph`:
		H := 2.525379371610213
		r := rng.IntN(56)
		switch {
		case r < 17:
			return seed + `o`, H
		case r < 29:
			return seed + `a`, H
		case r < 41:
			return seed + `e`, H
		case r < 47:
			return seed + `i`, H
		case r < 51:
			return seed + `r`, H
		case r < 54:
			return seed + `y`, H
		case r < 55:
			return seed + `l`, H
		case r < 56:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `sc`:
		H := 2.5124428937801095
		r := rng.IntN(134)
		switch {
		case r < 40:
			return seed + `a`, H
		case r < 73:
			return seed + `o`, H
		case r < 99:
			return seed + `r`, H
		case r < 112:
			return seed + `u`, H
		case r < 121:
			return seed + `h`, H
		case r < 129:
			return seed + `e`, H
		case r < 133:
			return seed + `i`, H
		case r < 134:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `vo`:
		H := 2.871054088933562
		r := rng.IntN(57)
		switch {
		case r < 16:
			return seed + `r`, H
		case r < 27:
			return seed + `l`, H
		case r < 34:
			return seed + `c`, H
		case r < 40:
			return seed + `u`, H
		case r < 45:
			return seed + `t`, H
		case r < 49:
			return seed + `i`, H
		case r < 53:
			return seed + `k`, H
		case r < 55:
			return seed + `w`, H
		case r < 57:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `c`:
		H := 2.671815163549163
		r := rng.IntN(751)
		switch {
		case r < 224:
			return seed + `o`, H
		case r < 389:
			return seed + `a`, H
		case r < 496:
			return seed + `r`, H
		case r < 593:
			return seed + `h`, H
		case r < 656:
			return seed + `l`, H
		case r < 708:
			return seed + `u`, H
		case r < 727:
			return seed + `i`, H
		case r < 743:
			return seed + `e`, H
		case r < 751:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `g`:
		H := 2.6968078160957982
		r := rng.IntN(304)
		switch {
		case r < 97:
			return seed + `r`, H
		case r < 142:
			return seed + `a`, H
		case r < 185:
			return seed + `l`, H
		case r < 218:
			return seed + `e`, H
		case r < 250:
			return seed + `o`, H
		case r < 277:
			return seed + `u`, H
		case r < 301:
			return seed + `i`, H
		case r < 303:
			return seed + `n`, H
		case r < 304:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `sk`:
		H := 2.0077378472545835
		r := rng.IntN(61)
		switch {
		case r < 32:
			return seed + `i`, H
		case r < 46:
			return seed + `e`, H
		case r < 54:
			return seed + `y`, H
		case r < 56:
			return seed + `a`, H
		case r < 57:
			return seed + `n`, H
		case r < 58:
			return seed + `w`, H
		case r < 59:
			return seed + `l`, H
		case r < 60:
			return seed + `t`, H
		case r < 61:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `lp`:
		H := 3.0503018554349826
		r := rng.IntN(19)
		switch {
		case r < 4:
			return seed + `h`, H
		case r < 7:
			return seed + `i`, H
		case r < 9:
			return seed + `f`, H
		case r < 11:
			return seed + `a`, H
		case r < 13:
			return seed + `e`, H
		case r < 15:
			return seed + `l`, H
		case r < 17:
			return seed + `t`, H
		case r < 18:
			return seed + `r`, H
		case r < 19:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ya`:
		H := 2.7153266987718103
		r := rng.IntN(19)
		switch {
		case r < 6:
			return seed + `r`, H
		case r < 11:
			return seed + `b`, H
		case r < 13:
			return seed + `l`, H
		case r < 14:
			return seed + `n`, H
		case r < 15:
			return seed + `w`, H
		case r < 16:
			return seed + `g`, H
		case r < 17:
			return seed + `h`, H
		case r < 18:
			return seed + `p`, H
		case r < 19:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ny`:
		H := 3.039148671903072
		r := rng.IntN(14)
		switch {
		case r < 3:
			return seed + `m`, H
		case r < 5:
			return seed + `w`, H
		case r < 7:
			return seed + `t`, H
		case r < 9:
			return seed + `o`, H
		case r < 10:
			return seed + `h`, H
		case r < 11:
			return seed + `x`, H
		case r < 12:
			return seed + `l`, H
		case r < 13:
			return seed + `p`, H
		case r < 14:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `yp`:
		H := 2.7132696895151076
		r := rng.IntN(25)
		switch {
		case r < 8:
			return seed + `e`, H
		case r < 14:
			return seed + `n`, H
		case r < 17:
			return seed + `t`, H
		case r < 19:
			return seed + `h`, H
		case r < 21:
			return seed + `o`, H
		case r < 22:
			return seed + `a`, H
		case r < 23:
			return seed + `l`, H
		case r < 24:
			return seed + `i`, H
		case r < 25:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ht`:
		H := 2.6292492238560348
		r := rng.IntN(19)
		switch {
		case r < 8:
			return seed + `e`, H
		case r < 11:
			return seed + `l`, H
		case r < 13:
			return seed + `i`, H
		case r < 14:
			return seed + `n`, H
		case r < 15:
			return seed + `f`, H
		case r < 16:
			return seed + `w`, H
		case r < 17:
			return seed + `y`, H
		case r < 18:
			return seed + `r`, H
		case r < 19:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rc`:
		H := 2.8029839504926644
		r := rng.IntN(53)
		switch {
		case r < 15:
			return seed + `h`, H
		case r < 26:
			return seed + `e`, H
		case r < 34:
			return seed + `o`, H
		case r < 39:
			return seed + `u`, H
		case r < 43:
			return seed + `a`, H
		case r < 47:
			return seed + `i`, H
		case r < 50:
			return seed + `l`, H
		case r < 52:
			return seed + `t`, H
		case r < 53:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ai`:
		H := 2.315377442431913
		r := rng.IntN(132)
		switch {
		case r < 56:
			return seed + `n`, H
		case r < 88:
			return seed + `l`, H
		case r < 102:
			return seed + `r`, H
		case r < 115:
			return seed + `d`, H
		case r < 123:
			return seed + `m`, H
		case r < 128:
			return seed + `s`, H
		case r < 130:
			return seed + `t`, H
		case r < 131:
			return seed + `k`, H
		case r < 132:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `rb`:
		H := 2.7675793883891737
		r := rng.IntN(47)
		switch {
		case r < 12:
			return seed + `o`, H
		case r < 23:
			return seed + `i`, H
		case r < 30:
			return seed + `a`, H
		case r < 36:
			return seed + `e`, H
		case r < 40:
			return seed + `l`, H
		case r < 43:
			return seed + `y`, H
		case r < 45:
			return seed + `u`, H
		case r < 46:
			return seed + `r`, H
		case r < 47:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ft`:
		H := 2.4952508233936745
		r := rng.IntN(33)
		switch {
		case r < 13:
			return seed + `e`, H
		case r < 20:
			return seed + `i`, H
		case r < 25:
			return seed + `y`, H
		case r < 28:
			return seed + `l`, H
		case r < 29:
			return seed + `n`, H
		case r < 30:
			return seed + `w`, H
		case r < 31:
			return seed + `h`, H
		case r < 32:
			return seed + `s`, H
		case r < 33:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ff`:
		H := 2.723780466448913
		r := rng.IntN(63)
		switch {
		case r < 16:
			return seed + `e`, H
		case r < 30:
			return seed + `i`, H
		case r < 44:
			return seed + `l`, H
		case r < 49:
			return seed + `o`, H
		case r < 53:
			return seed + `u`, H
		case r < 57:
			return seed + `y`, H
		case r < 60:
			return seed + `r`, H
		case r < 62:
			return seed + `a`, H
		case r < 63:
			return seed + `n`, H
		default:
			panic("unexpected rand num")
		}
	case `p`:
		H := 2.5923931878457345
		r := rng.IntN(578)
		switch {
		case r < 166:
			return seed + `r`, H
		case r < 302:
			return seed + `a`, H
		case r < 379:
			return seed + `o`, H
		case r < 447:
			return seed + `e`, H
		case r < 506:
			return seed + `u`, H
		case r < 562:
			return seed + `l`, H
		case r < 574:
			return seed + `h`, H
		case r < 577:
			return seed + `y`, H
		case r < 578:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `eb`:
		H := 2.887941726415332
		r := rng.IntN(47)
		switch {
		case r < 10:
			return seed + `a`, H
		case r < 20:
			return seed + `o`, H
		case r < 29:
			return seed + `u`, H
		case r < 34:
			return seed + `r`, H
		case r < 38:
			return seed + `l`, H
		case r < 41:
			return seed + `i`, H
		case r < 43:
			return seed + `t`, H
		case r < 45:
			return seed + `e`, H
		case r < 47:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `tt`:
		H := 2.3319551423402523
		r := rng.IntN(127)
		switch {
		case r < 59:
			return seed + `e`, H
		case r < 80:
			return seed + `i`, H
		case r < 96:
			return seed + `l`, H
		case r < 106:
			return seed + `a`, H
		case r < 116:
			return seed + `o`, H
		case r < 123:
			return seed + `y`, H
		case r < 125:
			return seed + `r`, H
		case r < 126:
			return seed + `h`, H
		case r < 127:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rp`:
		H := 2.9031815050305063
		r := rng.IntN(41)
		switch {
		case r < 8:
			return seed + `l`, H
		case r < 16:
			return seed + `o`, H
		case r < 22:
			return seed + `e`, H
		case r < 28:
			return seed + `a`, H
		case r < 33:
			return seed + `i`, H
		case r < 37:
			return seed + `h`, H
		case r < 39:
			return seed + `r`, H
		case r < 40:
			return seed + `n`, H
		case r < 41:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `od`:
		H := 2.7181256737103596
		r := rng.IntN(50)
		switch {
		case r < 13:
			return seed + `e`, H
		case r < 25:
			return seed + `i`, H
		case r < 33:
			return seed + `y`, H
		case r < 40:
			return seed + `u`, H
		case r < 43:
			return seed + `g`, H
		case r < 46:
			return seed + `o`, H
		case r < 48:
			return seed + `l`, H
		case r < 49:
			return seed + `a`, H
		case r < 50:
			return seed + `d`, H
		default:
			panic("unexpected rand num")
		}
	case `io`:
		H := 1.3890034152520174
		r := rng.IntN(187)
		switch {
		case r < 140:
			return seed + `n`, H
		case r < 163:
			return seed + `u`, H
		case r < 169:
			return seed + `r`, H
		case r < 175:
			return seed + `l`, H
		case r < 180:
			return seed + `t`, H
		case r < 183:
			return seed + `c`, H
		case r < 185:
			return seed + `d`, H
		case r < 186:
			return seed + `x`, H
		case r < 187:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ib`:
		H := 2.5862770862160236
		r := rng.IntN(51)
		switch {
		case r < 19:
			return seed + `l`, H
		case r < 30:
			return seed + `e`, H
		case r < 36:
			return seed + `u`, H
		case r < 41:
			return seed + `b`, H
		case r < 44:
			return seed + `i`, H
		case r < 47:
			return seed + `r`, H
		case r < 49:
			return seed + `a`, H
		case r < 50:
			return seed + `o`, H
		case r < 51:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `sp`:
		H := 2.8235716352955578
		r := rng.IntN(160)
		switch {
		case r < 38:
			return seed + `o`, H
		case r < 73:
			return seed + `e`, H
		case r < 102:
			return seed + `i`, H
		case r < 119:
			return seed + `l`, H
		case r < 136:
			return seed + `r`, H
		case r < 145:
			return seed + `a`, H
		case r < 152:
			return seed + `h`, H
		case r < 156:
			return seed + `u`, H
		case r < 159:
			return seed + `y`, H
		case r < 160:
			return seed + `n`, H
		default:
			panic("unexpected rand num")
		}
	case `oy`:
		H := 2.9638096525384383
		r := rng.IntN(23)
		switch {
		case r < 6:
			return seed + `e`, H
		case r < 11:
			return seed + `a`, H
		case r < 14:
			return seed + `o`, H
		case r < 16:
			return seed + `i`, H
		case r < 18:
			return seed + `s`, H
		case r < 19:
			return seed + `n`, H
		case r < 20:
			return seed + `f`, H
		case r < 21:
			return seed + `l`, H
		case r < 22:
			return seed + `m`, H
		case r < 23:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ts`:
		H := 3.135821824148036
		r := rng.IntN(22)
		switch {
		case r < 5:
			return seed + `h`, H
		case r < 8:
			return seed + `i`, H
		case r < 11:
			return seed + `m`, H
		case r < 13:
			return seed + `y`, H
		case r < 15:
			return seed + `u`, H
		case r < 17:
			return seed + `k`, H
		case r < 19:
			return seed + `o`, H
		case r < 20:
			return seed + `c`, H
		case r < 21:
			return seed + `e`, H
		case r < 22:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ct`:
		H := 2.5852558586975434
		r := rng.IntN(118)
		switch {
		case r < 44:
			return seed + `i`, H
		case r < 67:
			return seed + `o`, H
		case r < 83:
			return seed + `e`, H
		case r < 96:
			return seed + `a`, H
		case r < 106:
			return seed + `u`, H
		case r < 110:
			return seed + `l`, H
		case r < 113:
			return seed + `s`, H
		case r < 116:
			return seed + `r`, H
		case r < 117:
			return seed + `f`, H
		case r < 118:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `we`:
		H := 2.748621296142989
		r := rng.IntN(58)
		switch {
		case r < 18:
			return seed + `r`, H
		case r < 27:
			return seed + `l`, H
		case r < 35:
			return seed + `a`, H
		case r < 43:
			return seed + `e`, H
		case r < 51:
			return seed + `d`, H
		case r < 54:
			return seed + `n`, H
		case r < 55:
			return seed + `i`, H
		case r < 56:
			return seed + `s`, H
		case r < 57:
			return seed + `b`, H
		case r < 58:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ds`:
		H := 3.0464393446710147
		r := rng.IntN(20)
		switch {
		case r < 6:
			return seed + `t`, H
		case r < 8:
			return seed + `h`, H
		case r < 10:
			return seed + `i`, H
		case r < 12:
			return seed + `e`, H
		case r < 14:
			return seed + `m`, H
		case r < 16:
			return seed + `c`, H
		case r < 17:
			return seed + `f`, H
		case r < 18:
			return seed + `a`, H
		case r < 19:
			return seed + `o`, H
		case r < 20:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `oe`:
		H := 3.2516291673878226
		r := rng.IntN(12)
		switch {
		case r < 2:
			return seed + `s`, H
		case r < 4:
			return seed + `r`, H
		case r < 5:
			return seed + `n`, H
		case r < 6:
			return seed + `x`, H
		case r < 7:
			return seed + `d`, H
		case r < 8:
			return seed + `y`, H
		case r < 9:
			return seed + `t`, H
		case r < 10:
			return seed + `m`, H
		case r < 11:
			return seed + `i`, H
		case r < 12:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `nc`:
		H := 2.8557990432188514
		r := rng.IntN(207)
		switch {
		case r < 66:
			return seed + `e`, H
		case r < 109:
			return seed + `h`, H
		case r < 130:
			return seed + `i`, H
		case r < 149:
			return seed + `o`, H
		case r < 165:
			return seed + `y`, H
		case r < 177:
			return seed + `l`, H
		case r < 188:
			return seed + `t`, H
		case r < 196:
			return seed + `r`, H
		case r < 202:
			return seed + `a`, H
		case r < 207:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ia`:
		H := 2.7871615786315345
		r := rng.IntN(106)
		switch {
		case r < 26:
			return seed + `l`, H
		case r < 48:
			return seed + `n`, H
		case r < 66:
			return seed + `t`, H
		case r < 84:
			return seed + `b`, H
		case r < 93:
			return seed + `r`, H
		case r < 97:
			return seed + `g`, H
		case r < 100:
			return seed + `c`, H
		case r < 102:
			return seed + `s`, H
		case r < 104:
			return seed + `m`, H
		case r < 106:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `lt`:
		H := 2.9388580483450712
		r := rng.IntN(36)
		switch {
		case r < 11:
			return seed + `i`, H
		case r < 17:
			return seed + `e`, H
		case r < 21:
			return seed + `y`, H
		case r < 25:
			return seed + `r`, H
		case r < 28:
			return seed + `h`, H
		case r < 30:
			return seed + `u`, H
		case r < 32:
			return seed + `a`, H
		case r < 34:
			return seed + `z`, H
		case r < 35:
			return seed + `o`, H
		case r < 36:
			return seed + `t`, H
		default:
			panic("unexpected rand num")
		}
	case `ef`:
		H := 3.1182289091344035
		r := rng.IntN(91)
		switch {
		case r < 17:
			return seed + `u`, H
		case r < 33:
			return seed + `i`, H
		case r < 45:
			return seed + `e`, H
		case r < 56:
			return seed + `o`, H
		case r < 66:
			return seed + `l`, H
		case r < 73:
			return seed + `r`, H
		case r < 79:
			return seed + `a`, H
		case r < 85:
			return seed + `t`, H
		case r < 89:
			return seed + `f`, H
		case r < 91:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `ou`:
		H := 2.433480709472623
		r := rng.IntN(258)
		switch {
		case r < 73:
			return seed + `s`, H
		case r < 146:
			return seed + `t`, H
		case r < 211:
			return seed + `n`, H
		case r < 227:
			return seed + `r`, H
		case r < 237:
			return seed + `c`, H
		case r < 244:
			return seed + `g`, H
		case r < 250:
			return seed + `p`, H
		case r < 253:
			return seed + `d`, H
		case r < 256:
			return seed + `b`, H
		case r < 258:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `nk`:
		H := 2.955870327431456
		r := rng.IntN(61)
		switch {
		case r < 17:
			return seed + `i`, H
		case r < 30:
			return seed + `e`, H
		case r < 38:
			return seed + `l`, H
		case r < 44:
			return seed + `y`, H
		case r < 49:
			return seed + `n`, H
		case r < 52:
			return seed + `s`, H
		case r < 54:
			return seed + `h`, H
		case r < 56:
			return seed + `a`, H
		case r < 58:
			return seed + `m`, H
		case r < 60:
			return seed + `b`, H
		case r < 61:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ly`:
		H := 3.272769892034794
		r := rng.IntN(26)
		switch {
		case r < 5:
			return seed + `i`, H
		case r < 9:
			return seed + `r`, H
		case r < 12:
			return seed + `s`, H
		case r < 15:
			return seed + `z`, H
		case r < 17:
			return seed + `e`, H
		case r < 19:
			return seed + `a`, H
		case r < 21:
			return seed + `m`, H
		case r < 23:
			return seed + `g`, H
		case r < 24:
			return seed + `w`, H
		case r < 25:
			return seed + `o`, H
		case r < 26:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `um`:
		H := 2.7569798111220467
		r := rng.IntN(99)
		switch {
		case r < 29:
			return seed + `b`, H
		case r < 49:
			return seed + `p`, H
		case r < 64:
			return seed + `m`, H
		case r < 78:
			return seed + `e`, H
		case r < 87:
			return seed + `i`, H
		case r < 92:
			return seed + `o`, H
		case r < 94:
			return seed + `s`, H
		case r < 96:
			return seed + `a`, H
		case r < 97:
			return seed + `n`, H
		case r < 98:
			return seed + `d`, H
		case r < 99:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `eo`:
		H := 3.1748858687242363
		r := rng.IntN(23)
		switch {
		case r < 6:
			return seed + `l`, H
		case r < 9:
			return seed + `r`, H
		case r < 12:
			return seed + `c`, H
		case r < 14:
			return seed + `n`, H
		case r < 16:
			return seed + `d`, H
		case r < 18:
			return seed + `m`, H
		case r < 19:
			return seed + `t`, H
		case r < 20:
			return seed + `p`, H
		case r < 21:
			return seed + `g`, H
		case r < 22:
			return seed + `v`, H
		case r < 23:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `gg`:
		H := 2.5365464726603135
		r := rng.IntN(59)
		switch {
		case r < 19:
			return seed + `l`, H
		case r < 34:
			return seed + `e`, H
		case r < 47:
			return seed + `i`, H
		case r < 51:
			return seed + `y`, H
		case r < 53:
			return seed + `a`, H
		case r < 54:
			return seed + `n`, H
		case r < 55:
			return seed + `o`, H
		case r < 56:
			return seed + `r`, H
		case r < 57:
			return seed + `s`, H
		case r < 58:
			return seed + `p`, H
		case r < 59:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `be`:
		H := 2.6259431593969045
		r := rng.IntN(69)
		switch {
		case r < 27:
			return seed + `r`, H
		case r < 43:
			return seed + `d`, H
		case r < 50:
			return seed + `l`, H
		case r < 56:
			return seed + `a`, H
		case r < 59:
			return seed + `t`, H
		case r < 61:
			return seed + `n`, H
		case r < 63:
			return seed + `e`, H
		case r < 65:
			return seed + `s`, H
		case r < 67:
			return seed + `c`, H
		case r < 68:
			return seed + `y`, H
		case r < 69:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `jo`:
		H := 2.9275496777656036
		r := rng.IntN(32)
		switch {
		case r < 11:
			return seed + `y`, H
		case r < 17:
			return seed + `i`, H
		case r < 20:
			return seed + `l`, H
		case r < 22:
			return seed + `r`, H
		case r < 24:
			return seed + `k`, H
		case r < 26:
			return seed + `c`, H
		case r < 28:
			return seed + `g`, H
		case r < 29:
			return seed + `h`, H
		case r < 30:
			return seed + `t`, H
		case r < 31:
			return seed + `v`, H
		case r < 32:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ug`:
		H := 2.8739810481273578
		r := rng.IntN(34)
		switch {
		case r < 10:
			return seed + `g`, H
		case r < 19:
			return seed + `h`, H
		case r < 22:
			return seed + `n`, H
		case r < 25:
			return seed + `a`, H
		case r < 27:
			return seed + `u`, H
		case r < 29:
			return seed + `l`, H
		case r < 30:
			return seed + `e`, H
		case r < 31:
			return seed + `o`, H
		case r < 32:
			return seed + `s`, H
		case r < 33:
			return seed + `i`, H
		case r < 34:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ex`:
		H := 2.823204116207561
		r := rng.IntN(102)
		switch {
		case r < 30:
			return seed + `p`, H
		case r < 50:
			return seed + `t`, H
		case r < 69:
			return seed + `c`, H
		case r < 77:
			return seed + `e`, H
		case r < 85:
			return seed + `i`, H
		case r < 91:
			return seed + `a`, H
		case r < 95:
			return seed + `o`, H
		case r < 98:
			return seed + `h`, H
		case r < 100:
			return seed + `u`, H
		case r < 101:
			return seed + `f`, H
		case r < 102:
			return seed + `q`, H
		default:
			panic("unexpected rand num")
		}
	case `va`:
		H := 3.0359223229700754
		r := rng.IntN(108)
		switch {
		case r < 28:
			return seed + `l`, H
		case r < 47:
			return seed + `t`, H
		case r < 59:
			return seed + `r`, H
		case r < 70:
			return seed + `b`, H
		case r < 80:
			return seed + `n`, H
		case r < 89:
			return seed + `g`, H
		case r < 97:
			return seed + `c`, H
		case r < 103:
			return seed + `s`, H
		case r < 105:
			return seed + `i`, H
		case r < 107:
			return seed + `p`, H
		case r < 108:
			return seed + `d`, H
		default:
			panic("unexpected rand num")
		}
	case `ed`:
		H := 2.9112023765018518
		r := rng.IntN(74)
		switch {
		case r < 25:
			return seed + `i`, H
		case r < 35:
			return seed + `u`, H
		case r < 43:
			return seed + `e`, H
		case r < 51:
			return seed + `a`, H
		case r < 59:
			return seed + `g`, H
		case r < 63:
			return seed + `d`, H
		case r < 67:
			return seed + `o`, H
		case r < 70:
			return seed + `y`, H
		case r < 72:
			return seed + `l`, H
		case r < 73:
			return seed + `t`, H
		case r < 74:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ss`:
		H := 2.826274120140731
		r := rng.IntN(71)
		switch {
		case r < 19:
			return seed + `e`, H
		case r < 38:
			return seed + `i`, H
		case r < 47:
			return seed + `u`, H
		case r < 54:
			return seed + `a`, H
		case r < 59:
			return seed + `o`, H
		case r < 63:
			return seed + `y`, H
		case r < 66:
			return seed + `l`, H
		case r < 68:
			return seed + `p`, H
		case r < 69:
			return seed + `f`, H
		case r < 70:
			return seed + `w`, H
		case r < 71:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ue`:
		H := 3.066701577503766
		r := rng.IntN(39)
		switch {
		case r < 9:
			return seed + `l`, H
		case r < 17:
			return seed + `n`, H
		case r < 21:
			return seed + `d`, H
		case r < 25:
			return seed + `e`, H
		case r < 29:
			return seed + `s`, H
		case r < 33:
			return seed + `a`, H
		case r < 35:
			return seed + `r`, H
		case r < 36:
			return seed + `f`, H
		case r < 37:
			return seed + `t`, H
		case r < 38:
			return seed + `u`, H
		case r < 39:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `of`:
		H := 3.1173143543437773
		r := rng.IntN(26)
		switch {
		case r < 7:
			return seed + `f`, H
		case r < 11:
			return seed + `e`, H
		case r < 14:
			return seed + `i`, H
		case r < 17:
			return seed + `a`, H
		case r < 19:
			return seed + `o`, H
		case r < 21:
			return seed + `t`, H
		case r < 22:
			return seed + `s`, H
		case r < 23:
			return seed + `y`, H
		case r < 24:
			return seed + `r`, H
		case r < 25:
			return seed + `u`, H
		case r < 26:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `wo`:
		H := 2.4661778189349435
		r := rng.IntN(67)
		switch {
		case r < 33:
			return seed + `r`, H
		case r < 43:
			return seed + `o`, H
		case r < 51:
			return seed + `m`, H
		case r < 54:
			return seed + `v`, H
		case r < 57:
			return seed + `b`, H
		case r < 59:
			return seed + `w`, H
		case r < 61:
			return seed + `k`, H
		case r < 63:
			return seed + `u`, H
		case r < 65:
			return seed + `l`, H
		case r < 66:
			return seed + `n`, H
		case r < 67:
			return seed + `f`, H
		default:
			panic("unexpected rand num")
		}
	case `me`:
		H := 2.700171980219928
		r := rng.IntN(131)
		switch {
		case r < 38:
			return seed + `n`, H
		case r < 70:
			return seed + `r`, H
		case r < 93:
			return seed + `d`, H
		case r < 104:
			return seed + `t`, H
		case r < 113:
			return seed + `s`, H
		case r < 121:
			return seed + `l`, H
		case r < 126:
			return seed + `a`, H
		case r < 128:
			return seed + `g`, H
		case r < 129:
			return seed + `o`, H
		case r < 130:
			return seed + `m`, H
		case r < 131:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `og`:
		H := 3.2145272645029523
		r := rng.IntN(60)
		switch {
		case r < 10:
			return seed + `y`, H
		case r < 18:
			return seed + `r`, H
		case r < 26:
			return seed + `g`, H
		case r < 33:
			return seed + `e`, H
		case r < 40:
			return seed + `a`, H
		case r < 46:
			return seed + `i`, H
		case r < 52:
			return seed + `u`, H
		case r < 56:
			return seed + `n`, H
		case r < 58:
			return seed + `l`, H
		case r < 59:
			return seed + `w`, H
		case r < 60:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `bi`:
		H := 3.167327740603032
		r := rng.IntN(72)
		switch {
		case r < 15:
			return seed + `l`, H
		case r < 29:
			return seed + `t`, H
		case r < 41:
			return seed + `n`, H
		case r < 48:
			return seed + `d`, H
		case r < 53:
			return seed + `c`, H
		case r < 57:
			return seed + `e`, H
		case r < 61:
			return seed + `r`, H
		case r < 65:
			return seed + `a`, H
		case r < 67:
			return seed + `o`, H
		case r < 69:
			return seed + `s`, H
		case r < 71:
			return seed + `g`, H
		case r < 72:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `em`:
		H := 2.909604688717815
		r := rng.IntN(135)
		switch {
		case r < 28:
			return seed + `i`, H
		case r < 50:
			return seed + `e`, H
		case r < 72:
			return seed + `p`, H
		case r < 92:
			return seed + `o`, H
		case r < 111:
			return seed + `b`, H
		case r < 126:
			return seed + `a`, H
		case r < 128:
			return seed + `y`, H
		case r < 130:
			return seed + `l`, H
		case r < 132:
			return seed + `u`, H
		case r < 133:
			return seed + `n`, H
		case r < 134:
			return seed + `s`, H
		case r < 135:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `au`:
		H := 3.046087347807784
		r := rng.IntN(69)
		switch {
		case r < 15:
			return seed + `t`, H
		case r < 28:
			return seed + `n`, H
		case r < 39:
			return seed + `d`, H
		case r < 48:
			return seed + `s`, H
		case r < 55:
			return seed + `c`, H
		case r < 60:
			return seed + `g`, H
		case r < 63:
			return seed + `r`, H
		case r < 65:
			return seed + `l`, H
		case r < 66:
			return seed + `p`, H
		case r < 67:
			return seed + `z`, H
		case r < 68:
			return seed + `v`, H
		case r < 69:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `aw`:
		H := 2.978305516512272
		r := rng.IntN(35)
		switch {
		case r < 12:
			return seed + `a`, H
		case r < 18:
			return seed + `l`, H
		case r < 22:
			return seed + `n`, H
		case r < 25:
			return seed + `e`, H
		case r < 27:
			return seed + `f`, H
		case r < 29:
			return seed + `k`, H
		case r < 30:
			return seed + `d`, H
		case r < 31:
			return seed + `o`, H
		case r < 32:
			return seed + `h`, H
		case r < 33:
			return seed + `s`, H
		case r < 34:
			return seed + `i`, H
		case r < 35:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `ew`:
		H := 2.8625074707625155
		r := rng.IntN(44)
		switch {
		case r < 16:
			return seed + `a`, H
		case r < 24:
			return seed + `i`, H
		case r < 30:
			return seed + `e`, H
		case r < 33:
			return seed + `o`, H
		case r < 35:
			return seed + `r`, H
		case r < 37:
			return seed + `m`, H
		case r < 39:
			return seed + `l`, H
		case r < 40:
			return seed + `n`, H
		case r < 41:
			return seed + `d`, H
		case r < 42:
			return seed + `h`, H
		case r < 43:
			return seed + `y`, H
		case r < 44:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ec`:
		H := 3.0339701786495312
		r := rng.IntN(189)
		switch {
		case r < 64:
			return seed + `t`, H
		case r < 90:
			return seed + `o`, H
		case r < 106:
			return seed + `e`, H
		case r < 122:
			return seed + `i`, H
		case r < 135:
			return seed + `k`, H
		case r < 148:
			return seed + `a`, H
		case r < 161:
			return seed + `l`, H
		case r < 172:
			return seed + `u`, H
		case r < 179:
			return seed + `r`, H
		case r < 184:
			return seed + `h`, H
		case r < 188:
			return seed + `y`, H
		case r < 189:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `oc`:
		H := 2.8369155250951916
		r := rng.IntN(91)
		switch {
		case r < 37:
			return seed + `k`, H
		case r < 49:
			return seed + `a`, H
		case r < 57:
			return seed + `c`, H
		case r < 64:
			return seed + `r`, H
		case r < 70:
			return seed + `e`, H
		case r < 76:
			return seed + `i`, H
		case r < 81:
			return seed + `t`, H
		case r < 85:
			return seed + `u`, H
		case r < 87:
			return seed + `o`, H
		case r < 89:
			return seed + `h`, H
		case r < 90:
			return seed + `y`, H
		case r < 91:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `mu`:
		H := 3.0506400538703033
		r := rng.IntN(80)
		switch {
		case r < 22:
			return seed + `s`, H
		case r < 36:
			return seed + `l`, H
		case r < 47:
			return seed + `t`, H
		case r < 55:
			return seed + `m`, H
		case r < 61:
			return seed + `r`, H
		case r < 66:
			return seed + `n`, H
		case r < 71:
			return seed + `g`, H
		case r < 74:
			return seed + `d`, H
		case r < 77:
			return seed + `c`, H
		case r < 78:
			return seed + `f`, H
		case r < 79:
			return seed + `p`, H
		case r < 80:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `uc`:
		H := 2.8631704063613026
		r := rng.IntN(87)
		switch {
		case r < 31:
			return seed + `k`, H
		case r < 44:
			return seed + `h`, H
		case r < 57:
			return seed + `t`, H
		case r < 66:
			return seed + `e`, H
		case r < 71:
			return seed + `i`, H
		case r < 75:
			return seed + `a`, H
		case r < 78:
			return seed + `c`, H
		case r < 81:
			return seed + `l`, H
		case r < 83:
			return seed + `o`, H
		case r < 85:
			return seed + `u`, H
		case r < 86:
			return seed + `y`, H
		case r < 87:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `os`:
		H := 2.8544000807530825
		r := rng.IntN(128)
		switch {
		case r < 33:
			return seed + `e`, H
		case r < 66:
			return seed + `t`, H
		case r < 84:
			return seed + `s`, H
		case r < 101:
			return seed + `i`, H
		case r < 108:
			return seed + `a`, H
		case r < 112:
			return seed + `p`, H
		case r < 116:
			return seed + `m`, H
		case r < 120:
			return seed + `u`, H
		case r < 123:
			return seed + `h`, H
		case r < 126:
			return seed + `y`, H
		case r < 127:
			return seed + `o`, H
		case r < 128:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `xi`:
		H := 3.1329440449809582
		r := rng.IntN(26)
		switch {
		case r < 8:
			return seed + `d`, H
		case r < 12:
			return seed + `s`, H
		case r < 15:
			return seed + `m`, H
		case r < 17:
			return seed + `n`, H
		case r < 19:
			return seed + `c`, H
		case r < 20:
			return seed + `f`, H
		case r < 21:
			return seed + `o`, H
		case r < 22:
			return seed + `r`, H
		case r < 23:
			return seed + `t`, H
		case r < 24:
			return seed + `a`, H
		case r < 25:
			return seed + `g`, H
		case r < 26:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `ju`:
		H := 3.2165060421774707
		r := rng.IntN(42)
		switch {
		case r < 8:
			return seed + `n`, H
		case r < 16:
			return seed + `r`, H
		case r < 22:
			return seed + `s`, H
		case r < 27:
			return seed + `d`, H
		case r < 30:
			return seed + `i`, H
		case r < 33:
			return seed + `m`, H
		case r < 36:
			return seed + `g`, H
		case r < 38:
			return seed + `b`, H
		case r < 39:
			return seed + `j`, H
		case r < 40:
			return seed + `k`, H
		case r < 41:
			return seed + `v`, H
		case r < 42:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `th`:
		H := 2.880261933124179
		r := rng.IntN(129)
		switch {
		case r < 34:
			return seed + `e`, H
		case r < 64:
			return seed + `i`, H
		case r < 83:
			return seed + `r`, H
		case r < 98:
			return seed + `o`, H
		case r < 107:
			return seed + `a`, H
		case r < 114:
			return seed + `y`, H
		case r < 119:
			return seed + `l`, H
		case r < 124:
			return seed + `u`, H
		case r < 126:
			return seed + `w`, H
		case r < 127:
			return seed + `p`, H
		case r < 128:
			return seed + `m`, H
		case r < 129:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `cu`:
		H := 2.8392825556691563
		r := rng.IntN(124)
		switch {
		case r < 38:
			return seed + `r`, H
		case r < 64:
			return seed + `l`, H
		case r < 84:
			return seed + `s`, H
		case r < 96:
			return seed + `p`, H
		case r < 105:
			return seed + `t`, H
		case r < 110:
			return seed + `b`, H
		case r < 113:
			return seed + `f`, H
		case r < 116:
			return seed + `d`, H
		case r < 119:
			return seed + `e`, H
		case r < 122:
			return seed + `m`, H
		case r < 123:
			return seed + `c`, H
		case r < 124:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `ld`:
		H := 3.133525936753556
		r := rng.IntN(33)
		switch {
		case r < 9:
			return seed + `e`, H
		case r < 15:
			return seed + `l`, H
		case r < 18:
			return seed + `f`, H
		case r < 21:
			return seed + `o`, H
		case r < 24:
			return seed + `i`, H
		case r < 27:
			return seed + `c`, H
		case r < 28:
			return seed + `n`, H
		case r < 29:
			return seed + `h`, H
		case r < 30:
			return seed + `s`, H
		case r < 31:
			return seed + `y`, H
		case r < 32:
			return seed + `a`, H
		case r < 33:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ki`:
		H := 2.1171650938789814
		r := rng.IntN(129)
		switch {
		case r < 79:
			return seed + `n`, H
		case r < 96:
			return seed + `l`, H
		case r < 102:
			return seed + `m`, H
		case r < 107:
			return seed + `e`, H
		case r < 112:
			return seed + `s`, H
		case r < 117:
			return seed + `t`, H
		case r < 120:
			return seed + `r`, H
		case r < 123:
			return seed + `p`, H
		case r < 125:
			return seed + `d`, H
		case r < 127:
			return seed + `c`, H
		case r < 128:
			return seed + `w`, H
		case r < 129:
			return seed + `i`, H
		default:
			panic("unexpected rand num")
		}
	case `i`:
		H := 2.5497605472988316
		r := rng.IntN(115)
		switch {
		case r < 59:
			return seed + `m`, H
		case r < 71:
			return seed + `d`, H
		case r < 80:
			return seed + `r`, H
		case r < 88:
			return seed + `s`, H
		case r < 94:
			return seed + `c`, H
		case r < 99:
			return seed + `t`, H
		case r < 103:
			return seed + `g`, H
		case r < 106:
			return seed + `o`, H
		case r < 109:
			return seed + `p`, H
		case r < 112:
			return seed + `l`, H
		case r < 114:
			return seed + `v`, H
		case r < 115:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ui`:
		H := 3.141352734607186
		r := rng.IntN(73)
		switch {
		case r < 17:
			return seed + `t`, H
		case r < 27:
			return seed + `s`, H
		case r < 36:
			return seed + `n`, H
		case r < 44:
			return seed + `r`, H
		case r < 52:
			return seed + `c`, H
		case r < 59:
			return seed + `d`, H
		case r < 66:
			return seed + `l`, H
		case r < 68:
			return seed + `v`, H
		case r < 70:
			return seed + `g`, H
		case r < 71:
			return seed + `e`, H
		case r < 72:
			return seed + `p`, H
		case r < 73:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `om`:
		H := 2.9118142888347975
		r := rng.IntN(122)
		switch {
		case r < 29:
			return seed + `p`, H
		case r < 51:
			return seed + `e`, H
		case r < 72:
			return seed + `i`, H
		case r < 91:
			return seed + `a`, H
		case r < 104:
			return seed + `m`, H
		case r < 110:
			return seed + `b`, H
		case r < 114:
			return seed + `y`, H
		case r < 117:
			return seed + `o`, H
		case r < 119:
			return seed + `f`, H
		case r < 120:
			return seed + `n`, H
		case r < 121:
			return seed + `r`, H
		case r < 122:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ep`:
		H := 3.292881430273992
		r := rng.IntN(93)
		switch {
		case r < 16:
			return seed + `t`, H
		case r < 29:
			return seed + `r`, H
		case r < 41:
			return seed + `i`, H
		case r < 52:
			return seed + `a`, H
		case r < 63:
			return seed + `l`, H
		case r < 71:
			return seed + `u`, H
		case r < 78:
			return seed + `e`, H
		case r < 84:
			return seed + `o`, H
		case r < 87:
			return seed + `h`, H
		case r < 90:
			return seed + `s`, H
		case r < 92:
			return seed + `p`, H
		case r < 93:
			return seed + `n`, H
		default:
			panic("unexpected rand num")
		}
	case `ic`:
		H := 3.0986089299037656
		r := rng.IntN(198)
		switch {
		case r < 41:
			return seed + `a`, H
		case r < 78:
			return seed + `k`, H
		case r < 107:
			return seed + `e`, H
		case r < 134:
			return seed + `i`, H
		case r < 151:
			return seed + `s`, H
		case r < 166:
			return seed + `t`, H
		case r < 179:
			return seed + `o`, H
		case r < 184:
			return seed + `y`, H
		case r < 189:
			return seed + `l`, H
		case r < 193:
			return seed + `h`, H
		case r < 197:
			return seed + `u`, H
		case r < 198:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `us`:
		H := 2.823674243280801
		r := rng.IntN(187)
		switch {
		case r < 52:
			return seed + `t`, H
		case r < 93:
			return seed + `e`, H
		case r < 126:
			return seed + `h`, H
		case r < 148:
			return seed + `i`, H
		case r < 162:
			return seed + `a`, H
		case r < 171:
			return seed + `k`, H
		case r < 175:
			return seed + `p`, H
		case r < 178:
			return seed + `s`, H
		case r < 180:
			return seed + `y`, H
		case r < 182:
			return seed + `l`, H
		case r < 184:
			return seed + `u`, H
		case r < 186:
			return seed + `b`, H
		case r < 187:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ja`:
		H := 3.494680368408909
		r := rng.IntN(28)
		switch {
		case r < 4:
			return seed + `w`, H
		case r < 8:
			return seed + `i`, H
		case r < 12:
			return seed + `c`, H
		case r < 15:
			return seed + `r`, H
		case r < 17:
			return seed + `n`, H
		case r < 19:
			return seed + `y`, H
		case r < 21:
			return seed + `m`, H
		case r < 23:
			return seed + `u`, H
		case r < 24:
			return seed + `s`, H
		case r < 25:
			return seed + `z`, H
		case r < 26:
			return seed + `v`, H
		case r < 27:
			return seed + `l`, H
		case r < 28:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `rm`:
		H := 2.9183678050834025
		r := rng.IntN(67)
		switch {
		case r < 17:
			return seed + `a`, H
		case r < 31:
			return seed + `o`, H
		case r < 42:
			return seed + `e`, H
		case r < 53:
			return seed + `i`, H
		case r < 57:
			return seed + `l`, H
		case r < 59:
			return seed + `f`, H
		case r < 61:
			return seed + `u`, H
		case r < 62:
			return seed + `h`, H
		case r < 63:
			return seed + `y`, H
		case r < 64:
			return seed + `r`, H
		case r < 65:
			return seed + `c`, H
		case r < 66:
			return seed + `p`, H
		case r < 67:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `eg`:
		H := 3.2163328922854326
		r := rng.IntN(70)
		switch {
		case r < 17:
			return seed + `a`, H
		case r < 27:
			return seed + `i`, H
		case r < 36:
			return seed + `g`, H
		case r < 44:
			return seed + `r`, H
		case r < 51:
			return seed + `o`, H
		case r < 57:
			return seed + `u`, H
		case r < 60:
			return seed + `e`, H
		case r < 63:
			return seed + `l`, H
		case r < 65:
			return seed + `n`, H
		case r < 67:
			return seed + `w`, H
		case r < 68:
			return seed + `y`, H
		case r < 69:
			return seed + `m`, H
		case r < 70:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ua`:
		H := 2.9201179414976686
		r := rng.IntN(77)
		switch {
		case r < 25:
			return seed + `l`, H
		case r < 40:
			return seed + `t`, H
		case r < 52:
			return seed + `r`, H
		case r < 57:
			return seed + `d`, H
		case r < 62:
			return seed + `b`, H
		case r < 66:
			return seed + `n`, H
		case r < 69:
			return seed + `i`, H
		case r < 71:
			return seed + `k`, H
		case r < 73:
			return seed + `c`, H
		case r < 74:
			return seed + `h`, H
		case r < 75:
			return seed + `s`, H
		case r < 76:
			return seed + `g`, H
		case r < 77:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `oa`:
		H := 3.0807669557620305
		r := rng.IntN(86)
		switch {
		case r < 18:
			return seed + `d`, H
		case r < 36:
			return seed + `t`, H
		case r < 49:
			return seed + `s`, H
		case r < 62:
			return seed + `r`, H
		case r < 68:
			return seed + `c`, H
		case r < 73:
			return seed + `k`, H
		case r < 76:
			return seed + `n`, H
		case r < 79:
			return seed + `m`, H
		case r < 82:
			return seed + `l`, H
		case r < 83:
			return seed + `f`, H
		case r < 84:
			return seed + `g`, H
		case r < 85:
			return seed + `u`, H
		case r < 86:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ip`:
		H := 2.9084361115661546
		r := rng.IntN(58)
		switch {
		case r < 20:
			return seed + `p`, H
		case r < 33:
			return seed + `e`, H
		case r < 39:
			return seed + `t`, H
		case r < 42:
			return seed + `o`, H
		case r < 45:
			return seed + `h`, H
		case r < 48:
			return seed + `s`, H
		case r < 51:
			return seed + `l`, H
		case r < 53:
			return seed + `a`, H
		case r < 54:
			return seed + `f`, H
		case r < 55:
			return seed + `i`, H
		case r < 56:
			return seed + `c`, H
		case r < 57:
			return seed + `m`, H
		case r < 58:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `id`:
		H := 2.437836737262423
		r := rng.IntN(149)
		switch {
		case r < 78:
			return seed + `e`, H
		case r < 98:
			return seed + `i`, H
		case r < 109:
			return seed + `d`, H
		case r < 119:
			return seed + `a`, H
		case r < 128:
			return seed + `g`, H
		case r < 135:
			return seed + `l`, H
		case r < 138:
			return seed + `o`, H
		case r < 141:
			return seed + `y`, H
		case r < 144:
			return seed + `u`, H
		case r < 146:
			return seed + `n`, H
		case r < 147:
			return seed + `s`, H
		case r < 148:
			return seed + `t`, H
		case r < 149:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `hu`:
		H := 3.2003975417055566
		r := rng.IntN(80)
		switch {
		case r < 21:
			return seed + `m`, H
		case r < 35:
			return seed + `n`, H
		case r < 47:
			return seed + `r`, H
		case r < 54:
			return seed + `s`, H
		case r < 58:
			return seed + `f`, H
		case r < 62:
			return seed + `d`, H
		case r < 66:
			return seed + `t`, H
		case r < 69:
			return seed + `c`, H
		case r < 72:
			return seed + `l`, H
		case r < 74:
			return seed + `a`, H
		case r < 76:
			return seed + `p`, H
		case r < 78:
			return seed + `g`, H
		case r < 80:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `gu`:
		H := 3.2352429992546594
		r := rng.IntN(69)
		switch {
		case r < 13:
			return seed + `l`, H
		case r < 24:
			return seed + `i`, H
		case r < 33:
			return seed + `r`, H
		case r < 41:
			return seed + `e`, H
		case r < 48:
			return seed + `s`, H
		case r < 55:
			return seed + `a`, H
		case r < 61:
			return seed + `m`, H
		case r < 64:
			return seed + `t`, H
		case r < 65:
			return seed + `n`, H
		case r < 66:
			return seed + `o`, H
		case r < 67:
			return seed + `y`, H
		case r < 68:
			return seed + `p`, H
		case r < 69:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `pu`:
		H := 2.953209386359383
		r := rng.IntN(89)
		switch {
		case r < 26:
			return seed + `r`, H
		case r < 44:
			return seed + `l`, H
		case r < 56:
			return seed + `s`, H
		case r < 67:
			return seed + `t`, H
		case r < 74:
			return seed + `n`, H
		case r < 77:
			return seed + `p`, H
		case r < 80:
			return seed + `m`, H
		case r < 82:
			return seed + `z`, H
		case r < 84:
			return seed + `g`, H
		case r < 86:
			return seed + `b`, H
		case r < 87:
			return seed + `d`, H
		case r < 88:
			return seed + `e`, H
		case r < 89:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `im`:
		H := 2.9007368337224446
		r := rng.IntN(181)
		switch {
		case r < 55:
			return seed + `p`, H
		case r < 86:
			return seed + `m`, H
		case r < 114:
			return seed + `e`, H
		case r < 138:
			return seed + `a`, H
		case r < 154:
			return seed + `i`, H
		case r < 160:
			return seed + `b`, H
		case r < 165:
			return seed + `o`, H
		case r < 170:
			return seed + `u`, H
		case r < 173:
			return seed + `n`, H
		case r < 176:
			return seed + `s`, H
		case r < 178:
			return seed + `y`, H
		case r < 180:
			return seed + `l`, H
		case r < 181:
			return seed + `w`, H
		default:
			panic("unexpected rand num")
		}
	case `nt`:
		H := 3.0811605328891254
		r := rng.IntN(219)
		switch {
		case r < 52:
			return seed + `i`, H
		case r < 98:
			return seed + `e`, H
		case r < 127:
			return seed + `a`, H
		case r < 152:
			return seed + `r`, H
		case r < 168:
			return seed + `l`, H
		case r < 180:
			return seed + `o`, H
		case r < 192:
			return seed + `h`, H
		case r < 202:
			return seed + `y`, H
		case r < 209:
			return seed + `u`, H
		case r < 215:
			return seed + `s`, H
		case r < 217:
			return seed + `w`, H
		case r < 218:
			return seed + `d`, H
		case r < 219:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `pi`:
		H := 2.9688147163910608
		r := rng.IntN(129)
		switch {
		case r < 49:
			return seed + `n`, H
		case r < 65:
			return seed + `l`, H
		case r < 79:
			return seed + `r`, H
		case r < 90:
			return seed + `e`, H
		case r < 98:
			return seed + `c`, H
		case r < 104:
			return seed + `d`, H
		case r < 110:
			return seed + `s`, H
		case r < 115:
			return seed + `o`, H
		case r < 120:
			return seed + `t`, H
		case r < 124:
			return seed + `a`, H
		case r < 127:
			return seed + `p`, H
		case r < 128:
			return seed + `f`, H
		case r < 129:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `mp`:
		H := 3.2580382343434384
		r := rng.IntN(143)
		switch {
		case r < 29:
			return seed + `l`, H
		case r < 49:
			return seed + `o`, H
		case r < 68:
			return seed + `e`, H
		case r < 83:
			return seed + `i`, H
		case r < 98:
			return seed + `a`, H
		case r < 110:
			return seed + `r`, H
		case r < 121:
			return seed + `t`, H
		case r < 129:
			return seed + `h`, H
		case r < 136:
			return seed + `u`, H
		case r < 140:
			return seed + `s`, H
		case r < 141:
			return seed + `n`, H
		case r < 142:
			return seed + `f`, H
		case r < 143:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `st`:
		H := 2.956505271065868
		r := rng.IntN(496)
		switch {
		case r < 106:
			return seed + `a`, H
		case r < 200:
			return seed + `e`, H
		case r < 286:
			return seed + `i`, H
		case r < 356:
			return seed + `r`, H
		case r < 413:
			return seed + `o`, H
		case r < 445:
			return seed + `u`, H
		case r < 467:
			return seed + `y`, H
		case r < 478:
			return seed + `l`, H
		case r < 483:
			return seed + `n`, H
		case r < 488:
			return seed + `f`, H
		case r < 492:
			return seed + `b`, H
		case r < 494:
			return seed + `w`, H
		case r < 496:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ie`:
		H := 3.0687146768399054
		r := rng.IntN(136)
		switch {
		case r < 37:
			return seed + `d`, H
		case r < 64:
			return seed + `r`, H
		case r < 83:
			return seed + `n`, H
		case r < 92:
			return seed + `s`, H
		case r < 101:
			return seed + `v`, H
		case r < 110:
			return seed + `l`, H
		case r < 118:
			return seed + `w`, H
		case r < 124:
			return seed + `t`, H
		case r < 129:
			return seed + `f`, H
		case r < 133:
			return seed + `c`, H
		case r < 134:
			return seed + `k`, H
		case r < 135:
			return seed + `m`, H
		case r < 136:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rs`:
		H := 2.9825091489647586
		r := rng.IntN(79)
		switch {
		case r < 23:
			return seed + `e`, H
		case r < 37:
			return seed + `t`, H
		case r < 49:
			return seed + `i`, H
		case r < 57:
			return seed + `h`, H
		case r < 64:
			return seed + `u`, H
		case r < 69:
			return seed + `a`, H
		case r < 72:
			return seed + `o`, H
		case r < 74:
			return seed + `l`, H
		case r < 75:
			return seed + `n`, H
		case r < 76:
			return seed + `w`, H
		case r < 77:
			return seed + `d`, H
		case r < 78:
			return seed + `p`, H
		case r < 79:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `pe`:
		H := 2.860807140148557
		r := rng.IntN(243)
		switch {
		case r < 85:
			return seed + `r`, H
		case r < 130:
			return seed + `n`, H
		case r < 163:
			return seed + `d`, H
		case r < 180:
			return seed + `t`, H
		case r < 196:
			return seed + `c`, H
		case r < 210:
			return seed + `a`, H
		case r < 223:
			return seed + `l`, H
		case r < 230:
			return seed + `s`, H
		case r < 235:
			return seed + `e`, H
		case r < 237:
			return seed + `z`, H
		case r < 239:
			return seed + `g`, H
		case r < 241:
			return seed + `b`, H
		case r < 242:
			return seed + `w`, H
		case r < 243:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `rn`:
		H := 2.9893847578098818
		r := rng.IntN(49)
		switch {
		case r < 14:
			return seed + `e`, H
		case r < 25:
			return seed + `i`, H
		case r < 34:
			return seed + `a`, H
		case r < 37:
			return seed + `f`, H
		case r < 39:
			return seed + `o`, H
		case r < 41:
			return seed + `b`, H
		case r < 42:
			return seed + `n`, H
		case r < 43:
			return seed + `h`, H
		case r < 44:
			return seed + `s`, H
		case r < 45:
			return seed + `y`, H
		case r < 46:
			return seed + `c`, H
		case r < 47:
			return seed + `m`, H
		case r < 48:
			return seed + `l`, H
		case r < 49:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `bu`:
		H := 3.292370012087556
		r := rng.IntN(89)
		switch {
		case r < 24:
			return seed + `l`, H
		case r < 37:
			return seed + `n`, H
		case r < 46:
			return seed + `c`, H
		case r < 54:
			return seed + `s`, H
		case r < 61:
			return seed + `r`, H
		case r < 67:
			return seed + `f`, H
		case r < 72:
			return seed + `d`, H
		case r < 77:
			return seed + `t`, H
		case r < 80:
			return seed + `i`, H
		case r < 83:
			return seed + `g`, H
		case r < 86:
			return seed + `b`, H
		case r < 87:
			return seed + `p`, H
		case r < 88:
			return seed + `m`, H
		case r < 89:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `it`:
		H := 3.0957275161189104
		r := rng.IntN(266)
		switch {
		case r < 75:
			return seed + `y`, H
		case r < 118:
			return seed + `e`, H
		case r < 156:
			return seed + `a`, H
		case r < 184:
			return seed + `i`, H
		case r < 207:
			return seed + `t`, H
		case r < 221:
			return seed + `u`, H
		case r < 234:
			return seed + `c`, H
		case r < 245:
			return seed + `o`, H
		case r < 250:
			return seed + `h`, H
		case r < 255:
			return seed + `r`, H
		case r < 259:
			return seed + `z`, H
		case r < 263:
			return seed + `l`, H
		case r < 265:
			return seed + `s`, H
		case r < 266:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ig`:
		H := 2.973794367913592
		r := rng.IntN(109)
		switch {
		case r < 42:
			return seed + `h`, H
		case r < 55:
			return seed + `a`, H
		case r < 66:
			return seed + `n`, H
		case r < 75:
			return seed + `g`, H
		case r < 82:
			return seed + `e`, H
		case r < 89:
			return seed + `u`, H
		case r < 95:
			return seed + `o`, H
		case r < 100:
			return seed + `i`, H
		case r < 102:
			return seed + `r`, H
		case r < 104:
			return seed + `m`, H
		case r < 106:
			return seed + `l`, H
		case r < 107:
			return seed + `s`, H
		case r < 108:
			return seed + `y`, H
		case r < 109:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ch`:
		H := 2.7422368334249807
		r := rng.IntN(197)
		switch {
		case r < 56:
			return seed + `e`, H
		case r < 109:
			return seed + `a`, H
		case r < 144:
			return seed + `i`, H
		case r < 162:
			return seed + `o`, H
		case r < 172:
			return seed + `u`, H
		case r < 178:
			return seed + `y`, H
		case r < 183:
			return seed + `l`, H
		case r < 188:
			return seed + `b`, H
		case r < 190:
			return seed + `n`, H
		case r < 192:
			return seed + `r`, H
		case r < 194:
			return seed + `m`, H
		case r < 195:
			return seed + `w`, H
		case r < 196:
			return seed + `t`, H
		case r < 197:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ac`:
		H := 3.038335326308152
		r := rng.IntN(278)
		switch {
		case r < 81:
			return seed + `k`, H
		case r < 133:
			return seed + `t`, H
		case r < 166:
			return seed + `e`, H
		case r < 199:
			return seed + `h`, H
		case r < 225:
			return seed + `i`, H
		case r < 235:
			return seed + `a`, H
		case r < 244:
			return seed + `c`, H
		case r < 252:
			return seed + `y`, H
		case r < 259:
			return seed + `o`, H
		case r < 266:
			return seed + `u`, H
		case r < 272:
			return seed + `r`, H
		case r < 275:
			return seed + `q`, H
		case r < 277:
			return seed + `l`, H
		case r < 278:
			return seed + `s`, H
		default:
			panic("unexpected rand num")
		}
	case `ce`:
		H := 3.3511760926287546
		r := rng.IntN(88)
		switch {
		case r < 17:
			return seed + `r`, H
		case r < 30:
			return seed + `n`, H
		case r < 43:
			return seed + `d`, H
		case r < 54:
			return seed + `l`, H
		case r < 62:
			return seed + `s`, H
		case r < 68:
			return seed + `p`, H
		case r < 72:
			return seed + `i`, H
		case r < 76:
			return seed + `a`, H
		case r < 79:
			return seed + `e`, H
		case r < 82:
			return seed + `t`, H
		case r < 85:
			return seed + `m`, H
		case r < 86:
			return seed + `f`, H
		case r < 87:
			return seed + `c`, H
		case r < 88:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ol`:
		H := 3.263456186229182
		r := rng.IntN(174)
		switch {
		case r < 31:
			return seed + `o`, H
		case r < 58:
			return seed + `l`, H
		case r < 83:
			return seed + `d`, H
		case r < 105:
			return seed + `e`, H
		case r < 125:
			return seed + `i`, H
		case r < 143:
			return seed + `a`, H
		case r < 150:
			return seed + `y`, H
		case r < 157:
			return seed + `v`, H
		case r < 163:
			return seed + `t`, H
		case r < 167:
			return seed + `u`, H
		case r < 170:
			return seed + `k`, H
		case r < 172:
			return seed + `f`, H
		case r < 173:
			return seed + `s`, H
		case r < 174:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `su`:
		H := 3.1629317012719462
		r := rng.IntN(166)
		switch {
		case r < 45:
			return seed + `b`, H
		case r < 80:
			return seed + `r`, H
		case r < 98:
			return seed + `p`, H
		case r < 110:
			return seed + `l`, H
		case r < 121:
			return seed + `a`, H
		case r < 131:
			return seed + `i`, H
		case r < 140:
			return seed + `s`, H
		case r < 147:
			return seed + `m`, H
		case r < 153:
			return seed + `f`, H
		case r < 157:
			return seed + `e`, H
		case r < 160:
			return seed + `d`, H
		case r < 163:
			return seed + `c`, H
		case r < 165:
			return seed + `g`, H
		case r < 166:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `fa`:
		H := 3.2099575789421095
		r := rng.IntN(80)
		switch {
		case r < 23:
			return seed + `c`, H
		case r < 36:
			return seed + `n`, H
		case r < 45:
			return seed + `l`, H
		case r < 52:
			return seed + `s`, H
		case r < 58:
			return seed + `m`, H
		case r < 63:
			return seed + `v`, H
		case r < 67:
			return seed + `i`, H
		case r < 70:
			return seed + `r`, H
		case r < 73:
			return seed + `b`, H
		case r < 75:
			return seed + `d`, H
		case r < 77:
			return seed + `u`, H
		case r < 78:
			return seed + `t`, H
		case r < 79:
			return seed + `z`, H
		case r < 80:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ab`:
		H := 2.0438133422038423
		r := rng.IntN(220)
		switch {
		case r < 143:
			return seed + `l`, H
		case r < 158:
			return seed + `i`, H
		case r < 171:
			return seed + `o`, H
		case r < 184:
			return seed + `b`, H
		case r < 194:
			return seed + `s`, H
		case r < 202:
			return seed + `r`, H
		case r < 207:
			return seed + `a`, H
		case r < 211:
			return seed + `e`, H
		case r < 214:
			return seed + `d`, H
		case r < 216:
			return seed + `y`, H
		case r < 217:
			return seed + `n`, H
		case r < 218:
			return seed + `m`, H
		case r < 219:
			return seed + `g`, H
		case r < 220:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `fi`:
		H := 3.287172025028723
		r := rng.IntN(141)
		switch {
		case r < 34:
			return seed + `n`, H
		case r < 57:
			return seed + `e`, H
		case r < 75:
			return seed + `l`, H
		case r < 88:
			return seed + `s`, H
		case r < 99:
			return seed + `d`, H
		case r < 108:
			return seed + `r`, H
		case r < 115:
			return seed + `c`, H
		case r < 122:
			return seed + `g`, H
		case r < 128:
			return seed + `t`, H
		case r < 133:
			return seed + `x`, H
		case r < 137:
			return seed + `f`, H
		case r < 139:
			return seed + `a`, H
		case r < 140:
			return seed + `v`, H
		case r < 141:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `et`:
		H := 3.247370579505989
		r := rng.IntN(156)
		switch {
		case r < 32:
			return seed + `e`, H
		case r < 60:
			return seed + `i`, H
		case r < 83:
			return seed + `t`, H
		case r < 98:
			return seed + `h`, H
		case r < 112:
			return seed + `r`, H
		case r < 125:
			return seed + `a`, H
		case r < 136:
			return seed + `o`, H
		case r < 141:
			return seed + `y`, H
		case r < 146:
			return seed + `u`, H
		case r < 150:
			return seed + `c`, H
		case r < 152:
			return seed + `f`, H
		case r < 153:
			return seed + `d`, H
		case r < 154:
			return seed + `s`, H
		case r < 155:
			return seed + `z`, H
		case r < 156:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `te`:
		H := 2.3840358108683493
		r := rng.IntN(383)
		switch {
		case r < 168:
			return seed + `r`, H
		case r < 271:
			return seed + `d`, H
		case r < 318:
			return seed + `n`, H
		case r < 336:
			return seed + `e`, H
		case r < 344:
			return seed + `m`, H
		case r < 351:
			return seed + `a`, H
		case r < 357:
			return seed + `l`, H
		case r < 362:
			return seed + `x`, H
		case r < 366:
			return seed + `s`, H
		case r < 370:
			return seed + `c`, H
		case r < 374:
			return seed + `p`, H
		case r < 378:
			return seed + `g`, H
		case r < 380:
			return seed + `t`, H
		case r < 382:
			return seed + `b`, H
		case r < 383:
			return seed + `w`, H
		default:
			panic("unexpected rand num")
		}
	case `ir`:
		H := 3.203988838355143
		r := rng.IntN(110)
		switch {
		case r < 39:
			return seed + `e`, H
		case r < 52:
			return seed + `i`, H
		case r < 64:
			return seed + `t`, H
		case r < 71:
			return seed + `r`, H
		case r < 77:
			return seed + `c`, H
		case r < 82:
			return seed + `s`, H
		case r < 87:
			return seed + `a`, H
		case r < 92:
			return seed + `m`, H
		case r < 96:
			return seed + `d`, H
		case r < 99:
			return seed + `k`, H
		case r < 102:
			return seed + `u`, H
		case r < 105:
			return seed + `l`, H
		case r < 107:
			return seed + `y`, H
		case r < 109:
			return seed + `p`, H
		case r < 110:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `s`:
		H := 3.614984946742652
		r := rng.IntN(1087)
		switch {
		case r < 207:
			return seed + `t`, H
		case r < 322:
			return seed + `h`, H
		case r < 437:
			return seed + `u`, H
		case r < 541:
			return seed + `p`, H
		case r < 633:
			return seed + `a`, H
		case r < 718:
			return seed + `c`, H
		case r < 790:
			return seed + `e`, H
		case r < 847:
			return seed + `l`, H
		case r < 902:
			return seed + `i`, H
		case r < 947:
			return seed + `n`, H
		case r < 985:
			return seed + `k`, H
		case r < 1020:
			return seed + `w`, H
		case r < 1050:
			return seed + `m`, H
		case r < 1073:
			return seed + `q`, H
		case r < 1087:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `wi`:
		H := 3.2044098392197946
		r := rng.IntN(112)
		switch {
		case r < 33:
			return seed + `n`, H
		case r < 51:
			return seed + `l`, H
		case r < 65:
			return seed + `s`, H
		case r < 75:
			return seed + `r`, H
		case r < 83:
			return seed + `d`, H
		case r < 89:
			return seed + `m`, H
		case r < 94:
			return seed + `f`, H
		case r < 99:
			return seed + `t`, H
		case r < 102:
			return seed + `e`, H
		case r < 104:
			return seed + `c`, H
		case r < 106:
			return seed + `p`, H
		case r < 108:
			return seed + `z`, H
		case r < 110:
			return seed + `g`, H
		case r < 111:
			return seed + `k`, H
		case r < 112:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ba`:
		H := 3.20045667779631
		r := rng.IntN(199)
		switch {
		case r < 48:
			return seed + `c`, H
		case r < 77:
			return seed + `r`, H
		case r < 103:
			return seed + `n`, H
		case r < 127:
			return seed + `t`, H
		case r < 146:
			return seed + `l`, H
		case r < 162:
			return seed + `g`, H
		case r < 175:
			return seed + `s`, H
		case r < 182:
			return seed + `b`, H
		case r < 188:
			return seed + `k`, H
		case r < 193:
			return seed + `d`, H
		case r < 195:
			return seed + `f`, H
		case r < 196:
			return seed + `y`, H
		case r < 197:
			return seed + `m`, H
		case r < 198:
			return seed + `z`, H
		case r < 199:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `rt`:
		H := 3.4318639824861625
		r := rng.IntN(122)
		switch {
		case r < 24:
			return seed + `e`, H
		case r < 42:
			return seed + `a`, H
		case r < 59:
			return seed + `i`, H
		case r < 71:
			return seed + `h`, H
		case r < 83:
			return seed + `l`, H
		case r < 91:
			return seed + `u`, H
		case r < 97:
			return seed + `o`, H
		case r < 103:
			return seed + `y`, H
		case r < 108:
			return seed + `n`, H
		case r < 113:
			return seed + `s`, H
		case r < 115:
			return seed + `w`, H
		case r < 117:
			return seed + `r`, H
		case r < 119:
			return seed + `c`, H
		case r < 121:
			return seed + `z`, H
		case r < 122:
			return seed + `f`, H
		default:
			panic("unexpected rand num")
		}
	case `hi`:
		H := 2.918107201430624
		r := rng.IntN(129)
		switch {
		case r < 54:
			return seed + `n`, H
		case r < 70:
			return seed + `l`, H
		case r < 80:
			return seed + `r`, H
		case r < 90:
			return seed + `c`, H
		case r < 99:
			return seed + `p`, H
		case r < 106:
			return seed + `f`, H
		case r < 112:
			return seed + `e`, H
		case r < 117:
			return seed + `m`, H
		case r < 120:
			return seed + `s`, H
		case r < 123:
			return seed + `v`, H
		case r < 125:
			return seed + `t`, H
		case r < 126:
			return seed + `d`, H
		case r < 127:
			return seed + `h`, H
		case r < 128:
			return seed + `a`, H
		case r < 129:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `fo`:
		H := 2.981045025728019
		r := rng.IntN(86)
		switch {
		case r < 26:
			return seed + `o`, H
		case r < 45:
			return seed + `r`, H
		case r < 58:
			return seed + `l`, H
		case r < 65:
			return seed + `u`, H
		case r < 70:
			return seed + `n`, H
		case r < 74:
			return seed + `c`, H
		case r < 76:
			return seed + `s`, H
		case r < 78:
			return seed + `i`, H
		case r < 80:
			return seed + `g`, H
		case r < 81:
			return seed + `w`, H
		case r < 82:
			return seed + `d`, H
		case r < 83:
			return seed + `e`, H
		case r < 84:
			return seed + `y`, H
		case r < 85:
			return seed + `a`, H
		case r < 86:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ag`:
		H := 2.8283474600883034
		r := rng.IntN(179)
		switch {
		case r < 76:
			return seed + `e`, H
		case r < 99:
			return seed + `g`, H
		case r < 115:
			return seed + `i`, H
		case r < 131:
			return seed + `r`, H
		case r < 144:
			return seed + `n`, H
		case r < 154:
			return seed + `o`, H
		case r < 160:
			return seed + `u`, H
		case r < 165:
			return seed + `s`, H
		case r < 169:
			return seed + `a`, H
		case r < 173:
			return seed + `m`, H
		case r < 175:
			return seed + `p`, H
		case r < 176:
			return seed + `f`, H
		case r < 177:
			return seed + `w`, H
		case r < 178:
			return seed + `h`, H
		case r < 179:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `ll`:
		H := 2.9528008740663543
		r := rng.IntN(137)
		switch {
		case r < 34:
			return seed + `y`, H
		case r < 64:
			return seed + `e`, H
		case r < 93:
			return seed + `i`, H
		case r < 106:
			return seed + `a`, H
		case r < 116:
			return seed + `o`, H
		case r < 120:
			return seed + `f`, H
		case r < 123:
			return seed + `s`, H
		case r < 126:
			return seed + `u`, H
		case r < 128:
			return seed + `n`, H
		case r < 130:
			return seed + `d`, H
		case r < 132:
			return seed + `h`, H
		case r < 134:
			return seed + `p`, H
		case r < 135:
			return seed + `w`, H
		case r < 136:
			return seed + `r`, H
		case r < 137:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ru`:
		H := 3.385847293011174
		r := rng.IntN(148)
		switch {
		case r < 33:
			return seed + `s`, H
		case r < 57:
			return seed + `n`, H
		case r < 76:
			return seed + `m`, H
		case r < 89:
			return seed + `d`, H
		case r < 101:
			return seed + `b`, H
		case r < 112:
			return seed + `c`, H
		case r < 119:
			return seed + `t`, H
		case r < 125:
			return seed + `p`, H
		case r < 130:
			return seed + `f`, H
		case r < 135:
			return seed + `e`, H
		case r < 140:
			return seed + `g`, H
		case r < 143:
			return seed + `i`, H
		case r < 146:
			return seed + `l`, H
		case r < 147:
			return seed + `r`, H
		case r < 148:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `am`:
		H := 3.1846111780304867
		r := rng.IntN(178)
		switch {
		case r < 40:
			return seed + `e`, H
		case r < 73:
			return seed + `i`, H
		case r < 103:
			return seed + `p`, H
		case r < 122:
			return seed + `b`, H
		case r < 135:
			return seed + `a`, H
		case r < 145:
			return seed + `o`, H
		case r < 155:
			return seed + `m`, H
		case r < 162:
			return seed + `u`, H
		case r < 167:
			return seed + `l`, H
		case r < 170:
			return seed + `s`, H
		case r < 172:
			return seed + `n`, H
		case r < 174:
			return seed + `y`, H
		case r < 176:
			return seed + `r`, H
		case r < 177:
			return seed + `t`, H
		case r < 178:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `ge`:
		H := 2.964543201961648
		r := rng.IntN(132)
		switch {
		case r < 34:
			return seed + `r`, H
		case r < 62:
			return seed + `n`, H
		case r < 90:
			return seed + `d`, H
		case r < 100:
			return seed + `t`, H
		case r < 107:
			return seed + `o`, H
		case r < 113:
			return seed + `s`, H
		case r < 118:
			return seed + `e`, H
		case r < 123:
			return seed + `l`, H
		case r < 126:
			return seed + `a`, H
		case r < 127:
			return seed + `f`, H
		case r < 128:
			return seed + `y`, H
		case r < 129:
			return seed + `i`, H
		case r < 130:
			return seed + `c`, H
		case r < 131:
			return seed + `m`, H
		case r < 132:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `as`:
		H := 3.0133424407336613
		r := rng.IntN(280)
		switch {
		case r < 79:
			return seed + `t`, H
		case r < 135:
			return seed + `h`, H
		case r < 172:
			return seed + `s`, H
		case r < 202:
			return seed + `i`, H
		case r < 229:
			return seed + `e`, H
		case r < 242:
			return seed + `p`, H
		case r < 253:
			return seed + `c`, H
		case r < 261:
			return seed + `k`, H
		case r < 265:
			return seed + `o`, H
		case r < 269:
			return seed + `a`, H
		case r < 273:
			return seed + `m`, H
		case r < 276:
			return seed + `u`, H
		case r < 278:
			return seed + `y`, H
		case r < 279:
			return seed + `q`, H
		case r < 280:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `mi`:
		H := 2.8826950013624906
		r := rng.IntN(150)
		switch {
		case r < 54:
			return seed + `n`, H
		case r < 77:
			return seed + `s`, H
		case r < 97:
			return seed + `c`, H
		case r < 115:
			return seed + `t`, H
		case r < 128:
			return seed + `l`, H
		case r < 132:
			return seed + `d`, H
		case r < 135:
			return seed + `f`, H
		case r < 138:
			return seed + `x`, H
		case r < 140:
			return seed + `e`, H
		case r < 142:
			return seed + `a`, H
		case r < 144:
			return seed + `z`, H
		case r < 146:
			return seed + `g`, H
		case r < 148:
			return seed + `u`, H
		case r < 149:
			return seed + `r`, H
		case r < 150:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `rd`:
		H := 3.531539645867806
		r := rng.IntN(52)
		switch {
		case r < 11:
			return seed + `e`, H
		case r < 19:
			return seed + `i`, H
		case r < 24:
			return seed + `r`, H
		case r < 29:
			return seed + `l`, H
		case r < 33:
			return seed + `o`, H
		case r < 36:
			return seed + `w`, H
		case r < 39:
			return seed + `c`, H
		case r < 41:
			return seed + `h`, H
		case r < 43:
			return seed + `s`, H
		case r < 45:
			return seed + `y`, H
		case r < 47:
			return seed + `a`, H
		case r < 49:
			return seed + `u`, H
		case r < 50:
			return seed + `n`, H
		case r < 51:
			return seed + `d`, H
		case r < 52:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ns`:
		H := 3.452447729749734
		r := rng.IntN(128)
		switch {
		case r < 24:
			return seed + `e`, H
		case r < 46:
			return seed + `t`, H
		case r < 60:
			return seed + `i`, H
		case r < 73:
			return seed + `u`, H
		case r < 84:
			return seed + `o`, H
		case r < 94:
			return seed + `h`, H
		case r < 103:
			return seed + `a`, H
		case r < 108:
			return seed + `c`, H
		case r < 113:
			return seed + `p`, H
		case r < 117:
			return seed + `l`, H
		case r < 120:
			return seed + `w`, H
		case r < 123:
			return seed + `m`, H
		case r < 125:
			return seed + `n`, H
		case r < 127:
			return seed + `f`, H
		case r < 128:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `fe`:
		H := 3.3380070663648165
		r := rng.IntN(84)
		switch {
		case r < 20:
			return seed + `r`, H
		case r < 31:
			return seed + `n`, H
		case r < 42:
			return seed + `c`, H
		case r < 50:
			return seed + `d`, H
		case r < 58:
			return seed + `s`, H
		case r < 65:
			return seed + `e`, H
		case r < 70:
			return seed + `m`, H
		case r < 74:
			return seed + `l`, H
		case r < 77:
			return seed + `t`, H
		case r < 79:
			return seed + `a`, H
		case r < 80:
			return seed + `w`, H
		case r < 81:
			return seed + `h`, H
		case r < 82:
			return seed + `i`, H
		case r < 83:
			return seed + `v`, H
		case r < 84:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ve`:
		H := 2.0357494674844125
		r := rng.IntN(277)
		switch {
		case r < 169:
			return seed + `r`, H
		case r < 212:
			return seed + `n`, H
		case r < 229:
			return seed + `d`, H
		case r < 245:
			return seed + `l`, H
		case r < 254:
			return seed + `s`, H
		case r < 262:
			return seed + `t`, H
		case r < 265:
			return seed + `y`, H
		case r < 268:
			return seed + `a`, H
		case r < 270:
			return seed + `i`, H
		case r < 272:
			return seed + `g`, H
		case r < 273:
			return seed + `w`, H
		case r < 274:
			return seed + `e`, H
		case r < 275:
			return seed + `h`, H
		case r < 276:
			return seed + `m`, H
		case r < 277:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ke`:
		H := 3.1007666652190577
		r := rng.IntN(113)
		switch {
		case r < 30:
			return seed + `r`, H
		case r < 54:
			return seed + `d`, H
		case r < 69:
			return seed + `n`, H
		case r < 82:
			return seed + `t`, H
		case r < 88:
			return seed + `l`, H
		case r < 93:
			return seed + `w`, H
		case r < 98:
			return seed + `e`, H
		case r < 103:
			return seed + `s`, H
		case r < 105:
			return seed + `y`, H
		case r < 107:
			return seed + `p`, H
		case r < 109:
			return seed + `b`, H
		case r < 110:
			return seed + `o`, H
		case r < 111:
			return seed + `h`, H
		case r < 112:
			return seed + `m`, H
		case r < 113:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ap`:
		H := 2.8684848816641035
		r := rng.IntN(143)
		switch {
		case r < 60:
			return seed + `p`, H
		case r < 82:
			return seed + `e`, H
		case r < 93:
			return seed + `s`, H
		case r < 103:
			return seed + `t`, H
		case r < 111:
			return seed + `h`, H
		case r < 119:
			return seed + `a`, H
		case r < 125:
			return seed + `i`, H
		case r < 130:
			return seed + `r`, H
		case r < 134:
			return seed + `l`, H
		case r < 137:
			return seed + `o`, H
		case r < 139:
			return seed + `d`, H
		case r < 140:
			return seed + `n`, H
		case r < 141:
			return seed + `k`, H
		case r < 142:
			return seed + `y`, H
		case r < 143:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `gi`:
		H := 3.019193877281908
		r := rng.IntN(102)
		switch {
		case r < 43:
			return seed + `n`, H
		case r < 51:
			return seed + `s`, H
		case r < 59:
			return seed + `l`, H
		case r < 66:
			return seed + `g`, H
		case r < 72:
			return seed + `d`, H
		case r < 78:
			return seed + `c`, H
		case r < 84:
			return seed + `b`, H
		case r < 89:
			return seed + `v`, H
		case r < 92:
			return seed + `e`, H
		case r < 95:
			return seed + `z`, H
		case r < 97:
			return seed + `r`, H
		case r < 99:
			return seed + `m`, H
		case r < 100:
			return seed + `f`, H
		case r < 101:
			return seed + `o`, H
		case r < 102:
			return seed + `a`, H
		default:
			panic("unexpected rand num")
		}
	case `ob`:
		H := 3.2815794566249576
		r := rng.IntN(69)
		switch {
		case r < 14:
			return seed + `s`, H
		case r < 25:
			return seed + `i`, H
		case r < 36:
			return seed + `b`, H
		case r < 45:
			return seed + `e`, H
		case r < 51:
			return seed + `l`, H
		case r < 56:
			return seed + `a`, H
		case r < 60:
			return seed + `t`, H
		case r < 62:
			return seed + `o`, H
		case r < 63:
			return seed + `n`, H
		case r < 64:
			return seed + `w`, H
		case r < 65:
			return seed + `j`, H
		case r < 66:
			return seed + `y`, H
		case r < 67:
			return seed + `c`, H
		case r < 68:
			return seed + `v`, H
		case r < 69:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `vi`:
		H := 3.350995470822189
		r := rng.IntN(189)
		switch {
		case r < 45:
			return seed + `n`, H
		case r < 77:
			return seed + `s`, H
		case r < 94:
			return seed + `d`, H
		case r < 110:
			return seed + `t`, H
		case r < 125:
			return seed + `a`, H
		case r < 138:
			return seed + `o`, H
		case r < 149:
			return seed + `v`, H
		case r < 159:
			return seed + `e`, H
		case r < 168:
			return seed + `c`, H
		case r < 177:
			return seed + `l`, H
		case r < 183:
			return seed + `r`, H
		case r < 186:
			return seed + `g`, H
		case r < 187:
			return seed + `p`, H
		case r < 188:
			return seed + `x`, H
		case r < 189:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `he`:
		H := 3.1299537542181337
		r := rng.IntN(227)
		switch {
		case r < 56:
			return seed + `r`, H
		case r < 109:
			return seed + `a`, H
		case r < 138:
			return seed + `d`, H
		case r < 157:
			return seed + `l`, H
		case r < 173:
			return seed + `s`, H
		case r < 186:
			return seed + `m`, H
		case r < 197:
			return seed + `e`, H
		case r < 207:
			return seed + `n`, H
		case r < 212:
			return seed + `w`, H
		case r < 216:
			return seed + `f`, H
		case r < 219:
			return seed + `o`, H
		case r < 221:
			return seed + `t`, H
		case r < 223:
			return seed + `c`, H
		case r < 225:
			return seed + `v`, H
		case r < 227:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `nu`:
		H := 3.4253282651341452
		r := rng.IntN(57)
		switch {
		case r < 14:
			return seed + `m`, H
		case r < 25:
			return seed + `t`, H
		case r < 30:
			return seed + `s`, H
		case r < 34:
			return seed + `a`, H
		case r < 38:
			return seed + `g`, H
		case r < 41:
			return seed + `e`, H
		case r < 44:
			return seed + `r`, H
		case r < 47:
			return seed + `c`, H
		case r < 49:
			return seed + `o`, H
		case r < 51:
			return seed + `l`, H
		case r < 52:
			return seed + `n`, H
		case r < 53:
			return seed + `f`, H
		case r < 54:
			return seed + `i`, H
		case r < 55:
			return seed + `p`, H
		case r < 56:
			return seed + `z`, H
		case r < 57:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `lu`:
		H := 3.5210705994541325
		r := rng.IntN(127)
		switch {
		case r < 27:
			return seed + `s`, H
		case r < 41:
			return seed + `t`, H
		case r < 54:
			return seed + `m`, H
		case r < 66:
			return seed + `n`, H
		case r < 77:
			return seed + `r`, H
		case r < 87:
			return seed + `c`, H
		case r < 96:
			return seed + `d`, H
		case r < 104:
			return seed + `e`, H
		case r < 110:
			return seed + `g`, H
		case r < 116:
			return seed + `b`, H
		case r < 120:
			return seed + `x`, H
		case r < 123:
			return seed + `a`, H
		case r < 124:
			return seed + `f`, H
		case r < 125:
			return seed + `i`, H
		case r < 126:
			return seed + `k`, H
		case r < 127:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `so`:
		H := 2.98458260307347
		r := rng.IntN(58)
		switch {
		case r < 16:
			return seed + `l`, H
		case r < 31:
			return seed + `n`, H
		case r < 42:
			return seed + `r`, H
		case r < 44:
			return seed + `i`, H
		case r < 46:
			return seed + `t`, H
		case r < 48:
			return seed + `u`, H
		case r < 49:
			return seed + `f`, H
		case r < 50:
			return seed + `w`, H
		case r < 51:
			return seed + `d`, H
		case r < 52:
			return seed + `o`, H
		case r < 53:
			return seed + `a`, H
		case r < 54:
			return seed + `c`, H
		case r < 55:
			return seed + `p`, H
		case r < 56:
			return seed + `m`, H
		case r < 57:
			return seed + `v`, H
		case r < 58:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `at`:
		H := 2.98197939609681
		r := rng.IntN(506)
		switch {
		case r < 182:
			return seed + `e`, H
		case r < 271:
			return seed + `i`, H
		case r < 319:
			return seed + `t`, H
		case r < 365:
			return seed + `o`, H
		case r < 395:
			return seed + `c`, H
		case r < 420:
			return seed + `h`, H
		case r < 441:
			return seed + `u`, H
		case r < 460:
			return seed + `r`, H
		case r < 479:
			return seed + `a`, H
		case r < 485:
			return seed + `l`, H
		case r < 490:
			return seed + `s`, H
		case r < 494:
			return seed + `n`, H
		case r < 498:
			return seed + `f`, H
		case r < 501:
			return seed + `w`, H
		case r < 504:
			return seed + `b`, H
		case r < 506:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `op`:
		H := 3.240850086008522
		r := rng.IntN(106)
		switch {
		case r < 29:
			return seed + `e`, H
		case r < 50:
			return seed + `p`, H
		case r < 63:
			return seed + `i`, H
		case r < 69:
			return seed + `o`, H
		case r < 75:
			return seed + `s`, H
		case r < 81:
			return seed + `l`, H
		case r < 87:
			return seed + `u`, H
		case r < 92:
			return seed + `h`, H
		case r < 96:
			return seed + `y`, H
		case r < 99:
			return seed + `a`, H
		case r < 101:
			return seed + `t`, H
		case r < 102:
			return seed + `w`, H
		case r < 103:
			return seed + `d`, H
		case r < 104:
			return seed + `k`, H
		case r < 105:
			return seed + `c`, H
		case r < 106:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `oo`:
		H := 3.489791824124952
		r := rng.IntN(195)
		switch {
		case r < 40:
			return seed + `t`, H
		case r < 67:
			return seed + `k`, H
		case r < 92:
			return seed + `n`, H
		case r < 111:
			return seed + `m`, H
		case r < 124:
			return seed + `d`, H
		case r < 137:
			return seed + `r`, H
		case r < 149:
			return seed + `f`, H
		case r < 161:
			return seed + `l`, H
		case r < 171:
			return seed + `p`, H
		case r < 177:
			return seed + `s`, H
		case r < 183:
			return seed + `z`, H
		case r < 188:
			return seed + `g`, H
		case r < 191:
			return seed + `v`, H
		case r < 193:
			return seed + `i`, H
		case r < 194:
			return seed + `e`, H
		case r < 195:
			return seed + `c`, H
		default:
			panic("unexpected rand num")
		}
	case `go`:
		H := 3.4079256616775075
		r := rng.IntN(67)
		switch {
		case r < 16:
			return seed + `n`, H
		case r < 25:
			return seed + `o`, H
		case r < 33:
			return seed + `r`, H
		case r < 40:
			return seed + `t`, H
		case r < 47:
			return seed + `l`, H
		case r < 51:
			return seed + `u`, H
		case r < 54:
			return seed + `s`, H
		case r < 57:
			return seed + `i`, H
		case r < 60:
			return seed + `a`, H
		case r < 61:
			return seed + `w`, H
		case r < 62:
			return seed + `d`, H
		case r < 63:
			return seed + `e`, H
		case r < 64:
			return seed + `p`, H
		case r < 65:
			return seed + `m`, H
		case r < 66:
			return seed + `g`, H
		case r < 67:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `es`:
		H := 2.2514347845496214
		r := rng.IntN(405)
		switch {
		case r < 232:
			return seed + `s`, H
		case r < 299:
			return seed + `t`, H
		case r < 323:
			return seed + `i`, H
		case r < 341:
			return seed + `e`, H
		case r < 352:
			return seed + `h`, H
		case r < 362:
			return seed + `o`, H
		case r < 372:
			return seed + `c`, H
		case r < 380:
			return seed + `p`, H
		case r < 388:
			return seed + `u`, H
		case r < 394:
			return seed + `a`, H
		case r < 399:
			return seed + `k`, H
		case r < 401:
			return seed + `m`, H
		case r < 402:
			return seed + `d`, H
		case r < 403:
			return seed + `y`, H
		case r < 404:
			return seed + `q`, H
		case r < 405:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `ci`:
		H := 3.465666807234371
		r := rng.IntN(132)
		switch {
		case r < 31:
			return seed + `n`, H
		case r < 52:
			return seed + `t`, H
		case r < 64:
			return seed + `d`, H
		case r < 75:
			return seed + `s`, H
		case r < 85:
			return seed + `a`, H
		case r < 93:
			return seed + `e`, H
		case r < 101:
			return seed + `o`, H
		case r < 108:
			return seed + `l`, H
		case r < 114:
			return seed + `r`, H
		case r < 119:
			return seed + `f`, H
		case r < 123:
			return seed + `v`, H
		case r < 126:
			return seed + `p`, H
		case r < 129:
			return seed + `m`, H
		case r < 130:
			return seed + `z`, H
		case r < 131:
			return seed + `u`, H
		case r < 132:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `sh`:
		H := 3.039460068107811
		r := rng.IntN(233)
		switch {
		case r < 48:
			return seed + `o`, H
		case r < 91:
			return seed + `i`, H
		case r < 133:
			return seed + `e`, H
		case r < 174:
			return seed + `a`, H
		case r < 192:
			return seed + `r`, H
		case r < 202:
			return seed + `y`, H
		case r < 212:
			return seed + `u`, H
		case r < 218:
			return seed + `l`, H
		case r < 223:
			return seed + `b`, H
		case r < 227:
			return seed + `c`, H
		case r < 228:
			return seed + `n`, H
		case r < 229:
			return seed + `d`, H
		case r < 230:
			return seed + `h`, H
		case r < 231:
			return seed + `s`, H
		case r < 232:
			return seed + `t`, H
		case r < 233:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ea`:
		H := 3.286394666608718
		r := rng.IntN(296)
		switch {
		case r < 59:
			return seed + `r`, H
		case r < 110:
			return seed + `d`, H
		case r < 159:
			return seed + `t`, H
		case r < 195:
			return seed + `s`, H
		case r < 218:
			return seed + `l`, H
		case r < 237:
			return seed + `m`, H
		case r < 255:
			return seed + `c`, H
		case r < 264:
			return seed + `k`, H
		case r < 271:
			return seed + `n`, H
		case r < 278:
			return seed + `v`, H
		case r < 283:
			return seed + `p`, H
		case r < 288:
			return seed + `b`, H
		case r < 291:
			return seed + `f`, H
		case r < 293:
			return seed + `w`, H
		case r < 295:
			return seed + `g`, H
		case r < 296:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `wa`:
		H := 3.402126204758526
		r := rng.IntN(120)
		switch {
		case r < 25:
			return seed + `r`, H
		case r < 48:
			return seed + `s`, H
		case r < 67:
			return seed + `y`, H
		case r < 77:
			return seed + `l`, H
		case r < 83:
			return seed + `n`, H
		case r < 89:
			return seed + `t`, H
		case r < 95:
			return seed + `g`, H
		case r < 100:
			return seed + `k`, H
		case r < 105:
			return seed + `b`, H
		case r < 109:
			return seed + `v`, H
		case r < 112:
			return seed + `d`, H
		case r < 114:
			return seed + `f`, H
		case r < 116:
			return seed + `w`, H
		case r < 117:
			return seed + `i`, H
		case r < 118:
			return seed + `c`, H
		case r < 119:
			return seed + `p`, H
		case r < 120:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `ay`:
		H := 3.9021061909687025
		r := rng.IntN(53)
		switch {
		case r < 6:
			return seed + `e`, H
		case r < 12:
			return seed + `b`, H
		case r < 17:
			return seed + `s`, H
		case r < 22:
			return seed + `i`, H
		case r < 26:
			return seed + `m`, H
		case r < 29:
			return seed + `f`, H
		case r < 32:
			return seed + `d`, H
		case r < 35:
			return seed + `r`, H
		case r < 38:
			return seed + `t`, H
		case r < 41:
			return seed + `a`, H
		case r < 44:
			return seed + `l`, H
		case r < 46:
			return seed + `o`, H
		case r < 48:
			return seed + `c`, H
		case r < 50:
			return seed + `p`, H
		case r < 51:
			return seed + `w`, H
		case r < 52:
			return seed + `h`, H
		case r < 53:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `up`:
		H := 3.696077868346799
		r := rng.IntN(86)
		switch {
		case r < 12:
			return seed + `e`, H
		case r < 23:
			return seed + `p`, H
		case r < 33:
			return seed + `t`, H
		case r < 42:
			return seed + `s`, H
		case r < 50:
			return seed + `r`, H
		case r < 58:
			return seed + `l`, H
		case r < 63:
			return seed + `h`, H
		case r < 68:
			return seed + `i`, H
		case r < 72:
			return seed + `c`, H
		case r < 75:
			return seed + `y`, H
		case r < 77:
			return seed + `w`, H
		case r < 79:
			return seed + `o`, H
		case r < 81:
			return seed + `a`, H
		case r < 83:
			return seed + `b`, H
		case r < 84:
			return seed + `f`, H
		case r < 85:
			return seed + `d`, H
		case r < 86:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `no`:
		H := 3.578556177289845
		r := rng.IntN(101)
		switch {
		case r < 19:
			return seed + `w`, H
		case r < 34:
			return seed + `t`, H
		case r < 47:
			return seed + `r`, H
		case r < 55:
			return seed + `u`, H
		case r < 62:
			return seed + `s`, H
		case r < 69:
			return seed + `p`, H
		case r < 76:
			return seed + `m`, H
		case r < 83:
			return seed + `l`, H
		case r < 86:
			return seed + `o`, H
		case r < 89:
			return seed + `g`, H
		case r < 92:
			return seed + `x`, H
		case r < 95:
			return seed + `v`, H
		case r < 97:
			return seed + `n`, H
		case r < 98:
			return seed + `e`, H
		case r < 99:
			return seed + `y`, H
		case r < 100:
			return seed + `i`, H
		case r < 101:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `tu`:
		H := 3.249185448350641
		r := rng.IntN(135)
		switch {
		case r < 50:
			return seed + `r`, H
		case r < 63:
			return seed + `a`, H
		case r < 74:
			return seed + `d`, H
		case r < 83:
			return seed + `m`, H
		case r < 91:
			return seed + `n`, H
		case r < 99:
			return seed + `b`, H
		case r < 105:
			return seed + `s`, H
		case r < 111:
			return seed + `t`, H
		case r < 116:
			return seed + `f`, H
		case r < 121:
			return seed + `c`, H
		case r < 126:
			return seed + `p`, H
		case r < 128:
			return seed + `e`, H
		case r < 130:
			return seed + `i`, H
		case r < 132:
			return seed + `l`, H
		case r < 133:
			return seed + `o`, H
		case r < 134:
			return seed + `g`, H
		case r < 135:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `da`:
		H := 3.488123022570739
		r := rng.IntN(107)
		switch {
		case r < 19:
			return seed + `n`, H
		case r < 37:
			return seed + `y`, H
		case r < 53:
			return seed + `r`, H
		case r < 67:
			return seed + `t`, H
		case r < 74:
			return seed + `b`, H
		case r < 80:
			return seed + `i`, H
		case r < 85:
			return seed + `l`, H
		case r < 89:
			return seed + `c`, H
		case r < 92:
			return seed + `d`, H
		case r < 95:
			return seed + `s`, H
		case r < 98:
			return seed + `u`, H
		case r < 100:
			return seed + `w`, H
		case r < 102:
			return seed + `z`, H
		case r < 104:
			return seed + `g`, H
		case r < 105:
			return seed + `f`, H
		case r < 106:
			return seed + `m`, H
		case r < 107:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `pa`:
		H := 3.51517726223104
		r := rng.IntN(214)
		switch {
		case r < 48:
			return seed + `r`, H
		case r < 74:
			return seed + `n`, H
		case r < 99:
			return seed + `s`, H
		case r < 120:
			return seed + `c`, H
		case r < 139:
			return seed + `t`, H
		case r < 155:
			return seed + `y`, H
		case r < 165:
			return seed + `l`, H
		case r < 174:
			return seed + `d`, H
		case r < 183:
			return seed + `i`, H
		case r < 190:
			return seed + `p`, H
		case r < 197:
			return seed + `v`, H
		case r < 203:
			return seed + `b`, H
		case r < 207:
			return seed + `m`, H
		case r < 211:
			return seed + `g`, H
		case r < 212:
			return seed + `w`, H
		case r < 213:
			return seed + `j`, H
		case r < 214:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ng`:
		H := 3.17500625470462
		r := rng.IntN(145)
		switch {
		case r < 39:
			return seed + `e`, H
		case r < 72:
			return seed + `l`, H
		case r < 87:
			return seed + `r`, H
		case r < 98:
			return seed + `i`, H
		case r < 108:
			return seed + `o`, H
		case r < 118:
			return seed + `u`, H
		case r < 124:
			return seed + `s`, H
		case r < 130:
			return seed + `a`, H
		case r < 134:
			return seed + `y`, H
		case r < 136:
			return seed + `n`, H
		case r < 138:
			return seed + `t`, H
		case r < 140:
			return seed + `b`, H
		case r < 141:
			return seed + `f`, H
		case r < 142:
			return seed + `w`, H
		case r < 143:
			return seed + `d`, H
		case r < 144:
			return seed + `h`, H
		case r < 145:
			return seed + `m`, H
		default:
			panic("unexpected rand num")
		}
	case `mo`:
		H := 3.623419066149199
		r := rng.IntN(203)
		switch {
		case r < 41:
			return seed + `n`, H
		case r < 72:
			return seed + `r`, H
		case r < 92:
			return seed + `t`, H
		case r < 108:
			return seed + `o`, H
		case r < 124:
			return seed + `u`, H
		case r < 138:
			return seed + `v`, H
		case r < 150:
			return seed + `s`, H
		case r < 161:
			return seed + `l`, H
		case r < 171:
			return seed + `d`, H
		case r < 178:
			return seed + `c`, H
		case r < 183:
			return seed + `i`, H
		case r < 188:
			return seed + `k`, H
		case r < 193:
			return seed + `b`, H
		case r < 196:
			return seed + `m`, H
		case r < 198:
			return seed + `w`, H
		case r < 200:
			return seed + `a`, H
		case r < 202:
			return seed + `g`, H
		case r < 203:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `ho`:
		H := 3.5246233511381844
		r := rng.IntN(127)
		switch {
		case r < 20:
			return seed + `r`, H
		case r < 38:
			return seed + `w`, H
		case r < 55:
			return seed + `o`, H
		case r < 71:
			return seed + `n`, H
		case r < 85:
			return seed + `u`, H
		case r < 97:
			return seed + `l`, H
		case r < 103:
			return seed + `p`, H
		case r < 108:
			return seed + `t`, H
		case r < 112:
			return seed + `s`, H
		case r < 115:
			return seed + `e`, H
		case r < 118:
			return seed + `v`, H
		case r < 120:
			return seed + `k`, H
		case r < 122:
			return seed + `m`, H
		case r < 123:
			return seed + `y`, H
		case r < 124:
			return seed + `i`, H
		case r < 125:
			return seed + `c`, H
		case r < 126:
			return seed + `g`, H
		case r < 127:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `el`:
		H := 3.4963439564868266
		r := rng.IntN(231)
		switch {
		case r < 44:
			return seed + `e`, H
		case r < 88:
			return seed + `i`, H
		case r < 116:
			return seed + `l`, H
		case r < 138:
			return seed + `y`, H
		case r < 153:
			return seed + `d`, H
		case r < 168:
			return seed + `a`, H
		case r < 182:
			return seed + `o`, H
		case r < 191:
			return seed + `t`, H
		case r < 199:
			return seed + `f`, H
		case r < 207:
			return seed + `p`, H
		case r < 214:
			return seed + `u`, H
		case r < 220:
			return seed + `v`, H
		case r < 223:
			return seed + `s`, H
		case r < 225:
			return seed + `n`, H
		case r < 227:
			return seed + `c`, H
		case r < 229:
			return seed + `m`, H
		case r < 230:
			return seed + `k`, H
		case r < 231:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `ow`:
		H := 3.593817601124209
		r := rng.IntN(102)
		switch {
		case r < 22:
			return seed + `n`, H
		case r < 38:
			return seed + `e`, H
		case r < 47:
			return seed + `d`, H
		case r < 56:
			return seed + `i`, H
		case r < 65:
			return seed + `l`, H
		case r < 72:
			return seed + `s`, H
		case r < 79:
			return seed + `b`, H
		case r < 82:
			return seed + `f`, H
		case r < 85:
			return seed + `a`, H
		case r < 88:
			return seed + `p`, H
		case r < 91:
			return seed + `m`, H
		case r < 93:
			return seed + `w`, H
		case r < 95:
			return seed + `r`, H
		case r < 97:
			return seed + `y`, H
		case r < 99:
			return seed + `c`, H
		case r < 100:
			return seed + `o`, H
		case r < 101:
			return seed + `t`, H
		case r < 102:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ul`:
		H := 3.3475651004206015
		r := rng.IntN(140)
		switch {
		case r < 33:
			return seed + `a`, H
		case r < 63:
			return seed + `l`, H
		case r < 81:
			return seed + `t`, H
		case r < 93:
			return seed + `e`, H
		case r < 101:
			return seed + `p`, H
		case r < 108:
			return seed + `s`, H
		case r < 114:
			return seed + `f`, H
		case r < 120:
			return seed + `i`, H
		case r < 123:
			return seed + `y`, H
		case r < 126:
			return seed + `k`, H
		case r < 129:
			return seed + `g`, H
		case r < 132:
			return seed + `b`, H
		case r < 134:
			return seed + `m`, H
		case r < 136:
			return seed + `u`, H
		case r < 137:
			return seed + `d`, H
		case r < 138:
			return seed + `o`, H
		case r < 139:
			return seed + `c`, H
		case r < 140:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ot`:
		H := 3.4151997626610338
		r := rng.IntN(138)
		switch {
		case r < 27:
			return seed + `i`, H
		case r < 50:
			return seed + `e`, H
		case r < 70:
			return seed + `h`, H
		case r < 90:
			return seed + `t`, H
		case r < 100:
			return seed + `a`, H
		case r < 109:
			return seed + `o`, H
		case r < 114:
			return seed + `l`, H
		case r < 118:
			return seed + `s`, H
		case r < 121:
			return seed + `r`, H
		case r < 124:
			return seed + `c`, H
		case r < 127:
			return seed + `p`, H
		case r < 130:
			return seed + `b`, H
		case r < 132:
			return seed + `w`, H
		case r < 134:
			return seed + `y`, H
		case r < 135:
			return seed + `n`, H
		case r < 136:
			return seed + `m`, H
		case r < 137:
			return seed + `g`, H
		case r < 138:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ca`:
		H := 3.301798879691526
		r := rng.IntN(317)
		switch {
		case r < 71:
			return seed + `r`, H
		case r < 126:
			return seed + `l`, H
		case r < 176:
			return seed + `t`, H
		case r < 210:
			return seed + `n`, H
		case r < 236:
			return seed + `p`, H
		case r < 255:
			return seed + `s`, H
		case r < 267:
			return seed + `m`, H
		case r < 277:
			return seed + `b`, H
		case r < 286:
			return seed + `u`, H
		case r < 294:
			return seed + `d`, H
		case r < 302:
			return seed + `v`, H
		case r < 307:
			return seed + `c`, H
		case r < 311:
			return seed + `k`, H
		case r < 313:
			return seed + `g`, H
		case r < 314:
			return seed + `f`, H
		case r < 315:
			return seed + `w`, H
		case r < 316:
			return seed + `h`, H
		case r < 317:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `po`:
		H := 3.4558028582947937
		r := rng.IntN(189)
		switch {
		case r < 47:
			return seed + `s`, H
		case r < 77:
			return seed + `r`, H
		case r < 94:
			return seed + `i`, H
		case r < 110:
			return seed + `l`, H
		case r < 123:
			return seed + `u`, H
		case r < 135:
			return seed + `n`, H
		case r < 145:
			return seed + `t`, H
		case r < 155:
			return seed + `p`, H
		case r < 164:
			return seed + `w`, H
		case r < 172:
			return seed + `o`, H
		case r < 178:
			return seed + `k`, H
		case r < 181:
			return seed + `d`, H
		case r < 183:
			return seed + `e`, H
		case r < 185:
			return seed + `c`, H
		case r < 186:
			return seed + `f`, H
		case r < 187:
			return seed + `a`, H
		case r < 188:
			return seed + `g`, H
		case r < 189:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `du`:
		H := 3.4682415026057893
		r := rng.IntN(71)
		switch {
		case r < 20:
			return seed + `c`, H
		case r < 30:
			return seed + `r`, H
		case r < 38:
			return seed + `l`, H
		case r < 43:
			return seed + `s`, H
		case r < 48:
			return seed + `p`, H
		case r < 52:
			return seed + `e`, H
		case r < 55:
			return seed + `a`, H
		case r < 58:
			return seed + `m`, H
		case r < 60:
			return seed + `i`, H
		case r < 62:
			return seed + `t`, H
		case r < 64:
			return seed + `b`, H
		case r < 65:
			return seed + `f`, H
		case r < 66:
			return seed + `d`, H
		case r < 67:
			return seed + `o`, H
		case r < 68:
			return seed + `h`, H
		case r < 69:
			return seed + `k`, H
		case r < 70:
			return seed + `g`, H
		case r < 71:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ck`:
		H := 3.3939366348296245
		r := rng.IntN(135)
		switch {
		case r < 30:
			return seed + `e`, H
		case r < 51:
			return seed + `i`, H
		case r < 72:
			return seed + `l`, H
		case r < 91:
			return seed + `s`, H
		case r < 100:
			return seed + `y`, H
		case r < 104:
			return seed + `n`, H
		case r < 108:
			return seed + `w`, H
		case r < 112:
			return seed + `t`, H
		case r < 116:
			return seed + `a`, H
		case r < 120:
			return seed + `b`, H
		case r < 123:
			return seed + `f`, H
		case r < 126:
			return seed + `p`, H
		case r < 129:
			return seed + `u`, H
		case r < 131:
			return seed + `r`, H
		case r < 132:
			return seed + `d`, H
		case r < 133:
			return seed + `o`, H
		case r < 134:
			return seed + `h`, H
		case r < 135:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `il`:
		H := 3.0714574670306756
		r := rng.IntN(249)
		switch {
		case r < 57:
			return seed + `l`, H
		case r < 106:
			return seed + `y`, H
		case r < 147:
			return seed + `e`, H
		case r < 186:
			return seed + `i`, H
		case r < 201:
			return seed + `d`, H
		case r < 213:
			return seed + `t`, H
		case r < 225:
			return seed + `a`, H
		case r < 235:
			return seed + `o`, H
		case r < 237:
			return seed + `s`, H
		case r < 239:
			return seed + `u`, H
		case r < 241:
			return seed + `b`, H
		case r < 242:
			return seed + `n`, H
		case r < 243:
			return seed + `w`, H
		case r < 244:
			return seed + `h`, H
		case r < 245:
			return seed + `r`, H
		case r < 246:
			return seed + `k`, H
		case r < 247:
			return seed + `c`, H
		case r < 248:
			return seed + `m`, H
		case r < 249:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ni`:
		H := 3.452831684393273
		r := rng.IntN(202)
		switch {
		case r < 52:
			return seed + `n`, H
		case r < 80:
			return seed + `s`, H
		case r < 104:
			return seed + `t`, H
		case r < 125:
			return seed + `c`, H
		case r < 137:
			return seed + `f`, H
		case r < 148:
			return seed + `m`, H
		case r < 158:
			return seed + `a`, H
		case r < 167:
			return seed + `o`, H
		case r < 176:
			return seed + `z`, H
		case r < 181:
			return seed + `p`, H
		case r < 186:
			return seed + `v`, H
		case r < 191:
			return seed + `u`, H
		case r < 195:
			return seed + `l`, H
		case r < 197:
			return seed + `g`, H
		case r < 198:
			return seed + `d`, H
		case r < 199:
			return seed + `e`, H
		case r < 200:
			return seed + `q`, H
		case r < 201:
			return seed + `x`, H
		case r < 202:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `or`:
		H := 3.7615134537945343
		r := rng.IntN(290)
		switch {
		case r < 50:
			return seed + `t`, H
		case r < 86:
			return seed + `e`, H
		case r < 113:
			return seed + `i`, H
		case r < 139:
			return seed + `n`, H
		case r < 164:
			return seed + `a`, H
		case r < 187:
			return seed + `m`, H
		case r < 204:
			return seed + `k`, H
		case r < 219:
			return seed + `r`, H
		case r < 232:
			return seed + `d`, H
		case r < 244:
			return seed + `y`, H
		case r < 256:
			return seed + `s`, H
		case r < 264:
			return seed + `o`, H
		case r < 271:
			return seed + `c`, H
		case r < 278:
			return seed + `p`, H
		case r < 282:
			return seed + `b`, H
		case r < 285:
			return seed + `g`, H
		case r < 287:
			return seed + `f`, H
		case r < 289:
			return seed + `l`, H
		case r < 290:
			return seed + `w`, H
		default:
			panic("unexpected rand num")
		}
	case `ha`:
		H := 3.468397105262385
		r := rng.IntN(257)
		switch {
		case r < 66:
			return seed + `n`, H
		case r < 119:
			return seed + `r`, H
		case r < 138:
			return seed + `t`, H
		case r < 155:
			return seed + `p`, H
		case r < 171:
			return seed + `s`, H
		case r < 186:
			return seed + `m`, H
		case r < 197:
			return seed + `l`, H
		case r < 207:
			return seed + `b`, H
		case r < 215:
			return seed + `c`, H
		case r < 222:
			return seed + `d`, H
		case r < 229:
			return seed + `k`, H
		case r < 236:
			return seed + `z`, H
		case r < 242:
			return seed + `i`, H
		case r < 245:
			return seed + `f`, H
		case r < 248:
			return seed + `w`, H
		case r < 251:
			return seed + `v`, H
		case r < 254:
			return seed + `u`, H
		case r < 256:
			return seed + `g`, H
		case r < 257:
			return seed + `o`, H
		default:
			panic("unexpected rand num")
		}
	case `ri`:
		H := 3.8044324336764634
		r := rng.IntN(484)
		switch {
		case r < 94:
			return seed + `n`, H
		case r < 146:
			return seed + `c`, H
		case r < 193:
			return seed + `t`, H
		case r < 236:
			return seed + `s`, H
		case r < 269:
			return seed + `e`, H
		case r < 299:
			return seed + `v`, H
		case r < 325:
			return seed + `m`, H
		case r < 349:
			return seed + `d`, H
		case r < 369:
			return seed + `l`, H
		case r < 387:
			return seed + `a`, H
		case r < 405:
			return seed + `g`, H
		case r < 422:
			return seed + `p`, H
		case r < 438:
			return seed + `f`, H
		case r < 454:
			return seed + `b`, H
		case r < 467:
			return seed + `o`, H
		case r < 475:
			return seed + `z`, H
		case r < 480:
			return seed + `u`, H
		case r < 483:
			return seed + `k`, H
		case r < 484:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `o`:
		H := 2.9453946587990396
		r := rng.IntN(246)
		switch {
		case r < 87:
			return seed + `v`, H
		case r < 140:
			return seed + `u`, H
		case r < 167:
			return seed + `b`, H
		case r < 186:
			return seed + `p`, H
		case r < 201:
			return seed + `n`, H
		case r < 212:
			return seed + `c`, H
		case r < 219:
			return seed + `x`, H
		case r < 225:
			return seed + `m`, H
		case r < 229:
			return seed + `a`, H
		case r < 232:
			return seed + `o`, H
		case r < 235:
			return seed + `i`, H
		case r < 238:
			return seed + `l`, H
		case r < 240:
			return seed + `t`, H
		case r < 241:
			return seed + `w`, H
		case r < 242:
			return seed + `s`, H
		case r < 243:
			return seed + `k`, H
		case r < 244:
			return seed + `y`, H
		case r < 245:
			return seed + `z`, H
		case r < 246:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `ti`:
		H := 3.485360725030509
		r := rng.IntN(500)
		switch {
		case r < 118:
			return seed + `n`, H
		case r < 208:
			return seed + `o`, H
		case r < 273:
			return seed + `c`, H
		case r < 313:
			return seed + `v`, H
		case r < 337:
			return seed + `t`, H
		case r < 361:
			return seed + `l`, H
		case r < 384:
			return seed + `f`, H
		case r < 405:
			return seed + `m`, H
		case r < 425:
			return seed + `s`, H
		case r < 437:
			return seed + `g`, H
		case r < 448:
			return seed + `e`, H
		case r < 459:
			return seed + `r`, H
		case r < 470:
			return seed + `p`, H
		case r < 480:
			return seed + `z`, H
		case r < 488:
			return seed + `d`, H
		case r < 494:
			return seed + `a`, H
		case r < 497:
			return seed + `q`, H
		case r < 499:
			return seed + `b`, H
		case r < 500:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `si`:
		H := 3.7124503830288225
		r := rng.IntN(246)
		switch {
		case r < 43:
			return seed + `n`, H
		case r < 75:
			return seed + `o`, H
		case r < 100:
			return seed + `t`, H
		case r < 122:
			return seed + `v`, H
		case r < 143:
			return seed + `d`, H
		case r < 164:
			return seed + `l`, H
		case r < 180:
			return seed + `s`, H
		case r < 191:
			return seed + `z`, H
		case r < 202:
			return seed + `g`, H
		case r < 212:
			return seed + `b`, H
		case r < 221:
			return seed + `m`, H
		case r < 227:
			return seed + `x`, H
		case r < 232:
			return seed + `e`, H
		case r < 237:
			return seed + `c`, H
		case r < 239:
			return seed + `f`, H
		case r < 241:
			return seed + `r`, H
		case r < 243:
			return seed + `a`, H
		case r < 245:
			return seed + `p`, H
		case r < 246:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `la`:
		H := 3.734800815977329
		r := rng.IntN(393)
		switch {
		case r < 68:
			return seed + `t`, H
		case r < 132:
			return seed + `n`, H
		case r < 174:
			return seed + `r`, H
		case r < 212:
			return seed + `s`, H
		case r < 236:
			return seed + `y`, H
		case r < 257:
			return seed + `c`, H
		case r < 277:
			return seed + `m`, H
		case r < 295:
			return seed + `p`, H
		case r < 312:
			return seed + `d`, H
		case r < 326:
			return seed + `i`, H
		case r < 338:
			return seed + `u`, H
		case r < 350:
			return seed + `b`, H
		case r < 361:
			return seed + `g`, H
		case r < 371:
			return seed + `z`, H
		case r < 378:
			return seed + `v`, H
		case r < 384:
			return seed + `w`, H
		case r < 389:
			return seed + `k`, H
		case r < 392:
			return seed + `x`, H
		case r < 393:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `ga`:
		H := 3.4964001580415376
		r := rng.IntN(98)
		switch {
		case r < 18:
			return seed + `t`, H
		case r < 35:
			return seed + `l`, H
		case r < 50:
			return seed + `r`, H
		case r < 62:
			return seed + `n`, H
		case r < 69:
			return seed + `m`, H
		case r < 75:
			return seed + `g`, H
		case r < 79:
			return seed + `i`, H
		case r < 82:
			return seed + `z`, H
		case r < 85:
			return seed + `u`, H
		case r < 88:
			return seed + `b`, H
		case r < 90:
			return seed + `d`, H
		case r < 91:
			return seed + `f`, H
		case r < 92:
			return seed + `w`, H
		case r < 93:
			return seed + `e`, H
		case r < 94:
			return seed + `h`, H
		case r < 95:
			return seed + `s`, H
		case r < 96:
			return seed + `c`, H
		case r < 97:
			return seed + `p`, H
		case r < 98:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `sa`:
		H := 3.7462923073871974
		r := rng.IntN(159)
		switch {
		case r < 30:
			return seed + `l`, H
		case r < 56:
			return seed + `n`, H
		case r < 73:
			return seed + `b`, H
		case r < 84:
			return seed + `t`, H
		case r < 94:
			return seed + `g`, H
		case r < 102:
			return seed + `d`, H
		case r < 110:
			return seed + `v`, H
		case r < 117:
			return seed + `f`, H
		case r < 123:
			return seed + `i`, H
		case r < 129:
			return seed + `r`, H
		case r < 135:
			return seed + `m`, H
		case r < 140:
			return seed + `c`, H
		case r < 145:
			return seed + `u`, H
		case r < 149:
			return seed + `s`, H
		case r < 152:
			return seed + `w`, H
		case r < 155:
			return seed + `p`, H
		case r < 157:
			return seed + `y`, H
		case r < 158:
			return seed + `k`, H
		case r < 159:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `na`:
		H := 3.6919110777928674
		r := rng.IntN(161)
		switch {
		case r < 30:
			return seed + `t`, H
		case r < 53:
			return seed + `l`, H
		case r < 70:
			return seed + `n`, H
		case r < 84:
			return seed + `r`, H
		case r < 96:
			return seed + `m`, H
		case r < 107:
			return seed + `p`, H
		case r < 117:
			return seed + `g`, H
		case r < 126:
			return seed + `b`, H
		case r < 134:
			return seed + `c`, H
		case r < 140:
			return seed + `s`, H
		case r < 144:
			return seed + `d`, H
		case r < 148:
			return seed + `i`, H
		case r < 151:
			return seed + `v`, H
		case r < 154:
			return seed + `u`, H
		case r < 156:
			return seed + `w`, H
		case r < 158:
			return seed + `e`, H
		case r < 159:
			return seed + `f`, H
		case r < 160:
			return seed + `k`, H
		case r < 161:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ad`:
		H := 3.4091611715594885
		r := rng.IntN(149)
		switch {
		case r < 40:
			return seed + `e`, H
		case r < 69:
			return seed + `i`, H
		case r < 87:
			return seed + `d`, H
		case r < 100:
			return seed + `l`, H
		case r < 106:
			return seed + `s`, H
		case r < 112:
			return seed + `y`, H
		case r < 117:
			return seed + `o`, H
		case r < 122:
			return seed + `a`, H
		case r < 125:
			return seed + `n`, H
		case r < 128:
			return seed + `f`, H
		case r < 131:
			return seed + `w`, H
		case r < 134:
			return seed + `r`, H
		case r < 137:
			return seed + `p`, H
		case r < 140:
			return seed + `b`, H
		case r < 142:
			return seed + `c`, H
		case r < 144:
			return seed + `m`, H
		case r < 146:
			return seed + `g`, H
		case r < 148:
			return seed + `u`, H
		case r < 149:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ub`:
		H := 3.6756598041227804
		r := rng.IntN(89)
		switch {
		case r < 21:
			return seed + `b`, H
		case r < 33:
			return seed + `s`, H
		case r < 43:
			return seed + `l`, H
		case r < 51:
			return seed + `t`, H
		case r < 56:
			return seed + `i`, H
		case r < 61:
			return seed + `m`, H
		case r < 65:
			return seed + `d`, H
		case r < 69:
			return seed + `p`, H
		case r < 72:
			return seed + `e`, H
		case r < 75:
			return seed + `a`, H
		case r < 77:
			return seed + `w`, H
		case r < 79:
			return seed + `h`, H
		case r < 81:
			return seed + `r`, H
		case r < 83:
			return seed + `u`, H
		case r < 84:
			return seed + `f`, H
		case r < 85:
			return seed + `y`, H
		case r < 86:
			return seed + `j`, H
		case r < 87:
			return seed + `c`, H
		case r < 88:
			return seed + `z`, H
		case r < 89:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `bo`:
		H := 3.6735456800695463
		r := rng.IntN(153)
		switch {
		case r < 26:
			return seed + `o`, H
		case r < 46:
			return seed + `n`, H
		case r < 64:
			return seed + `a`, H
		case r < 80:
			return seed + `x`, H
		case r < 94:
			return seed + `r`, H
		case r < 106:
			return seed + `u`, H
		case r < 116:
			return seed + `t`, H
		case r < 123:
			return seed + `d`, H
		case r < 130:
			return seed + `b`, H
		case r < 136:
			return seed + `l`, H
		case r < 139:
			return seed + `w`, H
		case r < 142:
			return seed + `s`, H
		case r < 145:
			return seed + `g`, H
		case r < 147:
			return seed + `y`, H
		case r < 148:
			return seed + `f`, H
		case r < 149:
			return seed + `e`, H
		case r < 150:
			return seed + `j`, H
		case r < 151:
			return seed + `i`, H
		case r < 152:
			return seed + `k`, H
		case r < 153:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `ta`:
		H := 3.632864868080895
		r := rng.IntN(324)
		switch {
		case r < 45:
			return seed + `r`, H
		case r < 87:
			return seed + `l`, H
		case r < 126:
			return seed + `b`, H
		case r < 162:
			return seed + `t`, H
		case r < 195:
			return seed + `n`, H
		case r < 224:
			return seed + `i`, H
		case r < 246:
			return seed + `g`, H
		case r < 267:
			return seed + `c`, H
		case r < 279:
			return seed + `s`, H
		case r < 290:
			return seed + `p`, H
		case r < 300:
			return seed + `k`, H
		case r < 310:
			return seed + `m`, H
		case r < 314:
			return seed + `d`, H
		case r < 316:
			return seed + `f`, H
		case r < 318:
			return seed + `y`, H
		case r < 320:
			return seed + `u`, H
		case r < 321:
			return seed + `w`, H
		case r < 322:
			return seed + `e`, H
		case r < 323:
			return seed + `v`, H
		case r < 324:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `di`:
		H := 3.4494359973163755
		r := rng.IntN(288)
		switch {
		case r < 71:
			return seed + `s`, H
		case r < 134:
			return seed + `n`, H
		case r < 158:
			return seed + `a`, H
		case r < 178:
			return seed + `t`, H
		case r < 195:
			return seed + `c`, H
		case r < 211:
			return seed + `v`, H
		case r < 225:
			return seed + `l`, H
		case r < 234:
			return seed + `e`, H
		case r < 243:
			return seed + `m`, H
		case r < 250:
			return seed + `f`, H
		case r < 257:
			return seed + `o`, H
		case r < 264:
			return seed + `r`, H
		case r < 270:
			return seed + `z`, H
		case r < 276:
			return seed + `g`, H
		case r < 280:
			return seed + `p`, H
		case r < 284:
			return seed + `b`, H
		case r < 285:
			return seed + `w`, H
		case r < 286:
			return seed + `d`, H
		case r < 287:
			return seed + `x`, H
		case r < 288:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `se`:
		H := 3.6228884091123996
		r := rng.IntN(199)
		switch {
		case r < 32:
			return seed + `d`, H
		case r < 60:
			return seed + `r`, H
		case r < 86:
			return seed + `l`, H
		case r < 111:
			return seed + `n`, H
		case r < 127:
			return seed + `s`, H
		case r < 142:
			return seed + `c`, H
		case r < 156:
			return seed + `t`, H
		case r < 165:
			return seed + `m`, H
		case r < 171:
			return seed + `a`, H
		case r < 177:
			return seed + `v`, H
		case r < 182:
			return seed + `e`, H
		case r < 186:
			return seed + `p`, H
		case r < 189:
			return seed + `q`, H
		case r < 191:
			return seed + `y`, H
		case r < 193:
			return seed + `i`, H
		case r < 195:
			return seed + `u`, H
		case r < 196:
			return seed + `f`, H
		case r < 197:
			return seed + `w`, H
		case r < 198:
			return seed + `g`, H
		case r < 199:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ur`:
		H := 3.773670379758159
		r := rng.IntN(244)
		switch {
		case r < 63:
			return seed + `e`, H
		case r < 90:
			return seed + `i`, H
		case r < 110:
			return seed + `a`, H
		case r < 125:
			return seed + `s`, H
		case r < 139:
			return seed + `r`, H
		case r < 152:
			return seed + `g`, H
		case r < 162:
			return seed + `n`, H
		case r < 172:
			return seed + `v`, H
		case r < 181:
			return seed + `l`, H
		case r < 190:
			return seed + `b`, H
		case r < 198:
			return seed + `o`, H
		case r < 206:
			return seed + `t`, H
		case r < 214:
			return seed + `p`, H
		case r < 221:
			return seed + `f`, H
		case r < 227:
			return seed + `d`, H
		case r < 233:
			return seed + `c`, H
		case r < 237:
			return seed + `k`, H
		case r < 240:
			return seed + `y`, H
		case r < 242:
			return seed + `m`, H
		case r < 244:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `do`:
		H := 3.6985346528002148
		r := rng.IntN(103)
		switch {
		case r < 15:
			return seed + `r`, H
		case r < 29:
			return seed + `w`, H
		case r < 43:
			return seed + `o`, H
		case r < 56:
			return seed + `m`, H
		case r < 65:
			return seed + `n`, H
		case r < 73:
			return seed + `l`, H
		case r < 80:
			return seed + `c`, H
		case r < 83:
			return seed + `s`, H
		case r < 86:
			return seed + `i`, H
		case r < 89:
			return seed + `g`, H
		case r < 91:
			return seed + `d`, H
		case r < 93:
			return seed + `t`, H
		case r < 95:
			return seed + `z`, H
		case r < 97:
			return seed + `u`, H
		case r < 98:
			return seed + `f`, H
		case r < 99:
			return seed + `e`, H
		case r < 100:
			return seed + `k`, H
		case r < 101:
			return seed + `a`, H
		case r < 102:
			return seed + `v`, H
		case r < 103:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ut`:
		H := 3.8482247103998266
		r := rng.IntN(152)
		switch {
		case r < 27:
			return seed + `e`, H
		case r < 49:
			return seed + `i`, H
		case r < 64:
			return seed + `t`, H
		case r < 76:
			return seed + `s`, H
		case r < 86:
			return seed + `o`, H
		case r < 95:
			return seed + `a`, H
		case r < 103:
			return seed + `h`, H
		case r < 111:
			return seed + `l`, H
		case r < 117:
			return seed + `r`, H
		case r < 123:
			return seed + `b`, H
		case r < 128:
			return seed + `c`, H
		case r < 132:
			return seed + `p`, H
		case r < 136:
			return seed + `m`, H
		case r < 139:
			return seed + `f`, H
		case r < 142:
			return seed + `w`, H
		case r < 145:
			return seed + `d`, H
		case r < 147:
			return seed + `y`, H
		case r < 149:
			return seed + `g`, H
		case r < 151:
			return seed + `u`, H
		case r < 152:
			return seed + `n`, H
		default:
			panic("unexpected rand num")
		}
	case `e`:
		H := 3.440729337470061
		r := rng.IntN(398)
		switch {
		case r < 87:
			return seed + `n`, H
		case r < 169:
			return seed + `x`, H
		case r < 208:
			return seed + `m`, H
		case r < 245:
			return seed + `a`, H
		case r < 280:
			return seed + `l`, H
		case r < 301:
			return seed + `s`, H
		case r < 322:
			return seed + `v`, H
		case r < 334:
			return seed + `c`, H
		case r < 345:
			return seed + `r`, H
		case r < 355:
			return seed + `p`, H
		case r < 364:
			return seed + `d`, H
		case r < 372:
			return seed + `q`, H
		case r < 380:
			return seed + `g`, H
		case r < 387:
			return seed + `t`, H
		case r < 391:
			return seed + `f`, H
		case r < 394:
			return seed + `b`, H
		case r < 395:
			return seed + `e`, H
		case r < 396:
			return seed + `j`, H
		case r < 397:
			return seed + `i`, H
		case r < 398:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `ma`:
		H := 3.24717090849481
		r := rng.IntN(267)
		switch {
		case r < 75:
			return seed + `n`, H
		case r < 124:
			return seed + `t`, H
		case r < 166:
			return seed + `r`, H
		case r < 187:
			return seed + `g`, H
		case r < 203:
			return seed + `s`, H
		case r < 215:
			return seed + `l`, H
		case r < 225:
			return seed + `k`, H
		case r < 235:
			return seed + `c`, H
		case r < 239:
			return seed + `d`, H
		case r < 243:
			return seed + `i`, H
		case r < 247:
			return seed + `y`, H
		case r < 251:
			return seed + `j`, H
		case r < 255:
			return seed + `m`, H
		case r < 258:
			return seed + `b`, H
		case r < 260:
			return seed + `p`, H
		case r < 262:
			return seed + `z`, H
		case r < 264:
			return seed + `x`, H
		case r < 265:
			return seed + `h`, H
		case r < 266:
			return seed + `v`, H
		case r < 267:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `is`:
		H := 3.425586749295823
		r := rng.IntN(339)
		switch {
		case r < 84:
			return seed + `t`, H
		case r < 148:
			return seed + `h`, H
		case r < 178:
			return seed + `m`, H
		case r < 207:
			return seed + `e`, H
		case r < 230:
			return seed + `i`, H
		case r < 252:
			return seed + `p`, H
		case r < 269:
			return seed + `o`, H
		case r < 285:
			return seed + `s`, H
		case r < 298:
			return seed + `c`, H
		case r < 307:
			return seed + `a`, H
		case r < 315:
			return seed + `k`, H
		case r < 323:
			return seed + `l`, H
		case r < 327:
			return seed + `b`, H
		case r < 330:
			return seed + `f`, H
		case r < 332:
			return seed + `d`, H
		case r < 334:
			return seed + `y`, H
		case r < 336:
			return seed + `r`, H
		case r < 337:
			return seed + `w`, H
		case r < 338:
			return seed + `j`, H
		case r < 339:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `nd`:
		H := 3.5314039953405065
		r := rng.IntN(257)
		switch {
		case r < 75:
			return seed + `e`, H
		case r < 103:
			return seed + `i`, H
		case r < 130:
			return seed + `l`, H
		case r < 151:
			return seed + `a`, H
		case r < 170:
			return seed + `o`, H
		case r < 183:
			return seed + `s`, H
		case r < 194:
			return seed + `r`, H
		case r < 204:
			return seed + `u`, H
		case r < 213:
			return seed + `b`, H
		case r < 219:
			return seed + `f`, H
		case r < 225:
			return seed + `w`, H
		case r < 231:
			return seed + `p`, H
		case r < 237:
			return seed + `m`, H
		case r < 242:
			return seed + `n`, H
		case r < 247:
			return seed + `y`, H
		case r < 252:
			return seed + `c`, H
		case r < 254:
			return seed + `g`, H
		case r < 255:
			return seed + `d`, H
		case r < 256:
			return seed + `h`, H
		case r < 257:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `ee`:
		H := 3.7833598092940597
		r := rng.IntN(126)
		switch {
		case r < 20:
			return seed + `n`, H
		case r < 40:
			return seed + `d`, H
		case r < 53:
			return seed + `p`, H
		case r < 64:
			return seed + `z`, H
		case r < 74:
			return seed + `t`, H
		case r < 83:
			return seed + `l`, H
		case r < 91:
			return seed + `r`, H
		case r < 97:
			return seed + `m`, H
		case r < 102:
			return seed + `b`, H
		case r < 105:
			return seed + `f`, H
		case r < 108:
			return seed + `w`, H
		case r < 111:
			return seed + `s`, H
		case r < 114:
			return seed + `i`, H
		case r < 117:
			return seed + `k`, H
		case r < 120:
			return seed + `c`, H
		case r < 122:
			return seed + `a`, H
		case r < 123:
			return seed + `h`, H
		case r < 124:
			return seed + `x`, H
		case r < 125:
			return seed + `v`, H
		case r < 126:
			return seed + `g`, H
		default:
			panic("unexpected rand num")
		}
	case `lo`:
		H := 3.9481220227019564
		r := rng.IntN(203)
		switch {
		case r < 26:
			return seed + `g`, H
		case r < 48:
			return seed + `r`, H
		case r < 67:
			return seed + `n`, H
		case r < 84:
			return seed + `a`, H
		case r < 100:
			return seed + `t`, H
		case r < 115:
			return seed + `w`, H
		case r < 128:
			return seed + `o`, H
		case r < 141:
			return seed + `c`, H
		case r < 152:
			return seed + `p`, H
		case r < 162:
			return seed + `s`, H
		case r < 172:
			return seed + `u`, H
		case r < 179:
			return seed + `v`, H
		case r < 185:
			return seed + `y`, H
		case r < 190:
			return seed + `b`, H
		case r < 194:
			return seed + `d`, H
		case r < 196:
			return seed + `i`, H
		case r < 198:
			return seed + `q`, H
		case r < 200:
			return seed + `m`, H
		case r < 201:
			return seed + `f`, H
		case r < 202:
			return seed + `e`, H
		case r < 203:
			return seed + `h`, H
		default:
			panic("unexpected rand num")
		}
	case `a`:
		H := 3.8126739941897543
		r := rng.IntN(407)
		switch {
		case r < 74:
			return seed + `n`, H
		case r < 113:
			return seed + `m`, H
		case r < 148:
			return seed + `r`, H
		case r < 181:
			return seed + `l`, H
		case r < 211:
			return seed + `c`, H
		case r < 240:
			return seed + `p`, H
		case r < 264:
			return seed + `b`, H
		case r < 287:
			return seed + `f`, H
		case r < 310:
			return seed + `s`, H
		case r < 332:
			return seed + `t`, H
		case r < 352:
			return seed + `g`, H
		case r < 371:
			return seed + `u`, H
		case r < 384:
			return seed + `v`, H
		case r < 393:
			return seed + `w`, H
		case r < 397:
			return seed + `e`, H
		case r < 400:
			return seed + `i`, H
		case r < 402:
			return seed + `h`, H
		case r < 404:
			return seed + `q`, H
		case r < 405:
			return seed + `o`, H
		case r < 406:
			return seed + `j`, H
		case r < 407:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `de`:
		H := 3.6173342588664883
		r := rng.IntN(390)
		switch {
		case r < 97:
			return seed + `r`, H
		case r < 138:
			return seed + `d`, H
		case r < 177:
			return seed + `n`, H
		case r < 211:
			return seed + `f`, H
		case r < 241:
			return seed + `c`, H
		case r < 267:
			return seed + `l`, H
		case r < 288:
			return seed + `s`, H
		case r < 305:
			return seed + `p`, H
		case r < 321:
			return seed + `a`, H
		case r < 336:
			return seed + `v`, H
		case r < 350:
			return seed + `t`, H
		case r < 361:
			return seed + `m`, H
		case r < 370:
			return seed + `b`, H
		case r < 376:
			return seed + `e`, H
		case r < 380:
			return seed + `o`, H
		case r < 384:
			return seed + `g`, H
		case r < 386:
			return seed + `x`, H
		case r < 387:
			return seed + `h`, H
		case r < 388:
			return seed + `j`, H
		case r < 389:
			return seed + `i`, H
		case r < 390:
			return seed + `u`, H
		default:
			panic("unexpected rand num")
		}
	case `in`:
		H := 2.1017950678134936
		r := rng.IntN(1050)
		switch {
		case r < 665:
			return seed + `g`, H
		case r < 813:
			return seed + `e`, H
		case r < 855:
			return seed + `t`, H
		case r < 888:
			return seed + `i`, H
		case r < 918:
			return seed + `k`, H
		case r < 944:
			return seed + `a`, H
		case r < 969:
			return seed + `d`, H
		case r < 984:
			return seed + `n`, H
		case r < 998:
			return seed + `s`, H
		case r < 1012:
			return seed + `c`, H
		case r < 1021:
			return seed + `o`, H
		case r < 1028:
			return seed + `y`, H
		case r < 1035:
			return seed + `l`, H
		case r < 1039:
			return seed + `f`, H
		case r < 1042:
			return seed + `u`, H
		case r < 1044:
			return seed + `j`, H
		case r < 1046:
			return seed + `x`, H
		case r < 1047:
			return seed + `w`, H
		case r < 1048:
			return seed + `h`, H
		case r < 1049:
			return seed + `p`, H
		case r < 1050:
			return seed + `v`, H
		default:
			panic("unexpected rand num")
		}
	case `li`:
		H := 3.546418837133385
		r := rng.IntN(475)
		switch {
		case r < 157:
			return seed + `n`, H
		case r < 200:
			return seed + `c`, H
		case r < 235:
			return seed + `t`, H
		case r < 264:
			return seed + `k`, H
		case r < 293:
			return seed + `g`, H
		case r < 320:
			return seed + `s`, H
		case r < 341:
			return seed + `a`, H
		case r < 359:
			return seed + `f`, H
		case r < 377:
			return seed + `m`, H
		case r < 393:
			return seed + `v`, H
		case r < 408:
			return seed + `z`, H
		case r < 422:
			return seed + `e`, H
		case r < 433:
			return seed + `d`, H
		case r < 443:
			return seed + `o`, H
		case r < 453:
			return seed + `p`, H
		case r < 462:
			return seed + `b`, H
		case r < 466:
			return seed + `q`, H
		case r < 469:
			return seed + `r`, H
		case r < 472:
			return seed + `l`, H
		case r < 474:
			return seed + `u`, H
		case r < 475:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ne`:
		H := 3.0862784397322383
		r := rng.IntN(326)
		switch {
		case r < 133:
			return seed + `s`, H
		case r < 178:
			return seed + `r`, H
		case r < 217:
			return seed + `d`, H
		case r < 232:
			return seed + `t`, H
		case r < 247:
			return seed + `a`, H
		case r < 257:
			return seed + `n`, H
		case r < 267:
			return seed + `l`, H
		case r < 275:
			return seed + `w`, H
		case r < 283:
			return seed + `g`, H
		case r < 290:
			return seed + `e`, H
		case r < 297:
			return seed + `u`, H
		case r < 303:
			return seed + `y`, H
		case r < 308:
			return seed + `m`, H
		case r < 313:
			return seed + `x`, H
		case r < 316:
			return seed + `c`, H
		case r < 319:
			return seed + `v`, H
		case r < 322:
			return seed + `b`, H
		case r < 323:
			return seed + `o`, H
		case r < 324:
			return seed + `h`, H
		case r < 325:
			return seed + `q`, H
		case r < 326:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `on`:
		H := 3.7481496950737694
		r := rng.IntN(302)
		switch {
		case r < 59:
			return seed + `e`, H
		case r < 87:
			return seed + `i`, H
		case r < 115:
			return seed + `s`, H
		case r < 143:
			return seed + `g`, H
		case r < 166:
			return seed + `d`, H
		case r < 188:
			return seed + `t`, H
		case r < 208:
			return seed + `f`, H
		case r < 227:
			return seed + `a`, H
		case r < 244:
			return seed + `c`, H
		case r < 260:
			return seed + `o`, H
		case r < 275:
			return seed + `y`, H
		case r < 281:
			return seed + `l`, H
		case r < 285:
			return seed + `u`, H
		case r < 288:
			return seed + `n`, H
		case r < 291:
			return seed + `j`, H
		case r < 294:
			return seed + `v`, H
		case r < 296:
			return seed + `w`, H
		case r < 298:
			return seed + `z`, H
		case r < 300:
			return seed + `b`, H
		case r < 301:
			return seed + `k`, H
		case r < 302:
			return seed + `r`, H
		default:
			panic("unexpected rand num")
		}
	case `le`:
		H := 3.601021056216224
		r := rng.IntN(318)
		switch {
		case r < 69:
			return seed + `s`, H
		case r < 115:
			return seed + `d`, H
		case r < 153:
			return seed + `r`, H
		case r < 182:
			return seed + `t`, H
		case r < 204:
			return seed + `n`, H
		case r < 225:
			return seed + `c`, H
		case r < 243:
			return seed + `a`, H
		case r < 260:
			return seed + `g`, H
		case r < 273:
			return seed + `v`, H
		case r < 283:
			return seed + `e`, H
		case r < 290:
			return seed + `m`, H
		case r < 296:
			return seed + `x`, H
		case r < 301:
			return seed + `y`, H
		case r < 305:
			return seed + `p`, H
		case r < 308:
			return seed + `f`, H
		case r < 311:
			return seed + `u`, H
		case r < 313:
			return seed + `b`, H
		case r < 314:
			return seed + `w`, H
		case r < 315:
			return seed + `o`, H
		case r < 316:
			return seed + `h`, H
		case r < 317:
			return seed + `i`, H
		case r < 318:
			return seed + `l`, H
		default:
			panic("unexpected rand num")
		}
	case `al`:
		H := 3.693547477659404
		r := rng.IntN(250)
		switch {
		case r < 55:
			return seed + `l`, H
		case r < 102:
			return seed + `i`, H
		case r < 123:
			return seed + `a`, H
		case r < 140:
			return seed + `e`, H
		case r < 156:
			return seed + `o`, H
		case r < 172:
			return seed + `t`, H
		case r < 183:
			return seed + `u`, H
		case r < 192:
			return seed + `k`, H
		case r < 201:
			return seed + `m`, H
		case r < 208:
			return seed + `s`, H
		case r < 214:
			return seed + `y`, H
		case r < 220:
			return seed + `c`, H
		case r < 226:
			return seed + `v`, H
		case r < 230:
			return seed + `f`, H
		case r < 234:
			return seed + `d`, H
		case r < 238:
			return seed + `r`, H
		case r < 242:
			return seed + `p`, H
		case r < 244:
			return seed + `n`, H
		case r < 246:
			return seed + `g`, H
		case r < 248:
			return seed + `b`, H
		case r < 249:
			return seed + `w`, H
		case r < 250:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `ra`:
		H := 3.9497641727639463
		r := rng.IntN(471)
		switch {
		case r < 69:
			return seed + `n`, H
		case r < 125:
			return seed + `t`, H
		case r < 161:
			return seed + `c`, H
		case r < 195:
			return seed + `i`, H
		case r < 228:
			return seed + `d`, H
		case r < 260:
			return seed + `m`, H
		case r < 291:
			return seed + `l`, H
		case r < 321:
			return seed + `g`, H
		case r < 346:
			return seed + `v`, H
		case r < 369:
			return seed + `p`, H
		case r < 391:
			return seed + `s`, H
		case r < 410:
			return seed + `b`, H
		case r < 425:
			return seed + `y`, H
		case r < 439:
			return seed + `f`, H
		case r < 448:
			return seed + `r`, H
		case r < 456:
			return seed + `w`, H
		case r < 462:
			return seed + `z`, H
		case r < 466:
			return seed + `k`, H
		case r < 468:
			return seed + `u`, H
		case r < 469:
			return seed + `e`, H
		case r < 470:
			return seed + `o`, H
		case r < 471:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `ar`:
		H := 4.022304287988487
		r := rng.IntN(513)
		switch {
		case r < 79:
			return seed + `d`, H
		case r < 132:
			return seed + `t`, H
		case r < 174:
			return seed + `i`, H
		case r < 215:
			return seed + `m`, H
		case r < 251:
			return seed + `e`, H
		case r < 284:
			return seed + `r`, H
		case r < 316:
			return seed + `a`, H
		case r < 346:
			return seed + `y`, H
		case r < 371:
			return seed + `g`, H
		case r < 392:
			return seed + `l`, H
		case r < 409:
			return seed + `n`, H
		case r < 426:
			return seed + `c`, H
		case r < 442:
			return seed + `b`, H
		case r < 457:
			return seed + `k`, H
		case r < 470:
			return seed + `p`, H
		case r < 482:
			return seed + `o`, H
		case r < 493:
			return seed + `s`, H
		case r < 501:
			return seed + `f`, H
		case r < 508:
			return seed + `v`, H
		case r < 511:
			return seed + `w`, H
		case r < 512:
			return seed + `h`, H
		case r < 513:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case `to`:
		H := 2.943211643112022
		r := rng.IntN(200)
		switch {
		case r < 82:
			return seed + `r`, H
		case r < 116:
			return seed + `n`, H
		case r < 136:
			return seed + `p`, H
		case r < 154:
			return seed + `m`, H
		case r < 164:
			return seed + `o`, H
		case r < 170:
			return seed + `u`, H
		case r < 175:
			return seed + `c`, H
		case r < 179:
			return seed + `w`, H
		case r < 183:
			return seed + `l`, H
		case r < 185:
			return seed + `d`, H
		case r < 187:
			return seed + `i`, H
		case r < 189:
			return seed + `t`, H
		case r < 191:
			return seed + `x`, H
		case r < 192:
			return seed + `f`, H
		case r < 193:
			return seed + `e`, H
		case r < 194:
			return seed + `s`, H
		case r < 195:
			return seed + `k`, H
		case r < 196:
			return seed + `a`, H
		case r < 197:
			return seed + `z`, H
		case r < 198:
			return seed + `g`, H
		case r < 199:
			return seed + `v`, H
		case r < 200:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `en`:
		H := 3.239836470529042
		r := rng.IntN(441)
		switch {
		case r < 148:
			return seed + `t`, H
		case r < 217:
			return seed + `d`, H
		case r < 268:
			return seed + `c`, H
		case r < 306:
			return seed + `e`, H
		case r < 331:
			return seed + `s`, H
		case r < 351:
			return seed + `g`, H
		case r < 368:
			return seed + `a`, H
		case r < 384:
			return seed + `i`, H
		case r < 392:
			return seed + `n`, H
		case r < 400:
			return seed + `v`, H
		case r < 407:
			return seed + `o`, H
		case r < 414:
			return seed + `u`, H
		case r < 419:
			return seed + `j`, H
		case r < 424:
			return seed + `l`, H
		case r < 428:
			return seed + `r`, H
		case r < 431:
			return seed + `f`, H
		case r < 434:
			return seed + `z`, H
		case r < 436:
			return seed + `h`, H
		case r < 438:
			return seed + `y`, H
		case r < 439:
			return seed + `q`, H
		case r < 440:
			return seed + `k`, H
		case r < 441:
			return seed + `p`, H
		default:
			panic("unexpected rand num")
		}
	case `an`:
		H := 3.4086078047375628
		r := rng.IntN(553)
		switch {
		case r < 125:
			return seed + `d`, H
		case r < 238:
			return seed + `t`, H
		case r < 294:
			return seed + `c`, H
		case r < 347:
			return seed + `g`, H
		case r < 386:
			return seed + `k`, H
		case r < 422:
			return seed + `i`, H
		case r < 454:
			return seed + `n`, H
		case r < 473:
			return seed + `a`, H
		case r < 489:
			return seed + `y`, H
		case r < 503:
			return seed + `e`, H
		case r < 516:
			return seed + `o`, H
		case r < 528:
			return seed + `s`, H
		case r < 535:
			return seed + `l`, H
		case r < 541:
			return seed + `h`, H
		case r < 544:
			return seed + `u`, H
		case r < 546:
			return seed + `q`, H
		case r < 547:
			return seed + `f`, H
		case r < 548:
			return seed + `j`, H
		case r < 549:
			return seed + `p`, H
		case r < 550:
			return seed + `m`, H
		case r < 551:
			return seed + `z`, H
		case r < 552:
			return seed + `v`, H
		case r < 553:
			return seed + `b`, H
		default:
			panic("unexpected rand num")
		}
	case `co`:
		H := 3.4735036327365343
		r := rng.IntN(349)
		switch {
		case r < 93:
			return seed + `n`, H
		case r < 145:
			return seed + `m`, H
		case r < 196:
			return seed + `r`, H
		case r < 224:
			return seed + `u`, H
		case r < 252:
			return seed + `l`, H
		case r < 268:
			return seed + `a`, H
		case r < 281:
			return seed + `p`, H
		case r < 291:
			return seed + `s`, H
		case r < 299:
			return seed + `o`, H
		case r < 307:
			return seed + `t`, H
		case r < 315:
			return seed + `v`, H
		case r < 320:
			return seed + `d`, H
		case r < 324:
			return seed + `g`, H
		case r < 328:
			return seed + `b`, H
		case r < 331:
			return seed + `f`, H
		case r < 334:
			return seed + `e`, H
		case r < 337:
			return seed + `h`, H
		case r < 340:
			return seed + `i`, H
		case r < 343:
			return seed + `z`, H
		case r < 345:
			return seed + `y`, H
		case r < 347:
			return seed + `c`, H
		case r < 348:
			return seed + `w`, H
		case r < 349:
			return seed + `k`, H
		default:
			panic("unexpected rand num")
		}
	case `er`:
		H := 4.231365586934958
		r := rng.IntN(466)
		switch {
		case r < 48:
			return seed + `s`, H
		case r < 96:
			return seed + `i`, H
		case r < 134:
			return seed + `e`, H
		case r < 172:
			return seed + `a`, H
		case r < 205:
			return seed + `t`, H
		case r < 233:
			return seed + `y`, H
		case r < 256:
			return seed + `o`, H
		case r < 279:
			return seed + `r`, H
		case r < 301:
			return seed + `b`, H
		case r < 319:
			return seed + `n`, H
		case r < 337:
			return seed + `m`, H
		case r < 355:
			return seed + `g`, H
		case r < 372:
			return seed + `c`, H
		case r < 389:
			return seed + `v`, H
		case r < 405:
			return seed + `l`, H
		case r < 418:
			return seed + `p`, H
		case r < 430:
			return seed + `f`, H
		case r < 440:
			return seed + `d`, H
		case r < 449:
			return seed + `h`, H
		case r < 454:
			return seed + `u`, H
		case r < 458:
			return seed + `w`, H
		case r < 462:
			return seed + `k`, H
		case r < 466:
			return seed + `j`, H
		default:
			panic("unexpected rand num")
		}
	case `ro`:
		H := 4.143180722133236
		r := rng.IntN(388)
		switch {
		case r < 45:
			return seed + `u`, H
		case r < 84:
			return seed + `n`, H
		case r < 115:
			return seed + `w`, H
		case r < 146:
			return seed + `o`, H
		case r < 172:
			return seed + `p`, H
		case r < 196:
			return seed + `v`, H
		case r < 218:
			return seed + `c`, H
		case r < 239:
			return seed + `s`, H
		case r < 259:
			return seed + `a`, H
		case r < 278:
			return seed + `l`, H
		case r < 296:
			return seed + `b`, H
		case r < 313:
			return seed + `t`, H
		case r < 329:
			return seed + `m`, H
		case r < 343:
			return seed + `g`, H
		case r < 353:
			return seed + `d`, H
		case r < 361:
			return seed + `f`, H
		case r < 369:
			return seed + `i`, H
		case r < 374:
			return seed + `r`, H
		case r < 379:
			return seed + `x`, H
		case r < 383:
			return seed + `k`, H
		case r < 386:
			return seed + `z`, H
		case r < 387:
			return seed + `j`, H
		case r < 388:
			return seed + `y`, H
		default:
			panic("unexpected rand num")
		}
	case `un`:
		H := 4.013869088306411
		r := rng.IntN(583)
		switch {
		case r < 92:
			return seed + `d`, H
		case r < 160:
			return seed + `c`, H
		case r < 227:
			return seed + `t`, H
		case r < 281:
			return seed + `s`, H
		case r < 312:
			return seed + `i`, H
		case r < 338:
			return seed + `l`, H
		case r < 362:
			return seed + `w`, H
		case r < 385:
			return seed + `g`, H
		case r < 407:
			return seed + `e`, H
		case r < 429:
			return seed + `r`, H
		case r < 449:
			return seed + `f`, H
		case r < 469:
			return seed + `b`, H
		case r < 487:
			return seed + `k`, H
		case r < 505:
			return seed + `a`, H
		case r < 523:
			return seed + `p`, H
		case r < 539:
			return seed + `m`, H
		case r < 553:
			return seed + `n`, H
		case r < 566:
			return seed + `h`, H
		case r < 575:
			return seed + `v`, H
		case r < 578:
			return seed + `u`, H
		case r < 580:
			return seed + `o`, H
		case r < 581:
			return seed + `j`, H
		case r < 582:
			return seed + `q`, H
		case r < 583:
			return seed + `z`, H
		default:
			panic("unexpected rand num")
		}
	case `re`:
		H := 4.070035666531293
		r := rng.IntN(744)
		switch {
		case r < 88:
			return seed + `a`, H
		case r < 173:
			return seed + `s`, H
		case r < 232:
			return seed + `c`, H
		case r < 290:
			return seed + `e`, H
		case r < 346:
			return seed + `d`, H
		case r < 397:
			return seed + `t`, H
		case r < 441:
			return seed + `p`, H
		case r < 484:
			return seed + `n`, H
		case r < 527:
			return seed + `l`, H
		case r < 567:
			return seed + `f`, H
		case r < 600:
			return seed + `v`, H
		case r < 631:
			return seed + `m`, H
		case r < 653:
			return seed + `g`, H
		case r < 674:
			return seed + `w`, H
		case r < 694:
			return seed + `r`, H
		case r < 710:
			return seed + `b`, H
		case r < 718:
			return seed + `o`, H
		case r < 724:
			return seed + `h`, H
		case r < 729:
			return seed + `q`, H
		case r < 733:
			return seed + `i`, H
		case r < 737:
			return seed + `u`, H
		case r < 740:
			return seed + `j`, H
		case r < 742:
			return seed + `k`, H
		case r < 743:
			return seed + `y`, H
		case r < 744:
			return seed + `x`, H
		default:
			panic("unexpected rand num")
		}
	case ``:
		H := 4.2128960043149775
		r := rng.IntN(7776)
		switch {
		case r < 1087:
			return seed + `s`, H
		case r < 1838:
			return seed + `c`, H
		case r < 2416:
			return seed + `p`, H
		case r < 2963:
			return seed + `d`, H
		case r < 3476:
			return seed + `r`, H
		case r < 3952:
			return seed + `u`, H
		case r < 4359:
			return seed + `a`, H
		case r < 4757:
			return seed + `e`, H
		case r < 5119:
			return seed + `b`, H
		case r < 5442:
			return seed + `t`, H
		case r < 5750:
			return seed + `f`, H
		case r < 6054:
			return seed + `g`, H
		case r < 6349:
			return seed + `m`, H
		case r < 6598:
			return seed + `h`, H
		case r < 6844:
			return seed + `o`, H
		case r < 7038:
			return seed + `l`, H
		case r < 7193:
			return seed + `w`, H
		case r < 7325:
			return seed + `v`, H
		case r < 7440:
			return seed + `i`, H
		case r < 7537:
			return seed + `n`, H
		case r < 7633:
			return seed + `j`, H
		case r < 7688:
			return seed + `k`, H
		case r < 7725:
			return seed + `q`, H
		case r < 7752:
			return seed + `y`, H
		case r < 7774:
			return seed + `z`, H
		case r < 7776:
			return seed + `x`, H
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
