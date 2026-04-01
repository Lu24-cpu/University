#!/bin/bash
# Gradle Cache Management Script

case "$1" in
    "config")
        echo "Cleaning configuration cache only..."
        ./gradlew --no-configuration-cache clean build
        ;;
    "build")
        echo "Cleaning build cache..."
        ./gradlew cleanBuildCache
        ;;
    "deps")
        echo "Refreshing dependencies..."
        ./gradlew --refresh-dependencies build
        ;;
    "all")
        echo "Complete clean rebuild..."
        ./gradlew --stop
        rm -rf .gradle/ build/
        ./gradlew build
        ;;
    "nuclear")
        echo "Nuclear option - cleaning all Gradle caches..."
        ./gradlew --stop
        rm -rf ~/.gradle/caches/ ~/.gradle/daemon/ .gradle/ build/
        ./gradlew build
        ;;
    *)
        echo "Usage: $0 {config|build|deps|all|nuclear}"
        echo ""
        echo "config  - Clean configuration cache only"
        echo "build   - Clean build cache"
        echo "deps    - Refresh dependencies"
        echo "all     - Complete project clean rebuild"
        echo "nuclear - Clean all Gradle caches (global)"
        ;;
esac