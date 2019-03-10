#!/bin/bash

# Run examples/ from google/jsonnet/ and compare output to golden files.

MAIN="/tmp/$$.jsonnet_main"
OUT="/tmp/$$.out"
TRUE="/tmp/$$.true.golden"
go build -o "$MAIN" jsonnet_main/main.go
echo true > "$TRUE"

E=../../google/jsonnet/examples/
S=../../google/jsonnet/test_suite/
STATUS=0
for x in "$E"/*.jsonnet "$E"/terraform/*.jsonnet "$S"/*.jsonnet
do
	case $x in 
	  */top-level-*.jsonnet )
	    # These top-level things don't have a printable value.
	    continue
	    ;;
	  */error.* )
	    echo "== $x =="
	    if go run jsonnet_main/main.go "$x" >$OUT 2>&1
	    then
	        cat -nv $OUT
	    	echo "FAILED: That should have failed."
		STATUS=7
	    fi
	    ;;
          */stdlib.jsonnet | */tla.simple.jsonnet )
	    continue
	    ;;
	  * )
	    if test -f "$x.golden"
	    then
	      echo "== $x =="
	      (cd $(dirname "$x") ; $MAIN $(basename "$x")) >$OUT 2>&1
	      if diff -b "$x.golden" "$OUT"
	      then
	        : ok good
	      else
	        STATUS=$?
	        echo "FAILED: $x"
	      fi
	    else
	      echo "== $x =="
	      (cd $(dirname "$x") ; $MAIN $(basename "$x")) >$OUT 2>&1
	      if diff -b "$TRUE" "$OUT"
	      then
	        : ok good
	      else
	        STATUS=$?
	        echo "FAILED: $x"
	      fi
	    fi
	    ;;
	esac
done

case $STATUS in
  0 ) echo 'ALL OKAY' >&2 ;;
  * ) echo 'FAILED' >&2 ;;
esac

rm -f "$MAIN" "$TRUE" "$OUT"
exit $STATUS
