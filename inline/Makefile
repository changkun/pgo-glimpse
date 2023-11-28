step1:
	go build -gcflags -m main.go
step2:
	go tool objdump -s "main.main" main | grep CALL
step3:
	go build -gcflags -m main2.go
step4:
	go tool objdump -s "main.main" main2 | grep CALL