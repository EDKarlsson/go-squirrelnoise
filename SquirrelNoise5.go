package random

// -----------------------------------------------------------------------------------------------
// SquirrelNoise5.go
//

// ///////////////////////////////////////////////////////////////////////////////////////////////
// SquirrelNoise5 - Squirrel's Raw Noise utilities (version 5)
//
// This code is made available under the Creative Commons attribution 3.0 license (CC-BY-3.0 US):
//	Attribution in source code comments (even closed-source/commercial code) is sufficient.
//	License summary and text available at: https://creativecommons.org/licenses/by/3.0/us/
//
// These noise functions were written by Squirrel Eiserloh as a cheap and simple substitute for
//	the [sometimes awful] bit-noise sample code functions commonly found on the web, many of which
//	are hugely biased or terribly patterned, e.g. having bits which are on (or off) 75% or even
//	100% of the time (or are excessively overkill/slow for our needs, such as MD5 or SHA).
//
// Note: This is work in progress; not all functions have been tested.  Use at your own risk.
//	Please report any bugs, issues, or bothersome cases to SquirrelEiserloh at gmail.com.
//
// The following functions are all based on a simple bit-noise hash function which returns an
//	unsigned integer containing 32 reasonably-well-scrambled bits, based on a given (signed)
//	integer input parameter (position/index) and [optional] seed.  Kind of like looking up a
//	value in an infinitely large [non-existent] table of previously rolled random numbers.
//
// These functions are deterministic and random-access / order-independent (i.e. state-free),
//	so they are particularly well-suited for use in smoothed/fractal/simplex/Perlin noise
//	functions and out-of-order (or or-demand) procedural content generation (i.e. that mountain
//	village is the same whether you generated it first or last, ahead of time or just now).
//
// The N-dimensional variations simply hash their multidimensional coordinates down to a single
//	32-bit index and then proceed as usual, so while results are not unique they should
//	(hopefully) not seem locally predictable or repetitive.
//
// ///////////////////////////////////////////////////////////////////////////////////////////////

// -----------------------------------------------------------------------------------------------
// Raw pseudorandom noise functions (random-access / deterministic).  Basis of all other noise.
//
// func Get1dNoiseUint( index int32, seed uint32 = 0 )
// func Get2dNoiseUint( indexX int32, indexY int32, seed uint32 = 0 )
// func Get3dNoiseUint( indexX int32, indexY int32, indexZ int32, seed uint32 = 0 )
// func Get4dNoiseUint( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 =0 )

// -----------------------------------------------------------------------------------------------
// Same functions, mapped to floats in [0,1] for convenience.
//
// func Get1dNoiseZeroToOne( index int32, seed uint32 =0 )
// func Get2dNoiseZeroToOne( indexX int32, indexY int32, seed uint32 = 0 )
// func Get3dNoiseZeroToOne( indexX int32, indexY int32, indexZ int32, seed uint32 = 0 )
// func Get4dNoiseZeroToOne( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 = 0 )

// -----------------------------------------------------------------------------------------------
// Same functions, mapped to floats in [-1,1] for convenience.
//
// func Get1dNoiseNegOneToOne( index int32, seed uint32 =0 )
// func Get2dNoiseNegOneToOne( indexX int32, indexY int32, seed uint32 = 0 )
// func Get3dNoiseNegOneToOne( indexX int32, indexY int32, indexZ int32, seed uint32 = 0 )
// func Get4dNoiseNegOneToOne( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 = 0 )

// ///////////////////////////////////////////////////////////////////////////////////////////////
// Inline function definitions below
// ///////////////////////////////////////////////////////////////////////////////////////////////

// -----------------------------------------------------------------------------------------------
// Fast hash of an int32 into a different (unrecognizable) uint32.
//
// Returns an unsigned integer containing 32 reasonably-well-scrambled bits, based on the hash
//	of a given (signed) integer input parameter (position/index) and [optional] seed.  Kind of
//	like looking up a value in an infinitely large table of previously generated random numbers.
//
// I call this particular approach SquirrelNoise5 (5th iteration of my 1D raw noise function).
//
// Many thanks to Peter Schmidt-Nielsen whose outstanding analysis helped identify a weakness
//	in the SquirrelNoise3 code I originally used in my GDC 2017 talk, "Noise-based RNG".
//	Version 5 avoids a noise repetition found in version 3 at extremely high position values
//	caused by a lack of influence by some of the high input bits onto some of the low output bits.
//
// The revised SquirrelNoise5 function ensures all input bits affect all output bits, and to
//	(for me) a statistically acceptable degree.  I believe the worst-case here is in the amount
//	of influence input position bit #30 has on output noise bit #0 (49.99%, vs. 50% ideal).
//

