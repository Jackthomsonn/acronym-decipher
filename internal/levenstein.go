package internal

import "math"

var peq [0x10000]uint32

func mb(a, b string) int {
	n := len(a)
	m := len(b)
	lst := 1 << uint(n-1)
	pv := uint32(0)
	mv := uint32(0)
	sc := n
	i := n
	for i > 0 {
		i--
		peq[a[i]] |= 1 << uint(i)
	}
	for i = 0; i < m; i++ {
		eq := peq[b[i]]
		xv := eq | mv
		eq |= ((eq & pv) + pv) ^ pv
		mv |= ^(eq | pv)
		pv &= eq
		if mv&uint32(lst) != 0 {
			sc++
		}
		if pv&uint32(lst) != 0 {
			sc--
		}
		mv = (mv << 1) | 1
		pv = (pv << 1) | ^(xv | mv)
		mv &= xv
	}
	i = n
	for i > 0 {
		i--
		peq[a[i]] = 0
	}
	return sc
}

func mx(b, a string) int {
	n := len(a)
	m := len(b)
	var mhc []uint32
	var phc []uint32
	hsize := int(math.Ceil(float64(n) / 32))
	vsize := int(math.Ceil(float64(m) / 32))
	for i := 0; i < hsize; i++ {
		phc = append(phc, ^uint32(0))
		mhc = append(mhc, uint32(0))
	}
	j := 0
	for j < vsize-1 {
		mv := uint32(0)
		pv := ^uint32(0)
		start := j * 32
		vlen := int(math.Min(32, float64(m))) + start
		for k := start; k < vlen; k++ {
			peq[b[k]] |= 1 << uint(k)
		}
		for i := 0; i < n; i++ {
			eq := peq[a[i]]
			pb := (phc[i/32] >> uint(i)) & 1
			mb := (mhc[i/32] >> uint(i)) & 1
			xv := eq | mv
			xh := ((((eq | mb) & pv) + pv) ^ pv) | eq | mb
			ph := mv | ^(xh | pv)
			mh := pv & xh
			if ((ph >> 31) ^ pb) != 0 {
				phc[i/32] ^= 1 << uint(i)
			}
			if ((mh >> 31) ^ mb) != 0 {
				mhc[i/32] ^= 1 << uint(i)
			}
			ph = (ph << 1) | pb
			mh = (mh << 1) | mb
			pv = mh | ^(xv | ph)
			mv = ph & xv
		}
		for k := start; k < vlen; k++ {
			peq[b[k]] = 0
		}
		j++
	}
	mv := uint32(0)
	pv := ^uint32(0)
	start := j * 32
	vlen := int(math.Min(32, float64(m-start))) + start
	for k := start; k < vlen; k++ {
		peq[b[k]] |= 1 << uint(k)
	}
	score := m
	for i := 0; i < n; i++ {
		eq := peq[a[i]]
		pb := (phc[i/32] >> uint(i)) & 1
		mb := (mhc[i/32] >> uint(i)) & 1
		xv := eq | mv
		xh := ((((eq | mb) & pv) + pv) ^ pv) | eq | mb
		ph := mv | ^(xh | pv)
		mh := pv & xh
		score += int((ph >> (m - 1)) & 1)
		score -= int((mh >> (m - 1)) & 1)
		if ((ph >> 31) ^ pb) != 0 {
			phc[i/32] ^= 1 << uint(i)
		}
		if ((mh >> 31) ^ mb) != 0 {
			mhc[i/32] ^= 1 << uint(i)
		}
		ph = (ph << 1) | pb
		mh = (mh << 1) | mb
		pv = mh | ^(xv | ph)
		mv = ph & xv
	}
	for k := start; k < vlen; k++ {
		peq[b[k]] = 0
	}
	return score
}

func Distance(a, b string) int {
	if len(a) < len(b) {
		a, b = b, a
	}
	if len(b) == 0 {
		return len(a)
	}
	if len(a) <= 32 {
		return mb(a, b)
	}
	return mx(a, b)
}

func Closest(str string, arr []string) string {
	minDistance := math.Inf(1)
	var minIndex int
	for i, s := range arr {
		dist := float64(Distance(str, s))
		if dist < minDistance {
			minDistance = dist
			minIndex = i
		}
	}
	return arr[minIndex]
}
