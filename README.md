# mcdf [![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs] [![Travis](https://img.shields.io/travis/mahosi1/mcdf.svg?style=flat-square)][travis] [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[godocs]: http://godoc.org/github.com/mahosi1/mcdf
[travis]: https://travis-ci.org/mahosi1/mcdf
[license]: /LICENSE

`mcdf` is a Go library to support the reading, writing and updating of Microsoft Compound File Binary Format files. 

It supports read/write operations on streams and storages and traversal of structures tree.

Features
- COM Structured Storage
- Read / Write OLE compound files
- No external dependencies
- Version 4 of the Compound File Binary Format is supported

Read and manipulate files 
- Microsoft Word Document (DOC)
- Microsoft Excel Spreadsheet (XLS)
- Microsoft Powerpoint Presentation (PPT)
- Microsoft Outlook (MSG)
- MSN (Toolbar) (C:\Documents and Settings\%USERNAME%\Local Settings\Application Data\Microsoft\MSNe\msninfo.dat)
- Jump Lists
- StickyNotes.snt
- Thumbs.db
- Windows Installer (.msi) and patch file (.msp)
- Windows Search (srchadm.msc)