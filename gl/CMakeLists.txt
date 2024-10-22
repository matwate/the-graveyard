cmake_minimum_required(VERSION 3.14)
project(gl/ C GO)
set(CMAKE_GO_COMPILER "C:\\Program Files\\Go\\bin")
set(CMAKE_GO_COMPILER_FORCED ON)
find_package(assimp CONFIG REQUIRED)

add_executable(assimp main.go)
target_link_libraries(assimp PRIVATE ${ASSIMP_LIBRARIES})
target_include_directories(assimp PRIVATE ${ASSIMP_INCLUDE_DIRS})