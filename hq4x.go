package hq4x

import (
	"github.com/donatj/hqxgo/interpolation"
	"github.com/donatj/hqxgo/rgbYuv"
	"math"
)

// public class Hqx_4x extends Hqx {
/**
 * This is the extended Java port of the hq4x algorithm.
 * <b>The destination image must be exactly 4 times as large in both dimensions as the source image</b>
 * The Y, U, V, A parameters will be set as 48, 7, 6 and 0, respectively. Also, wrapping will be false.
 *
 * @param sp the source image data array in ARGB format
 * @param dp the destination image data array in ARGB format
 * @param Xres the horizontal resolution of the source image
 * @param Yres the vertical resolution of the source image
 *
 * @see #hq4x_32_rb(int[], int[], int, int, int, int, int, int, boolean, boolean)
 */
// public static void hq4x_32_rb(final int[] sp, final int[] dp, final int Xres, final int Yres) {
// 	hq4x_32_rb(sp, dp, Xres, Yres, 48, 7, 6, 0, false, false);
// }

/**
 * This and the next caseXXX methods were used to reduce the code size of the main
 * #hq4x_32_rb(int[], int[], int, int, int, int, int, int, boolean, boolean) method because of the Java 65K bytecode limit.
 * Only the necessary methods were created, to leave the maximum code on the original one to avoid excessive calling.
 * However, this is a very bad design (too much code in the same method)
 */
