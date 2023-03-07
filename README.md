# Bor Overview
Bor is the Golang implementation of the Polygon Matic protocol adapted to Algofoundry requirements.

It is a fork of [Polygon matic's Bor](https://github.com/maticnetwork/bor) with some changes.

### Changes from Bor

- capability to disable ancient db write
- real time pruning for active data (this can reduce mainnet data size up 30% - 40%)


### Submit an issue

- Create a [new issue](https://github.com/AlgoFoundry/bor/issues/new/choose)
- Comment on the issue (if you'd like to be assigned to it) - that way [our team can assign the issue to you](https://github.blog/2019-06-25-assign-issues-to-issue-commenters/).
- If you do not have a specific contribution in mind, you can also browse the issues labelled as `help wanted`
- Issues that additionally have the `good first issue` label are considered ideal for first-timers

### Fork the repository (repo)
- You can fork the repository from [here](https://github.com/AlgoFoundry/bor/fork).
- To [sync your fork with the latest changes](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork):

    ```
    $ git checkout develop
    $ git fetch upstream
    $ git merge upstream/develop
    ```

### Building the source

- Building `bor` requires both a Go (version 1.19 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

     ```shell
     $ make bor -B
     ```

### Make awesome changes!

1. Create new branch for your changes

    ```
    $ git checkout -b new_branch_name
    ```

2. Commit and prepare for pull request (PR). In your PR commit message, reference the issue it resolves (see [how to link a commit message to an issue using a keyword](https://docs.github.com/en/free-pro-team@latest/github/managing-your-work-on-github/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword).

    ```
    $ git commit -m "[#1234] brief description of changes"
    ```

3. Push to your GitHub account

    ```
    $ git push
    ```

### Submit your PR

- After your changes are commited to your GitHub fork, submit a pull request (PR) to the `develop` branch of the `AlgoFoundry/bor` repo
- In your PR description, reference the issue it resolves (see [linking a pull request to an issue using a keyword](https://docs.github.com/en/free-pro-team@latest/github/managing-your-work-on-github/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword))
  - ex. `[#1234] Updates out of date content`

### Wait for review

- The team reviews every PR
- Acceptable PRs will be approved & merged into the `develop` branch

<hr style="margin-top: 3em; margin-bottom: 3em;">

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.

<hr style="margin-top: 3em; margin-bottom: 3em;">