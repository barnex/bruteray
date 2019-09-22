set -e

pre=github.com/barnex/bruteray-private/v2
export GOPATH=/home/arne/src

for pkg in api builder color geom light material texture tracer; do
	godoc2ghmd -goroot $GOPATH -ex $pre/$pkg > $pkg/README.md
done


