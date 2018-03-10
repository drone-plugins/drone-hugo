# cbrgm/drone-hugo

[![GitHub release](https://img.shields.io/github/release/cbrgm/drone-hugo.svg)](https://github.com/cbrgm/drone-hugo/releases) ![](https://img.shields.io/badge/hugo%20version-v<HUGO_VERSION>-ff69b4.svg)
[![Docker Pulls](https://img.shields.io/docker/pulls/cbrgm/drone-hugo.svg)](https://hub.docker.com/r/cbrgm/drone-hugo/tags/)

**Automatically create static web page files using Hugo within your drone pipeline!**

cbrgm/drone-hugo is:

- **Easy** to implement in your existing pipeline using `.drone.yml`
- **Small** 21mb image size
- **Highly configurable**

## Basic Usage with Drone CI

The example below demonstrates how you can use the plugin to automatically create static web page files using Hugo. **It's as easy as pie!**

```yml
pipeline:
  hugo:
    image: cbrgm/drone-hugo:<HUGO_VERSION>
    validate: true
```

`validate` allows you to check your configuration file for errors before generating the files.

### Customize source, output, theme, config, layout OR content directory paths

You can customize the paths for e. g. the theme, layout, content directory and output directory and much more!

```yml
pipeline:
  hugo:
    image: cbrgm/drone-hugo:<HUGO_VERSION>
+   config: path/to/config
+   content: path/to/content/
+   layout: path/to/layout
+   output: path/to/public
+   source: path/to/source
+   theme: path/themes/THEMENAME/
    validate: true
```

### Set hostname (and path) to the root

You can also define a base URL directly in the pipeline, which is used when generating the files.

```yml
pipeline:
  hugo:
    image: cbrgm/drone-hugo:<HUGO_VERSION>
    config: path/to/config
    content: path/to/content/
    output: path/to/public
    source: path/to/source
    theme: path/themes/THEMENAME/
+   url: https://cbrgm.de
    validate: true
```

### Build sites and include expired, drafts or future content

You can set the `buildDrafts`, `buildExpired`, `buildFuture` settings to configure the generated files.

- `buildDrafts` - include content marked as draft
- `buildExpired` - include expired content
- `buildFuture` - include content with publishdate in the future

```yml
pipeline:
  hugo:
    image: cbrgm/drone-hugo:<HUGO_VERSION>
+   buildDrafts: true
+   buildExpired: true
+   buildFuture: true
    config: path/to/config
    content: path/to/content/
    output: path/to/public
    source: path/to/source
    theme: path/themes/THEMENAME/
    url: https://cbrgm.de
    validate: true
```

### **Example**: Generate Hugo static files and publish them to remote directory using scp

Here is a short example of how to define a pipeline that automatically generates the static web page files with Hugo and then copies them to a remote server via scp. This makes publishing websites a breeze!

```yml
pipeline:
  build:
    image: cbrgm/drone-hugo:<HUGO_VERSION>
    output: site # Output path
    validate: true
    when:
      branch: [ master ]
  publish:
    image: appleboy/drone-scp
    host: cbrgm.de
    username: webuser
    password: xxxxxxx
    port: 54321
    target: /var/www/ # Path to your web directory
    source: site/* # Copy all files from output path
```

You can also use secrets to hide credentials:

```yml
pipeline:
  build:
    image: cbrgm/drone-hugo:<HUGO_VERSION>
    output: site # Output path
    validate: true
    when:
      branch: [ master ]
  publish:
    image: appleboy/drone-scp
+   secrets: [ ssh_username, ssh_password ]
    host: cbrgm.de
-   username: webuser
-   password: xxxxxxx
    port: 54321
    target: /var/www/ # Path to your web directory
    source: site/* # Copy all files from output path
```

## Basic Usage using a Docker Container

```bash
docker run --rm \
  -e PLUGIN_BUILDDRAFTS=false \
  -e PLUGIN_BUILDEXPIRED=false \
  -e PLUGIN_BUILDFUTURE=false \
  -e PLUGIN_CONFIG=false \
  -e PLUGIN_CONTENT=false \
  -e PLUGIN_LAYOUT=false \
  -e PLUGIN_OUTPUT=false \
  -e PLUGIN_SOURCE=false \
  -e PLUGIN_THEME=false \
  -e PLUGIN_OUTPUT=false \
  -e PLUGIN_VALIDATE=false \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  cbrgm/drone-hugo:latest
```

## Parameter Reference

`buildDrafts` - include content marked as draft<br>
`buildExpired` - include expired content<br>
`buildFuture` - include content with publishdate in the future<br>
`config` - config file (default is path/config.yaml|json|toml)<br>
`content` - filesystem path to content directory<br>
`layout` - filesystem path to layout directory<br>
`output` - filesystem path to write files to<br>
`source` - filesystem path to read files relative from<br>
`theme` - theme to use (located in /themes/THEMENAME/)<br>
`url` - hostname (and path) to the root<br>
`validate` - validate config file before generation

## Contributing

You have suggestions for improvements or features you miss? You are welcome to express all your wishes here. Just create a new Issue and it will be taken care of quickly!

If you are a developer yourself, you can also contribute code! Further information will follow shortly.