func SquirrelNoise5(positionX int32, seed uint32) uint32 {
	var Sq5BitNoise1 uint32 = 0xd2a80a3f // 11010010101010000000101000111111
	var Sq5BitNoise2 uint32 = 0xa884f197 // 10101000100001001111000110010111
	var Sq5BitNoise3 uint32 = 0x6C736F4B // 01101100011100110110111101001011
	var Sq5BitNoise4 uint32 = 0xB79F3ABB // 10110111100111110011101010111011
	var Sq5BitNoise5 uint32 = 0x1b56c4f5 // 00011011010101101100010011110101

	var mangledBits uint32
	mangledBits = uint32(positionX)
	mangledBits *= Sq5BitNoise1
	mangledBits += seed
	mangledBits ^= mangledBits >> 9
	mangledBits += Sq5BitNoise2
	mangledBits ^= mangledBits >> 11
	mangledBits *= Sq5BitNoise3
	mangledBits ^= mangledBits >> 13
	mangledBits += Sq5BitNoise4
	mangledBits ^= mangledBits >> 15
	mangledBits *= Sq5BitNoise5
	mangledBits ^= mangledBits >> 17
	return mangledBits
}

// ------------------------------------------------------------------------------------------------

func Get1dNoiseUint(positionX int32, seed uint32) uint32 {
	return SquirrelNoise5(positionX, seed)
}

// -----------------------------------------------------------------------------------------------

func Get2dNoiseUint(indexX int32, indexY int32, seed uint32) uint32 {
	const PrimeNumber int32 = 198491317 // Large prime number with non-boring bits
	return SquirrelNoise5(indexX+(PrimeNumber*indexY), seed)
}

// -----------------------------------------------------------------------------------------------

func Get3dNoiseUint(indexX int32, indexY int32, indexZ int32, seed uint32) uint32 {
	const Prime1 int32 = 198491317 // Large prime number with non-boring bits
	const Prime2 int32 = 6542989   // Large prime number with distinct and non-boring bits
	return SquirrelNoise5(indexX+(Prime1*indexY)+(Prime2*indexZ), seed)
}

// -----------------------------------------------------------------------------------------------

func Get4dNoiseUint(indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32) uint32 {
	const Prime1 int32 = 198491317 // Large prime number with non-boring bits
	const Prime2 int32 = 6542989   // Large prime number with distinct and non-boring bits
	const Prime3 int32 = 357239    // Large prime number with distinct and non-boring bits
	return SquirrelNoise5(indexX+(Prime1*indexY)+(Prime2*indexZ)+(Prime3*indexT), seed)
}

// -----------------------------------------------------------------------------------------------

func Get1dNoiseZeroToOne(index int32, seed uint32) float32 {
	const OneOverMaxUint = 1.0 / float64(0xFFFFFFFF)
	return float32(OneOverMaxUint * float64(SquirrelNoise5(index, seed)))
}

// -----------------------------------------------------------------------------------------------

func Get2dNoiseZeroToOne(indexX int32, indexY int32, seed uint32) float32 {
	const OneOverMaxUint = 1.0 / float64(0xFFFFFFFF)
	return float32(OneOverMaxUint * float64(Get2dNoiseUint(indexX, indexY, seed)))
}

// -----------------------------------------------------------------------------------------------

func Get3dNoiseZeroToOne(indexX int32, indexY int32, indexZ int32, seed uint32) float32 {
	const OneOverMaxUint = 1.0 / float64(0xFFFFFFFF)
	return float32(OneOverMaxUint * float64(Get3dNoiseUint(indexX, indexY, indexZ, seed)))
}

// -----------------------------------------------------------------------------------------------

func Get4dNoiseZeroToOne(indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32) float32 {
	const OneOverMaxUint = 1.0 / float64(0xFFFFFFFF)
	return float32(OneOverMaxUint * float64(Get4dNoiseUint(indexX, indexY, indexZ, indexT, seed)))
}

// -----------------------------------------------------------------------------------------------

func Get1dNoiseNegOneToOne(index int32, seed uint32) float32 {
	const OneOverMaxInt = 1.0 / float64(0x7FFFFFFF)
	return float32(OneOverMaxInt * float64(int32(SquirrelNoise5(index, seed))))
}

// -----------------------------------------------------------------------------------------------

func Get2dNoiseNegOneToOne(indexX int32, indexY int32, seed uint32) float32 {
	const OneOverMaxInt = 1.0 / float64(0x7FFFFFFF)
	return float32(OneOverMaxInt * float64(int32(Get2dNoiseUint(indexX, indexY, seed))))
}

// -----------------------------------------------------------------------------------------------

func Get3dNoiseNegOneToOne(indexX int32, indexY int32, indexZ int32, seed uint32) float32 {
	const OneOverMaxInt = 1.0 / float64(0x7FFFFFFF)
	return float32(OneOverMaxInt * float64(int32(Get3dNoiseUint(indexX, indexY, indexZ, seed))))
}

// -----------------------------------------------------------------------------------------------

func Get4dNoiseNegOneToOne(indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32) float32 {
	const OneOverMaxInt = 1.0 / float64(0x7FFFFFFF)
	return float32(OneOverMaxInt * float64(int32(Get4dNoiseUint(indexX, indexY, indexZ, indexT, seed))))
}
