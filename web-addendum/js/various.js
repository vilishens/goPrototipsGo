var BIT_STATE_ACTIVE = 0x0001
var BIT_STATE_FREEZE = 0x0002

function SetInterv(nbr, todo, interv) {
    if(0 > nbr) {
        nbr = setInterval(todo, interv); 
    }    

    return nbr;
}

function UnsetInterv(nbr) {
    if(0 <= nbr) {
        clearInterval(nbr);
        nbr = -1;
    }    

    return nbr;
}

function IsPointFrozen(bit) {
    return (0 != (bit & BIT_STATE_FREEZE));
}

function IsPointActive(bit) {
    return (0 != (bit & BIT_STATE_ACTIVE));
}
