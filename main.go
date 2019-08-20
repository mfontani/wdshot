package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

const envUsername = "WDSHOT_USERNAME"
const envPassword = "WDSHOT_PASSWORD"

func printUsage() {
	_, err := fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	if err != nil {
		panic(err)
	}
	flag.PrintDefaults()
	fmt.Println("Note you will need to provide the following environment variables:")
	fmt.Println("- " + envUsername)
	fmt.Println("- " + envPassword)
}

func main() {

	flag.Usage = printUsage
	wantDebug := flag.Bool("debug", false, "whether to turn selenium debugging on")
	theURL := flag.String("url", "https://www.example.com/", "the URL to hit")
	theWDURL := flag.String("wdurl", "hub-cloud.browserstack.com/wd/hub", "the host/path to use to hit the remote Selenium WebDriver")
	theSelector := flag.String("css", ".srow.rt-3c", "the CSS selector to use")
	theScroll := flag.Int("scroll", -200, "how much to scroll after hitting the selector")
	thePNG := flag.String("png", "foo.png", "the output PNG file")
	theBrowser := flag.String("browser", "Firefox", "the browser to use")
	theBrowserVersion := flag.String("browserversion", "69.0 beta", "the browser version to use")
	theOS := flag.String("os", "Windows", "the OS to use")
	theOSVersion := flag.String("osversion", "10", "the OS version to use")
	theResolution := flag.String("resolution", "1280x1024", "the resolution to use")
	theTestName := flag.String("test", "BStack Sample Test", "the test name to use")
	theSleep := flag.Int64("sleep", 1, "the seconds to sleep before moving to the element (seconds)")
	flag.Parse()

	// pass sensitive info through the environment!
	WDUsername := os.Getenv(envUsername)
	if len(WDUsername) == 0 {
		panic("Need " + envUsername)
	}
	WDKey := os.Getenv(envPassword)
	if len(WDKey) == 0 {
		panic("Need " + envPassword)
	}

	f, err := os.Create(*thePNG)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	selenium.SetDebug(*wantDebug)

	caps := selenium.Capabilities{
		"browser":         *theBrowser,
		"browser_version": *theBrowserVersion,
		"os":              *theOS,
		"os_version":      *theOSVersion,
		"resolution":      *theResolution,
		"name":            *theTestName,
		// TODO more options via flags/command line
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://%s:%s@%s", WDUsername, WDKey, *theWDURL))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()
	if err = wd.Get(*theURL); err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(*theSleep) * time.Second)
	el, err := wd.FindElement(selenium.ByCSSSelector, *theSelector)
	if err != nil {
		panic(err)
	}
	_, err = wd.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{el})
	if err != nil {
		panic(err)
	}
	_, err = wd.ExecuteScript("window.scrollBy(0,arguments[0])", []interface{}{*theScroll})
	if err != nil {
		panic(err)
	}
	data, err := wd.Screenshot()
	if err != nil {
		panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}
