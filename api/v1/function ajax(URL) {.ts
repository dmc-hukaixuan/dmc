function ajax(URL) {
    return new Promise(function(resolve, reject){
        var req = new XMLHttpRequest();
        req.open('GET', URL, true);
        req.onload = function() {
            if ( req.status === 200 ) {
                resolve(req.responseText)
            } else {
                reject(new Error(req.statusText))
            }
        }
        req.onerror = function() {
            reject(new Error(req.statusText))
        };
        req.send;
    });
    var url = "";
    ajax(url).then(function onFullfiled(value ){

    }).catch(function onRejected(error){
        document.write("cuowu :" + error)
    })
