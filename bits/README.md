# bits, etc
Notes on bits, bytes, string encoding etc

## Bits

Mostly notes from Vladimir Vivien's amazing [Medium article](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)

### Bit operators

**&**  bitwise AND

**|**  bitwise OR

**^**  XOR

**&^**  AND NOT

**<<**  left shift

**>>**  right shift

#### Bitwise AND (&)
- can be used to clear a number of bits

Example ([also in playground](https://play.golang.org/p/lZPfWzTSoDL)):
```
    var a uint8 = 198
    fmt.Printf("a in binary: %08b\n", a)
    // Clear the last 4 bits
    a &= 0xF0 // 0xF0 => 11110000
    fmt.Printf("a with last 4 bits cleared: %08b\n", a)
    // Output:
    // a in binary: 11000110
    // a with last 4 bits cleared: 11000000
```

#### Bitwise OR (|)
- can be used to set specific bits

Example ([also in playground](https://play.golang.org/p/zQ7VCiEYYIg)):
```
    var b uint8 = 12
    fmt.Printf("b in binary: %08b\n", b)
    // Set the 2nd, 6th and 7th bit
    b |= 98 // 98 => 01100010
    fmt.Printf("b with 2nd, 6th, 7th bit set: %08b\n",b)
    // Output:
    // b in binary: 00001100
    // b with 2nd, 6th, 7th bit set: 01101110
```

#### Bitwise XOR (^)
- can be used to toggle bits

Example ([also in playground](https://play.golang.org/p/DUHOzrMl5vj)):
```
    var a uint8 = 230
    fmt.Printf("a in binary: %08b\n",a)
    // Toggle the last 4 bits
    a ^= 0x0F // 0x0F => 00001111
    fmt.Printf("a with last 4 bits toggled: %08b\n",a)
    // Output:
    // a in binary: 11100110
    // a with last 4 bits toggled: 11101001
```

- can be used to reverse bits (one's complement- TODO: read more on this)

Example ([also in playground](https://play.golang.org/p/mRBDlnZGXTu)):
```
    var a uint8 = 230
    fmt.Printf("a in binary: %08b\n",a)
    fmt.Printf("a reversed: %08b\n", ^a)
```
#### AND NOT (&^)
- can be used to clear bits

Example ([also in playground](https://play.golang.org/p/OVbaEUDf47Y)):
```
    var a uint8 = 230
    fmt.Printf("a in binary: %08b\n",a)
    // Clear the last 4 bits
    a &^= 0x0F // 0x0F => 00001111
    fmt.Printf("a with last 4 bits cleared: %08b\n",a)
    // Output:
    // a in binary: 11100110
    // a with last 4 bits cleared: 11100000
```
#### left shift and right shift (<<, >>)
- can be used to multipy / divide by a power of 2

Example ([also in playground](https://play.golang.org/p/4N9KZ_T_2oz)):
```
    var a uint8 = 3
    var b uint8 = 120
    var result uint8

    fmt.Printf("a: %08b, %d\n",a, a)
    // a multiplied by 2 (a * 2**1)
    result = a << 1 
    fmt.Printf("a times 2: %08b, %d\n", result, result)
    // a multiplied by 8 (a * 2**3)
    result = a << 3
    fmt.Printf("a times 8: %08b, %d\n", result, result)

    fmt.Printf("b: %08b, %d\n", b, b)
    // b divided by 2 (b / 2**1)
    result = b >> 1
    fmt.Printf("b divided by 2: %08b, %d\n", result, result)
    // b divided by 8 (b / 2**3)
    result = b >> 3
    fmt.Printf("b divided by 8: %08b, %d\n", result, result)

    // Output:
    // a: 00000011, 3
    // a times 2: 00000110, 6
    // a times 8: 00011000, 24
    // b: 01111000, 120
    // b divided by 2: 00111100, 60
    // b divided by 8: 00001111, 15
```
- can also be used to manipulate bits at a specific position

Example ([also in playground](https://play.golang.org/p/PCpapq2XSiW)):
```
    var a uint8 = 145
    fmt.Printf("a in binary: %08b\n",a)

    // Set the third bit
    result := a | (1<<2) // (1<<2 => 0100- the last bit shifted by 2 positions left)
    fmt.Printf("a with third bit set: %-8b\n", result)

    // Check if 5th bit is set
    isSet := a&(1<<4) != 0 // (1<<4 => 00010000- the last bit shifted by 4 positions left)
    fmt.Printf("5th bit of a is set: %t\n", isSet)

    // Unset the 5th bit
    result = a &^ (1<<4)
    fmt.Printf("a with fifth bit unset: %08b\n", result)
```
	
	
	// Output:
	// a in binary: 10010001
	// a with third bit set: 10010101
	// 5th bit of a is set: true
	// a with fifth bit unset: 10000001



## Strings, encoding, etc

Notes from Joel Spolsky's amazing [blog post](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/):

- Unicode is a character set where every letter is encoded to a Unicode code point (i.e `U+0061`)
- UTF-8 is encoding where every Unicode code point can be represented by 1 - 6 bytes (All ASCII chars are represented using a single byte)
- For browsers, email clients, etc to be able to display strings correctly, it is important that the files etc contain information about how the strings are encoded. For emails, it should be in the header (`charset="UTF-8"`), for webpages use the `<meta>` tag.

Notes from Rob Pike's [blog post](https://blog.golang.org/strings):

- In Go a string is a slice of bytes
- Go source code is encoded using UTF-8
- A string value can contain any bytes including ones that are not valid UTF-8 ```s := "\xAD\x01"```
- A raw string cannot contain escape sequences and is therefore always valid UTF-8 ``    s := `raw string`    ``
- A Unicode code point is a rune in Go. Alias for rune is int32. A character is rune constant ``    r := `c`    ``
- When ranging over a string, the range goes over runes not bytes

## Miscellaneous

Go [format verbs](https://golang.org/pkg/fmt/) seem to be a nice way to quickly see a value converted across different formats.

```

    var b uint8 = 0x41
    fmt.Printf("Value: '%#v' of type: '%T'\n", b, b)
    fmt.Printf("In binary form: '%08b'\n",b)
    fmt.Printf("In decimal form: '%d'\n",b)
    fmt.Printf("As a Unicode code point: '%U'\n",b)
    fmt.Printf("Unicode code point and representation: '%#U'\n",b)
    fmt.Printf("The corresponding character: '%c'\n",b)
    // Output:
    // Value: '0x41' of type: 'uint8'
    // In binary form: '01000001'
    // In decimal form: '65'
    // As a Unicode code point: 'U+0041'
    // Unicode code point and representation: 'U+0041 'A''
    // The corresponding character: 'A'

```
[The above in Go Playground](https://play.golang.org/p/6I6RDt_Es26)
