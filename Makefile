warmreload:
	curl https://github.com/schollz/warmreload/releases/download/v0.1.0/warmreload > /home/we/dust/code/warmreload/warmreload
	chmod +x /home/we/dust/code/warmreload/warmreload

build:
	docker build --progress=plain -t warmreload .

clean:
	rm -rf warmreload
