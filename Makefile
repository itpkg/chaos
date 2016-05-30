dst=dist

dist:
	mkdir -pv ${dst}
	cd app && gradle build
	cd front-emberjs && ember build --env production
	cp -rv app/build/libs/itpkg-*.jar app/config $(dst)
	cp -rv front-emberjs/dist $(dst)/public



clean:
	gradle clean
	-rm -r $(dst) front-emberjs/dist




