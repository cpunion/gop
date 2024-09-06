Go+ Specification for STEM Education
=====

## Comments

Comments serve as program documentation. There are three forms:

* _Line comments_ start with the character sequence `//` and stop at the end of the line.
* _Line comments_ start with the character sequence `#` and stop at the end of the line.
* _General comments_ start with the character sequence `/*` and stop with the first subsequent character sequence `*/`.

A _general comment_ containing no newlines acts like a space. Any other comment acts like a newline.

```
# this is a line comment
// this is another line comment
/* this is a general comment */
```

## Literals

### Integer literals

An integer literal is a sequence of digits representing an [integer constant](). An optional prefix sets a non-decimal base: 0b or 0B for binary, 0, 0o, or 0O for octal, and 0x or 0X for hexadecimal. A single 0 is considered a decimal zero. In hexadecimal literals, letters a through f and A through F represent values 10 through 15.

For readability, an underscore character _ may appear after a base prefix or between successive digits; such underscores do not change the literal's value.

```go
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
```

### Floating-point literals

A floating-point literal is a decimal or hexadecimal representation of a [floating-point constant]().

A decimal floating-point literal consists of an integer part (decimal digits), a decimal point, a fractional part (decimal digits), and an exponent part (e or E followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; one of the decimal point or the exponent part may be elided. An exponent value exp scales the mantissa (integer and fractional part) by 10<sup>exp</sup>.

A hexadecimal floating-point literal consists of a 0x or 0X prefix, an integer part (hexadecimal digits), a radix point, a fractional part (hexadecimal digits), and an exponent part (p or P followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; the radix point may be elided as well, but the exponent part is required. (This syntax matches the one given in IEEE 754-2008 §5.12.3.) An exponent value exp scales the mantissa (integer and fractional part) by 2<sup>exp</sup>.

For readability, an underscore character _ may appear after a base prefix or between successive digits; such underscores do not change the literal value.

```go
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375

0x15e-2      // == 0x15e - 2 (integer subtraction)

0x.p1        // invalid: mantissa has no digits
1p-2         // invalid: p exponent requires hexadecimal mantissa
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
1_.5         // invalid: _ must separate successive digits
1._5         // invalid: _ must separate successive digits
1.5_e1       // invalid: _ must separate successive digits
1.5e_1       // invalid: _ must separate successive digits
1.5e1_       // invalid: _ must separate successive digits
```

### Rational literals

TODO

```sh
1r       # bigint 1
2/3r     # bigrat 2/3
```

### Imaginary literals

An imaginary literal represents the imaginary part of a [complex constant](). It consists of an [integer](#integer-literals) or [floating-point](#floating-point-literals) literal followed by the lowercase letter _i_. The value of an imaginary literal is the value of the respective integer or floating-point literal multiplied by the imaginary unit _i_.

For backward compatibility, an imaginary literal's integer part consisting entirely of decimal digits (and possibly underscores) is considered a decimal integer, even if it starts with a leading 0.

```go
0i
0123i         // == 123i for backward-compatibility
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

### Boolean literals

The boolean truth values are represented by the predeclared constants `true` and `false`.

```go
true
false
```

### Rune literals

A rune literal represents a [rune constant](), an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes, as in `'x'` or `'\n'`. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats.

The simplest form represents the single character within the quotes; since Go+ source text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent a single integer value. For instance, the literal `'a'` holds a single byte representing a literal a, Unicode `U+0061`, value 0x61, while `'ä'` holds two bytes (0xc3 0xa4) representing a literal a-dieresis, `U+00E4`, value 0xe4.

Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: `\x` followed by exactly two hexadecimal digits; `\u` followed by exactly four hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a plain backslash `\` followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.

Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and `\U` represent Unicode code points so within them some values are illegal, in particular those above 0x10FFFF and surrogate halves.

After a backslash, certain single-character escapes represent special values:

```
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000B vertical tab
\\   U+005C backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
```

An unrecognized character following a backslash in a rune literal is illegal.

```go
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\k'         // illegal: k is not recognized after a backslash
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\400'       // illegal: octal value over 255
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

### String literals

A string literal represents a [string constant]() obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

Raw string literals are character sequences between back quotes, as in \`foo\`. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters (`'\r'`) inside raw string literals are discarded from the raw string value.

Interpreted string literals are character sequences between double quotes, as in `"bar"`. Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in [rune literals](#rune-literals) (except that `\'` is illegal and `\"` is legal), with the same restrictions. The three-digit octal (`\nnn`) and two-digit hexadecimal (`\xnn`) escapes represent individual bytes of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual characters. Thus inside a string literal `\377` and `\xFF` represent a single byte of value 0xFF=255, while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes 0xc3 0xbf of the UTF-8 encoding of character `U+00FF`.

```go
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:

```go
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

If the source code represents a character as two code points, such as a combining form involving an accent and a letter, the result will be an error if placed in a rune literal (it is not a single code point), and will appear as two code points if placed in a string literal.

### Special literals

TODO

```go
nil
iota
```

## Constants

There are _boolean constants_, _rune constants_, _integer constants_, _floating-point constants_, _complex constants_, and _string constants_. Rune, integer, floating-point, and complex constants are collectively called _numeric constants_.

A constant value is represented by a [rune](#rune-literals), [integer](#integer-literals), [floating-point](#floating-point-literals), [imaginary](#imaginary-literals), [boolean](#boolean-literals) or [string](#string-literals) literal, an identifier denoting a constant, a [constant expression](), a [conversion](#conversions) with a result that is a constant, or the result value of some built-in functions such as `min` or `max` applied to constant arguments, `unsafe.Sizeof` applied to [certain values](), `cap` or `len` applied to [some expressions](), `real` and `imag` applied to a complex constant and complex applied to numeric constants. The predeclared identifier `iota` denotes an integer constant.

Not all literals are constants. For example:

```sh
nil
1r
2/3r
```

In general, complex constants are a form of [constant expression]() and are discussed in that section.

Numeric constants represent exact values of arbitrary precision and do not overflow. Consequently, there are no constants denoting the IEEE-754 negative zero, infinity, and not-a-number values.

Constants may be [typed](#types) or _untyped_. Literal constants (including `true`, `false`, `iota`), and certain [constant expressions]() containing only untyped constant operands are untyped.

A constant may be given a type explicitly by a [constant declaration]() or [conversion](#conversions), or implicitly when used in a [variable declaration]() or an [assignment statement]() or as an operand in an [expression](#expressions). It is an error if the constant value cannot be [represented]() as a value of the respective type.

An untyped constant has a _default type_ which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a [short variable declaration]() such as `i := 0` where there is no explicit type. The default type of an untyped constant is `bool`, `rune`, `int`, `float64`, `complex128`, or `string` respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.


## Variables

A variable is a storage location for holding a value. The set of permissible values is determined by the variable's [type](#types).

A [variable declaration]() or, for function parameters and results, the signature of a [function declaration]() or [function literal]() reserves storage for a named variable. Calling the built-in function [new]() or taking the address of a [composite literal]() allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) [pointer indirection](#address-operators).

_Structured_ variables of [array](#array-types), [slice](#slice-types), and [class](#classes) types have elements and fields that may be [addressed](#address-operators) individually. Each such element acts like a variable.

The _static type_ (or just _type_) of a variable is the type given in its declaration, the type provided in the new call or composite literal, or the type of an element of a class variable. Variables of interface type also have a distinct _dynamic type_, which is the (non-interface) type of the value assigned to the variable at run time (unless the value is the predeclared identifier `nil`, which has no type). The dynamic type may vary during execution but values stored in interface variables are always [assignable]() to the static type of the variable.

```go
var x any  // x is nil and has static type any
var v *T   // v has value nil, static type *T
x = 42     // x has value 42 and dynamic type int
x = v      // x has value (*T)(nil) and dynamic type *T
```

A variable's value is retrieved by referring to the variable in an [expression](#expressions); it is the most recent value [assigned]() to the variable. If a variable has not yet been assigned a value, its value is the [zero value]() for its type.


## Types

### Boolean types

A _boolean type_ represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is `bool`; it is a defined type.

```go
bool
```

### Numeric types

An _integer_, _floating-point_, _rational_ or _complex_ type represents the set of integer, floating-point, or complex values, respectively. They are collectively called _numeric types_. The predeclared architecture-independent numeric types are:

```go
uint8       // the set of all unsigned  8-bit integers (0 to 255)
uint16      // the set of all unsigned 16-bit integers (0 to 65535)
uint32      // the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      // the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        // the set of all signed  8-bit integers (-128 to 127)
int16       // the set of all signed 16-bit integers (-32768 to 32767)
int32       // the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       // the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     // the set of all IEEE-754 32-bit floating-point numbers
float64     // the set of all IEEE-754 64-bit floating-point numbers

complex64   // the set of all complex numbers with float32 real and imaginary parts
complex128  // the set of all complex numbers with float64 real and imaginary parts

byte        // alias for uint8
rune        // alias for int32
```

The value of an _n_-bit integer is n bits wide and represented using [two's complement arithmetic](https://en.wikipedia.org/wiki/Two's_complement).

There is also a set of predeclared integer types with implementation-specific sizes:

```go
uint     // either 32 or 64 bits
int      // same size as uint
uintptr  // an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

To avoid portability issues all numeric types are defined types and thus distinct except _byte_, which is an [alias]() for _uint8_, and _rune_, which is an alias for _int32_. Explicit conversions are required when different numeric types are mixed in an expression or assignment. For instance, _int32_ and _int_ are not the same type even though they may have the same size on a particular architecture.

TODO:

```go
bigint  // TODO
bigrat  // TODO
```

### String types

A _string type_ represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is `string`; it is a defined type.

```go
string
```

The length of a string `s` can be discovered using the built-in function [len](). The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer [indices]() `0` through `len(s)-1`. It is illegal to take the address of such an element; if `s[i]` is the i'th byte of a string, `&s[i]` is invalid.


### Array types

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.

```go
[N]T
```

The length is part of the array's type; it must evaluate to a non-negative [constant]() [representable]() by a value of type int. The length of array `a` can be discovered using the built-in function [len](). The elements can be addressed by integer [indices]() `0` through `len(a)-1`. Array types are always one-dimensional but may be composed to form multi-dimensional types.

```go
[32]byte
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

### Pointer types

A _pointer_ type denotes the set of all pointers to [variables]() of a given type, called the base type of the pointer. The value of an uninitialized pointer is `nil`.

```go
*T
```

For example:

```go
*Point
*[4]int
```

### Slice types

A _slice_ is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is `nil`.

```go
[]T
```

The length of a slice `s` can be discovered by the built-in function [len](); unlike with arrays it may change during execution. The elements can be addressed by integer [indices]() `0` through `len(s)-1`. The slice index of a given element may be less than the index of the same element in the underlying array.

A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

The array underlying `a` slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by [slicing]() a new one from the original slice. The capacity of a slice a can be discovered using the built-in function `cap(a)`.

A new, initialized slice value for a given element type `T` may be made using the built-in function [make](), which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with make always allocates a new, hidden array to which the returned slice value refers. That is, executing

```go
make([]T, length, capacity)
```

produces the same slice as allocating an array and [slicing]() it, so these two expressions are equivalent:

```
make([]int, 50, 100)
new([100]int)[0:50]
```

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.

### Map types

A _map_ is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is `nil`.

```go
map[KeyT]ElemT
```

The comparison operators `==` and `!=` must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a run-time panic.

```go
map[string]int
map[*T]string
map[string]any
```

The number of map elements is called its length. For a map `m`, it can be discovered using the built-in function [len]() and may change during execution. Elements may be added during execution using [assignments]() and retrieved with [index expressions](); they may be removed with the [delete]() and [clear]() built-in function.

A new, empty map value is made using the built-in function [make](), which takes the map type and an optional capacity hint as arguments:

```go
make(map[string]int)
make(map[string]int, 100)
```

The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of nil maps. A nil map is equivalent to an empty map except that no elements may be added.


### Function types

A _function_ type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is `nil`.

```go
func(parameters) results
```

Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-[blank]() names in the signature must be [unique](). If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

The final incoming parameter in a function signature may have a type prefixed with `...`. A function with such a parameter is called _variadic_ and may be invoked with zero or more arguments for that parameter.

```go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...any) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

### Interface types

#### Built-in interfaces

TODO:

```go
error
any
```

### Classes

TODO (classfile)


## Expressions

### Commands and calls

TODO

```go
echo "Hello world"
echo("Hello world")
```

### Built-in functions

TODO

### Operators

Operators combine operands into expressions.

Binary operators:

```go
|| && == != < <= > >=
+ - * / %
| & ^ &^ << >>
```

Unary operators:

```go
+ - ! ^ * &
```

#### Operator precedence

_Unary operators_ have the highest precedence. As the ++ and -- operators form statements, not expressions, they fall outside the operator hierarchy. As a consequence, statement *p++ is the same as (*p)++.

There are five precedence levels for _binary operators_. Multiplication operators bind strongest, followed by addition operators, comparison operators, && (logical AND), and finally || (logical OR):

```
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

Binary operators of the same precedence associate from left to right. For instance, `x / y * z` is the same as `(x / y) * z`.

```go
+x                         // x
42 + a - b                 // (42 + a) - b
23 + 3*x[i]                // 23 + (3 * x[i])
x <= f()                   // x <= f()
^a >> b                    // (^a) >> b
f() || g()                 // f() || g()
x == y+1 && <-chanInt > 0  // (x == (y+1)) && ((<-chanInt) > 0)
```

#### Arithmetic operators

_Arithmetic operators_ apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (+, -, *, /) apply to [integer](), [floating-point](), [rational]() and [complex]() types; + also applies to [strings](). The bitwise logical and shift operators apply to integers only.

```
+    sum                    integers (including bigint), floats, bigrat, complex values, strings
-    difference             integers (including bigint), floats, bigrat, complex values
*    product                integers (including bigint), floats, bigrat, complex values
/    quotient               integers (including bigint), floats, bigrat, complex values
%    remainder              integers (including bigint)

&    bitwise AND            integers (including bigint)
|    bitwise OR             integers (including bigint)
^    bitwise XOR            integers (including bigint)
&^   bit clear (AND NOT)    integers (including bigint)

<<   left shift             integer << integer >= 0
>>   right shift            integer >> integer >= 0
```

TODO

#### Comparison operators

_Comparison operators_ compare two operands and yield an untyped boolean value.

```go
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

In any comparison, the first operand must be [assignable]() to the type of the second operand, or vice versa.

The equality operators == and != apply to operands of comparable types. The ordering operators <, <=, >, and >= apply to operands of ordered types. These terms and the result of the comparisons are defined as follows:

* Boolean types are comparable. Two boolean values are equal if they are either both true or both false.
* Integer types are comparable and ordered. Two integer values are compared in the usual way.
* Floating-point types are comparable and ordered. Two floating-point values are compared as defined by the IEEE-754 standard.
* Complex types are comparable. Two complex values u and v are equal if both real(u) == real(v) and imag(u) == imag(v).
* String types are comparable and ordered. Two string values are compared lexically byte-wise.
* Pointer types are comparable. Two pointer values are equal if they point to the same variable or if both have value `nil`. Pointers to distinct [zero-size]() variables may or may not be equal.
* Interface types are comparable. Two interface values are equal if they have [identical]() dynamic types and equal dynamic values or if both have value `nil`.
* A value x of non-interface type X and a value t of interface type T can be compared if type X is comparable and X [implements]() T. They are equal if t's dynamic type is identical to X and t's dynamic value is equal to x.
* Array types are comparable if their array element types are comparable. Two array values are equal if their corresponding element values are equal. The elements are compared in ascending index order, and comparison stops as soon as two element values differ (or all elements have been compared).

A comparison of two interface values with identical dynamic types causes a [run-time panic]() if that type is not comparable. This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.

Slice, map, and function types are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`. Comparison of pointer, channel, and interface values to `nil` is also allowed and follows from the general rules above.

#### Logical operators

Logical operators apply to [boolean]() values and yield a result of the same type as the operands. The left operand is evaluated, and then the right if the condition requires it.

```
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

#### Address operators

For an operand x of type T, the address operation &x generates a pointer of type *T to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, x may also be a (possibly parenthesized) [composite literal](). If the evaluation of x would cause a [run-time panic](), then the evaluation of &x does too.

For an operand x of pointer type *T, the pointer indirection *x denotes the [variable]() of type T pointed to by x. If x is nil, an attempt to evaluate *x will cause a [run-time panic]().

```go
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

#### Conversions

A _conversion_ changes the [type](#types) of an expression to the type specified by the conversion. A conversion may appear literally in the source, or it may be _implied_ by the context in which an expression appears.

An _explicit conversion_ is an expression of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type `T`.

```go
T(x)
```

If the type starts with the operator * or <-, or if the type starts with the keyword func and has no result list, it must be parenthesized when necessary to avoid ambiguity:

```go
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
```

A [constant]() value `x` can be converted to type `T` if `x` is [representable]() by a value of `T`. As a special case, an integer constant `x` can be explicitly converted to a [string type]() using the [same rule]() as for non-constant `x`.

Converting a constant to a type yields a typed constant.

```go
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
myString("foo" + "bar")  // "foobar" of type myString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
```

##### Conversions between numeric types

For the conversion of non-constant numeric values, the following rules apply:

* When converting between [integer types](#numeric-types), if the value is a signed integer, it is sign extended to implicit infinite precision; otherwise it is zero extended. It is then truncated to fit in the result type's size. For example, if v := uint16(0x10F0), then uint32(int8(v)) == 0xFFFFFFF0. The conversion always yields a valid value; there is no indication of overflow.
* When converting a [floating-point number](#numeric-types) to an integer, the fraction is discarded (truncation towards zero).
* When converting an integer or floating-point number to a floating-point type, or a [complex number](#numeric-types) to another complex type, the result value is rounded to the precision specified by the destination type. For instance, the value of a variable x of type float32 may be stored using additional precision beyond that of an IEEE-754 32-bit number, but float32(x) represents the result of rounding x's value to 32-bit precision. Similarly, x + 0.1 may use more than 32 bits of precision, but float32(x + 0.1) does not.

In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.

##### Conversions to and from a string type

TODO

##### Conversions from slice to array or array pointer

TODO

### Constant expressions

Constant expressions may contain only [constant](#constants) operands and are evaluated at compile time.

Untyped boolean, numeric, and string constants may be used as operands wherever it is legal to use an operand of boolean, numeric, or string type, respectively.

A constant [comparison](#comparison-operators) always yields an untyped boolean constant. If the left operand of a constant [shift expression](#operators) is an untyped constant, the result is an integer constant; otherwise it is a constant of the same type as the left operand, which must be of [integer type](#numeric-types).

Any other operation on untyped constants results in an untyped constant of the same kind; that is, a boolean, integer, floating-point, complex, or string constant. If the untyped operands of a binary operation (other than a shift) are of different kinds, the result is of the operand's kind that appears later in this list: integer, rune, floating-point, complex. For example, an untyped integer constant divided by an untyped complex constant yields an untyped complex constant.

```go
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
```

Applying the built-in function `complex` to untyped integer, rune, or floating-point constants yields an untyped complex constant.

```go
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)
```

Constant expressions are always evaluated exactly; intermediate values and the constants themselves may require precision significantly larger than supported by any predeclared type in the language. The following are legal declarations:

```go
const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge >> 98  // Four == 4                                (type int8)
```

The divisor of a constant division or remainder operation must not be zero:

```go
3.14 / 0.0   // illegal: division by zero
```

The values of typed constants must always be accurately [representable]() by values of the constant type. The following constant expressions are illegal:

```go
uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)
```

The mask used by the unary bitwise complement operator ^ matches the rule for non-constants: the mask is all 1s for unsigned constants and -1 for signed and untyped constants.

```go
^1         // untyped integer constant, equal to -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2
```

### Short variable declarations

A short variable declaration uses the syntax:

```go
varName[, ...] = expression[, ...]
```

It is shorthand for a regular [variable declaration]() with initializer expressions but no types:

```go
var varName[, ...] = expression[, ...]
```

For example:

```go
i, j := 0, 10
f := func() int { return 7 }
ints := make([]int)
r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
_, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate
```

Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-[blank]() variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original. The non-blank variable names on the left side of := must be [unique]().

```go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
x, y, x := 1, 2, 3                        // illegal: x repeated on left side of :=
```

Short variable declarations may appear only inside functions. In some contexts such as the initializers for "[if]()", "[for]()", or "[switch]()" statements, they can be used to declare local temporary variables.


### Slice literals

TODO

```go
[expression1, ...]
```

For example:

```go
[]                   // []any
[1, 2, 3]            // []int
[10, 3.14, 200]      // []float64
["Hello", "world"]   // []string
["Hello", 100, true] // []any
```

The type of slice literals can be inferred from the context:

```go
func echoF32s(vals []float32) {
	echo vals
}

echo [10, 3.14, 200]           // []float64
echoF32s [10, 3.14, 200]       // []float32

var a []any = [10, 3.14, 200]  // []any
echo a
```

### Map literals

TODO

```go
{key1: value1, ...}
```

For example:

```go
{}                           // map[string]any
{"Monday": 1, "Sunday": 7}   // map[string]int
{1: 100, 3: 3.14, 5: 10}     // map[int]float64
```

The type of map literals can be inferred from the context:

```go
func echoS2f32(vals map[string]float32) {
	echo vals
}

echo {"Monday": 1, "Sunday": 7}
echoS2f32 {"Monday": 1, "Sunday": 7}

var a map[string]any = {"Monday": 1, "Sunday": 7}
echo a
```

## Statements

Statements control execution.

```go
Statement =
	Declaration | SimpleStmt | IfStmt | ForStmt | SwitchStmt |
    LabeledStmt | BreakStmt | ContinueStmt | FallthroughStmt | GotoStmt |
	ReturnStmt | DeferStmt | Block .

SimpleStmt = EmptyStmt | ExpressionStmt | IncDecStmt | Assignment | ShortVarDecl .
```

### Empty statements

The empty statement does nothing.

```go
EmptyStmt = .
```

### Expression statements

With the exception of specific built-in functions, function and method [calls](#commands-and-calls) and [receive operations]() can appear in statement context. Such statements may be parenthesized.

```go
ExpressionStmt = Expression .
```

The following built-in functions are not permitted in statement context:

```go
append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice unsafe.SliceData unsafe.String unsafe.StringData
```

For example:

```go
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

### IncDec statements

The "++" and "--" statements increment or decrement their operands by the untyped constant 1. As with an assignment, the operand must be addressable or a map index expression.

```go
IncDecStmt = Expression ( "++" | "--" ) .
```

The following [assignment statements]() are semantically equivalent:

```go
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

### Assignment statements

An assignment replaces the current value stored in a [variable](#variables) with a new value specified by an [expression](#expressions). An assignment statement may assign a single value to a single variable, or multiple values to a matching number of variables.

```go
Assignment = ExpressionList assign_op ExpressionList .
ExpressionList = Expression { "," Expression } .
```

Here `assign_op` can be:

```go
= += -= |= ^= *= /= %= <<= >>= &= &^=
```

Each left-hand side operand must be [addressable](#address-operators), a map index expression, or (for = assignments only) the [blank identifier](). Operands may be parenthesized.

```go
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

An _assignment operation_ x _op=_ y where op is a binary [arithmetic operator](#arithmetic-operators) is equivalent to x = x op (y) but evaluates x only once. The op= construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier.

```go
a[i] <<= 2
i &^= 1<<n
```

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a channel or [map](#map-types) operation, or a [type assertion](). The number of operands on the left hand side must match the number of values. For instance, if f is a function returning two values,

```go
x, y = f()
```

assigns the first value to x and the second to y. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the nth expression on the right is assigned to the nth operand on the left:

```go
one, two, three = '一', '二', '三'
```

The [blank identifier]() provides a way to ignore right-hand side values in an assignment:

```go
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
```

The assignment proceeds in two phases. First, the operands of [index expressions]() and [pointer indirections]() (including implicit pointer indirections in selectors) on the left and the expressions on the right are all [evaluated in the usual order](). Second, the assignments are carried out in left-to-right order.

```go
a, b = b, a  // exchange a and b

x := [1, 2, 3]
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

i = 2
x = [3, 5, 7]
for i, x[i] <- x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x is [3, 5, 3]
```

In assignments, each value must be [assignable]() to the type of the operand to which it is assigned, with the following special cases:

* Any typed value may be assigned to the blank identifier.
* If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first implicitly [converted](#conversions) to its [default type](#constants).
* If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first implicitly converted to type bool.


### If statements

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.

```go
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
```

For example:

```go
if x > 1 {
	x = 1
}
```

The expression may be preceded by a simple statement, which executes before the expression is evaluated.

```go
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

### For statements

A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.

```go
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
```

#### For statements with single condition

In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value true.

```go
for a < b {
	a *= 2
}
```

#### For statements with for clause

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement. The init statement may be a [short variable declaration](#short-variable-declarations), but the post statement must not.

```go
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
```

For example:

```go
for i := 0; i < 10; i++ {
	f(i)
}
```

If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the [semicolons]() are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value true.

```go
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

Each iteration has its own separate declared variable (or variables) [Go 1.22](). The variable used by the first iteration is declared by the init statement. The variable used by each subsequent iteration is declared implicitly before executing the post statement and initialized to the value of the previous iteration's variable at that moment.

```go
var prints []func()
for i := 0; i < 5; i++ {
	prints = append(prints, func() { println(i) })
	i++
}
for _, p := range prints {
	p()
}
```

prints

```
1
3
5
```

Prior to [Go 1.22], iterations share one set of variables instead of having their own separate variables. In that case, the example above prints

```
6
6
6
```

#### For statements with range clause

TODO

### Switch statements

"Switch" statements provide multi-way execution. An expression or type is compared to the "cases" inside the "switch" to determine which branch to execute.

```go
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
```

There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.

#### Expression switches

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value true.

```go
ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
```

If the switch expression evaluates to an untyped constant, it is first implicitly [converted](#conversions) to its [default type](#constants). The predeclared untyped value nil cannot be used as a switch expression. The switch expression type must be [comparable](#comparison-operators).

If a case expression is untyped, it is first implicitly [converted](#conversions) to the type of the switch expression. For each (possibly converted) case expression x and the value t of the switch expression, x == t must be a valid [comparison](#comparison-operators).

In other words, the switch expression is treated as if it were used to declare and initialize a temporary variable t without explicit type; it is that value of t against which each case expression x is tested for equality.

In a case or default clause, the last non-empty statement may be a (possibly [labeled](#labeled-statements)) ["fallthrough" statement]() to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

The switch expression may be preceded by a simple statement, which executes before the expression is evaluated.

```go
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

Implementation restriction: A compiler may disallow multiple case expressions evaluating to the same constant. For instance, the current compilers disallow duplicate integer, floating point, or string constants in case expressions.

#### Type switches

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a [type assertion]() using the keyword type rather than an actual type:

```go
switch x.(type) {
// cases
}
```

Cases then match actual types T against the dynamic type of the expression x. As with type assertions, x must be of interface type, but not a type parameter, and each non-interface type T listed in a case must implement the type of x. The types listed in the cases of a type switch must all be different.

```go
TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
```

The TypeSwitchGuard may include a [short variable declaration](#short-variable-declarations). When that form is used, the variable is declared at the end of the TypeSwitchCase in the [implicit block]() of each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.

Instead of a type, a case may use the predeclared identifier [nil](); that case is selected when the expression in the TypeSwitchGuard is a `nil` interface value. There may be at most one `nil` case.

Given an expression x of type `any`, the following type switch:

```
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (any)
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (any)
default:
	printString("don't know the type")     // type of i is type of x (any)
}
```

could be rewritten:

```go
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (any)
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (any)
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (any)
		printString("don't know the type")
	}
}
```

The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.

The "fallthrough" statement is not permitted in a type switch.


### Labeled statements

A labeled statement may be the target of a goto, break or continue statement.

```go
LabeledStmt = Label ":" Statement .
Label       = identifier .
```

For example:

```go
Error:
	log.Panic("error encountered")
```

### Break statements

A "break" statement terminates execution of the innermost "[for](#for-statements)" or "[switch](#switch-statements)" statement within the same function.

```go
BreakStmt = "break" [ Label ] .
```

If there is a label, it must be that of an enclosing "for" or "switch" statement, and that is the one whose execution terminates.

```go
OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
```

### Continue statements

A "continue" statement begins the next iteration of the innermost enclosing "[for](#for-statements)" loop by advancing control to the end of the loop block. The "for" loop must be within the same function.

```go
ContinueStmt = "continue" [ Label ] .
```

If there is a label, it must be that of an enclosing "for" statement, and that is the one whose execution advances.

```go
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
```

### Fallthrough statements

A "fallthrough" statement transfers control to the first statement of the next case clause in an expression "[switch](#switch-statements)" statement. It may be used only as the final non-empty statement in such a clause.

```go
FallthroughStmt = "fallthrough" .
```

### Goto statements

A "goto" statement transfers control to the statement with the corresponding label within the same function.

```go
GotoStmt = "goto" Label .
```

For example:

```go
goto Error
```

Executing the "goto" statement must not cause any variables to come into [scope]() that were not already in scope at the point of the goto. For instance, this example:

```go
	goto L  // BAD
	v := 3
L:
```

is erroneous because the jump to label L skips the creation of v.

A "goto" statement outside a [block]() cannot jump to a label inside that block. For instance, this example:

```go
if n%2 == 1 {
	goto L1  // BAD
}
for n > 0 {
	f()
	n--
L1:
	f()
	n--
}
```

is erroneous because the label L1 is inside the "for" statement's block but the goto is not.

### Return statements

A "return" statement in a function F terminates the execution of F, and optionally provides one or more result values. Any functions [deferred](#defer-statements) by F are executed before F returns to its caller.

```go
ReturnStmt = "return" [ ExpressionList ] .
```

In a function without a result type, a "return" statement must not specify any result values.

```go
func noResult() {
	return
}
```

There are three ways to return values from a function with a result type:

* The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and [assignable]() to the corresponding element of the function's result type.

```go
func simpleF() int {
	return 2
}

func complexF1() (re float64, im float64) {
	return -7.0, -4.0
}
```

* The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply.

```go
func complexF2() (re float64, im float64) {
	return complexF1()
}
```

* The expression list may be empty if the function's result type specifies names for its [result parameters](). The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables.

```go
func complexF3() (re float64, im float64) {
	re = 7.0
	im = 4.0
	return
}

func (devnull) Write(p []byte) (n int, _ error) {
	n = len(p)
	return
}
```

Regardless of how they are declared, all the result values are initialized to the [zero values]() for their type upon entry to the function. A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

Implementation restriction: A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) with the same name as a result parameter is in [scope]() at the place of the return.

```go
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
```

### Defer statements

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a [return statement](#return-statements), reached the end of its [function body](), or because the corresponding goroutine is [panicking]().

```go
DeferStmt = "defer" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](#expression-statements).

Each time a "defer" statement executes, the function value and parameters to the call are [evaluated as usual]() and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. That is, if the surrounding function returns through an explicit [return statement](#return-statements), deferred functions are executed after any result parameters are set by that return statement but before the function returns to its caller. If a deferred function value evaluates to nil, execution [panics]() when the function is invoked, not when the "defer" statement is executed.

For instance, if the deferred function is a [function literal]() and the surrounding function has [named result parameters]() that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. (See also the section on [handling panics]().)

```go
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
```

### Terminating statements

TODO


## Built-in functions

Built-in functions are [predeclared](). They are called like any other function but some of them accept a type instead of an expression as the first argument.

The built-in functions do not have standard Go types, so they can only appear in [call expressions](#commands-and-calls); they cannot be used as function values.

### Appending to and copying slices

The built-in functions `append` and `copy` assist in common slice operations. For both functions, the result is independent of whether the memory referenced by the arguments overlaps.

The [variadic](#function-types) function append appends zero or more values x to a slice s and returns the resulting slice of the same type as s. The [core type]() of s must be a slice of type `[]E`. The values x are passed to a parameter of type `...E` and the respective [parameter passing rules]() apply. As a special case, if the core type of s is []byte, append also accepts a second argument with core type [bytestring]() followed by `...`. This form appends the bytes of the byte slice or string.

```go
append(s S, x ...E) S  // core type of S is []E
```

If the capacity of s is not large enough to fit the additional values, append [allocates]() a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array.

```go
s0 := [0, 0]
s1 := append(s0, 2)                // append a single element     s1 is [0, 0, 2]
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 is [0, 0, 2, 3, 5, 7]
s3 := append(s2, s0...)            // append a slice              s3 is [0, 0, 2, 3, 5, 7, 0, 0]
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 is [3, 5, 7, 2, 3, 5, 7, 0, 0]

var t []any
t = append(t, 42, 3.1415, "foo")   //                             t is [42, 3.1415, "foo"]

var b []byte
b = append(b, "bar"...)            // append string contents      b is []byte("bar")
```

The function copy copies slice elements from a source src to a destination dst and returns the number of elements copied. The [core types]() of both arguments must be slices with [identical]() element type. The number of elements copied is the minimum of `len(src)` and `len(dst)`. As a special case, if the destination's core type is `[]byte`, copy also accepts a source argument with core type [bytestring](). This form copies the bytes from the byte slice or string into the byte slice.

```go
copy(dst, src []T) int
copy(dst []byte, src string) int
```

Examples:

```go
a := [0, 1, 2, 3, 4, 5, 6, 7]
s := make([]int, 6)
b := make([]byte, 5)
n1 := copy(s, a)                // n1 == 6, s is []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s is []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b is []byte("Hello")
```

### Clear

The built-in function `clear` takes an argument of `map` or `slice` and deletes or zeroes out all elements.

```
Call        Argument type     Result

clear(m)    map[K]T           deletes all entries, resulting in an
                              empty map (len(m) == 0)

clear(s)    []T               sets all elements up to the length of
                              s to the zero value of T
```

If the map or slice is `nil`, `clear` is a no-op.

### Manipulating complex numbers

TODO
