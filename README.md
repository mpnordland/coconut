coconut
=======

**WARNING!!! If you are using coconut, you MUST upgrade to the latest git revision. There is a bug in earlier versions that only allows you to log in with a correct user name and an incorrect password. This includes empty passwords.**

A simple markdown based blog engine.

The configuration file is conf.yaml. It *must* exist because there are no builtin defaults.

Articles are in Markdown and should be placed in `articles` with the file extension ".md". The theme files are in `static/theme`. Any static files like CSS and images should be placed in `static`. The login and publish pages are in there as well. Pages are in Markdown, with urls and file paths specified in conf.yaml. Note that all page file paths are relative to `static`.

You can use [coconut-post](https://github.com/mpnordland/coconut-post) to partially automate creating posts.

A performance note: Coconut uses bcrypt for the hashed passwords in the config file. It's apart of the very basic login system used to protect the very basic publishing page.
This may make it unsuitable for low powered servers. However, use of the login/publishing system is optional; no performance penalty will be incurred if they are not used.
Another option may be to use a small work factor. Such a decision should be made based upon how often you back up stuff and how valuble the content on the blog is. The publishing system at present willingly overwrites articles.

![Screenshot of the default coconut setup](https://drive.google.com/file/d/0B_dqqSENmE0CUzhWOVFsSmx4em8/edit?usp=sharing "coconut")

