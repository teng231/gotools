package stringfolding

import (
	"log"
	"math/rand"
	"testing"
)

func Test_Genint(t *testing.T) {
	ins := []string{
		"s",
		"2sfs",
		"aerf@asf34",
		"IN9IJA7UK-eWgMeiWDQAfGV1gTqScpHTEQIcqYIyXiebQw2ukC4t6C3ciNK7zj2N",
		"7TSpmkmjnZTqddE525VMteynA7tBBMMmvqVLjafWA4p3RZbKckNoO5I_D6Xj01D8",
		"-OQ9ujOuhaDNlbkwh1Vt2epclcfktREwy-RDQxLa1XBjJhKhUfpFt8AA8Pt7_aDT",
		"iQJq7lpbb9AjdFR1lEpfu_tokDsqX8vTCGSCU3YejvAdkHhqLbPOeaSa8O235fuu",
		"3jx4jbpbUJ72bDFeZuSt-M1pUeMUCR0CBfozuYKs1jkTuthyi9CDG3vdkWR3mLWA",
		"kCsgjUYOTBCCRXvpkUNfBQ",
		"d2SF02L4ztIWXIhCBHmyS8OoydNjOCe2SnH-seYiEM7d5ipVWWfic8_3b7kzH3xW",
		"VUNaSAlFAdjL2j4090QEqjI5RE8OPo7W3Vaa2TAN4LxW9lj9a11qNGGpEE38NwOP",
		"2xemeeuXhjTU4E1QOBHAYBFAcdf45-lauAxLR9bSDgeCDcz-WIHbcgOOT0_u3qih",
		"om7uKXGejg2xAfmVyeh1ZV1B5uAs6q1qQH8km2LQA4vxRMiXJ4eTCLSAsqd0vzSi",
		"9nestG3QF4gCTaV7U516LQ",
		"tTKMukrv2LkixHAUHKud6g",
		"VH93Xv86JsX1UbiU8Xml8g",
		"4O6G3hBBEuDGlhiBlfYdQA",
		"cJJLkWHDuWIBk38g_LBt0w",
		"5WpBp_RcnP3TyZqIbc2Vkg",
		"ohTpTYMLYThedtuVjdbJwA",
		"W415H1DQFlEUmWh6Hw9xDA",
		"Byx5GlF0148FB6gHnVieYQ",
		"LOU0A6gEU0DbDVXhICzQsA",
		"64pbhgbqgDt5hYZnCuaMYA",
		"qyq-22pVll5GyJgFPChvlw",
		"BDAkygBkUwkj7kdBohmKZQ",
		"7Lc1VYMmu7XfS_Qw2lnTrw",
		"S_joC_HcL_ifQjV8vKty3w",
		"Fv830CiNL7cixHAUHKud6g",
		"nxEsSNiGyoSvmf5JoX1SNg",
		"tALFcYesAh9TB_pF5HQGoW3Ey-KGbJqavN_ksd1sB132l98udsUBWex6z9mhdEml",
		"_juNVLSDLldmcKUPvLsazp26MqdUv0lYNblEgzgzHY07QR3XzmEMgmi3IBYdkcH3",
		"daE9eOirxdZtWfiPulU1Vg",
		"D2UpsmOeAfMN0ylv2-3Kkw",
		"k5u_ATe7DpFYUZR26HqwYA",
		"d5TivpOOD_FC69bb_RWH-Q",
		"Wd5n6oiYzAK7q4Lm1JFGhA",
		"NyPkQPvLrujVIf0aGjALP1bNo4UMvL5Ee07uKCqJdwHGFZgNQZd2yLC2cO6Z-4ax",
		"gSb120WPOoP9JYsidHV1mA7R5JuiciY2x2-ktrlhAN5TWfy_9FQmLCKlVSO049qq",
		"IDACWQzM32l7xDzmrLX_bA",
		"O9LA2m-dHJsmDUE9yYqumg",
		"PuxN9nMR_7MGkAj1hDhiCg",
		"HiEnWcBcBXfAvx7ibvtUw_rgJ1AFEuDoDZZI4xsZVLrfjHv4WrJH-ZM2se-BlzHH",
		"GKZz96WIdfABlWUBBNlajQ",
		"_UphevgFBUjnm46A1W83zg",
		"Dk_fTWD1LZqqpfDACAh38g",
		"WC1Rm5PLCTsCqur-Bbq2-4nOdzhBOFtgg7TLZLfDNIMzfWIxMBskzGKU7ITrx1oG",
		"yqm2uIAWqA816oLuEsLANTGio0ZLIRYHxsmXSetFzhoag2eKBo843FpO46x-Cvj2",
		"-bxnuY4howSpwDG0psAD2Q",
		"4UcPjC4oKXUK5QTZTFLFBg",
		"OHsaxL8ZKMqfiuUNRjzP_w",
		"j9vIFGqyqUDnVVyTM707eg",
		"HFgPqRZyAN0CVAItM0aZkQ",
		"Hw2dBbbLCuRBQh1h6BPZfA",
		"tlC25TGYLhK2ZwmAfItZ0g",
		"IrlrkKiDN9Eja0sXWrriDw",
		"HShopNERzqeUPLOtNDMLcw",
		"SaHWorddjJRp86Ja76E8LA",
		"c0CXTa_KMi6Tzd8tIIuGGA",
		"LL2aj20IjUZsD4wUeZ0hSSLmEhaZXEyyanwAJI-6b_c0lAXVIiPIpc7Z0V7bFZHX",
		"Qaagix9goxc3VuMGuk74jHpz7QZ6jACdQQ2Idk4xrAIWu5xo88RbWdNnAsvnfW75",
		"fGSNDAUrk6smNb2Sd27InUY46Kj_SVdZ9wAq0qiElckqNU18s_rPt9iP36gGIfsm",
	}
	for i := 0; i < len(ins); i++ {
		log.Print(ins[i], ":", Generate(ins[i], 1000000))
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestGenDuplicate(t *testing.T) {
	m := map[int64]string{}
	checkDuplicate := map[string]bool{}
	var item string
	for i := 0; i < 5e7; i++ {
		for {
			x := RandStringRunes(12)
			if _, has := checkDuplicate[x]; has {
				log.Print("duplicate")
				continue
			}
			checkDuplicate[x] = true
			item = x
			break
		}
		num := Generate(item, 1e8)
		m[num] = item
	}
	log.Print(len(m), len(checkDuplicate), float32(len(m))/float32(5e7))
}
