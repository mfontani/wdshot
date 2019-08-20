# wdshot

`wdshot` is a quick and dirty Go program to connect to a remote Selenium WebDriver instance and take a screenshot.

# How to use it

Assuming you're using BrowserStack's Selenium Web Driver (the default):

- Go to https://www.browserstack.com/accounts/settings
- Get your "Automate" "Username" (will be `WDSHOT_USERNAME`) and "Access Key" (will be `WDSHOT_PASSWORD`)
- `go get github.com/mfontani/wdshot`
- `wdshot -help`

If you're using another Selenium WebDriver, fiddle with the options.

# Caveats

This is a complete work-in-progress, and I might even stop working on it tomorrow, or tonight.

# Copyright and License

`wdshot` is Copyright (c) 2019, Marco Fontani MFONTANI@cpan.org

It is released under the MIT license - see the `LICENSE` file in this repository/directory.
