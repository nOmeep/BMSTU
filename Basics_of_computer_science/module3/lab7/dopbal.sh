#!/bin/bash

#/home/sevastyan/* - запускаю эту

find $1 -type f -name *.[ch] | grep -v '\.svn' | xargs cat | sed '/^\s*$/d' | wc -l

