export GOPATH=$(pwd)

# Build the front-end application

yarn --cwd ../front run build

# Copy files from front/dist to public

rm -rf ./public
mkdir public
cp -r ../front/dist/* ./public/

# Generate assets from front-end build

./bin/statik -f

go build -o app

cp app ../bin
