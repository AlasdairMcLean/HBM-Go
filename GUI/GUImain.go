package main

import (
	"fmt"
	"hbm/plot"
	"hbm/util"
	"time"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/gtk"
)

const (
	//WINWID is the width of the window by default
	WINWID = 800
	//WINHEI is the height of the window by default
	WINHEI = 600
)

func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, box := setupWindow("HBM plot", 800, 600)
	da, err := gtk.DrawingAreaNew() // make a new gtk drawing area where the image will go
	if err != nil {
		panic(err)
	}
	box.PackStart(da, true, true, 0) // 'pack' the drawing area from the top of the screen
	win.Add(box)                     // add the box that holds the drawing area to the window
	win.ShowAll()                    // reveal the window

	file := getfile()
	// as an example, read TSMEP1.csv
	pts := hbmutil.ReadMEPcsvf(file) //use the csv parse library in hbmutil to return an array of points and corresponding mep amplitudes
	if err != nil {
		panic(err)
	}
	t0 := time.Now()
	ptsi := pts.ToMati() // round the float64 pts down to int
	xcol := 0
	ycol := 1
	Cmap := "hot"
	a := hbmutil.MattoImgi(ptsi, xcol, ycol, 256, 256) // use the integer-based points-to-image function in the hbmutil package to return a 256 by 256 px image
	img := a.ToMatff()                                 // return the image to float64 as Cairo needs
	MEPs := pts.Getcol(3)                              // get the 4th column
	hbmplot.Drawim(da, img, Cmap)                      // draws image on the drawing area using hbmutil package
	canvsc := 1                                        // keep track of how scaled we are
	keyMap := map[uint]func(){                         // map keyboard keys to a function
		gdk.KEY_equal: func() { // If the enter key is pressed, do the following:
			canvsc++                                  // increase the canvas scale so we keep track of what the user is looking at
			hbmplot.ZoomCanvas(da, img, canvsc, Cmap) // and zoom the canvas in appropriately
		},
		gdk.KEY_minus: func() { // If the minus key is pressed, do the following:
			if canvsc > 1 { // so long as we aren't trying to zoom out from full scale (which would result in dividing by zero)
				canvsc--                                  // decrease the canvas scale so we keep track of what the user is looking at
				hbmplot.ZoomCanvas(da, img, canvsc, Cmap) // and zoom out the canvas appropriately
			}
		},
		gdk.KEY_Insert: func() { // If the insert key is pressed, do the following:
			t1 := time.Now()
			if xcol < 2 { // rotate through the x columns
				xcol++ //increment the excel column to represent on the x-axis.
			} else { // if we have the last column on the x-axis,
				xcol = 0 // then reset back to the first column so we don't trail off into empty columns
			}
			fmt.Println("Col switch: ", time.Since(t1))

			t1 = time.Now()
			a = hbmutil.MattoImgi(ptsi, xcol, ycol, 256, 256) // change the points from the excel column into an image matrix
			fmt.Println("Matrix to imgi: ", time.Since(t1))

			t1 = time.Now()
			img = a.ToMatff() // change the image matrix to a double precision matrix for cairo
			fmt.Println("To matff: ", time.Since(t1))

			t1 = time.Now()
			MEPs = pts.Getcol(3) // get the MEPs from the parsed points
			fmt.Println("getcol: ", time.Since(t1))

			t1 = time.Now()
			hbmplot.Drawim(da, img, Cmap) // draw the image on the screen
			fmt.Println("drawim: ", time.Since(t1))
		},
		gdk.KEY_Delete: func() { // if the delete key is pressed, do the following:
			if ycol < 2 { // iterate through the column to plot on the y-axis
				ycol++ // keep track of the column we are graphing
			} else { // but if we are already plotting the last column,
				ycol = 0 // then reset back to the first column
			}
			a = hbmutil.MattoImgi(ptsi, xcol, ycol, 256, 256) // convert the points from the excel sheet into an image matrix
			img = a.ToMatff()                                 // convert that image matrix into double precision for cairo
			MEPs = pts.Getcol(3)                              // get the MEP column
			hbmplot.Drawim(da, img, Cmap)                     // draw the image on the screen
		},
	}
	fmt.Println("should only run once: ", time.Since(t0))

	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) { // capture key-press events for the key commands that are already enabled
		t0 := time.Now()
		keyEvent := &gdk.EventKey{ev}                          // capture the key from the queue
		if action, found := keyMap[keyEvent.KeyVal()]; found { // if the key pressed is in fact one of the valid key commands,
			t1 := time.Now()
			action() // do the function associated with that key (zoom, switch column, ect)
			fmt.Println("total func: ", time.Since(t1))
			t1 = time.Now()
			win.QueueDraw() // redraw the screen to update for changes
			fmt.Println("queuedraw: ", time.Since(t1))

		}
		fmt.Println("connect: ", time.Since(t0))
		fmt.Println("")
	})

	gtk.Main() // start the main loop, waiting for user to close the window
}

