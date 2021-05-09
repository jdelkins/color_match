#!/usr/bin/env zsh

for cs in ~/.zsh/themes/*; do
	sed -i -e '/_mycolor_ansi/d' $cs
	# loadcolorstyle is defined in my ~/.zshenv
	loadcolorstyle $(basename $cs)
	echo 'typeset -gxA _mycolor_ansi' >>$cs
	for i in `seq 0 255`; do
		./color_match -a $i -f $'_mycolor_ansi['$i$']="%s"\n' >>$cs
	done
done
