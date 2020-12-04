//Goroutines:
//From the official go documentation: “A goroutine is a lightweight thread of execution”.
//Goroutines are lighter than a thread so managing them is comparatively less resource intensive.
//To make sure main function in a program waits for the goroutines to complete, you will need some way for
//the goroutines to tell it that they are done executing, that's where channels can help us.

//Channels:
//Channels are used when you want to pass in results or errors or any other kind of information from one goroutine to another.
//Say there is a channel ch of type int If you want to send something to a channel, the syntax is ch <- 1 if you want to receive
//something from the channel it will be var := <- ch. This recieves from the channel and stores the value in var.

//Context:
//A way to think about context package in go is that it allows you to pass in a “context” to your program. Context like a timeout or deadline or a
//channel to indicate stop working and return. For instance, if you are doing a web request or running a system command,
//it is usually a good idea to have a timeout for production-grade systems. Because, if an API you depend on is running slow,
//you would not want to back up requests on your system, because, it may end up increasing the load and degrading the performance of all the requests you serve.
//Resulting in a cascading effect. This is where a timeout or deadline context can come in handy.

package port

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// var scanme = map[int]string{ // maps each key with a value for example in this case it maps 80 to http.
// 	80:  "http",
// 	443: "https", // the ports to scan
// }

func ScanPort(hostname string, port int) bool { // refering to the each variable with their variable type. 'bool' is for the return statement.
	if port == 80 {
		address := hostname + ":" + strconv.Itoa(port)               // 'strconv.Itoa' is for converting Integers to Strings
		conn, err := net.DialTimeout("tcp", address, 60*time.Second) // the type of connection TCP/UDP, the host URL/IP ADDRESS with the port, the max time for each connection.
		if err != nil {
			return false
		}
		defer conn.Close() // 'defer' wait until the entire function is done and then execute the given command.
	}
	if port == 443 {
		address := hostname + ":" + strconv.Itoa(port)               // 'strconv.Itoa' is for converting Integers to Strings
		conn, err := net.DialTimeout("tcp", address, 60*time.Second) // the type of connection TCP/UDP, the host URL/IP ADDRESS with the port, the max time for each connection.
		if err != nil {
			return false
		}
		defer conn.Close() // 'defer' wait until the entire function is done and then execute the given command.
	}

	return true
}

// type PortOpen struct { // this is a custom type which means that i can refer PortOpen to use this type and then use the arguments passed in as a structure for different variables.
// 	Start, End int // the arguments are start and end which means that the variable will follow this type.
// }

func GetPort(hostname string, ports int) { // This function takes in the hostname as a string and a ports as PortOpen whose type is defined by us in PortOpen which actually represents a array/list that has to be fed in for a range of ports.

	v := ScanPort(hostname, ports) // refer the ScanPort function.
	if v == true {                 // if v has a value then print the printf statement.
		ctx, cancel := chromedp.NewContext(context.Background()) // this where a new context is defined and context.Background gives a new context fresh to us to work on.
		defer cancel()                                           // the cancel function if the context fails since blocking a context will lead to the program to behave unexpectedly.
		var buf []byte
		if ports == 80 {
			if err := chromedp.Run(ctx, fullScreenshot("http://"+hostname, 1080, &buf)); err != nil { // chromedp.Run will run the function fullScreenshot as a process of the context that we just declared.
				log.Fatal(err)
			}
			if err := ioutil.WriteFile(hostname, buf, 0644); err != nil { //the i/o package is used to make input/output to perform a task. We have put thr WriteFile function which writes file to the directory the script has been run.
				fmt.Println("Not Found")
			}
		} else if ports == 443 {
			if err := chromedp.Run(ctx, fullScreenshot("https://"+hostname, 1080, &buf)); err != nil { // chromedp.Run will run the function fullScreenshot as a process of the context that we just declared.
				log.Fatal(err)
			}
			if err := ioutil.WriteFile(hostname, buf, 0644); err != nil { //the i/o package is used to make input/output to perform a task. We have put thr WriteFile function which writes file to the directory the script has been run.
				fmt.Println("Not Found")
			}
		}
	}
}
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
