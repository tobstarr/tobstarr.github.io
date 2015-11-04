<!DOCTYPE html>
<html>
<head>
    <title>tobstarr.com</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/bootstrap.min.css" type="text/css" media="screen" title="no title" charset="utf-8">
    <link rel="stylesheet" href="/css/default.css" type="text/css" media="screen" title="no title" charset="utf-8">
</head>
<body>
  <div class="container">
    <div class="header index_link">
      <a href="/index.html">tobstarr.com</a>
    </div>
    <div class="header_links">
      <a href="/dotfiles.html">Dotfiles</a>
      |
      <a href="/cheats.html">Cheats</a>
    </div>
    <hr />
<p>#!/bin/zsh</p>

<p>SLEEP=${SLEEP:-5}</p>

<p>while true; do
  res=$($@)
  clear
  date
  echo $res
  sleep $SLEEP
done</p>
</div>
</body>
</html>
