---
title: "Image variants"
layout: default
sort: 1
---

# Image variants
{: .no_toc}

## Table of contents
{: .no_toc .text-delta}

1. TOC
{:toc}

---

NFD currently offers two variants of the container image. The "full" variant is
currently deployed by default. Released container images are available for
x86_64 and Arm64 architectures.

## Full

This image is based on [debian:bullseye-slim](https://hub.docker.com/_/debian)
and contains a full Linux system for running shell-based nfd-worker hooks and
doing live debugging and diagnosis of the NFD images.

## Minimal

This is a minimal image based on
[gcr.io/distroless/base](https://github.com/GoogleContainerTools/distroless/blob/master/base/README.md)
and only supports running statically linked binaries.

The container image tag has suffix `-minimal`
(e.g. `{{ site.container_image }}-minimal`)
