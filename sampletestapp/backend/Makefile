moduleowner = github.com/elmunyeco/
packagename = sampletestapp
containercli = podman

# Build And Development
init:
	@ go mod init $(moduleowner)$(packagename)
	@ go mod vendor 
update:
	@ go mod vendor
clean:
	@ sudo rm -rf $(packagename).bin $(packagename).exe cover.txt cover.html cover.out build
run:
	@ go build -o ./build/$(packagename).bin $(moduleowner)$(packagename)/cmd/$(packagename)  && ./build/$(packagename).bin	 
image-build: 
	@ $(containercli) build -t elmunyeco/$(packagename):latest .
image-push: 
	@ make image-build && $(containercli) push elmunyeco/$(packagename):latest
container-run: 
	@ $(containercli) run -p 8080:8080 elmunyeco/$(packagename):latest
