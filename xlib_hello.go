package main

import (
	"github.com/vbsw/xlib"
	"log"
//	"net"
	"os"
//	"strings"
//	"time"
)


/* Create window, log X11 events and display pressed keys.
 *
 * Usage:
 *   gcc -o xlib_hello xlib_hello.c -lX11
 *   ./xlib_hello
#include <X11/Xlib.h>
#include <X11/Xutil.h>

#include <string.h>
#include <stdio.h>

//for sleep
#include <unistd.h>

//for system
#include <stdlib.h>

static const char *event_names[] = {
    "",
    "",
    "KeyPress",
    "KeyRelease",
    "ButtonPress",
    "ButtonRelease",
    "MotionNotify",
    "EnterNotify",
    "LeaveNotify",
    "FocusIn",
    "FocusOut",
    "KeymapNotify",
    "Expose",
    "GraphicsExpose",
    "NoExpose",
    "VisibilityNotify",
    "CreateNotify",
    "DestroyNotify",
    "UnmapNotify",
    "MapNotify",
    "MapRequest",
    "ReparentNotify",
    "ConfigureNotify",
    "ConfigureRequest",
    "GravityNotify",
    "ResizeRequest",
    "CirculateNotify",
    "CirculateRequest",
    "PropertyNotify",
    "SelectionClear",
    "SelectionRequest",
    "SelectionNotify",
    "ColormapNotify",
    "ClientMessage",
    "MappingNotify"};

struct windowWithDeletePackage
{
    Window window;
    Atom wm_delete;
};

void createWindowWithVMDelete(struct windowWithDeletePackage *new_window_package, Display *display, int screen, Window parent_window)
{
    int x = 0;
    int y = 0;

    unsigned int width = 400;
    unsigned int height = 40;

    unsigned int border_width = 1;

    unsigned int border_color = BlackPixel(display, screen);
    unsigned int background_color = WhitePixel(display, screen);

    // Create window
    new_window_package->window = XCreateSimpleWindow(display, parent_window,
                                                     x,
                                                     y,
                                                     width,
                                                     height,
                                                     border_width,
                                                     border_color,
                                                     background_color);

    long event_mask = ExposureMask | KeyPressMask | KeyReleaseMask | ButtonPressMask | ButtonReleaseMask | FocusChangeMask;

    // Select window events
    XSelectInput(display, new_window_package->window, event_mask);

    // Make window visible
    XMapWindow(display, new_window_package->window);

    // Set window title
    XStoreName(display, new_window_package->window, "Hello, World!");

    // Get WM_DELETE_WINDOW atom
    new_window_package->wm_delete = XInternAtom(display, "WM_DELETE_WINDOW", True);

    // Subscribe WM_DELETE_WINDOW message
    XSetWMProtocols(display, new_window_package->window, &(new_window_package->wm_delete), 1);
}

void handleEventLoop(struct windowWithDeletePackage *target_window_package, Display *display, int screen)
{
    GC gc = DefaultGC(display, screen);
    char msg[1024] = "";
    char key[32];

    char buf[1024];
    sprintf(buf, "lsof -p %d", getpid());
    system(buf);

    // Event loop
    for (;;)
    {
        // Get next event from queue
        XEvent event;
        XNextEvent(display, &event);

        // Print event type
        printf("got event: %s\n", event_names[event.type]);

        // Keyboard
        if (event.type == KeyPress)
        {
            int len = XLookupString(&event.xkey, key, sizeof(key) - 1, 0, 0);
            key[len] = 0;

            if (strlen(msg) > 50)
                msg[0] = 0;

            strcat(msg, key);
            strcat(msg, " ");
        }

        // Refresh
        if (event.type == KeyPress || event.type == Expose)
        {
            XClearWindow(display, target_window_package->window);
            XDrawString(display, target_window_package->window, gc, 10, 20, msg, strlen(msg));
        }

        // Close button
        if (event.type == ClientMessage)
        {
            if (event.xclient.data.l[0] == target_window_package->wm_delete)
            {
                printf("handling WM_DELETE_WINDOW from WM protocol with data [0x%lX]\n", target_window_package->wm_delete);
                printf("Collect Trash and destroy every thing should cost 1 second\n");
                sleep(1);
                break;
            }
        }
    }
}

int main(int argc, char **argv)
{
    Display *display = XOpenDisplay(NULL);
    if (display == NULL)
    {
        printf("OpenDisplay for Current Dissplay Failed\n");
        display = XOpenDisplay(":0");
        if (display == NULL)
        {
            printf("OpenDisplay for :0 Failed\n");
            return 1;
        }
        else
        {
            printf("OpenDisplay for :0 success\n");
        }
    }
    else
    {
        printf("OpenDisplay for Current Dissplay success\n");
    }

    int screen = DefaultScreen(display);

    Window parent_window = DefaultRootWindow(display);

    struct windowWithDeletePackage hello_window;
    createWindowWithVMDelete(&hello_window, display, screen, parent_window);

    handleEventLoop(&hello_window, display, screen);

    XCloseDisplay(display);
    return 0;
}
 */
