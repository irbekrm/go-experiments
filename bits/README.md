# bits, etc
Notes on bits, bytes, string encoding etc

## Bits

Mostly notes from Vladimir Vivien's amazing [Medium article](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)

### Bit operators

**&** bitwise AND

**|** bitwise OR

**^** XOR

**&^** AND NOT

**<<<** left shift

**>>>** right shift


## Strings, encoding

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
    fmt.Printf("In binary form: '%b'\n",b)
    fmt.Printf("In decimal form: '%d'\n",b)
    fmt.Printf("As a Unicode code point: '%U'\n",b)
    fmt.Printf("The corresponding character: '%c'\n",b)
    // Output:
    // Value: '0x41' of type: 'uint8'
    // In binary form: '1000001'
    // In decimal form: '65'
    // As a Unicode code point: 'U+0041'
    // The corresponding character: 'A'
    
```
[The above in Go Playground](https://play.golang.org/p/JJu8hpFc7YV)
