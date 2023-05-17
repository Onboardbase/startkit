# Overview
obbkit is a CLI tool for generating boilerplates with onboardbase preconfigured in them as a SecretOps infrastucture.

# Usage
Run to see all the commands available
```
obbkit help
```

This will display information about the CLI

![alt obb help](./docs/images/obbkit_help.png)

To get more information about a command run the command with `-h`, for example:
```
obbkit init -h
```

## Initializing a Project
Run the init command:
```
obbkit init
```
After providing the project name, you can select a project type:
![alt Init](./docs/images/obb_init.png)

**Note**: *You must have the build tools needed to create the project you selected. For example, if you selected a project type based on NodeJS e.g NestJs, Next.js, you must have Node.js already installed on your computer*

After selecting a project type, you will be prompted to complete the flow for creating project:

![alt Init](./docs/images/init_project.png)

## Initializing a Project From Git Repository
**Note**: *You must have git installed and configure on your machine to do this*

obbkit can also automatically setup a project directrly from a github repo. It will attempt to clone the repository and then configure Onboardbase inside of it. This means that you must have the permission to clone the repository if it's a private repository.

To initialize from a git reposity run
```
obbkit init --from-git [git-repository]

e.g

obbkit init --from-git https://github.com/Onboardbase/Flask-Starterkit.git
```

## Contributing

