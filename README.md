IT-PACKAGE
---

## Install
    pacman -S jdk8-openjdk gradle
    archlinux-java set java-8-openjdk

## Build
    gradle projects
    gradle tasks
    gradle build
    cd app && java -jar build/libs/itpkg-*.jar
    
### Run
    cd app && gradle bootRun
