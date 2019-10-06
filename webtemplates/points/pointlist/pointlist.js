var URL_POINT_LIST="/pointlist/data";
var POINT_LIST_OBJ = 'pointList';
var ITEM_ID_PREFIX = 'ptItem';
var ITEM_BUTTON_ID_PREFIX = 'ptBtnItem';

var ITEM_CLASS_DEFAULT = 'btn-outline-secondary';
var ITEM_CLASS_SIGNED = 'btn-outline-success';
var ITEM_CLASS_DISCONNECTED = 'btn-outline-secondary button-blink';
var ITEM_CLASS_FROZEN = 'btn-outline-danger button-blink';

var URL_LIST_ACTION_CONFIG = 'pointlist/act/cfg/';

var allD = {};
var pointStates = {};
var pointStatus = {};

function makeList() {
    handlePointList()
    var nbr = SetInterv(-5, "handlePointList()", 1000);   // 1 sec
}

function handlePointList() {
 
    allD = {};

    $.ajax({
        url: URL_POINT_LIST,
        type: 'post',
        data: allD, //JSON.stringify(d), 
        dataType: 'json',
        contentType: 'application/json;charset=utf-8',
        async: true,
        timeout: 500,   // 0.5 second
        success : function(data, status, xhr) {
            allD = data;
            drawPointList();
        },
//        error : function(request,error) {
//            alert("Error: "+error);
//        },
    });
}

function drawPointList() {

    removeListItems();

    var wasName = "";

    for (ind in allD["List"]) {
        var name = allD["List"][ind];

        var isNew = newItem(name);
        var isChanged = !isNew && changedItem(name);

        if(isChanged) {
            var kor = 3;
        }
 
        var strHtml = "";
        if(isNew || isChanged) {
            strHtml = itemDataSpanHTML(name);

            if(isNew) {
                var strSpan = '<span id="'+listItemId(name)+'">' + strHtml + '</span>';

                if("" == wasName) {
                    $('#'+POINT_LIST_OBJ).prepend(strSpan);
                } else {
                    var spanItem = $('#'+POINT_LIST_OBJ).find('#' + listItemId(wasName));
                    $(strSpan).insertAfter(spanItem);
                }
            } else {
                $('#' + listItemId(name)).html(strHtml);
            }
        }

//        pointStates[name] = allD["Data"][name]["State"]

        saveStatus(allD["Data"]);


//        var vato = !(pointStates[name] == thisState );

//        return  !(pointStates[name] == thisState );


        wasName = name;
    }
}

function saveStatus(d) {

    pointStatus = {};

    for(name in d) {
        pointStatus[name] = {};
    }    

    for(name in d) {
        pointStatus[name]["Point"] = d[name]["Point"];
        pointStatus[name]["Disconnected"] = d[name]["Disconnected"];
        pointStatus[name]["Frozen"] = d[name]["Frozen"];
        pointStatus[name]["Signed"] = d[name]["Signed"];
        pointStatus[name]["State"] = d[name]["State"];
        pointStatus[name]["Type"] = d[name]["Type"];
    }
}

function removeListItems() {
    var obj = $('#'+POINT_LIST_OBJ);
    var search = '[id^="'+ITEM_ID_PREFIX+'"]';
    var items = obj.find(search);

    items.each(function(){

        var id = $(this).attr('id');
        var name = id.substr(ITEM_ID_PREFIX.length)

        var found = false;
        for (ind in allD["List"]) {
            it = allD["List"][ind];

            if(it == name) {
                found = true;
                break;
            }
        }

        if (!found) {
            // this item is not in the list of data
            change = true;
            $(this).remove();
        } 
    });
}

function itemDataSpanHTML(name) {

    var d = allD["Data"][name];
    var cl = itemDataClass(name);
    var str = '';

//    str += '<span id="'+listItemId(name)+'">'
    str += '<div class="container">';
  	str += '    <div class="row pointListItemDiv">';
    str += '        <div class="dropright">';
    str += '            <button class="btn dropdown-toggle '+cl+' pointListItem" type="button" id="'+listItemBtnId(name)+'" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">';
    str += '                '+d["Point"];
    str += '            </button>';
    str += '            <ul class="dropdown-menu multi-level" role="menu" aria-labelledby="dropdownMenu">';



//    str += '                <li class="dropdown-item"><a href="#">Some action</a></li>';
//    str += '                <li class="dropdown-item"><a href="#">Some other action</a></li>';
//    str += '                <li class="dropdown-divider"></li>';
    str += '                <li class="dropdown-submenu">';
    str += '                    <a  class="dropdown-item" tabindex="-1" href="#">Configuration</a>';
    str += '                    <ul class="dropdown-menu">';
    str += configChoices(name);

    /*
                                if (name in allD["Data"]) {

                                    var tData = allD["Data"][name];


                                    for(k in tData["CfgList"]) {
                                        var typ = tData["CfgList"][k];

                                        var tName = tData["CfgInfo"][typ]["Name"];
                                        str += '    <li class="dropdown-item"><a tabindex="-1" href="pointlist/act/cfg/'+name+'/'+Number(typ)+'">' + tName +'</a></li>';






                                    }
                                }

*/
    str += '                        <li class="dropdown-item"><a tabindex="-1" href="#">Second level</a></li>';
    str += '                        <li class="dropdown-submenu">';
    str += '                            <a class="dropdown-item" href="#">Even More..</a>';
    str += '                            <ul class="dropdown-menu">';
    str += '                                <li class="dropdown-item"><a href="#">3rd level</a></li>';
    str += '                                <li class="dropdown-submenu"><a class="dropdown-item" href="#">another level</a>';
    str += '                                    <ul class="dropdown-menu">';
    str += '                                        <li class="dropdown-item"><a href="#">4th level</a></li>';
    str += '                                        <li class="dropdown-item"><a href="#">4th level</a></li>';
    str += '                                        <li class="dropdown-item"><a href="#">4th level</a></li>';
    str += '                                    </ul>';
    str += '                                </li>';
    str += '                                <li class="dropdown-item"><a href="#">3rd level</a></li>';
    str += '                            </ul>';
    str += '                        </li>';
    str += '                        <li class="dropdown-item"><a href="#">Second level</a></li>';
    str += '                        <li class="dropdown-item"><a href="#">Second level</a></li>';
    str += '                    </ul>';
    str += '                </li>';
    if(isDisconnected(name)) {
    str += '                <li class="dropdown-item"><a href="pointlist/act/rescan/'+name+'">Rescan</a></li>';
//    str += '                <li class="dropdown-item"><a href="#">Some other action</a></li>';
    }
    str += '            </ul>';           
    str += '        </div>';
    str += '    </div>';
    str += '</div>';
//  str += '</span>'

    return str;

}

