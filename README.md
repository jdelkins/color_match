# match_color

A small utility written in [go](http://golang.org) to help manage my personal
shell color preferences.

## Background and description

I have a fairly complicated and tightly-bound system for managing color themes.
This system revolves around a set of theme files that set a few color
environment variables, which I use in a wide array of app configurations to set
up consistent color schemes. I sometimes have a need to adapt various
configuration schemes to this one. For example: sometimes colors are given
with ansi 256 color indices. This kind of thing can involve finding the closest
match to a color from a smallish set of theme colors. This simple utility does
that using [CIELAB](https://en.wikipedia.org/wiki/CIELAB_color_space)
colorspace matching, which, I've learned, is the Right Way to do colorspace
distance calculations. The hard work (not that hard really, but why re-invent
the wheel) is done by the excellent
[go-colorful](https://github.com/lucasb-eyer/go-colorful) library; this utility
just presents a handy interface for my purposes.

The particular use case prompting this project is
[powerlevel10k](https://github.com/romkatv/powerlevel10k), which, though an
amazing suite of hacks, has, IMO, poorly-designed (if flexible) color spec
approach involving environment variables set to xterm 256 color numbers. Easy
enough to set those to my color theme variables, but which variable to use,
without getting in to re-designing the whole look? This tool can be used easily
to build a map of xterm color numbers to the closest matching color from my
color theme.

## How to use

Don't. The program is uselss unless using my particular color environment
variable scheme, namely a suite of varibles named`$mycolor_<colorname>` with
24-bit rgb hex triplet values of the form `#rrggbb`.

If you happen to be developing some kind of similar project, one
possibly-reusable bit of work here is that the program initializes a map of
color structs based on the 256 standard xterm ansi color numbers.

See `match_color --help` or read the source for usage guidance.

## Author

| Method    | Address        |
|-----------|----------------|
| Real life | Joel Elkins    |
| Email     | joel@elkins.co |
| Github    | @jdelkins      |

