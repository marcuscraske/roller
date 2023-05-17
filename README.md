# Roller
<img src="roller.svg" />

Paving the road, with the ability to create and update projects from a template repo.


## Commands
Pull down initial template:
````
roller create <git repo url> <optional: git tag/branch, otherwise default branch>
````

_Note: you can run this in an existing git working tree, and it will clone into the current directory. Otherwise, the
template is cloned as per normal git clone._

Update from template repo:
````
roller update <optional: git tag/branch, otherwise uses default branch>
````

Performs a user-defined action:
````
roller [action]
````

Synchronises template changes with the provided target directory, useful when building and testing a template:

````
roller sync <optional: target dir, defaults to 'roller_output', relative to working dir>
````

Command-line options:
- `--survey` - forces a survey, even if there's no new template variables
- `--survey=skip` - skips the survey


## Support
Please raise an issue, or use Discord TBD.
