## Contributing 

### Issues

Use the search tool before opening a new issue.

Please provide source code and commit sha if you found a bug.

Review existing issues and provide feedback or react to them.

### Pull requests

Open your pull request against ```develop```.

You should add/modify tests to cover your proposed code changes.

Tests coverage should never go down from 90%.

If your pull request contains a new feature, please document it on the README.

### Branching

The ```main``` is a **regular branch** which always contains the latest
**stable** codebase and must **never** be broken.

The ```develop``` is a **regular branch** which always contains the latest
**development** codebase and **eventually** can be broken. But you'll need to
accept the **sombrero of shame** if you do that.

The ```release``` is a **regular branch** which contains a specific release
version. You must use the following name convention: **release-X.Y.Z**, where
X, Y and Z are: major, minor and patch [version numbers](#versioning).

The ```experimental``` is a **temporary branch** which contains a new feature or
ideia. You must use the following name convention: **experimental-brief-description**.

The ```feature``` is a **temporary branch** which contains a new feature under
development that latter will be merged against the development branch. You must
use the following name convention: **feature-brief-description**.

The ```bugfix``` is a **temporary branch** which contains necessary fix to be
applied **after** a specific release to be merged against the development branch.
You must use the following name convention: **bugfix-brief-description**.

The ```hotfix``` is a **temporary branch** which contains a critical fix to be
applied **immediately** and merged against the main and the development branches.
You must use the following name convention: **hotfix-brief-description**.

Feel free to apply the labels from GitHub to the branches, they are very helpful.

### Versioning

The project uses the [semantic versioning 2.0.0](https://semver.org) in
order to control the version numbers.

### Commiting

The ```main```, ```develop``` and ```release``` branches have protection rules
against **push**.

In order to contribute you must create a new branch following the [branching](#branching)
guideline and once your work is done, open a **pull request** from your branch
to the **develop** branch.

The **pull request** will trigger the test suites automatically
and the code must **pass all the tests** and also be reviewed and approved
before merged in the **develop** branch (or even ``main`` or ``release`` in case
of a ``*fix``).

Feel free to apply the labels from GitHub to the pull requests, they are very helpful.
