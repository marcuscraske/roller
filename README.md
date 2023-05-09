# Replicator

## Naming
rep
replicator
pave
proad
paved
roller...winner!


## Commands
Pull down initial files:
````
roller create <git repo url>
````

Update from template repo:
````
roller update <optional: reference/version, otherwise uses default branch>
````

Performs a defined action:
````
roller [action]
````


## Files
- `replicator.yaml` - template configuration


## replicator.yaml

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


## Replicator Algorithm
- Pull down template repo to a tmp dir
  - `git clone <uri>`
- New?
  - Prompt if files exist
    - Aborted? -> exit
  - Display vim for `replicator.yaml` file in tmp
- Existing?
  - Check whether `replicator.yaml` missing fields compared to tmp
  - Copy `replicator.yaml` file to tmp, add missing fields
- Apply templating using variables
- New?
  - Copy files to target dir
  - Exit.
- Existing?
  - Check for files changed between template and target path
    - Remove ignored files
    - `git diff --no-index ~/git/tmp/a ~/git/tmp2/a` - builds patch file
  - Collect list of files deleted in the template repo but exist in target path
    - Remove ignored files
  - Prompt user whether to apply all the diffs + deleted files
    - Could have mode to step through each file and y/n each file?
  - Apply changes
    - Delete files from target path that no longer exist in the template
    - Apply patches to changed files


## Templating Engine
- <https://pkg.go.dev/text/template>
- <https://github.com/flosch/pongo2>


## Kudos
Shout-out to Cruft for inspiration:

<https://cruft.github.io/cruft/>
