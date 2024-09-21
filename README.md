# Go - Templ - Tailwind -- Template Project

This project serves as a template project for simple applications using Go (programming language), Templ (templating library), and Tailwind (CSS framework) including testing & release workflows.\
For the database, [Sqlite][sqlite] is used.

[Goreleaser][goreleaser] with the [Ko][ko] integration is used to 1) create GitHub releases and 2) build and push Docker images.

## Development Stack

The dev stack uses [Air][air], [Templ][templ], and [Tailwind][tailwind]. \
Air is used to automatically reload the application when any Go file changed. \
Templ is the templating framework used.\
Tailwind generates css styles fom HTML classes.

### Structure

* `main.go`: Routes & DB initialization
* `/internal/persistence`: DB access
* `/internal/views`: Frontend files
* `/static/css`: Style files. `style.css` is generated by Tailwind from `input.css`.

### Development
Execute the following commands in parallel:

* `air`: Restarts the app if any Go files change
* `templ generate --watch`: Compile *templ files to Go source
* `npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch`: Re-compile any css files if any template files change.

### GitHub
GitHub integrations are provided with two workflows.

#### [integration.yaml](./github/workflows/integration.yaml)
Executes tests & vets all pushes.

#### [release.yaml](./github/workflows/release.yaml)
Used to create releases. The pipeline is run on every push. If no tag is present, the release is not performed but validated.
If a tag in the form of `v*` is present, a GitHub release is created, as well as the Docker images are created and pushed.

### Kubernetes

[./deployment/deployment.yaml](./deployment/deployment.yaml) contains a basic Kubernetes deployment file.


[goreleaser]: https://goreleaser.com/
[ko]: https://ko.build/
[sqlite]: https://www.sqlite.org/
[air]: https://github.com/air-verse/air
[templ]: https://templ.guide/
[tailwind]: https://tailwindcss.com/
