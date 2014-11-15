# pivgit


Tool to augment git branch output with PivotalTracker story names if the id of the story is part of the branch name.

## Motivation

We use gitflow at work and we create essentially a branch per ticket and given that branch descriptions were not
remotely persisted by git we needed a way to make more meaning out of our branch names without them becomming 
excessivly long.

So we started using the following convention:
piv-####### where # was the story id from PivotelTracker.

Which soon became tiresome.

## To Build

PivGit is written in GOLANG http://www.golang.org you need to download and install the latest to build this.

```
mkdir -p ~/pivgit/src
cd ~/pivgit
export GOPATH=`pwd`
cd src
git clone < -- whatever github tells you -- > .

# Make sure you put the dot on the end when cloning so that it wil go into the src directory.

cd ..

go install github.com/R358/pivgit

# Set it up for you to use, make sure ~/bin exists

cp cp bin/pivgit ~/bin/

````

## How it works

Pivgit works by tokenising the output of git branch (or anything else piped into it). 

Then when it finds a number that matches a story id it has harvested from PivotelTracker using your access Token it will
substitute the name of the story after that token. Everything else is parsed through.

When tokenising pivgit removes all non numeric characters so all you have to do is add the story id 
to the branch name and it will find it.

**NB:**
At the moment it makes some assumptions about the tokens so by removing all non numeric characters it may
inadvertently join multiple numbers together and therefore not find a story id.

**For example:**

round_2_of_piv-1234567

Will end up being 2123456 with the 2 added to the front of the number.

## Usage:


If you are on your own machine, then the simplest thing to do would be add the following to your .bash_profile
```
export PIVGIT_TOKEN=--- my token ---
```
If you don't do this then you will need to add the token as the first parameter whenever you call pivgit

To use:
```
git branch | pivgit [Token see above]
```

Before and after

```
Without pivgit:

> git branch
* develop
  master
  piv-92133556
  piv-72424294

With pivgit:

> git branch | pivgit Krispy

* develop
master
piv-92133556 Reset credits list and title list.
piv-72424294 Integrate xxxx into the app.
```
