package com.example

class GatewayEvent( val op: Int?, val d: String,val s: Int?,val t: String?) {

    override fun toString(): String {
        return "GatewayEvent(op=$op, d='$d', s=$s, t=$t)"
    }

}