//setupWindow will perform all the startup boilerplate for the main window of the application
func setupWindow(title string, wid, hei int) (*gtk.Window, *gtk.Box) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL) //create a new toplevel window
	if err != nil {
		pmsg := fmt.Sprintf("Unable to create window: %v", err)
		panic(pmsg) // capture the error if for some reason the window cannot be created
	}
	gtk.WindowSetDefaultIconFromFile("./brainimg.png") // use the brainimg.png file as an icon
	win.SetTitle(title)                                // sets title to that specified as input to setupWindow function
	win.Connect("destroy", func() {                    //destroy the window if the user clicks the x
		gtk.MainQuit()
	})
	win.SetDefaultSize(wid, hei)                     // set the default size to that specified in setupWindow input
	win.SetPosition(gtk.WIN_POS_CENTER)              // centers the window on the screen
	mainbox := makevGTKBox()                         // make a new box for packing elements
	setupBoxToolBar(mainbox, "File", "Edit", "Help") // make an example toolbar
	return win, mainbox                              // return the window and the main box unit to hold the elements inside
}

func setupBoxToolBar(box *gtk.Box, label ...string) { //sets up the example toolbar. will have callback functions later.
	guitoolbar, err := gtk.ToolbarNew() // create a new empty gtk toolbar object
	if err != nil {
		panic(err)
	}

	for i, v := range label { //for each label specified in setupBoxToolBar input
		newbutton, _ := gtk.ToolButtonNew(nil, v) //create the new button with label equal to string
		guitoolbar.Insert(newbutton, i)           // insert that button into the toolbar
	}
	box.PackStart(guitoolbar, false, false, 0) //pack all the buttons in at the left side of the bar
}

func makevGTKBox() *gtk.Box { //boilerplate code for a new box to hold all the objects
	newbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1) //make a vertical box (means that startpack will pack from the top)
	if err != nil {
		panic(err) // capture the error if for some reason the box cannot be created
	}

	return newbox // return the newly created box
}

func getfile() string {
	dlgwin, err := gtk.WindowNew(gtk.WINDOW_POPUP) //create a new toplevel window
	if err != nil {
		fmt.Printf("\nUnable to create window: ")
		panic(err)
	}
	dlgwin.SetTitle("Pick a csv file") // sets title to that specified as input to setupWindow function
	dlgwin.Connect("destroy", func() { //destroy the window if the user clicks the x
		gtk.MainQuit()
	})
	dlgwin.SetDefaultSize(WINWID/2, WINHEI/2)               // set the default size to that specified in setupWindow input
	dlgwin.SetPosition(gtk.WIN_POS_CENTER)                  // center the window
	filechooser, _ := gtk.FileChooserDialogNewWith2Buttons( // set up the file chooser window
		"Select MEPs...", // Set the title of the window
		dlgwin,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Cancel", // use a button if the user wishes to cancel the file open operation
		gtk.RESPONSE_DELETE_EVENT,
		"Open", // use a button for opening the file
		gtk.RESPONSE_ACCEPT)
	filter, _ := gtk.FileFilterNew()
	filter.AddPattern("*.csv")
	filter.SetName(".csv")
	filechooser.AddFilter(filter)

	switcher := filechooser.Run()
	filename := filechooser.GetFilename()
	filechooser.Destroy()
	if switcher != -3 {

		//Nothing more to do here
		filename = ""
		return ""
	}
	return filename
}

/*import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/gillesdemey/go-dicom"
)

func main() {

	bytes, err := ioutil.ReadFile("TAIZA_GRESPAN_SANTOS0100.dcm")
	if err != nil {
		panic(err)
	}

	parser, err := dicom.NewParser()
	if err != nil {
		panic(err)
	}

	data, c := parser.Parse(bytes)

	dcm := &dicom.DicomFile{}
	gw := new(sync.WaitGroup)
	dcm.Discard(c, gw)
	gw.Wait()
	for _, elem := range data.Elements {
		fmt.Printf("%+v\n", &elem)
	}
	pdata := data.Elements[52]
	//pdatain := pdata.Value
	//pdatain2 := (pdatain[0]).(string)
	pdatain2 := pdata.Value
	pdatain3 := pdatain2[0].(string)
	pdatain4 := []int32(pdatain3)
	fmt.Println("---")
	fmt.Println(pdatain4)
	fmt.Println("---")
	fmt.Println(len(pdatain4))
	fmt.Println(len(pdatain2))
	fmt.Println("---")
	var joined []string
	for i := 0; i < len(pdatain4); i++ {
		strings.Join(joined, string(pdatain4[i]))
	}
	fmt.Println(joined)
}
*/
