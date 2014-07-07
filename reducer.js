#!/usr/bin/env node

var split = require('split')

process.stdin.resume()

var path_hash = {}
process.stdin
.pipe(split())
.on('data', function (d) {
  d = d.toString()
  path_hash[d+''] = Number(path_hash[d+''] ? path_hash[d+'']+1 : 1)
})
.on('end', function () {
  var ok = Object.keys(path_hash) 
  if (ok.length > 0) {
    ok.sort(function (a,b) {
      return path_hash[a] < path_hash[b] ? 1 : -1
    }).slice(0,20).forEach(function (o) {
      process.stdout.write(path_hash[o] +  " " + o + "\n")
    })
  }
})
