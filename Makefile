goes-boot:
	env -u CFLAGS -u CPPFLAGS -u CXXFLAGS goes-build -x -z -v coreboot-platina-mk1.rom coreboot-example-amd64.rom goes-boot goes-boot-platina-mk1

install:
#	$(INSTALL) -m 0644 -d $(DESTDIR)/boot/goes
#	$(INSTALL) goes-boot $(DESTDIR)/boot/goes

clean:
	rm -f *.rom debian/debhelper-build-stamp debian/files debian/*.substvars *.vmlinuz *.xz goes-boot goes-boot-platina-mk1 goes-bootrom goes-bootrom-platina-mk1
	rm -rf debian/.debhelper debian/goes-boot-mk1 debian/goes-boot-example-amd64

binpkg-deb:
	debuild -i -I -Iworktrees --lintian-opts --profile debian

.PHONY: goes-boot install binpkg-deb clean
