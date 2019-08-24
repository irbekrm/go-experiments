# bits
Notes on bits, bytes, string encoding etc

## Encoding

Notes from Joel Spolsky's amazing [blog post](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/):

- Unicode is a character set where every letter is encoded to a Unicode code point (i.e `U+0061`)
- UTF-8 is encoding where every Unicode code point can be represented by 1 - 6 bytes (All ASCII chars are represented using a single byte)
- For browsers, email clients, etc to be able to display strings correctly, it is important that the files etc contain information about how the strings are encoded. For emails, it should be in the header (`charset="UTF-8"`), for webpages use the `<meta>` tag.

## Miscellaneous

Go [format verbs](https://golang.org/pkg/fmt/) seem to be a nice way to quickly see a value converted across different formats.

```

    var b uint8 = 0x41
    fmt.Printf("Value: '%#v' of type: '%T'\n", b, b)
    fmt.Printf("In binary form: %b\n",b)
    fmt.Printf("In decimal form: %d\n",b)
    fmt.Printf("As a Unicode code point: %U\n",b)
    fmt.Printf("The corresponding character: %c\n",b)
    
```
[The above in Go Playground](https://play.golang.org/p/TKnk0iUCyu9)
