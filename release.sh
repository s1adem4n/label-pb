make -j4 all
TAG=$(git describe --tags --abbrev=0)
echo "Tag: $TAG"

zip -r releases/${TAG}.zip build/
gh release create $TAG releases/${TAG}.zip -t $TAG