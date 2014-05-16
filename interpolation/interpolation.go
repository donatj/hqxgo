package interpolation

/**
* Helper class to interpolate colors. Nothing to see here, move along...
 */

const mask4 int = -16777216 //0xFF000000
const mask2 int = 65280     //0x0000FF00
const mask13 int = 16711935 //0x00FF00FF

// return statements:
//	 1. line: green
//	 2. line: red and blue
//	 3. line: alpha

func Mix3To1(c1 int, c2 int) int {
	//return (c1*3+c2) >> 2;
	if c1 == c2 {
		return c1
	}
	return ((((c1&mask2)*3 + (c2 & mask2)) >> 2) & mask2) |
		((((c1&mask13)*3 + (c2 & mask13)) >> 2) & mask13) |
		((((c1&mask4)>>2)*3 + ((c2 & mask4) >> 2)) & mask4)
}

func Mix2To1To1(c1 int, c2 int, c3 int) int {
	//return (c1*2+c2+c3) >> 2;
	return ((((c1&mask2)*2 + (c2 & mask2) + (c3 & mask2)) >> 2) & mask2) |
		((((c1&mask13)*2 + (c2 & mask13) + (c3 & mask13)) >> 2) & mask13) |
		((((c1&mask4)>>2)*2 + ((c2 & mask4) >> 2) + ((c3 & mask4) >> 2)) & mask4)
}

func Mix7To1(c1 int, c2 int) int {
	//return (c1*7+c2)/8;
	if c1 == c2 {
		return c1
	}
	return ((((c1&mask2)*7 + (c2 & mask2)) >> 3) & mask2) |
		((((c1&mask13)*7 + (c2 & mask13)) >> 3) & mask13) |
		((((c1&mask4)>>3)*7 + ((c2 & mask4) >> 3)) & mask4)
}

func Mix2To7To7(c1 int, c2 int, c3 int) int {
	//return (c1*2+(c2+c3)*7)/16;
	return ((((c1&mask2)*2 + (c2&mask2)*7 + (c3&mask2)*7) >> 4) & mask2) |
		((((c1&mask13)*2 + (c2&mask13)*7 + (c3&mask13)*7) >> 4) & mask13) |
		((((c1&mask4)>>4)*2 + ((c2&mask4)>>4)*7 + ((c3&mask4)>>4)*7) & mask4)
}

func MixEven(c1 int, c2 int) int {
	//return (c1+c2) >> 1;
	if c1 == c2 {
		return c1
	}
	return ((((c1 & mask2) + (c2 & mask2)) >> 1) & mask2) |
		((((c1 & mask13) + (c2 & mask13)) >> 1) & mask13) |
		((((c1 & mask4) >> 1) + ((c2 & mask4) >> 1)) & mask4)
}

func Mix4To2To1(c1 int, c2 int, c3 int) int {
	//return (c1*5+c2*2+c3)/8;
	return ((((c1&mask2)*5 + (c2&mask2)*2 + (c3 & mask2)) >> 3) & mask2) |
		((((c1&mask13)*5 + (c2&mask13)*2 + (c3 & mask13)) >> 3) & mask13) |
		((((c1&mask4)>>3)*5 + ((c2&mask4)>>3)*2 + ((c3 & mask4) >> 3)) & mask4)
}

func Mix6To1To1(c1 int, c2 int, c3 int) int {
	//return (c1*6+c2+c3)/8;
	return ((((c1&mask2)*6 + (c2 & mask2) + (c3 & mask2)) >> 3) & mask2) |
		((((c1&mask13)*6 + (c2 & mask13) + (c3 & mask13)) >> 3) & mask13) |
		((((c1&mask4)>>3)*6 + ((c2 & mask4) >> 3) + ((c3 & mask4) >> 3)) & mask4)
}

func Mix5To3(c1 int, c2 int) int {
	//return (c1*5+c2*3)/8;
	if c1 == c2 {
		return c1
	}
	return ((((c1&mask2)*5 + (c2&mask2)*3) >> 3) & mask2) |
		((((c1&mask13)*5 + (c2&mask13)*3) >> 3) & mask13) |
		((((c1&mask4)>>3)*5 + ((c2&mask4)>>3)*3) & mask4)
}

func Mix2To3To3(c1 int, c2 int, c3 int) int {
	//return (c1*2+(c2+c3)*3)/8;
	return ((((c1&mask2)*2 + (c2&mask2)*3 + (c3&mask2)*3) >> 3) & mask2) |
		((((c1&mask13)*2 + (c2&mask13)*3 + (c3&mask13)*3) >> 3) & mask13) |
		((((c1&mask4)>>3)*2 + ((c2&mask4)>>3)*3 + ((c3&mask4)>>3)*3) & mask4)
}

func Mix14To1To1(c1 int, c2 int, c3 int) int {
	//return (c1*14+c2+c3)/16;
	return ((((c1&mask2)*14 + (c2 & mask2) + (c3 & mask2)) >> 4) & mask2) |
		((((c1&mask13)*14 + (c2 & mask13) + (c3 & mask13)) >> 4) & mask13) |
		((((c1&mask4)>>4)*14 + ((c2 & mask4) >> 4) + ((c3 & mask4) >> 4)) & mask4)
}
