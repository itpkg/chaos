
export function parseUrl() {
    var query = window.location.search.substring(1);
    var vars = query.split("&");
    var params = {};
    for (var i=0;i < vars.length;i++) {
      var pair = vars[i].split("=");
      if(pair.length === 2){
        params[pair[0]]=pair[1];
      }
    }
    return params;
}

export function parseXhr(xhr){
  //console.log(xhr);
  return xhr.errors.map(function(e){
     return e.detail;
   });
}
