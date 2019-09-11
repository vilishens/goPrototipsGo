var AJAX_LET_TIMEOUT = 500; // 0.5 sec

// send the handle request to the server and don't get data back
function LetAjax(urlStr, timeOut) {
    var data = {}
    $.ajax({  
        url: urlStr,
        type: 'post',
        data: data, 
        dataType: 'json',
        contentType: 'application/json;charset=utf-8',
        async: true,
        timeout: timeOut,   // 500 == 0.5 second
        success : function(data, status, xhr) {
            return true;
        },
        error : function(request,error) {
            return false;

//            alert("Request: "+JSON.stringify(request)+", Error: "+error);
        },
    });
}

function DoAjax(urlStr, data, timeOut) {
    $.ajax({  
        url: urlStr,
        type: 'post',
        data: data, 
        dataType: 'json',
        contentType: 'application/json;charset=utf-8',
        async: true,
        timeout: timeOut,   // 500 == 0.5 second
        success : function(data, status, xhr) {
            return;
        },
        error : function(request,error) {
            alert("Request: "+JSON.stringify(request)+", Error: "+error);
        },
    });
}

function ReturnData(url, d) {
    $.ajax({  
        url: url,
        type: 'post',
        data: JSON.stringify(d), 
        dataType: 'json',
        contentType: 'application/json;charset=utf-8',
        async: true,
        timeout: 500,   // 0.5 second
        success : function(data, status, xhr) {
 //           alert("Data "+ data + " STATUS " + status + " XHR " +xhr);
            return;
        },
        error : function(request,error) {
            alert("Request: "+JSON.stringify(request)+", Error: "+error);
        },
    });
}