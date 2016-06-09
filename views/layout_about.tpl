<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Optiroute</title>
    {{ .HeadScripts }}
    <link rel="stylesheet" href="https://code.getmdl.io/1.1.3/material.grey-red.min.css" />
    <link rel="stylesheet" href="/static/css/about.css"/>
</head>
<body>
<div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
    <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">
            <!-- Title -->
            <span class="mdl-layout-title">
                 <a class="mdl-navigation__link" href="/">
                <img id="logo" src="/static/img/logo.png" alt="Optiroute"/>
                </a>
            </span>
            <!-- Add spacer, to align navigation to the right -->
            <div class="mdl-layout-spacer"></div>
            <!-- Navigation. We hide it in small screens. -->
            <nav class="mdl-navigation mdl-layout--large-screen-only">
                <a class="mdl-navigation__link" href="/">Home</a>
            </nav>
        </div>
    </header>

    <main class="mdl-layout__content">
        <div class="page-content">
            {{ .LayoutContent }}
        </div>
    </main>
</div>
{{ .Footer }}
<script src="/static/js/AboutHelpers.js"></script>
</body>
</html>