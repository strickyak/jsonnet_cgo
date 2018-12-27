#!/bin/bash

# To help me keep up-to-date with jsonnet,
# this will compare our copies of jsonnet files
# with the ones in a jsonnet directory.
# See "Usage:" a few lines below.

WRITE=
case "$1" in
	-w )
		WRITE=1
		shift
		;;
esac

case "$#/$1" in
	0/ )
		set ../../google/jsonnet/
		;;
	1/*/jsonnet/ )
		: ok
		;;
	* )
		echo >&2 '
Usage:
	sh  $0  ?-w?  ?/path/to/jsonnet/?

This command takes one argument, the jsonnet repository directory,
ending in /jsonnet/.  The default is ../../google/jsonnet/.

If -w is provided, then it will attempt to update files
with differences from the source.
'
		exit 13
		;;
esac

J="$1"
test -d "$J"

for x in \
	ast.h \
	desugarer.cpp \
	desugarer.h \
	formatter.cpp \
	formatter.h \
	json.h \
	json.hpp \
	lexer.cpp \
	lexer.h \
	libjsonnet.cpp \
	libjsonnet.h \
	libjsonnet_fmt.h \
	md5.cpp \
	md5.h \
	parser.cpp \
	parser.h \
	pass.cpp \
	pass.h \
	state.h \
	static_analysis.cpp \
	static_analysis.h \
	static_error.h \
	std.jsonnet.h \
	string_utils.cpp \
	string_utils.h \
	unicode.h \
	vm.cpp \
	vm.h \
	#
do
	ok=false
	F=
	for subdir in core cpp third_party/md5 third_party/json include
	do
		test -f "$J/$subdir/$x" && F="$J/$subdir/$x"
		if cmp "$J/$subdir/$x" "./$x" 2>/dev/null
		then
	    		ok=true
			break
		fi
	done

	if $ok
	then
		echo "ok: $x"
	else
		echo "******** NOT OK: $x"
		test -n "$WRITE" && cp -v "$F" "./$x"
	fi
done
