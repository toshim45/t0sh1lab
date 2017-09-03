name := """learn-play-rest-jv"""

version := "1.0-SNAPSHOT"

lazy val root = (project in file(".")).enablePlugins(PlayJava)

scalaVersion := "2.11.7"

libraryDependencies ++= Seq(
  javaJpa,
  "org.postgresql" % "postgresql" % "9.4.1209",
  "org.hibernate" % "hibernate-entitymanager" % "5.2.1.Final",
  "dom4j" % "dom4j" % "1.6.1", // https://stackoverflow.com/questions/38278199/play-framework-inject-error
  cache,
  javaWs,
  filters,
  evolutions
)
