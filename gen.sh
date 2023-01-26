# #! /bin/env bash

# SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

# FILE=db.sqlite

# REGEN=0
# while getopts "r" opt; do
# 	case "$opt" in
# 		r) REGEN=1
# 	esac
# done

# popd &> /dev/null

# if [[ "$REGEN" = "1" ]]; then
# 	echo "Regenerating database..."
# 	if [[ -f $FILE ]]; then
# 		(set -x; rm $FILE)
# 	fi
# 	(set -x; sqlite3 $FILE < $SRC/sql/schema.sql)
# fi

rm -rf db.sqlite && sqlite3 db.sqlite < sql/schema.sql