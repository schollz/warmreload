warmreload:
	rm -rf /home/we/dust/code/warmreload/warmreload
	cd /home/we/dust/code/warmreload && wget https://github.com/schollz/warmreload/releases/download/v0.1.0/warmreload
	chmod +x /home/we/dust/code/warmreload/warmreload

build:
	docker build --progress=plain -t warmreload .

clean:
	rm -rf warmreload
