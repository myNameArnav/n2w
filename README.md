# n2w

A simple Go program that converts a number into its English word representation.

## Usage

```
Usage: n2w <number>
```

## Example

```
$ n2w 1234567890
One hundred twenty three thousand four hundred fifty six thousand seven hundred eighty nine
$ n2w -1234567890
Negative one hundred twenty three thousand four hundred fifty six thousand seven hundred eighty nine
$ n2w 12345678901
Error: number too large, exceeds defined thousands scale
```