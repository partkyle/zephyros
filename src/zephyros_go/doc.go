/*
	Example script:

		 import (
			 "fmt"
			 . "./zephyros_go"
		 )

		 func main() {
			 API.Bind("d", []string{"cmd", "shift"}, func() {
				 API.Alert("LIKE", 1)

				 win := API.FocusedWindow()
				 fmt.Println(win.Title())

				 f := win.TopLeft()
				 f.X += 10
				 win.SetTopLeft(f)

				 API.ChooseFrom([]string{"foo", "bar"}, "title", 20, 20, func(i int) {
					 fmt.Println(i)
				 })
			 })

			 API.Listen("app_launched", func(app App) {
				 API.Alert(app.Title(), 1)
			 })

			 ListenForCallbacks()
		 }

	A note about resources:

	Every window/app/screen is a resource, and they each have methods Retain and Release.

	- These methods must be used when you want to keep a refernce around longer than a single callback.
	- Retain increments the retain-count and release decrements it. When it reaches 0, it will be garbage-collected after 5 seconds.
	- When you first get a resource back, it starts with a retain-count of 0.

*/
package zephyros_go
