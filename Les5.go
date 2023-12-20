package main

import "fmt"

func king(pos int) uint64 {

	k := uint64(1 << pos)
	k1 := k & 0xfefefefefefefefe
	k2 := k & 0x7f7f7f7f7f7f7f7f
	mask :=
		(k1 << 7) | (k << 8) | (k2 << 9) |
			(k1 >> 1) | (k2 << 1) |
			(k1 >> 9) | (k >> 8) | (k2 >> 7)

	return mask

}

func knight(pos int) uint64 {
	k := uint64(1 << pos)
	kA := k & 0xfefefefefefefefe
	kAB := k & 0xfcfcfcfcfcfcfcfc
	kH := k & 0x7f7f7f7f7f7f7f7f
	kGH := k & 0x3f3f3f3f3f3f3f3f
	mask := (kA << 15) | (kH << 17) |
		(kAB << 6) | (kGH << 10) |
		(kAB >> 10) | (kGH >> 6) |
		(kA >> 17) | (kH >> 15)

	return mask

}

func castle(pos int) uint64 {
	horiz := (pos / 8) * 8
	vert := pos % 8
	maskHoriz := uint64(0xFF)
	maskVert := uint64(0x101010101010101)

	mask := (maskHoriz << horiz) ^ (maskVert << vert)

	return mask
}

func countBitDiv(mask uint64) int {
	count := 0
	for mask > 0 {
		if (mask & 1) == 1 {
			count++
		}
		mask >>= 1
	}
	return count
}

func countBitSubt(mask uint64) int {
	count := 0
	for mask > 0 {
		mask &= mask - 1
		count++
	}
	return count
}

func countBitCahce(mask uint64) int {
	count := 0
	var cache [256]int
	for i := 0; i < 256; i++ {
		cache[i] = countBitSubt(uint64(i))
	}
	for mask > 0 {
		count += cache[mask&255]
		mask >>= 8
	}
	return count
}

func main() {
	pos := 14
	maskKing := king(pos)
	maskKnight := knight(pos)
	maskCastle := castle(pos)
	fmt.Println("маска для короля = ", maskKing)
	fmt.Println("количество ходов делением =", countBitDiv(maskKing))
	fmt.Println("маска для коня = ", maskKnight)
	fmt.Println("количество ходов вычитанием 1 =", countBitSubt(maskKnight))
	fmt.Println("маска для ладьи = ", maskCastle)
	fmt.Println("количество ходов кэшированием =", countBitCahce(maskCastle))
}
