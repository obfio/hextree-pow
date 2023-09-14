# Why
I saw that a service named [hextree.io](https://hextree.io) added a POW challenge to their password-reset page, for fun, I decided to look into it. I'm typically not very good at POW stuff when it comes to anti-bots, I struggle quite a lot, so I figured this was some decent practice.

# Explanation
When you try to password reset, it loads a webworker script that fetches a config from the server and does the POW challenge. It also appears to be doing some checks to see if you're running the script in a browser or not. After it gets the config, it starts the SHA256 cycle for POW, this SHA256 function returns the uint32's, not a string as you may expect, that's actually something I completely overlooked at first. That entire aspect is why I chose to fork [crypto/sha256](https://github.com/golang/go/tree/master/src/crypto/sha256) instead of using it as is, I didn't see any exposed methods that would return the uint32's. After it gets that uint32, it does the `>>>` operation, this doesn't exist in golang so I made a helper func for it. Next it converts the number to binary and adds padding up to 32 length with "0", so we convert out uint32 to a binary string with `fmt.Sprintf("%b", our_int)`, then add padding with `str = strings.Repeat("0", 32-len(str)) + str`. Next the POW has this function that checks if the resulting binary string is "valid", this check basically boils down to, does it have x repeating 0's prefix'd, where x is config.difficulty, returned from the server. Then that's it, if it's "valid", we add the itteration (i) to the output array and continue. This is a pretty simple POW, I had a lot of fun with this so I'll have to look more into how these hashing methods actually work code-wise in the future!