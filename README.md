tunnelchan
==========

Tunnelchan is a simple bot to merge IRC channels with Discord channels. It is
useful for communities which want to be accessible to both those stuck in their
anachronistic ways and to the most naive corporate techo-optimists without
splitting the community.

Currently the bot is an early stage of development. It works and *should* scale,
but it has many limitations:

* it makes no attempt to limit the messages sent to Discord, potentially breaking rate limit rules
* it does not handle many type of Discord messages properly (such as those containing newlines)
* it can not connect to IRC servers requiring authentication
* error messages are primitive
* and on and on

None of these is a hard fix, I just need to find the time to address them.

Setting up
----------

Setting it up is fairly simple. Running `go get github.com/serbuvlad/tunnelchan` shoud fetch
the code, compile it, and put the result in `$GOPATH/bin` (or `$HOME/go/bin`, if `GOPATH` is not set),
so add that direcotry to your `PATH`.

Next, you need a config file, so copy and paste the contents
of this repo's `cfg.yaml.example`, changing things as appropriate.

Finally, run

    $ tunnelchan -cfg mycfg.yaml

to put launch the bot. It will not deamonize itself and will log to stderr.