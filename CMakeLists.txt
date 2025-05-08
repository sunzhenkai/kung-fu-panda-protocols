cmake_minimum_required(VERSION 3.20)
project(kung-fu-panda-protocols C CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_FLAGS "${CMAKE_CPP_FLAGS} -O2 -fPIC -std=c++17")
option(BUILD_TESTING "build unittests" OFF)
option(GEN_GO "build unittests" OFF)
include(GNUInstallDirs)
if(BUILD_TESTING)
  set(CMAKE_INSTALL_PREFIX ${CMAKE_SOURCE_DIR}/build/install/)
endif()

find_package(protobuf CONFIG REQUIRED)

set(THIRDPARTY_LIBRARIES protobuf::libprotobuf protobuf::libprotobuf-lite)

# generate protocols
include(3rd/vcpkg-base/cmakes/tools.cmake)
include(3rd/vcpkg-base/cmakes/protobufv2.cmake)
file(GLOB_RECURSE PB_SOURCE_FILES "protos/*.proto")

generate_protobuf_message(
  SRCS
  PB_SRCS
  HEADERS
  PB_HEADERS
  OUTPUT
  ${PROJECT_BINARY_DIR}/gen/protobuf/cpp
  IMPORT
  ${PROJECT_SOURCE_DIR}/
  FILES
  ${PB_SOURCE_FILES})

if(GEN_GO)
  generate_protobuf_message_go(
    GEN_GRPC
    OUTPUT
    ${PROJECT_BINARY_DIR}/gen/protobuf/go
    IMPORT
    ${PROJECT_SOURCE_DIR}/
    FILES
    ${PB_SOURCE_FILES}
    GO_OPT_MODULE
    "github.com/sunzhenkai/kung-fu-panda-protocols/go")
endif()

add_library(${PROJECT_NAME} ${PB_SRCS})
target_link_libraries(${PROJECT_NAME} ${THIRDPARTY_LIBRARIES})
target_include_directories(${PROJECT_NAME} PRIVATE ${PB_HEADERS})

install(
  TARGETS ${PROJECT_NAME}
  EXPORT ${PROJECT_NAME}
  RUNTIME DESTINATION "${CMAKE_INSTALL_BINDIR}"
  ARCHIVE DESTINATION "${CMAKE_INSTALL_LIBDIR}"
  LIBRARY DESTINATION "${CMAKE_INSTALL_LIBDIR}"
  INCLUDES
  DESTINATION "${CMAKE_INSTALL_INCLUDEDIR}")

foreach(hd IN LISTS PB_HEADERS)
  install(
    DIRECTORY ${hd}/
    DESTINATION include
    FILES_MATCHING
    PATTERN "*.h")
endforeach()
install(
  EXPORT ${PROJECT_NAME}
  DESTINATION share/${PROJECT_NAME}
  FILE ${PROJECT_NAME}Config.cmake
  NAMESPACE ${PROJECT_NAME}::)
