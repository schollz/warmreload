warmreload:
	docker build --progress=plain -t warmreload .

clean:
	rm -rf warmreload
