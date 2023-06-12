# token-studio-graph-engine-go
Hacky WIP implementation of the [Token Studio Graph Engine ](https://github.com/tokens-studio/graph-engine) in GO which compiles with gomobile to android and iOS.

## Usage Android
Download Build artefact (or eventually a release 🤪) and add the AAR/JAR to your projects libs folder. Configure your build.gradle.
```
dependencies {
    implementation (name: 'token_studio_graph_engine', ext: 'aar')
}
```
