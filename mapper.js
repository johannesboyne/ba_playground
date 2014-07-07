#!/usr/bin/env node

var split = require('split')

process.stdin.setEncoding('utf8')
process.stdin.resume()

process.stdin
.pipe(split())
.on('data', function (d) {
  d = d.toString()
  var slices = d.split(" ")
  if (slices.length >= 6)
    process.stdout.write(slices[6] + "\n") 
})
