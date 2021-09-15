# SquirrelNoise5 - Squirrel's Raw Noise utilities (version 5)

> Erik Karlsson: I'm not the original author of this code. Just wanted to convert this into Go for a side project. Will make this a module at some point to be importable.

This code is made available under the Creative Commons attribution 3.0 license (CC-BY-3.0 US):
Attribution in source code comments (even closed-source/commercial code) is sufficient. License summary and text
available at: https://creativecommons.org/licenses/by/3.0/us/

These noise functions were written by Squirrel Eiserloh as a cheap and simple substitute for the [sometimes awful]
bit-noise sample code functions commonly found on the web, many of which are hugely biased or terribly patterned, e.g.
having bits which are on (or off) 75% or even 100% of the time (or are excessively overkill/slow for our needs, such as
MD5 or SHA).

Note: This is work in progress; not all functions have been tested. Use at your own risk. Please report any bugs,
issues, or bothersome cases to SquirrelEiserloh at gmail.com.

The following functions are all based on a simple bit-noise hash function which returns an unsigned integer containing
32 reasonably-well-scrambled bits, based on a given (signed)
integer input parameter (position/index) and [optional] seed. Kind of like looking up a value in an infinitely
large [non-existent] table of previously rolled random numbers.

These functions are deterministic and random-access / order-independent (i.e. state-free), so they are particularly
well-suited for use in smoothed/fractal/simplex/Perlin noise functions and out-of-order (or or-demand) procedural
content generation (i.e. that mountain village is the same whether you generated it first or last, ahead of time or just
now).

The N-dimensional variations simply hash their multidimensional coordinates down to a single 32-bit index and then
proceed as usual, so while results are not unique they should
(hopefully) not seem locally predictable or repetitive.

-----------------------------------------------------------------------------------------------

### Raw pseudorandom noise functions (random-access / deterministic). Basis of all other noise.

```go
func Get1dNoiseUint( index int32, seed uint32 ) uint32
func Get2dNoiseUint( indexX int32, indexY int32, seed uint32 ) uint32
func Get3dNoiseUint( indexX int32, indexY int32, indexZ int32, seed uint32 ) uint32
func Get4dNoiseUint( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 ) uint32
```

-----------------------------------------------------------------------------------------------

### Same functions, mapped to floats in [0,1] for convenience.

```go
func Get1dNoiseZeroToOne( index int32, seed uint32 ) float32
func Get2dNoiseZeroToOne( indexX int32, indexY int32, seed uint32 ) float32
func Get3dNoiseZeroToOne( indexX int32, indexY int32, indexZ int32, seed uint32 ) float32
func Get4dNoiseZeroToOne( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 ) float32
```

-----------------------------------------------------------------------------------------------

### Same functions, mapped to floats in [-1,1] for convenience.

```go
func Get1dNoiseNegOneToOne( index int32, seed uint32 ) float32
func Get2dNoiseNegOneToOne( indexX int32, indexY int32, seed uint32 ) float32
func Get3dNoiseNegOneToOne( indexX int32, indexY int32, indexZ int32, seed uint32 ) float32
func Get4dNoiseNegOneToOne( indexX int32, indexY int32, indexZ int32, indexT int32, seed uint32 ) float32
```

## Inline function definitions below

-----------------------------------------------------------------------------------------------
### Fast hash of an int32 into a different (unrecognizable) uint32.

Returns an unsigned integer containing 32 reasonably-well-scrambled bits, based on the hash of a given (signed) integer
input parameter (position/index) and [optional] seed. Kind of like looking up a value in an infinitely large table of
previously generated random numbers.

I call this particular approach SquirrelNoise5 (5th iteration of my 1D raw noise function).

Many thanks to Peter Schmidt-Nielsen whose outstanding analysis helped identify a weakness in the SquirrelNoise3 code I
originally used in my GDC 2017 talk, "Noise-based RNG". Version 5 avoids a noise repetition found in version 3 at
extremely high position values caused by a lack of influence by some of the high input bits onto some of the low output
bits.

The revised SquirrelNoise5 function ensures all input bits affect all output bits, and to
(for me) a statistically acceptable degree. I believe the worst-case here is in the amount of influence input position
bit #30 has on output noise bit #0 (49.99%, vs. 50% ideal).
