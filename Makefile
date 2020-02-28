build-pi:
	cd pi && GOOS=linux GOARCH=arm GOARM=5 go build -o ../CLOCK

deploy: build-pi
	scp ./CLOCK pi@192.168.0.52:/home/pi/CLOCK
