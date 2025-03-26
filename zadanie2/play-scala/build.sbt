name := """my-play-scala-app"""
organization := "myFirstScalaApp"

version := "1.0"

lazy val root = (project in file(".")).enablePlugins(PlayScala)

scalaVersion := "2.13.16"

libraryDependencies += guice
libraryDependencies += "org.scalatestplus.play" %% "scalatestplus-play" % "7.0.1" % Test
libraryDependencies += "com.typesafe.play" %% "play-json" % "2.10.0-RC5"


// Adds additional packages into Twirl
//TwirlKeys.templateImports += "myFirstScalaApp.controllers._"

// Adds additional packages into conf/routes
// play.sbt.routes.RoutesKeys.routesImport += "myFirstScalaApp.binders._"
