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

func ScanPort(hostname string, port int) bool { // refering to the each variable with their variable type. 'bool' is for the return statement.
	if port == 80 {
		address := hostname + ":" + strconv.Itoa(port)               
		conn, err := net.DialTimeout("tcp", address, 60*time.Second) 
		if err != nil {
			return false
		}
		defer conn.Close() // 'defer' wait until the entire function is done and then execute the given command.
	}
	if port == 443 {
		address := hostname + ":" + strconv.Itoa(port)               
		conn, err := net.DialTimeout("tcp", address, 60*time.Second) 
		if err != nil {
			return false
		}
		defer conn.Close() // 'defer' wait until the entire function is done and then execute the given command.
	}

	return true
}

func GetPort(hostname string, ports int) { 
	v := ScanPort(hostname, ports) 
	if v == true {                 
		ctx, cancel := chromedp.NewContext(context.Background()) 
		defer cancel()                                           
		var buf []byte
		if ports == 80 {
			if err := chromedp.Run(ctx, fullScreenshot("http://"+hostname, 1080, &buf)); err != nil { 
				log.Fatal(err)
			}
			if err := ioutil.WriteFile(hostname+"http", buf, 0644); err != nil { 
				fmt.Println("Not Found")
			}
		} else if ports == 443 {
			if err := chromedp.Run(ctx, fullScreenshot("https://"+hostname, 1080, &buf)); err != nil { 
				log.Fatal(err)
			}
			if err := ioutil.WriteFile(hostname+"https", buf, 0644); err != nil { 
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