func main() {
	 //display:=xlib.XOpenDisplay("")
	 display:=xlib.XOpenDisplay(":0")
	 if display == nil {
		 log.Println("OpenDisplay for Current Dissplay Failed");
		 display = xlib.XOpenDisplay(":0");
		 if display == nil {
			 log.Println("OpenDisplay for :0 Failed\n");
			 os.Exit( 1);
		 } else {
			 log.Println("OpenDisplay for :0 success\n");
		 }
	 } else {
		 log.Println("OpenDisplay for Current Dissplay success\n");
	 }
    //defer xlib.XCloseDisplay(display);

    screen := xlib.XDefaultScreenOfDisplay(display);

    parent_window :=xlib.XRootWindowOfScreen(screen);

    //Window parent_window = xlib.DefaultRootWindow(display);

    //struct windowWithDeletePackage hello_window;
    //createWindowWithVMDelete(&hello_window, display, screen, parent_window);

    /*************/
    x := 0;
    y := 0;

    var width uint = 400;
    var height uint = 40;

    var border_width uint = 1;

    //unsigned int border_color = BlackPixel(display, screen);
    var border_color uint64 = 99999
    //unsigned int background_color = WhitePixel(display, screen);
    var background_color uint64= 1111

    // Create window
    window := xlib.XCreateSimpleWindow(display, parent_window,
                                                     x,
                                                     y,
                                                     width,
                                                     height,
                                                     border_width,
                                                     border_color,
                                                     background_color);

     event_mask := xlib.ExposureMask | xlib.KeyPressMask | xlib.KeyReleaseMask | xlib.ButtonPressMask | xlib.ButtonReleaseMask | xlib.FocusChangeMask;

    // Select window events
    xlib.XSelectInput(display, window, event_mask);

    // Make window visible
    xlib.XMapWindow(display, window);

    /*
    // Set window title
    //xlib.XStoreName(display, window, "Hello, World!");

    // Get WM_DELETE_WINDOW atom
    wm_delete := xlib.XInternAtom(display, "WM_DELETE_WINDOW", True);

    // Subscribe WM_DELETE_WINDOW message
    XSetWMProtocols(display, new_window_package->window, &(new_window_package->wm_delete), 1);
    */

    //time.Sleep(time.Duration(10) * time.Second)

    /*************/
    /*************/
    /*************/
    /*************/

    //handleEventLoop(&hello_window, display, screen);


	 //Display *display = XOpenDisplay(NULL);
	 /*
	 server, err := net.Listen("tcp", "0.0.0.0:8888")
	 if err != nil {
		 log.Fatal("start server failed!\n")
		 os.Exit(1)
	 }
	 defer server.Close()

	 log.Println("server is running...")
	 for {
		 client, err := server.Accept()
		 if err != nil {
			 log.Fatal("Accept error\n")
			 continue
		 }

		 log.Println("the client is connectted...")
		 go recvMessage(client)
	 }
	 */

//	 gc := xlib.DefaultGC(display, screen);
    /*
    char msg[1024] = "";
    char key[32];

    char buf[1024];
    sprintf(buf, "lsof -p %d", getpid());
    system(buf);
    */

    // Event loop
    for  {
        // Get next event from queue
	event:= xlib.XNextEvent(display );
	if event != nil{
        // Print event type
        log.Printf("got event: %d\n", event.Type());
	}else{
        log.Printf("got nil event\n");
	}
    }

}
