package server

import (
  "encoding/json"
  "fmt"
	"github.com/gorilla/mux"
  "github.com/Rookii/paicehusk"
	"html/template"
	"net/http"
  "strings"
)

var Templates *template.Template

func init() {
	t, err := template.New("foo").Parse(homePage)
	if err != nil {
		panic(err)
	}
	Templates = t
	router := mux.NewRouter()

	// Root handler
	router.HandleFunc("/", home)
  router.HandleFunc("/stem", stem)
	http.Handle("/", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "T", nil)
}

func stem(w http.ResponseWriter, r *http.Request) {
  out := make(map[string]string)
  decoder := json.NewDecoder(r.Body)
  build := new(CalcBuild)
  
  if err := decoder.Decode(build); err != nil {
    fmt.Println(err)
  }

  for _, word := range in {
    out[word] = paicehusk.DefaultRules.Stem(word)
  }
  for _, stem := range out {
    fmt.Println(stem)
      fmt.Fprintf(w, stem)
  }
}

var homePage = `
{{define "T"}}

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Bootstrap, from Twitter</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!-- Le styles -->
    <link href="/assets/css/bootstrap.css" rel="stylesheet">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
    <style type="text/css">
      body {
        padding-top: 60px;
        padding-bottom: 40px;
      }
    </style>
    <link href="/assets/css/bootstrap-responsive.css" rel="stylesheet">
    <script>
        $(document).ready(function() {
            $('#stemButton').click(function() {
                stem();
              })
          })

        function stem(){
            var text = $("#input").val();
            var input = {input: text}
            $.ajax({
                type: "POST",
                url: "/stem",
                data: JSON.stringify(input)
            }).done(function(data) {
                $("#output").val(data);
            });
        };
    </script>
  </head>

  <body>

    <div class="navbar navbar-inverse navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container">
          <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </a>
          <a class="brand" href="#">Paice/Husk Stemmer</a>
        </div>
      </div>
    </div>

    <div class="container">
      <!-- Example row of columns -->
      <div class="row">
        <div class="span6">
          <h2>Text to Stem</h2>
          <form>
          <textarea rows="3" id="input"></textarea>
          </form>
          <button id="stemButton" class="btn">Submit</button>
        </div>
        <div class="span6">
          <h2>Stemmed Output</h2>
          <textarea rows="3" id="output"></textarea>
       </div>
      </div>

      <hr>

      <footer>
        <p>&copy; Oh hai 2012</p>
      </footer>

    </div> <!-- /container -->

    <!-- Le javascript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->


  </body>
</html>

{{end}}
`
