
# Coding Guidelines

The Mercado Pago Go SDK is a collaborative effort from the start. The SDK team believes that contributions from different developers will enrich its feature set and make it more relevant to the community.

However, absorbing all contributions as-is, while expedient, might lead to difficulties in maintaining the codebase if left unchecked. A collaborative codebase often establishes guidelines for contributors to ensure that the code remains maintainable over time. The effort to maintain the SDK is no different in this regard, so a bit of guidance is in order.

The purpose of this guide is to set a baseline for contributions. These guidelines are not intended to limit the tools at your disposal or rewire the way you think but rather to encourage good neighbor behavior.

## Language Guidelines

We use **English** language. This is to be consistent everywhere and to be considerate of developers who do not speak our native language.

Therefore: source code, comments, documentation, commit messages, review comments, and any other kind of contribution *MUST* use the English language.

Typos are unavoidable, but try to reduce them by using a spellchecker. Most IDEs can be configured to run one automatically.

## Code Guidelines

To contribute to the project, you need to have the [pre-commit tool](https://pre-commit.com/) installed on your machine to check code style and formatting.

1. To check if pre-commit is installed, you must successfully run the following command:

```shell
$ pre-commit --version
```
pre-commit 2.17.0

2. After pre-commit is installed on your machine, within the SDK project folder, you must run the following command to set up the git hook scripts of the project:

```shell
$ pre-commit install
```

After that, git hooks will run automatically before every commit command.

In general, be conscious when contributing and try to follow the same style that the code in the SDK already has. If you have any doubts, just ask us!

These rules will be enforced automatically when making a pull request, and checks will fail if you do not follow them, resulting in your contribution being automatically rejected until fixed.

## Comment Guidelines

Comments in code are a hard thing to write, not because the words are difficult to produce but because it is hard to make relevant comments. Too much of it, and people do not read comments (and it obfuscates code reading), and too little of it gives you no recourse but to read large portions of the codebase to understand what a feature/code block is doing. Both situations are undesirable, and efforts should be made at all times to have a pleasant comment reading experience.

As a general rule, you would have to comment on decisions you made while coding that are not part of any specification.

In particular, you should always comment on any decision that:

* Departs from common wisdom or convention (The **why's** are necessary).
* Takes a significant amount of time to produce. A good rule of thumb here is that if you spent more than 1 hour thinking about how to produce a fragment of code that took 2 minutes of wrist time to write, you should document your thinking to aid the reader and allow for validation.
* Needs to preserve properties of the implementation. This is the case of performance-sensitive portions of the codebase, goroutines synchronization, implementations of security primitives, congestion control algorithms, etc.

As a general rule of what not to comment, you should avoid:

* Commenting on the structure of programs that is already part of a convention, specified or otherwise.
* Having pedantic explanations of behavior that can be found by immediate examination of the surrounding code artifacts.
* Commenting on behavior you cannot attest.

### Branching Guidelines

Currently, `main` is our only long-term branch. Below are a few suggestions for short-term branch names:

* `hotfix/something-needs-fix`: Small routine patches in code for an existing feature.
* `feature/something-new`: A new feature or a change in an existing feature. Beware of breaking changes that would require a major version bump.
* `doc/improves-documentation-for-this-feature`: If you add or change documentation with no impact on the source code.

### Git Guidelines

All commits **SHOULD** follow the [seven rules of a great Git commit message](https://chris.beams.io/posts/git-commit):

1. Separate subject from body with a blank line.
2. Limit the subject line to 72 characters.
3. Capitalize the subject line.
4. Do not end the subject line with a period.
5. Use the imperative mood in the subject line.
6. Wrap the body at 72 characters.
7. Use the body to explain what and why vs. how.

Commits such as "fix tests", "now it's working", and many other common messages we usually find in code **WON'T** be accepted.

Ideally, we would like to enforce these rules, but we are realistic and understand that it might be a big change for some people. So, unless deviating heavily from what was stated, we might accept your commits even if not following these rules perfectly.

Please avoid taking too much time to deliver code and always [rebase](https://git-scm.com/docs/git-rebase) your code to avoid reverse merge commits.
