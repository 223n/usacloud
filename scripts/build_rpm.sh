#!/bin/bash
set -e
set -x

SOURCE_DIR="package/rpm"
DESTINATION_DIR="package/rpm-build"

: "prepare rpm build..."
    rm -rf rpmbuild/
    mkdir -p rpmbuild/RPMS/{x86_64,noarch}
    rm -rf repos/centos
    mkdir -p repos/centos/{x86_64,noarch}
    rm -rf "${DESTINATION_DIR}"
    mkdir -p "${DESTINATION_DIR}/src"
    cp contrib/completion/bash/usacloud "${DESTINATION_DIR}/src/usacloud_bash_completion"
    cp "${SOURCE_DIR}/usacloud.spec" "${DESTINATION_DIR}/usacloud.spec"

: "building i386..."
    unzip -oq bin/usacloud_linux-386.zip -d bin/
	docker run --rm \
	    -v "$PWD":/workdir \
	    -v "$PWD/rpmbuild":/rpmbuild \
	    sacloud/usacloud:rpm-build \
	        --define "_sourcedir /workdir/package/rpm-build/src" \
	        --define "_builddir /workdir/bin" \
	        --define "_version ${CURRENT_VERSION}" \
	        --define "buildarch noarch" \
	        --target noarch \
	        -bb package/rpm-build/usacloud.spec

: "building x86_64..."
    unzip -oq bin/usacloud_linux-amd64.zip -d bin/
	docker run --rm \
	    -v "$PWD":/workdir \
	    -v "$PWD/rpmbuild":/rpmbuild \
	    sacloud/usacloud:rpm-build \
	        --define "_sourcedir /workdir/package/rpm-build/src" \
	        --define "_builddir /workdir/bin" \
	        --define "_version ${CURRENT_VERSION}" \
	        --define "buildarch x86_64" \
	        --target x86_64 \
	        -bb package/rpm-build/usacloud.spec

: "create yum repo..."
    cp -rf rpmbuild/RPMS/noarch/* repos/centos/noarch/
    cp -rf rpmbuild/RPMS/x86_64/* repos/centos/x86_64/
	docker run --rm \
	    -v "$PWD/repos/centos/noarch":/workdir \
	    --entrypoint createrepo \
	    sacloud/usacloud:rpm-build \
	        -v /workdir
	docker run --rm \
	    -v "$PWD/repos/centos/x86_64":/workdir \
	    --entrypoint createrepo \
	    sacloud/usacloud:rpm-build \
	        -v /workdir
