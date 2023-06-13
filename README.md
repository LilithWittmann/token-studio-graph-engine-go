# token-studio-graph-engine-go
Hacky WIP implementation of the [Token Studio Graph Engine ](https://github.com/tokens-studio/graph-engine) in GO which compiles with gomobile to android and iOS.

## Usage Android
Download Build artefact (or eventually a release 🤪) and add the AAR/JAR to your projects libs folder. Configure your build.gradle.
```
dependencies {
    implementation (name: 'token_studio_graph_engine', ext: 'aar')
}
```
Execute a graph: (Yes currently the result object is a json object 🙃)
```kotlin
var input = applicationContext.assets.open("graphs/math.json")
var g = Graph(input.readBytes())

var result = g.executeToJSON()
var resultString = result.toString(Charsets.UTF_8)
```

