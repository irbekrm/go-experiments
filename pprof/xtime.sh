#!/bin/sh

# User time, system time, real time, memory usage
/usr/bin/time -f '%Uu %Ss %er %MkB %C' "$@"