# yenc-variant-go
A yEnc encoding / decoding package without header and footer.



## yEnc encoding

* Fetch the character
* Add 42
* Check for NULL, TAB(ascii 9), LF(ascii 10), CR (ascii 13) and '='
* If one of the critical chars encounters then write '=' as escape
  character to the output stream followed by the critical+64.
  (NULL -> =@,   TAB --> =I,  LF --> =J,  CR --> =M,  '=' --> =}

<http://www.yenc.org/yEnc-draft-1.txt>
