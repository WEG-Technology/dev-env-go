# Development Environment Setup

This project is designed to set up a development environment using Ansible. It installs Ansible, clones a repository, fetches a configuration YAML file, and runs an Ansible playbook. The project uses Go programming language.


## Requirements

- Go (version 1.16 or higher)
- Git
- Homebrew (for Ansible installation)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/your-repo.git
    cd development-env
    ```

2. Initialize the Go module:
    ```sh
    go mod init development-env
    ```

3. Download dependencies:
    ```sh
    go mod tidy
    ```

## Usage

Run the project using the following command:

```sh
go run cmd/main.go -url=<YAML_FILE_URL> -tempDir=<TEMPORARY_DIRECTORY> -repoURL=<REPOSITORY_URL>
```

## Example

```sh
go run cmd/main.go -url=https://example.com/config.yaml -tempDir=/tmp/my-development-environment -repoURL=https://github.com/my-repo.git
```


Arguments
-url: The URL of the YAML configuration file. (Required)
-tempDir: The temporary directory for loading files. (Default: /tmp/development-environment)
-repoURL: The URL of the repository to clone. (Default: https://github.com/your-repo.git)
