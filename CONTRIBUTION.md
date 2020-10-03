# Contribution Guide

You'll add a new topic in the new directory unless it's fix or addition to the existing topic.  Make sure to follow the naming convention.

If you add a new topic then add its name and link in `README.md`.

Add the references in `REFERENCE.md` file.

Follow the following step to contribute.

1. Fork this repo.

2. Clone your fork into your computer.

3. Add remote upstream:
    * `git remote add upstream https://github.com/pratikjagrut/go-tutorial.git`

4. Fetch upstream remote:
    * `git fetch upstream`

5. Create a local branch to work on your issue:
    * `git checkout -b your_branch master` your_branch tracking to your fork master branch. 
    * `git checkout -b your_branch upstream/master` your_branch tracking to the upstream master branch

6. After finishing the work:
    * `git status` to check changed files.
    * `git add changed_files`
    * `git commit` or `git commit -m 'your commit message'`
    * `git push origin your_branch`
    * The above step will print a link to create a pull request or go to your fork on GitHub and create a pull request from just pushed branch.

7. Your pull request will be reviewed by the maintainer and other contributors.