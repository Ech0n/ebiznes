package com.example

import io.ktor.client.*
import io.ktor.client.call.body
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.websocket.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.websocket.*
import kotlinx.coroutines.*
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import io.github.cdimascio.dotenv.Dotenv

val dotenv = Dotenv.load()
val USER_ID = dotenv.get("DISCORDUSERID")
val DISCORD_BOT_ID = dotenv.get("DISCORDBOTID")
val DISCORD_BOT_TOKEN = dotenv.get("DISCORDBOTTOKEN")
var heartbeatInterval = 40000
var isIdentified = false
var isReady = false

fun main() = runBlocking {
    if(DISCORD_BOT_TOKEN == null || USER_ID == null  || DISCORD_BOT_ID == null ){
        error("\n--------------------------------------------\nNot all env variables set. Check README.md for instructions\n----------------------------------------")
    }
    val clientForHttpRequests = HttpClient(CIO)

    val gatewayRequestResponse = getGateway(clientForHttpRequests)
    if(gatewayRequestResponse.status !=  HttpStatusCode.OK) {
        println("Error getting gateway: ${gatewayRequestResponse.status}")
        return@runBlocking
    }
    val channelId = makeDmChannel(clientForHttpRequests)
    sendMessageToClient(
        clientForHttpRequests,
        "Hey!",
        channelId
    )
    val bodyAsString: String = gatewayRequestResponse.body<String>()
    val jsonElement = Json.parseToJsonElement(bodyAsString)
    val url: String =
        jsonElement.jsonObject["url"]?.jsonPrimitive?.content ?: error("URL is missing or not a string")
    println("Connecting to gateway $url")

    runGatewayLoop(url, clientForHttpRequests)

    clientForHttpRequests.close()
}
suspend fun getGateway(client: HttpClient): io.ktor.client.statement.HttpResponse {
    val url = "https://discord.com/api/gateway/bot"
    val response = client.get(url) {
        header(HttpHeaders.Authorization, "Bot $DISCORD_BOT_TOKEN")
    }
    return response
}

suspend fun makeDmChannel(client: HttpClient): String {
    val url = "https://discord.com/api/v10/users/@me/channels"

    val response = client.post(url) {
        header(HttpHeaders.Authorization, "Bot $DISCORD_BOT_TOKEN")
        contentType(ContentType.Application.Json)
        setBody("""{"recipient_id":"$USER_ID"}""")
    }
//    println("response: ${response.status}")
//    println("response: ${response.body<String>()}")
    val jsonResponse = Json.parseToJsonElement(response.body<String>())
    val channelId = jsonResponse.jsonObject["id"]!!.jsonPrimitive.content
    return channelId
}

suspend fun sendMessageToClient(client: HttpClient, messageToClient: String, channelId : String) {

    val url = "https://discord.com/api/v10/channels/$channelId/messages"

    val response = client.post(url) {
        header(HttpHeaders.Authorization, "Bot $DISCORD_BOT_TOKEN")
        contentType(ContentType.Application.Json)
        setBody("""{"content":"$messageToClient"}""")
    }
//    println("response: ${response.status}")
//    println("response: ${response.body<String>()}")
}

fun runGatewayLoop(url: String, clientForHttpRequests: HttpClient)   {
    val client = HttpClient(CIO)
    {
        install(WebSockets) {
            pingIntervalMillis = 20_000
        }
    }

    runBlocking {
        client.webSocket(url) {
            val session = this
            while (true) {
                val othersMessage = incoming.receive() as? Frame.Text
                val messageText = othersMessage?.readText()
                if (messageText != null) {
                    val ge = parseMessage(messageText)
                    when (ge.op) {
                        10 -> {
                            handleHello(ge, session)
                        }

                        11 -> {
                            handleHeartbeatACK()
                        }
                        0 ->{
                            when(ge.t){
                                """"READY"""" -> {
                                    println("Gateway is READY")
                                    isReady = true
                                }
                                """"MESSAGE_CREATE""""-> handleMessageCreated(ge, clientForHttpRequests)
                                else -> println("Unhandled event: ${ge.t}")
                            }
                        }
                        else -> println("Unhandled OP code: ${ge.op}")
                    }
                }
            }
        }
    }
    client.close()
}

