all:
	go build -buildmode=c-shared -o lib/libton.so $$PWD/go
	
