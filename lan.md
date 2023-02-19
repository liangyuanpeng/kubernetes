
KUBE_BUILD_PLATFORMS=linux/amd64 make all WHAT=cmd/kube-apiserver GOFLAGS=-v GOGCFLAGS="-N -l"


KUBE_BUILD_PLATFORMS=linux/amd64 WHAT=cmd/kube-apiserver GOFLAGS=-v GOGCFLAGS="-N -l" make quick-release 


KUBE_BUILD_PLATFORMS=linux/amd64  make quick-release  WHAT=cmd/kube-apiserver GOFLAGS=-v GOGCFLAGS="-N -l"


KUBE_BUILD_PLATFORMS=linux/amd64 WHAT=cmd/kube-apiserver GOFLAGS=-v GOGCFLAGS="-N -l" make release-images


KUBE_BUILD_PLATFORMS=linux/amd64  make release-images WHAT=cmd/kube-apiserver GOFLAGS=-v GOGCFLAGS="-N -l"

