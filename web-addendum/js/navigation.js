$("#stationrescanwhole").on('click', function(){

    var urlStr = '/station/act/rescanwhole';

    DoAjax(urlStr, {}, 500);
})

$("#stationreboot").on('click', function(){

    var urlStr = '/station/act/reboot';
    
    DoAjax(urlStr, {}, 500);
})

$("#stationexit").on('click', function(){

    var urlStr = '/station/act/exit';
    
    DoAjax(urlStr, {}, 500);
})

$("#stationrestart").on('click', function(){

    var urlStr = '/station/act/restart';
    
    DoAjax(urlStr, {}, 500);
})

$("#stationshutdown").on('click', function(){

    var urlStr = '/station/act/shutdown';
    
    DoAjax(urlStr, {}, 500);
})
