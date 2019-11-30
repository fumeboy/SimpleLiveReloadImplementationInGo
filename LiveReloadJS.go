package SimpleLiveReloadImplementationInGo

const script = `let ws = new WebSocket("ws://127.0.0.1%s/livereload");
ws.onmessage = function(event) {
    var data = event.data;
console.log(data);
    window.location.reload(true)
};`
