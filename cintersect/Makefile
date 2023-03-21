.PHONY: clean test

all: test

intersections: intersection.c
	$(CC) $^ -o $@

test: intersections golden.txt
	./intersections < ../samples/fromTest.txt > actual.txt
	diff golden.txt actual.txt
