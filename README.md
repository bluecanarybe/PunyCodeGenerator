# PunyCodeGenerator
A Go tool to automatically list visually indistinguishable permutations of domain names using IDN and punycode.

Inspired by:
  - https://www.wordfence.com/blog/2017/04/chrome-firefox-unicode-phishing/
  - https://www.xudongz.com/blog/2017/idn-phishing/
  - https://github.com/honoki/punycoder

### Usage

1. Clone tool `git clone https://github.com/bluecanarybe/PunyCodeGenerator.git`
2. Build tool with `go build`
3. Run tool with `./PunyCodeGenerator <domain>`

### Example

```
./PunyCodeGenerator google.com
Genereated punycode domains for google.com :
ɡoogle.com          xn--oogle-qmc.com
gοogle.com          xn--gogle-rce.com
gοogle.com          xn--gogle-rce.com
ɡoogle.com          xn--oogle-qmc.com
googӏe.com          xn--googe-hof.com
googlе.com          xn--googl-3we.com
```
