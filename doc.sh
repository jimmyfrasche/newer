#!/bin/sh
#This isn't necessary since mango and autoreadme are essentially instantaneous,
#but if nothing else it's a reasonable example
trim() {
	<README.md.template sed -n "$@" | grep -v %USAGE
}
if newer newer.go README.md.template doc.sh $(which newer) $(which autoreadme) $(which mango) -- newer.1 README.md
then
	mango >newer.1

	(
		trim "1,/%USAGE%/p"
		newer 2>&1 | sed -n "2,\$p"
		trim "/%USAGE%/,\$p"
	) >template
	autoreadme -f -template template
	rm template
fi
