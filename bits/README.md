# bits
Notes on bits, bytes, string encoding etc

## Encoding

Notes from Joel Spolsky's amazing [blog post](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/):

- Unicode is a character set where every letter is encoded to a Unicode code point (i.e `U+0061`)
- UTF-8 is encoding where every Unicode code point can be represented by 1 - 6 bytes (All ASCII chars are represented using a single byte)
- For browsers, email clients, etc to be able to display strings correctly, it is important that the files etc contain information about how the strings are encoded. For emails, it should be in the header (`charset="UTF-8"`), for webpages use the `<meta>` tag.
