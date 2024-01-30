TAGS=$(git tag --list)
echo "Tags:"
echo "$TAGS"

# check if there are uncommited changes
if [[ $(git status --porcelain) ]]; then
  echo "There are uncommited changes. Commit them first."
  exit 1
fi

TAG=$(git describe --tags --abbrev=0 2>/dev/null)
if [[ -z "$TAG" ]]; then
    echo "No tag found. Enter a tag:"
    read TAG
    git tag $TAG
else 
    echo "Tag found: $TAG"
    echo "Do you want to release this tag? (y/n)"
    read ANSWER
    if [[ $ANSWER =~ ^[Yy]$ ]]; then
        echo "Releasing tag $TAG"
    else
        echo "Enter a tag:"
        read TAG
        git tag $TAG
    fi
fi

# check if tag is already released
if [[ $(gh release list | grep $TAG) ]]; then
  echo "Tag $TAG already released"
  exit 1
fi

# check if tag is present in origin
if [[ ! $(git ls-remote --tags origin | grep $TAG) ]]; then
  git push --tags
fi

# check if zip does not exist
if [[ ! -f releases/${TAG}.zip ]] || [[ $@ =~ "-f" ]]; then
  rm -rf build
  rm -rf frontend/build

  make -j4 all
  zip -r releases/${TAG}.zip build/
fi

gh release create $TAG releases/${TAG}.zip