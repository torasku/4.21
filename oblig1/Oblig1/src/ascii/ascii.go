package ascii

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func IterateOverExtendedASCIIStringLiteral() {

	strArray := []string{"80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "8A", "8B", "8C", "8D", "8E", "8F", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "9A", "9B", "9C", "9D", "9E", "9F", "A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9", "AA", "AB", "AC", "AD", "AE", "AF", "B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9", "BA", "BB", "BC", "BD", "BE", "BF", "C0", "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "CA", "CB", "CC", "CD", "CE", "CF", "D0", "D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "DA", "DB", "DC", "DD", "DE", "DF", "E0", "E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8", "E9", "EA", "EB", "EC", "ED", "EF", "F0", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "FA", "FB", "FC", "FD", "FE", "FF"}
	/* itererer gjennom alle hexa verdier i utvidet ASCII array
	og konverterer hver verdi til tegn og binaert for alt printes*/
	for _, hexV := range strArray {
		bin, err := hexToBin(hexV)
		if err != nil {
			log.Fatalln(err)
		}
		bs, err := hex.DecodeString(hexV)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\t%v\t%v\n", hexV, string(bs), bin)
	}
}

func ExtendedASCIIText(b []byte) string {

	strArray := []string{"80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "8A", "8B", "8C", "8D", "8E", "8F", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "9A", "9B", "9C", "9D", "9E", "9F", "A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9", "AA", "AB", "AC", "AD", "AE", "AF", "B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9", "BA", "BB", "BC", "BD", "BE", "BF", "C0", "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "CA", "CB", "CC", "CD", "CE", "CF", "D0", "D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "DA", "DB", "DC", "DD", "DE", "DF", "E0", "E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8", "E9", "EA", "EB", "EC", "ED", "EF", "F0", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "FA", "FB", "FC", "FD", "FE", "FF"}
	hexa := hex.EncodeToString(b)
	var slice []string
	var extSlice []string
	var signs string

	/* iterer gjennom input verdier (2 av gangen) og legger til
	i slice. */
	for i := 0; i < len(hexa); i += 2 {
		slice = append(slice, hexa[i:i+2])
		//fmt.Println(slice)
	}
	/* itererer gjennom slice og array som inneholder alle verdier
	i utvidet ASCII og sjekker om noen av disse ligger i slice.
	Hvis ja, legges denne verdien til i extSlice. */
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(strArray); j++ {
			if strings.ContainsAny(strArray[j], slice[i]) {
				extSlice = append(extSlice, slice[i])
				break
			}
		}
	}
	if extSlice == nil {
		fmt.Println("No values represented by the extended ASCII set found.")
	}
	//fmt.Println(extSlice)
	for _, v := range extSlice {
		si, err := hex.DecodeString(v)
		if err != nil {
			panic(err)
		}
		signs += string(si)
		//fmt.Println(string(si))
	}
	return signs
}

/*metode som tar inn hexa verdier, parser dem og
returnerer et print format for binaere tall*/
func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}
	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui), nil
}
