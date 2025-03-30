val ktor_version: String by project
val logback_version: String by project

plugins {
    application
//    id("org.jetbrains.kotlin.jvm") version "1.9.21"
    kotlin("jvm") version "2.1.0"
}

application {
    mainClass.set("com.example.ApplicationKt")
}

repositories {
    mavenCentral()
    maven { url = uri("https://maven.pkg.jetbrains.space/public/p/ktor/eap") }
}

tasks.named<JavaExec>("run") {
    standardInput = System.`in`
}

dependencies {
    implementation("io.ktor:ktor-client-core:$ktor_version")
    implementation("io.ktor:ktor-client-cio:$ktor_version")
    implementation("io.ktor:ktor-client-websockets:$ktor_version")
    implementation("io.ktor:ktor-client-logging:$ktor_version")
    implementation("io.ktor:ktor-http-cio:$ktor_version")
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:1.6.0")
    implementation("io.ktor:ktor-http:$ktor_version")
    implementation("io.github.cdimascio:dotenv-kotlin:6.5.1")
    implementation("ch.qos.logback:logback-classic:$logback_version")
}