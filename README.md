#newer
Command newer(1) takes two lists of files and returns true if any of the files in the first list are newer than any of the files in the second list.

Download:
```shell
go get github.com/jimmyfrasche/newer
```
If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install) first
* * *
Usage:
```
Usage: newer left+ [-- right+]
	where left and right are the names of files

If -- is not included the last value in left becomes the value of right.
newer compares the modtimes of left to the modtimes of right,
and returns true if any file in right is newer than any file in left.
```

Command newer(1) takes two lists of files and returns true if any of the files
in the first list are newer than any of the files in the second list.

newer(1) is the essential logic of many build tools.
It is useful when you have a few processes that could benefit
from such logic, but are not quite ready to use a full build tool.

All files in both lists must exist.

EXAMPLES

```
newer a b && echo b is newer than a || echo a is newer than b

if newer a b
then
	rebuild-b-from a
fi
```

See if any of a, b, or c are newer than d

```
newer a b c d
```

See if a is newer than any of b, c, or d

```
newer a -- b c d
```

See if either a or b is newer than eitehr of c or d

```
newer a b -- c d
```

##EXIT STATUS
If there is a usage error or a file cannot be stated, exit status is 2.

If the second list contains an item newer than any item in the first,
exit status is 0.

Otherwise, the exit status is 1.



Another example can be found in doc.sh in this repository.

* * *
Automatically generated by [autoreadme](https://github.com/jimmyfrasche/autoreadme) on 2016.07.03
