<!DOCTYPE html>
<html lang="en">
  <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>OCT</title>
  <link href="/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">
  <link href="/font-awesome/css/font-awesome.css" rel="stylesheet">
  <link href="/css/style.css" rel="stylesheet" type="text/css" media="all">
  <link rel="shortcut icon" href="/img/logo/favicon.ico" type="image/x-icon">
  </head>
  <body>
    <div>
      <nav class="navbar navbar-default navbar-fixed-top">
        <div class="container">
          <div class="navbar-text">
              <a class="navbar-link link-logo" href="/"></a>
          </div>
	  <p class="navbar-text navbar-right">
	  <a href="/"><i class="fa fa-home fa-1x"> Home</i></a>
	  <span class="separator"></span>
	  <a href="/news"><i class="fa fa-bolt fa-1x"> News</i></a>
	  <span class="separator"></span>
<!--	  <span class="separator"></span>
	  <a href="/cases"><i class="fa fa-database fa-1x"> Cases</i></a>
-->
	  <a href="https://github.com/huawei-openlab/oct/"><i class="fa fa-github fa-1x"> Github</i></a>
	  <span class="separator"></span>
	  <a href="https://github.com/huawei-openlab/oct/blob/master/README.md"><i class="fa fa-file-text-o fa-1x"> Docs</i></a>
	  </p>
        </div>
      </nav>

      <div class="head-empty">
      </div>

      <div class="container">
       {{range .News}}
   	  <h2>
	    {{ .Title }}
	  </h2>
	  <div class="post-meta">
	    <span class="post-date">
	      {{ .Date  }}
	    </span>
	    <a href="https://github.com/{{ .Author }}" class="post-author">
	      <img src="https://github.com/{{ .Author }}.png" class="avatar" alt="{{ .Author }} avatar" width="24" height="24">
	      {{ .Author }}
	    </a>
	  </div>
	  <div class="post-content">
	    {{ .Content }}
	  </div>
       {{else}}<div><strong>No news is a bad news.</strong></div>
       {{end}}
      </div>

    </div>

    <div class="footer">
	    <div class="separator-hr"></div>
	    <div class="row">
		    <div class="col-md-6">
			    <span class="separator"></span>
			    Join OCT
			    <a href="https://github.com/huawei-openlab/oct/" class="btn btn-default btn-github"><i class="fa fa-github fa-1x"> </i></a>
			    <a href="https://github.com/huawei-openlab/oct/blob/master/README.md" class="btn btn-default btn-docs"><i class="fa fa-file-text-o fa-1x"></i> </a>
		    </div>
		    <div class="col-md-6">
			    <p class="copyright">Copyright<i class="fa fa-copyright"></i> 2015 Huawei Openlab. All Rights Reserved.</p>
		    </div>
	    </div>
    </div>
  </body>
</html>
