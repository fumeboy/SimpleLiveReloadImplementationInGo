package SimpleLiveReloadImplementationInGo

func Run(srcPath, dstPath, port string){
	lr := server(port, dstPath)
	go lr.watch(srcPath)
	lr.run()
}