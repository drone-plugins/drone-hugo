.PHONY: all test push build release clean

README_TEMPLATE=./docs/tmpl.md

all: release

test: Dockerfile drone-hugo.sh
	docker build -t "cbrgm/drone-hugo:$(hugo)_test" --build-arg HUGO_VERSION="$(hugo)" .

build: Dockerfile drone-hugo.sh
	docker build -t "cbrgm/drone-hugo:$(hugo)" --build-arg HUGO_VERSION="$(hugo)" .
	docker build -t "cbrgm/drone-hugo:latest" --build-arg HUGO_VERSION="$(hugo)" .

push: build
	docker push "cbrgm/drone-hugo:$(hugo)"
	docker push "cbrgm/drone-hugo:latest"

release: $(README_TEMPLATE) test build push clean
	sed 's/<HUGO_VERSION>/$(hugo)/g' $(README_TEMPLATE) > README.md
	git add .
	git commit -m "Updated to the latest Hugo version v.$(hugo), see https://github.com/gohugoio/hugo/releases"
	git push origin master

clean:
	docker rmi cbrgm/drone-hugo:$(hugo)
	docker rmi cbrgm/drone-hugo:latest
	docker rmi cbrgm/drone-hugo:$(hugo)_test