func case0(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case2(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case16(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case64(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case8(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case3(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case6(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case20(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case144(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
}

func case192(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
}

func case96(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case40(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case9(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case66(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case24(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case7(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case148(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
}

func case224(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
}

func case41(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
	dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
	dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
}

func case67(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case70(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case28(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
	dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case152(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
}

func case194(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
}

func case98(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case56(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
	dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

func case25(dp []int, dpIdx int, dpL int, w [9]int) {
	dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
	dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
	dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
	dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
	dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
	dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
	dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
	dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
	dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
	dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
	dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
	dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
}

/**
 * This is the extended Java port of the hq4x algorithm.
 * <b>The destination image must be exactly 4 times as large in both dimensions as the source image</b>
 * @param sp the source image data array in ARGB format
 * @param dp the destination image data array in ARGB format
 * @param Xres the horizontal resolution of the source image
 * @param Yres the vertical resolution of the source image
 * @param trY the Y (luminance) threshold
 * @param trU the U (chrominance) threshold
 * @param trV the V (chrominance) threshold
 * @param trA the A (transparency) threshold
 * @param wrapX used for images that can be seamlessly repeated horizontally
 * @param wrapY used for images that can be seamlessly repeated vertically
 */
func hq4x_32_rb(sp []int, dp []int, Xres int, Yres int, trY int, trU int, trV int, trA int, wrapX bool, wrapY bool) {
	spIdx := 0
	dpIdx := 0
	//Don't shift trA, as it uses shift right instead of a mask for comparisons.
	trY <<= 2 * 8
	trU <<= 1 * 8
	dpL := Xres * 4

	var prevline int
	var nextline int
	w := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for j := 0; j < Yres; j++ {

		if j > 0 {
			prevline = -Xres
		} else {
			if wrapY {
				prevline = Xres * (Yres - 1)
			} else {
				prevline = 0
			}
		}

		if j < Yres-1 {
			nextline = Xres
		} else {
			if wrapY {
				nextline = -(Xres * (Yres - 1))
			} else {
				nextline = 0
			}
		}

		for i := 0; i < Xres; i++ {
			w[1] = sp[spIdx+prevline]
			w[4] = sp[spIdx]
			w[7] = sp[spIdx+nextline]

			if i > 0 {
				w[0] = sp[spIdx+prevline-1]
				w[3] = sp[spIdx-1]
				w[6] = sp[spIdx+nextline-1]
			} else {
				if wrapX {
					w[0] = sp[spIdx+prevline+Xres-1]
					w[3] = sp[spIdx+Xres-1]
					w[6] = sp[spIdx+nextline+Xres-1]
				} else {
					w[0] = w[1]
					w[3] = w[4]
					w[6] = w[7]
				}
			}

			if i < Xres-1 {
				w[2] = sp[spIdx+prevline+1]
				w[5] = sp[spIdx+1]
				w[8] = sp[spIdx+nextline+1]
			} else {
				if wrapX {
					w[2] = sp[spIdx+prevline-Xres+1]
					w[5] = sp[spIdx-Xres+1]
					w[8] = sp[spIdx+nextline-Xres+1]
				} else {
					w[2] = w[1]
					w[5] = w[4]
					w[8] = w[7]
				}
			}

			pattern := 0
			flag := 1

			for k := 0; k < 9; k++ {
				if k == 4 {
					continue
				}

				if w[k] != w[4] {
					if diff(w[4], w[k], trY, trU, trV, trA) {
						pattern |= flag
					}
				}
				flag <<= 1
			}

			switch pattern {
			case 0:
			case 1:
			case 4:
			case 32:
			case 128:
			case 5:
			case 132:
			case 160:
			case 33:
			case 129:
			case 36:
			case 133:
			case 164:
			case 161:
			case 37:
			case 165:
				{
					case0(dp, dpIdx, dpL, w)
					break
				}
			case 2:
			case 34:
			case 130:
			case 162:
				{
					case2(dp, dpIdx, dpL, w)
					break
				}
			case 16:
			case 17:
			case 48:
			case 49:
				{
					case16(dp, dpIdx, dpL, w)
					break
				}
			case 64:
			case 65:
			case 68:
			case 69:
				{
					case64(dp, dpIdx, dpL, w)
					break
				}
			case 8:
			case 12:
			case 136:
			case 140:
				{
					case8(dp, dpIdx, dpL, w)
					break
				}
			case 3:
			case 35:
			case 131:
			case 163:
				{
					case3(dp, dpIdx, dpL, w)
					break
				}
			case 6:
			case 38:
			case 134:
			case 166:
				{
					case6(dp, dpIdx, dpL, w)
					break
				}
			case 20:
			case 21:
			case 52:
			case 53:
				{
					case20(dp, dpIdx, dpL, w)
					break
				}
			case 144:
			case 145:
			case 176:
			case 177:
				{
					case144(dp, dpIdx, dpL, w)
					break
				}
			case 192:
			case 193:
			case 196:
			case 197:
				{
					case192(dp, dpIdx, dpL, w)
					break
				}
			case 96:
			case 97:
			case 100:
			case 101:
				{
					case96(dp, dpIdx, dpL, w)
					break
				}
			case 40:
			case 44:
			case 168:
			case 172:
				{
					case40(dp, dpIdx, dpL, w)
					break
				}
			case 9:
			case 13:
			case 137:
			case 141:
				{
					case9(dp, dpIdx, dpL, w)
					break
				}
			case 18:
			case 50:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 80:
			case 81:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 72:
			case 76:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 10:
			case 138:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+1] = w[4]
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 66:
				{
					case66(dp, dpIdx, dpL, w)
					break
				}
			case 24:
				{
					case24(dp, dpIdx, dpL, w)
					break
				}
			case 7:
			case 39:
			case 135:
				{
					case7(dp, dpIdx, dpL, w)
					break
				}
			case 148:
			case 149:
			case 180:
				{
					case148(dp, dpIdx, dpL, w)
					break
				}
			case 224:
			case 228:
			case 225:
				{
					case224(dp, dpIdx, dpL, w)
					break
				}
			case 41:
			case 169:
			case 45:
				{
					case41(dp, dpIdx, dpL, w)
					break
				}
			case 22:
			case 54:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 208:
			case 209:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 104:
			case 108:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 11:
			case 139:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 19:
			case 51:
				{
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
						dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+1] = interpolation.Mix3To1(w[1], w[4])
						dp[dpIdx+2] = interpolation.Mix5To3(w[1], w[5])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix2To1To1(w[5], w[4], w[1])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 146:
			case 178:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					} else {
						dp[dpIdx+2] = interpolation.Mix2To1To1(w[1], w[4], w[5])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[5], w[1])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 84:
			case 85:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[5], w[4])
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[5], w[7])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix2To1To1(w[7], w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 112:
			case 113:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix2To1To1(w[5], w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[7], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 200:
			case 204:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix2To1To1(w[3], w[4], w[7])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 73:
			case 77:
				{
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[3], w[4])
						dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[3], w[7])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix2To1To1(w[7], w[4], w[3])
					}
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 42:
			case 170:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
						dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix2To1To1(w[1], w[4], w[3])
						dp[dpIdx+dpL] = interpolation.Mix5To3(w[3], w[1])
						dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 14:
			case 142:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix5To3(w[1], w[3])
						dp[dpIdx+2] = interpolation.Mix3To1(w[1], w[4])
						dp[dpIdx+3] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix2To1To1(w[3], w[4], w[1])
						dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					}
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 67:
				{
					case67(dp, dpIdx, dpL, w)
					break
				}
			case 70:
				{
					case70(dp, dpIdx, dpL, w)
					break
				}
			case 28:
				{
					case28(dp, dpIdx, dpL, w)
					break
				}
			case 152:
				{
					case152(dp, dpIdx, dpL, w)
					break
				}
			case 194:
				{
					case194(dp, dpIdx, dpL, w)
					break
				}
			case 98:
				{
					case98(dp, dpIdx, dpL, w)
					break
				}
			case 56:
				{
					case56(dp, dpIdx, dpL, w)
					break
				}
			case 25:
				{
					case25(dp, dpIdx, dpL, w)
					break
				}
			case 26:
			case 31:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 82:
			case 214:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 88:
			case 248:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 74:
			case 107:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 27:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 86:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 216:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 106:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 30:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 210:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 120:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 75:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 29:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 198:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 184:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 99:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 57:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 71:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 156:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 226:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 60:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 195:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 102:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 153:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 58:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 83:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 92:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 202:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 78:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 154:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 114:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					break
				}
			case 89:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 90:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 55:
			case 23:
				{
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
						dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+1] = interpolation.Mix3To1(w[1], w[4])
						dp[dpIdx+2] = interpolation.Mix5To3(w[1], w[5])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix2To1To1(w[5], w[4], w[1])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 182:
			case 150:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					} else {
						dp[dpIdx+2] = interpolation.Mix2To1To1(w[1], w[4], w[5])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[5], w[1])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 213:
			case 212:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[5], w[4])
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[5], w[7])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix2To1To1(w[7], w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 241:
			case 240:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix2To1To1(w[5], w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[7], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 236:
			case 232:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix2To1To1(w[3], w[4], w[7])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 109:
			case 105:
				{
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[3], w[4])
						dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[3], w[7])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix2To1To1(w[7], w[4], w[3])
					}
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 171:
			case 43:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
						dp[dpIdx+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix2To1To1(w[1], w[4], w[3])
						dp[dpIdx+dpL] = interpolation.Mix5To3(w[3], w[1])
						dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 143:
			case 15:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
						dp[dpIdx+dpL] = w[4]
						dp[dpIdx+dpL+1] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix5To3(w[1], w[3])
						dp[dpIdx+2] = interpolation.Mix3To1(w[1], w[4])
						dp[dpIdx+3] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix2To1To1(w[3], w[4], w[1])
						dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					}
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 124:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 203:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 62:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 211:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 118:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 217:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 110:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 155:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 188:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 185:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 61:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 157:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 103:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 227:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 230:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 199:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 220:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 158:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 234:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 242:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					break
				}
			case 59:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 121:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 87:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 79:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 122:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 94:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL+2] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 218:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 91:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL+1] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 229:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 167:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 173:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 181:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 186:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 115:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					break
				}
			case 93:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 206:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 205:
			case 201:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 174:
			case 46:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
						dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL+1] = w[4]
					}
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 179:
			case 147:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
						dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					} else {
						dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 117:
			case 116:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					} else {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					break
				}
			case 189:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 231:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 126:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 219:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 125:
				{
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix3To1(w[4], w[3])
						dp[dpIdx+dpL] = interpolation.Mix3To1(w[3], w[4])
						dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[3], w[7])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix2To1To1(w[7], w[4], w[3])
					}
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 221:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix3To1(w[4], w[5])
						dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[5], w[4])
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[5], w[7])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix2To1To1(w[7], w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 207:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
						dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
						dp[dpIdx+dpL] = w[4]
						dp[dpIdx+dpL+1] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix5To3(w[1], w[3])
						dp[dpIdx+2] = interpolation.Mix3To1(w[1], w[4])
						dp[dpIdx+3] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+dpL] = interpolation.Mix2To1To1(w[3], w[4], w[1])
						dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					}
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 238:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.Mix2To1To1(w[3], w[4], w[7])
						dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[7])
					}
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 190:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					} else {
						dp[dpIdx+2] = interpolation.Mix2To1To1(w[1], w[4], w[5])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[5], w[1])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 187:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
						dp[dpIdx+dpL+1] = w[4]
						dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.Mix2To1To1(w[1], w[4], w[3])
						dp[dpIdx+dpL] = interpolation.Mix5To3(w[3], w[1])
						dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
						dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix3To1(w[4], w[3])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 243:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
						dp[dpIdx+dpL+dpL+3] = interpolation.Mix2To1To1(w[5], w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix3To1(w[4], w[7])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[7], w[5])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 119:
				{
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
						dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+2] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix3To1(w[4], w[1])
						dp[dpIdx+1] = interpolation.Mix3To1(w[1], w[4])
						dp[dpIdx+2] = interpolation.Mix5To3(w[1], w[5])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
						dp[dpIdx+dpL+3] = interpolation.Mix2To1To1(w[5], w[4], w[1])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 237:
			case 233:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[5])
					dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[1])
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 175:
			case 47:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix6To1To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					break
				}
			case 183:
			case 151:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 245:
			case 244:
				{
					dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[3])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix6To1To1(w[4], w[3], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 250:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					break
				}
			case 123:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 95:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 222:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 252:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix4To2To1(w[4], w[1], w[0])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 249:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix4To2To1(w[4], w[1], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					break
				}
			case 235:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[2])
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 111:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix4To2To1(w[4], w[5], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 63:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix4To2To1(w[4], w[7], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 159:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix4To2To1(w[4], w[7], w[6])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 215:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 246:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix4To2To1(w[4], w[3], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 254:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[0])
					dp[dpIdx+1] = interpolation.Mix3To1(w[4], w[0])
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = interpolation.Mix3To1(w[4], w[0])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[0])
					dp[dpIdx+dpL+2] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 253:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+1] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+2] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[1])
					dp[dpIdx+dpL] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+3] = interpolation.Mix7To1(w[4], w[1])
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 251:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[2])
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[2])
					dp[dpIdx+dpL+3] = interpolation.Mix3To1(w[4], w[2])
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					break
				}
			case 239:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					dp[dpIdx+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[5])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[5])
					break
				}
			case 127:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+2] = w[4]
						dp[dpIdx+3] = w[4]
						dp[dpIdx+dpL+3] = w[4]
					} else {
						dp[dpIdx+2] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+3] = interpolation.MixEven(w[1], w[5])
						dp[dpIdx+dpL+3] = interpolation.MixEven(w[5], w[4])
					}
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL] = w[4]
						dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					} else {
						dp[dpIdx+dpL+dpL] = interpolation.MixEven(w[3], w[4])
						dp[dpIdx+dpL+dpL+dpL] = interpolation.MixEven(w[7], w[3])
						dp[dpIdx+dpL+dpL+dpL+1] = interpolation.MixEven(w[7], w[4])
					}
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix3To1(w[4], w[8])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[8])
					break
				}
			case 191:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+2] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+3] = interpolation.Mix7To1(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+2] = interpolation.Mix5To3(w[4], w[7])
					dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix5To3(w[4], w[7])
					break
				}
			case 223:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
						dp[dpIdx+1] = w[4]
						dp[dpIdx+dpL] = w[4]
					} else {
						dp[dpIdx] = interpolation.MixEven(w[1], w[3])
						dp[dpIdx+1] = interpolation.MixEven(w[1], w[4])
						dp[dpIdx+dpL] = interpolation.MixEven(w[3], w[4])
					}
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix3To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[6])
					dp[dpIdx+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+3] = w[4]
						dp[dpIdx+dpL+dpL+dpL+2] = w[4]
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+3] = interpolation.MixEven(w[5], w[4])
						dp[dpIdx+dpL+dpL+dpL+2] = interpolation.MixEven(w[7], w[4])
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.MixEven(w[7], w[5])
					}
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[6])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix3To1(w[4], w[6])
					break
				}
			case 247:
				{
					dp[dpIdx] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix5To3(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+1] = interpolation.Mix7To1(w[4], w[3])
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			case 255:
				{
					if diff(w[3], w[1], trY, trU, trV, trA) {
						dp[dpIdx] = w[4]
					} else {
						dp[dpIdx] = interpolation.Mix2To1To1(w[4], w[1], w[3])
					}
					dp[dpIdx+1] = w[4]
					dp[dpIdx+2] = w[4]
					if diff(w[1], w[5], trY, trU, trV, trA) {
						dp[dpIdx+3] = w[4]
					} else {
						dp[dpIdx+3] = interpolation.Mix2To1To1(w[4], w[1], w[5])
					}
					dp[dpIdx+dpL] = w[4]
					dp[dpIdx+dpL+1] = w[4]
					dp[dpIdx+dpL+2] = w[4]
					dp[dpIdx+dpL+3] = w[4]
					dp[dpIdx+dpL+dpL] = w[4]
					dp[dpIdx+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+2] = w[4]
					dp[dpIdx+dpL+dpL+3] = w[4]
					if diff(w[7], w[3], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL] = interpolation.Mix2To1To1(w[4], w[7], w[3])
					}
					dp[dpIdx+dpL+dpL+dpL+1] = w[4]
					dp[dpIdx+dpL+dpL+dpL+2] = w[4]
					if diff(w[5], w[7], trY, trU, trV, trA) {
						dp[dpIdx+dpL+dpL+dpL+3] = w[4]
					} else {
						dp[dpIdx+dpL+dpL+dpL+3] = interpolation.Mix2To1To1(w[4], w[7], w[5])
					}
					break
				}
			}
			spIdx++
			dpIdx += 4
		}
		dpIdx += (dpL * 3)
	}
}

const Ymask int = 16711680 //0x00FF0000;
const Umask int = 65280    //0x0000FF00;
const Vmask int = 255      //0x000000FF;

func diff(c1 int, c2 int, trY int, trU int, trV int, trA int) bool {
	YUV1 := rgbYuv.GetYuv(c1)
	YUV2 := rgbYuv.GetYuv(c2)

	return ((math.Abs(float64((YUV1&Ymask)-(YUV2&Ymask))) > float64(trY)) ||
		(math.Abs(float64((YUV1&Umask)-(YUV2&Umask))) > float64(trU)) ||
		(math.Abs(float64((YUV1&Vmask)-(YUV2&Vmask))) > float64(trV)) ||
		(math.Abs(float64(((c1 >> 24) - (c2 >> 24)))) > float64(trA)))
}