private fun DefaultClientWebSocketSession.handleHello(
    ge: GatewayEvent,
    session: DefaultClientWebSocketSession
) {
    println("Received OP code 10: Hello received from Discord server.")
    heartbeatInterval = parseHello(ge.d)
    launch {
        heartbeat(session)
    }
}

private suspend fun DefaultClientWebSocketSession.handleHeartbeatACK() {
    if (!isIdentified) {
        val message =
            """{"op":2,"d":{"token":"$DISCORD_BOT_TOKEN", "properties":{"os":"windows"},"intents":${1 shl 12}}} """
        send(Frame.Text(message))
//        print("sent $message")
        isIdentified = true
    }
}

private suspend fun handleMessageCreated(
    ge: GatewayEvent,
    clientForHttpRequests: HttpClient
) {
    if (!isReady) {
        println("gateway not ready yet")
        return
    }
    val (content, channelId, senderId) = parseMessageCreateEvent(ge)
    if (senderId == DISCORD_BOT_ID) {
        println("ignoring message sent by bot")
        return
    }
    println("recieved message from discord user: $content")
    if (content != null) {
        val processedContent = content.trim().lowercase()
        if (!processedContent.startsWith("list")) {
            return
        }
        val words = processedContent.split(" ")
        if (words.size < 2) {
            return
        }
        if (words[1] == "categories") {
            val messageToClient = Products.map.keys.joinToString(separator = ", ")
            sendMessageToClient(clientForHttpRequests, messageToClient, channelId)
        }
        if (Products.map.containsKey(words[1])) {
            val messageToClient = Products.map[words[1]]?.joinToString(separator = ", ")
            sendMessageToClient(clientForHttpRequests, messageToClient.toString(), channelId)
        }
    }
}

fun parseHello(d: String): Int {
    val jsonElement = Json.parseToJsonElement(d)
    val heartbeatInterval = jsonElement.jsonObject["heartbeat_interval"]?.jsonPrimitive?.content?.toIntOrNull()
        ?: error("Missing or invalid heartbeat_interval in Hello event")
    println("Parsed heartbeat_interval: $heartbeatInterval")
    return heartbeatInterval
}
fun parseMessage(messageText: String): GatewayEvent {
//    println("parsing: $messageText")
    val jsonElement = Json.parseToJsonElement(messageText)
    val op = jsonElement.jsonObject["op"]?.jsonPrimitive?.content?.toIntOrNull()
    val d = jsonElement.jsonObject["d"].toString()
    val s = jsonElement.jsonObject["s"]?.jsonPrimitive?.content?.toIntOrNull()
    val t = jsonElement.jsonObject["t"]?.toString()

    if(op == null){
        println("PARSING ERROR")
        return GatewayEvent(-1,"",null,null)
    }

//    println("op: $op, d: $d, s: $s, t: $t")
    println("recieved: op $op, t:$t")
    return GatewayEvent(op,d,s,t)
}
private fun parseMessageCreateEvent(ge: GatewayEvent): Triple<String?, String, String> {
    val message = ge.d
    val jsonMessage = Json.parseToJsonElement(message)
    val content = jsonMessage.jsonObject["content"]?.jsonPrimitive?.content
    val channelId = jsonMessage.jsonObject["channel_id"]?.jsonPrimitive?.content.toString()
    val author = jsonMessage.jsonObject["author"].toString()
    val authorJson = Json.parseToJsonElement(author)
    val senderId = authorJson.jsonObject["id"]?.jsonPrimitive?.content.toString()
    return Triple(content, channelId, senderId)
}

suspend fun heartbeat(session: WebSocketSession){
    while (true) {
        val heartbeatPayload = """{"op":1, "d":null}"""
        session.send(Frame.Text(heartbeatPayload))
        println("Sent heartbeat payload: $heartbeatPayload")
        delay(heartbeatInterval.toLong())
    }
}
