# Roller

## Commands
Pull down initial template:
````
roller create <git repo url>
````

Update from template repo:
````
roller update <optional: reference/version, otherwise uses default branch>
````

Performs a user-defined action:
````
roller [action]
````

### Development
Synchronises template changes with the provided target directory, useful when building and testing a template:

````
roller sync <optional: target dir, defaults to 'roller_output', relative to working dir>
````


## Files
- `roller.yaml` - template configuration
- `.roller.state.yaml` - tracks state changes (tracked files, used to detect deleted files)


## roller.yaml

````
template:
  repo: [url to git repo]
  vars:
    <free-form fields>
  ignore:
    <pattern of files to ignore>

action:
  [action name: cant be a reserved keyword]:
    shell: [shell command here]
    working_dir: [working dir, optional, defaults to root dir]
````


## Templating Engine
- <https://pkg.go.dev/text/template>
- <https://github.com/flosch/pongo2>


## Kudos
Shout-out to Cruft for inspiration:

<https://cruft.github.io/cruft/>