function newItem(name) {
    return 0 == itemObject(name).length; 
}

function itemObject(name) {
    return $('#'+POINT_LIST_OBJ).find("#"+listItemId(name));
}

function changedItem(name) {

    if(name in pointStatus) {
        var now = allD["Data"][name]["State"];
        var was = pointStatus[name]; 

        for(fld in now) {
            if (now[fld] != was[fld]) {
                return true;
            }
        }
    }

    return false;
}

function itemObjectButton(name) {
    var obj = $('#'+POINT_LIST_OBJ).find("#"+listItemId(name));
    var btn = obj.find('#'+listItemBtnId(name));

    return btn;
}

function listItemId(name) {
    return prefixNameId(ITEM_ID_PREFIX,name);
}

function listItemBtnId(name) {
    return prefixNameId(ITEM_BUTTON_ID_PREFIX,name);
}

function prefixNameId(prefix, name) {
    return prefix + name;
}

function itemDataClass(name) {

    var cl = ITEM_CLASS_DEFAULT;

    if(isDisconnected(name)) {
        cl = ITEM_CLASS_DISCONNECTED;
    } else if(isFrozen(name)) {
        cl = ITEM_CLASS_FROZEN;
    } 
       else if(isSigned(name)) {
        cl = ITEM_CLASS_SIGNED;
    }

    return cl;
}

function isDisconnected(name) {
    var item = allD["Data"][name];

    return (item["Signed"] && item["Disconnected"]);
}

function isFrozen(name) {
    var item = allD["Data"][name];

    return (item["Signed"] && item["Frozen"]);
}

function isSigned(name) {
    var item = allD["Data"][name];

    return (item["Signed"] && !item["Disconnected"]);
}

function configChoices(name) {
    var str = "";

    if (name in allD["Data"]) {

        var objCfgCds = allD["Data"][name]["CfgList"];
        var objCfgItems = allD["Data"][name]["CfgInfo"];


        for(k in objCfgCds) {
            var cfgType = objCfgCds[k];

            var itemName = objCfgItems[cfgType]["Name"].trim();
            if(0 == itemName.length) {
                return "";
            }
            itemName = itemName.charAt(0).toUpperCase() + itemName.slice(1);
            
            str += '<li class="dropdown-item">';
            str += '    <a tabindex="-1" href="' +URL_LIST_ACTION_CONFIG+name+'/'+Number(cfgType)+'">' + itemName +'</a>';
            str += '</li>';       
        }    
    }

    return str;
}

//#############################################################
//#############################################################
//#############################################################

/*
function hasMyClasses(obj, cl) {

    var arr = cl.split(" ");

    for(ind in arr) {
        if(!obj.hasClass(arr[ind])) {
            return false;
        }
    }

    return true;
}
*/

function bootstrapa_menu(d, cl, name) {
/*    
    <div class="container">
  	<div class="row">
        <h2>Multi level dropdown menu in Bootstrap</h2>
        <hr>
        <div class="dropdown">
            <button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
              Dropdown
            </button>
            <ul class="dropdown-menu multi-level" role="menu" aria-labelledby="dropdownMenu">
                <li class="dropdown-item"><a href="#">Some action</a></li>
                <li class="dropdown-item"><a href="#">Some other action</a></li>
                <li class="dropdown-divider"></li>
                <li class="dropdown-submenu">
                  <a  class="dropdown-item" tabindex="-1" href="#">Hover me for more options</a>
                  <ul class="dropdown-menu">
                    <li class="dropdown-item"><a tabindex="-1" href="#">Second level</a></li>
                    <li class="dropdown-submenu">
                      <a class="dropdown-item" href="#">Even More..</a>
                      <ul class="dropdown-menu">
                          <li class="dropdown-item"><a href="#">3rd level</a></li>
                            <li class="dropdown-submenu"><a class="dropdown-item" href="#">another level</a>
                            <ul class="dropdown-menu">
                                <li class="dropdown-item"><a href="#">4th level</a></li>
                                <li class="dropdown-item"><a href="#">4th level</a></li>
                                <li class="dropdown-item"><a href="#">4th level</a></li>
                            </ul>
                          </li>
                            <li class="dropdown-item"><a href="#">3rd level</a></li>
                      </ul>
                    </li>
                    <li class="dropdown-item"><a href="#">Second level</a></li>
                    <li class="dropdown-item"><a href="#">Second level</a></li>
                  </ul>
                </li>
              </ul>
        </div>
    </div>
</div>
*/
}
