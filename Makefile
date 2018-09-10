# Makefile for generating the cross platform binaries and tarballs

DIST=distributions

all:
	@echo "Building label-hook..."; 
	go get && go build .

.PHONY: release
release: distribution \
	label-hook_darwin_386 \
	label-hook_darwin_amd64 \
	label-hook_freebsd_386 \
	label-hook_freebsd_amd64 \
	label-hook_freebsd_arm \
	label-hook_linux_386 \
	label-hook_linux_amd64 \
	label-hook_linux_arm \
	label-hook_linux_arm64 \
	label-hook_netbsd_386 \
	label-hook_netbsd_amd64 \
	label-hook_netbsd_arm \
	label-hook_openbsd_386 \
	label-hook_openbsd_amd64 \
	label-hook_openbsd_arm \
	label-hook_windows_386 \
	label-hook_windows_amd64 \
	cleanup

distribution:
	@mkdir -p distributions; \
	rm -f ${DIST}/*

cleanup:
	rm -rf executables

label-hook_darwin_386:
	$(eval name=label-hook_darwin_386)
	GOOS=darwin GOARCH=386 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_darwin_amd64:
	$(eval name=label-hook_darwin_amd64)
	GOOS=darwin GOARCH=amd64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_freebsd_386:
	$(eval name=label-hook_freebsd_386)
	GOOS=freebsd GOARCH=386 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_freebsd_amd64:
	$(eval name=label-hook_freebsd_amd64)
	GOOS=freebsd GOARCH=amd64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_freebsd_arm:
	$(eval name=label-hook_freebsd_arm)
	GOOS=freebsd GOARCH=arm go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_linux_386:
	$(eval name=label-hook_linux_386)
	GOOS=linux GOARCH=386 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_linux_arm:
	$(eval name=label-hook_linux_arm)
	GOOS=linux GOARCH=arm go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_linux_arm64:
	$(eval name=label-hook_linux_arm64)
	GOOS=linux GOARCH=arm64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_netbsd_386:
	$(eval name=label-hook_netbsd_386)
	GOOS=netbsd GOARCH=386 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_netbsd_amd64:
	$(eval name=label-hook_netbsd_amd64)
	GOOS=netbsd GOARCH=amd64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_netbsd_arm:
	$(eval name=label-hook_netbsd_arm)
	GOOS=netbsd GOARCH=arm go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_openbsd_386:
	$(eval name=label-hook_openbsd_386)
	GOOS=openbsd GOARCH=386 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_openbsd_amd64:
	$(eval name=label-hook_openbsd_amd64)
	GOOS=openbsd GOARCH=amd64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_openbsd_arm:
	$(eval name=label-hook_openbsd_arm)
	GOOS=openbsd GOARCH=arm go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_windows_386:
	$(eval name=label-hook_windows_386)
	GOOS=windows GOARCH=386 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}

label-hook_windows_amd64:
	$(eval name=label-hook_windows_amd64)
	GOOS=windows GOARCH=amd64 go build -o executables/${name} .
	tar -czvf ${DIST}/${name}.tar.gz -C executables/ ${name}