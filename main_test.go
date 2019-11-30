package SimpleLiveReloadImplementationInGo

import "testing"

func Test(t *testing.T){
	srcPath := "./example/"
	dstPath := srcPath

	// Start LiveReload server
	lr := server(":8080", dstPath)
	go lr.watch(srcPath)
	lr.run()
}
