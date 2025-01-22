package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	deci := AtoiBase(nbr, baseFrom)
	return AtoiBase(deci, baseTo)
}
